package deasciifier

import (
	"testing"
)

func TestDeasciify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"acimasizca acelya gorunen bir sacmaliktansa acilip sacilmak...", "acımasızca açelya görünen bir saçmalıktansa açılıp saçılmak..."},
		{"Acisindan bagirip cagirarak sacma sozler soylemek.", "Acısından bağırıp çağırarak saçma sözler söylemek."},
		{"Bogurtuler opucukler.", "Böğürtüler öpücükler."},
		{"BUYUKCE BIR TOPAC TOPARLAGI VE DE YUMAGI yumagi.", "BÜYÜKÇE BİR TOPAÇ TOPARLAĞI VE DE YUMAĞI yumağı."},
		{"Bilgisayarlarda uc adet bellek turu bulunur. Islemci icerisinde yer alan yazmaclar, son derece hizli ancak cok sinirli hafizaya sahiptirler. Islemcinin cok daha yavas olan ana bellege olan erisim gereksinimini gidermek icin kullanilirlar. Ana bellek ise Rastgele erisimli bellek (REB veya RAM, Random Access Memory) ve Salt okunur bellek (SOB veya ROM, Read Only Memory) olmak uzere ikiye ayrilir. RAM'a istenildigi zaman yazilabilir ve icerigi ancak guc surdugu surece korunur. ROM'sa sâdece okunabilen ve onceden yerlestirilmis bilgiler icerir. Bu icerigi gucten bagimsiz olarak korur. Ornegin herhangi bir veri veya komut RAM'da bulunurken, bilgisayar donanimini duzenleyen BIOS ROM'da yer alir.", "Bilgisayarlarda üç adet bellek turu bulunur. İşlemci içerisinde yer alan yazmaçlar, son derece hızlı ancak çok sınırlı hafızaya sahiptirler. İşlemcinin çok daha yavaş olan ana bellege olan erişim gereksinimini gidermek için kullanılırlar. Ana bellek ise Rastgele erişimli bellek (REB veya RAM, Random Access Memory) ve Salt okunur bellek (SOB veya ROM, Read Only Memory) olmak üzere ikiye ayrılır. RAM'a istenildiği zaman yazılabilir ve içeriği ancak güç sürdüğü sürece korunur. ROM'sa sâdece okunabilen ve önceden yerleştirilmiş bilgiler içerir. Bu içeriği güçten bağımsız olarak korur. Örneğin herhangi bir veri veya komut RAM'da bulunurken, bilgisayar donanımını düzenleyen BİOS ROM'da yer alır."},
		{"dunyanin en guzel sehri istanbul", "dünyanın en güzel şehri istanbul"},
	}

	for _, test := range tests {
		result := Deasciify(test.input)
		if result != test.expected {
			t.Errorf("Deasciify(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}
