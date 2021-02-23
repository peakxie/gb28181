%%{# -*-ragel-*-

machine base;



action mark {
	mark = p
}

action start {
	amt = 0
}

action append {
	buf[amt] = fc
	amt++
}

action hexHi {
	hex = unhex(fc) * 16
}

action hexLo {
	hex += unhex(fc)
	buf[amt] = hex
	amt++
}

action name {
	name = string(data[mark:p])
}

action addparam {
	params.M.Store(name, string(buf[0:amt]))	
}


SP              = " ";
HTAB            = "\t";
CR              = "\r";
LF              = "\n";
DQUOTE          = "\"";
CRLF            = CR LF;
WSP             = SP | HTAB;
LWS             = ( WSP* CRLF )? WSP+;
SWS             = LWS?;

LWSCRLF_append  =  CR @append LF @append;
LWS_append      = ( WSP* @append LWSCRLF_append )? WSP+ @append;

UTF8_CONT       = 0x80..0xBF @append;
UTF8_NONASCII   = 0xC0..0xDF @append UTF8_CONT {1}
                | 0xE0..0xEF @append UTF8_CONT {2}
                | 0xF0..0xF7 @append UTF8_CONT {3}
                | 0xF8..0xFb @append UTF8_CONT {4}
                | 0xFC..0xFD @append UTF8_CONT {5};
UTF8            = 0x21..0x7F @append | UTF8_NONASCII;

mUTF8_CONT      = 0x80..0xBF;
mUTF8_NONASCII  = 0xC0..0xDF mUTF8_CONT {1}
                | 0xE0..0xEF mUTF8_CONT {2}
                | 0xF0..0xF7 mUTF8_CONT {3}
                | 0xF8..0xFb mUTF8_CONT {4}
                | 0xFC..0xFD mUTF8_CONT {5};
mUTF8           = 0x21..0x7F | mUTF8_NONASCII;

# https://tools.ietf.org/html/rfc3261#section-25.1
reserved        = ";" | "/" | "?" | ":" | "@" | "&" | "=" | "+" | "$" | "," ;
mark            = "-" | "_" | "." | "!" | "~" | "*" | "'" | "(" | ")" ;
unreserved      = alnum | mark ;
tokenc          = alnum | "-" | "." | "!" | "%" | "*" | "_" | "+" | "`"
                | "'" | "~" ;
separators      = "("  | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\\"
                | "\"" | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP
                | HTAB ;
wordc           = alnum | "-" | "." | "!" | "%" | "*" | "_" | "+" | "`"
                | "'" | "~" | "(" | ")" | "<" | ">" | ":" | "\\" | "\""
                | "/" | "[" | "]" | "?" | "{" | "}" ;
schmchars       = alnum | "+" | "-" | "." ;
word            = wordc+;
STAR            = SWS "*" SWS;
SLASH           = SWS "/" SWS;
EQUAL           = SWS "=" SWS;
LPAREN          = SWS "(" SWS;
RPAREN          = SWS ")" SWS;
RAQUOT          = ">" SWS;
LAQUOT          = SWS "<";
COMMA           = SWS "," SWS;
SEMI            = SWS ";" SWS;
COLON           = SWS ":" SWS;
HCOLON          = WSP* ":" SWS;
LDQUOT          = SWS "\"";
RDQUOT          = "\"" SWS;
escaped         = "%" ( xdigit @hexHi ) ( xdigit @hexLo ) ;
ipv4c           = digit | "." ;
ipv6c           = xdigit | "." | ":" ;
hostc           = alnum | "-" | "." ;
token           = tokenc+;
tokenhost       = ( tokenc | "[" | "]" | ":" )+;
reasonc         = UTF8_NONASCII | ( reserved | unreserved | SP | HTAB ) @append;
reasonmc        = escaped | reasonc;
cid             = word ( "@" word )?;
hval            = ( mUTF8 | LWS )* >mark;

schemec         = alnum | "+" | "-" | ".";
scheme          = alpha schemec*;
uric            = reserved | unreserved | "%" | "[" | "]";
uri             = scheme ":" uric+;

# Quoted strings can have just about anything, including backslash escapes,
# which aren't quite as fancy as the ones you'd see in programming.
qdtextc         = 0x21 | 0x23..0x5B | 0x5D..0x7E;
qdtext          = UTF8_NONASCII | LWS_append | qdtextc @append;
quoted_pair     = "\\" ( 0x00..0x09 | 0x0B..0x0C | 0x0E..0x7F ) @append;
quoted_content  = ( qdtext | quoted_pair )* >start;
quoted_string   = DQUOTE quoted_content DQUOTE;
unquoted_string = ( token LWS )+;

# Parameter Parsing
#
# Parameters can be used by vias and addresses, but not URIs. They can look
# like=this or like="this". The =value part is optional.
param_name      = token >mark %name;
param_content   = tokenhost @append;
param_value     = param_content | quoted_string;
param           = param_name >start (EQUAL param_value)?;

WWWAuthenticateMessage = SWS "Digest" LWS param %addparam (COMMA param %addparam)*; 

}%%
