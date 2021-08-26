package ResourceServer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// findMimeTypeAndCheck get the MIME-Type and file path and test again expected result
func findMimeTypeAndCheck(t *testing.T, rURL, expectedPath, expectedMimeType string) {
	path, mimeType := findMimeType("./test_data" + rURL)
	assert.Equal(t, "./test_data"+expectedPath, path)
	assert.Equal(t, expectedMimeType, mimeType)
}

func TestNoneMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.none", "/dummy.none", "application/")
	findMimeTypeAndCheck(t, "/sub/dummy.none", "/sub/dummy.none", "application/")
}

func TestHtmlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/", "/index.html", "text/html")
	findMimeTypeAndCheck(t, "/second.html", "/second.html", "text/html")
	findMimeTypeAndCheck(t, "/sub", "/sub/index.html", "text/html")
	findMimeTypeAndCheck(t, "/sub/", "/sub/index.html", "text/html")
}

func TestDwgMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dwg", "/dummy.dwg", "application/acad")
	findMimeTypeAndCheck(t, "/sub/dummy.dwg", "/sub/dummy.dwg", "application/acad")
}

func TestAsdMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.asd", "/dummy.asd", "application/astound")
	findMimeTypeAndCheck(t, "/sub/dummy.asd", "/sub/dummy.asd", "application/astound")
}

func TestAsnMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.asn", "/dummy.asn", "application/astound")
	findMimeTypeAndCheck(t, "/sub/dummy.asn", "/sub/dummy.asn", "application/astound")
}

func TestTspMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tsp", "/dummy.tsp", "application/dsptype")
	findMimeTypeAndCheck(t, "/sub/dummy.tsp", "/sub/dummy.tsp", "application/dsptype")
}

func TestDxfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dxf", "/dummy.dxf", "application/dxf")
	findMimeTypeAndCheck(t, "/sub/dummy.dxf", "/sub/dummy.dxf", "application/dxf")
}

func TestRegMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.reg", "/dummy.reg", "application/force-download")
	findMimeTypeAndCheck(t, "/sub/dummy.reg", "/sub/dummy.reg", "application/force-download")
}

func TestSplMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.spl", "/dummy.spl", "application/futuresplash")
	findMimeTypeAndCheck(t, "/sub/dummy.spl", "/sub/dummy.spl", "application/futuresplash")
}

func TestGzMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.gz", "/dummy.gz", "application/gzip")
	findMimeTypeAndCheck(t, "/sub/dummy.gz", "/sub/dummy.gz", "application/gzip")
}

func TestJsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.js", "/dummy.js", "application/javascript")
	findMimeTypeAndCheck(t, "/sub/dummy.js", "/sub/dummy.js", "application/javascript")
}

func TestJsonMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.json", "/dummy.json", "application/json")
	findMimeTypeAndCheck(t, "/sub/dummy.json", "/sub/dummy.json", "application/json")
}

func TestPtlkMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ptlk", "/dummy.ptlk", "application/listenup")
	findMimeTypeAndCheck(t, "/sub/dummy.ptlk", "/sub/dummy.ptlk", "application/listenup")
}

func TestHqxMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.hqx", "/dummy.hqx", "application/mac-binhex40")
	findMimeTypeAndCheck(t, "/sub/dummy.hqx", "/sub/dummy.hqx", "application/mac-binhex40")
}

func TestMbdMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mbd", "/dummy.mbd", "application/mbedlet")
	findMimeTypeAndCheck(t, "/sub/dummy.mbd", "/sub/dummy.mbd", "application/mbedlet")
}

func TestMifMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mif", "/dummy.mif", "application/mif")
	findMimeTypeAndCheck(t, "/sub/dummy.mif", "/sub/dummy.mif", "application/mif")
}

func TestXlsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.xls", "/dummy.xls", "application/msexcel")
	findMimeTypeAndCheck(t, "/sub/dummy.xls", "/sub/dummy.xls", "application/msexcel")
}

func TestXlaMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.xla", "/dummy.xla", "application/msexcel")
	findMimeTypeAndCheck(t, "/sub/dummy.xla", "/sub/dummy.xla", "application/msexcel")
}

func TestChmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.chm", "/dummy.chm", "application/mshelp")
	findMimeTypeAndCheck(t, "/sub/dummy.chm", "/sub/dummy.chm", "application/mshelp")
}

func TestHlpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.hlp", "/dummy.hlp", "application/mshelp")
	findMimeTypeAndCheck(t, "/sub/dummy.hlp", "/sub/dummy.hlp", "application/mshelp")
}

func TestPptMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ppt", "/dummy.ppt", "application/mspowerpoint")
	findMimeTypeAndCheck(t, "/sub/dummy.ppt", "/sub/dummy.ppt", "application/mspowerpoint")
}

func TestPpzMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ppz", "/dummy.ppz", "application/mspowerpoint")
	findMimeTypeAndCheck(t, "/sub/dummy.ppz", "/sub/dummy.ppz", "application/mspowerpoint")
}

func TestPpsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.pps", "/dummy.pps", "application/mspowerpoint")
	findMimeTypeAndCheck(t, "/sub/dummy.pps", "/sub/dummy.pps", "application/mspowerpoint")
}

func TestPotMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.pot", "/dummy.pot", "application/mspowerpoint")
	findMimeTypeAndCheck(t, "/sub/dummy.pot", "/sub/dummy.pot", "application/mspowerpoint")
}

func TestDocMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.doc", "/dummy.doc", "application/msword")
	findMimeTypeAndCheck(t, "/sub/dummy.doc", "/sub/dummy.doc", "application/msword")
}

func TestDotMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dot", "/dummy.dot", "application/msword")
	findMimeTypeAndCheck(t, "/sub/dummy.dot", "/sub/dummy.dot", "application/msword")
}

func TestBinMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.bin", "/dummy.bin", "application/octet-stream")
	findMimeTypeAndCheck(t, "/sub/dummy.bin", "/sub/dummy.bin", "application/octet-stream")
}

func TestFileMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.file", "/dummy.file", "application/octet-stream")
	findMimeTypeAndCheck(t, "/sub/dummy.file", "/sub/dummy.file", "application/octet-stream")
}

func TestComMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.com", "/dummy.com", "application/octet-stream")
	findMimeTypeAndCheck(t, "/sub/dummy.com", "/sub/dummy.com", "application/octet-stream")
}

func TestClassMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.class", "/dummy.class", "application/octet-stream")
	findMimeTypeAndCheck(t, "/sub/dummy.class", "/sub/dummy.class", "application/octet-stream")
}

func TestIniMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ini", "/dummy.ini", "application/octet-stream")
	findMimeTypeAndCheck(t, "/sub/dummy.ini", "/sub/dummy.ini", "application/octet-stream")
}

func TestOdaMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.oda", "/dummy.oda", "application/oda")
	findMimeTypeAndCheck(t, "/sub/dummy.oda", "/sub/dummy.oda", "application/oda")
}

func TestPdfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.pdf", "/dummy.pdf", "application/pdf")
	findMimeTypeAndCheck(t, "/sub/dummy.pdf", "/sub/dummy.pdf", "application/pdf")
}

func TestAiMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ai", "/dummy.ai", "application/postscript")
	findMimeTypeAndCheck(t, "/sub/dummy.ai", "/sub/dummy.ai", "application/postscript")
}

func TestEpsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.eps", "/dummy.eps", "application/postscript")
	findMimeTypeAndCheck(t, "/sub/dummy.eps", "/sub/dummy.eps", "application/postscript")
}

func TestPsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ps", "/dummy.ps", "application/postscript")
	findMimeTypeAndCheck(t, "/sub/dummy.ps", "/sub/dummy.ps", "application/postscript")
}

func TestRtcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.rtc", "/dummy.rtc", "application/rtc")
	findMimeTypeAndCheck(t, "/sub/dummy.rtc", "/sub/dummy.rtc", "application/rtc")
}

func TestRtfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.rtf", "/dummy.rtf", "application/rtf")
	findMimeTypeAndCheck(t, "/sub/dummy.rtf", "/sub/dummy.rtf", "application/rtf")
}

func TestSmpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.smp", "/dummy.smp", "application/studiom")
	findMimeTypeAndCheck(t, "/sub/dummy.smp", "/sub/dummy.smp", "application/studiom")
}

func TestTbkMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tbk", "/dummy.tbk", "application/toolbook")
	findMimeTypeAndCheck(t, "/sub/dummy.tbk", "/sub/dummy.tbk", "application/toolbook")
}

func TestOdcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odc", "/dummy.odc", "application/vnd.oasis.opendocument.chart")
	findMimeTypeAndCheck(t, "/sub/dummy.odc", "/sub/dummy.odc", "application/vnd.oasis.opendocument.chart")
}

func TestOdfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odf", "/dummy.odf", "application/vnd.oasis.opendocument.formula")
	findMimeTypeAndCheck(t, "/sub/dummy.odf", "/sub/dummy.odf", "application/vnd.oasis.opendocument.formula")
}

func TestOdgMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odg", "/dummy.odg", "application/vnd.oasis.opendocument.graphics")
	findMimeTypeAndCheck(t, "/sub/dummy.odg", "/sub/dummy.odg", "application/vnd.oasis.opendocument.graphics")
}

func TestOdiMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odi", "/dummy.odi", "application/vnd.oasis.opendocument.image")
	findMimeTypeAndCheck(t, "/sub/dummy.odi", "/sub/dummy.odi", "application/vnd.oasis.opendocument.image")
}

func TestOdpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odp", "/dummy.odp", "application/vnd.oasis.opendocument.presentation")
	findMimeTypeAndCheck(t, "/sub/dummy.odp", "/sub/dummy.odp", "application/vnd.oasis.opendocument.presentation")
}

func TestOdsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ods", "/dummy.ods", "application/vnd.oasis.opendocument.spreadsheet")
	findMimeTypeAndCheck(t, "/sub/dummy.ods", "/sub/dummy.ods", "application/vnd.oasis.opendocument.spreadsheet")
}

func TestOdtMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odt", "/dummy.odt", "application/vnd.oasis.opendocument.text")
	findMimeTypeAndCheck(t, "/sub/dummy.odt", "/sub/dummy.odt", "application/vnd.oasis.opendocument.text")
}

func TestOdmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.odm", "/dummy.odm", "application/vnd.oasis.opendocument.text-master")
	findMimeTypeAndCheck(t, "/sub/dummy.odm", "/sub/dummy.odm", "application/vnd.oasis.opendocument.text-master")
}

func TestWmlcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wmlc", "/dummy.wmlc", "application/vnd.wap.wmlc")
	findMimeTypeAndCheck(t, "/sub/dummy.wmlc", "/sub/dummy.wmlc", "application/vnd.wap.wmlc")
}

func TestWmlscMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wmlsc", "/dummy.wmlsc", "application/vnd.wap.wmlscriptc")
	findMimeTypeAndCheck(t, "/sub/dummy.wmlsc", "/sub/dummy.wmlsc", "application/vnd.wap.wmlscriptc")
}

func TestVmdMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.vmd", "/dummy.vmd", "application/vocaltec-media-desc")
	findMimeTypeAndCheck(t, "/sub/dummy.vmd", "/sub/dummy.vmd", "application/vocaltec-media-desc")
}

func TestVmfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.vmf", "/dummy.vmf", "application/vocaltec-media-file")
	findMimeTypeAndCheck(t, "/sub/dummy.vmf", "/sub/dummy.vmf", "application/vocaltec-media-file")
}

func TestBcpioMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.bcpio", "/dummy.bcpio", "application/x-bcpio")
	findMimeTypeAndCheck(t, "/sub/dummy.bcpio", "/sub/dummy.bcpio", "application/x-bcpio")
}

func TestZMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.z", "/dummy.z", "application/x-compress")
	findMimeTypeAndCheck(t, "/sub/dummy.z", "/sub/dummy.z", "application/x-compress")
}

func TestCpioMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.cpio", "/dummy.cpio", "application/x-cpio")
	findMimeTypeAndCheck(t, "/sub/dummy.cpio", "/sub/dummy.cpio", "application/x-cpio")
}

func TestCshMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.csh", "/dummy.csh", "application/x-csh")
	findMimeTypeAndCheck(t, "/sub/dummy.csh", "/sub/dummy.csh", "application/x-csh")
}

func TestDcrMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dcr", "/dummy.dcr", "application/x-director")
	findMimeTypeAndCheck(t, "/sub/dummy.dcr", "/sub/dummy.dcr", "application/x-director")
}

func TestDirMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dir", "/dummy.dir", "application/x-director")
	findMimeTypeAndCheck(t, "/sub/dummy.dir", "/sub/dummy.dir", "application/x-director")
}

func TestDxrMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dxr", "/dummy.dxr", "application/x-director")
	findMimeTypeAndCheck(t, "/sub/dummy.dxr", "/sub/dummy.dxr", "application/x-director")
}

func TestDviMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dvi", "/dummy.dvi", "application/x-dvi")
	findMimeTypeAndCheck(t, "/sub/dummy.dvi", "/sub/dummy.dvi", "application/x-dvi")
}

func TestEvyMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.evy", "/dummy.evy", "application/x-envoy")
	findMimeTypeAndCheck(t, "/sub/dummy.evy", "/sub/dummy.evy", "application/x-envoy")
}

func TestGtarMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.gtar", "/dummy.gtar", "application/x-gtar")
	findMimeTypeAndCheck(t, "/sub/dummy.gtar", "/sub/dummy.gtar", "application/x-gtar")
}

func TestHdfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.hdf", "/dummy.hdf", "application/x-hdf")
	findMimeTypeAndCheck(t, "/sub/dummy.hdf", "/sub/dummy.hdf", "application/x-hdf")
}

func TestPhpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.php", "/dummy.php", "application/x-httpd-php")
	findMimeTypeAndCheck(t, "/sub/dummy.php", "/sub/dummy.php", "application/x-httpd-php")
}

func TestPhtmlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.phtml", "/dummy.phtml", "application/x-httpd-php")
	findMimeTypeAndCheck(t, "/sub/dummy.phtml", "/sub/dummy.phtml", "application/x-httpd-php")
}

func TestLatexMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.latex", "/dummy.latex", "application/x-latex")
	findMimeTypeAndCheck(t, "/sub/dummy.latex", "/sub/dummy.latex", "application/x-latex")
}

func TestXmlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.xml", "/dummy.xml", "application/xml")
	findMimeTypeAndCheck(t, "/sub/dummy.xml", "/sub/dummy.xml", "application/xml")
}

func TestNcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.nc", "/dummy.nc", "application/x-netcdf")
	findMimeTypeAndCheck(t, "/sub/dummy.nc", "/sub/dummy.nc", "application/x-netcdf")
}

func TestCdfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.cdf", "/dummy.cdf", "application/x-netcdf")
	findMimeTypeAndCheck(t, "/sub/dummy.cdf", "/sub/dummy.cdf", "application/x-netcdf")
}

func TestNscMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.nsc", "/dummy.nsc", "application/x-nschat")
	findMimeTypeAndCheck(t, "/sub/dummy.nsc", "/sub/dummy.nsc", "application/x-nschat")
}

func TestShMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sh", "/dummy.sh", "application/x-sh")
	findMimeTypeAndCheck(t, "/sub/dummy.sh", "/sub/dummy.sh", "application/x-sh")
}

func TestSharMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.shar", "/dummy.shar", "application/x-shar")
	findMimeTypeAndCheck(t, "/sub/dummy.shar", "/sub/dummy.shar", "application/x-shar")
}

func TestSwfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.swf", "/dummy.swf", "application/x-shockwave-flash")
	findMimeTypeAndCheck(t, "/sub/dummy.swf", "/sub/dummy.swf", "application/x-shockwave-flash")
}

func TestCabMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.cab", "/dummy.cab", "application/x-shockwave-flash")
	findMimeTypeAndCheck(t, "/sub/dummy.cab", "/sub/dummy.cab", "application/x-shockwave-flash")
}

func TestSprMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.spr", "/dummy.spr", "application/x-sprite")
	findMimeTypeAndCheck(t, "/sub/dummy.spr", "/sub/dummy.spr", "application/x-sprite")
}

func TestSpriteMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sprite", "/dummy.sprite", "application/x-sprite")
	findMimeTypeAndCheck(t, "/sub/dummy.sprite", "/sub/dummy.sprite", "application/x-sprite")
}

func TestSitMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sit", "/dummy.sit", "application/x-stuffit")
	findMimeTypeAndCheck(t, "/sub/dummy.sit", "/sub/dummy.sit", "application/x-stuffit")
}

func TestScaMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sca", "/dummy.sca", "application/x-supercard")
	findMimeTypeAndCheck(t, "/sub/dummy.sca", "/sub/dummy.sca", "application/x-supercard")
}

func TestSv4cpioMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sv4cpio", "/dummy.sv4cpio", "application/x-sv4cpio")
	findMimeTypeAndCheck(t, "/sub/dummy.sv4cpio", "/sub/dummy.sv4cpio", "application/x-sv4cpio")
}

func TestSv4crcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sv4crc", "/dummy.sv4crc", "application/x-sv4crc")
	findMimeTypeAndCheck(t, "/sub/dummy.sv4crc", "/sub/dummy.sv4crc", "application/x-sv4crc")
}

func TestTarMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tar", "/dummy.tar", "application/x-tar")
	findMimeTypeAndCheck(t, "/sub/dummy.tar", "/sub/dummy.tar", "application/x-tar")
}

func TestTclMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tcl", "/dummy.tcl", "application/x-tcl")
	findMimeTypeAndCheck(t, "/sub/dummy.tcl", "/sub/dummy.tcl", "application/x-tcl")
}

func TestTexMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tex", "/dummy.tex", "application/x-tex")
	findMimeTypeAndCheck(t, "/sub/dummy.tex", "/sub/dummy.tex", "application/x-tex")
}

func TestTexinfoMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.texinfo", "/dummy.texinfo", "application/x-texinfo")
	findMimeTypeAndCheck(t, "/sub/dummy.texinfo", "/sub/dummy.texinfo", "application/x-texinfo")
}

func TestTexiMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.texi", "/dummy.texi", "application/x-texinfo")
	findMimeTypeAndCheck(t, "/sub/dummy.texi", "/sub/dummy.texi", "application/x-texinfo")
}

func TestTMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.t", "/dummy.t", "application/x-troff")
	findMimeTypeAndCheck(t, "/sub/dummy.t", "/sub/dummy.t", "application/x-troff")
}

func TestTrMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tr", "/dummy.tr", "application/x-troff")
	findMimeTypeAndCheck(t, "/sub/dummy.tr", "/sub/dummy.tr", "application/x-troff")
}

func TestRoffMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.roff", "/dummy.roff", "application/x-troff")
	findMimeTypeAndCheck(t, "/sub/dummy.roff", "/sub/dummy.roff", "application/x-troff")
}

func TestUstarMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ustar", "/dummy.ustar", "application/x-ustar")
	findMimeTypeAndCheck(t, "/sub/dummy.ustar", "/sub/dummy.ustar", "application/x-ustar")
}

func TestSrcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.src", "/dummy.src", "application/x-wais-source")
	findMimeTypeAndCheck(t, "/sub/dummy.src", "/sub/dummy.src", "application/x-wais-source")
}

func TestZipMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.zip", "/dummy.zip", "application/zip")
	findMimeTypeAndCheck(t, "/sub/dummy.zip", "/sub/dummy.zip", "application/zip")
}

func TestAuMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.au", "/dummy.au", "audio/basic")
	findMimeTypeAndCheck(t, "/sub/dummy.au", "/sub/dummy.au", "audio/basic")
}

func TestSndMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.snd", "/dummy.snd", "audio/basic")
	findMimeTypeAndCheck(t, "/sub/dummy.snd", "/sub/dummy.snd", "audio/basic")
}

func TestEsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.es", "/dummy.es", "audio/echospeech")
	findMimeTypeAndCheck(t, "/sub/dummy.es", "/sub/dummy.es", "audio/echospeech")
}

func TestMp3MimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mp3", "/dummy.mp3", "audio/mpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.mp3", "/sub/dummy.mp3", "audio/mpeg")
}

func TestTsiMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tsi", "/dummy.tsi", "audio/tsplayer")
	findMimeTypeAndCheck(t, "/sub/dummy.tsi", "/sub/dummy.tsi", "audio/tsplayer")
}

func TestVoxMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.vox", "/dummy.vox", "audio/voxware")
	findMimeTypeAndCheck(t, "/sub/dummy.vox", "/sub/dummy.vox", "audio/voxware")
}

func TestWavMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wav", "/dummy.wav", "audio/wav")
	findMimeTypeAndCheck(t, "/sub/dummy.wav", "/sub/dummy.wav", "audio/wav")
}

func TestAifMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.aif", "/dummy.aif", "audio/x-aiff")
	findMimeTypeAndCheck(t, "/sub/dummy.aif", "/sub/dummy.aif", "audio/x-aiff")
}

func TestAiffMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.aiff", "/dummy.aiff", "audio/x-aiff")
	findMimeTypeAndCheck(t, "/sub/dummy.aiff", "/sub/dummy.aiff", "audio/x-aiff")
}

func TestAifcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.aifc", "/dummy.aifc", "audio/x-aiff")
	findMimeTypeAndCheck(t, "/sub/dummy.aifc", "/sub/dummy.aifc", "audio/x-aiff")
}

func TestDusMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dus", "/dummy.dus", "audio/x-dspeech")
	findMimeTypeAndCheck(t, "/sub/dummy.dus", "/sub/dummy.dus", "audio/x-dspeech")
}

func TestChtMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.cht", "/dummy.cht", "audio/x-dspeech")
	findMimeTypeAndCheck(t, "/sub/dummy.cht", "/sub/dummy.cht", "audio/x-dspeech")
}

func TestMidMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mid", "/dummy.mid", "audio/x-midi")
	findMimeTypeAndCheck(t, "/sub/dummy.mid", "/sub/dummy.mid", "audio/x-midi")
}

func TestMidiMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.midi", "/dummy.midi", "audio/x-midi")
	findMimeTypeAndCheck(t, "/sub/dummy.midi", "/sub/dummy.midi", "audio/x-midi")
}

func TestMp2MimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mp2", "/dummy.mp2", "audio/x-mpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.mp2", "/sub/dummy.mp2", "audio/x-mpeg")
}

func TestRamMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ram", "/dummy.ram", "audio/x-pn-realaudio")
	findMimeTypeAndCheck(t, "/sub/dummy.ram", "/sub/dummy.ram", "audio/x-pn-realaudio")
}

func TestRaMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ra", "/dummy.ra", "audio/x-pn-realaudio")
	findMimeTypeAndCheck(t, "/sub/dummy.ra", "/sub/dummy.ra", "audio/x-pn-realaudio")
}

func TestRpmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.rpm", "/dummy.rpm", "audio/x-pn-realaudio-plugin")
	findMimeTypeAndCheck(t, "/sub/dummy.rpm", "/sub/dummy.rpm", "audio/x-pn-realaudio-plugin")
}

func TestStreamMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.stream", "/dummy.stream", "audio/x-qt-stream")
	findMimeTypeAndCheck(t, "/sub/dummy.stream", "/sub/dummy.stream", "audio/x-qt-stream")
}

func TestDwfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.dwf", "/dummy.dwf", "drawing/x-dwf")
	findMimeTypeAndCheck(t, "/sub/dummy.dwf", "/sub/dummy.dwf", "drawing/x-dwf")
}

func TestBmpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.bmp", "/dummy.bmp", "image/bmp")
	findMimeTypeAndCheck(t, "/sub/dummy.bmp", "/sub/dummy.bmp", "image/bmp")
}

func TestCodMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.cod", "/dummy.cod", "image/cis-cod")
	findMimeTypeAndCheck(t, "/sub/dummy.cod", "/sub/dummy.cod", "image/cis-cod")
}

func TestRasMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ras", "/dummy.ras", "image/cmu-raster")
	findMimeTypeAndCheck(t, "/sub/dummy.ras", "/sub/dummy.ras", "image/cmu-raster")
}

func TestFifMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.fif", "/dummy.fif", "image/fif")
	findMimeTypeAndCheck(t, "/sub/dummy.fif", "/sub/dummy.fif", "image/fif")
}

func TestGifMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.gif", "/dummy.gif", "image/gif")
	findMimeTypeAndCheck(t, "/sub/dummy.gif", "/sub/dummy.gif", "image/gif")
}

func TestIefMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ief", "/dummy.ief", "image/ief")
	findMimeTypeAndCheck(t, "/sub/dummy.ief", "/sub/dummy.ief", "image/ief")
}

func TestJpegMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.jpeg", "/dummy.jpeg", "image/jpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.jpeg", "/sub/dummy.jpeg", "image/jpeg")
}

func TestJpgMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.jpg", "/dummy.jpg", "image/jpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.jpg", "/sub/dummy.jpg", "image/jpeg")
}

func TestJpeMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.jpe", "/dummy.jpe", "image/jpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.jpe", "/sub/dummy.jpe", "image/jpeg")
}

func TestPngMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.png", "/dummy.png", "image/png")
	findMimeTypeAndCheck(t, "/sub/dummy.png", "/sub/dummy.png", "image/png")
}

func TestSvgMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.svg", "/dummy.svg", "image/svg+xml")
	findMimeTypeAndCheck(t, "/sub/dummy.svg", "/sub/dummy.svg", "image/svg+xml")
}

func TestTiffMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tiff", "/dummy.tiff", "image/tiff")
	findMimeTypeAndCheck(t, "/sub/dummy.tiff", "/sub/dummy.tiff", "image/tiff")
}

func TestTifMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tif", "/dummy.tif", "image/tiff")
	findMimeTypeAndCheck(t, "/sub/dummy.tif", "/sub/dummy.tif", "image/tiff")
}

func TestMcfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mcf", "/dummy.mcf", "image/vasa")
	findMimeTypeAndCheck(t, "/sub/dummy.mcf", "/sub/dummy.mcf", "image/vasa")
}

func TestWbmpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wbmp", "/dummy.wbmp", "image/vnd.wap.wbmp")
	findMimeTypeAndCheck(t, "/sub/dummy.wbmp", "/sub/dummy.wbmp", "image/vnd.wap.wbmp")
}

func TestFh4MimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.fh4", "/dummy.fh4", "image/x-freehand")
	findMimeTypeAndCheck(t, "/sub/dummy.fh4", "/sub/dummy.fh4", "image/x-freehand")
}

func TestFh5MimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.fh5", "/dummy.fh5", "image/x-freehand")
	findMimeTypeAndCheck(t, "/sub/dummy.fh5", "/sub/dummy.fh5", "image/x-freehand")
}

func TestFhcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.fhc", "/dummy.fhc", "image/x-freehand")
	findMimeTypeAndCheck(t, "/sub/dummy.fhc", "/sub/dummy.fhc", "image/x-freehand")
}

func TestIcoMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ico", "/dummy.ico", "image/x-icon")
	findMimeTypeAndCheck(t, "/sub/dummy.ico", "/sub/dummy.ico", "image/x-icon")
}

func TestPnmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.pnm", "/dummy.pnm", "image/x-portable-anymap")
	findMimeTypeAndCheck(t, "/sub/dummy.pnm", "/sub/dummy.pnm", "image/x-portable-anymap")
}

func TestPbmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.pbm", "/dummy.pbm", "image/x-portable-bitmap")
	findMimeTypeAndCheck(t, "/sub/dummy.pbm", "/sub/dummy.pbm", "image/x-portable-bitmap")
}

func TestPgmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.pgm", "/dummy.pgm", "image/x-portable-graymap")
	findMimeTypeAndCheck(t, "/sub/dummy.pgm", "/sub/dummy.pgm", "image/x-portable-graymap")
}

func TestPpmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ppm", "/dummy.ppm", "image/x-portable-pixmap")
	findMimeTypeAndCheck(t, "/sub/dummy.ppm", "/sub/dummy.ppm", "image/x-portable-pixmap")
}

func TestRgbMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.rgb", "/dummy.rgb", "image/x-rgb")
	findMimeTypeAndCheck(t, "/sub/dummy.rgb", "/sub/dummy.rgb", "image/x-rgb")
}

func TestXwdMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.xwd", "/dummy.xwd", "image/x-windowdump")
	findMimeTypeAndCheck(t, "/sub/dummy.xwd", "/sub/dummy.xwd", "image/x-windowdump")
}

func TestXbmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.xbm", "/dummy.xbm", "image/x-xbitmap")
	findMimeTypeAndCheck(t, "/sub/dummy.xbm", "/sub/dummy.xbm", "image/x-xbitmap")
}

func TestXpmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.xpm", "/dummy.xpm", "image/x-xpixmap")
	findMimeTypeAndCheck(t, "/sub/dummy.xpm", "/sub/dummy.xpm", "image/x-xpixmap")
}

func TestWrlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wrl", "/dummy.wrl", "model/vrml")
	findMimeTypeAndCheck(t, "/sub/dummy.wrl", "/sub/dummy.wrl", "model/vrml")
}

func TestIcsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ics", "/dummy.ics", "text/calendar")
	findMimeTypeAndCheck(t, "/sub/dummy.ics", "/sub/dummy.ics", "text/calendar")
}

func TestCsvMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.csv", "/dummy.csv", "text/comma-separated-values")
	findMimeTypeAndCheck(t, "/sub/dummy.csv", "/sub/dummy.csv", "text/comma-separated-values")
}

func TestCssMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.css", "/dummy.css", "text/css")
	findMimeTypeAndCheck(t, "/sub/dummy.css", "/sub/dummy.css", "text/css")
}

func TestHtmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.htm", "/dummy.htm", "text/html")
	findMimeTypeAndCheck(t, "/sub/dummy.htm", "/sub/dummy.htm", "text/html")
}

func TestShtmlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.shtml", "/dummy.shtml", "text/html")
	findMimeTypeAndCheck(t, "/sub/dummy.shtml", "/sub/dummy.shtml", "text/html")
}

func TestTxtMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.txt", "/dummy.txt", "text/plain")
	findMimeTypeAndCheck(t, "/sub/dummy.txt", "/sub/dummy.txt", "text/plain")
}

func TestRtxMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.rtx", "/dummy.rtx", "text/richtext")
	findMimeTypeAndCheck(t, "/sub/dummy.rtx", "/sub/dummy.rtx", "text/richtext")
}

func TestTsvMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.tsv", "/dummy.tsv", "text/tab-separated-values")
	findMimeTypeAndCheck(t, "/sub/dummy.tsv", "/sub/dummy.tsv", "text/tab-separated-values")
}

func TestWmlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wml", "/dummy.wml", "text/vnd.wap.wml")
	findMimeTypeAndCheck(t, "/sub/dummy.wml", "/sub/dummy.wml", "text/vnd.wap.wml")
}

func TestWmlsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.wmls", "/dummy.wmls", "text/vnd.wap.wmlscript")
	findMimeTypeAndCheck(t, "/sub/dummy.wmls", "/sub/dummy.wmls", "text/vnd.wap.wmlscript")
}

func TestEtxMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.etx", "/dummy.etx", "text/x-setext")
	findMimeTypeAndCheck(t, "/sub/dummy.etx", "/sub/dummy.etx", "text/x-setext")
}

func TestSgmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sgm", "/dummy.sgm", "text/x-sgml")
	findMimeTypeAndCheck(t, "/sub/dummy.sgm", "/sub/dummy.sgm", "text/x-sgml")
}

func TestSgmlMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.sgml", "/dummy.sgml", "text/x-sgml")
	findMimeTypeAndCheck(t, "/sub/dummy.sgml", "/sub/dummy.sgml", "text/x-sgml")
}

func TestTalkMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.talk", "/dummy.talk", "text/x-speech")
	findMimeTypeAndCheck(t, "/sub/dummy.talk", "/sub/dummy.talk", "text/x-speech")
}

func TestSpcMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.spc", "/dummy.spc", "text/x-speech")
	findMimeTypeAndCheck(t, "/sub/dummy.spc", "/sub/dummy.spc", "text/x-speech")
}

func Test3gpMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.3gp", "/dummy.3gp", "video/3gpp")
	findMimeTypeAndCheck(t, "/sub/dummy.3gp", "/sub/dummy.3gp", "video/3gpp")
}

func TestMp4MimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mp4", "/dummy.mp4", "video/mp4")
	findMimeTypeAndCheck(t, "/sub/dummy.mp4", "/sub/dummy.mp4", "video/mp4")
}

func TestMpegMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mpeg", "/dummy.mpeg", "video/mpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.mpeg", "/sub/dummy.mpeg", "video/mpeg")
}

func TestMpgMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mpg", "/dummy.mpg", "video/mpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.mpg", "/sub/dummy.mpg", "video/mpeg")
}

func TestMpeMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mpe", "/dummy.mpe", "video/mpeg")
	findMimeTypeAndCheck(t, "/sub/dummy.mpe", "/sub/dummy.mpe", "video/mpeg")
}

func TestOggMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ogg", "/dummy.ogg", "video/ogg")
	findMimeTypeAndCheck(t, "/sub/dummy.ogg", "/sub/dummy.ogg", "video/ogg")
}

func TestOgvMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.ogv", "/dummy.ogv", "video/ogg")
	findMimeTypeAndCheck(t, "/sub/dummy.ogv", "/sub/dummy.ogv", "video/ogg")
}

func TestQtMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.qt", "/dummy.qt", "video/quicktime")
	findMimeTypeAndCheck(t, "/sub/dummy.qt", "/sub/dummy.qt", "video/quicktime")
}

func TestMovMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.mov", "/dummy.mov", "video/quicktime")
	findMimeTypeAndCheck(t, "/sub/dummy.mov", "/sub/dummy.mov", "video/quicktime")
}

func TestVivMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.viv", "/dummy.viv", "video/vnd.vivo")
	findMimeTypeAndCheck(t, "/sub/dummy.viv", "/sub/dummy.viv", "video/vnd.vivo")
}

func TestVivoMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.vivo", "/dummy.vivo", "video/vnd.vivo")
	findMimeTypeAndCheck(t, "/sub/dummy.vivo", "/sub/dummy.vivo", "video/vnd.vivo")
}

func TestWebmMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.webm", "/dummy.webm", "video/webm")
	findMimeTypeAndCheck(t, "/sub/dummy.webm", "/sub/dummy.webm", "video/webm")
}

func TestAviMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.avi", "/dummy.avi", "video/x-msvideo")
	findMimeTypeAndCheck(t, "/sub/dummy.avi", "/sub/dummy.avi", "video/x-msvideo")
}

func TestMovieMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.movie", "/dummy.movie", "video/x-sgi-movie")
	findMimeTypeAndCheck(t, "/sub/dummy.movie", "/sub/dummy.movie", "video/x-sgi-movie")
}

func TestVtsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.vts", "/dummy.vts", "workbook/formulaone")
	findMimeTypeAndCheck(t, "/sub/dummy.vts", "/sub/dummy.vts", "workbook/formulaone")
}

func TestVttsMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.vtts", "/dummy.vtts", "workbook/formulaone")
	findMimeTypeAndCheck(t, "/sub/dummy.vtts", "/sub/dummy.vtts", "workbook/formulaone")
}

func TestMfMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy=mf", "/dummy=mf", "x-world/x-3dmf")
	findMimeTypeAndCheck(t, "/sub/dummy=mf", "/sub/dummy=mf", "x-world/x-3dmf")
}

func TestMMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy=m", "/dummy=m", "x-world/x-3dmf")
	findMimeTypeAndCheck(t, "/sub/dummy=m", "/sub/dummy=m", "x-world/x-3dmf")
}

func TestQd3dMimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.qd3d", "/dummy.qd3d", "x-world/x-3dmf")
	findMimeTypeAndCheck(t, "/sub/dummy.qd3d", "/sub/dummy.qd3d", "x-world/x-3dmf")
}

func TestQd3MimeType(t *testing.T) {
	findMimeTypeAndCheck(t, "/dummy.qd3", "/dummy.qd3", "x-world/x-3dmf")
	findMimeTypeAndCheck(t, "/sub/dummy.qd3", "/sub/dummy.qd3", "x-world/x-3dmf")
}
