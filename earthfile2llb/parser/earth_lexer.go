// Code generated from earthfile2llb/parser/EarthLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 522,
	8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4,
	4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10,
	4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4,
	16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21,
	9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9,
	26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31,
	4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4,
	37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42,
	9, 42, 4, 43, 9, 43, 4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9,
	47, 4, 48, 9, 48, 4, 49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52,
	4, 53, 9, 53, 4, 54, 9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4,
	58, 9, 58, 4, 59, 9, 59, 4, 60, 9, 60, 4, 61, 9, 61, 3, 2, 6, 2, 130, 10,
	2, 13, 2, 14, 2, 131, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 6, 16, 271, 10, 16, 13, 16, 14,
	16, 272, 3, 16, 3, 16, 3, 17, 5, 17, 278, 10, 17, 3, 17, 5, 17, 281, 10,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 288, 10, 18, 12, 18, 14,
	18, 291, 11, 18, 3, 18, 6, 18, 294, 10, 18, 13, 18, 14, 18, 295, 3, 19,
	3, 19, 3, 19, 5, 19, 301, 10, 19, 3, 20, 3, 20, 7, 20, 305, 10, 20, 12,
	20, 14, 20, 308, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 405, 10, 40, 3, 41, 3, 41, 5, 41, 409,
	10, 41, 3, 41, 3, 41, 7, 41, 413, 10, 41, 12, 41, 14, 41, 416, 11, 41,
	3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 7, 42, 424, 10, 42, 12, 42, 14,
	42, 427, 11, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 5, 45,
	436, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49,
	5, 49, 457, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 54, 5, 54, 479, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 5, 57,
	496, 10, 57, 3, 57, 3, 57, 7, 57, 500, 10, 57, 12, 57, 14, 57, 503, 11,
	57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 5, 60, 512, 10, 60,
	3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 2, 2, 62,
	8, 5, 10, 6, 12, 7, 14, 8, 16, 9, 18, 10, 20, 11, 22, 12, 24, 13, 26, 14,
	28, 15, 30, 16, 32, 17, 34, 18, 36, 19, 38, 20, 40, 21, 42, 2, 44, 2, 46,
	2, 48, 2, 50, 2, 52, 2, 54, 2, 56, 2, 58, 2, 60, 2, 62, 2, 64, 2, 66, 2,
	68, 2, 70, 2, 72, 2, 74, 2, 76, 2, 78, 2, 80, 22, 82, 23, 84, 24, 86, 25,
	88, 2, 90, 2, 92, 2, 94, 2, 96, 2, 98, 2, 100, 2, 102, 2, 104, 2, 106,
	26, 108, 27, 110, 2, 112, 2, 114, 2, 116, 28, 118, 2, 120, 2, 122, 2, 124,
	2, 126, 2, 8, 2, 3, 4, 5, 6, 7, 11, 6, 2, 47, 48, 50, 59, 67, 92, 99, 124,
	3, 2, 67, 92, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 3, 2, 36, 36,
	7, 2, 11, 12, 15, 15, 34, 34, 36, 36, 93, 93, 6, 2, 11, 12, 15, 15, 34,
	34, 36, 36, 8, 2, 11, 12, 15, 15, 34, 34, 36, 36, 63, 63, 93, 93, 7, 2,
	11, 12, 15, 15, 34, 34, 36, 36, 63, 63, 2, 531, 2, 8, 3, 2, 2, 2, 2, 10,
	3, 2, 2, 2, 2, 12, 3, 2, 2, 2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2,
	18, 3, 2, 2, 2, 2, 20, 3, 2, 2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2,
	2, 26, 3, 2, 2, 2, 2, 28, 3, 2, 2, 2, 2, 30, 3, 2, 2, 2, 2, 32, 3, 2, 2,
	2, 2, 34, 3, 2, 2, 2, 2, 36, 3, 2, 2, 2, 2, 38, 3, 2, 2, 2, 2, 40, 3, 2,
	2, 2, 3, 46, 3, 2, 2, 2, 3, 48, 3, 2, 2, 2, 3, 50, 3, 2, 2, 2, 3, 52, 3,
	2, 2, 2, 3, 54, 3, 2, 2, 2, 3, 56, 3, 2, 2, 2, 3, 58, 3, 2, 2, 2, 3, 60,
	3, 2, 2, 2, 3, 62, 3, 2, 2, 2, 3, 64, 3, 2, 2, 2, 3, 66, 3, 2, 2, 2, 3,
	68, 3, 2, 2, 2, 3, 70, 3, 2, 2, 2, 3, 72, 3, 2, 2, 2, 3, 74, 3, 2, 2, 2,
	3, 76, 3, 2, 2, 2, 3, 78, 3, 2, 2, 2, 4, 80, 3, 2, 2, 2, 4, 82, 3, 2, 2,
	2, 4, 84, 3, 2, 2, 2, 4, 86, 3, 2, 2, 2, 4, 94, 3, 2, 2, 2, 4, 96, 3, 2,
	2, 2, 5, 98, 3, 2, 2, 2, 5, 100, 3, 2, 2, 2, 5, 102, 3, 2, 2, 2, 5, 104,
	3, 2, 2, 2, 6, 106, 3, 2, 2, 2, 6, 108, 3, 2, 2, 2, 6, 110, 3, 2, 2, 2,
	6, 112, 3, 2, 2, 2, 6, 114, 3, 2, 2, 2, 7, 116, 3, 2, 2, 2, 7, 118, 3,
	2, 2, 2, 7, 124, 3, 2, 2, 2, 7, 126, 3, 2, 2, 2, 8, 129, 3, 2, 2, 2, 10,
	137, 3, 2, 2, 2, 12, 144, 3, 2, 2, 2, 14, 151, 3, 2, 2, 2, 16, 167, 3,
	2, 2, 2, 18, 180, 3, 2, 2, 2, 20, 186, 3, 2, 2, 2, 22, 192, 3, 2, 2, 2,
	24, 198, 3, 2, 2, 2, 26, 206, 3, 2, 2, 2, 28, 216, 3, 2, 2, 2, 30, 229,
	3, 2, 2, 2, 32, 241, 3, 2, 2, 2, 34, 255, 3, 2, 2, 2, 36, 270, 3, 2, 2,
	2, 38, 277, 3, 2, 2, 2, 40, 293, 3, 2, 2, 2, 42, 300, 3, 2, 2, 2, 44, 302,
	3, 2, 2, 2, 46, 309, 3, 2, 2, 2, 48, 314, 3, 2, 2, 2, 50, 319, 3, 2, 2,
	2, 52, 324, 3, 2, 2, 2, 54, 329, 3, 2, 2, 2, 56, 334, 3, 2, 2, 2, 58, 339,
	3, 2, 2, 2, 60, 344, 3, 2, 2, 2, 62, 349, 3, 2, 2, 2, 64, 354, 3, 2, 2,
	2, 66, 359, 3, 2, 2, 2, 68, 364, 3, 2, 2, 2, 70, 369, 3, 2, 2, 2, 72, 374,
	3, 2, 2, 2, 74, 379, 3, 2, 2, 2, 76, 384, 3, 2, 2, 2, 78, 388, 3, 2, 2,
	2, 80, 392, 3, 2, 2, 2, 82, 396, 3, 2, 2, 2, 84, 400, 3, 2, 2, 2, 86, 408,
	3, 2, 2, 2, 88, 419, 3, 2, 2, 2, 90, 430, 3, 2, 2, 2, 92, 432, 3, 2, 2,
	2, 94, 435, 3, 2, 2, 2, 96, 442, 3, 2, 2, 2, 98, 446, 3, 2, 2, 2, 100,
	451, 3, 2, 2, 2, 102, 456, 3, 2, 2, 2, 104, 463, 3, 2, 2, 2, 106, 467,
	3, 2, 2, 2, 108, 471, 3, 2, 2, 2, 110, 473, 3, 2, 2, 2, 112, 478, 3, 2,
	2, 2, 114, 485, 3, 2, 2, 2, 116, 489, 3, 2, 2, 2, 118, 495, 3, 2, 2, 2,
	120, 506, 3, 2, 2, 2, 122, 508, 3, 2, 2, 2, 124, 511, 3, 2, 2, 2, 126,
	518, 3, 2, 2, 2, 128, 130, 9, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 131,
	3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 134, 7, 60, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 8, 2, 2, 2,
	136, 9, 3, 2, 2, 2, 137, 138, 7, 72, 2, 2, 138, 139, 7, 84, 2, 2, 139,
	140, 7, 81, 2, 2, 140, 141, 7, 79, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143,
	8, 3, 3, 2, 143, 11, 3, 2, 2, 2, 144, 145, 7, 69, 2, 2, 145, 146, 7, 81,
	2, 2, 146, 147, 7, 82, 2, 2, 147, 148, 7, 91, 2, 2, 148, 149, 3, 2, 2,
	2, 149, 150, 8, 4, 3, 2, 150, 13, 3, 2, 2, 2, 151, 152, 7, 85, 2, 2, 152,
	153, 7, 67, 2, 2, 153, 154, 7, 88, 2, 2, 154, 155, 7, 71, 2, 2, 155, 156,
	7, 34, 2, 2, 156, 157, 7, 67, 2, 2, 157, 158, 7, 84, 2, 2, 158, 159, 7,
	86, 2, 2, 159, 160, 7, 75, 2, 2, 160, 161, 7, 72, 2, 2, 161, 162, 7, 67,
	2, 2, 162, 163, 7, 69, 2, 2, 163, 164, 7, 86, 2, 2, 164, 165, 3, 2, 2,
	2, 165, 166, 8, 5, 3, 2, 166, 15, 3, 2, 2, 2, 167, 168, 7, 85, 2, 2, 168,
	169, 7, 67, 2, 2, 169, 170, 7, 88, 2, 2, 170, 171, 7, 71, 2, 2, 171, 172,
	7, 34, 2, 2, 172, 173, 7, 75, 2, 2, 173, 174, 7, 79, 2, 2, 174, 175, 7,
	67, 2, 2, 175, 176, 7, 73, 2, 2, 176, 177, 7, 71, 2, 2, 177, 178, 3, 2,
	2, 2, 178, 179, 8, 6, 3, 2, 179, 17, 3, 2, 2, 2, 180, 181, 7, 84, 2, 2,
	181, 182, 7, 87, 2, 2, 182, 183, 7, 80, 2, 2, 183, 184, 3, 2, 2, 2, 184,
	185, 8, 7, 3, 2, 185, 19, 3, 2, 2, 2, 186, 187, 7, 71, 2, 2, 187, 188,
	7, 80, 2, 2, 188, 189, 7, 88, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 8,
	8, 4, 2, 191, 21, 3, 2, 2, 2, 192, 193, 7, 67, 2, 2, 193, 194, 7, 84, 2,
	2, 194, 195, 7, 73, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 8, 9, 4, 2, 197,
	23, 3, 2, 2, 2, 198, 199, 7, 68, 2, 2, 199, 200, 7, 87, 2, 2, 200, 201,
	7, 75, 2, 2, 201, 202, 7, 78, 2, 2, 202, 203, 7, 70, 2, 2, 203, 204, 3,
	2, 2, 2, 204, 205, 8, 10, 3, 2, 205, 25, 3, 2, 2, 2, 206, 207, 7, 89, 2,
	2, 207, 208, 7, 81, 2, 2, 208, 209, 7, 84, 2, 2, 209, 210, 7, 77, 2, 2,
	210, 211, 7, 70, 2, 2, 211, 212, 7, 75, 2, 2, 212, 213, 7, 84, 2, 2, 213,
	214, 3, 2, 2, 2, 214, 215, 8, 11, 3, 2, 215, 27, 3, 2, 2, 2, 216, 217,
	7, 71, 2, 2, 217, 218, 7, 80, 2, 2, 218, 219, 7, 86, 2, 2, 219, 220, 7,
	84, 2, 2, 220, 221, 7, 91, 2, 2, 221, 222, 7, 82, 2, 2, 222, 223, 7, 81,
	2, 2, 223, 224, 7, 75, 2, 2, 224, 225, 7, 80, 2, 2, 225, 226, 7, 86, 2,
	2, 226, 227, 3, 2, 2, 2, 227, 228, 8, 12, 3, 2, 228, 29, 3, 2, 2, 2, 229,
	230, 7, 73, 2, 2, 230, 231, 7, 75, 2, 2, 231, 232, 7, 86, 2, 2, 232, 233,
	7, 34, 2, 2, 233, 234, 7, 69, 2, 2, 234, 235, 7, 78, 2, 2, 235, 236, 7,
	81, 2, 2, 236, 237, 7, 80, 2, 2, 237, 238, 7, 71, 2, 2, 238, 239, 3, 2,
	2, 2, 239, 240, 8, 13, 3, 2, 240, 31, 3, 2, 2, 2, 241, 242, 7, 70, 2, 2,
	242, 243, 7, 81, 2, 2, 243, 244, 7, 69, 2, 2, 244, 245, 7, 77, 2, 2, 245,
	246, 7, 71, 2, 2, 246, 247, 7, 84, 2, 2, 247, 248, 7, 34, 2, 2, 248, 249,
	7, 78, 2, 2, 249, 250, 7, 81, 2, 2, 250, 251, 7, 67, 2, 2, 251, 252, 7,
	70, 2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 8, 14, 3, 2, 254, 33, 3, 2, 2,
	2, 255, 256, 7, 70, 2, 2, 256, 257, 7, 81, 2, 2, 257, 258, 7, 69, 2, 2,
	258, 259, 7, 77, 2, 2, 259, 260, 7, 71, 2, 2, 260, 261, 7, 84, 2, 2, 261,
	262, 7, 34, 2, 2, 262, 263, 7, 82, 2, 2, 263, 264, 7, 87, 2, 2, 264, 265,
	7, 78, 2, 2, 265, 266, 7, 78, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 8,
	15, 3, 2, 268, 35, 3, 2, 2, 2, 269, 271, 9, 3, 2, 2, 270, 269, 3, 2, 2,
	2, 271, 272, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	274, 3, 2, 2, 2, 274, 275, 8, 16, 5, 2, 275, 37, 3, 2, 2, 2, 276, 278,
	5, 40, 18, 2, 277, 276, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 280, 3,
	2, 2, 2, 279, 281, 5, 44, 20, 2, 280, 279, 3, 2, 2, 2, 280, 281, 3, 2,
	2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 5, 42, 19, 2, 283, 39, 3, 2, 2, 2,
	284, 294, 9, 4, 2, 2, 285, 289, 7, 94, 2, 2, 286, 288, 9, 4, 2, 2, 287,
	286, 3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290,
	3, 2, 2, 2, 290, 292, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 294, 5, 42,
	19, 2, 293, 284, 3, 2, 2, 2, 293, 285, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2,
	295, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 41, 3, 2, 2, 2, 297, 301,
	9, 5, 2, 2, 298, 299, 7, 15, 2, 2, 299, 301, 7, 12, 2, 2, 300, 297, 3,
	2, 2, 2, 300, 298, 3, 2, 2, 2, 301, 43, 3, 2, 2, 2, 302, 306, 7, 37, 2,
	2, 303, 305, 10, 5, 2, 2, 304, 303, 3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306,
	304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 45, 3, 2, 2, 2, 308, 306, 3,
	2, 2, 2, 309, 310, 5, 8, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312, 8, 21, 6,
	2, 312, 313, 8, 21, 2, 2, 313, 47, 3, 2, 2, 2, 314, 315, 5, 10, 3, 2, 315,
	316, 3, 2, 2, 2, 316, 317, 8, 22, 7, 2, 317, 318, 8, 22, 3, 2, 318, 49,
	3, 2, 2, 2, 319, 320, 5, 12, 4, 2, 320, 321, 3, 2, 2, 2, 321, 322, 8, 23,
	8, 2, 322, 323, 8, 23, 3, 2, 323, 51, 3, 2, 2, 2, 324, 325, 5, 14, 5, 2,
	325, 326, 3, 2, 2, 2, 326, 327, 8, 24, 9, 2, 327, 328, 8, 24, 3, 2, 328,
	53, 3, 2, 2, 2, 329, 330, 5, 16, 6, 2, 330, 331, 3, 2, 2, 2, 331, 332,
	8, 25, 10, 2, 332, 333, 8, 25, 3, 2, 333, 55, 3, 2, 2, 2, 334, 335, 5,
	18, 7, 2, 335, 336, 3, 2, 2, 2, 336, 337, 8, 26, 11, 2, 337, 338, 8, 26,
	3, 2, 338, 57, 3, 2, 2, 2, 339, 340, 5, 20, 8, 2, 340, 341, 3, 2, 2, 2,
	341, 342, 8, 27, 12, 2, 342, 343, 8, 27, 4, 2, 343, 59, 3, 2, 2, 2, 344,
	345, 5, 22, 9, 2, 345, 346, 3, 2, 2, 2, 346, 347, 8, 28, 13, 2, 347, 348,
	8, 28, 4, 2, 348, 61, 3, 2, 2, 2, 349, 350, 5, 24, 10, 2, 350, 351, 3,
	2, 2, 2, 351, 352, 8, 29, 14, 2, 352, 353, 8, 29, 3, 2, 353, 63, 3, 2,
	2, 2, 354, 355, 5, 26, 11, 2, 355, 356, 3, 2, 2, 2, 356, 357, 8, 30, 15,
	2, 357, 358, 8, 30, 3, 2, 358, 65, 3, 2, 2, 2, 359, 360, 5, 28, 12, 2,
	360, 361, 3, 2, 2, 2, 361, 362, 8, 31, 16, 2, 362, 363, 8, 31, 3, 2, 363,
	67, 3, 2, 2, 2, 364, 365, 5, 30, 13, 2, 365, 366, 3, 2, 2, 2, 366, 367,
	8, 32, 17, 2, 367, 368, 8, 32, 3, 2, 368, 69, 3, 2, 2, 2, 369, 370, 5,
	32, 14, 2, 370, 371, 3, 2, 2, 2, 371, 372, 8, 33, 18, 2, 372, 373, 8, 33,
	3, 2, 373, 71, 3, 2, 2, 2, 374, 375, 5, 34, 15, 2, 375, 376, 3, 2, 2, 2,
	376, 377, 8, 34, 19, 2, 377, 378, 8, 34, 3, 2, 378, 73, 3, 2, 2, 2, 379,
	380, 5, 36, 16, 2, 380, 381, 3, 2, 2, 2, 381, 382, 8, 35, 20, 2, 382, 383,
	8, 35, 5, 2, 383, 75, 3, 2, 2, 2, 384, 385, 5, 38, 17, 2, 385, 386, 3,
	2, 2, 2, 386, 387, 8, 36, 21, 2, 387, 77, 3, 2, 2, 2, 388, 389, 5, 40,
	18, 2, 389, 390, 3, 2, 2, 2, 390, 391, 8, 37, 22, 2, 391, 79, 3, 2, 2,
	2, 392, 393, 7, 93, 2, 2, 393, 394, 3, 2, 2, 2, 394, 395, 8, 38, 23, 2,
	395, 81, 3, 2, 2, 2, 396, 397, 5, 84, 40, 2, 397, 398, 7, 63, 2, 2, 398,
	399, 5, 86, 41, 2, 399, 83, 3, 2, 2, 2, 400, 401, 7, 47, 2, 2, 401, 402,
	7, 47, 2, 2, 402, 404, 3, 2, 2, 2, 403, 405, 5, 86, 41, 2, 404, 403, 3,
	2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 85, 3, 2, 2, 2, 406, 409, 5, 90, 43,
	2, 407, 409, 5, 88, 42, 2, 408, 406, 3, 2, 2, 2, 408, 407, 3, 2, 2, 2,
	409, 414, 3, 2, 2, 2, 410, 413, 5, 92, 44, 2, 411, 413, 5, 88, 42, 2, 412,
	410, 3, 2, 2, 2, 412, 411, 3, 2, 2, 2, 413, 416, 3, 2, 2, 2, 414, 412,
	3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 417, 3, 2, 2, 2, 416, 414, 3, 2,
	2, 2, 417, 418, 8, 41, 24, 2, 418, 87, 3, 2, 2, 2, 419, 425, 7, 36, 2,
	2, 420, 424, 10, 6, 2, 2, 421, 422, 7, 94, 2, 2, 422, 424, 7, 36, 2, 2,
	423, 420, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 424, 427, 3, 2, 2, 2, 425,
	423, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 428, 3, 2, 2, 2, 427, 425,
	3, 2, 2, 2, 428, 429, 7, 36, 2, 2, 429, 89, 3, 2, 2, 2, 430, 431, 10, 7,
	2, 2, 431, 91, 3, 2, 2, 2, 432, 433, 10, 8, 2, 2, 433, 93, 3, 2, 2, 2,
	434, 436, 5, 40, 18, 2, 435, 434, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436,
	437, 3, 2, 2, 2, 437, 438, 5, 42, 19, 2, 438, 439, 3, 2, 2, 2, 439, 440,
	8, 45, 21, 2, 440, 441, 8, 45, 25, 2, 441, 95, 3, 2, 2, 2, 442, 443, 5,
	40, 18, 2, 443, 444, 3, 2, 2, 2, 444, 445, 8, 46, 22, 2, 445, 97, 3, 2,
	2, 2, 446, 447, 7, 93, 2, 2, 447, 448, 3, 2, 2, 2, 448, 449, 8, 47, 26,
	2, 449, 450, 8, 47, 23, 2, 450, 99, 3, 2, 2, 2, 451, 452, 5, 86, 41, 2,
	452, 453, 3, 2, 2, 2, 453, 454, 8, 48, 27, 2, 454, 101, 3, 2, 2, 2, 455,
	457, 5, 40, 18, 2, 456, 455, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 458,
	3, 2, 2, 2, 458, 459, 5, 42, 19, 2, 459, 460, 3, 2, 2, 2, 460, 461, 8,
	49, 21, 2, 461, 462, 8, 49, 25, 2, 462, 103, 3, 2, 2, 2, 463, 464, 5, 40,
	18, 2, 464, 465, 3, 2, 2, 2, 465, 466, 8, 50, 22, 2, 466, 105, 3, 2, 2,
	2, 467, 468, 7, 95, 2, 2, 468, 469, 3, 2, 2, 2, 469, 470, 8, 51, 25, 2,
	470, 107, 3, 2, 2, 2, 471, 472, 7, 46, 2, 2, 472, 109, 3, 2, 2, 2, 473,
	474, 5, 88, 42, 2, 474, 475, 3, 2, 2, 2, 475, 476, 8, 53, 27, 2, 476, 111,
	3, 2, 2, 2, 477, 479, 5, 40, 18, 2, 478, 477, 3, 2, 2, 2, 478, 479, 3,
	2, 2, 2, 479, 480, 3, 2, 2, 2, 480, 481, 5, 42, 19, 2, 481, 482, 3, 2,
	2, 2, 482, 483, 8, 54, 21, 2, 483, 484, 8, 54, 25, 2, 484, 113, 3, 2, 2,
	2, 485, 486, 5, 40, 18, 2, 486, 487, 3, 2, 2, 2, 487, 488, 8, 55, 22, 2,
	488, 115, 3, 2, 2, 2, 489, 490, 7, 63, 2, 2, 490, 491, 3, 2, 2, 2, 491,
	492, 8, 56, 24, 2, 492, 117, 3, 2, 2, 2, 493, 496, 5, 120, 58, 2, 494,
	496, 5, 88, 42, 2, 495, 493, 3, 2, 2, 2, 495, 494, 3, 2, 2, 2, 496, 501,
	3, 2, 2, 2, 497, 500, 5, 122, 59, 2, 498, 500, 5, 88, 42, 2, 499, 497,
	3, 2, 2, 2, 499, 498, 3, 2, 2, 2, 500, 503, 3, 2, 2, 2, 501, 499, 3, 2,
	2, 2, 501, 502, 3, 2, 2, 2, 502, 504, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2,
	504, 505, 8, 57, 27, 2, 505, 119, 3, 2, 2, 2, 506, 507, 10, 9, 2, 2, 507,
	121, 3, 2, 2, 2, 508, 509, 10, 10, 2, 2, 509, 123, 3, 2, 2, 2, 510, 512,
	5, 40, 18, 2, 511, 510, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 513, 3,
	2, 2, 2, 513, 514, 5, 42, 19, 2, 514, 515, 3, 2, 2, 2, 515, 516, 8, 60,
	21, 2, 516, 517, 8, 60, 25, 2, 517, 125, 3, 2, 2, 2, 518, 519, 5, 40, 18,
	2, 519, 520, 3, 2, 2, 2, 520, 521, 8, 61, 22, 2, 521, 127, 3, 2, 2, 2,
	31, 2, 3, 4, 5, 6, 7, 129, 131, 272, 277, 280, 289, 293, 295, 300, 306,
	404, 408, 412, 414, 423, 425, 435, 456, 478, 495, 499, 501, 511, 28, 7,
	3, 2, 7, 5, 2, 7, 7, 2, 7, 4, 2, 9, 5, 2, 9, 6, 2, 9, 7, 2, 9, 8, 2, 9,
	9, 2, 9, 10, 2, 9, 11, 2, 9, 12, 2, 9, 13, 2, 9, 14, 2, 9, 15, 2, 9, 16,
	2, 9, 17, 2, 9, 18, 2, 9, 19, 2, 9, 20, 2, 9, 21, 2, 7, 6, 2, 4, 5, 2,
	6, 2, 2, 9, 22, 2, 9, 25, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "RECIPE", "COMMAND_ARGS", "COMMAND_ARGS_ATOMS_ONLY", "COMMAND_BRACKETS",
	"COMMAND_ARGS_KEY_VALUE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'FROM'", "'COPY'", "'SAVE ARTIFACT'", "'SAVE IMAGE'",
	"'RUN'", "'ENV'", "'ARG'", "'BUILD'", "'WORKDIR'", "'ENTRYPOINT'", "'GIT CLONE'",
	"'DOCKER LOAD'", "'DOCKER PULL'", "", "", "", "'['", "", "", "", "']'",
	"','", "'='",
}

var lexerSymbolicNames = []string{
	"", "INDENT", "DEDENT", "Target", "FROM", "COPY", "SAVE_ARTIFACT", "SAVE_IMAGE",
	"RUN", "ENV", "ARG", "BUILD", "WORKDIR", "ENTRYPOINT", "GIT_CLONE", "DOCKER_LOAD",
	"DOCKER_PULL", "Command", "NL", "WS", "OPEN_BRACKET", "FlagKeyValue", "FlagKey",
	"Atom", "CLOSE_BRACKET", "COMMA", "EQUALS",
}

var lexerRuleNames = []string{
	"Target", "FROM", "COPY", "SAVE_ARTIFACT", "SAVE_IMAGE", "RUN", "ENV",
	"ARG", "BUILD", "WORKDIR", "ENTRYPOINT", "GIT_CLONE", "DOCKER_LOAD", "DOCKER_PULL",
	"Command", "NL", "WS", "CRLF", "COMMENT", "Target_R", "FROM_R", "COPY_R",
	"SAVE_ARTIFACT_R", "SAVE_IMAGE_R", "RUN_R", "ENV_R", "ARG_R", "BUILD_R",
	"WORKDIR_R", "ENTRYPOINT_R", "GIT_CLONE_R", "DOCKER_LOAD_R", "DOCKER_PULL_R",
	"Command_R", "NL_R", "WS_R", "OPEN_BRACKET", "FlagKeyValue", "FlagKey",
	"Atom", "QuotedAtom", "NonWSNLQuoteBracket", "NonWSNLQuote", "NL_C", "WS_C",
	"OPEN_BRACKET_CAAO", "Atom_CAAO", "NL_CAAO", "WS_CAAO", "CLOSE_BRACKET",
	"COMMA", "Atom_CB", "NL_CB", "WS_CB", "EQUALS", "Atom_CAKV", "NonWSNLQuoteBracket_CAKV",
	"NonWSNLQuote_CAKV", "NL_CAKV", "WS_CAKC",
}

type EarthLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEarthLexer(input antlr.CharStream) *EarthLexer {

	l := new(EarthLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "EarthLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EarthLexer tokens.
const (
	EarthLexerINDENT        = 1
	EarthLexerDEDENT        = 2
	EarthLexerTarget        = 3
	EarthLexerFROM          = 4
	EarthLexerCOPY          = 5
	EarthLexerSAVE_ARTIFACT = 6
	EarthLexerSAVE_IMAGE    = 7
	EarthLexerRUN           = 8
	EarthLexerENV           = 9
	EarthLexerARG           = 10
	EarthLexerBUILD         = 11
	EarthLexerWORKDIR       = 12
	EarthLexerENTRYPOINT    = 13
	EarthLexerGIT_CLONE     = 14
	EarthLexerDOCKER_LOAD   = 15
	EarthLexerDOCKER_PULL   = 16
	EarthLexerCommand       = 17
	EarthLexerNL            = 18
	EarthLexerWS            = 19
	EarthLexerOPEN_BRACKET  = 20
	EarthLexerFlagKeyValue  = 21
	EarthLexerFlagKey       = 22
	EarthLexerAtom          = 23
	EarthLexerCLOSE_BRACKET = 24
	EarthLexerCOMMA         = 25
	EarthLexerEQUALS        = 26
)

// EarthLexer modes.
const (
	EarthLexerRECIPE = iota + 1
	EarthLexerCOMMAND_ARGS
	EarthLexerCOMMAND_ARGS_ATOMS_ONLY
	EarthLexerCOMMAND_BRACKETS
	EarthLexerCOMMAND_ARGS_KEY_VALUE
)
