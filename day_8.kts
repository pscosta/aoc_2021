data class Display(val patterns: List<Set<Char>>, val output: List<String>)
fun regexFor(vararg pos: Set<Char>) = (pos.joinToString("") { "(?=.*$it)" } + ".+").toRegex()

val displays = java.io.File("/in/input8.txt")
    .readLines()
    .map { it.split(" | ") }
    .map { Display(it[0].split(" ").map { it.toSet() }, it[1].split(" ")) }

fun sol2(): Int = displays.sumOf { display ->
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
    val mid = `2`.intersect(`4` - `1`)
    val down = `length=5`.first { ((it - `4`) - top).size == 1 }.also { (it - `4`) - top }
    val tL = (`4` - `1`) - `2`
    val tR = `1`.intersect(`2`)
    val dL = ((`2` - `4`) - top) - down
    val dR = `1` - tR

    display.output.joinToString("") { with(it) {
        when {
            length == 2 -> "1"
            length == 4 -> "4"
            length == 3 -> "7"
            length == 7 -> "8"
            regexFor(down, tR, dL, tL, dR, top).matches(it) -> "0"
            regexFor(dL, down, top, mid, tR).matches(it) -> "2"
            regexFor(mid, dR, down, tR, top).matches(it) && length == 5 -> "3"
            regexFor(down, top, mid, dR, tL).matches(it) && length == 5 -> "5"
            regexFor(down, top, mid, dR, tL, dL).matches(it) -> "6"
            regexFor(down, tL, mid, tR, dR, top).matches(it) -> "9"
            else -> throw RuntimeException(it)
        }}
    }.toInt()
}

println("Sol1: " + displays.flatMap { it.output }.count { it.length in setOf(2, 3, 4, 7) })
println("Sol2: " + sol2())
