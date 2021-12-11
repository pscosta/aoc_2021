import java.io.File
import java.util.*

private val lines = File("/in/input10.txt").readLines()

private val syntax = mapOf('(' to ')', '[' to ']', '{' to '}', '<' to '>')
private val illegalPoints = mapOf(')' to 3, ']' to 57, '}' to 1197, '>' to 25137)
private val legalPoints = mapOf(')' to 1, ']' to 2, '}' to 3, '>' to 4)

fun sol1() = lines.sumOf { l ->
    val parsed = LinkedList<Char>()
    l.sumOf { c ->
        when {
            c in syntax.keys -> (0).also { parsed.addLast(c) }
            c != syntax[parsed.removeLast()] -> illegalPoints[c]
            else -> 0
        }
    }
}

fun sol2() = lines.mapNotNull {
        val parsed = LinkedList<Char>()
        for (c in it) when {
            c in syntax.keys -> parsed.addLast(syntax[c])
            c != parsed.removeLast() -> return@mapNotNull null
        }
        parsed.foldRight(0L) { c, partial -> (legalPoints[c]!! + partial * 5) }
    }
    .sorted()
    .let { it[it.size / 2] }

sol1()
sol2()