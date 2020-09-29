// Code generated by "stringer -type=RelocType"; DO NOT EDIT.

package objabi

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[R_ADDR-1]
	_ = x[R_ADDRPOWER-2]
	_ = x[R_ADDRARM64-3]
	_ = x[R_ADDRMIPS-4]
	_ = x[R_ADDROFF-5]
	_ = x[R_WEAKADDROFF-6]
	_ = x[R_SIZE-7]
	_ = x[R_CALL-8]
	_ = x[R_CALLARM-9]
	_ = x[R_CALLARM64-10]
	_ = x[R_CALLIND-11]
	_ = x[R_CALLPOWER-12]
	_ = x[R_CALLMIPS-13]
	_ = x[R_CALLRISCV-14]
	_ = x[R_CONST-15]
	_ = x[R_PCREL-16]
	_ = x[R_TLS_LE-17]
	_ = x[R_TLS_IE-18]
	_ = x[R_GOTOFF-19]
	_ = x[R_PLT0-20]
	_ = x[R_PLT1-21]
	_ = x[R_PLT2-22]
	_ = x[R_USEFIELD-23]
	_ = x[R_USETYPE-24]
	_ = x[R_USEIFACE-25]
	_ = x[R_METHODOFF-26]
	_ = x[R_POWER_TOC-27]
	_ = x[R_GOTPCREL-28]
	_ = x[R_JMPMIPS-29]
	_ = x[R_DWARFSECREF-30]
	_ = x[R_DWARFFILEREF-31]
	_ = x[R_ARM64_TLS_LE-32]
	_ = x[R_ARM64_TLS_IE-33]
	_ = x[R_ARM64_GOTPCREL-34]
	_ = x[R_ARM64_GOT-35]
	_ = x[R_ARM64_PCREL-36]
	_ = x[R_ARM64_LDST8-37]
	_ = x[R_ARM64_LDST32-38]
	_ = x[R_ARM64_LDST64-39]
	_ = x[R_ARM64_LDST128-40]
	_ = x[R_POWER_TLS_LE-41]
	_ = x[R_POWER_TLS_IE-42]
	_ = x[R_POWER_TLS-43]
	_ = x[R_ADDRPOWER_DS-44]
	_ = x[R_ADDRPOWER_GOT-45]
	_ = x[R_ADDRPOWER_PCREL-46]
	_ = x[R_ADDRPOWER_TOCREL-47]
	_ = x[R_ADDRPOWER_TOCREL_DS-48]
	_ = x[R_RISCV_PCREL_ITYPE-49]
	_ = x[R_RISCV_PCREL_STYPE-50]
	_ = x[R_PCRELDBL-51]
	_ = x[R_ADDRMIPSU-52]
	_ = x[R_ADDRMIPSTLS-53]
	_ = x[R_ADDRCUOFF-54]
	_ = x[R_WASMIMPORT-55]
	_ = x[R_XCOFFREF-56]
}

const _RelocType_name = "R_ADDRR_ADDRPOWERR_ADDRARM64R_ADDRMIPSR_ADDROFFR_WEAKADDROFFR_SIZER_CALLR_CALLARMR_CALLARM64R_CALLINDR_CALLPOWERR_CALLMIPSR_CALLRISCVR_CONSTR_PCRELR_TLS_LER_TLS_IER_GOTOFFR_PLT0R_PLT1R_PLT2R_USEFIELDR_USETYPER_USEIFACER_METHODOFFR_POWER_TOCR_GOTPCRELR_JMPMIPSR_DWARFSECREFR_DWARFFILEREFR_ARM64_TLS_LER_ARM64_TLS_IER_ARM64_GOTPCRELR_ARM64_GOTR_ARM64_PCRELR_ARM64_LDST8R_ARM64_LDST32R_ARM64_LDST64R_ARM64_LDST128R_POWER_TLS_LER_POWER_TLS_IER_POWER_TLSR_ADDRPOWER_DSR_ADDRPOWER_GOTR_ADDRPOWER_PCRELR_ADDRPOWER_TOCRELR_ADDRPOWER_TOCREL_DSR_RISCV_PCREL_ITYPER_RISCV_PCREL_STYPER_PCRELDBLR_ADDRMIPSUR_ADDRMIPSTLSR_ADDRCUOFFR_WASMIMPORTR_XCOFFREF"

var _RelocType_index = [...]uint16{0, 6, 17, 28, 38, 47, 60, 66, 72, 81, 92, 101, 112, 122, 133, 140, 147, 155, 163, 171, 177, 183, 189, 199, 208, 218, 229, 240, 250, 259, 272, 286, 300, 314, 330, 341, 354, 367, 381, 395, 410, 424, 438, 449, 463, 478, 495, 513, 534, 553, 572, 582, 593, 606, 617, 629, 639}

func (i RelocType) String() string {
	i -= 1
	if i < 0 || i >= RelocType(len(_RelocType_index)-1) {
		return "RelocType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RelocType_name[_RelocType_index[i]:_RelocType_index[i+1]]
}
