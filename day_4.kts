import java.io.File

fun day4() {
    val input = File("in/input4.txt").readLines()
    val drawNums = input.first().split(",").map { it.toInt() }
    var boards = input.drop(2).filterNot { it.isEmpty() }
        .chunked(5)
        .map { it.map { it.trim().replace("  ", " ").split(" ").map { it.toInt() } } }

    val drawnNums = mutableListOf<Int>()

    run sol1@{
        drawNums.forEach { drawn ->
            drawnNums.add(drawn)
            boards.firstOrNull { it.bingo(drawnNums) }?.also {
                println("sol1: ${it.score(drawn, drawnNums)}").also { return@sol1 }
            }
        }
    }
    run sol2@{
        drawNums.forEach { drawn ->
            drawnNums.add(drawn)
            when (boards.size == 1 && boards[0].bingo(drawnNums)) {
                true -> println("sol2: ${boards[0].score(drawn, drawnNums)}").also { return@sol2 }
                else -> boards = boards.filterNot { it.bingo(drawnNums) }
            }
        }
    }
}

fun List<List<Int>>.transpose() = (0 until this[0].size).map { this.map { l -> l[it] } }
fun List<List<Int>>.score(winner: Int, nums: List<Int>) = this.sumOf { it.filterNot { nums.contains(it) }.sum() } * winner
fun List<List<Int>>.bingo(nums: List<Int>) = this.any { it.marked(nums) } || this.transpose().any { it.marked(nums) }
fun List<Int>.marked(nums: List<Int>) = nums.containsAll(this)