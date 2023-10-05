package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"bufio"
	"strconv"
)

func main() {
	// ディレクトリリストのファイル名
    dirFilename := "file.txt"

	// 検索文字列リストのファイル名
    wordFilename := "name.txt"

	// ディレクトリリストファイルを読み込む
	dirlist := readfile(dirFilename)

	// 検索文字列リストファイルを読み込む
	wordlist := readfile(wordFilename)

	// ディレクトリ配下のファイル一覧を取得する
	filelist := getDirFile(dirlist)

	// ファイル一覧を読み込みキーワードにマッチするかどうか検証
	for _, value := range filelist {
		for _, word := range wordlist {
	    	if strings.Contains(strings.ToLower(value), strings.ToLower(word)) == true {
		    	fmt.Println(word + "\t" + strconv.FormatInt(getFileSize(value), 10) + "\t" + value)
		    	break
			}
		}
	}
}

// テキストファイルを1行ずつ読み込む
func readfile(filename string) []string {
	var strlist []string
    // ファイルオープン
    fp, ferr := os.Open(filename)
    if ferr != nil {
        // エラー処理
        fmt.Println(ferr)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    for scanner.Scan() {
        // ここで一行ずつ処理
        strlist = append(strlist, scanner.Text())
    }

    if ferr = scanner.Err(); ferr != nil {
        // エラー処理
        fmt.Println(ferr)
    }
    return strlist
}

// ディレクトリとサブディレクトリからファイル一覧を取得する
func getDirFile(pathlist []string) []string {
	var filelist []string
	count := 0

	for _, path := range pathlist {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				filelist = append(filelist, path)
				count++
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}

		// ディレクトリ毎の件数を表示
		fmt.Println(path + " => " + strconv.Itoa(count))
		count = 0;
	}

	// ファイル一覧を返却
    return filelist
}

// ファイルサイズを取得する
func getFileSize(path string) int64 {
    // ファイルパスからファイル情報取得
    fileinfo, staterr := os.Stat(path)

    if staterr != nil {
        fmt.Println(staterr)
        return -1
    }
 
    // ファイルサイズを返却
    return fileinfo.Size()
}
