// long variable declaration in if blocks
package main

func main() {
    if exp {
        var x int
        var y = 23
        var z int = 23

        var x1, x2 int
        var y1, y2 = 23, 23
        var z1, z2 int = 23, 23

        var (
            x3, x4 int
            x5 int
            y3, y4 = 23, 23
            y5 = 23
            z3, z4 int = 23, 23
            z4 int = 23
        )
    }
}
