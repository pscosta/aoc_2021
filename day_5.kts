import java.io.File
import kotlin.math.abs

data class Line(val x1: Int, val y1: Int, val x2: Int, val y2: Int) {
    val xxs = setOf(x1, x2).sorted()
    val yys = setOf(y1, y2).sorted()
}

val lines = File("/in/input5.txt").readLines()
    .mapNotNull { """(\d+),(\d+) -> (\d+),(\d+)""".toRegex().find(it)?.destructured }
    .map { (x1, y1, x2, y2) -> Line(x1.toInt(), y1.toInt(), x2.toInt(), y2.toInt()) }

val maxX = lines.maxOf { maxOf(it.x1, it.x2) }
val maxY = lines.maxOf { maxOf(it.y1, it.y2) }
val matrix1 = Array(maxX) { IntArray(maxX + 1) { 0 } }
val matrix2 = Array(maxX) { IntArray(maxX + 1) { 0 } }

fun day5() {
    intersections(diagonals = false, matrix1)
    intersections(diagonals = true, matrix2)
    println("Sol1: " + matrix1.sumOf { it.count { it > 1 } })
    println("Sol2: " + matrix2.sumOf { it.count { it > 1 } })
}

fun intersections(diagonals: Boolean, matrix: Array<IntArray>) = lines.forEach { l ->
    when {
        l.x1 == l.x2 -> for (y in l.yys[0]..l.yys[1]) matrix[y][l.x1] += 1
        l.y1 == l.y2 -> for (x in l.xxs[0]..l.xxs[1]) matrix[l.y1][x] += 1
        else -> if (diagonals) {
            var x = l.x1
            var y = l.y1
            for (it in 0..abs(l.x1 - l.x2)) {
                matrix[y][x] += 1
                if (l.x2 > x) ++x else --x
                if (l.y2 > y) ++y else --y
            }
        }
    }
}
