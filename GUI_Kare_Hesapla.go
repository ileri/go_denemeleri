package main
import (
	"fmt"
	//"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"strconv"
)

/*
________________________________________________
|0         BASLIK - PencereOlustur()     _ [] X|
|______________________________________________|
|            MENU - MenuOlustur()              |
|______________________________________________|
|            CERCEVE CerceveOlustur()          |
|  __________________________________________  |
| |     Alt Alta Düzen AltAltaDuzenOlustur() | |
| |                                          | |
| |  _Kutu_________________________________  | |
| | |                                      | | |
| | |                                      | | |
| | |                                      | | |
| | |            KutuOlustur()             | | |
| | |                                      | | |
| | |                                      | | |
| | |______________________________________| | |
| |__________________________________________| |
|______________________________________________|

*/
func main(){
	pencere := PencereOlustur()
	menubar := MenuOlustur()
	cerceve := CerceveOlustur(menubar)
	pencere.Add(cerceve)
	pencere.ShowAll()
	gtk.Main()
}

func PencereOlustur() *gtk.Window {
	//--------------------------------------------------------
	// Pencere oluşturuluyor
	//--------------------------------------------------------
	gtk.Init(nil) //GTK yı başlatıyoruz
	pencere := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //Pencere oluşturuyoruz
	pencere.SetPosition(gtk.WIN_POS_CENTER) //Penceremizin açılış konumunu belirliyoruz
	pencere.SetTitle("Serhat Celil İLERİ") //Pencere BAŞLIĞI
	pencere.SetSizeRequest(300, 300)
	return pencere
}

func MenuOlustur() *gtk.MenuBar {
	//--------------------------------------------------------
	// Menü Oluşturuluyor
	//--------------------------------------------------------
	var menuitem *gtk.MenuItem
	menubar := gtk.NewMenuBar()
	//--------------------------------------------------------
	// Menü Eleman Ekleme
	//--------------------------------------------------------
	cascademenu := gtk.NewMenuItemWithMnemonic("_Dosya")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.NewMenuItemWithMnemonic("Ç_ıkış")
	menuitem.Connect("activate", func() {
			gtk.MainQuit()
	})
	submenu.Append(menuitem)
	return menubar
}

func CerceveOlustur(menubar *gtk.MenuBar) *gtk.VBox {
	altalta_kutu := AltAltaDuzenOlustur()
	cerceve := gtk.NewVBox(false, 1)
	cerceve.PackStart(menubar, false, false, 0)
	cerceve.Add(altalta_kutu)
	return cerceve
}

func AltAltaDuzenOlustur() *gtk.VPaned{
	//--------------------------------------------------------
	// AltAlta Kutu Düzeni Ekliyoruz
	//--------------------------------------------------------
	altalta_kutu := gtk.NewVPaned()
	ic_cerceve1 := KutuOlustur()
	altalta_kutu.Pack1(ic_cerceve1, false, false)
	return altalta_kutu
}

func KutuOlustur() *gtk.Frame{
	//--------------------------------------------------------
	// İç Çerçeve Ekliyoruz
	//--------------------------------------------------------
	ic_cerceve1 := gtk.NewFrame("Hesapla")
	//--------------------------------------------------------
	// Çerçeve İçindekiler
	//--------------------------------------------------------
	girdi := gtk.NewEntry()
	girdi.SetText("")
	buton := gtk.NewButtonWithLabel("Karesini Hesapla")
	yazi := gtk.NewLabel("-")
	yazi.ModifyFontEasy("DejaVu Serif 15")
	buton.Clicked(func() {
		girilen:=girdi.GetText()
		sayi, err := strconv.Atoi(girilen)
		if err != nil {
			fmt.Println(strconv.Atoi, err)
			yazi.SetText("Lütfen bir sayı girin.")
		}else{
			karesi := sayi*sayi
			yazi.SetText(strconv.Itoa(karesi))
		}
	})

	cerceve_ici_kutu1 := gtk.NewVBox(false, 1)
	cerceve_ici_kutu1.Add(girdi)
	cerceve_ici_kutu1.Add(buton)
	cerceve_ici_kutu1.Add(yazi)
	ic_cerceve1.Add(cerceve_ici_kutu1)
	return ic_cerceve1
}
