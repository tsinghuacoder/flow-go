// Code generated from parser/Strictus.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 48, 298,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 7, 38, 208, 10, 38, 12, 38, 14, 38, 211, 11, 38, 3, 39, 5, 39,
	214, 10, 39, 3, 40, 3, 40, 5, 40, 218, 10, 40, 3, 41, 3, 41, 7, 41, 222,
	10, 41, 12, 41, 14, 41, 225, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 6, 42,
	231, 10, 42, 13, 42, 14, 42, 232, 3, 43, 3, 43, 3, 43, 3, 43, 6, 43, 239,
	10, 43, 13, 43, 14, 43, 240, 3, 44, 3, 44, 3, 44, 3, 44, 6, 44, 247, 10,
	44, 13, 44, 14, 44, 248, 3, 45, 3, 45, 3, 45, 7, 45, 254, 10, 45, 12, 45,
	14, 45, 257, 11, 45, 3, 46, 6, 46, 260, 10, 46, 13, 46, 14, 46, 261, 3,
	46, 3, 46, 3, 47, 6, 47, 267, 10, 47, 13, 47, 14, 47, 268, 3, 47, 3, 47,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 7, 48, 278, 10, 48, 12, 48, 14, 48,
	281, 11, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3,
	49, 7, 49, 292, 10, 49, 12, 49, 14, 49, 295, 11, 49, 3, 49, 3, 49, 3, 279,
	2, 50, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75,
	39, 77, 2, 79, 2, 81, 40, 83, 41, 85, 42, 87, 43, 89, 44, 91, 45, 93, 46,
	95, 47, 97, 48, 3, 2, 12, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59,
	4, 2, 50, 59, 97, 97, 4, 2, 50, 51, 97, 97, 4, 2, 50, 57, 97, 97, 6, 2,
	50, 59, 67, 72, 97, 97, 99, 104, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 6, 2, 2, 2, 11, 11, 13, 14, 34, 34, 4, 2, 12, 12,
	15, 15, 2, 307, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 81, 3, 2,
	2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3,
	2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97,
	3, 2, 2, 2, 3, 99, 3, 2, 2, 2, 5, 101, 3, 2, 2, 2, 7, 103, 3, 2, 2, 2,
	9, 105, 3, 2, 2, 2, 11, 107, 3, 2, 2, 2, 13, 109, 3, 2, 2, 2, 15, 111,
	3, 2, 2, 2, 17, 113, 3, 2, 2, 2, 19, 115, 3, 2, 2, 2, 21, 118, 3, 2, 2,
	2, 23, 121, 3, 2, 2, 2, 25, 123, 3, 2, 2, 2, 27, 125, 3, 2, 2, 2, 29, 128,
	3, 2, 2, 2, 31, 131, 3, 2, 2, 2, 33, 133, 3, 2, 2, 2, 35, 135, 3, 2, 2,
	2, 37, 138, 3, 2, 2, 2, 39, 141, 3, 2, 2, 2, 41, 143, 3, 2, 2, 2, 43, 145,
	3, 2, 2, 2, 45, 147, 3, 2, 2, 2, 47, 149, 3, 2, 2, 2, 49, 151, 3, 2, 2,
	2, 51, 153, 3, 2, 2, 2, 53, 155, 3, 2, 2, 2, 55, 157, 3, 2, 2, 2, 57, 161,
	3, 2, 2, 2, 59, 165, 3, 2, 2, 2, 61, 172, 3, 2, 2, 2, 63, 176, 3, 2, 2,
	2, 65, 180, 3, 2, 2, 2, 67, 183, 3, 2, 2, 2, 69, 188, 3, 2, 2, 2, 71, 194,
	3, 2, 2, 2, 73, 199, 3, 2, 2, 2, 75, 205, 3, 2, 2, 2, 77, 213, 3, 2, 2,
	2, 79, 217, 3, 2, 2, 2, 81, 219, 3, 2, 2, 2, 83, 226, 3, 2, 2, 2, 85, 234,
	3, 2, 2, 2, 87, 242, 3, 2, 2, 2, 89, 250, 3, 2, 2, 2, 91, 259, 3, 2, 2,
	2, 93, 266, 3, 2, 2, 2, 95, 272, 3, 2, 2, 2, 97, 287, 3, 2, 2, 2, 99, 100,
	7, 60, 2, 2, 100, 4, 3, 2, 2, 2, 101, 102, 7, 46, 2, 2, 102, 6, 3, 2, 2,
	2, 103, 104, 7, 93, 2, 2, 104, 8, 3, 2, 2, 2, 105, 106, 7, 95, 2, 2, 106,
	10, 3, 2, 2, 2, 107, 108, 7, 125, 2, 2, 108, 12, 3, 2, 2, 2, 109, 110,
	7, 127, 2, 2, 110, 14, 3, 2, 2, 2, 111, 112, 7, 63, 2, 2, 112, 16, 3, 2,
	2, 2, 113, 114, 7, 65, 2, 2, 114, 18, 3, 2, 2, 2, 115, 116, 7, 126, 2,
	2, 116, 117, 7, 126, 2, 2, 117, 20, 3, 2, 2, 2, 118, 119, 7, 40, 2, 2,
	119, 120, 7, 40, 2, 2, 120, 22, 3, 2, 2, 2, 121, 122, 7, 48, 2, 2, 122,
	24, 3, 2, 2, 2, 123, 124, 7, 61, 2, 2, 124, 26, 3, 2, 2, 2, 125, 126, 7,
	63, 2, 2, 126, 127, 7, 63, 2, 2, 127, 28, 3, 2, 2, 2, 128, 129, 7, 35,
	2, 2, 129, 130, 7, 63, 2, 2, 130, 30, 3, 2, 2, 2, 131, 132, 7, 62, 2, 2,
	132, 32, 3, 2, 2, 2, 133, 134, 7, 64, 2, 2, 134, 34, 3, 2, 2, 2, 135, 136,
	7, 62, 2, 2, 136, 137, 7, 63, 2, 2, 137, 36, 3, 2, 2, 2, 138, 139, 7, 64,
	2, 2, 139, 140, 7, 63, 2, 2, 140, 38, 3, 2, 2, 2, 141, 142, 7, 45, 2, 2,
	142, 40, 3, 2, 2, 2, 143, 144, 7, 47, 2, 2, 144, 42, 3, 2, 2, 2, 145, 146,
	7, 44, 2, 2, 146, 44, 3, 2, 2, 2, 147, 148, 7, 49, 2, 2, 148, 46, 3, 2,
	2, 2, 149, 150, 7, 39, 2, 2, 150, 48, 3, 2, 2, 2, 151, 152, 7, 35, 2, 2,
	152, 50, 3, 2, 2, 2, 153, 154, 7, 42, 2, 2, 154, 52, 3, 2, 2, 2, 155, 156,
	7, 43, 2, 2, 156, 54, 3, 2, 2, 2, 157, 158, 7, 104, 2, 2, 158, 159, 7,
	119, 2, 2, 159, 160, 7, 112, 2, 2, 160, 56, 3, 2, 2, 2, 161, 162, 7, 114,
	2, 2, 162, 163, 7, 119, 2, 2, 163, 164, 7, 100, 2, 2, 164, 58, 3, 2, 2,
	2, 165, 166, 7, 116, 2, 2, 166, 167, 7, 103, 2, 2, 167, 168, 7, 118, 2,
	2, 168, 169, 7, 119, 2, 2, 169, 170, 7, 116, 2, 2, 170, 171, 7, 112, 2,
	2, 171, 60, 3, 2, 2, 2, 172, 173, 7, 110, 2, 2, 173, 174, 7, 103, 2, 2,
	174, 175, 7, 118, 2, 2, 175, 62, 3, 2, 2, 2, 176, 177, 7, 120, 2, 2, 177,
	178, 7, 99, 2, 2, 178, 179, 7, 116, 2, 2, 179, 64, 3, 2, 2, 2, 180, 181,
	7, 107, 2, 2, 181, 182, 7, 104, 2, 2, 182, 66, 3, 2, 2, 2, 183, 184, 7,
	103, 2, 2, 184, 185, 7, 110, 2, 2, 185, 186, 7, 117, 2, 2, 186, 187, 7,
	103, 2, 2, 187, 68, 3, 2, 2, 2, 188, 189, 7, 121, 2, 2, 189, 190, 7, 106,
	2, 2, 190, 191, 7, 107, 2, 2, 191, 192, 7, 110, 2, 2, 192, 193, 7, 103,
	2, 2, 193, 70, 3, 2, 2, 2, 194, 195, 7, 118, 2, 2, 195, 196, 7, 116, 2,
	2, 196, 197, 7, 119, 2, 2, 197, 198, 7, 103, 2, 2, 198, 72, 3, 2, 2, 2,
	199, 200, 7, 104, 2, 2, 200, 201, 7, 99, 2, 2, 201, 202, 7, 110, 2, 2,
	202, 203, 7, 117, 2, 2, 203, 204, 7, 103, 2, 2, 204, 74, 3, 2, 2, 2, 205,
	209, 5, 77, 39, 2, 206, 208, 5, 79, 40, 2, 207, 206, 3, 2, 2, 2, 208, 211,
	3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 76, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 212, 214, 9, 2, 2, 2, 213, 212, 3, 2, 2, 2,
	214, 78, 3, 2, 2, 2, 215, 218, 9, 3, 2, 2, 216, 218, 5, 77, 39, 2, 217,
	215, 3, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 80, 3, 2, 2, 2, 219, 223, 9,
	3, 2, 2, 220, 222, 9, 4, 2, 2, 221, 220, 3, 2, 2, 2, 222, 225, 3, 2, 2,
	2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 82, 3, 2, 2, 2, 225,
	223, 3, 2, 2, 2, 226, 227, 7, 50, 2, 2, 227, 228, 7, 100, 2, 2, 228, 230,
	3, 2, 2, 2, 229, 231, 9, 5, 2, 2, 230, 229, 3, 2, 2, 2, 231, 232, 3, 2,
	2, 2, 232, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 84, 3, 2, 2, 2,
	234, 235, 7, 50, 2, 2, 235, 236, 7, 113, 2, 2, 236, 238, 3, 2, 2, 2, 237,
	239, 9, 6, 2, 2, 238, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 238,
	3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 86, 3, 2, 2, 2, 242, 243, 7, 50,
	2, 2, 243, 244, 7, 122, 2, 2, 244, 246, 3, 2, 2, 2, 245, 247, 9, 7, 2,
	2, 246, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248,
	249, 3, 2, 2, 2, 249, 88, 3, 2, 2, 2, 250, 251, 7, 50, 2, 2, 251, 255,
	9, 8, 2, 2, 252, 254, 9, 9, 2, 2, 253, 252, 3, 2, 2, 2, 254, 257, 3, 2,
	2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 90, 3, 2, 2, 2,
	257, 255, 3, 2, 2, 2, 258, 260, 9, 10, 2, 2, 259, 258, 3, 2, 2, 2, 260,
	261, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 263,
	3, 2, 2, 2, 263, 264, 8, 46, 2, 2, 264, 92, 3, 2, 2, 2, 265, 267, 9, 11,
	2, 2, 266, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2,
	268, 269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 8, 47, 2, 2, 271,
	94, 3, 2, 2, 2, 272, 273, 7, 49, 2, 2, 273, 274, 7, 44, 2, 2, 274, 279,
	3, 2, 2, 2, 275, 278, 5, 95, 48, 2, 276, 278, 11, 2, 2, 2, 277, 275, 3,
	2, 2, 2, 277, 276, 3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 280, 3, 2, 2,
	2, 279, 277, 3, 2, 2, 2, 280, 282, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 282,
	283, 7, 44, 2, 2, 283, 284, 7, 49, 2, 2, 284, 285, 3, 2, 2, 2, 285, 286,
	8, 48, 2, 2, 286, 96, 3, 2, 2, 2, 287, 288, 7, 49, 2, 2, 288, 289, 7, 49,
	2, 2, 289, 293, 3, 2, 2, 2, 290, 292, 10, 11, 2, 2, 291, 290, 3, 2, 2,
	2, 292, 295, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294,
	296, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 296, 297, 8, 49, 2, 2, 297, 98,
	3, 2, 2, 2, 16, 2, 209, 213, 217, 223, 232, 240, 248, 255, 261, 268, 277,
	279, 293, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "','", "'['", "']'", "'{'", "'}'", "'='", "'?'", "'||'", "'&&'",
	"'.'", "';'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'!'", "'('", "')'", "'fun'", "'pub'", "'return'",
	"'let'", "'var'", "'if'", "'else'", "'while'", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Negate", "OpenParen", "CloseParen", "Fun", "Pub", "Return",
	"Let", "Var", "If", "Else", "While", "True", "False", "Identifier", "DecimalLiteral",
	"BinaryLiteral", "OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral",
	"WS", "Terminator", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "Equal", "Unequal", "Less", "Greater", "LessEqual",
	"GreaterEqual", "Plus", "Minus", "Mul", "Div", "Mod", "Negate", "OpenParen",
	"CloseParen", "Fun", "Pub", "Return", "Let", "Var", "If", "Else", "While",
	"True", "False", "Identifier", "IdentifierHead", "IdentifierCharacter",
	"DecimalLiteral", "BinaryLiteral", "OctalLiteral", "HexadecimalLiteral",
	"InvalidNumberLiteral", "WS", "Terminator", "BlockComment", "LineComment",
}

type StrictusLexer struct {
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

func NewStrictusLexer(input antlr.CharStream) *StrictusLexer {

	l := new(StrictusLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Strictus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StrictusLexer tokens.
const (
	StrictusLexerT__0                 = 1
	StrictusLexerT__1                 = 2
	StrictusLexerT__2                 = 3
	StrictusLexerT__3                 = 4
	StrictusLexerT__4                 = 5
	StrictusLexerT__5                 = 6
	StrictusLexerT__6                 = 7
	StrictusLexerT__7                 = 8
	StrictusLexerT__8                 = 9
	StrictusLexerT__9                 = 10
	StrictusLexerT__10                = 11
	StrictusLexerT__11                = 12
	StrictusLexerEqual                = 13
	StrictusLexerUnequal              = 14
	StrictusLexerLess                 = 15
	StrictusLexerGreater              = 16
	StrictusLexerLessEqual            = 17
	StrictusLexerGreaterEqual         = 18
	StrictusLexerPlus                 = 19
	StrictusLexerMinus                = 20
	StrictusLexerMul                  = 21
	StrictusLexerDiv                  = 22
	StrictusLexerMod                  = 23
	StrictusLexerNegate               = 24
	StrictusLexerOpenParen            = 25
	StrictusLexerCloseParen           = 26
	StrictusLexerFun                  = 27
	StrictusLexerPub                  = 28
	StrictusLexerReturn               = 29
	StrictusLexerLet                  = 30
	StrictusLexerVar                  = 31
	StrictusLexerIf                   = 32
	StrictusLexerElse                 = 33
	StrictusLexerWhile                = 34
	StrictusLexerTrue                 = 35
	StrictusLexerFalse                = 36
	StrictusLexerIdentifier           = 37
	StrictusLexerDecimalLiteral       = 38
	StrictusLexerBinaryLiteral        = 39
	StrictusLexerOctalLiteral         = 40
	StrictusLexerHexadecimalLiteral   = 41
	StrictusLexerInvalidNumberLiteral = 42
	StrictusLexerWS                   = 43
	StrictusLexerTerminator           = 44
	StrictusLexerBlockComment         = 45
	StrictusLexerLineComment          = 46
)
