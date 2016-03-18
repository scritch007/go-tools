package crypto

import (
	"hash"
)

// Implementation of the boringssl EVP_BytesToKey library
// nbkey is the size of the key to generate
// niv is the size of the IV
// hash is the hash method
// salt provide empty bit if you don't need a salt
func EVP_BytesToKey(nkey, niv int, hash hash.Hash, salt, password []byte, count int) (key []byte, iv []byte){
	var md_buf []byte
	key = make([]byte, nkey)
	iv = make([]byte, niv)
	mds := hash.Size()
	keyIndex := 0
	ivIndex := 0
	salt = salt[:8] // This is because it uses PKCS5_SALT_LEN

	for ;;{
		hash.Reset()
		a := make([]byte, len(md_buf)+len(password)+len(salt))
    	copy(a, md_buf)
    	copy(a[len(md_buf):], password)
    	copy(a[len(md_buf)+len(password):], salt)
        hash.Write(a)
		md_buf = hash.Sum(nil)
		for i := 1; i < count; i++{
			hash.Reset()
			hash.Write(md_buf)
			md_buf = hash.Sum(nil)
		}
		i := 0
		if 0 != nkey{
			for;;{
			if (nkey == 0 || i == mds) {
			          break
		        }
			key[keyIndex] = md_buf[i]
			keyIndex++
		        nkey--
		        i++
			}
		}
		if 0 != niv && i != len(md_buf) {
			for ;; {
				if (niv == 0 || i == mds) {
					break
				}
				iv[ivIndex] = md_buf[i]
				ivIndex++
				niv--
			        i++
			}
		}
		if (nkey == 0 && niv == 0) {
			break
		}

	}
	return key, iv
}