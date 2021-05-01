package main

import (
	"errors"
	"fmt"
	"strings"
)

var ConvertibleFormatsMap = map[string][27]string{
	//Archive
	"7Z":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"ACE":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"ALZ":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"ARC":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"ARJ":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"BZ":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"BZ2":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"CAB":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"CPIO":    {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"DEB":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"DMG":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "TAR.LZO", "ZIP"},
	"GZ":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"IMG":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "TAR.LZO", "ZIP"},
	"ISO":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "TAR.LZO", "ZIP"},
	"JAR":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"LHA":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"LZ":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"LZMA":    {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"LZO":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"RAR":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"RPM":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"RZ":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.7Z":  {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.BZ":  {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.BZ2": {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.GZ":  {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.LZO": {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.XZ":  {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TAR.Z":   {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TBZ":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TBZ2":    {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TGZ":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TZ":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"TZO":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"XZ":      {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"Z":       {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	"ZIP":     {"7Z", "RAR", "TAR", "TAR.BZ2", "TAR.GZ", "ZIP"},
	//Audio
	"AAC":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"AC3":   {"AAC", "AC3", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"AIF":   {"AAC", "AIF", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"AIFC":  {"AAC", "AIFC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"AIFF":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"AMR":   {"AAC", "AIFF", "AMR", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"AU":    {"AAC", "AIFF", "AU", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"CAF":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"FLAC":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"M4A":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"M4B":   {"AAC", "AIFF", "FLAC", "M4A", "M4B", "MP3", "WAV", "WMA"},
	"MP3":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"OGA":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "OGA", "WAV", "WMA"},
	"SFARK": {"SF2", "SFARK"},
	"VOC":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "VOC", "WAV", "WMA"},
	"WAV":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	"WEBA":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WEBA", "WMA"},
	"WMA":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA"},
	//CAD
	"DWG": {"DWG", "DXF", "PDF", "BMP", "EPS", "GIF", "JPG", "PNG", "TIFF", "CGM", "SVG", "WMF"},
	"DXF": {"DWG", "DXF", "PDF", "BMP", "EPS", "GIF", "JPG", "PNG", "TIFF", "CGM", "SVG", "WMF"},
	//Document
	"ABW":   {"ABW", "DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG"},
	"DJVU":  {"DJVU", "PDF"},
	"DOC":   {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG", "XPS"},
	"DOCM":  {"DOC", "DOCM", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG", "XPS"},
	"DOCX":  {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "JPG", "PNG", "XPS"},
	"DOT":   {"DOC", "DOCX", "DOT", "HTML", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG", "XPS"},
	"DOTX":  {"DOC", "DOCX", "DOTX", "HTML", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG", "XPS"},
	"HTML":  {"DOC", "DOCX", "HTML", "MD", "ODT", "PDF", "RTF", "TEX", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "JPG", "PNG"},
	"HWP":   {"DOC", "DOCX", "HTML", "HWP", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG", "XPS"},
	"LWP":   {"DOC", "DOCX", "HTML", "LWP", "ODT", "PDF", "RTF", "TXT", "JPG", "PNG"},
	"MD":    {"DOCX", "HTML", "MD", "ODT", "PDF", "RST", "RTF", "TEX", "TXT"},
	"ODT":   {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "JPG", "PNG"},
	"PAGES": {"DOC", "DOCX", "PDF", "PAGES", "TXT", "EPUB", "JPG", "PNG"},
	"PDF":   {"DXF", "DOC", "DOCX", "HTML", "PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "EMF", "SVG", "WMF"},
	"RST":   {"DOCX", "HTML", "MD", "ODT", "PDF", "RST", "RTF", "TXT"},
	"RTF":   {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "JPG", "PNG"},
	"SDW":   {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "SDW", "TXT"},
	"TEX":   {"MD", "PDF", "TEX", "TXT"},
	"TXT":   {"DOC", "DOCX", "HTML", "MD", "ODT", "PDF", "RTF", "TEX", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "JPG", "PNG"},
	"WPD":   {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "WPD", "JPG", "PNG"},
	"WPS":   {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "WPS", "JPG", "PNG"},
	"ZABW":  {"DOC", "DOCX", "HTML", "ODT", "PDF", "RTF", "TXT", "ZABW", "JPG", "PNG"},
	//EBook
	"AZW":   {"PDF", "RTF", "TXT", "AZW", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"AZW3":  {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"AZW4":  {"PDF", "RTF", "TXT", "AZW3", "AZW4", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"CBC":   {"PDF", "RTF", "TXT", "AZW3", "CBC", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"CBR":   {"PDF", "RTF", "TXT", "AZW3", "CBR", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"CBZ":   {"PDF", "RTF", "TXT", "AZW3", "CBZ", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"CHM":   {"PDF", "RTF", "TXT", "AZW3", "CHM", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"EPUB":  {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"FB2":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "FB2", "LRF", "MOBI", "OEB", "PDB"},
	"HTM":   {"DOC", "DOCX", "ODT", "PDF", "RTF", "TXT", "AZW3", "EPUB", "HTM", "LRF", "MOBI", "OEB", "PDF"},
	"HTMLZ": {"PDF", "RTF", "TXT", "AZW3", "EPUB", "HTMLZ", "LRF", "MOBI", "OEB", "PDB"},
	"LIT":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LIT", "LRF", "MOBI", "OEB", "PDB"},
	"LRF":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"MOBI":  {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"PDB":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB"},
	"PML":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "PML"},
	"PRC":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "PRC"},
	"RB":    {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "RB"},
	"SNB":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "SNB"},
	"TCR":   {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "TCR"},
	"TXTZ":  {"PDF", "RTF", "TXT", "AZW3", "EPUB", "LRF", "MOBI", "OEB", "PDB", "TXTZ"},
	//Font
	"EOT":   {"EOT", "OTF", "TTF", "WOFF", "WOFF2", "SVG"},
	"OTF":   {"EOT", "OTF", "TTF", "WOFF", "WOFF2", "SVG"},
	"TTF":   {"EOT", "OTF", "TTF", "WOFF", "WOFF2", "SVG"},
	"WOFF":  {"EOT", "OTF", "TTF", "WOFF", "WOFF2", "SVG"},
	"WOFF2": {"EOT", "OTF", "TTF", "WOFF", "WOFF2", "SVG"},
	//Image
	"3FR":  {"PDF", "3FR", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"ARW":  {"PDF", "ARW", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"AVIF": {"PDF", "AVIF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"BMP":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"CR2":  {"PDF", "BMP", "CR2", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"CR3":  {"PDF", "BMP", "CR3", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"CRW":  {"PDF", "BMP", "CRW", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"DCR":  {"PDF", "BMP", "DCR", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"DNG":  {"PDF", "BMP", "DNG", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"EPS":  {"DXF", "PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "EMF", "SVG", "WMF"},
	"ERF":  {"PDF", "BMP", "EPS", "ERF", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"GIF":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"HEIC": {"PDF", "BMP", "EPS", "GIF", "HEIC", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"ICNS": {"PDF", "ICNS"},
	"ICO":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"JFIF": {"PDF", "BMP", "EPS", "GIF", "ICO", "JFIF", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"JPEG": {"PDF", "BMP", "EPS", "GIF", "ICO", "JPEG", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"JPG":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"MOS":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "MOS", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"MRW":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "MRW", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"NEF":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "NEF", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"ODD":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"ORF":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "ORF", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"PEF":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PEF", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"PNG":  {"PDF", "BMP", "EPS", "GIF", "ICNS", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"PPM":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PPM", "PS", "PSD", "TIFF", "WEBP"},
	"PS":   {"DXF", "HTML", "PDF", "TXT", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "ODP", "PPT", "PPTX", "EMF", "SVG", "WMF", "SWF"},
	"PSD":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"RAF":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "RAF", "TIFF", "WEBP"},
	"RAW":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "RAW", "TIFF", "WEBP"},
	"RAW2": {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "RAW2", "TIFF", "WEBP"},
	"TIF":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIF", "TIFF", "WEBP"},
	"TIFF": {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"WEBP": {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP"},
	"X3F":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "X3F"},
	"XCF":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "XCF"},
	"XPS":  {"PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "XPS", "SVG"},
	//Presentation
	"DPS":  {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "DPS", "ODP", "PPT", "PPTX", "SWF"},
	"KEY":  {"PDF", "JPG", "PNG", "PPT", "PPTX"},
	"ODP":  {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "ODP", "PPT", "PPTX", "SWF"},
	"POT":  {"PDF", "JPG", "PNG", "XPS", "ODP", "POT", "PPT", "PPTX", "EMF"},
	"POTX": {"PDF", "JPG", "PNG", "XPS", "ODP", "POTX", "PPT", "PPTX", "EMF"},
	"PPS":  {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "ODP", "PPS", "PPT", "PPTX", "SWF"},
	"PPSX": {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "ODP", "PPSX", "PPT", "PPTX", "SWF"},
	"PPT":  {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "XPS", "ODP", "PPT", "PPTX", "EMF", "SWF"},
	"PPTM": {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "XPS", "ODP", "PPT", "PPTM", "PPTX", "EMF", "SWF"},
	"PPTX": {"HTML", "PDF", "TXT", "EPS", "JPG", "PNG", "XPS", "ODP", "PPT", "PPTX", "EMF", "SWF"},
	"SDA":  {"HTML", "PDF", "EPS", "ODP", "PPT", "PPTX", "SDA", "SWF"},
	//Spreadsheet
	"CSV":     {"HTML", "PDF", "JPG", "PNG", "CSV", "ODS", "XLS", "XLSX"},
	"ET":      {"HTML", "PDF", "JPG", "PNG", "CSV", "ET", "ODS", "XLS", "XLSX"},
	"NUMBERS": {"PDF", "JPG", "PNG", "CSV", "XLS", "XLSX"},
	"ODS":     {"HTML", "PDF", "JPG", "PNG", "CSV", "ODS", "XLS", "XLSX"},
	"SDC":     {"HTML", "PDF", "CSV", "ODS", "SDC", "XLS", "XLSX"},
	"XLS":     {"HTML", "PDF", "JPG", "PNG", "XPS", "CSV", "ODS", "XLS", "XLSX"},
	"XLSM":    {"HTML", "PDF", "JPG", "PNG", "XPS", "CSV", "ODS", "XLS", "XLSM", "XLSX"},
	"XLSX":    {"HTML", "PDF", "JPG", "PNG", "XPS", "CSV", "ODS", "XLS", "XLSX"},
	//Vector
	"AI":   {"DXF", "PDF", "EPS", "PNG", "PS", "AI", "EMF", "SVG", "WMF"},
	"CDR":  {"DXF", "PDF", "EPS", "PNG", "PS", "CDR", "EMF", "SVG", "WMF"},
	"CGM":  {"DXF", "PDF", "EPS", "PNG", "PS", "CGM", "EMF", "SVG", "WMF"},
	"EMF":  {"DXF", "PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "EMF", "SVG", "WMF"},
	"SK":   {"DXF", "PDF", "EPS", "PNG", "PS", "EMF", "SK", "SVG", "WMF"},
	"SK1":  {"DXF", "PDF", "EPS", "PNG", "PS", "EMF", "SK1", "SVG", "WMF"},
	"SVG":  {"DXF", "PDF", "EOT", "OTF", "TTF", "WOFF", "WOFF2", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "EMF", "SVG", "WMF"},
	"SVGZ": {"DXF", "PDF", "BMP", "EPS", "GIF", "ICO", "JPG", "ODD", "PNG", "PS", "PSD", "TIFF", "WEBP", "EMF", "SVG", "WMF"},
	"VSD":  {"DXF", "PDF", "EPS", "PNG", "PS", "CGM", "EMF", "SVG", "WMF"},
	"WMF":  {"DXF", "PDF", "EPS", "PNG", "PS", "CGM", "EMF", "SVG", "WMF"},
	//Video
	"3G2":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "3G2", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"3GP":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "3GP", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"3GPP": {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "3GPP", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"AVI":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"CAVS": {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "CAVS", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"DV":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "DV", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"DVR":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "DVR", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"FLV":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"M2TS": {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "M2TS", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"M4V":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "M4V", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"MKV":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"MOD":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOD", "MOV", "MP4", "WEBM", "WMV"},
	"MOV":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"MP4":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"MPEG": {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "MPEG", "WEBM", "WMV"},
	"MPG":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "MPG", "WEBM", "WMV"},
	"MTS":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "MTS", "WEBM", "WMV"},
	"MXF":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "MXF", "WEBM", "WMV"},
	"OGG":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "OGG", "WEBM", "WMV"},
	"RM":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "RM", "WEBM", "WMV"},
	"RMVB": {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "RMVB", "WEBM", "WMV"},
	"SWF":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "SWF", "WEBM", "WMV"},
	"TS":   {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "TS", "WEBM", "WMV"},
	"VOB":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "VOB", "WEBM", "WMV"},
	"WEBM": {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"WMV":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV"},
	"WTV":  {"AAC", "AIFF", "FLAC", "M4A", "MP3", "WAV", "WMA", "GIF", "AVI", "FLV", "MKV", "MOV", "MP4", "WEBM", "WMV", "WTV"},
}

func CheckCanConvert(inputFormat, outputFormat string) error {
	var isSupported bool
	for _, format := range ConvertibleFormatsMap[strings.ToUpper(inputFormat)] {
		if outputFormat != strings.ToLower(format) {
			continue
		} else {
			isSupported = true
			break
		}
	}
	if isSupported {
		return nil
	} else {
		return errors.New("Can't convert")
	}
}

//PrintCorrespondingFormat print formats that inputFile is convertible when call convertibleformat flag
func PrintConvertibleFormat(inputFormat string) {
	fmt.Println(inputFormat + " >>")
	for i, format := range ConvertibleFormatsMap[strings.ToUpper(inputFormat)] {
		if format == "" {
			break
		}
		fmt.Print("・" + format + " ")
		if i%8 == 0 && i != 0 {
			fmt.Println()
		}
	}
}

//PrintFormats print Supported Format when call help flag
func PrintAllFormats() {
	formatKinds := [...]string{"Archive", "Audio", "Cad", "Document", "EBook", "Font", "Image", "PResentation", "Spredsheet", "Vector", "Video"}

	var formats [11][]string
	formats[0] = []string{"7Z", "ACE", "ALZ", "ARC", "ARJ", "BZ", "BZ2", "CAB", "CPIO", "DEB", "DMG", "GZ", "IMG", "ISO", "JAR", "LHA", "LZ", "LZMA", "LZO", "RAR", "RPM", "RZ", "TAR", "TAR.7Z", "TAR.BZ", "TAR.BZ2", "TAR.GZ", "TAR.LZO", "TAR.XZ", "TAR.Z", "TBZ", "TBZ2", "TGZ", "TZ", "TZO", "XZ", "Z", "ZIP"}
	formats[1] = []string{"AAC", "AC3", "AIF", "AIFC", "AIFF", "AMR", "AU", "CAF", "FLAC", "M4A", "M4B", "MP3", "OGA", "SFARK", "VOC", "WAV", "WEBA", "WMA"}
	formats[2] = []string{"DWG", "DXF"}
	formats[3] = []string{"ABW", "DJVU", "DOC", "DOCM", "DOCX", "DOT", "DOTX", "HTML", "HWP", "LWP", "MD", "ODT", "PAGES", "PDF", "RST", "RTF", "SDW", "TEX", "TXT", "WPD", "WPS", "ZABW"}
	formats[4] = []string{"AZW", "AZW3", "AZW4", "CBC", "CBR", "CBZ", "CHM", "EPUB", "FB2", "HTM", "HTMLZ", "LIT", "LRF", "MOBI", "PDB", "PML", "PRC", "RB", "SNB", "TCR", "TXTZ"}
	formats[5] = []string{"EOT", "OTF", "TTF", "WOFF", "WOFF2"}
	formats[6] = []string{"3FR", "ARW", "AVIF", "BMP", "CR2", "CR3", "CRW", "DCR", "DNG", "EPS", "ERF", "GIF", "HEIC", "ICNS", "ICO", "JFIF", "JPEG", "JPG", "MOS", "MRW", "NEF", "ODD", "ORF", "PEF", "PNG", "PPM", "PS", "PSD", "RAF", "RAW", "RW2", "TIF", "TIFF", "WEBP", "X3F", "XCF", "XPS"}
	formats[7] = []string{"DPS", "KEY", "ODP", "POT", "POTX", "PPS", "PPSX", "PPT", "PPTM", "PPTX", "SDA"}
	formats[8] = []string{"CSV", "ET", "NUMBERS", "ODS", "SDC", "XLS", "XLSM", "XLSX"}
	formats[9] = []string{"AI", "CDR", "CGM", "EMF", "SK", "SK1", "SVG", "SVGZ", "VSD", "WMF"}
	formats[10] = []string{"3G2", "3GP", "3GPP", "AVI", "CAVS", "DV", "DVR", "FLV", "M2TS", "M4V", "MKV", "MOD", "MOV", "MP4", "MPEG", "MPG", "MTS", "MXF", "OGG", "RM", "RMVB", "SWF", "TS", "VOB", "WEBM", "WMV", "WTV"}

	for i, kind := range formatKinds {
		fmt.Println(kind + " >")
		for j, format := range formats[i] {
			fmt.Print("・" + format)
			if j%8 == 0 && j != 0 {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}
