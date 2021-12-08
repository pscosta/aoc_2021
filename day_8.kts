import java.io.File

data class Display(val patterns: List<String>, val output: List<String>)

val displays = File("/in/input8.txt")
    .readLines()
    .map { it.split(" | ") }
    .map { Display(it.first().split(" "), it.last().split(" ")) }

fun day8() {
    val sol1 = displays.flatMap { it.output }.count { listOf(3, 4, 2, 7).contains(it.length) }

    val sol2 = displays.sumOf { display ->
        var one = emptySet<Char>()
        var seven = emptySet<Char>()
        var four = emptySet<Char>()

        display.patterns.forEach {
            when (it.length) {
                3 -> seven = it.toSet()
                4 -> four = it.toSet()
                2 -> one = it.toSet()
            }
        }

        val length5 = display.patterns.filter { it.length == 5 }.map { it.toSet() }
        val `4-1` = four.minus(one)
        val `2` = length5.first { it.plus(four).size == 7 }

        val topTop = seven.minus(one)
        val topDown = `2`.intersect(`4-1`)
        val topLeft = `4-1`.minus(`2`)
        val topRight = one.intersect(`2`)
        val downRight = one.minus(topRight)
        val downDown = length5.first { it.minus(four).minus(topTop).size == 1 }.also { it.minus(four).minus(topTop) }
        val downLeft = `2`.minus(four).minus(topTop).minus(downDown)

        val zero = "(?=.*${downDown})(?=.*${topRight})(?=.*${downLeft})(?=.*${topLeft})(?=.*${downRight})(?=.*${topTop}).+".toRegex()
        val two = "(?=.*${downLeft})(?=.*${downDown})(?=.*${topTop})(?=.*${topDown})(?=.*${topRight}).+".toRegex()
        val three = "(?=.*${topDown})(?=.*${downRight})(?=.*${downDown})(?=.*${topRight})(?=.*${topTop}).+".toRegex()
        val five = "(?=.*${downDown})(?=.*${topTop})(?=.*${topDown})(?=.*${downRight})(?=.*${topLeft}).+".toRegex()
        val six = "(?=.*${downDown})(?=.*${topTop})(?=.*${topDown})(?=.*${downRight})(?=.*${topLeft})(?=.*${downLeft}).+".toRegex()
        val nine = "(?=.*${downDown})(?=.*${topLeft})(?=.*${topDown})(?=.*${topRight})(?=.*${downRight})(?=.*${topTop}).+".toRegex()

        return@sumOf display.output.joinToString("") {
            when {
                it.length == 2 -> "1"
                it.length == 4 -> "4"
                it.length == 3 -> "7"
                it.length == 7 -> "8"
                zero.matches(it) -> "0"
                two.matches(it) -> "2"
                three.matches(it) && it.length == 5 -> "3"
                five.matches(it) && it.length == 5 -> "5"
                six.matches(it) -> "6"
                nine.matches(it) -> "9"
                else -> throw RuntimeException(it)
            }
        }.toInt()
    }

    println("Sol1: $sol1")
    println("Sol2: $sol2")
}

day8()
