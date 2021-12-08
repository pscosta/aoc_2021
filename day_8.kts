data class Display(val patterns: List<Set<Char>>, val output: List<String>)
fun regexFor(vararg pos: Set<Char>) = (pos.joinToString("") { "(?=.*$it)" } + ".+").toRegex()

val displays = java.io.File("/in/input8.txt")
    .readLines()
    .map { it.split(" | ") }
    .map { Display(it[0].split(" ").map { it.toSet() }, it[1].split(" ")) }

fun day8() {
    val sol2 = displays.sumOf { display ->
        lateinit var `1`: Set<Char>
        lateinit var `7`: Set<Char>
        lateinit var `4`: Set<Char>

        for (pattern in display.patterns) when (pattern.size) {
            3 -> `7` = pattern
            4 -> `4` = pattern
            2 -> `1` = pattern
        }

        val `length=5` = display.patterns.filter { it.size == 5 }.map { it.toSet() }
        val `2` = `length=5`.first { (it + `4`).size == 7 }

        val top = `7` - `1`
        val middle = `2`.intersect(`4` - `1`)
        val down = `length=5`.first { ((it - `4`) - top).size == 1 }.also { (it - `4`) - top }
        val tLeft = (`4` - `1`) - `2`
        val tRight = `1`.intersect(`2`)
        val dLeft = ((`2` - `4`) - top) - down
        val dRight = `1` - tRight

        val zero = regexFor(down, tRight, dLeft, tLeft, dRight, top)
        val two = regexFor(dLeft, down, top, middle, tRight)
        val three = regexFor(middle, dRight, down, tRight, top)
        val five = regexFor(down, top, middle, dRight, tLeft)
        val six = regexFor(down, top, middle, dRight, tLeft, dLeft)
        val nine = regexFor(down, tLeft, middle, tRight, dRight, top)

        display.output.joinToString("") { with(it) {
            when {
                length == 2 -> "1"
                length == 4 -> "4"
                length == 3 -> "7"
                length == 7 -> "8"
                zero.matches(it) -> "0"
                two.matches(it) -> "2"
                three.matches(it) && length == 5 -> "3"
                five.matches(it) && length == 5 -> "5"
                six.matches(it) -> "6"
                nine.matches(it) -> "9"
                else -> throw RuntimeException(it)
            }}
        }.toInt()
    }

    println("Sol1: " + displays.flatMap { it.output }.count { it.length in setOf(2, 3, 4, 7) })
    println("Sol2: $sol2")
}