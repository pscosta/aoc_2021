import java.io.File

val initialFish = File("/in/input6.txt")
    .readText()
    .split(",").map { it.toInt() }.toMutableList()

fun MutableMap<Int, Long>.incAt(pos: Int, amount: Long) = this.set(pos, (this[pos] ?: 0) + amount)

fun lanternfish(days: Int): Long {
    val newFish = mutableMapOf<Int, Long>()
    val allFish = initialFish.groupBy { it }
        .mapValues { it.value.size.toLong() }
        .toMutableMap()

    for (day in 0 until days) {
        allFish.map { (age, fishCount) ->
            when (age) {
                0 -> newFish.incAt(6, fishCount)
                else -> newFish.incAt(age - 1, fishCount)
            }
        }
        allFish[0]?.let { newFish[8] = it }

        allFish.clear()
        newFish.forEach { allFish.incAt(it.key, it.value) }
        newFish.clear()
    }
    return allFish.values.sum()
}

fun sol1() = lanternfish(80)
fun sol2() = lanternfish(256)
