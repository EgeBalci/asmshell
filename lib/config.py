#!/usr/bin/env python2
# -*- coding: utf-8 -*-
from unicorn import *
from unicorn.x86_const import *

VERSION='0.04'
ADDRESS = 0x1000000
MEM_SIZE = 2 * 1024 * 1024 # 2MB
STACK_SIZE = 64 + 16
ESP_OFFSET=0x300000
# for push/pop
MERGIN_OFFSET=32
# now UNUSED
HIST_SIZE = 256

# 32bit registers
X86_REGS = [UC_X86_REG_EAX,UC_X86_REG_EBX,UC_X86_REG_ECX,UC_X86_REG_EDX,UC_X86_REG_ESI,UC_X86_REG_EDI,UC_X86_REG_EIP, UC_X86_REG_ESP, UC_X86_REG_EBP, UC_X86_REG_EFLAGS,UC_X86_REG_CS,UC_X86_REG_SS,UC_X86_REG_DS,UC_X86_REG_ES,UC_X86_REG_FS,UC_X86_REG_GS]
