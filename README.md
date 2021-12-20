# Autolunar

Autolunar is a pseudorandom number generator (PRNG) based on cellular automata.

### How does it work?

In Autolunar, it is possible to configure one or more cellular automatons.
These cellular automata are defined by rules in the `rules` folder. It is possible to add as many rules as you want (see the "Rules" section for more explanations).

The generator then relies on seeds (examples available in the `seeds` folder) to execute the cellular automaton.
If several cellular automata have been declared, the generator will switch between the different automata randomly. At the end of the execution (defined by a sleep time in ms), the generator returns a value generated from the current state of the last cellular automaton that was executed.

### Rules

Inside the `rules` folder, you can add as many rules as you want.
These rules are based on the BxSy notation. 

All the current examples are 2D cellular automata rules, but it is possible to create rules in any dimension.
The rule file must be in JSON format.

| name       | `string` - name of the rule (e.g. "Game of Life")    |
|------------|------------------------------------------------------|
| model      | `string` - model type (e.g. "life-like)              |
|------------|------------------------------------------------------|
| birth      | `array` - values for which a new cell is born        |
|------------|------------------------------------------------------|
| survive    | `array` - values for which a cell survive            |
|------------|------------------------------------------------------|
| moore      | `number` - moore neighborhood                        |
|------------|------------------------------------------------------|
| dimensions | `number` - number of dimensions                      |
|------------|------------------------------------------------------|
| states     | `array` - possible values for a cell (e.g. `[0, 1]`) |
|------------|------------------------------------------------------|
| BxSy       | `string` - BxSy name (e.g. B3S23)                    |
|------------|------------------------------------------------------|

In the code, if the file is called `my_custom_rule.json` it is then possible to add this new rule in the generator in this way:
```golang
automaton, err := ReadRule("my_custom_rule")
if (err != nil) {
    return err
}

// add custom automaton
al.AddAutomaton(automaton, <your_seed>)
```

### Usage

Import:

```golang
import autolunar "github.com/Cadrew/autolunar/lib"
```

Simple usage use default settings:

```golang
al := autolunar.CreateGenerator()
err := al.SetDefault()
if err != nil {
    fmt.Println(err)
    return
}
// generate an int between 0 and 100
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

// add automaton with custom seed
al.AddAutomaton(automaton, [][]uint8{
    {1, 5}, {1, 6}, {2, 5}, {2, 6},
    {11, 5}, {11, 6}, {11, 7}, {12, 4}, {12, 8}, {13, 3},
    {13, 9}, {14, 3}, {14, 9}, {15, 6}, {16, 4}, {16, 8},
    {17, 5}, {17, 6}, {17, 7}, {18, 6},
    {21, 3}, {21, 4}, {21, 5}, {22, 3}, {22, 4}, {22, 5},
    {23, 2}, {23, 6}, {25, 1}, {25, 2}, {25, 6}, {25, 7},
    {35, 3}, {35, 4}, {36, 3}, {36, 4},
})
// you can add as many automaton you want
// the more automata you add, the better will be the result

// generate an int between 0 and 100
rng := al.Rand(0, 100)
```

### Build

```
go build
```

### TODOs

- Make output tests to check if this is cryptographic
- Find good seeds
- Read seeds from CSV
- Adjust default settings
- Optimize execution
- Write docs & readme
