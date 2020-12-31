from pathlib import Path
from typing import List
from collections import defaultdict
import numpy as np

input_file = Path.cwd()/"inputs"/"day22.input"

def part_a(deck_1, deck_2) -> int:
    score = score_deck(get_winning_deck(play_game(deck_1, deck_2)))
    return score

def part_b(deck_1, deck_2) -> int:
    return score_deck(get_winning_deck(play_recursive_game(deck_1, deck_2)))

def score_deck(deck):
    return sum([(i+1)*c for (i,c) in enumerate(deck[::-1])])

def get_winning_deck(decks):
    for deck in decks:
        if deck != []:
            return deck
    raise AttributeError(decks)

def hash_game(deck_1, deck_2):
    str_game = f'{",".join([str(c) for c in deck_1])}|{",".join([str(c) for c in deck_2])}'
    return hash(str_game)

def play_recursive_game(deck_1, deck_2):
    return play_subgame(deck_1, deck_2)

def play_subgame(deck_1, deck_2):
    moves_played = set()

    while deck_1 != [] and deck_2 != []:
        hash_current = hash_game(deck_1, deck_2)
        if hash_current in moves_played:
            return (deck_1, [])

        moves_played.add(hash_current)

        winner = determine_winner_turn(deck_1, deck_2)

        top_card_1 = deck_1[0]
        top_card_2 = deck_2[0]

        if winner == 1:
            deck_1 = deck_1[1:] + [top_card_1, top_card_2]
            deck_2 = deck_2[1:]
        else :
            deck_2 = deck_2[1:] + [top_card_2, top_card_1]
            deck_1 = deck_1[1:]
    print(f"Played subgame !")
    return deck_1, deck_2

def determine_winner_turn(deck_1, deck_2):
    top_card_1 = deck_1[0]
    top_card_2 = deck_2[0]

    if (top_card_1) <= (len(deck_1) - 1) and (top_card_2) <= (len(deck_2) - 1):
        return get_winner_id(play_subgame(deck_1[1:(1+top_card_1)], deck_2[1:(1+top_card_2)]))
    else:
        if top_card_1 > top_card_2 :
            return 1
        else :
            return 2

def get_winner_id(decks):
    for i, deck in enumerate(decks):
        if deck != []:
            return (i+1)
    return -1

def play_game(deck_1, deck_2):
    while deck_1 != [] and deck_2 != []:
        deck_1, deck_2 = play_turn(deck_1, deck_2)
    return deck_1, deck_2

def play_turn(deck_1, deck_2):
    top_card_1 = deck_1[0]
    top_card_2 = deck_2[0]
    if top_card_1 > top_card_2 :
        deck_1 = deck_1[1:] + [top_card_1, top_card_2]
        deck_2 = deck_2[1:]
    if top_card_2 > top_card_1 :
        deck_2 = deck_2[1:] + [top_card_2, top_card_1]
        deck_1 = deck_1[1:]
    return deck_1, deck_2


def parse_string(input_string : str):
    p1_string, p2_string = input_string.split("\n\n")
    return (parse_deck_string(p1_string), parse_deck_string(p2_string))

def parse_deck_string(deck_string : str) -> List[int]:
    return [int(x) for x in deck_string.split("\n")[1:]]

test_string = """Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10"""

test_string_should_finish ="""Player 1:
43
19

Player 2:
2
29
14"""

assert part_a(*parse_string(test_string)) == 306
assert part_b(*parse_string(test_string)) == 291
assert part_b(*parse_string(test_string_should_finish))

print("Passed tests !")

with open(input_file, "r") as f:
    input_string = f.read()

print(part_a(*parse_string(input_string)))
print(part_b(*parse_string(input_string)))
