package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
    Name string
    Value string
}

var requiredFields = []string{
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
}

func valid(p []field) bool {
    if len(p) < 7 {
        return false
    }
    for _, field := range requiredFields {
        if !hasField(p, field) {
            return false
        } else {
            f := getField(p, field)
            if !validField(f) {
                return false
            }
        }
    }
    return true
}

func getField(p []field, name string) field {
    for _, f := range p {
        if f.Name == name {
            return f
        }
    }
    return field{"", ""}
}

func validYear(y string, l int, u int) bool {
    year, err := strconv.Atoi(y)
    if err != nil {
        return false
    }
    return year >= l && year <= u
}

func validHeight(v string) bool {
    if strings.HasSuffix(v, "in") {
        s := strings.TrimSuffix(v, "in")
        i, err := strconv.Atoi(s)
        if err != nil {
            return false
        }
        return i >= 59 && i <= 76
    } else if strings.HasSuffix(v, "cm") {
        s := strings.TrimSuffix(v, "cm")
        i, err := strconv.Atoi(s)
        if err != nil {
            return false
        }
        return i >= 150 && i <= 193
    }
    return false
}

func validHairColor(v string) bool {
    match, err := regexp.MatchString("^#[a-f0-9]{6}$", v)
    if err != nil {
        return false
    }
    return match
}

func validEyecolor(v string) bool {
    validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
        for _, c := range validColors {
            if c == v {
                return true
            }
        }
    return false
}

func validField(f field) bool {
    switch f.Name {
    case "byr":
        return validYear(f.Value, 1920, 2020)
    case "iyr":
        return validYear(f.Value, 2010, 2020)
    case "eyr":
        return validYear(f.Value, 2020, 2030)
    case "hgt":
        return validHeight(f.Value)
    case "hcl":
        return validHairColor(f.Value)
    case "ecl":
        return validEyecolor(f.Value)
    case "pid":
        match, err := regexp.MatchString("^[0-9]{9}$", f.Value)
        if err != nil {
            return false
        }
        return match
    case "cid":
        return true
    default:
        return false
    }
}

func hasField(p []field, f string) bool {
    for _, field := range p {
        if f == field.Name {
            return true
        }
    }
    return false
}

func readInput(file string) []string {
    dat, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
    text := string(dat)
    return strings.Split(text, "\n\n")
}

func parseLine(line string) []field {
    fields := make([]field, 0)
    fieldsText := strings.Split(strings.ReplaceAll(line, "\n", " "), " ")
    for _, ft := range fieldsText {
        kv := strings.Split(ft, ":")
        f := field{}
        f.Name = kv[0]
        f.Value = kv[1]
        fields = append(fields, f)
    }
    return fields
}


func main(){
    lines := readInput("input")
    total := 0
    validPassports := 0
    for _, l := range(lines) {
        passport := parseLine(l)
        total++
        v := valid(passport)
        if v {
            validPassports++
        }
    }
    fmt.Printf("total passports: %d\n", total)
    fmt.Printf("valid passports: %d\n", validPassports)
    // 185 is too high
    // 183 is too low
    // 184 must be correct
} 