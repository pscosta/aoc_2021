import java.io.File

fun day1() {
    val input = File("in/input1.txt")
        .readLines()
        .map { it.toInt() }

    var currReading = input.first()
    var acc = 0

    input.drop(0).forEach {
        if (it > currReading) acc += 1
        currReading = it
    }
    println("sol1: $acc")

    val input2 = input.windowed(3, 1).map { it.sum() }
    currReading = input2.first()
    acc = 0

    input2.drop(0).forEach {
        if (it > currReading) acc += 1
        currReading = it
    }
    println("sol2: $acc")
}

day1()