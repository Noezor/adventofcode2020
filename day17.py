import itertools
from pathlib import Path
from typing import Dict, Tuple
from collections import defaultdict
from itertools import product
from copy import deepcopy

input_file = Path.cwd()/"inputs"/"day17.input"

ACTIVE = "#"
INACTIVE = "."
rec_dd = lambda: defaultdict(rec_dd)

class Environment:
    def __init__(self, env : Dict[int, Dict[int, Dict[int, bool]]], range_x, range_y, range_z):
        self.env = env
        self.range_x = range_x
        self.range_y = range_y
        self.range_z = range_z

    @staticmethod
    def from_string(strings):
        env = rec_dd()
        max_x = len(strings)

        z = 0
        for x, line in enumerate(strings):
            max_y = len(line)
            for y, c in enumerate(line):
                if c == ACTIVE:
                    env[x][y][z] = True

        new_env = Environment(env, (0-1, max_x+1), (0-1, max_y+1), (0-1,1+1))
        return new_env

    def is_active(self, pt):
        assert len(pt) == 3, pt
        return self.env[pt[0]][pt[1]][pt[2]]

    def get_next(self):
        new_environment = rec_dd()
        for pt in self :
            neighbor_coords = get_neighbors(pt)
            active_neighbors = [neighbor for neighbor in neighbor_coords if self.is_active(neighbor)]
            if self.is_active(pt):
                if len(active_neighbors) in (2,3) : 
                    new_environment[pt[0]][pt[1]][pt[2]] = True
                else:
                    new_environment[pt[0]][pt[1]][pt[2]] = False

            if not self.is_active(pt):
                if len(active_neighbors) == (3):
                    new_environment[pt[0]][pt[1]][pt[2]] = True
                else :
                    new_environment[pt[0]][pt[1]][pt[2]] = False

        return Environment(new_environment, self.new_range(self.range_x), self.new_range(self.range_y), self.new_range(self.range_z))

    def new_range(self, range):
        return (range[0] - 1, range[1] + 1)

    def __iter__(self):
        for x in range(*self.range_x):
            for y in range(*self.range_y):
                for z in range(*self.range_z):
                    yield (x,y,z)

    def __repr__(self) -> str:
        str = ""
        for z in range(*self.range_z):
            for x in range(*self.range_x):
                for y in range(*self.range_y):
                    if self.is_active((x,y,z)) :                    
                        str += ACTIVE   
                    else:
                        str += INACTIVE   
                str += "\n"
            str += f"z={z}\n"

        return str

    def count_active(self):
        return len([pt for pt in self if self.is_active(pt)])

class EnvironmentB:
    def __init__(self, env : Dict[int, Dict[int, Dict[int, bool]]], range_x, range_y, range_z, range_w):
        self.env = env
        self.range_x = range_x
        self.range_y = range_y
        self.range_z = range_z
        self.range_w = range_w


    @staticmethod
    def from_string(strings):
        env = rec_dd()
        max_x = len(strings)

        w = 0
        z = 0
        for x, line in enumerate(strings):
            max_y = len(line)
            for y, c in enumerate(line):
                if c == ACTIVE:
                    env[x][y][z][w] = True

        new_env = EnvironmentB(env, (0-1, max_x+1), (0-1, max_y+1), (0-1,1+1), (0-1, 1+1))
        return new_env

    def is_active(self, pt):
        assert len(pt) == 4, pt
        return self.env[pt[0]][pt[1]][pt[2]][pt[3]]

    def get_next(self):
        new_environment = rec_dd()
        for pt in self :
            neighbor_coords = get_neighbors(pt)
            active_neighbors = [neighbor for neighbor in neighbor_coords if self.is_active(neighbor)]
            if self.is_active(pt):
                if len(active_neighbors) in (2,3) : 
                    new_environment[pt[0]][pt[1]][pt[2]][pt[3]] = True
                else:
                    new_environment[pt[0]][pt[1]][pt[2]][pt[3]] = False

            if not self.is_active(pt):
                if len(active_neighbors) == (3):
                    new_environment[pt[0]][pt[1]][pt[2]][pt[3]] = True
                else :
                    new_environment[pt[0]][pt[1]][pt[2]][pt[3]] = False

        return EnvironmentB(new_environment, self.new_range(self.range_x), self.new_range(self.range_y), self.new_range(self.range_z), self.new_range(self.range_w))

    def new_range(self, range):
        return (range[0] - 1, range[1] + 1)

    def __iter__(self):
        for x in range(*self.range_x):
            for y in range(*self.range_y):
                for z in range(*self.range_z):
                    for w in range(*self.range_w):
                        yield (x,y,z,w)

    def count_active(self):
        return len([pt for pt in self if self.is_active(pt)])

def parse_file(filepath) -> Environment:
    with open(filepath, 'r') as f:
        strings = f.read().split("\n")
        env = Environment.from_string(strings)
    return env

def parse_file_b(filepath) -> EnvironmentB:
    with open(filepath, 'r') as f:
        strings = f.read().split("\n")
        env = EnvironmentB.from_string(strings)
    return env

def part_a(env : Environment) -> int :
    nb_simulations = 6

    for _ in range(nb_simulations):
        env = env.get_next()
    return env.count_active()

def part_b(env : EnvironmentB) -> int :
    nb_simulations = 6

    for _ in range(nb_simulations):
        env = env.get_next()
    return env.count_active()


def get_neighbors(pt):
    possible_coordinates = [(x-1,x,x+1) for x in pt]
    return [p for p in itertools.product(*possible_coordinates) if p != pt]

def max_distance(pt1, pt2):
    return max([abs(x-y) for (x,y) in zip (pt1,pt2)])

def are_neighbors(pt1, pt2):
    return max_distance(pt1, pt2) <= 1


assert len(get_neighbors((2,2,2))) == 26 
assert len(get_neighbors((2,2,2,2))) == 80 


test_input = """.#.
..#
###""".split("\n")

env = parse_file(input_file)
print(env)

assert (Environment.from_string(test_input).is_active((0,1,0)))
assert part_a(Environment.from_string(test_input)) == 112, part_a(Environment.from_string(test_input))
assert part_b(EnvironmentB.from_string(test_input)) == 848, part_b(EnvironmentB.from_string(test_input))


print("Passed tests !")


print(part_a(parse_file(input_file)))
print(part_b(parse_file_b(input_file)))