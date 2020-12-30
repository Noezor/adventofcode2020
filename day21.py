from collections import defaultdict
from pathlib import Path
from typing import List, Tuple

input_file = Path.cwd()/"inputs"/"day21.input"


def part_a(content) -> int:
    alergens_of_ingredient = get_possible_alergen_of_ingredient(content)
    ingredients_wo_alergens = [x for x in alergens_of_ingredient if len(alergens_of_ingredient[x]) == 0]
    count_appearance_ingredients_wo_alergen = 0
    for ingredients, _ in content:
        for ingredient in ingredients:
            if ingredient in ingredients_wo_alergens:
                count_appearance_ingredients_wo_alergen += 1
    return count_appearance_ingredients_wo_alergen

def part_b(content) -> str:
    alergens_of_ingredient = get_possible_alergen_of_ingredient(content)
    ingredient_and_alergens = greedily_assign_alergens(alergens_of_ingredient)
    return ",".join([ingredient for ingredient,_ in sorted(ingredient_and_alergens, key = lambda i_a : i_a[1])])

def greedily_assign_alergens(alergens_of_ingredients) -> List[Tuple[str,str]]:
    assigned_alergens = []
    assigned_ingredients = []
    while sum([len([alergen for alergen in alergens_of_ingredients[ingredient] if alergen not in assigned_alergens]) for ingredient in alergens_of_ingredients]) != 0:
        for ingredient in alergens_of_ingredients:
            possible_alergens = [alergen for alergen in alergens_of_ingredients[ingredient] if alergen not in assigned_alergens]
            if len(possible_alergens) == 1:
                assigned_alergens.append(possible_alergens[0])
                assigned_ingredients.append(ingredient)
    return list(zip(assigned_ingredients, assigned_alergens))


def get_possible_alergen_of_ingredient(content : List[Tuple[List[str], List[str]]]):
    ingredient_with_alergen = defaultdict(set)
    all_ingredients = set()
    for ingredients, alergens in content:
        all_ingredients.update(ingredients)
        for alergen in alergens:
            if alergen not in ingredient_with_alergen:
                ingredient_with_alergen[alergen].update(ingredients)
            else:
                ingredient_with_alergen[alergen] = ingredient_with_alergen[alergen].intersection(set(ingredients))
    return {ingredient:[alergen for alergen in ingredient_with_alergen if ingredient in ingredient_with_alergen[alergen]] for ingredient in all_ingredients}

def parse_string(input_string : str):
    return [parse_line(line) for line in input_string.split("\n")]

def parse_line(line : str) -> Tuple[List[str], List[str]]:
    split = line.split(" (contains ")
    ingredients = split[0].split(" ")
    alergens = split[1][:-1].split(", ")
    return ingredients, alergens

test_string = """mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)"""

print(parse_string(test_string))

assert part_a(parse_string(test_string)) == 5
assert part_b(parse_string(test_string)) == "mxmxvkd,sqjhc,fvjkl"


print("Passed tests !")

with open(input_file, "r") as f:
    input_string = f.read()

print(part_a(parse_string(input_string)))
print(part_b(parse_string(input_string)))
