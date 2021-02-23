package base 

import (
"errors"
"fmt"
)



var www_authenticate_start  int  = 1
var www_authenticate_first_final  int  = 37
var www_authenticate_error  int  = 0
var www_authenticate_en_main  int  = 1
var _www_authenticate_nfa_targs [] int8  = [] int8  { 0, 0  }
var _www_authenticate_nfa_offsets [] int8  = [] int8  { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0  }
var _www_authenticate_nfa_push_actions [] int8  = [] int8  { 0, 0  }
var _www_authenticate_nfa_pop_trans [] int8  = [] int8  { 0, 0  }
func ParseWwwAuth(data []byte) (params *Params, err error) {
	if data == nil {
		return nil, nil
	}
	params = new(Params)
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	var name string
	
	
	
	{
		cs = int(www_authenticate_start);
	}
	
	{
		if p == pe  {
			goto _test_eof;
			
		}
		switch cs  {
			case 1:
			goto st_case_1;
			case 0:
			goto st_case_0;
			case 2:
			goto st_case_2;
			case 3:
			goto st_case_3;
			case 4:
			goto st_case_4;
			case 5:
			goto st_case_5;
			case 6:
			goto st_case_6;
			case 7:
			goto st_case_7;
			case 8:
			goto st_case_8;
			case 9:
			goto st_case_9;
			case 10:
			goto st_case_10;
			case 11:
			goto st_case_11;
			case 12:
			goto st_case_12;
			case 13:
			goto st_case_13;
			case 14:
			goto st_case_14;
			case 37:
			goto st_case_37;
			case 15:
			goto st_case_15;
			case 16:
			goto st_case_16;
			case 17:
			goto st_case_17;
			case 18:
			goto st_case_18;
			case 19:
			goto st_case_19;
			case 20:
			goto st_case_20;
			case 21:
			goto st_case_21;
			case 22:
			goto st_case_22;
			case 38:
			goto st_case_38;
			case 23:
			goto st_case_23;
			case 24:
			goto st_case_24;
			case 25:
			goto st_case_25;
			case 26:
			goto st_case_26;
			case 27:
			goto st_case_27;
			case 28:
			goto st_case_28;
			case 29:
			goto st_case_29;
			case 30:
			goto st_case_30;
			case 39:
			goto st_case_39;
			case 31:
			goto st_case_31;
			case 32:
			goto st_case_32;
			case 33:
			goto st_case_33;
			case 34:
			goto st_case_34;
			case 35:
			goto st_case_35;
			case 36:
			goto st_case_36;
			
		}
		goto st_out;
		st1:
		p+= 1;
		if p == pe  {
			goto _test_eof1;
			
		}
		st_case_1:
		switch ( data[p ])  {
			case 9:
			{
				goto st1;
			}
			case 13:
			{
				goto st2;
			}
			case 32:
			{
				goto st1;
			}
			case 68:
			{
				goto st5;
			}
			
		}
		{
			goto st0;
		}
		st_case_0:
		st0:
		cs = 0;
		goto _out;
		st2:
		p+= 1;
		if p == pe  {
			goto _test_eof2;
			
		}
		st_case_2:
		if ( data[p ]) == 10  {
			{
				goto st3;
			}
			
		}
		{
			goto st0;
		}
		st3:
		p+= 1;
		if p == pe  {
			goto _test_eof3;
			
		}
		st_case_3:
		switch ( data[p ])  {
			case 9:
			{
				goto st4;
			}
			case 32:
			{
				goto st4;
			}
			
		}
		{
			goto st0;
		}
		st4:
		p+= 1;
		if p == pe  {
			goto _test_eof4;
			
		}
		st_case_4:
		switch ( data[p ])  {
			case 9:
			{
				goto st4;
			}
			case 32:
			{
				goto st4;
			}
			case 68:
			{
				goto st5;
			}
			
		}
		{
			goto st0;
		}
		st5:
		p+= 1;
		if p == pe  {
			goto _test_eof5;
			
		}
		st_case_5:
		if ( data[p ]) == 105  {
			{
				goto st6;
			}
			
		}
		{
			goto st0;
		}
		st6:
		p+= 1;
		if p == pe  {
			goto _test_eof6;
			
		}
		st_case_6:
		if ( data[p ]) == 103  {
			{
				goto st7;
			}
			
		}
		{
			goto st0;
		}
		st7:
		p+= 1;
		if p == pe  {
			goto _test_eof7;
			
		}
		st_case_7:
		if ( data[p ]) == 101  {
			{
				goto st8;
			}
			
		}
		{
			goto st0;
		}
		st8:
		p+= 1;
		if p == pe  {
			goto _test_eof8;
			
		}
		st_case_8:
		if ( data[p ]) == 115  {
			{
				goto st9;
			}
			
		}
		{
			goto st0;
		}
		st9:
		p+= 1;
		if p == pe  {
			goto _test_eof9;
			
		}
		st_case_9:
		if ( data[p ]) == 116  {
			{
				goto st10;
			}
			
		}
		{
			goto st0;
		}
		st10:
		p+= 1;
		if p == pe  {
			goto _test_eof10;
			
		}
		st_case_10:
		switch ( data[p ])  {
			case 9:
			{
				goto st11;
			}
			case 13:
			{
				goto st12;
			}
			case 32:
			{
				goto st11;
			}
			
		}
		{
			goto st0;
		}
		ctr52:
		{name = string(data[mark:p])
		}
		{params.M.Store(name, string(buf[0:amt]))	
		}
		
		
		goto st11;
		ctr56:
		{params.M.Store(name, string(buf[0:amt]))	
		}
		
		
		goto st11;
		st11:
		p+= 1;
		if p == pe  {
			goto _test_eof11;
			
		}
		st_case_11:
		switch ( data[p ])  {
			case 9:
			{
				goto st11;
			}
			case 13:
			{
				goto st12;
			}
			case 32:
			{
				goto st11;
			}
			case 33:
			{
				goto ctr13;
			}
			case 37:
			{
				goto ctr13;
			}
			case 39:
			{
				goto ctr13;
			}
			case 126:
			{
				goto ctr13;
			}
			
		}
		switch {
			case ( data[p ]) < 48 :
			{
				switch {
					case ( data[p ]) > 43 :
					{
						if 45 <= ( data[p ]) && ( data[p ]) <= 46  {
							{
								goto ctr13;
							}
							
						}
					} 
					case ( data[p ]) >= 42 :
					{
						goto ctr13;
					}
					
				}
			} 
			case ( data[p ]) > 57 :
			{
				switch {
					case ( data[p ]) > 90 :
					{
						if 95 <= ( data[p ]) && ( data[p ]) <= 122  {
							{
								goto ctr13;
							}
							
						}
					} 
					case ( data[p ]) >= 65 :
					{
						goto ctr13;
					}
					
				}
			} 
			default:
			{
				goto ctr13;
			}
			
		}
		{
			goto st0;
		}
		st12:
		p+= 1;
		if p == pe  {
			goto _test_eof12;
			
		}
		st_case_12:
		if ( data[p ]) == 10  {
			{
				goto st13;
			}
			
		}
		{
			goto st0;
		}
		st13:
		p+= 1;
		if p == pe  {
			goto _test_eof13;
			
		}
		st_case_13:
		switch ( data[p ])  {
			case 9:
			{
				goto st14;
			}
			case 32:
			{
				goto st14;
			}
			
		}
		{
			goto st0;
		}
		st14:
		p+= 1;
		if p == pe  {
			goto _test_eof14;
			
		}
		st_case_14:
		switch ( data[p ])  {
			case 9:
			{
				goto st14;
			}
			case 32:
			{
				goto st14;
			}
			case 33:
			{
				goto ctr13;
			}
			case 37:
			{
				goto ctr13;
			}
			case 39:
			{
				goto ctr13;
			}
			case 126:
			{
				goto ctr13;
			}
			
		}
		switch {
			case ( data[p ]) < 48 :
			{
				switch {
					case ( data[p ]) > 43 :
					{
						if 45 <= ( data[p ]) && ( data[p ]) <= 46  {
							{
								goto ctr13;
							}
							
						}
					} 
					case ( data[p ]) >= 42 :
					{
						goto ctr13;
					}
					
				}
			} 
			case ( data[p ]) > 57 :
			{
				switch {
					case ( data[p ]) > 90 :
					{
						if 95 <= ( data[p ]) && ( data[p ]) <= 122  {
							{
								goto ctr13;
							}
							
						}
					} 
					case ( data[p ]) >= 65 :
					{
						goto ctr13;
					}
					
				}
			} 
			default:
			{
				goto ctr13;
			}
			
		}
		{
			goto st0;
		}
		ctr13:
		{amt = 0
		}
		{mark = p
		}
		
		
		goto st37;
		st37:
		p+= 1;
		if p == pe  {
			goto _test_eof37;
			
		}
		st_case_37:
		switch ( data[p ])  {
			case 9:
			{
				goto ctr49;
			}
			case 13:
			{
				goto ctr50;
			}
			case 32:
			{
				goto ctr49;
			}
			case 33:
			{
				goto st37;
			}
			case 37:
			{
				goto st37;
			}
			case 39:
			{
				goto st37;
			}
			case 44:
			{
				goto ctr52;
			}
			case 61:
			{
				goto ctr53;
			}
			case 126:
			{
				goto st37;
			}
			
		}
		switch {
			case ( data[p ]) < 48 :
			{
				if 42 <= ( data[p ]) && ( data[p ]) <= 46  {
					{
						goto st37;
					}
					
				}
			} 
			case ( data[p ]) > 57 :
			{
				switch {
					case ( data[p ]) > 90 :
					{
						if 95 <= ( data[p ]) && ( data[p ]) <= 122  {
							{
								goto st37;
							}
							
						}
					} 
					case ( data[p ]) >= 65 :
					{
						goto st37;
					}
					
				}
			} 
			default:
			{
				goto st37;
			}
			
		}
		{
			goto st0;
		}
		ctr49:
		{name = string(data[mark:p])
		}
		{params.M.Store(name, string(buf[0:amt]))	
		}
		
		
		goto st15;
		st15:
		p+= 1;
		if p == pe  {
			goto _test_eof15;
			
		}
		st_case_15:
		switch ( data[p ])  {
			case 9:
			{
				goto st15;
			}
			case 13:
			{
				goto st16;
			}
			case 32:
			{
				goto st15;
			}
			case 44:
			{
				goto st11;
			}
			case 61:
			{
				goto st19;
			}
			
		}
		{
			goto st0;
		}
		ctr50:
		{name = string(data[mark:p])
		}
		{params.M.Store(name, string(buf[0:amt]))	
		}
		
		
		goto st16;
		st16:
		p+= 1;
		if p == pe  {
			goto _test_eof16;
			
		}
		st_case_16:
		if ( data[p ]) == 10  {
			{
				goto st17;
			}
			
		}
		{
			goto st0;
		}
		st17:
		p+= 1;
		if p == pe  {
			goto _test_eof17;
			
		}
		st_case_17:
		switch ( data[p ])  {
			case 9:
			{
				goto st18;
			}
			case 32:
			{
				goto st18;
			}
			
		}
		{
			goto st0;
		}
		st18:
		p+= 1;
		if p == pe  {
			goto _test_eof18;
			
		}
		st_case_18:
		switch ( data[p ])  {
			case 9:
			{
				goto st18;
			}
			case 32:
			{
				goto st18;
			}
			case 44:
			{
				goto st11;
			}
			case 61:
			{
				goto st19;
			}
			
		}
		{
			goto st0;
		}
		ctr53:
		{name = string(data[mark:p])
		}
		
		
		goto st19;
		st19:
		p+= 1;
		if p == pe  {
			goto _test_eof19;
			
		}
		st_case_19:
		switch ( data[p ])  {
			case 9:
			{
				goto st19;
			}
			case 13:
			{
				goto st20;
			}
			case 32:
			{
				goto st19;
			}
			case 33:
			{
				goto ctr22;
			}
			case 34:
			{
				goto st27;
			}
			case 37:
			{
				goto ctr22;
			}
			case 39:
			{
				goto ctr22;
			}
			case 93:
			{
				goto ctr22;
			}
			case 126:
			{
				goto ctr22;
			}
			
		}
		switch {
			case ( data[p ]) < 48 :
			{
				switch {
					case ( data[p ]) > 43 :
					{
						if 45 <= ( data[p ]) && ( data[p ]) <= 46  {
							{
								goto ctr22;
							}
							
						}
					} 
					case ( data[p ]) >= 42 :
					{
						goto ctr22;
					}
					
				}
			} 
			case ( data[p ]) > 58 :
			{
				switch {
					case ( data[p ]) > 91 :
					{
						if 95 <= ( data[p ]) && ( data[p ]) <= 122  {
							{
								goto ctr22;
							}
							
						}
					} 
					case ( data[p ]) >= 65 :
					{
						goto ctr22;
					}
					
				}
			} 
			default:
			{
				goto ctr22;
			}
			
		}
		{
			goto st0;
		}
		st20:
		p+= 1;
		if p == pe  {
			goto _test_eof20;
			
		}
		st_case_20:
		if ( data[p ]) == 10  {
			{
				goto st21;
			}
			
		}
		{
			goto st0;
		}
		st21:
		p+= 1;
		if p == pe  {
			goto _test_eof21;
			
		}
		st_case_21:
		switch ( data[p ])  {
			case 9:
			{
				goto st22;
			}
			case 32:
			{
				goto st22;
			}
			
		}
		{
			goto st0;
		}
		st22:
		p+= 1;
		if p == pe  {
			goto _test_eof22;
			
		}
		st_case_22:
		switch ( data[p ])  {
			case 9:
			{
				goto st22;
			}
			case 32:
			{
				goto st22;
			}
			case 33:
			{
				goto ctr22;
			}
			case 34:
			{
				goto st27;
			}
			case 37:
			{
				goto ctr22;
			}
			case 39:
			{
				goto ctr22;
			}
			case 93:
			{
				goto ctr22;
			}
			case 126:
			{
				goto ctr22;
			}
			
		}
		switch {
			case ( data[p ]) < 48 :
			{
				switch {
					case ( data[p ]) > 43 :
					{
						if 45 <= ( data[p ]) && ( data[p ]) <= 46  {
							{
								goto ctr22;
							}
							
						}
					} 
					case ( data[p ]) >= 42 :
					{
						goto ctr22;
					}
					
				}
			} 
			case ( data[p ]) > 58 :
			{
				switch {
					case ( data[p ]) > 91 :
					{
						if 95 <= ( data[p ]) && ( data[p ]) <= 122  {
							{
								goto ctr22;
							}
							
						}
					} 
					case ( data[p ]) >= 65 :
					{
						goto ctr22;
					}
					
				}
			} 
			default:
			{
				goto ctr22;
			}
			
		}
		{
			goto st0;
		}
		ctr22:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st38;
		st38:
		p+= 1;
		if p == pe  {
			goto _test_eof38;
			
		}
		st_case_38:
		switch ( data[p ])  {
			case 9:
			{
				goto ctr54;
			}
			case 13:
			{
				goto ctr55;
			}
			case 32:
			{
				goto ctr54;
			}
			case 33:
			{
				goto ctr22;
			}
			case 37:
			{
				goto ctr22;
			}
			case 39:
			{
				goto ctr22;
			}
			case 44:
			{
				goto ctr56;
			}
			case 93:
			{
				goto ctr22;
			}
			case 126:
			{
				goto ctr22;
			}
			
		}
		switch {
			case ( data[p ]) < 48 :
			{
				if 42 <= ( data[p ]) && ( data[p ]) <= 46  {
					{
						goto ctr22;
					}
					
				}
			} 
			case ( data[p ]) > 58 :
			{
				switch {
					case ( data[p ]) > 91 :
					{
						if 95 <= ( data[p ]) && ( data[p ]) <= 122  {
							{
								goto ctr22;
							}
							
						}
					} 
					case ( data[p ]) >= 65 :
					{
						goto ctr22;
					}
					
				}
			} 
			default:
			{
				goto ctr22;
			}
			
		}
		{
			goto st0;
		}
		ctr54:
		{params.M.Store(name, string(buf[0:amt]))	
		}
		
		
		goto st23;
		st23:
		p+= 1;
		if p == pe  {
			goto _test_eof23;
			
		}
		st_case_23:
		switch ( data[p ])  {
			case 9:
			{
				goto st23;
			}
			case 13:
			{
				goto st24;
			}
			case 32:
			{
				goto st23;
			}
			case 44:
			{
				goto st11;
			}
			
		}
		{
			goto st0;
		}
		ctr55:
		{params.M.Store(name, string(buf[0:amt]))	
		}
		
		
		goto st24;
		st24:
		p+= 1;
		if p == pe  {
			goto _test_eof24;
			
		}
		st_case_24:
		if ( data[p ]) == 10  {
			{
				goto st25;
			}
			
		}
		{
			goto st0;
		}
		st25:
		p+= 1;
		if p == pe  {
			goto _test_eof25;
			
		}
		st_case_25:
		switch ( data[p ])  {
			case 9:
			{
				goto st26;
			}
			case 32:
			{
				goto st26;
			}
			
		}
		{
			goto st0;
		}
		st26:
		p+= 1;
		if p == pe  {
			goto _test_eof26;
			
		}
		st_case_26:
		switch ( data[p ])  {
			case 9:
			{
				goto st26;
			}
			case 32:
			{
				goto st26;
			}
			case 44:
			{
				goto st11;
			}
			
		}
		{
			goto st0;
		}
		st27:
		p+= 1;
		if p == pe  {
			goto _test_eof27;
			
		}
		st_case_27:
		switch ( data[p ])  {
			case 9:
			{
				goto ctr30;
			}
			case 13:
			{
				goto ctr31;
			}
			case 34:
			{
				goto ctr32;
			}
			case 92:
			{
				goto ctr33;
			}
			
		}
		switch {
			case ( data[p ]) < 224 :
			{
				switch {
					case ( data[p ]) > 126 :
					{
						if 192 <= ( data[p ])  {
							{
								goto ctr34;
							}
							
						}
					} 
					case ( data[p ]) >= 32 :
					{
						goto ctr30;
					}
					
				}
			} 
			case ( data[p ]) > 239 :
			{
				switch {
					case ( data[p ]) < 248 :
					{
						{
							goto ctr36;
						}
					} 
					case ( data[p ]) > 251 :
					{
						if ( data[p ]) <= 253  {
							{
								goto ctr38;
							}
							
						}
					} 
					default:
					{
						goto ctr37;
					}
					
				}
			} 
			default:
			{
				goto ctr35;
			}
			
		}
		{
			goto st0;
		}
		ctr39:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st28;
		ctr30:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st28;
		st28:
		p+= 1;
		if p == pe  {
			goto _test_eof28;
			
		}
		st_case_28:
		switch ( data[p ])  {
			case 9:
			{
				goto ctr39;
			}
			case 13:
			{
				goto ctr40;
			}
			case 34:
			{
				goto st39;
			}
			case 92:
			{
				goto st31;
			}
			
		}
		switch {
			case ( data[p ]) < 224 :
			{
				switch {
					case ( data[p ]) > 126 :
					{
						if 192 <= ( data[p ])  {
							{
								goto ctr43;
							}
							
						}
					} 
					case ( data[p ]) >= 32 :
					{
						goto ctr39;
					}
					
				}
			} 
			case ( data[p ]) > 239 :
			{
				switch {
					case ( data[p ]) < 248 :
					{
						{
							goto ctr45;
						}
					} 
					case ( data[p ]) > 251 :
					{
						if ( data[p ]) <= 253  {
							{
								goto ctr47;
							}
							
						}
					} 
					default:
					{
						goto ctr46;
					}
					
				}
			} 
			default:
			{
				goto ctr44;
			}
			
		}
		{
			goto st0;
		}
		ctr40:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st29;
		ctr31:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st29;
		st29:
		p+= 1;
		if p == pe  {
			goto _test_eof29;
			
		}
		st_case_29:
		if ( data[p ]) == 10  {
			{
				goto ctr48;
			}
			
		}
		{
			goto st0;
		}
		ctr48:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st30;
		st30:
		p+= 1;
		if p == pe  {
			goto _test_eof30;
			
		}
		st_case_30:
		switch ( data[p ])  {
			case 9:
			{
				goto ctr39;
			}
			case 32:
			{
				goto ctr39;
			}
			
		}
		{
			goto st0;
		}
		ctr32:
		{amt = 0
		}
		
		
		goto st39;
		st39:
		p+= 1;
		if p == pe  {
			goto _test_eof39;
			
		}
		st_case_39:
		switch ( data[p ])  {
			case 9:
			{
				goto ctr54;
			}
			case 13:
			{
				goto ctr55;
			}
			case 32:
			{
				goto ctr54;
			}
			case 44:
			{
				goto ctr56;
			}
			
		}
		{
			goto st0;
		}
		ctr33:
		{amt = 0
		}
		
		
		goto st31;
		st31:
		p+= 1;
		if p == pe  {
			goto _test_eof31;
			
		}
		st_case_31:
		switch {
			case ( data[p ]) < 11 :
			{
				if ( data[p ]) <= 9  {
					{
						goto ctr39;
					}
					
				}
			} 
			case ( data[p ]) > 12 :
			{
				if 14 <= ( data[p ]) && ( data[p ]) <= 127  {
					{
						goto ctr39;
					}
					
				}
			} 
			default:
			{
				goto ctr39;
			}
			
		}
		{
			goto st0;
		}
		ctr43:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st32;
		ctr34:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st32;
		st32:
		p+= 1;
		if p == pe  {
			goto _test_eof32;
			
		}
		st_case_32:
		if 128 <= ( data[p ]) && ( data[p ]) <= 191  {
			{
				goto ctr39;
			}
			
		}
		{
			goto st0;
		}
		ctr44:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st33;
		ctr35:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st33;
		st33:
		p+= 1;
		if p == pe  {
			goto _test_eof33;
			
		}
		st_case_33:
		if 128 <= ( data[p ]) && ( data[p ]) <= 191  {
			{
				goto ctr43;
			}
			
		}
		{
			goto st0;
		}
		ctr45:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st34;
		ctr36:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st34;
		st34:
		p+= 1;
		if p == pe  {
			goto _test_eof34;
			
		}
		st_case_34:
		if 128 <= ( data[p ]) && ( data[p ]) <= 191  {
			{
				goto ctr44;
			}
			
		}
		{
			goto st0;
		}
		ctr46:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st35;
		ctr37:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st35;
		st35:
		p+= 1;
		if p == pe  {
			goto _test_eof35;
			
		}
		st_case_35:
		if 128 <= ( data[p ]) && ( data[p ]) <= 191  {
			{
				goto ctr45;
			}
			
		}
		{
			goto st0;
		}
		ctr47:
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st36;
		ctr38:
		{amt = 0
		}
		{buf[amt] = (( data[p ]))
			amt++
		}
		
		
		goto st36;
		st36:
		p+= 1;
		if p == pe  {
			goto _test_eof36;
			
		}
		st_case_36:
		if 128 <= ( data[p ]) && ( data[p ]) <= 191  {
			{
				goto ctr46;
			}
			
		}
		{
			goto st0;
		}
		st_out:
		_test_eof1: cs = 1;
		goto _test_eof; 
		_test_eof2: cs = 2;
		goto _test_eof; 
		_test_eof3: cs = 3;
		goto _test_eof; 
		_test_eof4: cs = 4;
		goto _test_eof; 
		_test_eof5: cs = 5;
		goto _test_eof; 
		_test_eof6: cs = 6;
		goto _test_eof; 
		_test_eof7: cs = 7;
		goto _test_eof; 
		_test_eof8: cs = 8;
		goto _test_eof; 
		_test_eof9: cs = 9;
		goto _test_eof; 
		_test_eof10: cs = 10;
		goto _test_eof; 
		_test_eof11: cs = 11;
		goto _test_eof; 
		_test_eof12: cs = 12;
		goto _test_eof; 
		_test_eof13: cs = 13;
		goto _test_eof; 
		_test_eof14: cs = 14;
		goto _test_eof; 
		_test_eof37: cs = 37;
		goto _test_eof; 
		_test_eof15: cs = 15;
		goto _test_eof; 
		_test_eof16: cs = 16;
		goto _test_eof; 
		_test_eof17: cs = 17;
		goto _test_eof; 
		_test_eof18: cs = 18;
		goto _test_eof; 
		_test_eof19: cs = 19;
		goto _test_eof; 
		_test_eof20: cs = 20;
		goto _test_eof; 
		_test_eof21: cs = 21;
		goto _test_eof; 
		_test_eof22: cs = 22;
		goto _test_eof; 
		_test_eof38: cs = 38;
		goto _test_eof; 
		_test_eof23: cs = 23;
		goto _test_eof; 
		_test_eof24: cs = 24;
		goto _test_eof; 
		_test_eof25: cs = 25;
		goto _test_eof; 
		_test_eof26: cs = 26;
		goto _test_eof; 
		_test_eof27: cs = 27;
		goto _test_eof; 
		_test_eof28: cs = 28;
		goto _test_eof; 
		_test_eof29: cs = 29;
		goto _test_eof; 
		_test_eof30: cs = 30;
		goto _test_eof; 
		_test_eof39: cs = 39;
		goto _test_eof; 
		_test_eof31: cs = 31;
		goto _test_eof; 
		_test_eof32: cs = 32;
		goto _test_eof; 
		_test_eof33: cs = 33;
		goto _test_eof; 
		_test_eof34: cs = 34;
		goto _test_eof; 
		_test_eof35: cs = 35;
		goto _test_eof; 
		_test_eof36: cs = 36;
		goto _test_eof; 
		
		_test_eof: {}
		if p == eof  {
			{
				switch cs  {
					case 38:
					fallthrough
					case 39:
					{params.M.Store(name, string(buf[0:amt]))	
					}
					
					break;
					case 37:
					{name = string(data[mark:p])
					}
					{params.M.Store(name, string(buf[0:amt]))	
					}
					
					break;
					
				}
			}
			
			
		}
		_out: {}
	}
	if cs < www_authenticate_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete data: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in data at pos %d: %s", p, data))
		}
	}
	
	return params, nil
}


