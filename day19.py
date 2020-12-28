from pathlib import Path
from typing import List, Tuple, Dict
from collections import defaultdict
import re

input_file = Path.cwd()/"inputs"/"day19.input"

def to_regex(rules) -> str:
    regex_rule = defaultdict(str)
    return resolve(rules, "0", regex_rule)

def to_regex_cheating(rules) -> str:
    regex_rule = defaultdict(str)
    return f"(({resolve(rules, '42', regex_rule)})+)(({resolve(rules, '31', regex_rule)})+)"

def resolve(rules : Dict[int, List[str]], idx : int, regex_rule : Dict[int, str]) -> str:
    if idx in regex_rule:
        return regex_rule[idx]

    rules_idx = rules[idx]
    rules_subs_wo_self = [rs for rs in split(rules_idx,"|") if idx not in rs]
    rules_subs_w_self = [rs for rs in split(rules_idx,"|") if idx in rs]
    assert len(rules_subs_w_self) <= 1
    
    regexs_wo_selfloop = [resolve_rules(rs, idx, rules, regex_rule) for rs in rules_subs_wo_self]
    regex_no_loop = f'{"|".join(regexs_wo_selfloop)}'
    if len(regexs_wo_selfloop) > 1:
        regex_no_loop = f"(?:{regex_no_loop})"

    if len(rules_subs_w_self) == 0 :
        regex_rule[idx] = regex_no_loop
        return regex_no_loop
    else:
        rules_self = rules_subs_w_self[0]
        i_self = rules_self.index(idx)

        prerules = rules_self[:i_self]
        postrules = rules_self[(i_self+1):]
        regex_pre = resolve_rules(prerules, idx, rules, regex_rule)        
        regex_post = resolve_rules(postrules, idx, rules, regex_rule)        

        regex_idx = regex_no_loop
        if regex_pre:
            regex_idx = f"({regex_pre})*{regex_idx}"
        if regex_post:
            regex_idx = f"{regex_idx}({regex_post})*"

        regex_rule[idx] = regex_idx
        return regex_idx

def resolve_rules(rs, idx, rules, regex_rule):
    assert idx not in rs
    assert "|" not in rs

    regex = ""
    for r in rs:
        if r not in ("a","b"):
            regex += resolve(rules, r, regex_rule)
        else:
            regex += r
    if len(rs) > 1:
        regex = f"{regex}"
    return regex

def split(sequence, sep):
    chunk = []
    for val in sequence:
        if val == sep:
            yield chunk
            chunk = []
        else:
            chunk.append(val)
    yield chunk

def part_a(rules, strings : List[str]) -> int:
    regex = re.compile(to_regex(rules))
    print(regex.pattern)
    matches = [x for x in strings if regex.fullmatch(x)]
    return len(matches)

def part_b(rules, strings : List[str]) -> int:
    pattern = to_regex_cheating(rules)
    print(pattern)
    matches = [x for x in strings if re.fullmatch(pattern, x)]
    return len(matches)

def parse_string(string) -> Tuple[List, List[str]]:
    strings = string.split("\n")
    index_break = [i for i,s in enumerate(strings) if s == ""][0]
    rules = parse_rules(strings[:index_break])
    strings = strings[(index_break+1):]
    return rules, strings

def parse_rules(str_rules):
    rules = defaultdict(list)
    for l in str_rules:
        index = l.split(" ")[0].rstrip(":")
        rules_index = l.split(" ")[1:]
        rules[index] = rules_index
    return rules

test_string_easy = """0: 1 2
1: a
2: 1 3
3: b

aa
abb
aab
baa
aba"""

test_string = """0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: a
5: b

ababbb
bababa
abbbab
aaabbb
aaaabbb"""

test_string_repeat = """0: a b | a 0 b

abb
aabb
a
b
ab"""

test_string_b = """42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: a
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: b
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba"""

test_string_real_input = """0: 8 11
8: 42
31: b
42 : a
11 : 42 31

a
b
ab
"""

assert (part_a(*parse_string(test_string))) == 2
assert (part_a(*parse_string(test_string_repeat))) == 3
assert (part_a(*parse_string(test_string_b))) == 3
assert (part_b(*parse_string(test_string_b))) == 12

print("Passed tests !")

print(part_a(*parse_string(test_string_real_input)))
print(part_b(*parse_string(test_string_real_input)))


with open(input_file, "r") as f: 
    input_string = f.read()

print(part_a(*parse_string(input_string)))
print(part_b(*parse_string(input_string)))
