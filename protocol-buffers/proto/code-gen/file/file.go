package file

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func WriteFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatal("Can't serialize!")
		return
	}

	if err = os.WriteFile(fname, out, 0644); err != nil {
		log.Fatal("Can't write to file!")
		return
	}

	fmt.Println("Data written!")
}

func ReadFile(fname string, pb proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatal("Can't read!")
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatal("Can't unmarshal!")
		return
	}

	fmt.Println(pb)
}

func ToJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	//out, err := protojson.Marshal(pb)
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatal("Can't read!")
		return ""
	}

	return string(out)
}

func FromJSON(in string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Can't unmarshal", err)
	}
}
