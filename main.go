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
	raw, _ := os.Open(os.Args[1])
	defer elfFile.Close()
	for _, s := range elfFile.Sections {
		if s.Flags&elf.SHF_COMPRESSED == 0 {
			continue
		}

		//reader := s.Open()
		raw.Seek(int64(s.Offset), os.SEEK_SET)

		var compression uint32
		if err := binary.Read(raw, binary.LittleEndian, &compression); err != nil {
			continue
		}

		//reader.Seek(int64(s.Offset)+24, os.SEEK_SET)

		fmt.Println("== section name:", s.Name, "compression:", compression)

		if elf.CompressionType(compression) != elf.COMPRESS_ZLIB {
			fmt.Println("!!!")
		}
	}
}
