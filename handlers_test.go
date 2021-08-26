package ResourceServer

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// requestAndCompare request a resource and test again expected result
func requestAndCompare(t *testing.T, method, rURL string, code int) {
	r, err := http.NewRequest(method, rURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	New("./test_data", false).ServeHTTP(w, r)

	assert.Equal(t, code, w.Code)
	if code == 404 {
		return
	}

	pathToResource, mimeType := findMimeType("./test_data" + rURL)
	data, err := ioutil.ReadFile(pathToResource)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, data, w.Body.Bytes())
	assert.Equal(t, mimeType, w.Header().Get("Content-Type"))
}

func TestHtmlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/", 200)
	requestAndCompare(t, "GET", "/second.html", 200)
	requestAndCompare(t, "GET", "/sub", 200)
	requestAndCompare(t, "GET", "/sub/", 200)
	requestAndCompare(t, "GET", "/not_exist.html", 404)
	requestAndCompare(t, "GET", "/not_exist", 404)
	requestAndCompare(t, "GET", "/not_exist/", 404)
}

func TestDwgRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dwg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dwg", 200)
	requestAndCompare(t, "GET", "/dummy2.dwg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dwg", 404)
}

func TestAsdRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.asd", 200)
	requestAndCompare(t, "GET", "/sub/dummy.asd", 200)
	requestAndCompare(t, "GET", "/dummy2.asd", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.asd", 404)
}

func TestAsnRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.asn", 200)
	requestAndCompare(t, "GET", "/sub/dummy.asn", 200)
	requestAndCompare(t, "GET", "/dummy2.asn", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.asn", 404)
}

func TestTspRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tsp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tsp", 200)
	requestAndCompare(t, "GET", "/dummy2.tsp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tsp", 404)
}

func TestDxfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dxf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dxf", 200)
	requestAndCompare(t, "GET", "/dummy2.dxf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dxf", 404)
}

func TestRegRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.reg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.reg", 200)
	requestAndCompare(t, "GET", "/dummy2.reg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.reg", 404)
}

func TestSplRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.spl", 200)
	requestAndCompare(t, "GET", "/sub/dummy.spl", 200)
	requestAndCompare(t, "GET", "/dummy2.spl", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.spl", 404)
}

func TestGzRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.gz", 200)
	requestAndCompare(t, "GET", "/sub/dummy.gz", 200)
	requestAndCompare(t, "GET", "/dummy2.gz", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.gz", 404)
}

func TestJsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.js", 200)
	requestAndCompare(t, "GET", "/sub/dummy.js", 200)
	requestAndCompare(t, "GET", "/dummy2.js", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.js", 404)
}

func TestJsonRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.json", 200)
	requestAndCompare(t, "GET", "/sub/dummy.json", 200)
	requestAndCompare(t, "GET", "/dummy2.json", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.json", 404)
}

func TestPtlkRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ptlk", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ptlk", 200)
	requestAndCompare(t, "GET", "/dummy2.ptlk", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ptlk", 404)
}

func TestHqxRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.hqx", 200)
	requestAndCompare(t, "GET", "/sub/dummy.hqx", 200)
	requestAndCompare(t, "GET", "/dummy2.hqx", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.hqx", 404)
}

func TestMbdRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mbd", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mbd", 200)
	requestAndCompare(t, "GET", "/dummy2.mbd", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mbd", 404)
}

func TestMifRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mif", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mif", 200)
	requestAndCompare(t, "GET", "/dummy2.mif", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mif", 404)
}

func TestXlsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.xls", 200)
	requestAndCompare(t, "GET", "/sub/dummy.xls", 200)
	requestAndCompare(t, "GET", "/dummy2.xls", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.xls", 404)
}

func TestXlaRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.xla", 200)
	requestAndCompare(t, "GET", "/sub/dummy.xla", 200)
	requestAndCompare(t, "GET", "/dummy2.xla", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.xla", 404)
}

func TestChmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.chm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.chm", 200)
	requestAndCompare(t, "GET", "/dummy2.chm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.chm", 404)
}

func TestHlpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.hlp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.hlp", 200)
	requestAndCompare(t, "GET", "/dummy2.hlp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.hlp", 404)
}

func TestPptRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ppt", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ppt", 200)
	requestAndCompare(t, "GET", "/dummy2.ppt", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ppt", 404)
}

func TestPpzRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ppz", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ppz", 200)
	requestAndCompare(t, "GET", "/dummy2.ppz", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ppz", 404)
}

func TestPpsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.pps", 200)
	requestAndCompare(t, "GET", "/sub/dummy.pps", 200)
	requestAndCompare(t, "GET", "/dummy2.pps", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.pps", 404)
}

func TestPotRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.pot", 200)
	requestAndCompare(t, "GET", "/sub/dummy.pot", 200)
	requestAndCompare(t, "GET", "/dummy2.pot", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.pot", 404)
}

func TestDocRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.doc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.doc", 200)
	requestAndCompare(t, "GET", "/dummy2.doc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.doc", 404)
}

func TestDotRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dot", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dot", 200)
	requestAndCompare(t, "GET", "/dummy2.dot", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dot", 404)
}

func TestBinRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.bin", 200)
	requestAndCompare(t, "GET", "/sub/dummy.bin", 200)
	requestAndCompare(t, "GET", "/dummy2.bin", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.bin", 404)
}

func TestFileRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.file", 200)
	requestAndCompare(t, "GET", "/sub/dummy.file", 200)
	requestAndCompare(t, "GET", "/dummy2.file", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.file", 404)
}

func TestComRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.com", 200)
	requestAndCompare(t, "GET", "/sub/dummy.com", 200)
	requestAndCompare(t, "GET", "/dummy2.com", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.com", 404)
}

func TestClassRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.class", 200)
	requestAndCompare(t, "GET", "/sub/dummy.class", 200)
	requestAndCompare(t, "GET", "/dummy2.class", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.class", 404)
}

func TestIniRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ini", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ini", 200)
	requestAndCompare(t, "GET", "/dummy2.ini", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ini", 404)
}

func TestOdaRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.oda", 200)
	requestAndCompare(t, "GET", "/sub/dummy.oda", 200)
	requestAndCompare(t, "GET", "/dummy2.oda", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.oda", 404)
}

func TestPdfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.pdf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.pdf", 200)
	requestAndCompare(t, "GET", "/dummy2.pdf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.pdf", 404)
}

func TestAiRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ai", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ai", 200)
	requestAndCompare(t, "GET", "/dummy2.ai", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ai", 404)
}

func TestEpsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.eps", 200)
	requestAndCompare(t, "GET", "/sub/dummy.eps", 200)
	requestAndCompare(t, "GET", "/dummy2.eps", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.eps", 404)
}

func TestPsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ps", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ps", 200)
	requestAndCompare(t, "GET", "/dummy2.ps", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ps", 404)
}

func TestRtcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.rtc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.rtc", 200)
	requestAndCompare(t, "GET", "/dummy2.rtc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.rtc", 404)
}

func TestRtfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.rtf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.rtf", 200)
	requestAndCompare(t, "GET", "/dummy2.rtf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.rtf", 404)
}

func TestSmpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.smp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.smp", 200)
	requestAndCompare(t, "GET", "/dummy2.smp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.smp", 404)
}

func TestTbkRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tbk", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tbk", 200)
	requestAndCompare(t, "GET", "/dummy2.tbk", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tbk", 404)
}

func TestOdcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odc", 200)
	requestAndCompare(t, "GET", "/dummy2.odc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odc", 404)
}

func TestOdfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odf", 200)
	requestAndCompare(t, "GET", "/dummy2.odf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odf", 404)
}

func TestOdgRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odg", 200)
	requestAndCompare(t, "GET", "/dummy2.odg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odg", 404)
}

func TestOdiRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odi", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odi", 200)
	requestAndCompare(t, "GET", "/dummy2.odi", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odi", 404)
}

func TestOdpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odp", 200)
	requestAndCompare(t, "GET", "/dummy2.odp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odp", 404)
}

func TestOdsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ods", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ods", 200)
	requestAndCompare(t, "GET", "/dummy2.ods", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ods", 404)
}

func TestOdtRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odt", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odt", 200)
	requestAndCompare(t, "GET", "/dummy2.odt", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odt", 404)
}

func TestOdmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.odm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.odm", 200)
	requestAndCompare(t, "GET", "/dummy2.odm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.odm", 404)
}

func TestWmlcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wmlc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wmlc", 200)
	requestAndCompare(t, "GET", "/dummy2.wmlc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wmlc", 404)
}

func TestWmlscRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wmlsc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wmlsc", 200)
	requestAndCompare(t, "GET", "/dummy2.wmlsc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wmlsc", 404)
}

func TestVmdRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.vmd", 200)
	requestAndCompare(t, "GET", "/sub/dummy.vmd", 200)
	requestAndCompare(t, "GET", "/dummy2.vmd", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.vmd", 404)
}

func TestVmfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.vmf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.vmf", 200)
	requestAndCompare(t, "GET", "/dummy2.vmf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.vmf", 404)
}

func TestBcpioRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.bcpio", 200)
	requestAndCompare(t, "GET", "/sub/dummy.bcpio", 200)
	requestAndCompare(t, "GET", "/dummy2.bcpio", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.bcpio", 404)
}

func TestZRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.z", 200)
	requestAndCompare(t, "GET", "/sub/dummy.z", 200)
	requestAndCompare(t, "GET", "/dummy2.z", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.z", 404)
}

func TestCpioRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.cpio", 200)
	requestAndCompare(t, "GET", "/sub/dummy.cpio", 200)
	requestAndCompare(t, "GET", "/dummy2.cpio", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.cpio", 404)
}

func TestCshRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.csh", 200)
	requestAndCompare(t, "GET", "/sub/dummy.csh", 200)
	requestAndCompare(t, "GET", "/dummy2.csh", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.csh", 404)
}

func TestDcrRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dcr", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dcr", 200)
	requestAndCompare(t, "GET", "/dummy2.dcr", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dcr", 404)
}

func TestDirRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dir", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dir", 200)
	requestAndCompare(t, "GET", "/dummy2.dir", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dir", 404)
}

func TestDxrRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dxr", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dxr", 200)
	requestAndCompare(t, "GET", "/dummy2.dxr", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dxr", 404)
}

func TestDviRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dvi", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dvi", 200)
	requestAndCompare(t, "GET", "/dummy2.dvi", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dvi", 404)
}

func TestEvyRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.evy", 200)
	requestAndCompare(t, "GET", "/sub/dummy.evy", 200)
	requestAndCompare(t, "GET", "/dummy2.evy", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.evy", 404)
}

func TestGtarRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.gtar", 200)
	requestAndCompare(t, "GET", "/sub/dummy.gtar", 200)
	requestAndCompare(t, "GET", "/dummy2.gtar", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.gtar", 404)
}

func TestHdfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.hdf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.hdf", 200)
	requestAndCompare(t, "GET", "/dummy2.hdf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.hdf", 404)
}

func TestPhpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.php", 200)
	requestAndCompare(t, "GET", "/sub/dummy.php", 200)
	requestAndCompare(t, "GET", "/dummy2.php", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.php", 404)
}

func TestPhtmlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.phtml", 200)
	requestAndCompare(t, "GET", "/sub/dummy.phtml", 200)
	requestAndCompare(t, "GET", "/dummy2.phtml", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.phtml", 404)
}

func TestLatexRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.latex", 200)
	requestAndCompare(t, "GET", "/sub/dummy.latex", 200)
	requestAndCompare(t, "GET", "/dummy2.latex", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.latex", 404)
}

func TestXmlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.xml", 200)
	requestAndCompare(t, "GET", "/sub/dummy.xml", 200)
	requestAndCompare(t, "GET", "/dummy2.xml", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.xml", 404)
}

func TestNcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.nc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.nc", 200)
	requestAndCompare(t, "GET", "/dummy2.nc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.nc", 404)
}

func TestCdfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.cdf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.cdf", 200)
	requestAndCompare(t, "GET", "/dummy2.cdf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.cdf", 404)
}

func TestNscRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.nsc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.nsc", 200)
	requestAndCompare(t, "GET", "/dummy2.nsc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.nsc", 404)
}

func TestShRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sh", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sh", 200)
	requestAndCompare(t, "GET", "/dummy2.sh", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sh", 404)
}

func TestSharRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.shar", 200)
	requestAndCompare(t, "GET", "/sub/dummy.shar", 200)
	requestAndCompare(t, "GET", "/dummy2.shar", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.shar", 404)
}

func TestSwfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.swf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.swf", 200)
	requestAndCompare(t, "GET", "/dummy2.swf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.swf", 404)
}

func TestCabRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.cab", 200)
	requestAndCompare(t, "GET", "/sub/dummy.cab", 200)
	requestAndCompare(t, "GET", "/dummy2.cab", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.cab", 404)
}

func TestSprRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.spr", 200)
	requestAndCompare(t, "GET", "/sub/dummy.spr", 200)
	requestAndCompare(t, "GET", "/dummy2.spr", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.spr", 404)
}

func TestSpriteRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sprite", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sprite", 200)
	requestAndCompare(t, "GET", "/dummy2.sprite", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sprite", 404)
}

func TestSitRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sit", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sit", 200)
	requestAndCompare(t, "GET", "/dummy2.sit", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sit", 404)
}

func TestScaRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sca", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sca", 200)
	requestAndCompare(t, "GET", "/dummy2.sca", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sca", 404)
}

func TestSv4cpioRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sv4cpio", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sv4cpio", 200)
	requestAndCompare(t, "GET", "/dummy2.sv4cpio", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sv4cpio", 404)
}

func TestSv4crcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sv4crc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sv4crc", 200)
	requestAndCompare(t, "GET", "/dummy2.sv4crc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sv4crc", 404)
}

func TestTarRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tar", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tar", 200)
	requestAndCompare(t, "GET", "/dummy2.tar", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tar", 404)
}

func TestTclRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tcl", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tcl", 200)
	requestAndCompare(t, "GET", "/dummy2.tcl", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tcl", 404)
}

func TestTexRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tex", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tex", 200)
	requestAndCompare(t, "GET", "/dummy2.tex", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tex", 404)
}

func TestTexinfoRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.texinfo", 200)
	requestAndCompare(t, "GET", "/sub/dummy.texinfo", 200)
	requestAndCompare(t, "GET", "/dummy2.texinfo", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.texinfo", 404)
}

func TestTexiRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.texi", 200)
	requestAndCompare(t, "GET", "/sub/dummy.texi", 200)
	requestAndCompare(t, "GET", "/dummy2.texi", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.texi", 404)
}

func TestTRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.t", 200)
	requestAndCompare(t, "GET", "/sub/dummy.t", 200)
	requestAndCompare(t, "GET", "/dummy2.t", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.t", 404)
}

func TestTrRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tr", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tr", 200)
	requestAndCompare(t, "GET", "/dummy2.tr", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tr", 404)
}

func TestRoffRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.roff", 200)
	requestAndCompare(t, "GET", "/sub/dummy.roff", 200)
	requestAndCompare(t, "GET", "/dummy2.roff", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.roff", 404)
}

func TestUstarRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ustar", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ustar", 200)
	requestAndCompare(t, "GET", "/dummy2.ustar", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ustar", 404)
}

func TestSrcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.src", 200)
	requestAndCompare(t, "GET", "/sub/dummy.src", 200)
	requestAndCompare(t, "GET", "/dummy2.src", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.src", 404)
}

func TestZipRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.zip", 200)
	requestAndCompare(t, "GET", "/sub/dummy.zip", 200)
	requestAndCompare(t, "GET", "/dummy2.zip", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.zip", 404)
}

func TestAuRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.au", 200)
	requestAndCompare(t, "GET", "/sub/dummy.au", 200)
	requestAndCompare(t, "GET", "/dummy2.au", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.au", 404)
}

func TestSndRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.snd", 200)
	requestAndCompare(t, "GET", "/sub/dummy.snd", 200)
	requestAndCompare(t, "GET", "/dummy2.snd", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.snd", 404)
}

func TestEsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.es", 200)
	requestAndCompare(t, "GET", "/sub/dummy.es", 200)
	requestAndCompare(t, "GET", "/dummy2.es", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.es", 404)
}

func TestMp3Request(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mp3", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mp3", 200)
	requestAndCompare(t, "GET", "/dummy2.mp3", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mp3", 404)
}

func TestTsiRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tsi", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tsi", 200)
	requestAndCompare(t, "GET", "/dummy2.tsi", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tsi", 404)
}

func TestVoxRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.vox", 200)
	requestAndCompare(t, "GET", "/sub/dummy.vox", 200)
	requestAndCompare(t, "GET", "/dummy2.vox", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.vox", 404)
}

func TestWavRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wav", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wav", 200)
	requestAndCompare(t, "GET", "/dummy2.wav", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wav", 404)
}

func TestAifRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.aif", 200)
	requestAndCompare(t, "GET", "/sub/dummy.aif", 200)
	requestAndCompare(t, "GET", "/dummy2.aif", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.aif", 404)
}

func TestAiffRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.aiff", 200)
	requestAndCompare(t, "GET", "/sub/dummy.aiff", 200)
	requestAndCompare(t, "GET", "/dummy2.aiff", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.aiff", 404)
}

func TestAifcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.aifc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.aifc", 200)
	requestAndCompare(t, "GET", "/dummy2.aifc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.aifc", 404)
}

func TestDusRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dus", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dus", 200)
	requestAndCompare(t, "GET", "/dummy2.dus", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dus", 404)
}

func TestChtRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.cht", 200)
	requestAndCompare(t, "GET", "/sub/dummy.cht", 200)
	requestAndCompare(t, "GET", "/dummy2.cht", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.cht", 404)
}

func TestMidRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mid", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mid", 200)
	requestAndCompare(t, "GET", "/dummy2.mid", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mid", 404)
}

func TestMidiRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.midi", 200)
	requestAndCompare(t, "GET", "/sub/dummy.midi", 200)
	requestAndCompare(t, "GET", "/dummy2.midi", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.midi", 404)
}

func TestMp2Request(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mp2", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mp2", 200)
	requestAndCompare(t, "GET", "/dummy2.mp2", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mp2", 404)
}

func TestRamRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ram", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ram", 200)
	requestAndCompare(t, "GET", "/dummy2.ram", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ram", 404)
}

func TestRaRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ra", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ra", 200)
	requestAndCompare(t, "GET", "/dummy2.ra", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ra", 404)
}

func TestRpmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.rpm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.rpm", 200)
	requestAndCompare(t, "GET", "/dummy2.rpm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.rpm", 404)
}

func TestStreamRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.stream", 200)
	requestAndCompare(t, "GET", "/sub/dummy.stream", 200)
	requestAndCompare(t, "GET", "/dummy2.stream", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.stream", 404)
}

func TestDwfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.dwf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.dwf", 200)
	requestAndCompare(t, "GET", "/dummy2.dwf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.dwf", 404)
}

func TestBmpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.bmp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.bmp", 200)
	requestAndCompare(t, "GET", "/dummy2.bmp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.bmp", 404)
}

func TestCodRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.cod", 200)
	requestAndCompare(t, "GET", "/sub/dummy.cod", 200)
	requestAndCompare(t, "GET", "/dummy2.cod", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.cod", 404)
}

func TestRasRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ras", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ras", 200)
	requestAndCompare(t, "GET", "/dummy2.ras", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ras", 404)
}

func TestFifRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.fif", 200)
	requestAndCompare(t, "GET", "/sub/dummy.fif", 200)
	requestAndCompare(t, "GET", "/dummy2.fif", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.fif", 404)
}

func TestGifRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.gif", 200)
	requestAndCompare(t, "GET", "/sub/dummy.gif", 200)
	requestAndCompare(t, "GET", "/dummy2.gif", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.gif", 404)
}

func TestIefRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ief", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ief", 200)
	requestAndCompare(t, "GET", "/dummy2.ief", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ief", 404)
}

func TestJpegRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.jpeg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.jpeg", 200)
	requestAndCompare(t, "GET", "/dummy2.jpeg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.jpeg", 404)
}

func TestJpgRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.jpg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.jpg", 200)
	requestAndCompare(t, "GET", "/dummy2.jpg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.jpg", 404)
}

func TestJpeRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.jpe", 200)
	requestAndCompare(t, "GET", "/sub/dummy.jpe", 200)
	requestAndCompare(t, "GET", "/dummy2.jpe", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.jpe", 404)
}

func TestPngRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.png", 200)
	requestAndCompare(t, "GET", "/sub/dummy.png", 200)
	requestAndCompare(t, "GET", "/dummy2.png", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.png", 404)
}

func TestSvgRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.svg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.svg", 200)
	requestAndCompare(t, "GET", "/dummy2.svg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.svg", 404)
}

func TestTiffRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tiff", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tiff", 200)
	requestAndCompare(t, "GET", "/dummy2.tiff", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tiff", 404)
}

func TestTifRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tif", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tif", 200)
	requestAndCompare(t, "GET", "/dummy2.tif", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tif", 404)
}

func TestMcfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mcf", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mcf", 200)
	requestAndCompare(t, "GET", "/dummy2.mcf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mcf", 404)
}

func TestWbmpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wbmp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wbmp", 200)
	requestAndCompare(t, "GET", "/dummy2.wbmp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wbmp", 404)
}

func TestFh4Request(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.fh4", 200)
	requestAndCompare(t, "GET", "/sub/dummy.fh4", 200)
	requestAndCompare(t, "GET", "/dummy2.fh4", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.fh4", 404)
}

func TestFh5Request(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.fh5", 200)
	requestAndCompare(t, "GET", "/sub/dummy.fh5", 200)
	requestAndCompare(t, "GET", "/dummy2.fh5", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.fh5", 404)
}

func TestFhcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.fhc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.fhc", 200)
	requestAndCompare(t, "GET", "/dummy2.fhc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.fhc", 404)
}

func TestIcoRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ico", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ico", 200)
	requestAndCompare(t, "GET", "/dummy2.ico", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ico", 404)
}

func TestPnmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.pnm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.pnm", 200)
	requestAndCompare(t, "GET", "/dummy2.pnm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.pnm", 404)
}

func TestPbmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.pbm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.pbm", 200)
	requestAndCompare(t, "GET", "/dummy2.pbm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.pbm", 404)
}

func TestPgmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.pgm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.pgm", 200)
	requestAndCompare(t, "GET", "/dummy2.pgm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.pgm", 404)
}

func TestPpmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ppm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ppm", 200)
	requestAndCompare(t, "GET", "/dummy2.ppm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ppm", 404)
}

func TestRgbRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.rgb", 200)
	requestAndCompare(t, "GET", "/sub/dummy.rgb", 200)
	requestAndCompare(t, "GET", "/dummy2.rgb", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.rgb", 404)
}

func TestXwdRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.xwd", 200)
	requestAndCompare(t, "GET", "/sub/dummy.xwd", 200)
	requestAndCompare(t, "GET", "/dummy2.xwd", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.xwd", 404)
}

func TestXbmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.xbm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.xbm", 200)
	requestAndCompare(t, "GET", "/dummy2.xbm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.xbm", 404)
}

func TestXpmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.xpm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.xpm", 200)
	requestAndCompare(t, "GET", "/dummy2.xpm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.xpm", 404)
}

func TestWrlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wrl", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wrl", 200)
	requestAndCompare(t, "GET", "/dummy2.wrl", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wrl", 404)
}

func TestIcsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ics", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ics", 200)
	requestAndCompare(t, "GET", "/dummy2.ics", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ics", 404)
}

func TestCsvRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.csv", 200)
	requestAndCompare(t, "GET", "/sub/dummy.csv", 200)
	requestAndCompare(t, "GET", "/dummy2.csv", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.csv", 404)
}

func TestCssRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.css", 200)
	requestAndCompare(t, "GET", "/sub/dummy.css", 200)
	requestAndCompare(t, "GET", "/dummy2.css", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.css", 404)
}

func TestHtmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.htm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.htm", 200)
	requestAndCompare(t, "GET", "/dummy2.htm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.htm", 404)
}

func TestShtmlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.shtml", 200)
	requestAndCompare(t, "GET", "/sub/dummy.shtml", 200)
	requestAndCompare(t, "GET", "/dummy2.shtml", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.shtml", 404)
}

func TestTxtRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.txt", 200)
	requestAndCompare(t, "GET", "/sub/dummy.txt", 200)
	requestAndCompare(t, "GET", "/dummy2.txt", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.txt", 404)
}

func TestRtxRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.rtx", 200)
	requestAndCompare(t, "GET", "/sub/dummy.rtx", 200)
	requestAndCompare(t, "GET", "/dummy2.rtx", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.rtx", 404)
}

func TestTsvRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.tsv", 200)
	requestAndCompare(t, "GET", "/sub/dummy.tsv", 200)
	requestAndCompare(t, "GET", "/dummy2.tsv", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.tsv", 404)
}

func TestWmlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wml", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wml", 200)
	requestAndCompare(t, "GET", "/dummy2.wml", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wml", 404)
}

func TestWmlsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.wmls", 200)
	requestAndCompare(t, "GET", "/sub/dummy.wmls", 200)
	requestAndCompare(t, "GET", "/dummy2.wmls", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.wmls", 404)
}

func TestEtxRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.etx", 200)
	requestAndCompare(t, "GET", "/sub/dummy.etx", 200)
	requestAndCompare(t, "GET", "/dummy2.etx", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.etx", 404)
}

func TestSgmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sgm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sgm", 200)
	requestAndCompare(t, "GET", "/dummy2.sgm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sgm", 404)
}

func TestSgmlRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.sgml", 200)
	requestAndCompare(t, "GET", "/sub/dummy.sgml", 200)
	requestAndCompare(t, "GET", "/dummy2.sgml", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.sgml", 404)
}

func TestTalkRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.talk", 200)
	requestAndCompare(t, "GET", "/sub/dummy.talk", 200)
	requestAndCompare(t, "GET", "/dummy2.talk", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.talk", 404)
}

func TestSpcRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.spc", 200)
	requestAndCompare(t, "GET", "/sub/dummy.spc", 200)
	requestAndCompare(t, "GET", "/dummy2.spc", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.spc", 404)
}

func Test3gpRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.3gp", 200)
	requestAndCompare(t, "GET", "/sub/dummy.3gp", 200)
	requestAndCompare(t, "GET", "/dummy2.3gp", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.3gp", 404)
}

func TestMp4Request(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mp4", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mp4", 200)
	requestAndCompare(t, "GET", "/dummy2.mp4", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mp4", 404)
}

func TestMpegRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mpeg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mpeg", 200)
	requestAndCompare(t, "GET", "/dummy2.mpeg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mpeg", 404)
}

func TestMpgRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mpg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mpg", 200)
	requestAndCompare(t, "GET", "/dummy2.mpg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mpg", 404)
}

func TestMpeRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mpe", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mpe", 200)
	requestAndCompare(t, "GET", "/dummy2.mpe", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mpe", 404)
}

func TestOggRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ogg", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ogg", 200)
	requestAndCompare(t, "GET", "/dummy2.ogg", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ogg", 404)
}

func TestOgvRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.ogv", 200)
	requestAndCompare(t, "GET", "/sub/dummy.ogv", 200)
	requestAndCompare(t, "GET", "/dummy2.ogv", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.ogv", 404)
}

func TestQtRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.qt", 200)
	requestAndCompare(t, "GET", "/sub/dummy.qt", 200)
	requestAndCompare(t, "GET", "/dummy2.qt", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.qt", 404)
}

func TestMovRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.mov", 200)
	requestAndCompare(t, "GET", "/sub/dummy.mov", 200)
	requestAndCompare(t, "GET", "/dummy2.mov", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.mov", 404)
}

func TestVivRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.viv", 200)
	requestAndCompare(t, "GET", "/sub/dummy.viv", 200)
	requestAndCompare(t, "GET", "/dummy2.viv", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.viv", 404)
}

func TestVivoRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.vivo", 200)
	requestAndCompare(t, "GET", "/sub/dummy.vivo", 200)
	requestAndCompare(t, "GET", "/dummy2.vivo", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.vivo", 404)
}

func TestWebmRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.webm", 200)
	requestAndCompare(t, "GET", "/sub/dummy.webm", 200)
	requestAndCompare(t, "GET", "/dummy2.webm", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.webm", 404)
}

func TestAviRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.avi", 200)
	requestAndCompare(t, "GET", "/sub/dummy.avi", 200)
	requestAndCompare(t, "GET", "/dummy2.avi", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.avi", 404)
}

func TestMovieRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.movie", 200)
	requestAndCompare(t, "GET", "/sub/dummy.movie", 200)
	requestAndCompare(t, "GET", "/dummy2.movie", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.movie", 404)
}

func TestVtsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.vts", 200)
	requestAndCompare(t, "GET", "/sub/dummy.vts", 200)
	requestAndCompare(t, "GET", "/dummy2.vts", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.vts", 404)
}

func TestVttsRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.vtts", 200)
	requestAndCompare(t, "GET", "/sub/dummy.vtts", 200)
	requestAndCompare(t, "GET", "/dummy2.vtts", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.vtts", 404)
}

func TestMfRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy=mf", 200)
	requestAndCompare(t, "GET", "/sub/dummy=mf", 200)
	requestAndCompare(t, "GET", "/dummy2=mf", 404)
	requestAndCompare(t, "GET", "/sub/dummy2=mf", 404)
}

func TestMRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy=m", 200)
	requestAndCompare(t, "GET", "/sub/dummy=m", 200)
	requestAndCompare(t, "GET", "/dummy2=m", 404)
	requestAndCompare(t, "GET", "/sub/dummy2=m", 404)
}

func TestQd3dRequest(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.qd3d", 200)
	requestAndCompare(t, "GET", "/sub/dummy.qd3d", 200)
	requestAndCompare(t, "GET", "/dummy2.qd3d", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.qd3d", 404)
}

func TestQd3Request(t *testing.T) {
	requestAndCompare(t, "GET", "/dummy.qd3", 200)
	requestAndCompare(t, "GET", "/sub/dummy.qd3", 200)
	requestAndCompare(t, "GET", "/dummy2.qd3", 404)
	requestAndCompare(t, "GET", "/sub/dummy2.qd3", 404)
}
