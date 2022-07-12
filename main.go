package main

import (
	"debug/elf"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	elfFile, err := elf.Open(os.Args[1])
	if err != nil {
		panic("could not open elf")
	}
	defer elfFile.Close()
	for _, s := range elfFile.Sections {
		if s.Flags&elf.SHF_COMPRESSED == 0 {
			continue
		}

		reader := s.Open()
		var compression uint32
		if err := binary.Read(reader, binary.LittleEndian, &compression); err != nil {
			continue
		}
		fmt.Println("== section name:", s.Name, "compression:", compression)

		if elf.CompressionType(compression) != elf.COMPRESS_ZLIB {
			fmt.Println("!!!")
		}
	}
}
