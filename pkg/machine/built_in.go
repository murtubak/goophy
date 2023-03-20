// package machine

// builtin_mapping := {
//     display       : display,
//     get_time      : get_time,
//     stringify     : stringify,
//     error         : error,
//     prompt        : prompt,
//     is_number     : is_number,
//     is_string     : is_string,
//     is_function   : x => typeof x === 'object' &&
//                          (x.tag == 'BUILTIN' ||
//                           x.tag == 'CLOSURE'),
//     is_boolean    : is_boolean,
//     is_undefined  : is_undefined,
//     parse_int     : parse_int,
//     char_at       : char_at,
//     arity         : x => typeof x === 'object'
//                          ? x.arity
//                          : error(x, 'arity expects function, received:'),
//     math_abs      : math_abs,
//     math_acos     : math_acos,
//     math_acosh    : math_acosh,
//     math_asin     : math_asin,
//     math_asinh    : math_asinh,
//     math_atan     : math_atan,
//     math_atanh    : math_atanh,
//     math_atan2    : math_atan2,
//     math_ceil     : math_ceil,
//     math_cbrt     : math_cbrt,
//     math_expm1    : math_expm1,
//     math_clz32    : math_clz32,
//     math_cos      : math_cos,
//     math_cosh     : math_cosh,
//     math_exp      : math_exp,
//     math_floor    : math_floor,
//     math_fround   : math_fround,
//     math_hypot    : math_hypot,
//     math_imul     : math_imul,
//     math_log      : math_log,
//     math_log1p    : math_log1p,
//     math_log2     : math_log2,
//     math_log10    : math_log10,
//     math_max      : math_max,
//     math_min      : math_min,
//     math_pow      : math_pow,
//     math_random   : math_random,
//     math_round    : math_round,
//     math_sign     : math_sign,
//     math_sin      : math_sin,
//     math_sinh     : math_sinh,
//     math_sqrt     : math_sqrt,
//     math_tanh     : math_tanh,
//     math_trunc    : math_trunc,
//     pair          : pair,
//     is_pair       : is_pair,
//     head          : head,
//     tail          : tail,
//     is_null       : is_null,
//     set_head      : set_head,
//     set_tail      : set_tail,
//     array_length  : array_length,
//     is_array      : is_array,
//     list          : list,
//     is_list       : is_list,
//     display_list  : display_list,
//     // from list libarary
//     equal         : equal,
//     length        : length,
//     list_to_string: list_to_string,
//     reverse       : reverse,
//     append        : append,
//     member        : member,
//     remove        : remove,
//     remove_all    : remove_all,
//     enum_list     : enum_list,
//     list_ref      : list_ref,
//     // misc
//     draw_data     : draw_data,
//     parse         : parse,
//     tokenize      : tokenize,
//     apply_in_underlying_javascript: apply_in_underlying_javascript
// }
