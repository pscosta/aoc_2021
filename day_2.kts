import java.io.File

fun day2Sol1() {
    var horiz = 0
    var depth = 0

    File("in/input2.txt")
        .readLines()
        .mapNotNull { "^(\\w+) (\\d+)$".toRegex().find(it)?.destructured }
        .forEach { (dir, amount) ->
            when (dir) {
                "forward" -> horiz += amount.toInt()
                "down" -> depth += amount.toInt()
                "up" -> depth -= amount.toInt()
            }
        }
        .also { println("sol1: ${horiz * depth}") }
}

fun day2Sol2() {
    var horiz = 0
    var depth = 0
    var aim = 0

    File("in/input2.txt")
        .readLines()
        .mapNotNull { "^(\\w+) (\\d+)$".toRegex().find(it)?.destructured }
        .forEach { (dir, amount) ->
            when (dir) {
                "down" -> aim += amount.toInt()
                "up" -> aim -= amount.toInt()
                "forward" -> {
                    horiz += amount.toInt()
                    depth += aim * amount.toInt()
                }
            }
        }
        .also { println("sol2: ${horiz * depth}") }
}

day2Sol1()
day2Sol2()