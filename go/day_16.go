package main

import (
	"bufio"
	"fmt"
	"math"
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

var packets []packet

func readInput() {
	var bin []string // binary slice
	pc := 0          // program counter
	file, _ := os.Open("input16.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		for _, hex := range str.Split(scanner.Text(), "") {
			bin = append(bin, hexToBits(hex)...)
		}
	}
	for pc < len(bin) {
		packets = append(packets, parsePacket(&pc, bin))
	}
}

func parsePacket(pc *int, bin []string) packet {
	var pkt packet
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
	prefix := -1

	for ; prefix != 0; *pc += 4 {
		prefix = parseInt(1, pc, bin)
		data = append(data, bin[(*pc):(*pc)+4]...)
	}
	return packet{data: parseInt(len(data), new(int), data)}
}

func parseInt(len int, pc *int, bin []string) int {
	res, _ := s.ParseInt(str.Join(bin[(*pc):(*pc)+len], ""), 2, 64)
	*pc += len
	return int(res)
}

func hexToBits(val string) []string {
	var bits []string
	bit, _ := s.ParseUint(val, 16, 32)
	for i := 0; i < 4; i++ {
		bits = append([]string{s.FormatUint(bit&0x1, 2)}, bits...)
		bit = bit >> 1
	}
	return bits
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
		sum := 0
		for _, sp := range pkt.subPackets {
			sum += eval(sp)
		}
		return sum
	case 1:
		prod := 1
		for _, sp := range pkt.subPackets {
			prod *= eval(sp)
		}
		return prod
	case 2:
		min := math.MaxInt
		for _, sp := range pkt.subPackets {
			if it := eval(sp); it < min {
				min = it
			}
		}
		return min
	case 3:
		max := 0
		for _, sp := range pkt.subPackets {
			if it := eval(sp); it > max {
				max = it
			}
		}
		return max
	case 4:
		return pkt.data
	case 5:
		if eval(pkt.subPackets[0]) > eval(pkt.subPackets[1]) {
			return 1
		}
		return 0
	case 6:
		if eval(pkt.subPackets[0]) < eval(pkt.subPackets[1]) {
			return 1
		}
		return 0
	case 7:
		if eval(pkt.subPackets[0]) == eval(pkt.subPackets[1]) {
			return 1
		}
		return 0
	default:
		return -1
	}
}

func main() {
	readInput()
	fmt.Printf("Sol1: %v\n", sumVersions(packets[0]))
	fmt.Printf("Sol2: %v\n", eval(packets[0]))
}
