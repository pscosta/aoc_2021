import java.io.File

lateinit var `🐙's`: List<List<`🐙`>>

fun loadInput() {
    `🐙's` = File("/in/input11.txt").readLines()
        .map { it.map { "$it".toInt() }.map { energy -> `🐙`(energy) } }

    `🐙's`.forEachIndexed { y, l ->
        l.forEachIndexed { x, it -> it.neighbours = neighboursFor(x, y) }
    }
}

data class `🐙`(var energy: Int, var neighbours: List<`🐙`> = emptyList()) {
    fun step() {
        if (++energy == 10) neighbours.forEach { it.step() }
    }
}

fun main() {
    loadInput().also { println("sol1: " + sol1()) }
    loadInput().also { println("sol2: " + sol2()) }
}

fun sol1(): Int = with(`🐙's`.flatten()) {
    var count = 0
    repeat(100) {
        forEach { if (it.energy > 9) it.energy = 0 }
        forEach { it.step() }
        count += count { it.energy > 9 }
    }
    return count
}

fun sol2(): Int = with(`🐙's`.flatten()) {
    var step = 0
    while (true) {
        forEach { if (it.energy > 9) it.energy = 0 }
        forEach { it.step() }
        val flashing = count { it.energy > 9 }
        step++
        if (flashing == size) break
    }
    return step
}

fun neighboursFor(x: Int, y: Int): List<`🐙`> {
    val neighbours = ArrayList<`🐙`>()
    if (x != 0) neighbours += `🐙's`[y][x - 1]
    if (x != `🐙's`[0].size - 1) neighbours += `🐙's`[y][x + 1]

    if (y != 0) {
        if (x != 0) neighbours += `🐙's`[y - 1][x - 1]
        neighbours += `🐙's`[y - 1][x]
        if (x != `🐙's`[0].size - 1) neighbours += `🐙's`[y - 1][x + 1]
    }

    if (y != `🐙's`.size - 1) {
        if (x != 0) neighbours += `🐙's`[y + 1][x - 1]
        neighbours += `🐙's`[y + 1][x]
        if (x != `🐙's`[0].size - 1) neighbours += `🐙's`[y + 1][x + 1]
    }
    return neighbours
}
