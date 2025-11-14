// -*- mode: go; -*-

// Integer constants defined in linux/input.h and linux/input-event-codes.h can be accessed
// directly as constants or through the provided reverse and forward mappings:

//
//   evdev.KEY_A  // 30
//   evdev.ecodes["KEY_A"]  // 30
//   evdev.KEY[30] // "KEY_A"
//   evdev.REL[0]  // "REL_X"
//   evdev.EV[evdev.EV_KEY]  // "EV_KEY"
//   evdev.ByEventType[EV_REL][0]  // "REL_X"
//
// Generated on: Linux 4.7.4-200.fc24.x86_64 #1 SMP Thu Sep 15 18:42:09 UTC 2016 x86_64

package evdev

import "strings"

const (
	EV_VERSION                   = 0x010001
	ID_BUS                       = 0
	ID_VENDOR                    = 1
	ID_PRODUCT                   = 2
	ID_VERSION                   = 3
	BUS_PCI                      = 0x01
	BUS_ISAPNP                   = 0x02
	BUS_USB                      = 0x03
	BUS_HIL                      = 0x04
	BUS_BLUETOOTH                = 0x05
	BUS_VIRTUAL                  = 0x06
	BUS_ISA                      = 0x10
	BUS_I8042                    = 0x11
	BUS_XTKBD                    = 0x12
	BUS_RS232                    = 0x13
	BUS_GAMEPORT                 = 0x14
	BUS_PARPORT                  = 0x15
	BUS_AMIGA                    = 0x16
	BUS_ADB                      = 0x17
	BUS_I2C                      = 0x18
	BUS_HOST                     = 0x19
	BUS_GSC                      = 0x1A
	BUS_ATARI                    = 0x1B
	BUS_SPI                      = 0x1C
	BUS_RMI                      = 0x1D
	BUS_CEC                      = 0x1E
	FF_STATUS_STOPPED            = 0x00
	FF_STATUS_PLAYING            = 0x01
	FF_STATUS_MAX                = 0x01
	FF_RUMBLE                    = 0x50
	FF_PERIODIC                  = 0x51
	FF_CONSTANT                  = 0x52
	FF_SPRING                    = 0x53
	FF_FRICTION                  = 0x54
	FF_DAMPER                    = 0x55
	FF_INERTIA                   = 0x56
	FF_RAMP                      = 0x57
	FF_EFFECT_MIN                = FF_RUMBLE
	FF_EFFECT_MAX                = FF_RAMP
	FF_SQUARE                    = 0x58
	FF_TRIANGLE                  = 0x59
	FF_SINE                      = 0x5a
	FF_SAW_UP                    = 0x5b
	FF_SAW_DOWN                  = 0x5c
	FF_CUSTOM                    = 0x5d
	FF_WAVEFORM_MIN              = FF_SQUARE
	FF_WAVEFORM_MAX              = FF_CUSTOM
	FF_GAIN                      = 0x60
	FF_AUTOCENTER                = 0x61
	FF_MAX_EFFECTS               = FF_GAIN
	FF_MAX                       = 0x7f
	EV_SYN                       = 0x00
	EV_KEY                       = 0x01
	EV_REL                       = 0x02
	EV_ABS                       = 0x03
	EV_MSC                       = 0x04
	EV_SW                        = 0x05
	EV_LED                       = 0x11
	EV_SND                       = 0x12
	EV_REP                       = 0x14
	EV_FF                        = 0x15
	EV_PWR                       = 0x16
	EV_FF_STATUS                 = 0x17
	EV_MAX                       = 0x1f
	SYN_REPORT                   = 0
	SYN_CONFIG                   = 1
	SYN_MT_REPORT                = 2
	SYN_DROPPED                  = 3
	SYN_MAX                      = 0xf
	KEY_RESERVED                 = 0
	KEY_ESC                      = 1
	KEY_1                        = 2
	KEY_2                        = 3
	KEY_3                        = 4
	KEY_4                        = 5
	KEY_5                        = 6
	KEY_6                        = 7
	KEY_7                        = 8
	KEY_8                        = 9
	KEY_9                        = 10
	KEY_0                        = 11
	KEY_MINUS                    = 12
	KEY_EQUAL                    = 13
	KEY_BACKSPACE                = 14
	KEY_TAB                      = 15
	KEY_Q                        = 16
	KEY_W                        = 17
	KEY_E                        = 18
	KEY_R                        = 19
	KEY_T                        = 20
	KEY_Y                        = 21
	KEY_U                        = 22
	KEY_I                        = 23
	KEY_O                        = 24
	KEY_P                        = 25
	KEY_LEFTBRACE                = 26
	KEY_RIGHTBRACE               = 27
	KEY_ENTER                    = 28
	KEY_LEFTCTRL                 = 29
	KEY_A                        = 30
	KEY_S                        = 31
	KEY_D                        = 32
	KEY_F                        = 33
	KEY_G                        = 34
	KEY_H                        = 35
	KEY_J                        = 36
	KEY_K                        = 37
	KEY_L                        = 38
	KEY_SEMICOLON                = 39
	KEY_APOSTROPHE               = 40
	KEY_GRAVE                    = 41
	KEY_LEFTSHIFT                = 42
	KEY_BACKSLASH                = 43
	KEY_Z                        = 44
	KEY_X                        = 45
	KEY_C                        = 46
	KEY_V                        = 47
	KEY_B                        = 48
	KEY_N                        = 49
	KEY_M                        = 50
	KEY_COMMA                    = 51
	KEY_DOT                      = 52
	KEY_SLASH                    = 53
	KEY_RIGHTSHIFT               = 54
	KEY_KPASTERISK               = 55
	KEY_LEFTALT                  = 56
	KEY_SPACE                    = 57
	KEY_CAPSLOCK                 = 58
	KEY_F1                       = 59
	KEY_F2                       = 60
	KEY_F3                       = 61
	KEY_F4                       = 62
	KEY_F5                       = 63
	KEY_F6                       = 64
	KEY_F7                       = 65
	KEY_F8                       = 66
	KEY_F9                       = 67
	KEY_F10                      = 68
	KEY_NUMLOCK                  = 69
	KEY_SCROLLLOCK               = 70
	KEY_KP7                      = 71
	KEY_KP8                      = 72
	KEY_KP9                      = 73
	KEY_KPMINUS                  = 74
	KEY_KP4                      = 75
	KEY_KP5                      = 76
	KEY_KP6                      = 77
	KEY_KPPLUS                   = 78
	KEY_KP1                      = 79
	KEY_KP2                      = 80
	KEY_KP3                      = 81
	KEY_KP0                      = 82
	KEY_KPDOT                    = 83
	KEY_ZENKAKUHANKAKU           = 85
	KEY_102ND                    = 86
	KEY_F11                      = 87
	KEY_F12                      = 88
	KEY_RO                       = 89
	KEY_KATAKANA                 = 90
	KEY_HIRAGANA                 = 91
	KEY_HENKAN                   = 92
	KEY_KATAKANAHIRAGANA         = 93
	KEY_MUHENKAN                 = 94
	KEY_KPJPCOMMA                = 95
	KEY_KPENTER                  = 96
	KEY_RIGHTCTRL                = 97
	KEY_KPSLASH                  = 98
	KEY_SYSRQ                    = 99
	KEY_RIGHTALT                 = 100
	KEY_LINEFEED                 = 101
	KEY_HOME                     = 102
	KEY_UP                       = 103
	KEY_PAGEUP                   = 104
	KEY_LEFT                     = 105
	KEY_RIGHT                    = 106
	KEY_END                      = 107
	KEY_DOWN                     = 108
	KEY_PAGEDOWN                 = 109
	KEY_INSERT                   = 110
	KEY_DELETE                   = 111
	KEY_MACRO                    = 112
	KEY_MUTE                     = 113
	KEY_VOLUMEDOWN               = 114
	KEY_VOLUMEUP                 = 115
	KEY_POWER                    = 116
	KEY_KPEQUAL                  = 117
	KEY_KPPLUSMINUS              = 118
	KEY_PAUSE                    = 119
	KEY_SCALE                    = 120
	KEY_KPCOMMA                  = 121
	KEY_HANGEUL                  = 122
	KEY_HANGUEL                  = KEY_HANGEUL
	KEY_HANJA                    = 123
	KEY_YEN                      = 124
	KEY_LEFTMETA                 = 125
	KEY_RIGHTMETA                = 126
	KEY_COMPOSE                  = 127
	KEY_STOP                     = 128
	KEY_AGAIN                    = 129
	KEY_PROPS                    = 130
	KEY_UNDO                     = 131
	KEY_FRONT                    = 132
	KEY_COPY                     = 133
	KEY_OPEN                     = 134
	KEY_PASTE                    = 135
	KEY_FIND                     = 136
	KEY_CUT                      = 137
	KEY_HELP                     = 138
	KEY_MENU                     = 139
	KEY_CALC                     = 140
	KEY_SETUP                    = 141
	KEY_SLEEP                    = 142
	KEY_WAKEUP                   = 143
	KEY_FILE                     = 144
	KEY_SENDFILE                 = 145
	KEY_DELETEFILE               = 146
	KEY_XFER                     = 147
	KEY_PROG1                    = 148
	KEY_PROG2                    = 149
	KEY_WWW                      = 150
	KEY_MSDOS                    = 151
	KEY_COFFEE                   = 152
	KEY_SCREENLOCK               = KEY_COFFEE
	KEY_ROTATE_DISPLAY           = 153
	KEY_DIRECTION                = KEY_ROTATE_DISPLAY
	KEY_CYCLEWINDOWS             = 154
	KEY_MAIL                     = 155
	KEY_BOOKMARKS                = 156
	KEY_COMPUTER                 = 157
	KEY_BACK                     = 158
	KEY_FORWARD                  = 159
	KEY_CLOSECD                  = 160
	KEY_EJECTCD                  = 161
	KEY_EJECTCLOSECD             = 162
	KEY_NEXTSONG                 = 163
	KEY_PLAYPAUSE                = 164
	KEY_PREVIOUSSONG             = 165
	KEY_STOPCD                   = 166
	KEY_RECORD                   = 167
	KEY_REWIND                   = 168
	KEY_PHONE                    = 169
	KEY_ISO                      = 170
	KEY_CONFIG                   = 171
	KEY_HOMEPAGE                 = 172
	KEY_REFRESH                  = 173
	KEY_EXIT                     = 174
	KEY_MOVE                     = 175
	KEY_EDIT                     = 176
	KEY_SCROLLUP                 = 177
	KEY_SCROLLDOWN               = 178
	KEY_KPLEFTPAREN              = 179
	KEY_KPRIGHTPAREN             = 180
	KEY_NEW                      = 181
	KEY_REDO                     = 182
	KEY_F13                      = 183
	KEY_F14                      = 184
	KEY_F15                      = 185
	KEY_F16                      = 186
	KEY_F17                      = 187
	KEY_F18                      = 188
	KEY_F19                      = 189
	KEY_F20                      = 190
	KEY_F21                      = 191
	KEY_F22                      = 192
	KEY_F23                      = 193
	KEY_F24                      = 194
	KEY_PLAYCD                   = 200
	KEY_PAUSECD                  = 201
	KEY_PROG3                    = 202
	KEY_PROG4                    = 203
	KEY_DASHBOARD                = 204
	KEY_SUSPEND                  = 205
	KEY_CLOSE                    = 206
	KEY_PLAY                     = 207
	KEY_FASTFORWARD              = 208
	KEY_BASSBOOST                = 209
	KEY_PRINT                    = 210
	KEY_HP                       = 211
	KEY_CAMERA                   = 212
	KEY_SOUND                    = 213
	KEY_QUESTION                 = 214
	KEY_EMAIL                    = 215
	KEY_CHAT                     = 216
	KEY_SEARCH                   = 217
	KEY_CONNECT                  = 218
	KEY_FINANCE                  = 219
	KEY_SPORT                    = 220
	KEY_SHOP                     = 221
	KEY_ALTERASE                 = 222
	KEY_CANCEL                   = 223
	KEY_BRIGHTNESSDOWN           = 224
	KEY_BRIGHTNESSUP             = 225
	KEY_MEDIA                    = 226
	KEY_SWITCHVIDEOMODE          = 227
	KEY_KBDILLUMTOGGLE           = 228
	KEY_KBDILLUMDOWN             = 229
	KEY_KBDILLUMUP               = 230
	KEY_SEND                     = 231
	KEY_REPLY                    = 232
	KEY_FORWARDMAIL              = 233
	KEY_SAVE                     = 234
	KEY_DOCUMENTS                = 235
	KEY_BATTERY                  = 236
	KEY_BLUETOOTH                = 237
	KEY_WLAN                     = 238
	KEY_UWB                      = 239
	KEY_UNKNOWN                  = 240
	KEY_VIDEO_NEXT               = 241
	KEY_VIDEO_PREV               = 242
	KEY_BRIGHTNESS_CYCLE         = 243
	KEY_BRIGHTNESS_AUTO          = 244
	KEY_BRIGHTNESS_ZERO          = KEY_BRIGHTNESS_AUTO
	KEY_DISPLAY_OFF              = 245
	KEY_WWAN                     = 246
	KEY_WIMAX                    = KEY_WWAN
	KEY_RFKILL                   = 247
	KEY_MICMUTE                  = 248
	BTN_MISC                     = 0x100
	BTN_0                        = 0x100
	BTN_1                        = 0x101
	BTN_2                        = 0x102
	BTN_3                        = 0x103
	BTN_4                        = 0x104
	BTN_5                        = 0x105
	BTN_6                        = 0x106
	BTN_7                        = 0x107
	BTN_8                        = 0x108
	BTN_9                        = 0x109
	BTN_MOUSE                    = 0x110
	BTN_LEFT                     = 0x110
	BTN_RIGHT                    = 0x111
	BTN_MIDDLE                   = 0x112
	BTN_SIDE                     = 0x113
	BTN_EXTRA                    = 0x114
	BTN_FORWARD                  = 0x115
	BTN_BACK                     = 0x116
	BTN_TASK                     = 0x117
	BTN_JOYSTICK                 = 0x120
	BTN_TRIGGER                  = 0x120
	BTN_THUMB                    = 0x121
	BTN_THUMB2                   = 0x122
	BTN_TOP                      = 0x123
	BTN_TOP2                     = 0x124
	BTN_PINKIE                   = 0x125
	BTN_BASE                     = 0x126
	BTN_BASE2                    = 0x127
	BTN_BASE3                    = 0x128
	BTN_BASE4                    = 0x129
	BTN_BASE5                    = 0x12a
	BTN_BASE6                    = 0x12b
	BTN_DEAD                     = 0x12f
	BTN_GAMEPAD                  = 0x130
	BTN_SOUTH                    = 0x130
	BTN_A                        = BTN_SOUTH
	BTN_EAST                     = 0x131
	BTN_B                        = BTN_EAST
	BTN_C                        = 0x132
	BTN_NORTH                    = 0x133
	BTN_X                        = BTN_NORTH
	BTN_WEST                     = 0x134
	BTN_Y                        = BTN_WEST
	BTN_Z                        = 0x135
	BTN_TL                       = 0x136
	BTN_TR                       = 0x137
	BTN_TL2                      = 0x138
	BTN_TR2                      = 0x139
	BTN_SELECT                   = 0x13a
	BTN_START                    = 0x13b
	BTN_MODE                     = 0x13c
	BTN_THUMBL                   = 0x13d
	BTN_THUMBR                   = 0x13e
	BTN_DIGI                     = 0x140
	BTN_TOOL_PEN                 = 0x140
	BTN_TOOL_RUBBER              = 0x141
	BTN_TOOL_BRUSH               = 0x142
	BTN_TOOL_PENCIL              = 0x143
	BTN_TOOL_AIRBRUSH            = 0x144
	BTN_TOOL_FINGER              = 0x145
	BTN_TOOL_MOUSE               = 0x146
	BTN_TOOL_LENS                = 0x147
	BTN_TOOL_QUINTTAP            = 0x148
	BTN_TOUCH                    = 0x14a
	BTN_STYLUS                   = 0x14b
	BTN_STYLUS2                  = 0x14c
	BTN_TOOL_DOUBLETAP           = 0x14d
	BTN_TOOL_TRIPLETAP           = 0x14e
	BTN_TOOL_QUADTAP             = 0x14f
	BTN_WHEEL                    = 0x150
	BTN_GEAR_DOWN                = 0x150
	BTN_GEAR_UP                  = 0x151
	KEY_OK                       = 0x160
	KEY_SELECT                   = 0x161
	KEY_GOTO                     = 0x162
	KEY_CLEAR                    = 0x163
	KEY_POWER2                   = 0x164
	KEY_OPTION                   = 0x165
	KEY_INFO                     = 0x166
	KEY_TIME                     = 0x167
	KEY_VENDOR                   = 0x168
	KEY_ARCHIVE                  = 0x169
	KEY_PROGRAM                  = 0x16a
	KEY_CHANNEL                  = 0x16b
	KEY_FAVORITES                = 0x16c
	KEY_EPG                      = 0x16d
	KEY_PVR                      = 0x16e
	KEY_MHP                      = 0x16f
	KEY_LANGUAGE                 = 0x170
	KEY_TITLE                    = 0x171
	KEY_SUBTITLE                 = 0x172
	KEY_ANGLE                    = 0x173
	KEY_ZOOM                     = 0x174
	KEY_MODE                     = 0x175
	KEY_KEYBOARD                 = 0x176
	KEY_SCREEN                   = 0x177
	KEY_PC                       = 0x178
	KEY_TV                       = 0x179
	KEY_TV2                      = 0x17a
	KEY_VCR                      = 0x17b
	KEY_VCR2                     = 0x17c
	KEY_SAT                      = 0x17d
	KEY_SAT2                     = 0x17e
	KEY_CD                       = 0x17f
	KEY_TAPE                     = 0x180
	KEY_RADIO                    = 0x181
	KEY_TUNER                    = 0x182
	KEY_PLAYER                   = 0x183
	KEY_TEXT                     = 0x184
	KEY_DVD                      = 0x185
	KEY_AUX                      = 0x186
	KEY_MP3                      = 0x187
	KEY_AUDIO                    = 0x188
	KEY_VIDEO                    = 0x189
	KEY_DIRECTORY                = 0x18a
	KEY_LIST                     = 0x18b
	KEY_MEMO                     = 0x18c
	KEY_CALENDAR                 = 0x18d
	KEY_RED                      = 0x18e
	KEY_GREEN                    = 0x18f
	KEY_YELLOW                   = 0x190
	KEY_BLUE                     = 0x191
	KEY_CHANNELUP                = 0x192
	KEY_CHANNELDOWN              = 0x193
	KEY_FIRST                    = 0x194
	KEY_LAST                     = 0x195
	KEY_AB                       = 0x196
	KEY_NEXT                     = 0x197
	KEY_RESTART                  = 0x198
	KEY_SLOW                     = 0x199
	KEY_SHUFFLE                  = 0x19a
	KEY_BREAK                    = 0x19b
	KEY_PREVIOUS                 = 0x19c
	KEY_DIGITS                   = 0x19d
	KEY_TEEN                     = 0x19e
	KEY_TWEN                     = 0x19f
	KEY_VIDEOPHONE               = 0x1a0
	KEY_GAMES                    = 0x1a1
	KEY_ZOOMIN                   = 0x1a2
	KEY_ZOOMOUT                  = 0x1a3
	KEY_ZOOMRESET                = 0x1a4
	KEY_WORDPROCESSOR            = 0x1a5
	KEY_EDITOR                   = 0x1a6
	KEY_SPREADSHEET              = 0x1a7
	KEY_GRAPHICSEDITOR           = 0x1a8
	KEY_PRESENTATION             = 0x1a9
	KEY_DATABASE                 = 0x1aa
	KEY_NEWS                     = 0x1ab
	KEY_VOICEMAIL                = 0x1ac
	KEY_ADDRESSBOOK              = 0x1ad
	KEY_MESSENGER                = 0x1ae
	KEY_DISPLAYTOGGLE            = 0x1af
	KEY_BRIGHTNESS_TOGGLE        = KEY_DISPLAYTOGGLE
	KEY_SPELLCHECK               = 0x1b0
	KEY_LOGOFF                   = 0x1b1
	KEY_DOLLAR                   = 0x1b2
	KEY_EURO                     = 0x1b3
	KEY_FRAMEBACK                = 0x1b4
	KEY_FRAMEFORWARD             = 0x1b5
	KEY_CONTEXT_MENU             = 0x1b6
	KEY_MEDIA_REPEAT             = 0x1b7
	KEY_10CHANNELSUP             = 0x1b8
	KEY_10CHANNELSDOWN           = 0x1b9
	KEY_IMAGES                   = 0x1ba
	KEY_DEL_EOL                  = 0x1c0
	KEY_DEL_EOS                  = 0x1c1
	KEY_INS_LINE                 = 0x1c2
	KEY_DEL_LINE                 = 0x1c3
	KEY_FN                       = 0x1d0
	KEY_FN_ESC                   = 0x1d1
	KEY_FN_F1                    = 0x1d2
	KEY_FN_F2                    = 0x1d3
	KEY_FN_F3                    = 0x1d4
	KEY_FN_F4                    = 0x1d5
	KEY_FN_F5                    = 0x1d6
	KEY_FN_F6                    = 0x1d7
	KEY_FN_F7                    = 0x1d8
	KEY_FN_F8                    = 0x1d9
	KEY_FN_F9                    = 0x1da
	KEY_FN_F10                   = 0x1db
	KEY_FN_F11                   = 0x1dc
	KEY_FN_F12                   = 0x1dd
	KEY_FN_1                     = 0x1de
	KEY_FN_2                     = 0x1df
	KEY_FN_D                     = 0x1e0
	KEY_FN_E                     = 0x1e1
	KEY_FN_F                     = 0x1e2
	KEY_FN_S                     = 0x1e3
	KEY_FN_B                     = 0x1e4
	KEY_BRL_DOT1                 = 0x1f1
	KEY_BRL_DOT2                 = 0x1f2
	KEY_BRL_DOT3                 = 0x1f3
	KEY_BRL_DOT4                 = 0x1f4
	KEY_BRL_DOT5                 = 0x1f5
	KEY_BRL_DOT6                 = 0x1f6
	KEY_BRL_DOT7                 = 0x1f7
	KEY_BRL_DOT8                 = 0x1f8
	KEY_BRL_DOT9                 = 0x1f9
	KEY_BRL_DOT10                = 0x1fa
	KEY_NUMERIC_0                = 0x200
	KEY_NUMERIC_1                = 0x201
	KEY_NUMERIC_2                = 0x202
	KEY_NUMERIC_3                = 0x203
	KEY_NUMERIC_4                = 0x204
	KEY_NUMERIC_5                = 0x205
	KEY_NUMERIC_6                = 0x206
	KEY_NUMERIC_7                = 0x207
	KEY_NUMERIC_8                = 0x208
	KEY_NUMERIC_9                = 0x209
	KEY_NUMERIC_STAR             = 0x20a
	KEY_NUMERIC_POUND            = 0x20b
	KEY_NUMERIC_A                = 0x20c
	KEY_NUMERIC_B                = 0x20d
	KEY_NUMERIC_C                = 0x20e
	KEY_NUMERIC_D                = 0x20f
	KEY_CAMERA_FOCUS             = 0x210
	KEY_WPS_BUTTON               = 0x211
	KEY_TOUCHPAD_TOGGLE          = 0x212
	KEY_TOUCHPAD_ON              = 0x213
	KEY_TOUCHPAD_OFF             = 0x214
	KEY_CAMERA_ZOOMIN            = 0x215
	KEY_CAMERA_ZOOMOUT           = 0x216
	KEY_CAMERA_UP                = 0x217
	KEY_CAMERA_DOWN              = 0x218
	KEY_CAMERA_LEFT              = 0x219
	KEY_CAMERA_RIGHT             = 0x21a
	KEY_ATTENDANT_ON             = 0x21b
	KEY_ATTENDANT_OFF            = 0x21c
	KEY_ATTENDANT_TOGGLE         = 0x21d
	KEY_LIGHTS_TOGGLE            = 0x21e
	BTN_DPAD_UP                  = 0x220
	BTN_DPAD_DOWN                = 0x221
	BTN_DPAD_LEFT                = 0x222
	BTN_DPAD_RIGHT               = 0x223
	KEY_ALS_TOGGLE               = 0x230
	KEY_BUTTONCONFIG             = 0x240
	KEY_TASKMANAGER              = 0x241
	KEY_JOURNAL                  = 0x242
	KEY_CONTROLPANEL             = 0x243
	KEY_APPSELECT                = 0x244
	KEY_SCREENSAVER              = 0x245
	KEY_VOICECOMMAND             = 0x246
	KEY_BRIGHTNESS_MIN           = 0x250
	KEY_BRIGHTNESS_MAX           = 0x251
	KEY_KBDINPUTASSIST_PREV      = 0x260
	KEY_KBDINPUTASSIST_NEXT      = 0x261
	KEY_KBDINPUTASSIST_PREVGROUP = 0x262
	KEY_KBDINPUTASSIST_NEXTGROUP = 0x263
	KEY_KBDINPUTASSIST_ACCEPT    = 0x264
	KEY_KBDINPUTASSIST_CANCEL    = 0x265
	KEY_RIGHT_UP                 = 0x266
	KEY_RIGHT_DOWN               = 0x267
	KEY_LEFT_UP                  = 0x268
	KEY_LEFT_DOWN                = 0x269
	KEY_ROOT_MENU                = 0x26a
	KEY_MEDIA_TOP_MENU           = 0x26b
	KEY_NUMERIC_11               = 0x26c
	KEY_NUMERIC_12               = 0x26d
	KEY_AUDIO_DESC               = 0x26e
	KEY_3D_MODE                  = 0x26f
	KEY_NEXT_FAVORITE            = 0x270
	KEY_STOP_RECORD              = 0x271
	KEY_PAUSE_RECORD             = 0x272
	KEY_VOD                      = 0x273
	KEY_UNMUTE                   = 0x274
	KEY_FASTREVERSE              = 0x275
	KEY_SLOWREVERSE              = 0x276
	KEY_DATA                     = 0x275
	BTN_TRIGGER_HAPPY            = 0x2c0
	BTN_TRIGGER_HAPPY1           = 0x2c0
	BTN_TRIGGER_HAPPY2           = 0x2c1
	BTN_TRIGGER_HAPPY3           = 0x2c2
	BTN_TRIGGER_HAPPY4           = 0x2c3
	BTN_TRIGGER_HAPPY5           = 0x2c4
	BTN_TRIGGER_HAPPY6           = 0x2c5
	BTN_TRIGGER_HAPPY7           = 0x2c6
	BTN_TRIGGER_HAPPY8           = 0x2c7
	BTN_TRIGGER_HAPPY9           = 0x2c8
	BTN_TRIGGER_HAPPY10          = 0x2c9
	BTN_TRIGGER_HAPPY11          = 0x2ca
	BTN_TRIGGER_HAPPY12          = 0x2cb
	BTN_TRIGGER_HAPPY13          = 0x2cc
	BTN_TRIGGER_HAPPY14          = 0x2cd
	BTN_TRIGGER_HAPPY15          = 0x2ce
	BTN_TRIGGER_HAPPY16          = 0x2cf
	BTN_TRIGGER_HAPPY17          = 0x2d0
	BTN_TRIGGER_HAPPY18          = 0x2d1
	BTN_TRIGGER_HAPPY19          = 0x2d2
	BTN_TRIGGER_HAPPY20          = 0x2d3
	BTN_TRIGGER_HAPPY21          = 0x2d4
	BTN_TRIGGER_HAPPY22          = 0x2d5
	BTN_TRIGGER_HAPPY23          = 0x2d6
	BTN_TRIGGER_HAPPY24          = 0x2d7
	BTN_TRIGGER_HAPPY25          = 0x2d8
	BTN_TRIGGER_HAPPY26          = 0x2d9
	BTN_TRIGGER_HAPPY27          = 0x2da
	BTN_TRIGGER_HAPPY28          = 0x2db
	BTN_TRIGGER_HAPPY29          = 0x2dc
	BTN_TRIGGER_HAPPY30          = 0x2dd
	BTN_TRIGGER_HAPPY31          = 0x2de
	BTN_TRIGGER_HAPPY32          = 0x2df
	BTN_TRIGGER_HAPPY33          = 0x2e0
	BTN_TRIGGER_HAPPY34          = 0x2e1
	BTN_TRIGGER_HAPPY35          = 0x2e2
	BTN_TRIGGER_HAPPY36          = 0x2e3
	BTN_TRIGGER_HAPPY37          = 0x2e4
	BTN_TRIGGER_HAPPY38          = 0x2e5
	BTN_TRIGGER_HAPPY39          = 0x2e6
	BTN_TRIGGER_HAPPY40          = 0x2e7
	KEY_MIN_INTERESTING          = KEY_MUTE
	KEY_MAX                      = 0x2ff
	REL_X                        = 0x00
	REL_Y                        = 0x01
	REL_Z                        = 0x02
	REL_RX                       = 0x03
	REL_RY                       = 0x04
	REL_RZ                       = 0x05
	REL_HWHEEL                   = 0x06
	REL_DIAL                     = 0x07
	REL_WHEEL                    = 0x08
	REL_MISC                     = 0x09
	REL_MAX                      = 0x0f
	ABS_X                        = 0x00
	ABS_Y                        = 0x01
	ABS_Z                        = 0x02
	ABS_RX                       = 0x03
	ABS_RY                       = 0x04
	ABS_RZ                       = 0x05
	ABS_THROTTLE                 = 0x06
	ABS_RUDDER                   = 0x07
	ABS_WHEEL                    = 0x08
	ABS_GAS                      = 0x09
	ABS_BRAKE                    = 0x0a
	ABS_HAT0X                    = 0x10
	ABS_HAT0Y                    = 0x11
	ABS_HAT1X                    = 0x12
	ABS_HAT1Y                    = 0x13
	ABS_HAT2X                    = 0x14
	ABS_HAT2Y                    = 0x15
	ABS_HAT3X                    = 0x16
	ABS_HAT3Y                    = 0x17
	ABS_PRESSURE                 = 0x18
	ABS_DISTANCE                 = 0x19
	ABS_TILT_X                   = 0x1a
	ABS_TILT_Y                   = 0x1b
	ABS_TOOL_WIDTH               = 0x1c
	ABS_VOLUME                   = 0x20
	ABS_MISC                     = 0x28
	ABS_MT_SLOT                  = 0x2f
	ABS_MT_TOUCH_MAJOR           = 0x30
	ABS_MT_TOUCH_MINOR           = 0x31
	ABS_MT_WIDTH_MAJOR           = 0x32
	ABS_MT_WIDTH_MINOR           = 0x33
	ABS_MT_ORIENTATION           = 0x34
	ABS_MT_POSITION_X            = 0x35
	ABS_MT_POSITION_Y            = 0x36
	ABS_MT_TOOL_TYPE             = 0x37
	ABS_MT_BLOB_ID               = 0x38
	ABS_MT_TRACKING_ID           = 0x39
	ABS_MT_PRESSURE              = 0x3a
	ABS_MT_DISTANCE              = 0x3b
	ABS_MT_TOOL_X                = 0x3c
	ABS_MT_TOOL_Y                = 0x3d
	ABS_MAX                      = 0x3f
	SW_LID                       = 0x00
	SW_TABLET_MODE               = 0x01
	SW_HEADPHONE_INSERT          = 0x02
	SW_RFKILL_ALL                = 0x03
	SW_RADIO                     = SW_RFKILL_ALL
	SW_MICROPHONE_INSERT         = 0x04
	SW_DOCK                      = 0x05
	SW_LINEOUT_INSERT            = 0x06
	SW_JACK_PHYSICAL_INSERT      = 0x07
	SW_VIDEOOUT_INSERT           = 0x08
	SW_CAMERA_LENS_COVER         = 0x09
	SW_KEYPAD_SLIDE              = 0x0a
	SW_FRONT_PROXIMITY           = 0x0b
	SW_ROTATE_LOCK               = 0x0c
	SW_LINEIN_INSERT             = 0x0d
	SW_MUTE_DEVICE               = 0x0e
	SW_PEN_INSERTED              = 0x0f
	SW_MAX                       = 0x0f
	MSC_SERIAL                   = 0x00
	MSC_PULSELED                 = 0x01
	MSC_GESTURE                  = 0x02
	MSC_RAW                      = 0x03
	MSC_SCAN                     = 0x04
	MSC_TIMESTAMP                = 0x05
	MSC_MAX                      = 0x07
	LED_NUML                     = 0x00
	LED_CAPSL                    = 0x01
	LED_SCROLLL                  = 0x02
	LED_COMPOSE                  = 0x03
	LED_KANA                     = 0x04
	LED_SLEEP                    = 0x05
	LED_SUSPEND                  = 0x06
	LED_MUTE                     = 0x07
	LED_MISC                     = 0x08
	LED_MAIL                     = 0x09
	LED_CHARGING                 = 0x0a
	LED_MAX                      = 0x0f
	REP_DELAY                    = 0x00
	REP_PERIOD                   = 0x01
	REP_MAX                      = 0x01
	SND_CLICK                    = 0x00
	SND_BELL                     = 0x01
	SND_TONE                     = 0x02
	SND_MAX                      = 0x07
)

var ecodes = map[string]int{
	"EV_VERSION":                   EV_VERSION,
	"ID_BUS":                       ID_BUS,
	"ID_VENDOR":                    ID_VENDOR,
	"ID_PRODUCT":                   ID_PRODUCT,
	"ID_VERSION":                   ID_VERSION,
	"BUS_PCI":                      BUS_PCI,
	"BUS_ISAPNP":                   BUS_ISAPNP,
	"BUS_USB":                      BUS_USB,
	"BUS_HIL":                      BUS_HIL,
	"BUS_BLUETOOTH":                BUS_BLUETOOTH,
	"BUS_VIRTUAL":                  BUS_VIRTUAL,
	"BUS_ISA":                      BUS_ISA,
	"BUS_I8042":                    BUS_I8042,
	"BUS_XTKBD":                    BUS_XTKBD,
	"BUS_RS232":                    BUS_RS232,
	"BUS_GAMEPORT":                 BUS_GAMEPORT,
	"BUS_PARPORT":                  BUS_PARPORT,
	"BUS_AMIGA":                    BUS_AMIGA,
	"BUS_ADB":                      BUS_ADB,
	"BUS_I2C":                      BUS_I2C,
	"BUS_HOST":                     BUS_HOST,
	"BUS_GSC":                      BUS_GSC,
	"BUS_ATARI":                    BUS_ATARI,
	"BUS_SPI":                      BUS_SPI,
	"BUS_RMI":                      BUS_RMI,
	"BUS_CEC":                      BUS_CEC,
	"FF_STATUS_STOPPED":            FF_STATUS_STOPPED,
	"FF_STATUS_PLAYING":            FF_STATUS_PLAYING,
	"FF_STATUS_MAX":                FF_STATUS_MAX,
	"FF_RUMBLE":                    FF_RUMBLE,
	"FF_PERIODIC":                  FF_PERIODIC,
	"FF_CONSTANT":                  FF_CONSTANT,
	"FF_SPRING":                    FF_SPRING,
	"FF_FRICTION":                  FF_FRICTION,
	"FF_DAMPER":                    FF_DAMPER,
	"FF_INERTIA":                   FF_INERTIA,
	"FF_RAMP":                      FF_RAMP,
	"FF_EFFECT_MIN":                FF_EFFECT_MIN,
	"FF_EFFECT_MAX":                FF_EFFECT_MAX,
	"FF_SQUARE":                    FF_SQUARE,
	"FF_TRIANGLE":                  FF_TRIANGLE,
	"FF_SINE":                      FF_SINE,
	"FF_SAW_UP":                    FF_SAW_UP,
	"FF_SAW_DOWN":                  FF_SAW_DOWN,
	"FF_CUSTOM":                    FF_CUSTOM,
	"FF_WAVEFORM_MIN":              FF_WAVEFORM_MIN,
	"FF_WAVEFORM_MAX":              FF_WAVEFORM_MAX,
	"FF_GAIN":                      FF_GAIN,
	"FF_AUTOCENTER":                FF_AUTOCENTER,
	"FF_MAX_EFFECTS":               FF_MAX_EFFECTS,
	"FF_MAX":                       FF_MAX,
	"EV_SYN":                       EV_SYN,
	"EV_KEY":                       EV_KEY,
	"EV_REL":                       EV_REL,
	"EV_ABS":                       EV_ABS,
	"EV_MSC":                       EV_MSC,
	"EV_SW":                        EV_SW,
	"EV_LED":                       EV_LED,
	"EV_SND":                       EV_SND,
	"EV_REP":                       EV_REP,
	"EV_FF":                        EV_FF,
	"EV_PWR":                       EV_PWR,
	"EV_FF_STATUS":                 EV_FF_STATUS,
	"EV_MAX":                       EV_MAX,
	"SYN_REPORT":                   SYN_REPORT,
	"SYN_CONFIG":                   SYN_CONFIG,
	"SYN_MT_REPORT":                SYN_MT_REPORT,
	"SYN_DROPPED":                  SYN_DROPPED,
	"SYN_MAX":                      SYN_MAX,
	"KEY_RESERVED":                 KEY_RESERVED,
	"KEY_ESC":                      KEY_ESC,
	"KEY_1":                        KEY_1,
	"KEY_2":                        KEY_2,
	"KEY_3":                        KEY_3,
	"KEY_4":                        KEY_4,
	"KEY_5":                        KEY_5,
	"KEY_6":                        KEY_6,
	"KEY_7":                        KEY_7,
	"KEY_8":                        KEY_8,
	"KEY_9":                        KEY_9,
	"KEY_0":                        KEY_0,
	"KEY_MINUS":                    KEY_MINUS,
	"KEY_EQUAL":                    KEY_EQUAL,
	"KEY_BACKSPACE":                KEY_BACKSPACE,
	"KEY_TAB":                      KEY_TAB,
	"KEY_Q":                        KEY_Q,
	"KEY_W":                        KEY_W,
	"KEY_E":                        KEY_E,
	"KEY_R":                        KEY_R,
	"KEY_T":                        KEY_T,
	"KEY_Y":                        KEY_Y,
	"KEY_U":                        KEY_U,
	"KEY_I":                        KEY_I,
	"KEY_O":                        KEY_O,
	"KEY_P":                        KEY_P,
	"KEY_LEFTBRACE":                KEY_LEFTBRACE,
	"KEY_RIGHTBRACE":               KEY_RIGHTBRACE,
	"KEY_ENTER":                    KEY_ENTER,
	"KEY_LEFTCTRL":                 KEY_LEFTCTRL,
	"KEY_A":                        KEY_A,
	"KEY_S":                        KEY_S,
	"KEY_D":                        KEY_D,
	"KEY_F":                        KEY_F,
	"KEY_G":                        KEY_G,
	"KEY_H":                        KEY_H,
	"KEY_J":                        KEY_J,
	"KEY_K":                        KEY_K,
	"KEY_L":                        KEY_L,
	"KEY_SEMICOLON":                KEY_SEMICOLON,
	"KEY_APOSTROPHE":               KEY_APOSTROPHE,
	"KEY_GRAVE":                    KEY_GRAVE,
	"KEY_LEFTSHIFT":                KEY_LEFTSHIFT,
	"KEY_BACKSLASH":                KEY_BACKSLASH,
	"KEY_Z":                        KEY_Z,
	"KEY_X":                        KEY_X,
	"KEY_C":                        KEY_C,
	"KEY_V":                        KEY_V,
	"KEY_B":                        KEY_B,
	"KEY_N":                        KEY_N,
	"KEY_M":                        KEY_M,
	"KEY_COMMA":                    KEY_COMMA,
	"KEY_DOT":                      KEY_DOT,
	"KEY_SLASH":                    KEY_SLASH,
	"KEY_RIGHTSHIFT":               KEY_RIGHTSHIFT,
	"KEY_KPASTERISK":               KEY_KPASTERISK,
	"KEY_LEFTALT":                  KEY_LEFTALT,
	"KEY_SPACE":                    KEY_SPACE,
	"KEY_CAPSLOCK":                 KEY_CAPSLOCK,
	"KEY_F1":                       KEY_F1,
	"KEY_F2":                       KEY_F2,
	"KEY_F3":                       KEY_F3,
	"KEY_F4":                       KEY_F4,
	"KEY_F5":                       KEY_F5,
	"KEY_F6":                       KEY_F6,
	"KEY_F7":                       KEY_F7,
	"KEY_F8":                       KEY_F8,
	"KEY_F9":                       KEY_F9,
	"KEY_F10":                      KEY_F10,
	"KEY_NUMLOCK":                  KEY_NUMLOCK,
	"KEY_SCROLLLOCK":               KEY_SCROLLLOCK,
	"KEY_KP7":                      KEY_KP7,
	"KEY_KP8":                      KEY_KP8,
	"KEY_KP9":                      KEY_KP9,
	"KEY_KPMINUS":                  KEY_KPMINUS,
	"KEY_KP4":                      KEY_KP4,
	"KEY_KP5":                      KEY_KP5,
	"KEY_KP6":                      KEY_KP6,
	"KEY_KPPLUS":                   KEY_KPPLUS,
	"KEY_KP1":                      KEY_KP1,
	"KEY_KP2":                      KEY_KP2,
	"KEY_KP3":                      KEY_KP3,
	"KEY_KP0":                      KEY_KP0,
	"KEY_KPDOT":                    KEY_KPDOT,
	"KEY_ZENKAKUHANKAKU":           KEY_ZENKAKUHANKAKU,
	"KEY_102ND":                    KEY_102ND,
	"KEY_F11":                      KEY_F11,
	"KEY_F12":                      KEY_F12,
	"KEY_RO":                       KEY_RO,
	"KEY_KATAKANA":                 KEY_KATAKANA,
	"KEY_HIRAGANA":                 KEY_HIRAGANA,
	"KEY_HENKAN":                   KEY_HENKAN,
	"KEY_KATAKANAHIRAGANA":         KEY_KATAKANAHIRAGANA,
	"KEY_MUHENKAN":                 KEY_MUHENKAN,
	"KEY_KPJPCOMMA":                KEY_KPJPCOMMA,
	"KEY_KPENTER":                  KEY_KPENTER,
	"KEY_RIGHTCTRL":                KEY_RIGHTCTRL,
	"KEY_KPSLASH":                  KEY_KPSLASH,
	"KEY_SYSRQ":                    KEY_SYSRQ,
	"KEY_RIGHTALT":                 KEY_RIGHTALT,
	"KEY_LINEFEED":                 KEY_LINEFEED,
	"KEY_HOME":                     KEY_HOME,
	"KEY_UP":                       KEY_UP,
	"KEY_PAGEUP":                   KEY_PAGEUP,
	"KEY_LEFT":                     KEY_LEFT,
	"KEY_RIGHT":                    KEY_RIGHT,
	"KEY_END":                      KEY_END,
	"KEY_DOWN":                     KEY_DOWN,
	"KEY_PAGEDOWN":                 KEY_PAGEDOWN,
	"KEY_INSERT":                   KEY_INSERT,
	"KEY_DELETE":                   KEY_DELETE,
	"KEY_MACRO":                    KEY_MACRO,
	"KEY_MUTE":                     KEY_MUTE,
	"KEY_VOLUMEDOWN":               KEY_VOLUMEDOWN,
	"KEY_VOLUMEUP":                 KEY_VOLUMEUP,
	"KEY_POWER":                    KEY_POWER,
	"KEY_KPEQUAL":                  KEY_KPEQUAL,
	"KEY_KPPLUSMINUS":              KEY_KPPLUSMINUS,
	"KEY_PAUSE":                    KEY_PAUSE,
	"KEY_SCALE":                    KEY_SCALE,
	"KEY_KPCOMMA":                  KEY_KPCOMMA,
	"KEY_HANGEUL":                  KEY_HANGEUL,
	"KEY_HANGUEL":                  KEY_HANGUEL,
	"KEY_HANJA":                    KEY_HANJA,
	"KEY_YEN":                      KEY_YEN,
	"KEY_LEFTMETA":                 KEY_LEFTMETA,
	"KEY_RIGHTMETA":                KEY_RIGHTMETA,
	"KEY_COMPOSE":                  KEY_COMPOSE,
	"KEY_STOP":                     KEY_STOP,
	"KEY_AGAIN":                    KEY_AGAIN,
	"KEY_PROPS":                    KEY_PROPS,
	"KEY_UNDO":                     KEY_UNDO,
	"KEY_FRONT":                    KEY_FRONT,
	"KEY_COPY":                     KEY_COPY,
	"KEY_OPEN":                     KEY_OPEN,
	"KEY_PASTE":                    KEY_PASTE,
	"KEY_FIND":                     KEY_FIND,
	"KEY_CUT":                      KEY_CUT,
	"KEY_HELP":                     KEY_HELP,
	"KEY_MENU":                     KEY_MENU,
	"KEY_CALC":                     KEY_CALC,
	"KEY_SETUP":                    KEY_SETUP,
	"KEY_SLEEP":                    KEY_SLEEP,
	"KEY_WAKEUP":                   KEY_WAKEUP,
	"KEY_FILE":                     KEY_FILE,
	"KEY_SENDFILE":                 KEY_SENDFILE,
	"KEY_DELETEFILE":               KEY_DELETEFILE,
	"KEY_XFER":                     KEY_XFER,
	"KEY_PROG1":                    KEY_PROG1,
	"KEY_PROG2":                    KEY_PROG2,
	"KEY_WWW":                      KEY_WWW,
	"KEY_MSDOS":                    KEY_MSDOS,
	"KEY_COFFEE":                   KEY_COFFEE,
	"KEY_SCREENLOCK":               KEY_SCREENLOCK,
	"KEY_ROTATE_DISPLAY":           KEY_ROTATE_DISPLAY,
	"KEY_DIRECTION":                KEY_DIRECTION,
	"KEY_CYCLEWINDOWS":             KEY_CYCLEWINDOWS,
	"KEY_MAIL":                     KEY_MAIL,
	"KEY_BOOKMARKS":                KEY_BOOKMARKS,
	"KEY_COMPUTER":                 KEY_COMPUTER,
	"KEY_BACK":                     KEY_BACK,
	"KEY_FORWARD":                  KEY_FORWARD,
	"KEY_CLOSECD":                  KEY_CLOSECD,
	"KEY_EJECTCD":                  KEY_EJECTCD,
	"KEY_EJECTCLOSECD":             KEY_EJECTCLOSECD,
	"KEY_NEXTSONG":                 KEY_NEXTSONG,
	"KEY_PLAYPAUSE":                KEY_PLAYPAUSE,
	"KEY_PREVIOUSSONG":             KEY_PREVIOUSSONG,
	"KEY_STOPCD":                   KEY_STOPCD,
	"KEY_RECORD":                   KEY_RECORD,
	"KEY_REWIND":                   KEY_REWIND,
	"KEY_PHONE":                    KEY_PHONE,
	"KEY_ISO":                      KEY_ISO,
	"KEY_CONFIG":                   KEY_CONFIG,
	"KEY_HOMEPAGE":                 KEY_HOMEPAGE,
	"KEY_REFRESH":                  KEY_REFRESH,
	"KEY_EXIT":                     KEY_EXIT,
	"KEY_MOVE":                     KEY_MOVE,
	"KEY_EDIT":                     KEY_EDIT,
	"KEY_SCROLLUP":                 KEY_SCROLLUP,
	"KEY_SCROLLDOWN":               KEY_SCROLLDOWN,
	"KEY_KPLEFTPAREN":              KEY_KPLEFTPAREN,
	"KEY_KPRIGHTPAREN":             KEY_KPRIGHTPAREN,
	"KEY_NEW":                      KEY_NEW,
	"KEY_REDO":                     KEY_REDO,
	"KEY_F13":                      KEY_F13,
	"KEY_F14":                      KEY_F14,
	"KEY_F15":                      KEY_F15,
	"KEY_F16":                      KEY_F16,
	"KEY_F17":                      KEY_F17,
	"KEY_F18":                      KEY_F18,
	"KEY_F19":                      KEY_F19,
	"KEY_F20":                      KEY_F20,
	"KEY_F21":                      KEY_F21,
	"KEY_F22":                      KEY_F22,
	"KEY_F23":                      KEY_F23,
	"KEY_F24":                      KEY_F24,
	"KEY_PLAYCD":                   KEY_PLAYCD,
	"KEY_PAUSECD":                  KEY_PAUSECD,
	"KEY_PROG3":                    KEY_PROG3,
	"KEY_PROG4":                    KEY_PROG4,
	"KEY_DASHBOARD":                KEY_DASHBOARD,
	"KEY_SUSPEND":                  KEY_SUSPEND,
	"KEY_CLOSE":                    KEY_CLOSE,
	"KEY_PLAY":                     KEY_PLAY,
	"KEY_FASTFORWARD":              KEY_FASTFORWARD,
	"KEY_BASSBOOST":                KEY_BASSBOOST,
	"KEY_PRINT":                    KEY_PRINT,
	"KEY_HP":                       KEY_HP,
	"KEY_CAMERA":                   KEY_CAMERA,
	"KEY_SOUND":                    KEY_SOUND,
	"KEY_QUESTION":                 KEY_QUESTION,
	"KEY_EMAIL":                    KEY_EMAIL,
	"KEY_CHAT":                     KEY_CHAT,
	"KEY_SEARCH":                   KEY_SEARCH,
	"KEY_CONNECT":                  KEY_CONNECT,
	"KEY_FINANCE":                  KEY_FINANCE,
	"KEY_SPORT":                    KEY_SPORT,
	"KEY_SHOP":                     KEY_SHOP,
	"KEY_ALTERASE":                 KEY_ALTERASE,
	"KEY_CANCEL":                   KEY_CANCEL,
	"KEY_BRIGHTNESSDOWN":           KEY_BRIGHTNESSDOWN,
	"KEY_BRIGHTNESSUP":             KEY_BRIGHTNESSUP,
	"KEY_MEDIA":                    KEY_MEDIA,
	"KEY_SWITCHVIDEOMODE":          KEY_SWITCHVIDEOMODE,
	"KEY_KBDILLUMTOGGLE":           KEY_KBDILLUMTOGGLE,
	"KEY_KBDILLUMDOWN":             KEY_KBDILLUMDOWN,
	"KEY_KBDILLUMUP":               KEY_KBDILLUMUP,
	"KEY_SEND":                     KEY_SEND,
	"KEY_REPLY":                    KEY_REPLY,
	"KEY_FORWARDMAIL":              KEY_FORWARDMAIL,
	"KEY_SAVE":                     KEY_SAVE,
	"KEY_DOCUMENTS":                KEY_DOCUMENTS,
	"KEY_BATTERY":                  KEY_BATTERY,
	"KEY_BLUETOOTH":                KEY_BLUETOOTH,
	"KEY_WLAN":                     KEY_WLAN,
	"KEY_UWB":                      KEY_UWB,
	"KEY_UNKNOWN":                  KEY_UNKNOWN,
	"KEY_VIDEO_NEXT":               KEY_VIDEO_NEXT,
	"KEY_VIDEO_PREV":               KEY_VIDEO_PREV,
	"KEY_BRIGHTNESS_CYCLE":         KEY_BRIGHTNESS_CYCLE,
	"KEY_BRIGHTNESS_AUTO":          KEY_BRIGHTNESS_AUTO,
	"KEY_BRIGHTNESS_ZERO":          KEY_BRIGHTNESS_ZERO,
	"KEY_DISPLAY_OFF":              KEY_DISPLAY_OFF,
	"KEY_WWAN":                     KEY_WWAN,
	"KEY_WIMAX":                    KEY_WIMAX,
	"KEY_RFKILL":                   KEY_RFKILL,
	"KEY_MICMUTE":                  KEY_MICMUTE,
	"BTN_MISC":                     BTN_MISC,
	"BTN_0":                        BTN_0,
	"BTN_1":                        BTN_1,
	"BTN_2":                        BTN_2,
	"BTN_3":                        BTN_3,
	"BTN_4":                        BTN_4,
	"BTN_5":                        BTN_5,
	"BTN_6":                        BTN_6,
	"BTN_7":                        BTN_7,
	"BTN_8":                        BTN_8,
	"BTN_9":                        BTN_9,
	"BTN_MOUSE":                    BTN_MOUSE,
	"BTN_LEFT":                     BTN_LEFT,
	"BTN_RIGHT":                    BTN_RIGHT,
	"BTN_MIDDLE":                   BTN_MIDDLE,
	"BTN_SIDE":                     BTN_SIDE,
	"BTN_EXTRA":                    BTN_EXTRA,
	"BTN_FORWARD":                  BTN_FORWARD,
	"BTN_BACK":                     BTN_BACK,
	"BTN_TASK":                     BTN_TASK,
	"BTN_JOYSTICK":                 BTN_JOYSTICK,
	"BTN_TRIGGER":                  BTN_TRIGGER,
	"BTN_THUMB":                    BTN_THUMB,
	"BTN_THUMB2":                   BTN_THUMB2,
	"BTN_TOP":                      BTN_TOP,
	"BTN_TOP2":                     BTN_TOP2,
	"BTN_PINKIE":                   BTN_PINKIE,
	"BTN_BASE":                     BTN_BASE,
	"BTN_BASE2":                    BTN_BASE2,
	"BTN_BASE3":                    BTN_BASE3,
	"BTN_BASE4":                    BTN_BASE4,
	"BTN_BASE5":                    BTN_BASE5,
	"BTN_BASE6":                    BTN_BASE6,
	"BTN_DEAD":                     BTN_DEAD,
	"BTN_GAMEPAD":                  BTN_GAMEPAD,
	"BTN_SOUTH":                    BTN_SOUTH,
	"BTN_A":                        BTN_A,
	"BTN_EAST":                     BTN_EAST,
	"BTN_B":                        BTN_B,
	"BTN_C":                        BTN_C,
	"BTN_NORTH":                    BTN_NORTH,
	"BTN_X":                        BTN_X,
	"BTN_WEST":                     BTN_WEST,
	"BTN_Y":                        BTN_Y,
	"BTN_Z":                        BTN_Z,
	"BTN_TL":                       BTN_TL,
	"BTN_TR":                       BTN_TR,
	"BTN_TL2":                      BTN_TL2,
	"BTN_TR2":                      BTN_TR2,
	"BTN_SELECT":                   BTN_SELECT,
	"BTN_START":                    BTN_START,
	"BTN_MODE":                     BTN_MODE,
	"BTN_THUMBL":                   BTN_THUMBL,
	"BTN_THUMBR":                   BTN_THUMBR,
	"BTN_DIGI":                     BTN_DIGI,
	"BTN_TOOL_PEN":                 BTN_TOOL_PEN,
	"BTN_TOOL_RUBBER":              BTN_TOOL_RUBBER,
	"BTN_TOOL_BRUSH":               BTN_TOOL_BRUSH,
	"BTN_TOOL_PENCIL":              BTN_TOOL_PENCIL,
	"BTN_TOOL_AIRBRUSH":            BTN_TOOL_AIRBRUSH,
	"BTN_TOOL_FINGER":              BTN_TOOL_FINGER,
	"BTN_TOOL_MOUSE":               BTN_TOOL_MOUSE,
	"BTN_TOOL_LENS":                BTN_TOOL_LENS,
	"BTN_TOOL_QUINTTAP":            BTN_TOOL_QUINTTAP,
	"BTN_TOUCH":                    BTN_TOUCH,
	"BTN_STYLUS":                   BTN_STYLUS,
	"BTN_STYLUS2":                  BTN_STYLUS2,
	"BTN_TOOL_DOUBLETAP":           BTN_TOOL_DOUBLETAP,
	"BTN_TOOL_TRIPLETAP":           BTN_TOOL_TRIPLETAP,
	"BTN_TOOL_QUADTAP":             BTN_TOOL_QUADTAP,
	"BTN_WHEEL":                    BTN_WHEEL,
	"BTN_GEAR_DOWN":                BTN_GEAR_DOWN,
	"BTN_GEAR_UP":                  BTN_GEAR_UP,
	"KEY_OK":                       KEY_OK,
	"KEY_SELECT":                   KEY_SELECT,
	"KEY_GOTO":                     KEY_GOTO,
	"KEY_CLEAR":                    KEY_CLEAR,
	"KEY_POWER2":                   KEY_POWER2,
	"KEY_OPTION":                   KEY_OPTION,
	"KEY_INFO":                     KEY_INFO,
	"KEY_TIME":                     KEY_TIME,
	"KEY_VENDOR":                   KEY_VENDOR,
	"KEY_ARCHIVE":                  KEY_ARCHIVE,
	"KEY_PROGRAM":                  KEY_PROGRAM,
	"KEY_CHANNEL":                  KEY_CHANNEL,
	"KEY_FAVORITES":                KEY_FAVORITES,
	"KEY_EPG":                      KEY_EPG,
	"KEY_PVR":                      KEY_PVR,
	"KEY_MHP":                      KEY_MHP,
	"KEY_LANGUAGE":                 KEY_LANGUAGE,
	"KEY_TITLE":                    KEY_TITLE,
	"KEY_SUBTITLE":                 KEY_SUBTITLE,
	"KEY_ANGLE":                    KEY_ANGLE,
	"KEY_ZOOM":                     KEY_ZOOM,
	"KEY_MODE":                     KEY_MODE,
	"KEY_KEYBOARD":                 KEY_KEYBOARD,
	"KEY_SCREEN":                   KEY_SCREEN,
	"KEY_PC":                       KEY_PC,
	"KEY_TV":                       KEY_TV,
	"KEY_TV2":                      KEY_TV2,
	"KEY_VCR":                      KEY_VCR,
	"KEY_VCR2":                     KEY_VCR2,
	"KEY_SAT":                      KEY_SAT,
	"KEY_SAT2":                     KEY_SAT2,
	"KEY_CD":                       KEY_CD,
	"KEY_TAPE":                     KEY_TAPE,
	"KEY_RADIO":                    KEY_RADIO,
	"KEY_TUNER":                    KEY_TUNER,
	"KEY_PLAYER":                   KEY_PLAYER,
	"KEY_TEXT":                     KEY_TEXT,
	"KEY_DVD":                      KEY_DVD,
	"KEY_AUX":                      KEY_AUX,
	"KEY_MP3":                      KEY_MP3,
	"KEY_AUDIO":                    KEY_AUDIO,
	"KEY_VIDEO":                    KEY_VIDEO,
	"KEY_DIRECTORY":                KEY_DIRECTORY,
	"KEY_LIST":                     KEY_LIST,
	"KEY_MEMO":                     KEY_MEMO,
	"KEY_CALENDAR":                 KEY_CALENDAR,
	"KEY_RED":                      KEY_RED,
	"KEY_GREEN":                    KEY_GREEN,
	"KEY_YELLOW":                   KEY_YELLOW,
	"KEY_BLUE":                     KEY_BLUE,
	"KEY_CHANNELUP":                KEY_CHANNELUP,
	"KEY_CHANNELDOWN":              KEY_CHANNELDOWN,
	"KEY_FIRST":                    KEY_FIRST,
	"KEY_LAST":                     KEY_LAST,
	"KEY_AB":                       KEY_AB,
	"KEY_NEXT":                     KEY_NEXT,
	"KEY_RESTART":                  KEY_RESTART,
	"KEY_SLOW":                     KEY_SLOW,
	"KEY_SHUFFLE":                  KEY_SHUFFLE,
	"KEY_BREAK":                    KEY_BREAK,
	"KEY_PREVIOUS":                 KEY_PREVIOUS,
	"KEY_DIGITS":                   KEY_DIGITS,
	"KEY_TEEN":                     KEY_TEEN,
	"KEY_TWEN":                     KEY_TWEN,
	"KEY_VIDEOPHONE":               KEY_VIDEOPHONE,
	"KEY_GAMES":                    KEY_GAMES,
	"KEY_ZOOMIN":                   KEY_ZOOMIN,
	"KEY_ZOOMOUT":                  KEY_ZOOMOUT,
	"KEY_ZOOMRESET":                KEY_ZOOMRESET,
	"KEY_WORDPROCESSOR":            KEY_WORDPROCESSOR,
	"KEY_EDITOR":                   KEY_EDITOR,
	"KEY_SPREADSHEET":              KEY_SPREADSHEET,
	"KEY_GRAPHICSEDITOR":           KEY_GRAPHICSEDITOR,
	"KEY_PRESENTATION":             KEY_PRESENTATION,
	"KEY_DATABASE":                 KEY_DATABASE,
	"KEY_NEWS":                     KEY_NEWS,
	"KEY_VOICEMAIL":                KEY_VOICEMAIL,
	"KEY_ADDRESSBOOK":              KEY_ADDRESSBOOK,
	"KEY_MESSENGER":                KEY_MESSENGER,
	"KEY_DISPLAYTOGGLE":            KEY_DISPLAYTOGGLE,
	"KEY_BRIGHTNESS_TOGGLE":        KEY_BRIGHTNESS_TOGGLE,
	"KEY_SPELLCHECK":               KEY_SPELLCHECK,
	"KEY_LOGOFF":                   KEY_LOGOFF,
	"KEY_DOLLAR":                   KEY_DOLLAR,
	"KEY_EURO":                     KEY_EURO,
	"KEY_FRAMEBACK":                KEY_FRAMEBACK,
	"KEY_FRAMEFORWARD":             KEY_FRAMEFORWARD,
	"KEY_CONTEXT_MENU":             KEY_CONTEXT_MENU,
	"KEY_MEDIA_REPEAT":             KEY_MEDIA_REPEAT,
	"KEY_10CHANNELSUP":             KEY_10CHANNELSUP,
	"KEY_10CHANNELSDOWN":           KEY_10CHANNELSDOWN,
	"KEY_IMAGES":                   KEY_IMAGES,
	"KEY_DEL_EOL":                  KEY_DEL_EOL,
	"KEY_DEL_EOS":                  KEY_DEL_EOS,
	"KEY_INS_LINE":                 KEY_INS_LINE,
	"KEY_DEL_LINE":                 KEY_DEL_LINE,
	"KEY_FN":                       KEY_FN,
	"KEY_FN_ESC":                   KEY_FN_ESC,
	"KEY_FN_F1":                    KEY_FN_F1,
	"KEY_FN_F2":                    KEY_FN_F2,
	"KEY_FN_F3":                    KEY_FN_F3,
	"KEY_FN_F4":                    KEY_FN_F4,
	"KEY_FN_F5":                    KEY_FN_F5,
	"KEY_FN_F6":                    KEY_FN_F6,
	"KEY_FN_F7":                    KEY_FN_F7,
	"KEY_FN_F8":                    KEY_FN_F8,
	"KEY_FN_F9":                    KEY_FN_F9,
	"KEY_FN_F10":                   KEY_FN_F10,
	"KEY_FN_F11":                   KEY_FN_F11,
	"KEY_FN_F12":                   KEY_FN_F12,
	"KEY_FN_1":                     KEY_FN_1,
	"KEY_FN_2":                     KEY_FN_2,
	"KEY_FN_D":                     KEY_FN_D,
	"KEY_FN_E":                     KEY_FN_E,
	"KEY_FN_F":                     KEY_FN_F,
	"KEY_FN_S":                     KEY_FN_S,
	"KEY_FN_B":                     KEY_FN_B,
	"KEY_BRL_DOT1":                 KEY_BRL_DOT1,
	"KEY_BRL_DOT2":                 KEY_BRL_DOT2,
	"KEY_BRL_DOT3":                 KEY_BRL_DOT3,
	"KEY_BRL_DOT4":                 KEY_BRL_DOT4,
	"KEY_BRL_DOT5":                 KEY_BRL_DOT5,
	"KEY_BRL_DOT6":                 KEY_BRL_DOT6,
	"KEY_BRL_DOT7":                 KEY_BRL_DOT7,
	"KEY_BRL_DOT8":                 KEY_BRL_DOT8,
	"KEY_BRL_DOT9":                 KEY_BRL_DOT9,
	"KEY_BRL_DOT10":                KEY_BRL_DOT10,
	"KEY_NUMERIC_0":                KEY_NUMERIC_0,
	"KEY_NUMERIC_1":                KEY_NUMERIC_1,
	"KEY_NUMERIC_2":                KEY_NUMERIC_2,
	"KEY_NUMERIC_3":                KEY_NUMERIC_3,
	"KEY_NUMERIC_4":                KEY_NUMERIC_4,
	"KEY_NUMERIC_5":                KEY_NUMERIC_5,
	"KEY_NUMERIC_6":                KEY_NUMERIC_6,
	"KEY_NUMERIC_7":                KEY_NUMERIC_7,
	"KEY_NUMERIC_8":                KEY_NUMERIC_8,
	"KEY_NUMERIC_9":                KEY_NUMERIC_9,
	"KEY_NUMERIC_STAR":             KEY_NUMERIC_STAR,
	"KEY_NUMERIC_POUND":            KEY_NUMERIC_POUND,
	"KEY_NUMERIC_A":                KEY_NUMERIC_A,
	"KEY_NUMERIC_B":                KEY_NUMERIC_B,
	"KEY_NUMERIC_C":                KEY_NUMERIC_C,
	"KEY_NUMERIC_D":                KEY_NUMERIC_D,
	"KEY_CAMERA_FOCUS":             KEY_CAMERA_FOCUS,
	"KEY_WPS_BUTTON":               KEY_WPS_BUTTON,
	"KEY_TOUCHPAD_TOGGLE":          KEY_TOUCHPAD_TOGGLE,
	"KEY_TOUCHPAD_ON":              KEY_TOUCHPAD_ON,
	"KEY_TOUCHPAD_OFF":             KEY_TOUCHPAD_OFF,
	"KEY_CAMERA_ZOOMIN":            KEY_CAMERA_ZOOMIN,
	"KEY_CAMERA_ZOOMOUT":           KEY_CAMERA_ZOOMOUT,
	"KEY_CAMERA_UP":                KEY_CAMERA_UP,
	"KEY_CAMERA_DOWN":              KEY_CAMERA_DOWN,
	"KEY_CAMERA_LEFT":              KEY_CAMERA_LEFT,
	"KEY_CAMERA_RIGHT":             KEY_CAMERA_RIGHT,
	"KEY_ATTENDANT_ON":             KEY_ATTENDANT_ON,
	"KEY_ATTENDANT_OFF":            KEY_ATTENDANT_OFF,
	"KEY_ATTENDANT_TOGGLE":         KEY_ATTENDANT_TOGGLE,
	"KEY_LIGHTS_TOGGLE":            KEY_LIGHTS_TOGGLE,
	"BTN_DPAD_UP":                  BTN_DPAD_UP,
	"BTN_DPAD_DOWN":                BTN_DPAD_DOWN,
	"BTN_DPAD_LEFT":                BTN_DPAD_LEFT,
	"BTN_DPAD_RIGHT":               BTN_DPAD_RIGHT,
	"KEY_ALS_TOGGLE":               KEY_ALS_TOGGLE,
	"KEY_BUTTONCONFIG":             KEY_BUTTONCONFIG,
	"KEY_TASKMANAGER":              KEY_TASKMANAGER,
	"KEY_JOURNAL":                  KEY_JOURNAL,
	"KEY_CONTROLPANEL":             KEY_CONTROLPANEL,
	"KEY_APPSELECT":                KEY_APPSELECT,
	"KEY_SCREENSAVER":              KEY_SCREENSAVER,
	"KEY_VOICECOMMAND":             KEY_VOICECOMMAND,
	"KEY_BRIGHTNESS_MIN":           KEY_BRIGHTNESS_MIN,
	"KEY_BRIGHTNESS_MAX":           KEY_BRIGHTNESS_MAX,
	"KEY_KBDINPUTASSIST_PREV":      KEY_KBDINPUTASSIST_PREV,
	"KEY_KBDINPUTASSIST_NEXT":      KEY_KBDINPUTASSIST_NEXT,
	"KEY_KBDINPUTASSIST_PREVGROUP": KEY_KBDINPUTASSIST_PREVGROUP,
	"KEY_KBDINPUTASSIST_NEXTGROUP": KEY_KBDINPUTASSIST_NEXTGROUP,
	"KEY_KBDINPUTASSIST_ACCEPT":    KEY_KBDINPUTASSIST_ACCEPT,
	"KEY_KBDINPUTASSIST_CANCEL":    KEY_KBDINPUTASSIST_CANCEL,
	"KEY_RIGHT_UP":                 KEY_RIGHT_UP,
	"KEY_RIGHT_DOWN":               KEY_RIGHT_DOWN,
	"KEY_LEFT_UP":                  KEY_LEFT_UP,
	"KEY_LEFT_DOWN":                KEY_LEFT_DOWN,
	"KEY_ROOT_MENU":                KEY_ROOT_MENU,
	"KEY_MEDIA_TOP_MENU":           KEY_MEDIA_TOP_MENU,
	"KEY_NUMERIC_11":               KEY_NUMERIC_11,
	"KEY_NUMERIC_12":               KEY_NUMERIC_12,
	"KEY_AUDIO_DESC":               KEY_AUDIO_DESC,
	"KEY_3D_MODE":                  KEY_3D_MODE,
	"KEY_NEXT_FAVORITE":            KEY_NEXT_FAVORITE,
	"KEY_STOP_RECORD":              KEY_STOP_RECORD,
	"KEY_PAUSE_RECORD":             KEY_PAUSE_RECORD,
	"KEY_VOD":                      KEY_VOD,
	"KEY_UNMUTE":                   KEY_UNMUTE,
	"KEY_FASTREVERSE":              KEY_FASTREVERSE,
	"KEY_SLOWREVERSE":              KEY_SLOWREVERSE,
	"KEY_DATA":                     KEY_DATA,
	"BTN_TRIGGER_HAPPY":            BTN_TRIGGER_HAPPY,
	"BTN_TRIGGER_HAPPY1":           BTN_TRIGGER_HAPPY1,
	"BTN_TRIGGER_HAPPY2":           BTN_TRIGGER_HAPPY2,
	"BTN_TRIGGER_HAPPY3":           BTN_TRIGGER_HAPPY3,
	"BTN_TRIGGER_HAPPY4":           BTN_TRIGGER_HAPPY4,
	"BTN_TRIGGER_HAPPY5":           BTN_TRIGGER_HAPPY5,
	"BTN_TRIGGER_HAPPY6":           BTN_TRIGGER_HAPPY6,
	"BTN_TRIGGER_HAPPY7":           BTN_TRIGGER_HAPPY7,
	"BTN_TRIGGER_HAPPY8":           BTN_TRIGGER_HAPPY8,
	"BTN_TRIGGER_HAPPY9":           BTN_TRIGGER_HAPPY9,
	"BTN_TRIGGER_HAPPY10":          BTN_TRIGGER_HAPPY10,
	"BTN_TRIGGER_HAPPY11":          BTN_TRIGGER_HAPPY11,
	"BTN_TRIGGER_HAPPY12":          BTN_TRIGGER_HAPPY12,
	"BTN_TRIGGER_HAPPY13":          BTN_TRIGGER_HAPPY13,
	"BTN_TRIGGER_HAPPY14":          BTN_TRIGGER_HAPPY14,
	"BTN_TRIGGER_HAPPY15":          BTN_TRIGGER_HAPPY15,
	"BTN_TRIGGER_HAPPY16":          BTN_TRIGGER_HAPPY16,
	"BTN_TRIGGER_HAPPY17":          BTN_TRIGGER_HAPPY17,
	"BTN_TRIGGER_HAPPY18":          BTN_TRIGGER_HAPPY18,
	"BTN_TRIGGER_HAPPY19":          BTN_TRIGGER_HAPPY19,
	"BTN_TRIGGER_HAPPY20":          BTN_TRIGGER_HAPPY20,
	"BTN_TRIGGER_HAPPY21":          BTN_TRIGGER_HAPPY21,
	"BTN_TRIGGER_HAPPY22":          BTN_TRIGGER_HAPPY22,
	"BTN_TRIGGER_HAPPY23":          BTN_TRIGGER_HAPPY23,
	"BTN_TRIGGER_HAPPY24":          BTN_TRIGGER_HAPPY24,
	"BTN_TRIGGER_HAPPY25":          BTN_TRIGGER_HAPPY25,
	"BTN_TRIGGER_HAPPY26":          BTN_TRIGGER_HAPPY26,
	"BTN_TRIGGER_HAPPY27":          BTN_TRIGGER_HAPPY27,
	"BTN_TRIGGER_HAPPY28":          BTN_TRIGGER_HAPPY28,
	"BTN_TRIGGER_HAPPY29":          BTN_TRIGGER_HAPPY29,
	"BTN_TRIGGER_HAPPY30":          BTN_TRIGGER_HAPPY30,
	"BTN_TRIGGER_HAPPY31":          BTN_TRIGGER_HAPPY31,
	"BTN_TRIGGER_HAPPY32":          BTN_TRIGGER_HAPPY32,
	"BTN_TRIGGER_HAPPY33":          BTN_TRIGGER_HAPPY33,
	"BTN_TRIGGER_HAPPY34":          BTN_TRIGGER_HAPPY34,
	"BTN_TRIGGER_HAPPY35":          BTN_TRIGGER_HAPPY35,
	"BTN_TRIGGER_HAPPY36":          BTN_TRIGGER_HAPPY36,
	"BTN_TRIGGER_HAPPY37":          BTN_TRIGGER_HAPPY37,
	"BTN_TRIGGER_HAPPY38":          BTN_TRIGGER_HAPPY38,
	"BTN_TRIGGER_HAPPY39":          BTN_TRIGGER_HAPPY39,
	"BTN_TRIGGER_HAPPY40":          BTN_TRIGGER_HAPPY40,
	"KEY_MIN_INTERESTING":          KEY_MIN_INTERESTING,
	"KEY_MAX":                      KEY_MAX,
	"REL_X":                        REL_X,
	"REL_Y":                        REL_Y,
	"REL_Z":                        REL_Z,
	"REL_RX":                       REL_RX,
	"REL_RY":                       REL_RY,
	"REL_RZ":                       REL_RZ,
	"REL_HWHEEL":                   REL_HWHEEL,
	"REL_DIAL":                     REL_DIAL,
	"REL_WHEEL":                    REL_WHEEL,
	"REL_MISC":                     REL_MISC,
	"REL_MAX":                      REL_MAX,
	"ABS_X":                        ABS_X,
	"ABS_Y":                        ABS_Y,
	"ABS_Z":                        ABS_Z,
	"ABS_RX":                       ABS_RX,
	"ABS_RY":                       ABS_RY,
	"ABS_RZ":                       ABS_RZ,
	"ABS_THROTTLE":                 ABS_THROTTLE,
	"ABS_RUDDER":                   ABS_RUDDER,
	"ABS_WHEEL":                    ABS_WHEEL,
	"ABS_GAS":                      ABS_GAS,
	"ABS_BRAKE":                    ABS_BRAKE,
	"ABS_HAT0X":                    ABS_HAT0X,
	"ABS_HAT0Y":                    ABS_HAT0Y,
	"ABS_HAT1X":                    ABS_HAT1X,
	"ABS_HAT1Y":                    ABS_HAT1Y,
	"ABS_HAT2X":                    ABS_HAT2X,
	"ABS_HAT2Y":                    ABS_HAT2Y,
	"ABS_HAT3X":                    ABS_HAT3X,
	"ABS_HAT3Y":                    ABS_HAT3Y,
	"ABS_PRESSURE":                 ABS_PRESSURE,
	"ABS_DISTANCE":                 ABS_DISTANCE,
	"ABS_TILT_X":                   ABS_TILT_X,
	"ABS_TILT_Y":                   ABS_TILT_Y,
	"ABS_TOOL_WIDTH":               ABS_TOOL_WIDTH,
	"ABS_VOLUME":                   ABS_VOLUME,
	"ABS_MISC":                     ABS_MISC,
	"ABS_MT_SLOT":                  ABS_MT_SLOT,
	"ABS_MT_TOUCH_MAJOR":           ABS_MT_TOUCH_MAJOR,
	"ABS_MT_TOUCH_MINOR":           ABS_MT_TOUCH_MINOR,
	"ABS_MT_WIDTH_MAJOR":           ABS_MT_WIDTH_MAJOR,
	"ABS_MT_WIDTH_MINOR":           ABS_MT_WIDTH_MINOR,
	"ABS_MT_ORIENTATION":           ABS_MT_ORIENTATION,
	"ABS_MT_POSITION_X":            ABS_MT_POSITION_X,
	"ABS_MT_POSITION_Y":            ABS_MT_POSITION_Y,
	"ABS_MT_TOOL_TYPE":             ABS_MT_TOOL_TYPE,
	"ABS_MT_BLOB_ID":               ABS_MT_BLOB_ID,
	"ABS_MT_TRACKING_ID":           ABS_MT_TRACKING_ID,
	"ABS_MT_PRESSURE":              ABS_MT_PRESSURE,
	"ABS_MT_DISTANCE":              ABS_MT_DISTANCE,
	"ABS_MT_TOOL_X":                ABS_MT_TOOL_X,
	"ABS_MT_TOOL_Y":                ABS_MT_TOOL_Y,
	"ABS_MAX":                      ABS_MAX,
	"SW_LID":                       SW_LID,
	"SW_TABLET_MODE":               SW_TABLET_MODE,
	"SW_HEADPHONE_INSERT":          SW_HEADPHONE_INSERT,
	"SW_RFKILL_ALL":                SW_RFKILL_ALL,
	"SW_RADIO":                     SW_RADIO,
	"SW_MICROPHONE_INSERT":         SW_MICROPHONE_INSERT,
	"SW_DOCK":                      SW_DOCK,
	"SW_LINEOUT_INSERT":            SW_LINEOUT_INSERT,
	"SW_JACK_PHYSICAL_INSERT":      SW_JACK_PHYSICAL_INSERT,
	"SW_VIDEOOUT_INSERT":           SW_VIDEOOUT_INSERT,
	"SW_CAMERA_LENS_COVER":         SW_CAMERA_LENS_COVER,
	"SW_KEYPAD_SLIDE":              SW_KEYPAD_SLIDE,
	"SW_FRONT_PROXIMITY":           SW_FRONT_PROXIMITY,
	"SW_ROTATE_LOCK":               SW_ROTATE_LOCK,
	"SW_LINEIN_INSERT":             SW_LINEIN_INSERT,
	"SW_MUTE_DEVICE":               SW_MUTE_DEVICE,
	"SW_PEN_INSERTED":              SW_PEN_INSERTED,
	"SW_MAX":                       SW_MAX,
	"MSC_SERIAL":                   MSC_SERIAL,
	"MSC_PULSELED":                 MSC_PULSELED,
	"MSC_GESTURE":                  MSC_GESTURE,
	"MSC_RAW":                      MSC_RAW,
	"MSC_SCAN":                     MSC_SCAN,
	"MSC_TIMESTAMP":                MSC_TIMESTAMP,
	"MSC_MAX":                      MSC_MAX,
	"LED_NUML":                     LED_NUML,
	"LED_CAPSL":                    LED_CAPSL,
	"LED_SCROLLL":                  LED_SCROLLL,
	"LED_COMPOSE":                  LED_COMPOSE,
	"LED_KANA":                     LED_KANA,
	"LED_SLEEP":                    LED_SLEEP,
	"LED_SUSPEND":                  LED_SUSPEND,
	"LED_MUTE":                     LED_MUTE,
	"LED_MISC":                     LED_MISC,
	"LED_MAIL":                     LED_MAIL,
	"LED_CHARGING":                 LED_CHARGING,
	"LED_MAX":                      LED_MAX,
	"REP_DELAY":                    REP_DELAY,
	"REP_PERIOD":                   REP_PERIOD,
	"REP_MAX":                      REP_MAX,
	"SND_CLICK":                    SND_CLICK,
	"SND_BELL":                     SND_BELL,
	"SND_TONE":                     SND_TONE,
	"SND_MAX":                      SND_MAX,
}

var KEY = map[int]string{}
var ABS = map[int]string{}
var REL = map[int]string{}
var SW = map[int]string{}
var MSC = map[int]string{}
var LED = map[int]string{}
var BTN = map[int]string{}
var REP = map[int]string{}
var SND = map[int]string{}
var ID = map[int]string{}
var EV = map[int]string{}
var BUS = map[int]string{}
var SYN = map[int]string{}
var FF = map[int]string{}

var ByEventType = map[int]map[int]string{
	EV_KEY: KEY,
	EV_ABS: ABS,
	EV_REL: REL,
	EV_SW:  SW,
	EV_MSC: MSC,
	EV_LED: LED,
	EV_REP: REP,
	EV_SND: SND,
	EV_SYN: SYN,
	EV_FF:  FF,
}

func init() {
	for code, value := range ecodes {
		switch {
		case strings.HasPrefix(code, "KEY"):
			KEY[value] = code
		case strings.HasPrefix(code, "ABS"):
			ABS[value] = code
		case strings.HasPrefix(code, "REL"):
			REL[value] = code
		case strings.HasPrefix(code, "SW"):
			SW[value] = code
		case strings.HasPrefix(code, "MSC"):
			MSC[value] = code
		case strings.HasPrefix(code, "LED"):
			LED[value] = code
		case strings.HasPrefix(code, "BTN"):
			BTN[value] = code
		case strings.HasPrefix(code, "SND"):
			SND[value] = code
		case strings.HasPrefix(code, "ID"):
			ID[value] = code
		case strings.HasPrefix(code, "EV"):
			EV[value] = code
		case strings.HasPrefix(code, "BUS"):
			BUS[value] = code
		case strings.HasPrefix(code, "SYN"):
			SYN[value] = code
		case strings.HasPrefix(code, "FF"):
			FF[value] = code
		}
	}
}
