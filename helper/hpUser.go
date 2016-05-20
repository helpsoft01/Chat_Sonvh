package helper

import (
	"fmt"
	user "../users"
	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeUser(b *flatbuffers.Builder, name []byte, id uint64) []byte {

	//re-use the already-allocated builder
	b.Reset()

	// create the name object and get its offset
	name_posstion := b.CreateByteString(name)

	//write  the user object
	user.UserStart(b)
	user.UserAddName(b, name_posstion)
	user.UserAddId(b, id)
	user_position := user.UserEnd(b)

	b.Finish(user_position)

	return b.Bytes[b.Head():]

}
func ReadUser(buf []byte) (name []byte, id uint64) {
	// initialize a user reader from the give buffer
	user := user.GetRootAsUser(buf, 0)

	// point the name variable to the bytes containing the encoded name
	name = user.Name()
	// copy the user's id(since this is just a unit64)
	id = user.Id()

	fmt.Println("name:", name, "-id:", id)
	return
}