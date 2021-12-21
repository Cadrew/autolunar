# Autolunar

Autolunar is a pseudorandom number generator (PRNG) based on cellular automata.

## How does it work?

In Autolunar, it is possible to configure one or more cellular automatons.
These cellular automata are defined by rules in the `rules` folder. It is possible to add as many rules as you want (see the "Rules" section for more explanations).

The generator then relies on seeds (examples available in the `seeds` folder) to execute the cellular automaton.
If several cellular automata have been declared, the generator will switch between the different automata randomly. At the end of the execution (defined by a sleep time in ms), the generator returns a value generated from the current state of the last cellular automaton that was executed.

Currently, only 2D cellular automata work.
To generate a number from a state of the cellular automaton, we proceed as follows:
- The grid of the automaton is 64x64 cells.
- Each cell has either the value 0 or the value 1.
- We create blocks of 8 bits (8 cells). There are thus 8 blocks of 8 bits.
- For each block, we calculate the power of 2 of the corresponding cell multiplied by the value of its cell.
- Finally, we calculate the final value by doing the XOR operation between all the blocks.

This process is the application of the one described in the paper available [here](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.759.6207&rep=rep1&type=pdf).

### Rules

Inside the `rules` directory, you can add as many rules as you want.
These rules are based on the BxSy notation. 

All the current examples are 2D cellular automata rules, but it is possible to create rules in any dimension.
The rule file must be in JSON format.

| property   | description                                          |
|------------|------------------------------------------------------|
| name       | `string` - name of the rule (e.g. "Game of Life")    |
| model      | `string` - model type (e.g. "life-like)              |
| birth      | `array` - values for which a new cell is born        |
| survive    | `array` - values for which a cell survive            |
| moore      | `number` - moore neighborhood                        |
| dimensions | `number` - number of dimensions                      |
| states     | `array` - possible values for a cell (e.g. `[0, 1]`) |
| BxSy       | `string` - BxSy name (e.g. B3S23)                    |

In the code, if the file is called `my_custom_rule.json` it is then possible to add this new rule in the generator in this way:
```golang
automaton, err := ReadRule("my_custom_rule")
if (err != nil) {
    return err
}

// add custom automaton rule
al.AddAutomaton(automaton, <your_seed>)
```

The default rule for the generator is `fredkin`. The reason for this is that this rule has the highest entropy rate. Thus, the generated chaos value is much larger with this generator, which allows to produce numbers that look statistically as random as possible.

### Seeds

Seeds are located inside the `seeds` directory. You can add as many seeds you want.

The seeds can be used for any type of cellular automaton as long as they respect the number of dimensions.
Seed files are simple CSV. The delimiter must be a comma `,`.

In the code, if the file is called `my_custom_seed.csv` it is then possible to add this new seed in the generator in this way:
```golang
seed, err := ReadSeed("my_custom_seed")
if (err != nil) {
    return err
}

// add custom seed
al.AddAutomaton(<your_automaton>, seed)
```

### Usage

```golang
import autolunar "github.com/Cadrew/autolunar/lib"
```

Simple usage with default settings:

```golang
al := autolunar.CreateGenerator()
err := al.SetDefault()
if err != nil {
    fmt.Println(err)
    return
}
// generate an int64 between 0 and 100
rng := al.Rand(0, 100)
```

You can use your custom settings:

```golang
al := autolunar.CreateGenerator()
// the larger the number, the longer the execution but the numbers generated will be much more random (default is 10)
al.sleep = 100 // in ms
automaton, err := ReadRule("fredkin") // example with fredkin, you can use the rules you want
if (err != nil) {
    return err
}
gun, err := ReadSeed("gun") // example with gun, you can use the seed you want
if err != nil {
    return err
}

// add automaton with custom seed
// you can add as many automaton you want
// the more automata you add, the better will be the result
al.AddAutomaton(fredkin, gun)

// generate an int64 between 0 and 100
rng := al.Rand(0, 100)
```

## Build

```
go build
```

## Statistical tests

In order to determine if the generator is a cryptographic generator, it needs to pass statistical tests.
In the `stats` directory, there are statistical reports for each of the tests.
It should be noted that the results of the tests depend strongly on the cellular automata used and the seeds.

### DIEHARD

<!-- In its default configuration (using fredkin and amoeba automata), the generator has passed the DIEHARD test, see the report diarhard.txt in the `stats` directory.
To pass this test, 35M random numbers were generated in a file `numbers.txt`. The file was tested with [dieharder](https://linux.die.net/man/1/dieharder) on Linux. -->

To install dieharder:
```
sudo apt-get install -y dieharder
```

The command to run the tests:
```
dieharder -g 202 -f numbers.txt -a
```

To understand the DIEHARD report, it is mainly necessary to read the name of the test that was passed in the `test_name` column as well as the `p-value` that was calculated.
If the `p-value` is exactly 0 or exactly 1, then the test is considered to have failed, otherwise it is considered to have passed.

## TODOs

- Find good seeds
- Adjust default settings
- Optimize execution
