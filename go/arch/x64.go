package arch

import (
    "github.com/keystone-engine/keystone/bindings/go/keystone"
    uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

// sample 
func SetX64() Machine {
    var mc Machine
    mc.Prompt = "(x64)> "
    mc.bit = 64
    mc.sp = uc.X86_REG_RSP
    mc.bp = uc.X86_REG_RBP

    mc.ks, _ = keystone.New(keystone.ARCH_X86, keystone.MODE_64)
    mc.ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_INTEL)

    mc.mu, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
    mc.mu.MemMap(0x0000, 0x200000)
    mc.mu.RegWrite(mc.sp, 0x100000)
    mc.mu.RegWrite(mc.bp, 0x80000)

    mc.oldMu, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
    mc.oldCtx, _ = mc.mu.ContextSave(nil)

    // if you want R8-15 register, add X86_REG_R8-15
    mc.regOrder = []string{
        "rax", "rip", "rbx", "eflags",
        "rcx", " cs", "rdx", " ss",
        "rsp", " ds", "rbp", " es",
        "rsi", " fs", "rdi", " gs",
    }
    mc.regs = map[string]int{
        "rax"    : uc.X86_REG_RAX,
        "rbx"    : uc.X86_REG_RBX,
        "rcx"    : uc.X86_REG_RCX,
        "rdx"    : uc.X86_REG_RDX,
        "rip"    : uc.X86_REG_RIP,
        "rsp"    : uc.X86_REG_RSP,
        "rbp"    : uc.X86_REG_RBP,
        "rsi"    : uc.X86_REG_RSI,
        "rdi"    : uc.X86_REG_RDI,
        "eflags" : uc.X86_REG_EFLAGS,
        " cs"    : uc.X86_REG_CS,
        " ss"    : uc.X86_REG_SS,
        " ds"    : uc.X86_REG_DS,
        " es"    : uc.X86_REG_ES,
        " fs"    : uc.X86_REG_FS,
        " gs"    : uc.X86_REG_GS,
    }
    return mc
}
