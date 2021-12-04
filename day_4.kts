import java.io.File

fun day4() {
    val input = File("in/input4.txt").readLines()
    val allDrawNums = input.first().split(",").map { it.toInt() }
    var boards = input.drop(2)
        .filterNot { it.isEmpty() }
        .chunked(5)
        .map { it.map { it.trim().replace("  ", " ").split(" ").map { it.toInt() } } }
        .toMutableList()

    val drawnNums = mutableListOf<Int>()

    run sol1@{
        allDrawNums.forEach { drawNum ->
            drawnNums.add(drawNum)
            boards.firstOrNull { it.isWinner(drawnNums) }?.also {
                println("sol1: ${it.score(drawNum, drawnNums)}").also { return@sol1 }
            }
        }
    }
    run sol2@{
        allDrawNums.forEach { drawNum ->
            drawnNums.add(drawNum)
            when (boards.size == 1 && boards.first().isWinner(drawnNums)) {
                true -> println("sol2: ${boards.first().score(drawNum, drawnNums)}").also { return@sol2 }
                else -> boards = boards.filterNot { it.isWinner(drawnNums) }.toMutableList()
            }
        }
    }
}

fun List<List<Int>>.transpose() = (0 until this[0].size).map { this.map { l -> l[it] } }
fun List<List<Int>>.score(winner: Int, drawnNums: List<Int>) = this.sumOf { it.filterNot { drawnNums.contains(it) }.sum() } * winner
fun List<List<Int>>.isWinner(drawnNums: List<Int>) = this.any { it.won(drawnNums) } || this.transpose().any { it.won(drawnNums) }
fun List<Int>.won(drawnNums: List<Int>) = drawnNums.containsAll(this)
