import java.io.File

val input = File("in/input3.txt").readLines()

fun day3Sol1() {
    val gama = (0 until input[0].length)
        .map { input.map { l -> l[it] } }
        .map { l ->
            if (l.count { it == '1' } > l.count { it == '0' }) '1' else '0'
        }

    val epsilon = gama.map { if (it == '0') '1' else '0' }

    println("Sol1: ${gama.joinToString("").toInt(2) * epsilon.joinToString("").toInt(2)}")
}

fun day3Sol2() {
    var o2 = input
    for (i in 0 until input[0].length) when (o2.size) {
        1 -> break
        else -> {
            val zeros = o2.count { it[i] == '0' }
            val ones = o2.count { it[i] == '1' }

            o2 = when {
                zeros > ones -> o2.filter { it[i] == '0' }
                else -> o2.filter { it[i] == '1' }
            }
        }
    }

    var co2 = input
    for (i in 0 until input[0].length) when (co2.size) {
        1 -> break
        else -> {
            val zeros = co2.count { it[i] == '0' }
            val ones = co2.count { it[i] == '1' }

            co2 = when {
                zeros > ones -> co2.filter { it[i] == '1' }
                else -> co2.filter { it[i] == '0' }
            }
        }
    }

    println("Sol2: ${o2[0].toInt(2) * co2[0].toInt(2)}")
}

day3Sol1()
day3Sol2()