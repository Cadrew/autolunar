# Autolunar

Autolunar is a cryptographic pseudorandom number generator (PRNG) based on cellular automata.

### Usage

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
