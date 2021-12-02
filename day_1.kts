import java.io.File

fun day1() {
    val input = File("in/input1.txt")
        .readLines()
        .map { it.toInt() }

    var acc = 0
    input.reduce { a, b -> if (b > a) acc++; b }
    println("sol1: $acc")

    acc = 0
    input.windowed(3, 1).map { it.sum() }.reduce { a, b -> if (b > a) acc++; b }
    println("sol2: $acc")
}

day1()