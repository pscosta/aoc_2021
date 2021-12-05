import java.io.File
import kotlin.math.abs

data class Line(val x1: Int, val y1: Int, val x2: Int, val y2: Int) {
    val xxs = setOf(x1, x2).sorted()
    val yys = setOf(y1, y2).sorted()
}

val lines = File("/Users/pco38/Library/Application Support/JetBrains/IntelliJIdea2021.2/scratches/2021/in/input5.txt").readLines()
    .mapNotNull { """(\d+),(\d+) -> (\d+),(\d+)""".toRegex().find(it)?.destructured }
    .map { (x1, y1, x2, y2) -> Line(x1.toInt(), y1.toInt(), x2.toInt(), y2.toInt()) }

val maxX = lines.maxOf { maxOf(it.x1, it.x2) }
val maxY = lines.maxOf { maxOf(it.y1, it.y2) }

fun day5() {
    println("Sol1: " + intersections(false, Array(maxY + 1) { IntArray(maxX + 1) }))
    println("Sol2: " + intersections(true, Array(maxY + 1) { IntArray(maxX + 1) }))
}

fun intersections(diagonals: Boolean, matrix: Array<IntArray>): Int {
    lines.forEach { with(it) {
        when {
            x1 == x2 -> for (y in yys[0]..yys[1]) matrix[y][x1] += 1
            y1 == y2 -> for (x in xxs[0]..xxs[1]) matrix[y1][x] += 1
            else -> if (diagonals) {
                var x = x1
                var y = y1
                for (i in 0..abs(x1 - x2)) {
                    matrix[y][x] += 1
                    if (x2 > x) ++x else --x
                    if (y2 > y) ++y else --y
                }
            }
        }
    }}
    return matrix.sumOf { it.count { it > 1 } }
}