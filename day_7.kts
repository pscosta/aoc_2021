import java.io.File
import kotlin.math.abs
import kotlin.math.pow

val crabs = File("/in/input7.txt").readText().split(",").map { it.toInt() }.sorted()

fun day7() {
    val avg = crabs.sum() / crabs.size
    val middle = crabs[crabs.size / 2]
    println("Sol1: " + crabs.sumOf { abs(middle - it) })
    println("Sol2: " + crabs.sumOf { (abs(avg - it) * abs(avg - it) + abs(avg - it)) / 2 })
}

day7()

