//Package fpcalc is Go wrapper for Chromaprint Library fpcalc
package fpcalc

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// FPCalc holds the structure for fingerprint metadata
type FPCalc struct {
	Timestamp float64 `json:"timestamp"`
	Duration  float64 `json:"duration"`
	FP        []int   `json:"fingerprint"`
	FPString  string  `json:"fingerprintString"`
}

// NewFPCalc returns FPCalc object
func NewFPCalc() *FPCalc {
	return &FPCalc{}
}

// GetFileFP calls the fpcalc application with the media file and media length in seconds
// returns the updated FPCalc metadata with fingerprints and updated duration
func (fp *FPCalc) GetFileFP(f string, l int) {

	log.Println("Opening " + f)

	// _ = l
	length := strconv.Itoa(l)
	// fpcalc -ts -overlap -raw -length <l> -json <f>
	//cmd := exec.Command("fpcalc", "-ts", "-overlap", "-raw", "-chunk", chunk, "-length", length, "-json", *wavPtr)

	// cmd := exec.Command("fpcalc", "-ts", "-overlap", "-raw", "-length", length, "-json", f)
	cmd := exec.Command("fpcalc", "-ts", "-overlap", "-raw", "-length", length, "-json", f)
	// cmd := exec.Command("fpcalc", "-ts", "-overlap", "-raw", "-json", f)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Run()
	// err := cmd.Run()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	t := stdout.Bytes()
	s := string(t[:])

	for {
		// look for {
		t1 := strings.Index(s, "{")
		if t1 == -1 {
			break
		}
		// look for }
		t2 := strings.Index(s, "}") + 1

		cut := s[t1:t2]
		err := json.Unmarshal([]byte(cut), &fp)
		if err != nil {
			log.Println(err)
		}

		//enc := b64.URLEncoding.EncodeToString([]byte(arrayToString(o.Fingerprint, " ")))
		//l := &Lookup{Fingerprint: enc, Timestamp: o.Timestamp}

		//c := NewBasicAuthClient("arvin", "arvin")
		//err = c.LookupFP(l)
		//if err != nil {
		//	log.Println(err)
		//}
		s = s[t2:]
	}

	var FPs []string
	for _, i := range fp.FP {
		FPs = append(FPs, strconv.Itoa(i))
	}
	fp.FPString = "{" + strings.Join(FPs, ",") + "}"

}

// GetFileFPFloat64 calls the fpcalc application with the media file and media length in seconds
// returns the updated FPCalc metadata with fingerprints and updated duration
func (fp *FPCalc) GetFileFPFloat64(f string, l float64) {

	length := strconv.FormatFloat(l, 'f', -1, 64)
	// fpcalc -ts -overlap -raw -length <l> -json <f>
	//cmd := exec.Command("fpcalc", "-ts", "-overlap", "-raw", "-chunk", chunk, "-length", length, "-json", *wavPtr)
	cmd := exec.Command("fpcalc", "-ts", "-overlap", "-raw", "-length", length, "-json", f)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Run()

	// err := cmd.Run()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	t := stdout.Bytes()
	s := string(t[:])

	for {
		// look for {
		t1 := strings.Index(s, "{")
		if t1 == -1 {
			break
		}
		// look for }
		t2 := strings.Index(s, "}") + 1

		cut := s[t1:t2]
		err := json.Unmarshal([]byte(cut), &fp)
		if err != nil {
			log.Println(err)
		}

		//enc := b64.URLEncoding.EncodeToString([]byte(arrayToString(o.Fingerprint, " ")))
		//l := &Lookup{Fingerprint: enc, Timestamp: o.Timestamp}

		//c := NewBasicAuthClient("arvin", "arvin")
		//err = c.LookupFP(l)
		//if err != nil {
		//	log.Println(err)
		//}
		s = s[t2:]
	}

	var FPs []string
	for _, i := range fp.FP {
		FPs = append(FPs, strconv.Itoa(i))
	}
	fp.FPString = "{" + strings.Join(FPs, ",") + "}"

}

//func arrayToString(a []int, delim string) string {
//	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
//}
