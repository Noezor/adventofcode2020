from pathlib import Path
from typing import List

from tqdm import tqdm

input_file = Path.cwd()/"inputs"/"day23.input"


def part_a(initial_setup) -> str:
    final_setup = play_game(initial_setup, 100)
    return str_after_one(final_setup)

def part_b(initial_setup) -> int:
    setup = initial_setup + [i for i in range(len(initial_setup)+1, 1000000+1)]
    final_setup = play_game(setup, 10000000)
    index_one = final_setup.index(1)
    return final_setup[(index_one+1)%len(final_setup)]*final_setup[(index_one+2)%len(final_setup)]

def str_after_one(setup : List) -> str:
    index_one = setup.index(1)
    return "".join([str(x) for x in (setup[(index_one+1):] + setup[:index_one])])

def play_game(setup, nb_moves):
    current_card = setup[0]
    for _ in tqdm(range(nb_moves)):
        setup, current_card = play_turn(setup, current_card)
    return setup

def play_turn(setup, current_card):
    idx_current_card = setup.index(current_card)
    picked, cards_left = pick_cards(setup, idx_current_card)
    # assert len(picked) == 3
    index_insert = cards_left.index(find_index_insert_from_picked(picked, current_card, len(setup)))
    next_setup = setup
    # next_setup = [*(cards_left[:(index_insert+1)]), *picked, *(cards_left[(index_insert+1):])]
    next_card = next_setup[(next_setup.index(current_card)+1)%len(next_setup)]
    return next_setup, next_card

def pick_cards(setup, idx_current_card):
    picked = setup[(idx_current_card+1):(idx_current_card+1+3)]
    left = setup
    left = [*(setup[:(idx_current_card+1)]), *(setup[(idx_current_card+1+3):])]
    if len(picked) != 3:
        to_pick = 3 - len(picked)
        picked = [*picked,*(setup[:to_pick])]
        left = left[to_pick:]
    return picked, left

def find_index_insert(cards_left : List[int], current_card : int) -> int:
    lower_than_current = [c for c in cards_left if c < current_card]
    if lower_than_current != []:
        return max(lower_than_current)
    else :
        return max(cards_left)

def find_index_insert_from_picked(picked : List[int], current_card : int, nb_cards) -> int:
    potential_picks_lower = [current_card-i for i in range(1,3+1) if current_card-i > 0 and current_card-i not in picked]
    if potential_picks_lower != []:
        return potential_picks_lower[0]
    potential_cards_upper = [nb_cards-i for i in range(3) if nb_cards-i not in picked]
    return potential_cards_upper[0]


def parse_string(input_string : str):
    return [int(x) for x in input_string]

test_string = """389125467"""


assert find_index_insert_from_picked([8, 9, 1], 3, 9) == find_index_insert([3,2,5,4,6,7], 3), (find_index_insert_from_picked([8, 9, 1], 3, 9),find_index_insert([3,2,5,4,6,7], 3)) 
assert str_after_one([5, 8, 3,  7,  4,  1,  9,  2,  6]) == "92658374"
# assert part_a(parse_string(test_string)) == "67384529", part_a(parse_string(test_string))
# assert part_b(parse_string(test_string)) == 149245887792
print("Passed tests !")

with open(input_file, "r") as f:
    input_string = f.read()

print(part_a(parse_string(input_string)))
print(part_b(parse_string(input_string)))
