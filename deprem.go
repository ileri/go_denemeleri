/* SON DEPREMLER BİLGİSİ
BOĞAZİÇİ ÜNİVERSİTESİ KANDİLLİ RASATHANESİ VE DEPREM ARAŞTIRMA ENSTİTÜSÜ (KRDAE)
BÖLGESEL DEPREM-TSUNAMİ İZLEME VE DEĞERLENDİRME MERKEZİ (BDTİM) 'nin
http://www.koeri.boun.edu.tr/scripts/lst0.asp adresinde yayınladığı deprem
bilgilerinin HTML olarak çekilmesi ve REGEX KULLANMADAN çözümlenmesi
İlk aşamada amaç sadece çözümlemek olduğu için strin olarak kullanılmıştır.
*/

package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "strings"
)
type deprem struct {
  tarih string
  saat string
  enlem string
  boylam string
  derinlik string
  md string
  ml string
  mw string
  yer string
}
var INDEX int
var SAYI int
func main() {
  fmt.Printf("Gösterilecek son deprem sayısı: ")
  fmt.Scanf("%d", &SAYI)
  var depremler [2000]deprem
  adres:="http://www.koeri.boun.edu.tr/scripts/lst0.asp"
  kod := kod_oku(adres)
  INDEX= baslangic_indexi(kod)
  satir := satir_oku(kod)
  depremler[0] = deprem_cozumle(satir)
  deprem_cozumle(satir)
  for sayi:=0; sayi < SAYI-1 ; sayi++ {
    yeni_satir:= satir_oku(kod)
    depremler[sayi+1] = deprem_cozumle(yeni_satir)
  }
  for k:=0; k<SAYI; k++ {
    fmt.Printf("%s bölgesinde %s şiddetinde deprem oldu.\n",depremler[k].yer,
      depremler[k].ml)
  }


}
func kod_oku(site_adresi string)(kod string){
  res, err := http.Get(site_adresi)
  if err != nil {
          fmt.Println("http.Get", err)
          return
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
          fmt.Println("ioutil.ReadAll", err)
          return
  }
  return string(body)
}
func baslangic_indexi(metin string)(index int){
  baslangic_deseni:="---------- --------  --------  -------   ----------    ------------    --------------                                  --------------"
  return strings.Index(metin, baslangic_deseni)+len(baslangic_deseni)+2

}
func satir_oku(metin string)(string){
  var tampon string
  for ; string(metin[INDEX]) != "\n" ; INDEX++{
    tampon += string(metin[INDEX])
  }
  INDEX += 1 // Bir alt satıra 1 karakter sonra geçiliyor
  return tampon
}
func bosluk_sil(metin string)(string){
  var tampon string
  for i:=0; i<len(metin) ; i++ {
    if string(metin[i]) == " " {
      tampon += "|"
      for ; string(metin[i]) == " "; i++ {}
      tampon += string(metin[i])
    }else{
      tampon += string(metin[i])
    }
  }
  return tampon
}

func deprem_cozumle(metin string)(deprem){
  var yeni deprem
  dizi:=strings.Split(bosluk_sil(metin), "|")
  yeni.tarih = dizi[0]
  yeni.saat = dizi[1]
  yeni.enlem = dizi[2]
  yeni.boylam = dizi[3]
  yeni.derinlik = dizi[4]
  yeni.md = dizi[5]
  yeni.ml = dizi[6]
  yeni.mw = dizi[7]
  yeni.yer = dizi[8]
  return yeni
}
