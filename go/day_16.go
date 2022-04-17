package main

import (
	"bufio"
	"fmt"
	. "github.com/pscosta/go-strm/strm"
	"os"
	s "strconv"
	str "strings"
)

type packet struct {
	version    int
	typeId     int
	data       int
	subPackets []packet
}

var bin []string // binary slice
var pc int       // program counter

func readInput() packet {
	file, _ := os.Open("input16.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, hex := range str.Split(scanner.Text(), "") {
			bin = append(bin, hexToBits(hex)...)
		}
	}
	// parses the outermost packet
	return parsePacket(&pc, bin)
}

func parsePacket(pc *int, bin []string) (pkt packet) {
	version := parseInt(3, pc, bin)
	typeId := parseInt(3, pc, bin)

	switch typeId {
	case 4:
		pkt = parseLiteral(pc, bin)
	default:
		switch parseInt(1, pc, bin) { // lengthTypeId
		case 0:
			pkt = parseLengthOperator(pc, bin)
		case 1:
			pkt = parseCountOperator(pc, bin)
		}
	}
	pkt.version, pkt.typeId = version, typeId
	return pkt
}

func parseCountOperator(pc *int, bin []string) packet {
	var subPackets []packet
	subPacketCount := parseInt(11, pc, bin)

	for i := 0; i < subPacketCount; i++ {
		subPackets = append(subPackets, parsePacket(pc, bin))
	}
	return packet{subPackets: subPackets}
}

func parseLengthOperator(pc *int, bin []string) packet {
	var subPackets []packet
	subPacketLen := parseInt(15, pc, bin)
	initPc := *pc

	for (*pc)-initPc < subPacketLen {
		packet := parsePacket(pc, bin)
		subPackets = append(subPackets, packet)
	}
	return packet{subPackets: subPackets}
}

func parseLiteral(pc *int, bin []string) packet {
	var data []string
	for prefix := -1; prefix != 0; *pc += 4 {
		prefix = parseInt(1, pc, bin)
		data = append(data, bin[*pc:*pc+4]...)
	}
	return packet{data: parseInt(len(data), new(int), data)}
}

func parseInt(len int, pc *int, bin []string) int {
	res, _ := s.ParseInt(str.Join(bin[*pc:*pc+len], ""), 2, 64)
	*pc += len // inc program counter with the read length
	return int(res)
}

func hexToBits(val string) (bits []string) {
	bit, _ := s.ParseUint(val, 16, 32)
	for i := 0; i < 4; i++ {
		bits = append([]string{s.FormatUint(bit&0x1, 2)}, bits...)
		bit = bit >> 1
	}
	return
}

func sumVersions(pkt packet) (acc int) {
	if len(pkt.subPackets) == 0 {
		return pkt.version
	} else {
		for _, sp := range pkt.subPackets {
			acc += sumVersions(sp)
		}
		return acc + pkt.version
	}
}

func eval(pkt packet) int {
	switch pkt.typeId {
	case 0:
		return From(pkt.subPackets).SumBy(eval)
	case 1:
		return Reduce(From(pkt.subPackets), func(acc int, p packet) int { return acc * eval(p) }, 1)
	case 2:
		return Min(Map(From(pkt.subPackets), eval))
	case 3:
		return Max(Map(From(pkt.subPackets), eval))
	case 4:
		return pkt.data
	case 5:
		return toInt(eval(pkt.subPackets[0]) > eval(pkt.subPackets[1]))
	case 6:
		return toInt(eval(pkt.subPackets[0]) < eval(pkt.subPackets[1]))
	case 7:
		return toInt(eval(pkt.subPackets[0]) == eval(pkt.subPackets[1]))
	}
	return -1
}

func toInt(pred bool) int {
	if pred {
		return 1
	}
	return 0
}

func main() {
	outerPacket := readInput()
	fmt.Printf("Sol1: %v\n", sumVersions(outerPacket))
	fmt.Printf("Sol2: %v\n", eval(outerPacket))
}
