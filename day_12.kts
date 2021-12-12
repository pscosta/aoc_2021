import java.io.File

val paths = mutableMapOf<String, Set<String>>()
val input = File("/in/input12.txt").readLines()
    .map { it.split("-") }.forEach { (k, v) ->
        paths[k] = paths.getOrDefault(k, setOf()) + v
        paths[v] = paths.getOrDefault(v, setOf()) + k
    }

fun dfs(cave: String, visited: List<String>): List<List<String>> = when (cave) {
    "end" -> listOf(visited + cave)
    else -> paths[cave]!!
        .filter { it != "start" }
        .filter { it.isBigCave() || it !in visited + cave }
        .flatMap { dfs(it, visited + cave) }
}

fun dfs2(cave: String, visited: List<String>): List<List<String>> = when (cave) {
    "end" -> listOf(visited + cave)
    else -> paths[cave]!!
        .filter { it != "start" }
        .filter { it !in (visited + cave) || it.isBigCave() || !(visited + cave).hasSmallCavesTwice() }
        .flatMap { dfs2(it, visited + cave) }
}

fun List<String>.hasSmallCavesTwice() = this.groupingBy { it }.eachCount().any { !it.key.isBigCave() && it.value >= 2 }
fun String.isBigCave() = this[0].isUpperCase()

fun sol1() = dfs("start", listOf()).size
fun sol2() = dfs2("start", listOf()).size
