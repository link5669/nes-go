package main

import (
	"strings"
)

var truth string = `C000  4C F5 C5  JMP $C5F5                       A:0 X:0 Y:0 P:24 SP:FD CYC:7
C5F5  A2 00     LDX #$00                        A:0 X:0 Y:0 P:24 SP:FD CYC:10
C5F7  86 00     STX $00 = 00                    A:0 X:0 Y:0 P:26 SP:FD CYC:12
C5F9  86 10     STX $10 = 00                    A:0 X:0 Y:0 P:26 SP:FD CYC:15
C5FB  86 11     STX $11 = 00                    A:0 X:0 Y:0 P:26 SP:FD CYC:18
C5FD  20 2D C7  JSR $C72D                       A:0 X:0 Y:0 P:26 SP:FD CYC:21
C72D  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:27
C72E  38        SEC                             A:0 X:0 Y:0 P:26 SP:FB CYC:29
C72F  B0 04     BCS $C735                       A:0 X:0 Y:0 P:27 SP:FB CYC:31
C735  EA        NOP                             A:0 X:0 Y:0 P:27 SP:FB CYC:34
C736  18        CLC                             A:0 X:0 Y:0 P:27 SP:FB CYC:36
C737  B0 03     BCS $C73C                       A:0 X:0 Y:0 P:26 SP:FB CYC:38
C739  4C 40 C7  JMP $C740                       A:0 X:0 Y:0 P:26 SP:FB CYC:40
C740  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:43
C741  38        SEC                             A:0 X:0 Y:0 P:26 SP:FB CYC:45
C742  90 03     BCC $C747                       A:0 X:0 Y:0 P:27 SP:FB CYC:47
C744  4C 4B C7  JMP $C74B                       A:0 X:0 Y:0 P:27 SP:FB CYC:49
C74B  EA        NOP                             A:0 X:0 Y:0 P:27 SP:FB CYC:52
C74C  18        CLC                             A:0 X:0 Y:0 P:27 SP:FB CYC:54
C74D  90 04     BCC $C753                       A:0 X:0 Y:0 P:26 SP:FB CYC:56
C753  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:59
C754  A9 00     LDA #$00                        A:0 X:0 Y:0 P:26 SP:FB CYC:61
C756  F0 04     BEQ $C75C                       A:0 X:0 Y:0 P:26 SP:FB CYC:63
C75C  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:66
C75D  A9 40     LDA #$40                        A:0 X:0 Y:0 P:26 SP:FB CYC:68
C75F  F0 03     BEQ $C764                       A:40 X:0 Y:0 P:24 SP:FB CYC:70
C761  4C 68 C7  JMP $C768                       A:40 X:0 Y:0 P:24 SP:FB CYC:72
C768  EA        NOP                             A:40 X:0 Y:0 P:24 SP:FB CYC:75
C769  A9 40     LDA #$40                        A:40 X:0 Y:0 P:24 SP:FB CYC:77
C76B  D0 04     BNE $C771                       A:40 X:0 Y:0 P:24 SP:FB CYC:79
C771  EA        NOP                             A:40 X:0 Y:0 P:24 SP:FB CYC:82
C772  A9 00     LDA #$00                        A:40 X:0 Y:0 P:24 SP:FB CYC:84
C774  D0 03     BNE $C779                       A:0 X:0 Y:0 P:26 SP:FB CYC:86
C776  4C 7D C7  JMP $C77D                       A:0 X:0 Y:0 P:26 SP:FB CYC:88
C77D  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:91
C77E  A9 FF     LDA #$FF                        A:0 X:0 Y:0 P:26 SP:FB CYC:93
C780  85 01     STA $01 = 00                    A:FF X:0 Y:0 P:A4 SP:FB CYC:95
C782  24 01     BIT $01 = FF                    A:FF X:0 Y:0 P:A4 SP:FB CYC:98
C784  70 04     BVS $C78A                       A:FF X:0 Y:0 P:E4 SP:FB CYC:101
C78A  EA        NOP                             A:FF X:0 Y:0 P:E4 SP:FB CYC:104
C78B  24 01     BIT $01 = FF                    A:FF X:0 Y:0 P:E4 SP:FB CYC:106
C78D  50 03     BVC $C792                       A:FF X:0 Y:0 P:E4 SP:FB CYC:109
C78F  4C 96 C7  JMP $C796                       A:FF X:0 Y:0 P:E4 SP:FB CYC:111
C796  EA        NOP                             A:FF X:0 Y:0 P:E4 SP:FB CYC:114
C797  A9 00     LDA #$00                        A:FF X:0 Y:0 P:E4 SP:FB CYC:116
C799  85 01     STA $01 = FF                    A:0 X:0 Y:0 P:66 SP:FB CYC:118
C79B  24 01     BIT $01 = 00                    A:0 X:0 Y:0 P:66 SP:FB CYC:121
C79D  50 04     BVC $C7A3                       A:0 X:0 Y:0 P:26 SP:FB CYC:124
C7A3  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:127
C7A4  24 01     BIT $01 = 00                    A:0 X:0 Y:0 P:26 SP:FB CYC:129
C7A6  70 03     BVS $C7AB                       A:0 X:0 Y:0 P:26 SP:FB CYC:132
C7A8  4C AF C7  JMP $C7AF                       A:0 X:0 Y:0 P:26 SP:FB CYC:134
C7AF  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:137
C7B0  A9 00     LDA #$00                        A:0 X:0 Y:0 P:26 SP:FB CYC:139
C7B2  10 04     BPL $C7B8                       A:0 X:0 Y:0 P:26 SP:FB CYC:141
C7B8  EA        NOP                             A:0 X:0 Y:0 P:26 SP:FB CYC:144
C7B9  A9 80     LDA #$80                        A:0 X:0 Y:0 P:26 SP:FB CYC:146
C7BB  10 03     BPL $C7C0                       A:80 X:0 Y:0 P:A4 SP:FB CYC:148
C7BD  4C D9 C7  JMP $C7D9                       A:80 X:0 Y:0 P:A4 SP:FB CYC:150
C7D9  EA        NOP                             A:80 X:0 Y:0 P:A4 SP:FB CYC:153
C7DA  60        RTS                             A:80 X:0 Y:0 P:A4 SP:FB CYC:155
C600  20 DB C7  JSR $C7DB                       A:80 X:0 Y:0 P:A4 SP:FD CYC:161
C7DB  EA        NOP                             A:80 X:0 Y:0 P:A4 SP:FB CYC:167
C7DC  A9 FF     LDA #$FF                        A:80 X:0 Y:0 P:A4 SP:FB CYC:169
C7DE  85 01     STA $01 = 00                    A:FF X:0 Y:0 P:A4 SP:FB CYC:171
C7E0  24 01     BIT $01 = FF                    A:FF X:0 Y:0 P:A4 SP:FB CYC:174
C7E2  A9 00     LDA #$00                        A:FF X:0 Y:0 P:E4 SP:FB CYC:177
C7E4  38        SEC                             A:0 X:0 Y:0 P:66 SP:FB CYC:179
C7E5  78        SEI                             A:0 X:0 Y:0 P:67 SP:FB CYC:181
C7E6  F8        SED                             A:0 X:0 Y:0 P:67 SP:FB CYC:183
C7E7  08        PHP                             A:0 X:0 Y:0 P:6F SP:FB CYC:185
C7E8  68        PLA                             A:0 X:0 Y:0 P:6F SP:FA CYC:188
C7E9  29 EF     AND #$EF                        A:7F X:0 Y:0 P:6D SP:FB CYC:192
C7EB  C9 6F     CMP #$6F                        A:6F X:0 Y:0 P:6D SP:FB CYC:194
C7ED  F0 04     BEQ $C7F3                       A:6F X:0 Y:0 P:6F SP:FB CYC:196
C7F3  EA        NOP                             A:6F X:0 Y:0 P:6F SP:FB CYC:199
C7F4  A9 40     LDA #$40                        A:6F X:0 Y:0 P:6F SP:FB CYC:201
C7F6  85 01     STA $01 = FF                    A:40 X:0 Y:0 P:6D SP:FB CYC:203
C7F8  24 01     BIT $01 = 40                    A:40 X:0 Y:0 P:6D SP:FB CYC:206
C7FA  D8        CLD                             A:40 X:0 Y:0 P:6D SP:FB CYC:209
C7FB  A9 10     LDA #$10                        A:40 X:0 Y:0 P:65 SP:FB CYC:211
C7FD  18        CLC                             A:10 X:0 Y:0 P:65 SP:FB CYC:213
C7FE  08        PHP                             A:10 X:0 Y:0 P:64 SP:FB CYC:215
C7FF  68        PLA                             A:10 X:0 Y:0 P:64 SP:FA CYC:218
C800  29 EF     AND #$EF                        A:74 X:0 Y:0 P:64 SP:FB CYC:222
C802  C9 64     CMP #$64                        A:64 X:0 Y:0 P:64 SP:FB CYC:224
C804  F0 04     BEQ $C80A                       A:64 X:0 Y:0 P:67 SP:FB CYC:226
C80A  EA        NOP                             A:64 X:0 Y:0 P:67 SP:FB CYC:229
C80B  A9 80     LDA #$80                        A:64 X:0 Y:0 P:67 SP:FB CYC:231
C80D  85 01     STA $01 = 40                    A:80 X:0 Y:0 P:E5 SP:FB CYC:233
C80F  24 01     BIT $01 = 80                    A:80 X:0 Y:0 P:E5 SP:FB CYC:236
C811  F8        SED                             A:80 X:0 Y:0 P:A5 SP:FB CYC:239
C812  A9 00     LDA #$00                        A:80 X:0 Y:0 P:AD SP:FB CYC:241
C814  38        SEC                             A:0 X:0 Y:0 P:2F SP:FB CYC:243
C815  08        PHP                             A:0 X:0 Y:0 P:2F SP:FB CYC:245
C816  68        PLA                             A:0 X:0 Y:0 P:2F SP:FA CYC:248
C817  29 EF     AND #$EF                        A:3F X:0 Y:0 P:2D SP:FB CYC:252
C819  C9 2F     CMP #$2F                        A:2F X:0 Y:0 P:2D SP:FB CYC:254
C81B  F0 04     BEQ $C821                       A:2F X:0 Y:0 P:2F SP:FB CYC:256
C821  EA        NOP                             A:2F X:0 Y:0 P:2F SP:FB CYC:259
C822  A9 FF     LDA #$FF                        A:2F X:0 Y:0 P:2F SP:FB CYC:261
C824  48        PHA                             A:FF X:0 Y:0 P:AD SP:FB CYC:263
C825  28        PLP                             A:FF X:0 Y:0 P:AD SP:FA CYC:266
C826  D0 09     BNE $C831                       A:FF X:0 Y:0 P:EF SP:FB CYC:270
C828  10 07     BPL $C831                       A:FF X:0 Y:0 P:EF SP:FB CYC:272
C82A  50 05     BVC $C831                       A:FF X:0 Y:0 P:EF SP:FB CYC:274
C82C  90 03     BCC $C831                       A:FF X:0 Y:0 P:EF SP:FB CYC:276
C82E  4C 35 C8  JMP $C835                       A:FF X:0 Y:0 P:EF SP:FB CYC:278
C835  EA        NOP                             A:FF X:0 Y:0 P:EF SP:FB CYC:281
C836  A9 04     LDA #$04                        A:FF X:0 Y:0 P:EF SP:FB CYC:283
C838  48        PHA                             A:4 X:0 Y:0 P:6D SP:FB CYC:285
C839  28        PLP                             A:4 X:0 Y:0 P:6D SP:FA CYC:288
C83A  F0 09     BEQ $C845                       A:4 X:0 Y:0 P:24 SP:FB CYC:292
C83C  30 07     BMI $C845                       A:4 X:0 Y:0 P:24 SP:FB CYC:294
C83E  70 05     BVS $C845                       A:4 X:0 Y:0 P:24 SP:FB CYC:296
C840  B0 03     BCS $C845                       A:4 X:0 Y:0 P:24 SP:FB CYC:298
C842  4C 49 C8  JMP $C849                       A:4 X:0 Y:0 P:24 SP:FB CYC:300
C849  EA        NOP                             A:4 X:0 Y:0 P:24 SP:FB CYC:303
C84A  F8        SED                             A:4 X:0 Y:0 P:24 SP:FB CYC:305
C84B  A9 FF     LDA #$FF                        A:4 X:0 Y:0 P:2C SP:FB CYC:307
C84D  85 01     STA $01 = 80                    A:FF X:0 Y:0 P:AC SP:FB CYC:309
C84F  24 01     BIT $01 = FF                    A:FF X:0 Y:0 P:AC SP:FB CYC:312
C851  18        CLC                             A:FF X:0 Y:0 P:EC SP:FB CYC:315
C852  A9 00     LDA #$00                        A:FF X:0 Y:0 P:EC SP:FB CYC:317
C854  48        PHA                             A:0 X:0 Y:0 P:6E SP:FB CYC:319
C855  A9 FF     LDA #$FF                        A:0 X:0 Y:0 P:6E SP:FA CYC:322
C857  68        PLA                             A:FF X:0 Y:0 P:EC SP:FA CYC:324
C858  D0 09     BNE $C863                       A:0 X:0 Y:0 P:6E SP:FB CYC:328
C85A  30 07     BMI $C863                       A:0 X:0 Y:0 P:6E SP:FB CYC:330
C85C  50 05     BVC $C863                       A:0 X:0 Y:0 P:6E SP:FB CYC:332
C85E  B0 03     BCS $C863                       A:0 X:0 Y:0 P:6E SP:FB CYC:334
C860  4C 67 C8  JMP $C867                       A:0 X:0 Y:0 P:6E SP:FB CYC:336
C867  EA        NOP                             A:0 X:0 Y:0 P:6E SP:FB CYC:339
C868  A9 00     LDA #$00                        A:0 X:0 Y:0 P:6E SP:FB CYC:341
C86A  85 01     STA $01 = FF                    A:0 X:0 Y:0 P:6E SP:FB CYC:343
C86C  24 01     BIT $01 = 00                    A:0 X:0 Y:0 P:6E SP:FB CYC:346
C86E  38        SEC                             A:0 X:0 Y:0 P:2E SP:FB CYC:349
C86F  A9 FF     LDA #$FF                        A:0 X:0 Y:0 P:2F SP:FB CYC:351
C871  48        PHA                             A:FF X:0 Y:0 P:AD SP:FB CYC:353
C872  A9 00     LDA #$00                        A:FF X:0 Y:0 P:AD SP:FA CYC:356
C874  68        PLA                             A:0 X:0 Y:0 P:2F SP:FA CYC:358
C875  F0 09     BEQ $C880                       A:FF X:0 Y:0 P:AD SP:FB CYC:362
C877  10 07     BPL $C880                       A:FF X:0 Y:0 P:AD SP:FB CYC:364
C879  70 05     BVS $C880                       A:FF X:0 Y:0 P:AD SP:FB CYC:366
C87B  90 03     BCC $C880                       A:FF X:0 Y:0 P:AD SP:FB CYC:368
C87D  4C 84 C8  JMP $C884                       A:FF X:0 Y:0 P:AD SP:FB CYC:370
C884  60        RTS                             A:FF X:0 Y:0 P:AD SP:FB CYC:373
C603  20 85 C8  JSR $C885                       A:FF X:0 Y:0 P:AD SP:FD CYC:379
C885  EA        NOP                             A:FF X:0 Y:0 P:AD SP:FB CYC:385
C886  18        CLC                             A:FF X:0 Y:0 P:AD SP:FB CYC:387
C887  A9 FF     LDA #$FF                        A:FF X:0 Y:0 P:AC SP:FB CYC:389
C889  85 01     STA $01 = 00                    A:FF X:0 Y:0 P:AC SP:FB CYC:391
C88B  24 01     BIT $01 = FF                    A:FF X:0 Y:0 P:AC SP:FB CYC:394
C88D  A9 55     LDA #$55                        A:FF X:0 Y:0 P:EC SP:FB CYC:397
C88F  09 AA     ORA #$AA                        A:55 X:0 Y:0 P:6C SP:FB CYC:399
C891  B0 0B     BCS $C89E                       A:FF X:0 Y:0 P:EC SP:FB CYC:401
C893  10 09     BPL $C89E                       A:FF X:0 Y:0 P:EC SP:FB CYC:403
C895  C9 FF     CMP #$FF                        A:FF X:0 Y:0 P:EC SP:FB CYC:405
C897  D0 05     BNE $C89E                       A:FF X:0 Y:0 P:6F SP:FB CYC:407
C899  50 03     BVC $C89E                       A:FF X:0 Y:0 P:6F SP:FB CYC:409
C89B  4C A2 C8  JMP $C8A2                       A:FF X:0 Y:0 P:6F SP:FB CYC:411
C8A2  EA        NOP                             A:FF X:0 Y:0 P:6F SP:FB CYC:414
C8A3  38        SEC                             A:FF X:0 Y:0 P:6F SP:FB CYC:416
C8A4  B8        CLV                             A:FF X:0 Y:0 P:6F SP:FB CYC:418
C8A5  A9 00     LDA #$00                        A:FF X:0 Y:0 P:2F SP:FB CYC:420
C8A7  09 00     ORA #$00                        A:0 X:0 Y:0 P:2F SP:FB CYC:422
C8A9  D0 09     BNE $C8B4                       A:0 X:0 Y:0 P:2F SP:FB CYC:424
C8AB  70 07     BVS $C8B4                       A:0 X:0 Y:0 P:2F SP:FB CYC:426
C8AD  90 05     BCC $C8B4                       A:0 X:0 Y:0 P:2F SP:FB CYC:428
C8AF  30 03     BMI $C8B4                       A:0 X:0 Y:0 P:2F SP:FB CYC:430
C8B1  4C B8 C8  JMP $C8B8                       A:0 X:0 Y:0 P:2F SP:FB CYC:432
C8B8  EA        NOP                             A:0 X:0 Y:0 P:2F SP:FB CYC:435
C8B9  18        CLC                             A:0 X:0 Y:0 P:2F SP:FB CYC:437
C8BA  24 01     BIT $01 = FF                    A:0 X:0 Y:0 P:2E SP:FB CYC:439
C8BC  A9 55     LDA #$55                        A:0 X:0 Y:0 P:EE SP:FB CYC:442
C8BE  29 AA     AND #$AA                        A:55 X:0 Y:0 P:6C SP:FB CYC:444
C8C0  D0 09     BNE $C8CB                       A:0 X:0 Y:0 P:6E SP:FB CYC:446
C8C2  50 07     BVC $C8CB                       A:0 X:0 Y:0 P:6E SP:FB CYC:448
C8C4  B0 05     BCS $C8CB                       A:0 X:0 Y:0 P:6E SP:FB CYC:450
C8C6  30 03     BMI $C8CB                       A:0 X:0 Y:0 P:6E SP:FB CYC:452
C8C8  4C CF C8  JMP $C8CF                       A:0 X:0 Y:0 P:6E SP:FB CYC:454
C8CF  EA        NOP                             A:0 X:0 Y:0 P:6E SP:FB CYC:457
C8D0  38        SEC                             A:0 X:0 Y:0 P:6E SP:FB CYC:459
C8D1  B8        CLV                             A:0 X:0 Y:0 P:6F SP:FB CYC:461
C8D2  A9 F8     LDA #$F8                        A:0 X:0 Y:0 P:2F SP:FB CYC:463
C8D4  29 EF     AND #$EF                        A:F8 X:0 Y:0 P:AD SP:FB CYC:465
C8D6  90 0B     BCC $C8E3                       A:E8 X:0 Y:0 P:AD SP:FB CYC:467
C8D8  10 09     BPL $C8E3                       A:E8 X:0 Y:0 P:AD SP:FB CYC:469
C8DA  C9 E8     CMP #$E8                        A:E8 X:0 Y:0 P:AD SP:FB CYC:471
C8DC  D0 05     BNE $C8E3                       A:E8 X:0 Y:0 P:2F SP:FB CYC:473
C8DE  70 03     BVS $C8E3                       A:E8 X:0 Y:0 P:2F SP:FB CYC:475
C8E0  4C E7 C8  JMP $C8E7                       A:E8 X:0 Y:0 P:2F SP:FB CYC:477
C8E7  EA        NOP                             A:E8 X:0 Y:0 P:2F SP:FB CYC:480
C8E8  18        CLC                             A:E8 X:0 Y:0 P:2F SP:FB CYC:482
C8E9  24 01     BIT $01 = FF                    A:E8 X:0 Y:0 P:2E SP:FB CYC:484
C8EB  A9 5F     LDA #$5F                        A:E8 X:0 Y:0 P:EC SP:FB CYC:487
C8ED  49 AA     EOR #$AA                        A:5F X:0 Y:0 P:6C SP:FB CYC:489
C8EF  B0 0B     BCS $C8FC                       A:F5 X:0 Y:0 P:EC SP:FB CYC:491
C8F1  10 09     BPL $C8FC                       A:F5 X:0 Y:0 P:EC SP:FB CYC:493
C8F3  C9 F5     CMP #$F5                        A:F5 X:0 Y:0 P:EC SP:FB CYC:495
C8F5  D0 05     BNE $C8FC                       A:F5 X:0 Y:0 P:6F SP:FB CYC:497
C8F7  50 03     BVC $C8FC                       A:F5 X:0 Y:0 P:6F SP:FB CYC:499
C8F9  4C 00 C9  JMP $C900                       A:F5 X:0 Y:0 P:6F SP:FB CYC:501
C900  EA        NOP                             A:F5 X:0 Y:0 P:6F SP:FB CYC:504
C901  38        SEC                             A:F5 X:0 Y:0 P:6F SP:FB CYC:506
C902  B8        CLV                             A:F5 X:0 Y:0 P:6F SP:FB CYC:508
C903  A9 70     LDA #$70                        A:F5 X:0 Y:0 P:2F SP:FB CYC:510
C905  49 70     EOR #$70                        A:70 X:0 Y:0 P:2D SP:FB CYC:512
C907  D0 09     BNE $C912                       A:0 X:0 Y:0 P:2F SP:FB CYC:514
C909  70 07     BVS $C912                       A:0 X:0 Y:0 P:2F SP:FB CYC:516
C90B  90 05     BCC $C912                       A:0 X:0 Y:0 P:2F SP:FB CYC:518
C90D  30 03     BMI $C912                       A:0 X:0 Y:0 P:2F SP:FB CYC:520
C90F  4C 16 C9  JMP $C916                       A:0 X:0 Y:0 P:2F SP:FB CYC:522
C916  EA        NOP                             A:0 X:0 Y:0 P:2F SP:FB CYC:525
C917  18        CLC                             A:0 X:0 Y:0 P:2F SP:FB CYC:527
C918  24 01     BIT $01 = FF                    A:0 X:0 Y:0 P:2E SP:FB CYC:529
C91A  A9 00     LDA #$00                        A:0 X:0 Y:0 P:EE SP:FB CYC:532
C91C  69 69     ADC #$69                        A:0 X:0 Y:0 P:6E SP:FB CYC:534
C91E  30 0B     BMI $C92B                       A:69 X:0 Y:0 P:2C SP:FB CYC:536
C920  B0 09     BCS $C92B                       A:69 X:0 Y:0 P:2C SP:FB CYC:538
C922  C9 69     CMP #$69                        A:69 X:0 Y:0 P:2C SP:FB CYC:540
C924  D0 05     BNE $C92B                       A:69 X:0 Y:0 P:2F SP:FB CYC:542
C926  70 03     BVS $C92B                       A:69 X:0 Y:0 P:2F SP:FB CYC:544
C928  4C 2F C9  JMP $C92F                       A:69 X:0 Y:0 P:2F SP:FB CYC:546
C92F  EA        NOP                             A:69 X:0 Y:0 P:2F SP:FB CYC:549
C930  38        SEC                             A:69 X:0 Y:0 P:2F SP:FB CYC:551
C931  F8        SED                             A:69 X:0 Y:0 P:2F SP:FB CYC:553
C932  24 01     BIT $01 = FF                    A:69 X:0 Y:0 P:2F SP:FB CYC:555
C934  A9 01     LDA #$01                        A:69 X:0 Y:0 P:ED SP:FB CYC:558
C936  69 69     ADC #$69                        A:1 X:0 Y:0 P:6D SP:FB CYC:560
C938  30 0B     BMI $C945                       A:6B X:0 Y:0 P:2C SP:FB CYC:562
C93A  B0 09     BCS $C945                       A:6B X:0 Y:0 P:2C SP:FB CYC:564
C93C  C9 6B     CMP #$6B                        A:6B X:0 Y:0 P:2C SP:FB CYC:566
C93E  D0 05     BNE $C945                       A:6B X:0 Y:0 P:2F SP:FB CYC:568
C940  70 03     BVS $C945                       A:6B X:0 Y:0 P:2F SP:FB CYC:570
C942  4C 49 C9  JMP $C949                       A:6B X:0 Y:0 P:2F SP:FB CYC:572
C949  EA        NOP                             A:6B X:0 Y:0 P:2F SP:FB CYC:575
C94A  D8        CLD                             A:6B X:0 Y:0 P:2F SP:FB CYC:577
C94B  38        SEC                             A:6B X:0 Y:0 P:27 SP:FB CYC:579
C94C  B8        CLV                             A:6B X:0 Y:0 P:27 SP:FB CYC:581
C94D  A9 7F     LDA #$7F                        A:6B X:0 Y:0 P:27 SP:FB CYC:583
C94F  69 7F     ADC #$7F                        A:7F X:0 Y:0 P:25 SP:FB CYC:585
C951  10 0B     BPL $C95E                       A:FF X:0 Y:0 P:E4 SP:FB CYC:587
C953  B0 09     BCS $C95E                       A:FF X:0 Y:0 P:E4 SP:FB CYC:589
C955  C9 FF     CMP #$FF                        A:FF X:0 Y:0 P:E4 SP:FB CYC:591
C957  D0 05     BNE $C95E                       A:FF X:0 Y:0 P:67 SP:FB CYC:593
C959  50 03     BVC $C95E                       A:FF X:0 Y:0 P:67 SP:FB CYC:595
C95B  4C 62 C9  JMP $C962                       A:FF X:0 Y:0 P:67 SP:FB CYC:597
C962  EA        NOP                             A:FF X:0 Y:0 P:67 SP:FB CYC:600
C963  18        CLC                             A:FF X:0 Y:0 P:67 SP:FB CYC:602
C964  24 01     BIT $01 = FF                    A:FF X:0 Y:0 P:66 SP:FB CYC:604
C966  A9 7F     LDA #$7F                        A:FF X:0 Y:0 P:E4 SP:FB CYC:607
C968  69 80     ADC #$80                        A:7F X:0 Y:0 P:64 SP:FB CYC:609
C96A  10 0B     BPL $C977                       A:FF X:0 Y:0 P:A4 SP:FB CYC:611
C96C  B0 09     BCS $C977                       A:FF X:0 Y:0 P:A4 SP:FB CYC:613
C96E  C9 FF     CMP #$FF                        A:FF X:0 Y:0 P:A4 SP:FB CYC:615
C970  D0 05     BNE $C977                       A:FF X:0 Y:0 P:27 SP:FB CYC:617
C972  70 03     BVS $C977                       A:FF X:0 Y:0 P:27 SP:FB CYC:619
C974  4C 7B C9  JMP $C97B                       A:FF X:0 Y:0 P:27 SP:FB CYC:621
C97B  EA        NOP                             A:FF X:0 Y:0 P:27 SP:FB CYC:624
C97C  38        SEC                             A:FF X:0 Y:0 P:27 SP:FB CYC:626
C97D  B8        CLV                             A:FF X:0 Y:0 P:27 SP:FB CYC:628
C97E  A9 7F     LDA #$7F                        A:FF X:0 Y:0 P:27 SP:FB CYC:630
C980  69 80     ADC #$80                        A:7F X:0 Y:0 P:25 SP:FB CYC:632
C982  D0 09     BNE $C98D                       A:0 X:0 Y:0 P:27 SP:FB CYC:634
C984  30 07     BMI $C98D                       A:0 X:0 Y:0 P:27 SP:FB CYC:636
C986  70 05     BVS $C98D                       A:0 X:0 Y:0 P:27 SP:FB CYC:638
C988  90 03     BCC $C98D                       A:0 X:0 Y:0 P:27 SP:FB CYC:640
C98A  4C 91 C9  JMP $C991                       A:0 X:0 Y:0 P:27 SP:FB CYC:642
C991  EA        NOP                             A:0 X:0 Y:0 P:27 SP:FB CYC:645
C992  38        SEC                             A:0 X:0 Y:0 P:27 SP:FB CYC:647
C993  B8        CLV                             A:0 X:0 Y:0 P:27 SP:FB CYC:649
C994  A9 9F     LDA #$9F                        A:0 X:0 Y:0 P:27 SP:FB CYC:651
C996  F0 09     BEQ $C9A1                       A:9F X:0 Y:0 P:A5 SP:FB CYC:653
C998  10 07     BPL $C9A1                       A:9F X:0 Y:0 P:A5 SP:FB CYC:655
C99A  70 05     BVS $C9A1                       A:9F X:0 Y:0 P:A5 SP:FB CYC:657
C99C  90 03     BCC $C9A1                       A:9F X:0 Y:0 P:A5 SP:FB CYC:659
C99E  4C A5 C9  JMP $C9A5                       A:9F X:0 Y:0 P:A5 SP:FB CYC:661
C9A5  EA        NOP                             A:9F X:0 Y:0 P:A5 SP:FB CYC:664
C9A6  18        CLC                             A:9F X:0 Y:0 P:A5 SP:FB CYC:666
C9A7  24 01     BIT $01 = FF                    A:9F X:0 Y:0 P:A4 SP:FB CYC:668
C9A9  A9 00     LDA #$00                        A:9F X:0 Y:0 P:E4 SP:FB CYC:671
C9AB  D0 09     BNE $C9B6                       A:0 X:0 Y:0 P:66 SP:FB CYC:673
C9AD  30 07     BMI $C9B6                       A:0 X:0 Y:0 P:66 SP:FB CYC:675
C9AF  50 05     BVC $C9B6                       A:0 X:0 Y:0 P:66 SP:FB CYC:677
C9B1  B0 03     BCS $C9B6                       A:0 X:0 Y:0 P:66 SP:FB CYC:679
C9B3  4C BA C9  JMP $C9BA                       A:0 X:0 Y:0 P:66 SP:FB CYC:681
C9BA  EA        NOP                             A:0 X:0 Y:0 P:66 SP:FB CYC:684
C9BB  24 01     BIT $01 = FF                    A:0 X:0 Y:0 P:66 SP:FB CYC:686
C9BD  A9 40     LDA #$40                        A:0 X:0 Y:0 P:E6 SP:FB CYC:689
C9BF  C9 40     CMP #$40                        A:40 X:0 Y:0 P:64 SP:FB CYC:691
C9C1  30 09     BMI $C9CC                       A:40 X:0 Y:0 P:67 SP:FB CYC:693
C9C3  90 07     BCC $C9CC                       A:40 X:0 Y:0 P:67 SP:FB CYC:695
C9C5  D0 05     BNE $C9CC                       A:40 X:0 Y:0 P:67 SP:FB CYC:697
C9C7  50 03     BVC $C9CC                       A:40 X:0 Y:0 P:67 SP:FB CYC:699
C9C9  4C D0 C9  JMP $C9D0                       A:40 X:0 Y:0 P:67 SP:FB CYC:701
C9D0  EA        NOP                             A:40 X:0 Y:0 P:67 SP:FB CYC:704
C9D1  B8        CLV                             A:40 X:0 Y:0 P:67 SP:FB CYC:706
C9D2  C9 3F     CMP #$3F                        A:40 X:0 Y:0 P:27 SP:FB CYC:708
C9D4  F0 09     BEQ $C9DF                       A:40 X:0 Y:0 P:25 SP:FB CYC:710
C9D6  30 07     BMI $C9DF                       A:40 X:0 Y:0 P:25 SP:FB CYC:712
C9D8  90 05     BCC $C9DF                       A:40 X:0 Y:0 P:25 SP:FB CYC:714
C9DA  70 03     BVS $C9DF                       A:40 X:0 Y:0 P:25 SP:FB CYC:716
C9DC  4C E3 C9  JMP $C9E3                       A:40 X:0 Y:0 P:25 SP:FB CYC:718
C9E3  EA        NOP                             A:40 X:0 Y:0 P:25 SP:FB CYC:721
C9E4  C9 41     CMP #$41                        A:40 X:0 Y:0 P:25 SP:FB CYC:723
C9E6  F0 07     BEQ $C9EF                       A:40 X:0 Y:0 P:A4 SP:FB CYC:725
C9E8  10 05     BPL $C9EF                       A:40 X:0 Y:0 P:A4 SP:FB CYC:727
C9EA  10 03     BPL $C9EF                       A:40 X:0 Y:0 P:A4 SP:FB CYC:729
C9EC  4C F3 C9  JMP $C9F3                       A:40 X:0 Y:0 P:A4 SP:FB CYC:731
C9F3  EA        NOP                             A:40 X:0 Y:0 P:A4 SP:FB CYC:734
C9F4  A9 80     LDA #$80                        A:40 X:0 Y:0 P:A4 SP:FB CYC:736
C9F6  C9 00     CMP #$00                        A:80 X:0 Y:0 P:A4 SP:FB CYC:738
C9F8  F0 07     BEQ $CA01                       A:80 X:0 Y:0 P:A5 SP:FB CYC:740
C9FA  10 05     BPL $CA01                       A:80 X:0 Y:0 P:A5 SP:FB CYC:742
C9FC  90 03     BCC $CA01                       A:80 X:0 Y:0 P:A5 SP:FB CYC:744
C9FE  4C 05 CA  JMP $CA05                       A:80 X:0 Y:0 P:A5 SP:FB CYC:746
CA05  EA        NOP                             A:80 X:0 Y:0 P:A5 SP:FB CYC:749
CA06  C9 80     CMP #$80                        A:80 X:0 Y:0 P:A5 SP:FB CYC:751
CA08  D0 07     BNE $CA11                       A:80 X:0 Y:0 P:27 SP:FB CYC:753
CA0A  30 05     BMI $CA11                       A:80 X:0 Y:0 P:27 SP:FB CYC:755
CA0C  90 03     BCC $CA11                       A:80 X:0 Y:0 P:27 SP:FB CYC:757
CA0E  4C 15 CA  JMP $CA15                       A:80 X:0 Y:0 P:27 SP:FB CYC:759
CA15  EA        NOP                             A:80 X:0 Y:0 P:27 SP:FB CYC:762
CA16  C9 81     CMP #$81                        A:80 X:0 Y:0 P:27 SP:FB CYC:764
CA18  B0 07     BCS $CA21                       A:80 X:0 Y:0 P:A4 SP:FB CYC:766
CA1A  F0 05     BEQ $CA21                       A:80 X:0 Y:0 P:A4 SP:FB CYC:768
CA1C  10 03     BPL $CA21                       A:80 X:0 Y:0 P:A4 SP:FB CYC:770
CA1E  4C 25 CA  JMP $CA25                       A:80 X:0 Y:0 P:A4 SP:FB CYC:772
CA25  EA        NOP                             A:80 X:0 Y:0 P:A4 SP:FB CYC:775
CA26  C9 7F     CMP #$7F                        A:80 X:0 Y:0 P:A4 SP:FB CYC:777
CA28  90 07     BCC $CA31                       A:80 X:0 Y:0 P:25 SP:FB CYC:779
CA2A  F0 05     BEQ $CA31                       A:80 X:0 Y:0 P:25 SP:FB CYC:781
CA2C  30 03     BMI $CA31                       A:80 X:0 Y:0 P:25 SP:FB CYC:783
CA2E  4C 35 CA  JMP $CA35                       A:80 X:0 Y:0 P:25 SP:FB CYC:785
CA35  EA        NOP                             A:80 X:0 Y:0 P:25 SP:FB CYC:788
CA36  24 01     BIT $01 = FF                    A:80 X:0 Y:0 P:25 SP:FB CYC:790
CA38  A0 40     LDY #$40                        A:80 X:0 Y:0 P:E5 SP:FB CYC:793
CA3A  C0 40     CPY #$40                        A:80 X:0 Y:40 P:65 SP:FB CYC:795
CA3C  D0 09     BNE $CA47                       A:80 X:0 Y:40 P:67 SP:FB CYC:797
CA3E  30 07     BMI $CA47                       A:80 X:0 Y:40 P:67 SP:FB CYC:799
CA40  90 05     BCC $CA47                       A:80 X:0 Y:40 P:67 SP:FB CYC:801
CA42  50 03     BVC $CA47                       A:80 X:0 Y:40 P:67 SP:FB CYC:803
CA44  4C 4B CA  JMP $CA4B                       A:80 X:0 Y:40 P:67 SP:FB CYC:805
CA4B  EA        NOP                             A:80 X:0 Y:40 P:67 SP:FB CYC:808
CA4C  B8        CLV                             A:80 X:0 Y:40 P:67 SP:FB CYC:810
CA4D  C0 3F     CPY #$3F                        A:80 X:0 Y:40 P:27 SP:FB CYC:812
CA4F  F0 09     BEQ $CA5A                       A:80 X:0 Y:40 P:25 SP:FB CYC:814
CA51  30 07     BMI $CA5A                       A:80 X:0 Y:40 P:25 SP:FB CYC:816
CA53  90 05     BCC $CA5A                       A:80 X:0 Y:40 P:25 SP:FB CYC:818
CA55  70 03     BVS $CA5A                       A:80 X:0 Y:40 P:25 SP:FB CYC:820
CA57  4C 5E CA  JMP $CA5E                       A:80 X:0 Y:40 P:25 SP:FB CYC:822
CA5E  EA        NOP                             A:80 X:0 Y:40 P:25 SP:FB CYC:825
CA5F  C0 41     CPY #$41                        A:80 X:0 Y:40 P:25 SP:FB CYC:827
CA61  F0 07     BEQ $CA6A                       A:80 X:0 Y:40 P:A4 SP:FB CYC:829
CA63  10 05     BPL $CA6A                       A:80 X:0 Y:40 P:A4 SP:FB CYC:831
CA65  10 03     BPL $CA6A                       A:80 X:0 Y:40 P:A4 SP:FB CYC:833
CA67  4C 6E CA  JMP $CA6E                       A:80 X:0 Y:40 P:A4 SP:FB CYC:835
CA6E  EA        NOP                             A:80 X:0 Y:40 P:A4 SP:FB CYC:838
CA6F  A0 80     LDY #$80                        A:80 X:0 Y:40 P:A4 SP:FB CYC:840
CA71  C0 00     CPY #$00                        A:80 X:0 Y:80 P:A4 SP:FB CYC:842
CA73  F0 07     BEQ $CA7C                       A:80 X:0 Y:80 P:A5 SP:FB CYC:844
CA75  10 05     BPL $CA7C                       A:80 X:0 Y:80 P:A5 SP:FB CYC:846
CA77  90 03     BCC $CA7C                       A:80 X:0 Y:80 P:A5 SP:FB CYC:848
CA79  4C 80 CA  JMP $CA80                       A:80 X:0 Y:80 P:A5 SP:FB CYC:850
CA80  EA        NOP                             A:80 X:0 Y:80 P:A5 SP:FB CYC:853
CA81  C0 80     CPY #$80                        A:80 X:0 Y:80 P:A5 SP:FB CYC:855
CA83  D0 07     BNE $CA8C                       A:80 X:0 Y:80 P:27 SP:FB CYC:857
CA85  30 05     BMI $CA8C                       A:80 X:0 Y:80 P:27 SP:FB CYC:859
CA87  90 03     BCC $CA8C                       A:80 X:0 Y:80 P:27 SP:FB CYC:861
CA89  4C 90 CA  JMP $CA90                       A:80 X:0 Y:80 P:27 SP:FB CYC:863
CA90  EA        NOP                             A:80 X:0 Y:80 P:27 SP:FB CYC:866
CA91  C0 81     CPY #$81                        A:80 X:0 Y:80 P:27 SP:FB CYC:868
CA93  B0 07     BCS $CA9C                       A:80 X:0 Y:80 P:A4 SP:FB CYC:870
CA95  F0 05     BEQ $CA9C                       A:80 X:0 Y:80 P:A4 SP:FB CYC:872
CA97  10 03     BPL $CA9C                       A:80 X:0 Y:80 P:A4 SP:FB CYC:874
CA99  4C A0 CA  JMP $CAA0                       A:80 X:0 Y:80 P:A4 SP:FB CYC:876
CAA0  EA        NOP                             A:80 X:0 Y:80 P:A4 SP:FB CYC:879
CAA1  C0 7F     CPY #$7F                        A:80 X:0 Y:80 P:A4 SP:FB CYC:881
CAA3  90 07     BCC $CAAC                       A:80 X:0 Y:80 P:25 SP:FB CYC:883
CAA5  F0 05     BEQ $CAAC                       A:80 X:0 Y:80 P:25 SP:FB CYC:885
CAA7  30 03     BMI $CAAC                       A:80 X:0 Y:80 P:25 SP:FB CYC:887
CAA9  4C B0 CA  JMP $CAB0                       A:80 X:0 Y:80 P:25 SP:FB CYC:889
CAB0  EA        NOP                             A:80 X:0 Y:80 P:25 SP:FB CYC:892
CAB1  24 01     BIT $01 = FF                    A:80 X:0 Y:80 P:25 SP:FB CYC:894
CAB3  A2 40     LDX #$40                        A:80 X:0 Y:80 P:E5 SP:FB CYC:897
CAB5  E0 40     CPX #$40                        A:80 X:40 Y:80 P:65 SP:FB CYC:899
CAB7  D0 09     BNE $CAC2                       A:80 X:40 Y:80 P:67 SP:FB CYC:901
CAB9  30 07     BMI $CAC2                       A:80 X:40 Y:80 P:67 SP:FB CYC:903
CABB  90 05     BCC $CAC2                       A:80 X:40 Y:80 P:67 SP:FB CYC:905
CABD  50 03     BVC $CAC2                       A:80 X:40 Y:80 P:67 SP:FB CYC:907
CABF  4C C6 CA  JMP $CAC6                       A:80 X:40 Y:80 P:67 SP:FB CYC:909
CAC6  EA        NOP                             A:80 X:40 Y:80 P:67 SP:FB CYC:912
CAC7  B8        CLV                             A:80 X:40 Y:80 P:67 SP:FB CYC:914
CAC8  E0 3F     CPX #$3F                        A:80 X:40 Y:80 P:27 SP:FB CYC:916
CACA  F0 09     BEQ $CAD5                       A:80 X:40 Y:80 P:25 SP:FB CYC:918
CACC  30 07     BMI $CAD5                       A:80 X:40 Y:80 P:25 SP:FB CYC:920
CACE  90 05     BCC $CAD5                       A:80 X:40 Y:80 P:25 SP:FB CYC:922
CAD0  70 03     BVS $CAD5                       A:80 X:40 Y:80 P:25 SP:FB CYC:924
CAD2  4C D9 CA  JMP $CAD9                       A:80 X:40 Y:80 P:25 SP:FB CYC:926
CAD9  EA        NOP                             A:80 X:40 Y:80 P:25 SP:FB CYC:929
CADA  E0 41     CPX #$41                        A:80 X:40 Y:80 P:25 SP:FB CYC:931
CADC  F0 07     BEQ $CAE5                       A:80 X:40 Y:80 P:A4 SP:FB CYC:933
CADE  10 05     BPL $CAE5                       A:80 X:40 Y:80 P:A4 SP:FB CYC:935
CAE0  10 03     BPL $CAE5                       A:80 X:40 Y:80 P:A4 SP:FB CYC:937
CAE2  4C E9 CA  JMP $CAE9                       A:80 X:40 Y:80 P:A4 SP:FB CYC:939
CAE9  EA        NOP                             A:80 X:40 Y:80 P:A4 SP:FB CYC:942
CAEA  A2 80     LDX #$80                        A:80 X:40 Y:80 P:A4 SP:FB CYC:944
CAEC  E0 00     CPX #$00                        A:80 X:80 Y:80 P:A4 SP:FB CYC:946
CAEE  F0 07     BEQ $CAF7                       A:80 X:80 Y:80 P:A5 SP:FB CYC:948
CAF0  10 05     BPL $CAF7                       A:80 X:80 Y:80 P:A5 SP:FB CYC:950
CAF2  90 03     BCC $CAF7                       A:80 X:80 Y:80 P:A5 SP:FB CYC:952
CAF4  4C FB CA  JMP $CAFB                       A:80 X:80 Y:80 P:A5 SP:FB CYC:954
CAFB  EA        NOP                             A:80 X:80 Y:80 P:A5 SP:FB CYC:957
CAFC  E0 80     CPX #$80                        A:80 X:80 Y:80 P:A5 SP:FB CYC:959
CAFE  D0 07     BNE $CB07                       A:80 X:80 Y:80 P:27 SP:FB CYC:961
CB00  30 05     BMI $CB07                       A:80 X:80 Y:80 P:27 SP:FB CYC:963
CB02  90 03     BCC $CB07                       A:80 X:80 Y:80 P:27 SP:FB CYC:965
CB04  4C 0B CB  JMP $CB0B                       A:80 X:80 Y:80 P:27 SP:FB CYC:967
CB0B  EA        NOP                             A:80 X:80 Y:80 P:27 SP:FB CYC:970
CB0C  E0 81     CPX #$81                        A:80 X:80 Y:80 P:27 SP:FB CYC:972
CB0E  B0 07     BCS $CB17                       A:80 X:80 Y:80 P:A4 SP:FB CYC:974
CB10  F0 05     BEQ $CB17                       A:80 X:80 Y:80 P:A4 SP:FB CYC:976
CB12  10 03     BPL $CB17                       A:80 X:80 Y:80 P:A4 SP:FB CYC:978
CB14  4C 1B CB  JMP $CB1B                       A:80 X:80 Y:80 P:A4 SP:FB CYC:980
CB1B  EA        NOP                             A:80 X:80 Y:80 P:A4 SP:FB CYC:983
CB1C  E0 7F     CPX #$7F                        A:80 X:80 Y:80 P:A4 SP:FB CYC:985
CB1E  90 07     BCC $CB27                       A:80 X:80 Y:80 P:25 SP:FB CYC:987
CB20  F0 05     BEQ $CB27                       A:80 X:80 Y:80 P:25 SP:FB CYC:989
CB22  30 03     BMI $CB27                       A:80 X:80 Y:80 P:25 SP:FB CYC:991
CB24  4C 2B CB  JMP $CB2B                       A:80 X:80 Y:80 P:25 SP:FB CYC:993
CB2B  EA        NOP                             A:80 X:80 Y:80 P:25 SP:FB CYC:996
CB2C  38        SEC                             A:80 X:80 Y:80 P:25 SP:FB CYC:998
CB2D  B8        CLV                             A:80 X:80 Y:80 P:25 SP:FB CYC:1000
CB2E  A2 9F     LDX #$9F                        A:80 X:80 Y:80 P:25 SP:FB CYC:1002
CB30  F0 09     BEQ $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB CYC:1004
CB32  10 07     BPL $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB CYC:1006
CB34  70 05     BVS $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB CYC:1008
CB36  90 03     BCC $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB CYC:1010
CB38  4C 3F CB  JMP $CB3F                       A:80 X:9F Y:80 P:A5 SP:FB CYC:1012
CB3F  EA        NOP                             A:80 X:9F Y:80 P:A5 SP:FB CYC:1015
CB40  18        CLC                             A:80 X:9F Y:80 P:A5 SP:FB CYC:1017
CB41  24 01     BIT $01 = FF                    A:80 X:9F Y:80 P:A4 SP:FB CYC:1019
CB43  A2 00     LDX #$00                        A:80 X:9F Y:80 P:E4 SP:FB CYC:1022
CB45  D0 09     BNE $CB50                       A:80 X:0 Y:80 P:66 SP:FB CYC:1024
CB47  30 07     BMI $CB50                       A:80 X:0 Y:80 P:66 SP:FB CYC:1026
CB49  50 05     BVC $CB50                       A:80 X:0 Y:80 P:66 SP:FB CYC:1028
CB4B  B0 03     BCS $CB50                       A:80 X:0 Y:80 P:66 SP:FB CYC:1030
CB4D  4C 54 CB  JMP $CB54                       A:80 X:0 Y:80 P:66 SP:FB CYC:1032
CB54  EA        NOP                             A:80 X:0 Y:80 P:66 SP:FB CYC:1035
CB55  38        SEC                             A:80 X:0 Y:80 P:66 SP:FB CYC:1037
CB56  B8        CLV                             A:80 X:0 Y:80 P:67 SP:FB CYC:1039
CB57  A0 9F     LDY #$9F                        A:80 X:0 Y:80 P:27 SP:FB CYC:1041
CB59  F0 09     BEQ $CB64                       A:80 X:0 Y:9F P:A5 SP:FB CYC:1043
CB5B  10 07     BPL $CB64                       A:80 X:0 Y:9F P:A5 SP:FB CYC:1045
CB5D  70 05     BVS $CB64                       A:80 X:0 Y:9F P:A5 SP:FB CYC:1047
CB5F  90 03     BCC $CB64                       A:80 X:0 Y:9F P:A5 SP:FB CYC:1049
CB61  4C 68 CB  JMP $CB68                       A:80 X:0 Y:9F P:A5 SP:FB CYC:1051
CB68  EA        NOP                             A:80 X:0 Y:9F P:A5 SP:FB CYC:1054
CB69  18        CLC                             A:80 X:0 Y:9F P:A5 SP:FB CYC:1056
CB6A  24 01     BIT $01 = FF                    A:80 X:0 Y:9F P:A4 SP:FB CYC:1058
CB6C  A0 00     LDY #$00                        A:80 X:0 Y:9F P:E4 SP:FB CYC:1061
CB6E  D0 09     BNE $CB79                       A:80 X:0 Y:0 P:66 SP:FB CYC:1063
CB70  30 07     BMI $CB79                       A:80 X:0 Y:0 P:66 SP:FB CYC:1065
CB72  50 05     BVC $CB79                       A:80 X:0 Y:0 P:66 SP:FB CYC:1067
CB74  B0 03     BCS $CB79                       A:80 X:0 Y:0 P:66 SP:FB CYC:1069
CB76  4C 7D CB  JMP $CB7D                       A:80 X:0 Y:0 P:66 SP:FB CYC:1071
CB7D  EA        NOP                             A:80 X:0 Y:0 P:66 SP:FB CYC:1074
CB7E  A9 55     LDA #$55                        A:80 X:0 Y:0 P:66 SP:FB CYC:1076
CB80  A2 AA     LDX #$AA                        A:55 X:0 Y:0 P:64 SP:FB CYC:1078
CB82  A0 33     LDY #$33                        A:55 X:AA Y:0 P:E4 SP:FB CYC:1080
CB84  C9 55     CMP #$55                        A:55 X:AA Y:33 P:64 SP:FB CYC:1082
CB86  D0 23     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB CYC:1084
CB88  E0 AA     CPX #$AA                        A:55 X:AA Y:33 P:67 SP:FB CYC:1086
CB8A  D0 1F     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB CYC:1088
CB8C  C0 33     CPY #$33                        A:55 X:AA Y:33 P:67 SP:FB CYC:1090
CB8E  D0 1B     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB CYC:1092
CB90  C9 55     CMP #$55                        A:55 X:AA Y:33 P:67 SP:FB CYC:1094
CB92  D0 17     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB CYC:1096
CB94  E0 AA     CPX #$AA                        A:55 X:AA Y:33 P:67 SP:FB CYC:1098
CB96  D0 13     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB CYC:1100
CB98  C0 33     CPY #$33                        A:55 X:AA Y:33 P:67 SP:FB CYC:1102
CB9A  D0 0F     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB CYC:1104
CB9C  C9 56     CMP #$56                        A:55 X:AA Y:33 P:67 SP:FB CYC:1106
CB9E  F0 0B     BEQ $CBAB                       A:55 X:AA Y:33 P:E4 SP:FB CYC:1108
CBA0  E0 AB     CPX #$AB                        A:55 X:AA Y:33 P:E4 SP:FB CYC:1110
CBA2  F0 07     BEQ $CBAB                       A:55 X:AA Y:33 P:E4 SP:FB CYC:1112
CBA4  C0 34     CPY #$34                        A:55 X:AA Y:33 P:E4 SP:FB CYC:1114
CBA6  F0 03     BEQ $CBAB                       A:55 X:AA Y:33 P:E4 SP:FB CYC:1116
CBA8  4C AF CB  JMP $CBAF                       A:55 X:AA Y:33 P:E4 SP:FB CYC:1118
CBAF  A0 71     LDY #$71                        A:55 X:AA Y:33 P:E4 SP:FB CYC:1121
CBB1  20 31 F9  JSR $F931                       A:55 X:AA Y:71 P:64 SP:FB CYC:1123
F931  24 01     BIT $01 = FF                    A:55 X:AA Y:71 P:64 SP:F9 CYC:1129
F933  A9 40     LDA #$40                        A:55 X:AA Y:71 P:E4 SP:F9 CYC:1132
F935  38        SEC                             A:40 X:AA Y:71 P:64 SP:F9 CYC:1134
F936  60        RTS                             A:40 X:AA Y:71 P:65 SP:F9 CYC:1136
CBB4  E9 40     SBC #$40                        A:40 X:AA Y:71 P:65 SP:FB CYC:1142
CBB6  20 37 F9  JSR $F937                       A:0 X:AA Y:71 P:27 SP:FB CYC:1144
F937  30 0B     BMI $F944                       A:0 X:AA Y:71 P:27 SP:F9 CYC:1150
F939  90 09     BCC $F944                       A:0 X:AA Y:71 P:27 SP:F9 CYC:1152
F93B  D0 07     BNE $F944                       A:0 X:AA Y:71 P:27 SP:F9 CYC:1154
F93D  70 05     BVS $F944                       A:0 X:AA Y:71 P:27 SP:F9 CYC:1156
F93F  C9 00     CMP #$00                        A:0 X:AA Y:71 P:27 SP:F9 CYC:1158
F941  D0 01     BNE $F944                       A:0 X:AA Y:71 P:27 SP:F9 CYC:1160
F943  60        RTS                             A:0 X:AA Y:71 P:27 SP:F9 CYC:1162
CBB9  C8        INY                             A:0 X:AA Y:71 P:27 SP:FB CYC:1168
CBBA  20 47 F9  JSR $F947                       A:0 X:AA Y:72 P:25 SP:FB CYC:1170
F947  B8        CLV                             A:0 X:AA Y:72 P:25 SP:F9 CYC:1176
F948  38        SEC                             A:0 X:AA Y:72 P:25 SP:F9 CYC:1178
F949  A9 40     LDA #$40                        A:0 X:AA Y:72 P:25 SP:F9 CYC:1180
F94B  60        RTS                             A:40 X:AA Y:72 P:25 SP:F9 CYC:1182
CBBD  E9 3F     SBC #$3F                        A:40 X:AA Y:72 P:25 SP:FB CYC:1188
CBBF  20 4C F9  JSR $F94C                       A:1 X:AA Y:72 P:25 SP:FB CYC:1190
F94C  F0 0B     BEQ $F959                       A:1 X:AA Y:72 P:25 SP:F9 CYC:1196
F94E  30 09     BMI $F959                       A:1 X:AA Y:72 P:25 SP:F9 CYC:1198
F950  90 07     BCC $F959                       A:1 X:AA Y:72 P:25 SP:F9 CYC:1200
F952  70 05     BVS $F959                       A:1 X:AA Y:72 P:25 SP:F9 CYC:1202
F954  C9 01     CMP #$01                        A:1 X:AA Y:72 P:25 SP:F9 CYC:1204
F956  D0 01     BNE $F959                       A:1 X:AA Y:72 P:27 SP:F9 CYC:1206
F958  60        RTS                             A:1 X:AA Y:72 P:27 SP:F9 CYC:1208
CBC2  C8        INY                             A:1 X:AA Y:72 P:27 SP:FB CYC:1214
CBC3  20 5C F9  JSR $F95C                       A:1 X:AA Y:73 P:25 SP:FB CYC:1216
F95C  A9 40     LDA #$40                        A:1 X:AA Y:73 P:25 SP:F9 CYC:1222
F95E  38        SEC                             A:40 X:AA Y:73 P:25 SP:F9 CYC:1224
F95F  24 01     BIT $01 = FF                    A:40 X:AA Y:73 P:25 SP:F9 CYC:1226
F961  60        RTS                             A:40 X:AA Y:73 P:E5 SP:F9 CYC:1229
CBC6  E9 41     SBC #$41                        A:40 X:AA Y:73 P:E5 SP:FB CYC:1235
CBC8  20 62 F9  JSR $F962                       A:FF X:AA Y:73 P:A4 SP:FB CYC:1237
F962  B0 0B     BCS $F96F                       A:FF X:AA Y:73 P:A4 SP:F9 CYC:1243
F964  F0 09     BEQ $F96F                       A:FF X:AA Y:73 P:A4 SP:F9 CYC:1245
F966  10 07     BPL $F96F                       A:FF X:AA Y:73 P:A4 SP:F9 CYC:1247
F968  70 05     BVS $F96F                       A:FF X:AA Y:73 P:A4 SP:F9 CYC:1249
F96A  C9 FF     CMP #$FF                        A:FF X:AA Y:73 P:A4 SP:F9 CYC:1251
F96C  D0 01     BNE $F96F                       A:FF X:AA Y:73 P:27 SP:F9 CYC:1253
F96E  60        RTS                             A:FF X:AA Y:73 P:27 SP:F9 CYC:1255
CBCB  C8        INY                             A:FF X:AA Y:73 P:27 SP:FB CYC:1261
CBCC  20 72 F9  JSR $F972                       A:FF X:AA Y:74 P:25 SP:FB CYC:1263
F972  18        CLC                             A:FF X:AA Y:74 P:25 SP:F9 CYC:1269
F973  A9 80     LDA #$80                        A:FF X:AA Y:74 P:24 SP:F9 CYC:1271
F975  60        RTS                             A:80 X:AA Y:74 P:A4 SP:F9 CYC:1273
CBCF  E9 00     SBC #$00                        A:80 X:AA Y:74 P:A4 SP:FB CYC:1279
CBD1  20 76 F9  JSR $F976                       A:7F X:AA Y:74 P:65 SP:FB CYC:1281
F976  90 05     BCC $F97D                       A:7F X:AA Y:74 P:65 SP:F9 CYC:1287
F978  C9 7F     CMP #$7F                        A:7F X:AA Y:74 P:65 SP:F9 CYC:1289
F97A  D0 01     BNE $F97D                       A:7F X:AA Y:74 P:67 SP:F9 CYC:1291
F97C  60        RTS                             A:7F X:AA Y:74 P:67 SP:F9 CYC:1293
CBD4  C8        INY                             A:7F X:AA Y:74 P:67 SP:FB CYC:1299
CBD5  20 80 F9  JSR $F980                       A:7F X:AA Y:75 P:65 SP:FB CYC:1301
F980  38        SEC                             A:7F X:AA Y:75 P:65 SP:F9 CYC:1307
F981  A9 81     LDA #$81                        A:7F X:AA Y:75 P:65 SP:F9 CYC:1309
F983  60        RTS                             A:81 X:AA Y:75 P:E5 SP:F9 CYC:1311
CBD8  E9 7F     SBC #$7F                        A:81 X:AA Y:75 P:E5 SP:FB CYC:1317
CBDA  20 84 F9  JSR $F984                       A:2 X:AA Y:75 P:65 SP:FB CYC:1319
F984  50 07     BVC $F98D                       A:2 X:AA Y:75 P:65 SP:F9 CYC:1325
F986  90 05     BCC $F98D                       A:2 X:AA Y:75 P:65 SP:F9 CYC:1327
F988  C9 02     CMP #$02                        A:2 X:AA Y:75 P:65 SP:F9 CYC:1329
F98A  D0 01     BNE $F98D                       A:2 X:AA Y:75 P:67 SP:F9 CYC:1331
F98C  60        RTS                             A:2 X:AA Y:75 P:67 SP:F9 CYC:1333
CBDD  60        RTS                             A:2 X:AA Y:75 P:67 SP:FB CYC:1339
C606  20 DE CB  JSR $CBDE                       A:2 X:AA Y:75 P:67 SP:FD CYC:1345
CBDE  EA        NOP                             A:2 X:AA Y:75 P:67 SP:FB CYC:1351
CBDF  A9 FF     LDA #$FF                        A:2 X:AA Y:75 P:67 SP:FB CYC:1353
CBE1  85 01     STA $01 = FF                    A:FF X:AA Y:75 P:E5 SP:FB CYC:1355
CBE3  A9 44     LDA #$44                        A:FF X:AA Y:75 P:E5 SP:FB CYC:1358
CBE5  A2 55     LDX #$55                        A:44 X:AA Y:75 P:65 SP:FB CYC:1360
CBE7  A0 66     LDY #$66                        A:44 X:55 Y:75 P:65 SP:FB CYC:1362
CBE9  E8        INX                             A:44 X:55 Y:66 P:65 SP:FB CYC:1364
CBEA  88        DEY                             A:44 X:56 Y:66 P:65 SP:FB CYC:1366
CBEB  E0 56     CPX #$56                        A:44 X:56 Y:65 P:65 SP:FB CYC:1368
CBED  D0 21     BNE $CC10                       A:44 X:56 Y:65 P:67 SP:FB CYC:1370
CBEF  C0 65     CPY #$65                        A:44 X:56 Y:65 P:67 SP:FB CYC:1372
CBF1  D0 1D     BNE $CC10                       A:44 X:56 Y:65 P:67 SP:FB CYC:1374
CBF3  E8        INX                             A:44 X:56 Y:65 P:67 SP:FB CYC:1376
CBF4  E8        INX                             A:44 X:57 Y:65 P:65 SP:FB CYC:1378
CBF5  88        DEY                             A:44 X:58 Y:65 P:65 SP:FB CYC:1380
CBF6  88        DEY                             A:44 X:58 Y:64 P:65 SP:FB CYC:1382
CBF7  E0 58     CPX #$58                        A:44 X:58 Y:63 P:65 SP:FB CYC:1384
CBF9  D0 15     BNE $CC10                       A:44 X:58 Y:63 P:67 SP:FB CYC:1386
CBFB  C0 63     CPY #$63                        A:44 X:58 Y:63 P:67 SP:FB CYC:1388
CBFD  D0 11     BNE $CC10                       A:44 X:58 Y:63 P:67 SP:FB CYC:1390
CBFF  CA        DEX                             A:44 X:58 Y:63 P:67 SP:FB CYC:1392
CC00  C8        INY                             A:44 X:57 Y:63 P:65 SP:FB CYC:1394
CC01  E0 57     CPX #$57                        A:44 X:57 Y:64 P:65 SP:FB CYC:1396
CC03  D0 0B     BNE $CC10                       A:44 X:57 Y:64 P:67 SP:FB CYC:1398
CC05  C0 64     CPY #$64                        A:44 X:57 Y:64 P:67 SP:FB CYC:1400
CC07  D0 07     BNE $CC10                       A:44 X:57 Y:64 P:67 SP:FB CYC:1402
CC09  C9 44     CMP #$44                        A:44 X:57 Y:64 P:67 SP:FB CYC:1404
CC0B  D0 03     BNE $CC10                       A:44 X:57 Y:64 P:67 SP:FB CYC:1406
CC0D  4C 14 CC  JMP $CC14                       A:44 X:57 Y:64 P:67 SP:FB CYC:1408
CC14  EA        NOP                             A:44 X:57 Y:64 P:67 SP:FB CYC:1411
CC15  38        SEC                             A:44 X:57 Y:64 P:67 SP:FB CYC:1413
CC16  A2 69     LDX #$69                        A:44 X:57 Y:64 P:67 SP:FB CYC:1415
CC18  A9 96     LDA #$96                        A:44 X:69 Y:64 P:65 SP:FB CYC:1417
CC1A  24 01     BIT $01 = FF                    A:96 X:69 Y:64 P:E5 SP:FB CYC:1419
CC1C  A0 FF     LDY #$FF                        A:96 X:69 Y:64 P:E5 SP:FB CYC:1422
CC1E  C8        INY                             A:96 X:69 Y:FF P:E5 SP:FB CYC:1424
CC1F  D0 3D     BNE $CC5E                       A:96 X:69 Y:0 P:67 SP:FB CYC:1426
CC21  30 3B     BMI $CC5E                       A:96 X:69 Y:0 P:67 SP:FB CYC:1428
CC23  90 39     BCC $CC5E                       A:96 X:69 Y:0 P:67 SP:FB CYC:1430
CC25  50 37     BVC $CC5E                       A:96 X:69 Y:0 P:67 SP:FB CYC:1432
CC27  C0 00     CPY #$00                        A:96 X:69 Y:0 P:67 SP:FB CYC:1434
CC29  D0 33     BNE $CC5E                       A:96 X:69 Y:0 P:67 SP:FB CYC:1436
CC2B  C8        INY                             A:96 X:69 Y:0 P:67 SP:FB CYC:1438
CC2C  F0 30     BEQ $CC5E                       A:96 X:69 Y:1 P:65 SP:FB CYC:1440
CC2E  30 2E     BMI $CC5E                       A:96 X:69 Y:1 P:65 SP:FB CYC:1442
CC30  90 2C     BCC $CC5E                       A:96 X:69 Y:1 P:65 SP:FB CYC:1444
CC32  50 2A     BVC $CC5E                       A:96 X:69 Y:1 P:65 SP:FB CYC:1446
CC34  18        CLC                             A:96 X:69 Y:1 P:65 SP:FB CYC:1448
CC35  B8        CLV                             A:96 X:69 Y:1 P:64 SP:FB CYC:1450
CC36  A0 00     LDY #$00                        A:96 X:69 Y:1 P:24 SP:FB CYC:1452
CC38  88        DEY                             A:96 X:69 Y:0 P:26 SP:FB CYC:1454
CC39  F0 23     BEQ $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB CYC:1456
CC3B  10 21     BPL $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB CYC:1458
CC3D  B0 1F     BCS $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB CYC:1460
CC3F  70 1D     BVS $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB CYC:1462
CC41  C0 FF     CPY #$FF                        A:96 X:69 Y:FF P:A4 SP:FB CYC:1464
CC43  D0 19     BNE $CC5E                       A:96 X:69 Y:FF P:27 SP:FB CYC:1466
CC45  18        CLC                             A:96 X:69 Y:FF P:27 SP:FB CYC:1468
CC46  88        DEY                             A:96 X:69 Y:FF P:26 SP:FB CYC:1470
CC47  F0 15     BEQ $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB CYC:1472
CC49  10 13     BPL $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB CYC:1474
CC4B  B0 11     BCS $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB CYC:1476
CC4D  70 0F     BVS $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB CYC:1478
CC4F  C0 FE     CPY #$FE                        A:96 X:69 Y:FE P:A4 SP:FB CYC:1480
CC51  D0 0B     BNE $CC5E                       A:96 X:69 Y:FE P:27 SP:FB CYC:1482
CC53  C9 96     CMP #$96                        A:96 X:69 Y:FE P:27 SP:FB CYC:1484
CC55  D0 07     BNE $CC5E                       A:96 X:69 Y:FE P:27 SP:FB CYC:1486
CC57  E0 69     CPX #$69                        A:96 X:69 Y:FE P:27 SP:FB CYC:1488
CC59  D0 03     BNE $CC5E                       A:96 X:69 Y:FE P:27 SP:FB CYC:1490
CC5B  4C 62 CC  JMP $CC62                       A:96 X:69 Y:FE P:27 SP:FB CYC:1492
CC62  EA        NOP                             A:96 X:69 Y:FE P:27 SP:FB CYC:1495
CC63  38        SEC                             A:96 X:69 Y:FE P:27 SP:FB CYC:1497
CC64  A0 69     LDY #$69                        A:96 X:69 Y:FE P:27 SP:FB CYC:1499
CC66  A9 96     LDA #$96                        A:96 X:69 Y:69 P:25 SP:FB CYC:1501
CC68  24 01     BIT $01 = FF                    A:96 X:69 Y:69 P:A5 SP:FB CYC:1503
CC6A  A2 FF     LDX #$FF                        A:96 X:69 Y:69 P:E5 SP:FB CYC:1506
CC6C  E8        INX                             A:96 X:FF Y:69 P:E5 SP:FB CYC:1508
CC6D  D0 3D     BNE $CCAC                       A:96 X:0 Y:69 P:67 SP:FB CYC:1510
CC6F  30 3B     BMI $CCAC                       A:96 X:0 Y:69 P:67 SP:FB CYC:1512
CC71  90 39     BCC $CCAC                       A:96 X:0 Y:69 P:67 SP:FB CYC:1514
CC73  50 37     BVC $CCAC                       A:96 X:0 Y:69 P:67 SP:FB CYC:1516
CC75  E0 00     CPX #$00                        A:96 X:0 Y:69 P:67 SP:FB CYC:1518
CC77  D0 33     BNE $CCAC                       A:96 X:0 Y:69 P:67 SP:FB CYC:1520
CC79  E8        INX                             A:96 X:0 Y:69 P:67 SP:FB CYC:1522
CC7A  F0 30     BEQ $CCAC                       A:96 X:1 Y:69 P:65 SP:FB CYC:1524
CC7C  30 2E     BMI $CCAC                       A:96 X:1 Y:69 P:65 SP:FB CYC:1526
CC7E  90 2C     BCC $CCAC                       A:96 X:1 Y:69 P:65 SP:FB CYC:1528
CC80  50 2A     BVC $CCAC                       A:96 X:1 Y:69 P:65 SP:FB CYC:1530
CC82  18        CLC                             A:96 X:1 Y:69 P:65 SP:FB CYC:1532
CC83  B8        CLV                             A:96 X:1 Y:69 P:64 SP:FB CYC:1534
CC84  A2 00     LDX #$00                        A:96 X:1 Y:69 P:24 SP:FB CYC:1536
CC86  CA        DEX                             A:96 X:0 Y:69 P:26 SP:FB CYC:1538
CC87  F0 23     BEQ $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB CYC:1540
CC89  10 21     BPL $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB CYC:1542
CC8B  B0 1F     BCS $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB CYC:1544
CC8D  70 1D     BVS $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB CYC:1546
CC8F  E0 FF     CPX #$FF                        A:96 X:FF Y:69 P:A4 SP:FB CYC:1548
CC91  D0 19     BNE $CCAC                       A:96 X:FF Y:69 P:27 SP:FB CYC:1550
CC93  18        CLC                             A:96 X:FF Y:69 P:27 SP:FB CYC:1552
CC94  CA        DEX                             A:96 X:FF Y:69 P:26 SP:FB CYC:1554
CC95  F0 15     BEQ $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB CYC:1556
CC97  10 13     BPL $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB CYC:1558
CC99  B0 11     BCS $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB CYC:1560
CC9B  70 0F     BVS $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB CYC:1562
CC9D  E0 FE     CPX #$FE                        A:96 X:FE Y:69 P:A4 SP:FB CYC:1564
CC9F  D0 0B     BNE $CCAC                       A:96 X:FE Y:69 P:27 SP:FB CYC:1566
CCA1  C9 96     CMP #$96                        A:96 X:FE Y:69 P:27 SP:FB CYC:1568
CCA3  D0 07     BNE $CCAC                       A:96 X:FE Y:69 P:27 SP:FB CYC:1570
CCA5  C0 69     CPY #$69                        A:96 X:FE Y:69 P:27 SP:FB CYC:1572
CCA7  D0 03     BNE $CCAC                       A:96 X:FE Y:69 P:27 SP:FB CYC:1574
CCA9  4C B0 CC  JMP $CCB0                       A:96 X:FE Y:69 P:27 SP:FB CYC:1576
CCB0  EA        NOP                             A:96 X:FE Y:69 P:27 SP:FB CYC:1579
CCB1  A9 85     LDA #$85                        A:96 X:FE Y:69 P:27 SP:FB CYC:1581
CCB3  A2 34     LDX #$34                        A:85 X:FE Y:69 P:A5 SP:FB CYC:1583
CCB5  A0 99     LDY #$99                        A:85 X:34 Y:69 P:25 SP:FB CYC:1585
CCB7  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB CYC:1587
CCB8  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB CYC:1589
CCBA  A8        TAY                             A:85 X:34 Y:99 P:E4 SP:FB CYC:1592
CCBB  F0 2E     BEQ $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB CYC:1594
CCBD  B0 2C     BCS $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB CYC:1596
CCBF  50 2A     BVC $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB CYC:1598
CCC1  10 28     BPL $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB CYC:1600
CCC3  C9 85     CMP #$85                        A:85 X:34 Y:85 P:E4 SP:FB CYC:1602
CCC5  D0 24     BNE $CCEB                       A:85 X:34 Y:85 P:67 SP:FB CYC:1604
CCC7  E0 34     CPX #$34                        A:85 X:34 Y:85 P:67 SP:FB CYC:1606
CCC9  D0 20     BNE $CCEB                       A:85 X:34 Y:85 P:67 SP:FB CYC:1608
CCCB  C0 85     CPY #$85                        A:85 X:34 Y:85 P:67 SP:FB CYC:1610
CCCD  D0 1C     BNE $CCEB                       A:85 X:34 Y:85 P:67 SP:FB CYC:1612
CCCF  A9 00     LDA #$00                        A:85 X:34 Y:85 P:67 SP:FB CYC:1614
CCD1  38        SEC                             A:0 X:34 Y:85 P:67 SP:FB CYC:1616
CCD2  B8        CLV                             A:0 X:34 Y:85 P:67 SP:FB CYC:1618
CCD3  A8        TAY                             A:0 X:34 Y:85 P:27 SP:FB CYC:1620
CCD4  D0 15     BNE $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1622
CCD6  90 13     BCC $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1624
CCD8  70 11     BVS $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1626
CCDA  30 0F     BMI $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1628
CCDC  C9 00     CMP #$00                        A:0 X:34 Y:0 P:27 SP:FB CYC:1630
CCDE  D0 0B     BNE $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1632
CCE0  E0 34     CPX #$34                        A:0 X:34 Y:0 P:27 SP:FB CYC:1634
CCE2  D0 07     BNE $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1636
CCE4  C0 00     CPY #$00                        A:0 X:34 Y:0 P:27 SP:FB CYC:1638
CCE6  D0 03     BNE $CCEB                       A:0 X:34 Y:0 P:27 SP:FB CYC:1640
CCE8  4C EF CC  JMP $CCEF                       A:0 X:34 Y:0 P:27 SP:FB CYC:1642
CCEF  EA        NOP                             A:0 X:34 Y:0 P:27 SP:FB CYC:1645
CCF0  A9 85     LDA #$85                        A:0 X:34 Y:0 P:27 SP:FB CYC:1647
CCF2  A2 34     LDX #$34                        A:85 X:34 Y:0 P:A5 SP:FB CYC:1649
CCF4  A0 99     LDY #$99                        A:85 X:34 Y:0 P:25 SP:FB CYC:1651
CCF6  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB CYC:1653
CCF7  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB CYC:1655
CCF9  AA        TAX                             A:85 X:34 Y:99 P:E4 SP:FB CYC:1658
CCFA  F0 2E     BEQ $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB CYC:1660
CCFC  B0 2C     BCS $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB CYC:1662
CCFE  50 2A     BVC $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB CYC:1664
CD00  10 28     BPL $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB CYC:1666
CD02  C9 85     CMP #$85                        A:85 X:85 Y:99 P:E4 SP:FB CYC:1668
CD04  D0 24     BNE $CD2A                       A:85 X:85 Y:99 P:67 SP:FB CYC:1670
CD06  E0 85     CPX #$85                        A:85 X:85 Y:99 P:67 SP:FB CYC:1672
CD08  D0 20     BNE $CD2A                       A:85 X:85 Y:99 P:67 SP:FB CYC:1674
CD0A  C0 99     CPY #$99                        A:85 X:85 Y:99 P:67 SP:FB CYC:1676
CD0C  D0 1C     BNE $CD2A                       A:85 X:85 Y:99 P:67 SP:FB CYC:1678
CD0E  A9 00     LDA #$00                        A:85 X:85 Y:99 P:67 SP:FB CYC:1680
CD10  38        SEC                             A:0 X:85 Y:99 P:67 SP:FB CYC:1682
CD11  B8        CLV                             A:0 X:85 Y:99 P:67 SP:FB CYC:1684
CD12  AA        TAX                             A:0 X:85 Y:99 P:27 SP:FB CYC:1686
CD13  D0 15     BNE $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1688
CD15  90 13     BCC $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1690
CD17  70 11     BVS $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1692
CD19  30 0F     BMI $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1694
CD1B  C9 00     CMP #$00                        A:0 X:0 Y:99 P:27 SP:FB CYC:1696
CD1D  D0 0B     BNE $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1698
CD1F  E0 00     CPX #$00                        A:0 X:0 Y:99 P:27 SP:FB CYC:1700
CD21  D0 07     BNE $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1702
CD23  C0 99     CPY #$99                        A:0 X:0 Y:99 P:27 SP:FB CYC:1704
CD25  D0 03     BNE $CD2A                       A:0 X:0 Y:99 P:27 SP:FB CYC:1706
CD27  4C 2E CD  JMP $CD2E                       A:0 X:0 Y:99 P:27 SP:FB CYC:1708
CD2E  EA        NOP                             A:0 X:0 Y:99 P:27 SP:FB CYC:1711
CD2F  A9 85     LDA #$85                        A:0 X:0 Y:99 P:27 SP:FB CYC:1713
CD31  A2 34     LDX #$34                        A:85 X:0 Y:99 P:A5 SP:FB CYC:1715
CD33  A0 99     LDY #$99                        A:85 X:34 Y:99 P:25 SP:FB CYC:1717
CD35  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB CYC:1719
CD36  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB CYC:1721
CD38  98        TYA                             A:85 X:34 Y:99 P:E4 SP:FB CYC:1724
CD39  F0 2E     BEQ $CD69                       A:99 X:34 Y:99 P:E4 SP:FB CYC:1726
CD3B  B0 2C     BCS $CD69                       A:99 X:34 Y:99 P:E4 SP:FB CYC:1728
CD3D  50 2A     BVC $CD69                       A:99 X:34 Y:99 P:E4 SP:FB CYC:1730
CD3F  10 28     BPL $CD69                       A:99 X:34 Y:99 P:E4 SP:FB CYC:1732
CD41  C9 99     CMP #$99                        A:99 X:34 Y:99 P:E4 SP:FB CYC:1734
CD43  D0 24     BNE $CD69                       A:99 X:34 Y:99 P:67 SP:FB CYC:1736
CD45  E0 34     CPX #$34                        A:99 X:34 Y:99 P:67 SP:FB CYC:1738
CD47  D0 20     BNE $CD69                       A:99 X:34 Y:99 P:67 SP:FB CYC:1740
CD49  C0 99     CPY #$99                        A:99 X:34 Y:99 P:67 SP:FB CYC:1742
CD4B  D0 1C     BNE $CD69                       A:99 X:34 Y:99 P:67 SP:FB CYC:1744
CD4D  A0 00     LDY #$00                        A:99 X:34 Y:99 P:67 SP:FB CYC:1746
CD4F  38        SEC                             A:99 X:34 Y:0 P:67 SP:FB CYC:1748
CD50  B8        CLV                             A:99 X:34 Y:0 P:67 SP:FB CYC:1750
CD51  98        TYA                             A:99 X:34 Y:0 P:27 SP:FB CYC:1752
CD52  D0 15     BNE $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1754
CD54  90 13     BCC $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1756
CD56  70 11     BVS $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1758
CD58  30 0F     BMI $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1760
CD5A  C9 00     CMP #$00                        A:0 X:34 Y:0 P:27 SP:FB CYC:1762
CD5C  D0 0B     BNE $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1764
CD5E  E0 34     CPX #$34                        A:0 X:34 Y:0 P:27 SP:FB CYC:1766
CD60  D0 07     BNE $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1768
CD62  C0 00     CPY #$00                        A:0 X:34 Y:0 P:27 SP:FB CYC:1770
CD64  D0 03     BNE $CD69                       A:0 X:34 Y:0 P:27 SP:FB CYC:1772
CD66  4C 6D CD  JMP $CD6D                       A:0 X:34 Y:0 P:27 SP:FB CYC:1774
CD6D  EA        NOP                             A:0 X:34 Y:0 P:27 SP:FB CYC:1777
CD6E  A9 85     LDA #$85                        A:0 X:34 Y:0 P:27 SP:FB CYC:1779
CD70  A2 34     LDX #$34                        A:85 X:34 Y:0 P:A5 SP:FB CYC:1781
CD72  A0 99     LDY #$99                        A:85 X:34 Y:0 P:25 SP:FB CYC:1783
CD74  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB CYC:1785
CD75  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB CYC:1787
CD77  8A        TXA                             A:85 X:34 Y:99 P:E4 SP:FB CYC:1790
CD78  F0 2E     BEQ $CDA8                       A:34 X:34 Y:99 P:64 SP:FB CYC:1792
CD7A  B0 2C     BCS $CDA8                       A:34 X:34 Y:99 P:64 SP:FB CYC:1794
CD7C  50 2A     BVC $CDA8                       A:34 X:34 Y:99 P:64 SP:FB CYC:1796
CD7E  30 28     BMI $CDA8                       A:34 X:34 Y:99 P:64 SP:FB CYC:1798
CD80  C9 34     CMP #$34                        A:34 X:34 Y:99 P:64 SP:FB CYC:1800
CD82  D0 24     BNE $CDA8                       A:34 X:34 Y:99 P:67 SP:FB CYC:1802
CD84  E0 34     CPX #$34                        A:34 X:34 Y:99 P:67 SP:FB CYC:1804
CD86  D0 20     BNE $CDA8                       A:34 X:34 Y:99 P:67 SP:FB CYC:1806
CD88  C0 99     CPY #$99                        A:34 X:34 Y:99 P:67 SP:FB CYC:1808
CD8A  D0 1C     BNE $CDA8                       A:34 X:34 Y:99 P:67 SP:FB CYC:1810
CD8C  A2 00     LDX #$00                        A:34 X:34 Y:99 P:67 SP:FB CYC:1812
CD8E  38        SEC                             A:34 X:0 Y:99 P:67 SP:FB CYC:1814
CD8F  B8        CLV                             A:34 X:0 Y:99 P:67 SP:FB CYC:1816
CD90  8A        TXA                             A:34 X:0 Y:99 P:27 SP:FB CYC:1818
CD91  D0 15     BNE $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1820
CD93  90 13     BCC $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1822
CD95  70 11     BVS $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1824
CD97  30 0F     BMI $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1826
CD99  C9 00     CMP #$00                        A:0 X:0 Y:99 P:27 SP:FB CYC:1828
CD9B  D0 0B     BNE $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1830
CD9D  E0 00     CPX #$00                        A:0 X:0 Y:99 P:27 SP:FB CYC:1832
CD9F  D0 07     BNE $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1834
CDA1  C0 99     CPY #$99                        A:0 X:0 Y:99 P:27 SP:FB CYC:1836
CDA3  D0 03     BNE $CDA8                       A:0 X:0 Y:99 P:27 SP:FB CYC:1838
CDA5  4C AC CD  JMP $CDAC                       A:0 X:0 Y:99 P:27 SP:FB CYC:1840
CDAC  EA        NOP                             A:0 X:0 Y:99 P:27 SP:FB CYC:1843
CDAD  BA        TSX                             A:0 X:0 Y:99 P:27 SP:FB CYC:1845
CDAE  8E FF 07  STX $07FF = 00                  A:0 X:FB Y:99 P:A5 SP:FB CYC:1847
CDB1  A0 33     LDY #$33                        A:0 X:FB Y:99 P:A5 SP:FB CYC:1851
CDB3  A2 69     LDX #$69                        A:0 X:FB Y:33 P:25 SP:FB CYC:1853
CDB5  A9 84     LDA #$84                        A:0 X:69 Y:33 P:25 SP:FB CYC:1855
CDB7  18        CLC                             A:84 X:69 Y:33 P:A5 SP:FB CYC:1857
CDB8  24 01     BIT $01 = FF                    A:84 X:69 Y:33 P:A4 SP:FB CYC:1859
CDBA  9A        TXS                             A:84 X:69 Y:33 P:E4 SP:FB CYC:1862
CDBB  F0 32     BEQ $CDEF                       A:84 X:69 Y:33 P:E4 SP:69 CYC:1864
CDBD  10 30     BPL $CDEF                       A:84 X:69 Y:33 P:E4 SP:69 CYC:1866
CDBF  B0 2E     BCS $CDEF                       A:84 X:69 Y:33 P:E4 SP:69 CYC:1868
CDC1  50 2C     BVC $CDEF                       A:84 X:69 Y:33 P:E4 SP:69 CYC:1870
CDC3  C9 84     CMP #$84                        A:84 X:69 Y:33 P:E4 SP:69 CYC:1872
CDC5  D0 28     BNE $CDEF                       A:84 X:69 Y:33 P:67 SP:69 CYC:1874
CDC7  E0 69     CPX #$69                        A:84 X:69 Y:33 P:67 SP:69 CYC:1876
CDC9  D0 24     BNE $CDEF                       A:84 X:69 Y:33 P:67 SP:69 CYC:1878
CDCB  C0 33     CPY #$33                        A:84 X:69 Y:33 P:67 SP:69 CYC:1880
CDCD  D0 20     BNE $CDEF                       A:84 X:69 Y:33 P:67 SP:69 CYC:1882
CDCF  A0 01     LDY #$01                        A:84 X:69 Y:33 P:67 SP:69 CYC:1884
CDD1  A9 04     LDA #$04                        A:84 X:69 Y:1 P:65 SP:69 CYC:1886
CDD3  38        SEC                             A:4 X:69 Y:1 P:65 SP:69 CYC:1888
CDD4  B8        CLV                             A:4 X:69 Y:1 P:65 SP:69 CYC:1890
CDD5  A2 00     LDX #$00                        A:4 X:69 Y:1 P:25 SP:69 CYC:1892
CDD7  BA        TSX                             A:4 X:0 Y:1 P:27 SP:69 CYC:1894
CDD8  F0 15     BEQ $CDEF                       A:4 X:69 Y:1 P:25 SP:69 CYC:1896
CDDA  30 13     BMI $CDEF                       A:4 X:69 Y:1 P:25 SP:69 CYC:1898
CDDC  90 11     BCC $CDEF                       A:4 X:69 Y:1 P:25 SP:69 CYC:1900
CDDE  70 0F     BVS $CDEF                       A:4 X:69 Y:1 P:25 SP:69 CYC:1902
CDE0  E0 69     CPX #$69                        A:4 X:69 Y:1 P:25 SP:69 CYC:1904
CDE2  D0 0B     BNE $CDEF                       A:4 X:69 Y:1 P:27 SP:69 CYC:1906
CDE4  C9 04     CMP #$04                        A:4 X:69 Y:1 P:27 SP:69 CYC:1908
CDE6  D0 07     BNE $CDEF                       A:4 X:69 Y:1 P:27 SP:69 CYC:1910
CDE8  C0 01     CPY #$01                        A:4 X:69 Y:1 P:27 SP:69 CYC:1912
CDEA  D0 03     BNE $CDEF                       A:4 X:69 Y:1 P:27 SP:69 CYC:1914
CDEC  4C F3 CD  JMP $CDF3                       A:4 X:69 Y:1 P:27 SP:69 CYC:1916
CDF3  AE FF 07  LDX $07FF = FB                  A:4 X:69 Y:1 P:27 SP:69 CYC:1919
CDF6  9A        TXS                             A:4 X:FB Y:1 P:A5 SP:69 CYC:1923
CDF7  60        RTS                             A:4 X:FB Y:1 P:A5 SP:FB CYC:1925
C609  20 F8 CD  JSR $CDF8                       A:4 X:FB Y:1 P:A5 SP:FD CYC:1931
CDF8  A9 FF     LDA #$FF                        A:4 X:FB Y:1 P:A5 SP:FB CYC:1937
CDFA  85 01     STA $01 = FF                    A:FF X:FB Y:1 P:A5 SP:FB CYC:1939
CDFC  BA        TSX                             A:FF X:FB Y:1 P:A5 SP:FB CYC:1942
CDFD  8E FF 07  STX $07FF = FB                  A:FF X:FB Y:1 P:A5 SP:FB CYC:1944
CE00  EA        NOP                             A:FF X:FB Y:1 P:A5 SP:FB CYC:1948
CE01  A2 80     LDX #$80                        A:FF X:FB Y:1 P:A5 SP:FB CYC:1950
CE03  9A        TXS                             A:FF X:80 Y:1 P:A5 SP:FB CYC:1952
CE04  A9 33     LDA #$33                        A:FF X:80 Y:1 P:A5 SP:80 CYC:1954
CE06  48        PHA                             A:33 X:80 Y:1 P:25 SP:80 CYC:1956
CE07  A9 69     LDA #$69                        A:33 X:80 Y:1 P:25 SP:7F CYC:1959
CE09  48        PHA                             A:69 X:80 Y:1 P:25 SP:7F CYC:1961
CE0A  BA        TSX                             A:69 X:80 Y:1 P:25 SP:7E CYC:1964
CE0B  E0 7E     CPX #$7E                        A:69 X:7E Y:1 P:25 SP:7E CYC:1966
CE0D  D0 20     BNE $CE2F                       A:69 X:7E Y:1 P:27 SP:7E CYC:1968
CE0F  68        PLA                             A:69 X:7E Y:1 P:27 SP:7E CYC:1970
CE10  C9 69     CMP #$69                        A:69 X:7E Y:1 P:25 SP:7F CYC:1974
CE12  D0 1B     BNE $CE2F                       A:69 X:7E Y:1 P:27 SP:7F CYC:1976
CE14  68        PLA                             A:69 X:7E Y:1 P:27 SP:7F CYC:1978
CE15  C9 33     CMP #$33                        A:33 X:7E Y:1 P:25 SP:80 CYC:1982
CE17  D0 16     BNE $CE2F                       A:33 X:7E Y:1 P:27 SP:80 CYC:1984
CE19  BA        TSX                             A:33 X:7E Y:1 P:27 SP:80 CYC:1986
CE1A  E0 80     CPX #$80                        A:33 X:80 Y:1 P:A5 SP:80 CYC:1988
CE1C  D0 11     BNE $CE2F                       A:33 X:80 Y:1 P:27 SP:80 CYC:1990
CE1E  AD 80 01  LDA $0180 = 33                  A:33 X:80 Y:1 P:27 SP:80 CYC:1992
CE21  C9 33     CMP #$33                        A:33 X:80 Y:1 P:25 SP:80 CYC:1996
CE23  D0 0A     BNE $CE2F                       A:33 X:80 Y:1 P:27 SP:80 CYC:1998
CE25  AD 7F 01  LDA $017F = 69                  A:33 X:80 Y:1 P:27 SP:80 CYC:2000
CE28  C9 69     CMP #$69                        A:69 X:80 Y:1 P:25 SP:80 CYC:2004
CE2A  D0 03     BNE $CE2F                       A:69 X:80 Y:1 P:27 SP:80 CYC:2006
CE2C  4C 33 CE  JMP $CE33                       A:69 X:80 Y:1 P:27 SP:80 CYC:2008
CE33  EA        NOP                             A:69 X:80 Y:1 P:27 SP:80 CYC:2011
CE34  A2 80     LDX #$80                        A:69 X:80 Y:1 P:27 SP:80 CYC:2013
CE36  9A        TXS                             A:69 X:80 Y:1 P:A5 SP:80 CYC:2015
CE37  20 3D CE  JSR $CE3D                       A:69 X:80 Y:1 P:A5 SP:80 CYC:2017
CE3D  BA        TSX                             A:69 X:80 Y:1 P:A5 SP:7E CYC:2023
CE3E  E0 7E     CPX #$7E                        A:69 X:7E Y:1 P:25 SP:7E CYC:2025
CE40  D0 19     BNE $CE5B                       A:69 X:7E Y:1 P:27 SP:7E CYC:2027
CE42  68        PLA                             A:69 X:7E Y:1 P:27 SP:7E CYC:2029
CE43  68        PLA                             A:39 X:7E Y:1 P:25 SP:7F CYC:2033
CE44  BA        TSX                             A:CE X:7E Y:1 P:A5 SP:80 CYC:2037
CE45  E0 80     CPX #$80                        A:CE X:80 Y:1 P:A5 SP:80 CYC:2039
CE47  D0 12     BNE $CE5B                       A:CE X:80 Y:1 P:27 SP:80 CYC:2041
CE49  A9 00     LDA #$00                        A:CE X:80 Y:1 P:27 SP:80 CYC:2043
CE4B  20 4E CE  JSR $CE4E                       A:0 X:80 Y:1 P:27 SP:80 CYC:2045
CE4E  68        PLA                             A:0 X:80 Y:1 P:27 SP:7E CYC:2051
CE4F  C9 4D     CMP #$4D                        A:4D X:80 Y:1 P:25 SP:7F CYC:2055
CE51  D0 08     BNE $CE5B                       A:4D X:80 Y:1 P:27 SP:7F CYC:2057
CE53  68        PLA                             A:4D X:80 Y:1 P:27 SP:7F CYC:2059
CE54  C9 CE     CMP #$CE                        A:CE X:80 Y:1 P:A5 SP:80 CYC:2063
CE56  D0 03     BNE $CE5B                       A:CE X:80 Y:1 P:27 SP:80 CYC:2065
CE58  4C 5F CE  JMP $CE5F                       A:CE X:80 Y:1 P:27 SP:80 CYC:2067
CE5F  EA        NOP                             A:CE X:80 Y:1 P:27 SP:80 CYC:2070
CE60  A9 CE     LDA #$CE                        A:CE X:80 Y:1 P:27 SP:80 CYC:2072
CE62  48        PHA                             A:CE X:80 Y:1 P:A5 SP:80 CYC:2074
CE63  A9 66     LDA #$66                        A:CE X:80 Y:1 P:A5 SP:7F CYC:2077
CE65  48        PHA                             A:66 X:80 Y:1 P:25 SP:7F CYC:2079
CE66  60        RTS                             A:66 X:80 Y:1 P:25 SP:7E CYC:2082
CE67  A2 77     LDX #$77                        A:66 X:80 Y:1 P:25 SP:80 CYC:2088
CE69  A0 69     LDY #$69                        A:66 X:77 Y:1 P:25 SP:80 CYC:2090
CE6B  18        CLC                             A:66 X:77 Y:69 P:25 SP:80 CYC:2092
CE6C  24 01     BIT $01 = FF                    A:66 X:77 Y:69 P:24 SP:80 CYC:2094
CE6E  A9 83     LDA #$83                        A:66 X:77 Y:69 P:E4 SP:80 CYC:2097
CE70  20 66 CE  JSR $CE66                       A:83 X:77 Y:69 P:E4 SP:80 CYC:2099
CE66  60        RTS                             A:83 X:77 Y:69 P:E4 SP:7E CYC:2105
CE73  F0 24     BEQ $CE99                       A:83 X:77 Y:69 P:E4 SP:80 CYC:2111
CE75  10 22     BPL $CE99                       A:83 X:77 Y:69 P:E4 SP:80 CYC:2113
CE77  B0 20     BCS $CE99                       A:83 X:77 Y:69 P:E4 SP:80 CYC:2115
CE79  50 1E     BVC $CE99                       A:83 X:77 Y:69 P:E4 SP:80 CYC:2117
CE7B  C9 83     CMP #$83                        A:83 X:77 Y:69 P:E4 SP:80 CYC:2119
CE7D  D0 1A     BNE $CE99                       A:83 X:77 Y:69 P:67 SP:80 CYC:2121
CE7F  C0 69     CPY #$69                        A:83 X:77 Y:69 P:67 SP:80 CYC:2123
CE81  D0 16     BNE $CE99                       A:83 X:77 Y:69 P:67 SP:80 CYC:2125
CE83  E0 77     CPX #$77                        A:83 X:77 Y:69 P:67 SP:80 CYC:2127
CE85  D0 12     BNE $CE99                       A:83 X:77 Y:69 P:67 SP:80 CYC:2129
CE87  38        SEC                             A:83 X:77 Y:69 P:67 SP:80 CYC:2131
CE88  B8        CLV                             A:83 X:77 Y:69 P:67 SP:80 CYC:2133
CE89  A9 00     LDA #$00                        A:83 X:77 Y:69 P:27 SP:80 CYC:2135
CE8B  20 66 CE  JSR $CE66                       A:0 X:77 Y:69 P:27 SP:80 CYC:2137
CE66  60        RTS                             A:0 X:77 Y:69 P:27 SP:7E CYC:2143
CE8E  D0 09     BNE $CE99                       A:0 X:77 Y:69 P:27 SP:80 CYC:2149
CE90  30 07     BMI $CE99                       A:0 X:77 Y:69 P:27 SP:80 CYC:2151
CE92  90 05     BCC $CE99                       A:0 X:77 Y:69 P:27 SP:80 CYC:2153
CE94  70 03     BVS $CE99                       A:0 X:77 Y:69 P:27 SP:80 CYC:2155
CE96  4C 9D CE  JMP $CE9D                       A:0 X:77 Y:69 P:27 SP:80 CYC:2157
CE9D  EA        NOP                             A:0 X:77 Y:69 P:27 SP:80 CYC:2160
CE9E  A9 CE     LDA #$CE                        A:0 X:77 Y:69 P:27 SP:80 CYC:2162
CEA0  48        PHA                             A:CE X:77 Y:69 P:A5 SP:80 CYC:2164
CEA1  A9 AE     LDA #$AE                        A:CE X:77 Y:69 P:A5 SP:7F CYC:2167
CEA3  48        PHA                             A:AE X:77 Y:69 P:A5 SP:7F CYC:2169
CEA4  A9 65     LDA #$65                        A:AE X:77 Y:69 P:A5 SP:7E CYC:2172
CEA6  48        PHA                             A:65 X:77 Y:69 P:25 SP:7E CYC:2174
CEA7  A9 55     LDA #$55                        A:65 X:77 Y:69 P:25 SP:7D CYC:2177
CEA9  A0 88     LDY #$88                        A:55 X:77 Y:69 P:25 SP:7D CYC:2179
CEAB  A2 99     LDX #$99                        A:55 X:77 Y:88 P:A5 SP:7D CYC:2181
CEAD  40        RTI                             A:55 X:99 Y:88 P:A5 SP:7D CYC:2183
CEAE  30 35     BMI $CEE5                       A:55 X:99 Y:88 P:65 SP:80 CYC:2189
CEB0  50 33     BVC $CEE5                       A:55 X:99 Y:88 P:65 SP:80 CYC:2191
CEB2  F0 31     BEQ $CEE5                       A:55 X:99 Y:88 P:65 SP:80 CYC:2193
CEB4  90 2F     BCC $CEE5                       A:55 X:99 Y:88 P:65 SP:80 CYC:2195
CEB6  C9 55     CMP #$55                        A:55 X:99 Y:88 P:65 SP:80 CYC:2197
CEB8  D0 2B     BNE $CEE5                       A:55 X:99 Y:88 P:67 SP:80 CYC:2199
CEBA  C0 88     CPY #$88                        A:55 X:99 Y:88 P:67 SP:80 CYC:2201
CEBC  D0 27     BNE $CEE5                       A:55 X:99 Y:88 P:67 SP:80 CYC:2203
CEBE  E0 99     CPX #$99                        A:55 X:99 Y:88 P:67 SP:80 CYC:2205
CEC0  D0 23     BNE $CEE5                       A:55 X:99 Y:88 P:67 SP:80 CYC:2207
CEC2  A9 CE     LDA #$CE                        A:55 X:99 Y:88 P:67 SP:80 CYC:2209
CEC4  48        PHA                             A:CE X:99 Y:88 P:E5 SP:80 CYC:2211
CEC5  A9 CE     LDA #$CE                        A:CE X:99 Y:88 P:E5 SP:7F CYC:2214
CEC7  48        PHA                             A:CE X:99 Y:88 P:E5 SP:7F CYC:2216
CEC8  A9 87     LDA #$87                        A:CE X:99 Y:88 P:E5 SP:7E CYC:2219
CECA  48        PHA                             A:87 X:99 Y:88 P:E5 SP:7E CYC:2221
CECB  A9 55     LDA #$55                        A:87 X:99 Y:88 P:E5 SP:7D CYC:2224
CECD  40        RTI                             A:55 X:99 Y:88 P:65 SP:7D CYC:2226
CECE  10 15     BPL $CEE5                       A:55 X:99 Y:88 P:A7 SP:80 CYC:2232
CED0  70 13     BVS $CEE5                       A:55 X:99 Y:88 P:A7 SP:80 CYC:2234
CED2  D0 11     BNE $CEE5                       A:55 X:99 Y:88 P:A7 SP:80 CYC:2236
CED4  90 0F     BCC $CEE5                       A:55 X:99 Y:88 P:A7 SP:80 CYC:2238
CED6  C9 55     CMP #$55                        A:55 X:99 Y:88 P:A7 SP:80 CYC:2240
CED8  D0 0B     BNE $CEE5                       A:55 X:99 Y:88 P:27 SP:80 CYC:2242
CEDA  C0 88     CPY #$88                        A:55 X:99 Y:88 P:27 SP:80 CYC:2244
CEDC  D0 07     BNE $CEE5                       A:55 X:99 Y:88 P:27 SP:80 CYC:2246
CEDE  E0 99     CPX #$99                        A:55 X:99 Y:88 P:27 SP:80 CYC:2248
CEE0  D0 03     BNE $CEE5                       A:55 X:99 Y:88 P:27 SP:80 CYC:2250
CEE2  4C E9 CE  JMP $CEE9                       A:55 X:99 Y:88 P:27 SP:80 CYC:2252
CEE9  AE FF 07  LDX $07FF = FB                  A:55 X:99 Y:88 P:27 SP:80 CYC:2255
CEEC  9A        TXS                             A:55 X:FB Y:88 P:A5 SP:80 CYC:2259
CEED  60        RTS                             A:55 X:FB Y:88 P:A5 SP:FB CYC:2261
C60C  20 EE CE  JSR $CEEE                       A:55 X:FB Y:88 P:A5 SP:FD CYC:2267
CEEE  A2 55     LDX #$55                        A:55 X:FB Y:88 P:A5 SP:FB CYC:2273
CEF0  A0 69     LDY #$69                        A:55 X:55 Y:88 P:25 SP:FB CYC:2275
CEF2  A9 FF     LDA #$FF                        A:55 X:55 Y:69 P:25 SP:FB CYC:2277
CEF4  85 01     STA $01 = FF                    A:FF X:55 Y:69 P:A5 SP:FB CYC:2279
CEF6  EA        NOP                             A:FF X:55 Y:69 P:A5 SP:FB CYC:2282
CEF7  24 01     BIT $01 = FF                    A:FF X:55 Y:69 P:A5 SP:FB CYC:2284
CEF9  38        SEC                             A:FF X:55 Y:69 P:E5 SP:FB CYC:2287
CEFA  A9 01     LDA #$01                        A:FF X:55 Y:69 P:E5 SP:FB CYC:2289
CEFC  4A        LSR A                           A:1 X:55 Y:69 P:65 SP:FB CYC:2291
CEFD  90 1D     BCC $CF1C                       A:0 X:55 Y:69 P:67 SP:FB CYC:2293
CEFF  D0 1B     BNE $CF1C                       A:0 X:55 Y:69 P:67 SP:FB CYC:2295
CF01  30 19     BMI $CF1C                       A:0 X:55 Y:69 P:67 SP:FB CYC:2297
CF03  50 17     BVC $CF1C                       A:0 X:55 Y:69 P:67 SP:FB CYC:2299
CF05  C9 00     CMP #$00                        A:0 X:55 Y:69 P:67 SP:FB CYC:2301
CF07  D0 13     BNE $CF1C                       A:0 X:55 Y:69 P:67 SP:FB CYC:2303
CF09  B8        CLV                             A:0 X:55 Y:69 P:67 SP:FB CYC:2305
CF0A  A9 AA     LDA #$AA                        A:0 X:55 Y:69 P:27 SP:FB CYC:2307
CF0C  4A        LSR A                           A:AA X:55 Y:69 P:A5 SP:FB CYC:2309
CF0D  B0 0D     BCS $CF1C                       A:55 X:55 Y:69 P:24 SP:FB CYC:2311
CF0F  F0 0B     BEQ $CF1C                       A:55 X:55 Y:69 P:24 SP:FB CYC:2313
CF11  30 09     BMI $CF1C                       A:55 X:55 Y:69 P:24 SP:FB CYC:2315
CF13  70 07     BVS $CF1C                       A:55 X:55 Y:69 P:24 SP:FB CYC:2317
CF15  C9 55     CMP #$55                        A:55 X:55 Y:69 P:24 SP:FB CYC:2319
CF17  D0 03     BNE $CF1C                       A:55 X:55 Y:69 P:27 SP:FB CYC:2321
CF19  4C 20 CF  JMP $CF20                       A:55 X:55 Y:69 P:27 SP:FB CYC:2323
CF20  EA        NOP                             A:55 X:55 Y:69 P:27 SP:FB CYC:2326
CF21  24 01     BIT $01 = FF                    A:55 X:55 Y:69 P:27 SP:FB CYC:2328
CF23  38        SEC                             A:55 X:55 Y:69 P:E5 SP:FB CYC:2331
CF24  A9 80     LDA #$80                        A:55 X:55 Y:69 P:E5 SP:FB CYC:2333
CF26  0A        ASL A                           A:80 X:55 Y:69 P:E5 SP:FB CYC:2335
CF27  90 1E     BCC $CF47                       A:0 X:55 Y:69 P:67 SP:FB CYC:2337
CF29  D0 1C     BNE $CF47                       A:0 X:55 Y:69 P:67 SP:FB CYC:2339
CF2B  30 1A     BMI $CF47                       A:0 X:55 Y:69 P:67 SP:FB CYC:2341
CF2D  50 18     BVC $CF47                       A:0 X:55 Y:69 P:67 SP:FB CYC:2343
CF2F  C9 00     CMP #$00                        A:0 X:55 Y:69 P:67 SP:FB CYC:2345
CF31  D0 14     BNE $CF47                       A:0 X:55 Y:69 P:67 SP:FB CYC:2347
CF33  B8        CLV                             A:0 X:55 Y:69 P:67 SP:FB CYC:2349
CF34  38        SEC                             A:0 X:55 Y:69 P:27 SP:FB CYC:2351
CF35  A9 55     LDA #$55                        A:0 X:55 Y:69 P:27 SP:FB CYC:2353
CF37  0A        ASL A                           A:55 X:55 Y:69 P:25 SP:FB CYC:2355
CF38  B0 0D     BCS $CF47                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2357
CF3A  F0 0B     BEQ $CF47                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2359
CF3C  10 09     BPL $CF47                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2361
CF3E  70 07     BVS $CF47                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2363
CF40  C9 AA     CMP #$AA                        A:AA X:55 Y:69 P:A4 SP:FB CYC:2365
CF42  D0 03     BNE $CF47                       A:AA X:55 Y:69 P:27 SP:FB CYC:2367
CF44  4C 4B CF  JMP $CF4B                       A:AA X:55 Y:69 P:27 SP:FB CYC:2369
CF4B  EA        NOP                             A:AA X:55 Y:69 P:27 SP:FB CYC:2372
CF4C  24 01     BIT $01 = FF                    A:AA X:55 Y:69 P:27 SP:FB CYC:2374
CF4E  38        SEC                             A:AA X:55 Y:69 P:E5 SP:FB CYC:2377
CF4F  A9 01     LDA #$01                        A:AA X:55 Y:69 P:E5 SP:FB CYC:2379
CF51  6A        ROR A                           A:1 X:55 Y:69 P:65 SP:FB CYC:2381
CF52  90 1E     BCC $CF72                       A:80 X:55 Y:69 P:E5 SP:FB CYC:2383
CF54  F0 1C     BEQ $CF72                       A:80 X:55 Y:69 P:E5 SP:FB CYC:2385
CF56  10 1A     BPL $CF72                       A:80 X:55 Y:69 P:E5 SP:FB CYC:2387
CF58  50 18     BVC $CF72                       A:80 X:55 Y:69 P:E5 SP:FB CYC:2389
CF5A  C9 80     CMP #$80                        A:80 X:55 Y:69 P:E5 SP:FB CYC:2391
CF5C  D0 14     BNE $CF72                       A:80 X:55 Y:69 P:67 SP:FB CYC:2393
CF5E  B8        CLV                             A:80 X:55 Y:69 P:67 SP:FB CYC:2395
CF5F  18        CLC                             A:80 X:55 Y:69 P:27 SP:FB CYC:2397
CF60  A9 55     LDA #$55                        A:80 X:55 Y:69 P:26 SP:FB CYC:2399
CF62  6A        ROR A                           A:55 X:55 Y:69 P:24 SP:FB CYC:2401
CF63  90 0D     BCC $CF72                       A:2A X:55 Y:69 P:25 SP:FB CYC:2403
CF65  F0 0B     BEQ $CF72                       A:2A X:55 Y:69 P:25 SP:FB CYC:2405
CF67  30 09     BMI $CF72                       A:2A X:55 Y:69 P:25 SP:FB CYC:2407
CF69  70 07     BVS $CF72                       A:2A X:55 Y:69 P:25 SP:FB CYC:2409
CF6B  C9 2A     CMP #$2A                        A:2A X:55 Y:69 P:25 SP:FB CYC:2411
CF6D  D0 03     BNE $CF72                       A:2A X:55 Y:69 P:27 SP:FB CYC:2413
CF6F  4C 76 CF  JMP $CF76                       A:2A X:55 Y:69 P:27 SP:FB CYC:2415
CF76  EA        NOP                             A:2A X:55 Y:69 P:27 SP:FB CYC:2418
CF77  24 01     BIT $01 = FF                    A:2A X:55 Y:69 P:27 SP:FB CYC:2420
CF79  38        SEC                             A:2A X:55 Y:69 P:E5 SP:FB CYC:2423
CF7A  A9 80     LDA #$80                        A:2A X:55 Y:69 P:E5 SP:FB CYC:2425
CF7C  2A        ROL A                           A:80 X:55 Y:69 P:E5 SP:FB CYC:2427
CF7D  90 1E     BCC $CF9D                       A:1 X:55 Y:69 P:65 SP:FB CYC:2429
CF7F  F0 1C     BEQ $CF9D                       A:1 X:55 Y:69 P:65 SP:FB CYC:2431
CF81  30 1A     BMI $CF9D                       A:1 X:55 Y:69 P:65 SP:FB CYC:2433
CF83  50 18     BVC $CF9D                       A:1 X:55 Y:69 P:65 SP:FB CYC:2435
CF85  C9 01     CMP #$01                        A:1 X:55 Y:69 P:65 SP:FB CYC:2437
CF87  D0 14     BNE $CF9D                       A:1 X:55 Y:69 P:67 SP:FB CYC:2439
CF89  B8        CLV                             A:1 X:55 Y:69 P:67 SP:FB CYC:2441
CF8A  18        CLC                             A:1 X:55 Y:69 P:27 SP:FB CYC:2443
CF8B  A9 55     LDA #$55                        A:1 X:55 Y:69 P:26 SP:FB CYC:2445
CF8D  2A        ROL A                           A:55 X:55 Y:69 P:24 SP:FB CYC:2447
CF8E  B0 0D     BCS $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2449
CF90  F0 0B     BEQ $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2451
CF92  10 09     BPL $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2453
CF94  70 07     BVS $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB CYC:2455
CF96  C9 AA     CMP #$AA                        A:AA X:55 Y:69 P:A4 SP:FB CYC:2457
CF98  D0 03     BNE $CF9D                       A:AA X:55 Y:69 P:27 SP:FB CYC:2459
CF9A  4C A1 CF  JMP $CFA1                       A:AA X:55 Y:69 P:27 SP:FB CYC:2461
CFA1  60        RTS                             A:AA X:55 Y:69 P:27 SP:FB CYC:2464
C60F  20 A2 CF  JSR $CFA2                       A:AA X:55 Y:69 P:27 SP:FD CYC:2470
CFA2  A5 00     LDA $00 = 00                    A:AA X:55 Y:69 P:27 SP:FB CYC:2476
CFA4  8D FF 07  STA $07FF = FB                  A:0 X:55 Y:69 P:27 SP:FB CYC:2479
CFA7  A9 00     LDA #$00                        A:0 X:55 Y:69 P:27 SP:FB CYC:2483
CFA9  85 80     STA $80 = 00                    A:0 X:55 Y:69 P:27 SP:FB CYC:2485
CFAB  A9 02     LDA #$02                        A:0 X:55 Y:69 P:27 SP:FB CYC:2488
CFAD  85 81     STA $81 = 00                    A:2 X:55 Y:69 P:25 SP:FB CYC:2490
CFAF  A9 FF     LDA #$FF                        A:2 X:55 Y:69 P:25 SP:FB CYC:2493
CFB1  85 01     STA $01 = FF                    A:FF X:55 Y:69 P:A5 SP:FB CYC:2495
CFB3  A9 00     LDA #$00                        A:FF X:55 Y:69 P:A5 SP:FB CYC:2498
CFB5  85 82     STA $82 = 00                    A:0 X:55 Y:69 P:27 SP:FB CYC:2500
CFB7  A9 03     LDA #$03                        A:0 X:55 Y:69 P:27 SP:FB CYC:2503
CFB9  85 83     STA $83 = 00                    A:3 X:55 Y:69 P:25 SP:FB CYC:2505
CFBB  85 84     STA $84 = 00                    A:3 X:55 Y:69 P:25 SP:FB CYC:2508
CFBD  A9 00     LDA #$00                        A:3 X:55 Y:69 P:25 SP:FB CYC:2511
CFBF  85 FF     STA $FF = 00                    A:0 X:55 Y:69 P:27 SP:FB CYC:2513
CFC1  A9 04     LDA #$04                        A:0 X:55 Y:69 P:27 SP:FB CYC:2516
CFC3  85 00     STA $00 = 00                    A:4 X:55 Y:69 P:25 SP:FB CYC:2518
CFC5  A9 5A     LDA #$5A                        A:4 X:55 Y:69 P:25 SP:FB CYC:2521
CFC7  8D 00 02  STA $0200 = 00                  A:5A X:55 Y:69 P:25 SP:FB CYC:2523
CFCA  A9 5B     LDA #$5B                        A:5A X:55 Y:69 P:25 SP:FB CYC:2527
CFCC  8D 00 03  STA $0300 = 00                  A:5B X:55 Y:69 P:25 SP:FB CYC:2529
CFCF  A9 5C     LDA #$5C                        A:5B X:55 Y:69 P:25 SP:FB CYC:2533
CFD1  8D 03 03  STA $0303 = 00                  A:5C X:55 Y:69 P:25 SP:FB CYC:2535
CFD4  A9 5D     LDA #$5D                        A:5C X:55 Y:69 P:25 SP:FB CYC:2539
CFD6  8D 00 04  STA $0400 = 00                  A:5D X:55 Y:69 P:25 SP:FB CYC:2541
CFD9  A2 00     LDX #$00                        A:5D X:55 Y:69 P:25 SP:FB CYC:2545
CFDB  A1 80     LDA ($80,X) @ 80 = 0200 = 5A    A:5D X:0 Y:69 P:27 SP:FB CYC:2547
CFDD  C9 5A     CMP #$5A                        A:5A X:0 Y:69 P:25 SP:FB CYC:2553
CFDF  D0 1F     BNE $D000                       A:5A X:0 Y:69 P:27 SP:FB CYC:2555
CFE1  E8        INX                             A:5A X:0 Y:69 P:27 SP:FB CYC:2557
CFE2  E8        INX                             A:5A X:1 Y:69 P:25 SP:FB CYC:2559
CFE3  A1 80     LDA ($80,X) @ 82 = 0300 = 5B    A:5A X:2 Y:69 P:25 SP:FB CYC:2561
CFE5  C9 5B     CMP #$5B                        A:5B X:2 Y:69 P:25 SP:FB CYC:2567
CFE7  D0 17     BNE $D000                       A:5B X:2 Y:69 P:27 SP:FB CYC:2569
CFE9  E8        INX                             A:5B X:2 Y:69 P:27 SP:FB CYC:2571
CFEA  A1 80     LDA ($80,X) @ 83 = 0303 = 5C    A:5B X:3 Y:69 P:25 SP:FB CYC:2573
CFEC  C9 5C     CMP #$5C                        A:5C X:3 Y:69 P:25 SP:FB CYC:2579
CFEE  D0 10     BNE $D000                       A:5C X:3 Y:69 P:27 SP:FB CYC:2581
CFF0  A2 00     LDX #$00                        A:5C X:3 Y:69 P:27 SP:FB CYC:2583
CFF2  A1 FF     LDA ($FF,X) @ FF = 0400 = 5D    A:5C X:0 Y:69 P:27 SP:FB CYC:2585
CFF4  C9 5D     CMP #$5D                        A:5D X:0 Y:69 P:25 SP:FB CYC:2591
CFF6  D0 08     BNE $D000                       A:5D X:0 Y:69 P:27 SP:FB CYC:2593
CFF8  A2 81     LDX #$81                        A:5D X:0 Y:69 P:27 SP:FB CYC:2595
CFFA  A1 FF     LDA ($FF,X) @ 80 = 0200 = 5A    A:5D X:81 Y:69 P:A5 SP:FB CYC:2597
CFFC  C9 5A     CMP #$5A                        A:5A X:81 Y:69 P:25 SP:FB CYC:2603
CFFE  F0 05     BEQ $D005                       A:5A X:81 Y:69 P:27 SP:FB CYC:2605
D005  A9 AA     LDA #$AA                        A:5A X:81 Y:69 P:27 SP:FB CYC:2608
D007  A2 00     LDX #$00                        A:AA X:81 Y:69 P:A5 SP:FB CYC:2610
D009  81 80     STA ($80,X) @ 80 = 0200 = 5A    A:AA X:0 Y:69 P:27 SP:FB CYC:2612
D00B  E8        INX                             A:AA X:0 Y:69 P:27 SP:FB CYC:2618
D00C  E8        INX                             A:AA X:1 Y:69 P:25 SP:FB CYC:2620
D00D  A9 AB     LDA #$AB                        A:AA X:2 Y:69 P:25 SP:FB CYC:2622
D00F  81 80     STA ($80,X) @ 82 = 0300 = 5B    A:AB X:2 Y:69 P:A5 SP:FB CYC:2624
D011  E8        INX                             A:AB X:2 Y:69 P:A5 SP:FB CYC:2630
D012  A9 AC     LDA #$AC                        A:AB X:3 Y:69 P:25 SP:FB CYC:2632
D014  81 80     STA ($80,X) @ 83 = 0303 = 5C    A:AC X:3 Y:69 P:A5 SP:FB CYC:2634
D016  A2 00     LDX #$00                        A:AC X:3 Y:69 P:A5 SP:FB CYC:2640
D018  A9 AD     LDA #$AD                        A:AC X:0 Y:69 P:27 SP:FB CYC:2642
D01A  81 FF     STA ($FF,X) @ FF = 0400 = 5D    A:AD X:0 Y:69 P:A5 SP:FB CYC:2644
D01C  AD 00 02  LDA $0200 = AA                  A:AD X:0 Y:69 P:A5 SP:FB CYC:2650
D01F  C9 AA     CMP #$AA                        A:AA X:0 Y:69 P:A5 SP:FB CYC:2654
D021  D0 15     BNE $D038                       A:AA X:0 Y:69 P:27 SP:FB CYC:2656
D023  AD 00 03  LDA $0300 = AB                  A:AA X:0 Y:69 P:27 SP:FB CYC:2658
D026  C9 AB     CMP #$AB                        A:AB X:0 Y:69 P:A5 SP:FB CYC:2662
D028  D0 0E     BNE $D038                       A:AB X:0 Y:69 P:27 SP:FB CYC:2664
D02A  AD 03 03  LDA $0303 = AC                  A:AB X:0 Y:69 P:27 SP:FB CYC:2666
D02D  C9 AC     CMP #$AC                        A:AC X:0 Y:69 P:A5 SP:FB CYC:2670
D02F  D0 07     BNE $D038                       A:AC X:0 Y:69 P:27 SP:FB CYC:2672
D031  AD 00 04  LDA $0400 = AD                  A:AC X:0 Y:69 P:27 SP:FB CYC:2674
D034  C9 AD     CMP #$AD                        A:AD X:0 Y:69 P:A5 SP:FB CYC:2678
D036  F0 05     BEQ $D03D                       A:AD X:0 Y:69 P:27 SP:FB CYC:2680
D03D  AD FF 07  LDA $07FF = 00                  A:AD X:0 Y:69 P:27 SP:FB CYC:2683
D040  85 00     STA $00 = 04                    A:0 X:0 Y:69 P:27 SP:FB CYC:2687
D042  A9 00     LDA #$00                        A:0 X:0 Y:69 P:27 SP:FB CYC:2690
D044  8D 00 03  STA $0300 = AB                  A:0 X:0 Y:69 P:27 SP:FB CYC:2692
D047  A9 AA     LDA #$AA                        A:0 X:0 Y:69 P:27 SP:FB CYC:2696
D049  8D 00 02  STA $0200 = AA                  A:AA X:0 Y:69 P:A5 SP:FB CYC:2698
D04C  A2 00     LDX #$00                        A:AA X:0 Y:69 P:A5 SP:FB CYC:2702
D04E  A0 5A     LDY #$5A                        A:AA X:0 Y:69 P:27 SP:FB CYC:2704
D050  20 B6 F7  JSR $F7B6                       A:AA X:0 Y:5A P:25 SP:FB CYC:2706
F7B6  18        CLC                             A:AA X:0 Y:5A P:25 SP:F9 CYC:2712
F7B7  A9 FF     LDA #$FF                        A:AA X:0 Y:5A P:24 SP:F9 CYC:2714
F7B9  85 01     STA $01 = FF                    A:FF X:0 Y:5A P:A4 SP:F9 CYC:2716
F7BB  24 01     BIT $01 = FF                    A:FF X:0 Y:5A P:A4 SP:F9 CYC:2719
F7BD  A9 55     LDA #$55                        A:FF X:0 Y:5A P:E4 SP:F9 CYC:2722
F7BF  60        RTS                             A:55 X:0 Y:5A P:64 SP:F9 CYC:2724
D053  01 80     ORA ($80,X) @ 80 = 0200 = AA    A:55 X:0 Y:5A P:64 SP:FB CYC:2730
D055  20 C0 F7  JSR $F7C0                       A:FF X:0 Y:5A P:E4 SP:FB CYC:2736
F7C0  B0 09     BCS $F7CB                       A:FF X:0 Y:5A P:E4 SP:F9 CYC:2742
F7C2  10 07     BPL $F7CB                       A:FF X:0 Y:5A P:E4 SP:F9 CYC:2744
F7C4  C9 FF     CMP #$FF                        A:FF X:0 Y:5A P:E4 SP:F9 CYC:2746
F7C6  D0 03     BNE $F7CB                       A:FF X:0 Y:5A P:67 SP:F9 CYC:2748
F7C8  50 01     BVC $F7CB                       A:FF X:0 Y:5A P:67 SP:F9 CYC:2750
F7CA  60        RTS                             A:FF X:0 Y:5A P:67 SP:F9 CYC:2752
D058  C8        INY                             A:FF X:0 Y:5A P:67 SP:FB CYC:2758
D059  20 CE F7  JSR $F7CE                       A:FF X:0 Y:5B P:65 SP:FB CYC:2760
F7CE  38        SEC                             A:FF X:0 Y:5B P:65 SP:F9 CYC:2766
F7CF  B8        CLV                             A:FF X:0 Y:5B P:65 SP:F9 CYC:2768
F7D0  A9 00     LDA #$00                        A:FF X:0 Y:5B P:25 SP:F9 CYC:2770
F7D2  60        RTS                             A:0 X:0 Y:5B P:27 SP:F9 CYC:2772
D05C  01 82     ORA ($82,X) @ 82 = 0300 = 00    A:0 X:0 Y:5B P:27 SP:FB CYC:2778
D05E  20 D3 F7  JSR $F7D3                       A:0 X:0 Y:5B P:27 SP:FB CYC:2784
F7D3  D0 07     BNE $F7DC                       A:0 X:0 Y:5B P:27 SP:F9 CYC:2790
F7D5  70 05     BVS $F7DC                       A:0 X:0 Y:5B P:27 SP:F9 CYC:2792
F7D7  90 03     BCC $F7DC                       A:0 X:0 Y:5B P:27 SP:F9 CYC:2794
F7D9  30 01     BMI $F7DC                       A:0 X:0 Y:5B P:27 SP:F9 CYC:2796
F7DB  60        RTS                             A:0 X:0 Y:5B P:27 SP:F9 CYC:2798
D061  C8        INY                             A:0 X:0 Y:5B P:27 SP:FB CYC:2804
D062  20 DF F7  JSR $F7DF                       A:0 X:0 Y:5C P:25 SP:FB CYC:2806
F7DF  18        CLC                             A:0 X:0 Y:5C P:25 SP:F9 CYC:2812
F7E0  24 01     BIT $01 = FF                    A:0 X:0 Y:5C P:24 SP:F9 CYC:2814
F7E2  A9 55     LDA #$55                        A:0 X:0 Y:5C P:E6 SP:F9 CYC:2817
F7E4  60        RTS                             A:55 X:0 Y:5C P:64 SP:F9 CYC:2819
D065  21 80     AND ($80,X) @ 80 = 0200 = AA    A:55 X:0 Y:5C P:64 SP:FB CYC:2825
D067  20 E5 F7  JSR $F7E5                       A:0 X:0 Y:5C P:66 SP:FB CYC:2831
F7E5  D0 07     BNE $F7EE                       A:0 X:0 Y:5C P:66 SP:F9 CYC:2837
F7E7  50 05     BVC $F7EE                       A:0 X:0 Y:5C P:66 SP:F9 CYC:2839
F7E9  B0 03     BCS $F7EE                       A:0 X:0 Y:5C P:66 SP:F9 CYC:2841
F7EB  30 01     BMI $F7EE                       A:0 X:0 Y:5C P:66 SP:F9 CYC:2843
F7ED  60        RTS                             A:0 X:0 Y:5C P:66 SP:F9 CYC:2845
D06A  C8        INY                             A:0 X:0 Y:5C P:66 SP:FB CYC:2851
D06B  A9 EF     LDA #$EF                        A:0 X:0 Y:5D P:64 SP:FB CYC:2853
D06D  8D 00 03  STA $0300 = 00                  A:EF X:0 Y:5D P:E4 SP:FB CYC:2855
D070  20 F1 F7  JSR $F7F1                       A:EF X:0 Y:5D P:E4 SP:FB CYC:2859
F7F1  38        SEC                             A:EF X:0 Y:5D P:E4 SP:F9 CYC:2865
F7F2  B8        CLV                             A:EF X:0 Y:5D P:E5 SP:F9 CYC:2867
F7F3  A9 F8     LDA #$F8                        A:EF X:0 Y:5D P:A5 SP:F9 CYC:2869
F7F5  60        RTS                             A:F8 X:0 Y:5D P:A5 SP:F9 CYC:2871
D073  21 82     AND ($82,X) @ 82 = 0300 = EF    A:F8 X:0 Y:5D P:A5 SP:FB CYC:2877
D075  20 F6 F7  JSR $F7F6                       A:E8 X:0 Y:5D P:A5 SP:FB CYC:2883
F7F6  90 09     BCC $F801                       A:E8 X:0 Y:5D P:A5 SP:F9 CYC:2889
F7F8  10 07     BPL $F801                       A:E8 X:0 Y:5D P:A5 SP:F9 CYC:2891
F7FA  C9 E8     CMP #$E8                        A:E8 X:0 Y:5D P:A5 SP:F9 CYC:2893
F7FC  D0 03     BNE $F801                       A:E8 X:0 Y:5D P:27 SP:F9 CYC:2895
F7FE  70 01     BVS $F801                       A:E8 X:0 Y:5D P:27 SP:F9 CYC:2897
F800  60        RTS                             A:E8 X:0 Y:5D P:27 SP:F9 CYC:2899
D078  C8        INY                             A:E8 X:0 Y:5D P:27 SP:FB CYC:2905
D079  20 04 F8  JSR $F804                       A:E8 X:0 Y:5E P:25 SP:FB CYC:2907
F804  18        CLC                             A:E8 X:0 Y:5E P:25 SP:F9 CYC:2913
F805  24 01     BIT $01 = FF                    A:E8 X:0 Y:5E P:24 SP:F9 CYC:2915
F807  A9 5F     LDA #$5F                        A:E8 X:0 Y:5E P:E4 SP:F9 CYC:2918
F809  60        RTS                             A:5F X:0 Y:5E P:64 SP:F9 CYC:2920
D07C  41 80     EOR ($80,X) @ 80 = 0200 = AA    A:5F X:0 Y:5E P:64 SP:FB CYC:2926
D07E  20 0A F8  JSR $F80A                       A:F5 X:0 Y:5E P:E4 SP:FB CYC:2932
F80A  B0 09     BCS $F815                       A:F5 X:0 Y:5E P:E4 SP:F9 CYC:2938
F80C  10 07     BPL $F815                       A:F5 X:0 Y:5E P:E4 SP:F9 CYC:2940
F80E  C9 F5     CMP #$F5                        A:F5 X:0 Y:5E P:E4 SP:F9 CYC:2942
F810  D0 03     BNE $F815                       A:F5 X:0 Y:5E P:67 SP:F9 CYC:2944
F812  50 01     BVC $F815                       A:F5 X:0 Y:5E P:67 SP:F9 CYC:2946
F814  60        RTS                             A:F5 X:0 Y:5E P:67 SP:F9 CYC:2948
D081  C8        INY                             A:F5 X:0 Y:5E P:67 SP:FB CYC:2954
D082  A9 70     LDA #$70                        A:F5 X:0 Y:5F P:65 SP:FB CYC:2956
D084  8D 00 03  STA $0300 = EF                  A:70 X:0 Y:5F P:65 SP:FB CYC:2958
D087  20 18 F8  JSR $F818                       A:70 X:0 Y:5F P:65 SP:FB CYC:2962
F818  38        SEC                             A:70 X:0 Y:5F P:65 SP:F9 CYC:2968
F819  B8        CLV                             A:70 X:0 Y:5F P:65 SP:F9 CYC:2970
F81A  A9 70     LDA #$70                        A:70 X:0 Y:5F P:25 SP:F9 CYC:2972
F81C  60        RTS                             A:70 X:0 Y:5F P:25 SP:F9 CYC:2974
D08A  41 82     EOR ($82,X) @ 82 = 0300 = 70    A:70 X:0 Y:5F P:25 SP:FB CYC:2980
D08C  20 1D F8  JSR $F81D                       A:0 X:0 Y:5F P:27 SP:FB CYC:2986
F81D  D0 07     BNE $F826                       A:0 X:0 Y:5F P:27 SP:F9 CYC:2992
F81F  70 05     BVS $F826                       A:0 X:0 Y:5F P:27 SP:F9 CYC:2994
F821  90 03     BCC $F826                       A:0 X:0 Y:5F P:27 SP:F9 CYC:2996
F823  30 01     BMI $F826                       A:0 X:0 Y:5F P:27 SP:F9 CYC:2998
F825  60        RTS                             A:0 X:0 Y:5F P:27 SP:F9 CYC:3000
D08F  C8        INY                             A:0 X:0 Y:5F P:27 SP:FB CYC:3006
D090  A9 69     LDA #$69                        A:0 X:0 Y:60 P:25 SP:FB CYC:3008
D092  8D 00 02  STA $0200 = AA                  A:69 X:0 Y:60 P:25 SP:FB CYC:3010
D095  20 29 F8  JSR $F829                       A:69 X:0 Y:60 P:25 SP:FB CYC:3014
F829  18        CLC                             A:69 X:0 Y:60 P:25 SP:F9 CYC:3020
F82A  24 01     BIT $01 = FF                    A:69 X:0 Y:60 P:24 SP:F9 CYC:3022
F82C  A9 00     LDA #$00                        A:69 X:0 Y:60 P:E4 SP:F9 CYC:3025
F82E  60        RTS                             A:0 X:0 Y:60 P:66 SP:F9 CYC:3027
D098  61 80     ADC ($80,X) @ 80 = 0200 = 69    A:0 X:0 Y:60 P:66 SP:FB CYC:3033
D09A  20 2F F8  JSR $F82F                       A:69 X:0 Y:60 P:24 SP:FB CYC:3039
F82F  30 09     BMI $F83A                       A:69 X:0 Y:60 P:24 SP:F9 CYC:3045
F831  B0 07     BCS $F83A                       A:69 X:0 Y:60 P:24 SP:F9 CYC:3047
F833  C9 69     CMP #$69                        A:69 X:0 Y:60 P:24 SP:F9 CYC:3049
F835  D0 03     BNE $F83A                       A:69 X:0 Y:60 P:27 SP:F9 CYC:3051
F837  70 01     BVS $F83A                       A:69 X:0 Y:60 P:27 SP:F9 CYC:3053
F839  60        RTS                             A:69 X:0 Y:60 P:27 SP:F9 CYC:3055
D09D  C8        INY                             A:69 X:0 Y:60 P:27 SP:FB CYC:3061
D09E  20 3D F8  JSR $F83D                       A:69 X:0 Y:61 P:25 SP:FB CYC:3063
F83D  38        SEC                             A:69 X:0 Y:61 P:25 SP:F9 CYC:3069
F83E  24 01     BIT $01 = FF                    A:69 X:0 Y:61 P:25 SP:F9 CYC:3071
F840  A9 00     LDA #$00                        A:69 X:0 Y:61 P:E5 SP:F9 CYC:3074
F842  60        RTS                             A:0 X:0 Y:61 P:67 SP:F9 CYC:3076
D0A1  61 80     ADC ($80,X) @ 80 = 0200 = 69    A:0 X:0 Y:61 P:67 SP:FB CYC:3082
D0A3  20 43 F8  JSR $F843                       A:6A X:0 Y:61 P:24 SP:FB CYC:3088
F843  30 09     BMI $F84E                       A:6A X:0 Y:61 P:24 SP:F9 CYC:3094
F845  B0 07     BCS $F84E                       A:6A X:0 Y:61 P:24 SP:F9 CYC:3096
F847  C9 6A     CMP #$6A                        A:6A X:0 Y:61 P:24 SP:F9 CYC:3098
F849  D0 03     BNE $F84E                       A:6A X:0 Y:61 P:27 SP:F9 CYC:3100
F84B  70 01     BVS $F84E                       A:6A X:0 Y:61 P:27 SP:F9 CYC:3102
F84D  60        RTS                             A:6A X:0 Y:61 P:27 SP:F9 CYC:3104
D0A6  C8        INY                             A:6A X:0 Y:61 P:27 SP:FB CYC:3110
D0A7  A9 7F     LDA #$7F                        A:6A X:0 Y:62 P:25 SP:FB CYC:3112
D0A9  8D 00 02  STA $0200 = 69                  A:7F X:0 Y:62 P:25 SP:FB CYC:3114
D0AC  20 51 F8  JSR $F851                       A:7F X:0 Y:62 P:25 SP:FB CYC:3118
F851  38        SEC                             A:7F X:0 Y:62 P:25 SP:F9 CYC:3124
F852  B8        CLV                             A:7F X:0 Y:62 P:25 SP:F9 CYC:3126
F853  A9 7F     LDA #$7F                        A:7F X:0 Y:62 P:25 SP:F9 CYC:3128
F855  60        RTS                             A:7F X:0 Y:62 P:25 SP:F9 CYC:3130
D0AF  61 80     ADC ($80,X) @ 80 = 0200 = 7F    A:7F X:0 Y:62 P:25 SP:FB CYC:3136
D0B1  20 56 F8  JSR $F856                       A:FF X:0 Y:62 P:E4 SP:FB CYC:3142
F856  10 09     BPL $F861                       A:FF X:0 Y:62 P:E4 SP:F9 CYC:3148
F858  B0 07     BCS $F861                       A:FF X:0 Y:62 P:E4 SP:F9 CYC:3150
F85A  C9 FF     CMP #$FF                        A:FF X:0 Y:62 P:E4 SP:F9 CYC:3152
F85C  D0 03     BNE $F861                       A:FF X:0 Y:62 P:67 SP:F9 CYC:3154
F85E  50 01     BVC $F861                       A:FF X:0 Y:62 P:67 SP:F9 CYC:3156
F860  60        RTS                             A:FF X:0 Y:62 P:67 SP:F9 CYC:3158
D0B4  C8        INY                             A:FF X:0 Y:62 P:67 SP:FB CYC:3164
D0B5  A9 80     LDA #$80                        A:FF X:0 Y:63 P:65 SP:FB CYC:3166
D0B7  8D 00 02  STA $0200 = 7F                  A:80 X:0 Y:63 P:E5 SP:FB CYC:3168
D0BA  20 64 F8  JSR $F864                       A:80 X:0 Y:63 P:E5 SP:FB CYC:3172
F864  18        CLC                             A:80 X:0 Y:63 P:E5 SP:F9 CYC:3178
F865  24 01     BIT $01 = FF                    A:80 X:0 Y:63 P:E4 SP:F9 CYC:3180
F867  A9 7F     LDA #$7F                        A:80 X:0 Y:63 P:E4 SP:F9 CYC:3183
F869  60        RTS                             A:7F X:0 Y:63 P:64 SP:F9 CYC:3185
D0BD  61 80     ADC ($80,X) @ 80 = 0200 = 80    A:7F X:0 Y:63 P:64 SP:FB CYC:3191
D0BF  20 6A F8  JSR $F86A                       A:FF X:0 Y:63 P:A4 SP:FB CYC:3197
F86A  10 09     BPL $F875                       A:FF X:0 Y:63 P:A4 SP:F9 CYC:3203
F86C  B0 07     BCS $F875                       A:FF X:0 Y:63 P:A4 SP:F9 CYC:3205
F86E  C9 FF     CMP #$FF                        A:FF X:0 Y:63 P:A4 SP:F9 CYC:3207
F870  D0 03     BNE $F875                       A:FF X:0 Y:63 P:27 SP:F9 CYC:3209
F872  70 01     BVS $F875                       A:FF X:0 Y:63 P:27 SP:F9 CYC:3211
F874  60        RTS                             A:FF X:0 Y:63 P:27 SP:F9 CYC:3213
D0C2  C8        INY                             A:FF X:0 Y:63 P:27 SP:FB CYC:3219
D0C3  20 78 F8  JSR $F878                       A:FF X:0 Y:64 P:25 SP:FB CYC:3221
F878  38        SEC                             A:FF X:0 Y:64 P:25 SP:F9 CYC:3227
F879  B8        CLV                             A:FF X:0 Y:64 P:25 SP:F9 CYC:3229
F87A  A9 7F     LDA #$7F                        A:FF X:0 Y:64 P:25 SP:F9 CYC:3231
F87C  60        RTS                             A:7F X:0 Y:64 P:25 SP:F9 CYC:3233
D0C6  61 80     ADC ($80,X) @ 80 = 0200 = 80    A:7F X:0 Y:64 P:25 SP:FB CYC:3239
D0C8  20 7D F8  JSR $F87D                       A:0 X:0 Y:64 P:27 SP:FB CYC:3245
F87D  D0 07     BNE $F886                       A:0 X:0 Y:64 P:27 SP:F9 CYC:3251
F87F  30 05     BMI $F886                       A:0 X:0 Y:64 P:27 SP:F9 CYC:3253
F881  70 03     BVS $F886                       A:0 X:0 Y:64 P:27 SP:F9 CYC:3255
F883  90 01     BCC $F886                       A:0 X:0 Y:64 P:27 SP:F9 CYC:3257
F885  60        RTS                             A:0 X:0 Y:64 P:27 SP:F9 CYC:3259
D0CB  C8        INY                             A:0 X:0 Y:64 P:27 SP:FB CYC:3265
D0CC  A9 40     LDA #$40                        A:0 X:0 Y:65 P:25 SP:FB CYC:3267
D0CE  8D 00 02  STA $0200 = 80                  A:40 X:0 Y:65 P:25 SP:FB CYC:3269
D0D1  20 89 F8  JSR $F889                       A:40 X:0 Y:65 P:25 SP:FB CYC:3273
F889  24 01     BIT $01 = FF                    A:40 X:0 Y:65 P:25 SP:F9 CYC:3279
F88B  A9 40     LDA #$40                        A:40 X:0 Y:65 P:E5 SP:F9 CYC:3282
F88D  60        RTS                             A:40 X:0 Y:65 P:65 SP:F9 CYC:3284
D0D4  C1 80     CMP ($80,X) @ 80 = 0200 = 40    A:40 X:0 Y:65 P:65 SP:FB CYC:3290
D0D6  20 8E F8  JSR $F88E                       A:40 X:0 Y:65 P:67 SP:FB CYC:3296
F88E  30 07     BMI $F897                       A:40 X:0 Y:65 P:67 SP:F9 CYC:3302
F890  90 05     BCC $F897                       A:40 X:0 Y:65 P:67 SP:F9 CYC:3304
F892  D0 03     BNE $F897                       A:40 X:0 Y:65 P:67 SP:F9 CYC:3306
F894  50 01     BVC $F897                       A:40 X:0 Y:65 P:67 SP:F9 CYC:3308
F896  60        RTS                             A:40 X:0 Y:65 P:67 SP:F9 CYC:3310
D0D9  C8        INY                             A:40 X:0 Y:65 P:67 SP:FB CYC:3316
D0DA  48        PHA                             A:40 X:0 Y:66 P:65 SP:FB CYC:3318
D0DB  A9 3F     LDA #$3F                        A:40 X:0 Y:66 P:65 SP:FA CYC:3321
D0DD  8D 00 02  STA $0200 = 40                  A:3F X:0 Y:66 P:65 SP:FA CYC:3323
D0E0  68        PLA                             A:3F X:0 Y:66 P:65 SP:FA CYC:3327
D0E1  20 9A F8  JSR $F89A                       A:40 X:0 Y:66 P:65 SP:FB CYC:3331
F89A  B8        CLV                             A:40 X:0 Y:66 P:65 SP:F9 CYC:3337
F89B  60        RTS                             A:40 X:0 Y:66 P:25 SP:F9 CYC:3339
D0E4  C1 80     CMP ($80,X) @ 80 = 0200 = 3F    A:40 X:0 Y:66 P:25 SP:FB CYC:3345
D0E6  20 9C F8  JSR $F89C                       A:40 X:0 Y:66 P:25 SP:FB CYC:3351
F89C  F0 07     BEQ $F8A5                       A:40 X:0 Y:66 P:25 SP:F9 CYC:3357
F89E  30 05     BMI $F8A5                       A:40 X:0 Y:66 P:25 SP:F9 CYC:3359
F8A0  90 03     BCC $F8A5                       A:40 X:0 Y:66 P:25 SP:F9 CYC:3361
F8A2  70 01     BVS $F8A5                       A:40 X:0 Y:66 P:25 SP:F9 CYC:3363
F8A4  60        RTS                             A:40 X:0 Y:66 P:25 SP:F9 CYC:3365
D0E9  C8        INY                             A:40 X:0 Y:66 P:25 SP:FB CYC:3371
D0EA  48        PHA                             A:40 X:0 Y:67 P:25 SP:FB CYC:3373
D0EB  A9 41     LDA #$41                        A:40 X:0 Y:67 P:25 SP:FA CYC:3376
D0ED  8D 00 02  STA $0200 = 3F                  A:41 X:0 Y:67 P:25 SP:FA CYC:3378
D0F0  68        PLA                             A:41 X:0 Y:67 P:25 SP:FA CYC:3382
D0F1  C1 80     CMP ($80,X) @ 80 = 0200 = 41    A:40 X:0 Y:67 P:25 SP:FB CYC:3386
D0F3  20 A8 F8  JSR $F8A8                       A:40 X:0 Y:67 P:A4 SP:FB CYC:3392
F8A8  F0 05     BEQ $F8AF                       A:40 X:0 Y:67 P:A4 SP:F9 CYC:3398
F8AA  10 03     BPL $F8AF                       A:40 X:0 Y:67 P:A4 SP:F9 CYC:3400
F8AC  10 01     BPL $F8AF                       A:40 X:0 Y:67 P:A4 SP:F9 CYC:3402
F8AE  60        RTS                             A:40 X:0 Y:67 P:A4 SP:F9 CYC:3404
D0F6  C8        INY                             A:40 X:0 Y:67 P:A4 SP:FB CYC:3410
D0F7  48        PHA                             A:40 X:0 Y:68 P:24 SP:FB CYC:3412
D0F8  A9 00     LDA #$00                        A:40 X:0 Y:68 P:24 SP:FA CYC:3415
D0FA  8D 00 02  STA $0200 = 41                  A:0 X:0 Y:68 P:26 SP:FA CYC:3417
D0FD  68        PLA                             A:0 X:0 Y:68 P:26 SP:FA CYC:3421
D0FE  20 B2 F8  JSR $F8B2                       A:40 X:0 Y:68 P:24 SP:FB CYC:3425
F8B2  A9 80     LDA #$80                        A:40 X:0 Y:68 P:24 SP:F9 CYC:3431
F8B4  60        RTS                             A:80 X:0 Y:68 P:A4 SP:F9 CYC:3433
D101  C1 80     CMP ($80,X) @ 80 = 0200 = 00    A:80 X:0 Y:68 P:A4 SP:FB CYC:3439
D103  20 B5 F8  JSR $F8B5                       A:80 X:0 Y:68 P:A5 SP:FB CYC:3445
F8B5  F0 05     BEQ $F8BC                       A:80 X:0 Y:68 P:A5 SP:F9 CYC:3451
F8B7  10 03     BPL $F8BC                       A:80 X:0 Y:68 P:A5 SP:F9 CYC:3453
F8B9  90 01     BCC $F8BC                       A:80 X:0 Y:68 P:A5 SP:F9 CYC:3455
F8BB  60        RTS                             A:80 X:0 Y:68 P:A5 SP:F9 CYC:3457
D106  C8        INY                             A:80 X:0 Y:68 P:A5 SP:FB CYC:3463
D107  48        PHA                             A:80 X:0 Y:69 P:25 SP:FB CYC:3465
D108  A9 80     LDA #$80                        A:80 X:0 Y:69 P:25 SP:FA CYC:3468
D10A  8D 00 02  STA $0200 = 00                  A:80 X:0 Y:69 P:A5 SP:FA CYC:3470
D10D  68        PLA                             A:80 X:0 Y:69 P:A5 SP:FA CYC:3474
D10E  C1 80     CMP ($80,X) @ 80 = 0200 = 80    A:80 X:0 Y:69 P:A5 SP:FB CYC:3478
D110  20 BF F8  JSR $F8BF                       A:80 X:0 Y:69 P:27 SP:FB CYC:3484
F8BF  D0 05     BNE $F8C6                       A:80 X:0 Y:69 P:27 SP:F9 CYC:3490
F8C1  30 03     BMI $F8C6                       A:80 X:0 Y:69 P:27 SP:F9 CYC:3492
F8C3  90 01     BCC $F8C6                       A:80 X:0 Y:69 P:27 SP:F9 CYC:3494
F8C5  60        RTS                             A:80 X:0 Y:69 P:27 SP:F9 CYC:3496
D113  C8        INY                             A:80 X:0 Y:69 P:27 SP:FB CYC:3502
D114  48        PHA                             A:80 X:0 Y:6A P:25 SP:FB CYC:3504
D115  A9 81     LDA #$81                        A:80 X:0 Y:6A P:25 SP:FA CYC:3507
D117  8D 00 02  STA $0200 = 80                  A:81 X:0 Y:6A P:A5 SP:FA CYC:3509
D11A  68        PLA                             A:81 X:0 Y:6A P:A5 SP:FA CYC:3513
D11B  C1 80     CMP ($80,X) @ 80 = 0200 = 81    A:80 X:0 Y:6A P:A5 SP:FB CYC:3517
D11D  20 C9 F8  JSR $F8C9                       A:80 X:0 Y:6A P:A4 SP:FB CYC:3523
F8C9  B0 05     BCS $F8D0                       A:80 X:0 Y:6A P:A4 SP:F9 CYC:3529
F8CB  F0 03     BEQ $F8D0                       A:80 X:0 Y:6A P:A4 SP:F9 CYC:3531
F8CD  10 01     BPL $F8D0                       A:80 X:0 Y:6A P:A4 SP:F9 CYC:3533
F8CF  60        RTS                             A:80 X:0 Y:6A P:A4 SP:F9 CYC:3535
D120  C8        INY                             A:80 X:0 Y:6A P:A4 SP:FB CYC:3541
D121  48        PHA                             A:80 X:0 Y:6B P:24 SP:FB CYC:3543
D122  A9 7F     LDA #$7F                        A:80 X:0 Y:6B P:24 SP:FA CYC:3546
D124  8D 00 02  STA $0200 = 81                  A:7F X:0 Y:6B P:24 SP:FA CYC:3548
D127  68        PLA                             A:7F X:0 Y:6B P:24 SP:FA CYC:3552
D128  C1 80     CMP ($80,X) @ 80 = 0200 = 7F    A:80 X:0 Y:6B P:A4 SP:FB CYC:3556
D12A  20 D3 F8  JSR $F8D3                       A:80 X:0 Y:6B P:25 SP:FB CYC:3562
F8D3  90 05     BCC $F8DA                       A:80 X:0 Y:6B P:25 SP:F9 CYC:3568
F8D5  F0 03     BEQ $F8DA                       A:80 X:0 Y:6B P:25 SP:F9 CYC:3570
F8D7  30 01     BMI $F8DA                       A:80 X:0 Y:6B P:25 SP:F9 CYC:3572
F8D9  60        RTS                             A:80 X:0 Y:6B P:25 SP:F9 CYC:3574
D12D  C8        INY                             A:80 X:0 Y:6B P:25 SP:FB CYC:3580
D12E  A9 40     LDA #$40                        A:80 X:0 Y:6C P:25 SP:FB CYC:3582
D130  8D 00 02  STA $0200 = 7F                  A:40 X:0 Y:6C P:25 SP:FB CYC:3584
D133  20 31 F9  JSR $F931                       A:40 X:0 Y:6C P:25 SP:FB CYC:3588
F931  24 01     BIT $01 = FF                    A:40 X:0 Y:6C P:25 SP:F9 CYC:3594
F933  A9 40     LDA #$40                        A:40 X:0 Y:6C P:E5 SP:F9 CYC:3597
F935  38        SEC                             A:40 X:0 Y:6C P:65 SP:F9 CYC:3599
F936  60        RTS                             A:40 X:0 Y:6C P:65 SP:F9 CYC:3601
D136  E1 80     SBC ($80,X) @ 80 = 0200 = 40    A:40 X:0 Y:6C P:65 SP:FB CYC:3607
D138  20 37 F9  JSR $F937                       A:0 X:0 Y:6C P:27 SP:FB CYC:3613
F937  30 0B     BMI $F944                       A:0 X:0 Y:6C P:27 SP:F9 CYC:3619
F939  90 09     BCC $F944                       A:0 X:0 Y:6C P:27 SP:F9 CYC:3621
F93B  D0 07     BNE $F944                       A:0 X:0 Y:6C P:27 SP:F9 CYC:3623
F93D  70 05     BVS $F944                       A:0 X:0 Y:6C P:27 SP:F9 CYC:3625
F93F  C9 00     CMP #$00                        A:0 X:0 Y:6C P:27 SP:F9 CYC:3627
F941  D0 01     BNE $F944                       A:0 X:0 Y:6C P:27 SP:F9 CYC:3629
F943  60        RTS                             A:0 X:0 Y:6C P:27 SP:F9 CYC:3631
D13B  C8        INY                             A:0 X:0 Y:6C P:27 SP:FB CYC:3637
D13C  A9 3F     LDA #$3F                        A:0 X:0 Y:6D P:25 SP:FB CYC:3639
D13E  8D 00 02  STA $0200 = 40                  A:3F X:0 Y:6D P:25 SP:FB CYC:3641
D141  20 47 F9  JSR $F947                       A:3F X:0 Y:6D P:25 SP:FB CYC:3645
F947  B8        CLV                             A:3F X:0 Y:6D P:25 SP:F9 CYC:3651
F948  38        SEC                             A:3F X:0 Y:6D P:25 SP:F9 CYC:3653
F949  A9 40     LDA #$40                        A:3F X:0 Y:6D P:25 SP:F9 CYC:3655
F94B  60        RTS                             A:40 X:0 Y:6D P:25 SP:F9 CYC:3657
D144  E1 80     SBC ($80,X) @ 80 = 0200 = 3F    A:40 X:0 Y:6D P:25 SP:FB CYC:3663
D146  20 4C F9  JSR $F94C                       A:1 X:0 Y:6D P:25 SP:FB CYC:3669
F94C  F0 0B     BEQ $F959                       A:1 X:0 Y:6D P:25 SP:F9 CYC:3675
F94E  30 09     BMI $F959                       A:1 X:0 Y:6D P:25 SP:F9 CYC:3677
F950  90 07     BCC $F959                       A:1 X:0 Y:6D P:25 SP:F9 CYC:3679
F952  70 05     BVS $F959                       A:1 X:0 Y:6D P:25 SP:F9 CYC:3681
F954  C9 01     CMP #$01                        A:1 X:0 Y:6D P:25 SP:F9 CYC:3683
F956  D0 01     BNE $F959                       A:1 X:0 Y:6D P:27 SP:F9 CYC:3685
F958  60        RTS                             A:1 X:0 Y:6D P:27 SP:F9 CYC:3687
D149  C8        INY                             A:1 X:0 Y:6D P:27 SP:FB CYC:3693
D14A  A9 41     LDA #$41                        A:1 X:0 Y:6E P:25 SP:FB CYC:3695
D14C  8D 00 02  STA $0200 = 3F                  A:41 X:0 Y:6E P:25 SP:FB CYC:3697
D14F  20 5C F9  JSR $F95C                       A:41 X:0 Y:6E P:25 SP:FB CYC:3701
F95C  A9 40     LDA #$40                        A:41 X:0 Y:6E P:25 SP:F9 CYC:3707
F95E  38        SEC                             A:40 X:0 Y:6E P:25 SP:F9 CYC:3709
F95F  24 01     BIT $01 = FF                    A:40 X:0 Y:6E P:25 SP:F9 CYC:3711
F961  60        RTS                             A:40 X:0 Y:6E P:E5 SP:F9 CYC:3714
D152  E1 80     SBC ($80,X) @ 80 = 0200 = 41    A:40 X:0 Y:6E P:E5 SP:FB CYC:3720
D154  20 62 F9  JSR $F962                       A:FF X:0 Y:6E P:A4 SP:FB CYC:3726
F962  B0 0B     BCS $F96F                       A:FF X:0 Y:6E P:A4 SP:F9 CYC:3732
F964  F0 09     BEQ $F96F                       A:FF X:0 Y:6E P:A4 SP:F9 CYC:3734
F966  10 07     BPL $F96F                       A:FF X:0 Y:6E P:A4 SP:F9 CYC:3736
F968  70 05     BVS $F96F                       A:FF X:0 Y:6E P:A4 SP:F9 CYC:3738
F96A  C9 FF     CMP #$FF                        A:FF X:0 Y:6E P:A4 SP:F9 CYC:3740
F96C  D0 01     BNE $F96F                       A:FF X:0 Y:6E P:27 SP:F9 CYC:3742
F96E  60        RTS                             A:FF X:0 Y:6E P:27 SP:F9 CYC:3744
D157  C8        INY                             A:FF X:0 Y:6E P:27 SP:FB CYC:3750
D158  A9 00     LDA #$00                        A:FF X:0 Y:6F P:25 SP:FB CYC:3752
D15A  8D 00 02  STA $0200 = 41                  A:0 X:0 Y:6F P:27 SP:FB CYC:3754
D15D  20 72 F9  JSR $F972                       A:0 X:0 Y:6F P:27 SP:FB CYC:3758
F972  18        CLC                             A:0 X:0 Y:6F P:27 SP:F9 CYC:3764
F973  A9 80     LDA #$80                        A:0 X:0 Y:6F P:26 SP:F9 CYC:3766
F975  60        RTS                             A:80 X:0 Y:6F P:A4 SP:F9 CYC:3768
D160  E1 80     SBC ($80,X) @ 80 = 0200 = 00    A:80 X:0 Y:6F P:A4 SP:FB CYC:3774
D162  20 76 F9  JSR $F976                       A:7F X:0 Y:6F P:65 SP:FB CYC:3780
F976  90 05     BCC $F97D                       A:7F X:0 Y:6F P:65 SP:F9 CYC:3786
F978  C9 7F     CMP #$7F                        A:7F X:0 Y:6F P:65 SP:F9 CYC:3788
F97A  D0 01     BNE $F97D                       A:7F X:0 Y:6F P:67 SP:F9 CYC:3790
F97C  60        RTS                             A:7F X:0 Y:6F P:67 SP:F9 CYC:3792
D165  C8        INY                             A:7F X:0 Y:6F P:67 SP:FB CYC:3798
D166  A9 7F     LDA #$7F                        A:7F X:0 Y:70 P:65 SP:FB CYC:3800
D168  8D 00 02  STA $0200 = 00                  A:7F X:0 Y:70 P:65 SP:FB CYC:3802
D16B  20 80 F9  JSR $F980                       A:7F X:0 Y:70 P:65 SP:FB CYC:3806
F980  38        SEC                             A:7F X:0 Y:70 P:65 SP:F9 CYC:3812
F981  A9 81     LDA #$81                        A:7F X:0 Y:70 P:65 SP:F9 CYC:3814
F983  60        RTS                             A:81 X:0 Y:70 P:E5 SP:F9 CYC:3816
D16E  E1 80     SBC ($80,X) @ 80 = 0200 = 7F    A:81 X:0 Y:70 P:E5 SP:FB CYC:3822
D170  20 84 F9  JSR $F984                       A:2 X:0 Y:70 P:65 SP:FB CYC:3828
F984  50 07     BVC $F98D                       A:2 X:0 Y:70 P:65 SP:F9 CYC:3834
F986  90 05     BCC $F98D                       A:2 X:0 Y:70 P:65 SP:F9 CYC:3836
F988  C9 02     CMP #$02                        A:2 X:0 Y:70 P:65 SP:F9 CYC:3838
F98A  D0 01     BNE $F98D                       A:2 X:0 Y:70 P:67 SP:F9 CYC:3840
F98C  60        RTS                             A:2 X:0 Y:70 P:67 SP:F9 CYC:3842
D173  60        RTS                             A:2 X:0 Y:70 P:67 SP:FB CYC:3848
C612  20 74 D1  JSR $D174                       A:2 X:0 Y:70 P:67 SP:FD CYC:3854
D174  A9 55     LDA #$55                        A:2 X:0 Y:70 P:67 SP:FB CYC:3860
D176  85 78     STA $78 = 00                    A:55 X:0 Y:70 P:65 SP:FB CYC:3862
D178  A9 FF     LDA #$FF                        A:55 X:0 Y:70 P:65 SP:FB CYC:3865
D17A  85 01     STA $01 = FF                    A:FF X:0 Y:70 P:E5 SP:FB CYC:3867
D17C  24 01     BIT $01 = FF                    A:FF X:0 Y:70 P:E5 SP:FB CYC:3870
D17E  A0 11     LDY #$11                        A:FF X:0 Y:70 P:E5 SP:FB CYC:3873
D180  A2 23     LDX #$23                        A:FF X:0 Y:11 P:65 SP:FB CYC:3875
D182  A9 00     LDA #$00                        A:FF X:23 Y:11 P:65 SP:FB CYC:3877
D184  A5 78     LDA $78 = 55                    A:0 X:23 Y:11 P:67 SP:FB CYC:3879
D186  F0 10     BEQ $D198                       A:55 X:23 Y:11 P:65 SP:FB CYC:3882
D188  30 0E     BMI $D198                       A:55 X:23 Y:11 P:65 SP:FB CYC:3884
D18A  C9 55     CMP #$55                        A:55 X:23 Y:11 P:65 SP:FB CYC:3886
D18C  D0 0A     BNE $D198                       A:55 X:23 Y:11 P:67 SP:FB CYC:3888
D18E  C0 11     CPY #$11                        A:55 X:23 Y:11 P:67 SP:FB CYC:3890
D190  D0 06     BNE $D198                       A:55 X:23 Y:11 P:67 SP:FB CYC:3892
D192  E0 23     CPX #$23                        A:55 X:23 Y:11 P:67 SP:FB CYC:3894
D194  50 02     BVC $D198                       A:55 X:23 Y:11 P:67 SP:FB CYC:3896
D196  F0 04     BEQ $D19C                       A:55 X:23 Y:11 P:67 SP:FB CYC:3898
D19C  A9 46     LDA #$46                        A:55 X:23 Y:11 P:67 SP:FB CYC:3901
D19E  24 01     BIT $01 = FF                    A:46 X:23 Y:11 P:65 SP:FB CYC:3903
D1A0  85 78     STA $78 = 55                    A:46 X:23 Y:11 P:E5 SP:FB CYC:3906
D1A2  F0 0A     BEQ $D1AE                       A:46 X:23 Y:11 P:E5 SP:FB CYC:3909
D1A4  10 08     BPL $D1AE                       A:46 X:23 Y:11 P:E5 SP:FB CYC:3911
D1A6  50 06     BVC $D1AE                       A:46 X:23 Y:11 P:E5 SP:FB CYC:3913
D1A8  A5 78     LDA $78 = 46                    A:46 X:23 Y:11 P:E5 SP:FB CYC:3915
D1AA  C9 46     CMP #$46                        A:46 X:23 Y:11 P:65 SP:FB CYC:3918
D1AC  F0 04     BEQ $D1B2                       A:46 X:23 Y:11 P:67 SP:FB CYC:3920
D1B2  A9 55     LDA #$55                        A:46 X:23 Y:11 P:67 SP:FB CYC:3923
D1B4  85 78     STA $78 = 46                    A:55 X:23 Y:11 P:65 SP:FB CYC:3925
D1B6  24 01     BIT $01 = FF                    A:55 X:23 Y:11 P:65 SP:FB CYC:3928
D1B8  A9 11     LDA #$11                        A:55 X:23 Y:11 P:E5 SP:FB CYC:3931
D1BA  A2 23     LDX #$23                        A:11 X:23 Y:11 P:65 SP:FB CYC:3933
D1BC  A0 00     LDY #$00                        A:11 X:23 Y:11 P:65 SP:FB CYC:3935
D1BE  A4 78     LDY $78 = 55                    A:11 X:23 Y:0 P:67 SP:FB CYC:3937
D1C0  F0 10     BEQ $D1D2                       A:11 X:23 Y:55 P:65 SP:FB CYC:3940
D1C2  30 0E     BMI $D1D2                       A:11 X:23 Y:55 P:65 SP:FB CYC:3942
D1C4  C0 55     CPY #$55                        A:11 X:23 Y:55 P:65 SP:FB CYC:3944
D1C6  D0 0A     BNE $D1D2                       A:11 X:23 Y:55 P:67 SP:FB CYC:3946
D1C8  C9 11     CMP #$11                        A:11 X:23 Y:55 P:67 SP:FB CYC:3948
D1CA  D0 06     BNE $D1D2                       A:11 X:23 Y:55 P:67 SP:FB CYC:3950
D1CC  E0 23     CPX #$23                        A:11 X:23 Y:55 P:67 SP:FB CYC:3952
D1CE  50 02     BVC $D1D2                       A:11 X:23 Y:55 P:67 SP:FB CYC:3954
D1D0  F0 04     BEQ $D1D6                       A:11 X:23 Y:55 P:67 SP:FB CYC:3956
D1D6  A0 46     LDY #$46                        A:11 X:23 Y:55 P:67 SP:FB CYC:3959
D1D8  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:65 SP:FB CYC:3961
D1DA  84 78     STY $78 = 55                    A:11 X:23 Y:46 P:E5 SP:FB CYC:3964
D1DC  F0 0A     BEQ $D1E8                       A:11 X:23 Y:46 P:E5 SP:FB CYC:3967
D1DE  10 08     BPL $D1E8                       A:11 X:23 Y:46 P:E5 SP:FB CYC:3969
D1E0  50 06     BVC $D1E8                       A:11 X:23 Y:46 P:E5 SP:FB CYC:3971
D1E2  A4 78     LDY $78 = 46                    A:11 X:23 Y:46 P:E5 SP:FB CYC:3973
D1E4  C0 46     CPY #$46                        A:11 X:23 Y:46 P:65 SP:FB CYC:3976
D1E6  F0 04     BEQ $D1EC                       A:11 X:23 Y:46 P:67 SP:FB CYC:3978
D1EC  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:67 SP:FB CYC:3981
D1EE  A9 55     LDA #$55                        A:11 X:23 Y:46 P:E5 SP:FB CYC:3984
D1F0  85 78     STA $78 = 46                    A:55 X:23 Y:46 P:65 SP:FB CYC:3986
D1F2  A0 11     LDY #$11                        A:55 X:23 Y:46 P:65 SP:FB CYC:3989
D1F4  A9 23     LDA #$23                        A:55 X:23 Y:11 P:65 SP:FB CYC:3991
D1F6  A2 00     LDX #$00                        A:23 X:23 Y:11 P:65 SP:FB CYC:3993
D1F8  A6 78     LDX $78 = 55                    A:23 X:0 Y:11 P:67 SP:FB CYC:3995
D1FA  F0 10     BEQ $D20C                       A:23 X:55 Y:11 P:65 SP:FB CYC:3998
D1FC  30 0E     BMI $D20C                       A:23 X:55 Y:11 P:65 SP:FB CYC:4000
D1FE  E0 55     CPX #$55                        A:23 X:55 Y:11 P:65 SP:FB CYC:4002
D200  D0 0A     BNE $D20C                       A:23 X:55 Y:11 P:67 SP:FB CYC:4004
D202  C0 11     CPY #$11                        A:23 X:55 Y:11 P:67 SP:FB CYC:4006
D204  D0 06     BNE $D20C                       A:23 X:55 Y:11 P:67 SP:FB CYC:4008
D206  C9 23     CMP #$23                        A:23 X:55 Y:11 P:67 SP:FB CYC:4010
D208  50 02     BVC $D20C                       A:23 X:55 Y:11 P:67 SP:FB CYC:4012
D20A  F0 04     BEQ $D210                       A:23 X:55 Y:11 P:67 SP:FB CYC:4014
D210  A2 46     LDX #$46                        A:23 X:55 Y:11 P:67 SP:FB CYC:4017
D212  24 01     BIT $01 = FF                    A:23 X:46 Y:11 P:65 SP:FB CYC:4019
D214  86 78     STX $78 = 55                    A:23 X:46 Y:11 P:E5 SP:FB CYC:4022
D216  F0 0A     BEQ $D222                       A:23 X:46 Y:11 P:E5 SP:FB CYC:4025
D218  10 08     BPL $D222                       A:23 X:46 Y:11 P:E5 SP:FB CYC:4027
D21A  50 06     BVC $D222                       A:23 X:46 Y:11 P:E5 SP:FB CYC:4029
D21C  A6 78     LDX $78 = 46                    A:23 X:46 Y:11 P:E5 SP:FB CYC:4031
D21E  E0 46     CPX #$46                        A:23 X:46 Y:11 P:65 SP:FB CYC:4034
D220  F0 04     BEQ $D226                       A:23 X:46 Y:11 P:67 SP:FB CYC:4036
D226  A9 C0     LDA #$C0                        A:23 X:46 Y:11 P:67 SP:FB CYC:4039
D228  85 78     STA $78 = 46                    A:C0 X:46 Y:11 P:E5 SP:FB CYC:4041
D22A  A2 33     LDX #$33                        A:C0 X:46 Y:11 P:E5 SP:FB CYC:4044
D22C  A0 88     LDY #$88                        A:C0 X:33 Y:11 P:65 SP:FB CYC:4046
D22E  A9 05     LDA #$05                        A:C0 X:33 Y:88 P:E5 SP:FB CYC:4048
D230  24 78     BIT $78 = C0                    A:5 X:33 Y:88 P:65 SP:FB CYC:4050
D232  10 10     BPL $D244                       A:5 X:33 Y:88 P:E7 SP:FB CYC:4053
D234  50 0E     BVC $D244                       A:5 X:33 Y:88 P:E7 SP:FB CYC:4055
D236  D0 0C     BNE $D244                       A:5 X:33 Y:88 P:E7 SP:FB CYC:4057
D238  C9 05     CMP #$05                        A:5 X:33 Y:88 P:E7 SP:FB CYC:4059
D23A  D0 08     BNE $D244                       A:5 X:33 Y:88 P:67 SP:FB CYC:4061
D23C  E0 33     CPX #$33                        A:5 X:33 Y:88 P:67 SP:FB CYC:4063
D23E  D0 04     BNE $D244                       A:5 X:33 Y:88 P:67 SP:FB CYC:4065
D240  C0 88     CPY #$88                        A:5 X:33 Y:88 P:67 SP:FB CYC:4067
D242  F0 04     BEQ $D248                       A:5 X:33 Y:88 P:67 SP:FB CYC:4069
D248  A9 03     LDA #$03                        A:5 X:33 Y:88 P:67 SP:FB CYC:4072
D24A  85 78     STA $78 = C0                    A:3 X:33 Y:88 P:65 SP:FB CYC:4074
D24C  A9 01     LDA #$01                        A:3 X:33 Y:88 P:65 SP:FB CYC:4077
D24E  24 78     BIT $78 = 03                    A:1 X:33 Y:88 P:65 SP:FB CYC:4079
D250  30 08     BMI $D25A                       A:1 X:33 Y:88 P:25 SP:FB CYC:4082
D252  70 06     BVS $D25A                       A:1 X:33 Y:88 P:25 SP:FB CYC:4084
D254  F0 04     BEQ $D25A                       A:1 X:33 Y:88 P:25 SP:FB CYC:4086
D256  C9 01     CMP #$01                        A:1 X:33 Y:88 P:25 SP:FB CYC:4088
D258  F0 04     BEQ $D25E                       A:1 X:33 Y:88 P:27 SP:FB CYC:4090
D25E  A0 7E     LDY #$7E                        A:1 X:33 Y:88 P:27 SP:FB CYC:4093
D260  A9 AA     LDA #$AA                        A:1 X:33 Y:7E P:25 SP:FB CYC:4095
D262  85 78     STA $78 = 03                    A:AA X:33 Y:7E P:A5 SP:FB CYC:4097
D264  20 B6 F7  JSR $F7B6                       A:AA X:33 Y:7E P:A5 SP:FB CYC:4100
F7B6  18        CLC                             A:AA X:33 Y:7E P:A5 SP:F9 CYC:4106
F7B7  A9 FF     LDA #$FF                        A:AA X:33 Y:7E P:A4 SP:F9 CYC:4108
F7B9  85 01     STA $01 = FF                    A:FF X:33 Y:7E P:A4 SP:F9 CYC:4110
F7BB  24 01     BIT $01 = FF                    A:FF X:33 Y:7E P:A4 SP:F9 CYC:4113
F7BD  A9 55     LDA #$55                        A:FF X:33 Y:7E P:E4 SP:F9 CYC:4116
F7BF  60        RTS                             A:55 X:33 Y:7E P:64 SP:F9 CYC:4118
D267  05 78     ORA $78 = AA                    A:55 X:33 Y:7E P:64 SP:FB CYC:4124
D269  20 C0 F7  JSR $F7C0                       A:FF X:33 Y:7E P:E4 SP:FB CYC:4127
F7C0  B0 09     BCS $F7CB                       A:FF X:33 Y:7E P:E4 SP:F9 CYC:4133
F7C2  10 07     BPL $F7CB                       A:FF X:33 Y:7E P:E4 SP:F9 CYC:4135
F7C4  C9 FF     CMP #$FF                        A:FF X:33 Y:7E P:E4 SP:F9 CYC:4137
F7C6  D0 03     BNE $F7CB                       A:FF X:33 Y:7E P:67 SP:F9 CYC:4139
F7C8  50 01     BVC $F7CB                       A:FF X:33 Y:7E P:67 SP:F9 CYC:4141
F7CA  60        RTS                             A:FF X:33 Y:7E P:67 SP:F9 CYC:4143
D26C  C8        INY                             A:FF X:33 Y:7E P:67 SP:FB CYC:4149
D26D  A9 00     LDA #$00                        A:FF X:33 Y:7F P:65 SP:FB CYC:4151
D26F  85 78     STA $78 = AA                    A:0 X:33 Y:7F P:67 SP:FB CYC:4153
D271  20 CE F7  JSR $F7CE                       A:0 X:33 Y:7F P:67 SP:FB CYC:4156
F7CE  38        SEC                             A:0 X:33 Y:7F P:67 SP:F9 CYC:4162
F7CF  B8        CLV                             A:0 X:33 Y:7F P:67 SP:F9 CYC:4164
F7D0  A9 00     LDA #$00                        A:0 X:33 Y:7F P:27 SP:F9 CYC:4166
F7D2  60        RTS                             A:0 X:33 Y:7F P:27 SP:F9 CYC:4168
D274  05 78     ORA $78 = 00                    A:0 X:33 Y:7F P:27 SP:FB CYC:4174
D276  20 D3 F7  JSR $F7D3                       A:0 X:33 Y:7F P:27 SP:FB CYC:4177
F7D3  D0 07     BNE $F7DC                       A:0 X:33 Y:7F P:27 SP:F9 CYC:4183
F7D5  70 05     BVS $F7DC                       A:0 X:33 Y:7F P:27 SP:F9 CYC:4185
F7D7  90 03     BCC $F7DC                       A:0 X:33 Y:7F P:27 SP:F9 CYC:4187
F7D9  30 01     BMI $F7DC                       A:0 X:33 Y:7F P:27 SP:F9 CYC:4189
F7DB  60        RTS                             A:0 X:33 Y:7F P:27 SP:F9 CYC:4191
D279  C8        INY                             A:0 X:33 Y:7F P:27 SP:FB CYC:4197
D27A  A9 AA     LDA #$AA                        A:0 X:33 Y:80 P:A5 SP:FB CYC:4199
D27C  85 78     STA $78 = 00                    A:AA X:33 Y:80 P:A5 SP:FB CYC:4201
D27E  20 DF F7  JSR $F7DF                       A:AA X:33 Y:80 P:A5 SP:FB CYC:4204
F7DF  18        CLC                             A:AA X:33 Y:80 P:A5 SP:F9 CYC:4210
F7E0  24 01     BIT $01 = FF                    A:AA X:33 Y:80 P:A4 SP:F9 CYC:4212
F7E2  A9 55     LDA #$55                        A:AA X:33 Y:80 P:E4 SP:F9 CYC:4215
F7E4  60        RTS                             A:55 X:33 Y:80 P:64 SP:F9 CYC:4217
D281  25 78     AND $78 = AA                    A:55 X:33 Y:80 P:64 SP:FB CYC:4223
D283  20 E5 F7  JSR $F7E5                       A:0 X:33 Y:80 P:66 SP:FB CYC:4226
F7E5  D0 07     BNE $F7EE                       A:0 X:33 Y:80 P:66 SP:F9 CYC:4232
F7E7  50 05     BVC $F7EE                       A:0 X:33 Y:80 P:66 SP:F9 CYC:4234
F7E9  B0 03     BCS $F7EE                       A:0 X:33 Y:80 P:66 SP:F9 CYC:4236
F7EB  30 01     BMI $F7EE                       A:0 X:33 Y:80 P:66 SP:F9 CYC:4238
F7ED  60        RTS                             A:0 X:33 Y:80 P:66 SP:F9 CYC:4240
D286  C8        INY                             A:0 X:33 Y:80 P:66 SP:FB CYC:4246
D287  A9 EF     LDA #$EF                        A:0 X:33 Y:81 P:E4 SP:FB CYC:4248
D289  85 78     STA $78 = AA                    A:EF X:33 Y:81 P:E4 SP:FB CYC:4250
D28B  20 F1 F7  JSR $F7F1                       A:EF X:33 Y:81 P:E4 SP:FB CYC:4253
F7F1  38        SEC                             A:EF X:33 Y:81 P:E4 SP:F9 CYC:4259
F7F2  B8        CLV                             A:EF X:33 Y:81 P:E5 SP:F9 CYC:4261
F7F3  A9 F8     LDA #$F8                        A:EF X:33 Y:81 P:A5 SP:F9 CYC:4263
F7F5  60        RTS                             A:F8 X:33 Y:81 P:A5 SP:F9 CYC:4265
D28E  25 78     AND $78 = EF                    A:F8 X:33 Y:81 P:A5 SP:FB CYC:4271
D290  20 F6 F7  JSR $F7F6                       A:E8 X:33 Y:81 P:A5 SP:FB CYC:4274
F7F6  90 09     BCC $F801                       A:E8 X:33 Y:81 P:A5 SP:F9 CYC:4280
F7F8  10 07     BPL $F801                       A:E8 X:33 Y:81 P:A5 SP:F9 CYC:4282
F7FA  C9 E8     CMP #$E8                        A:E8 X:33 Y:81 P:A5 SP:F9 CYC:4284
F7FC  D0 03     BNE $F801                       A:E8 X:33 Y:81 P:27 SP:F9 CYC:4286
F7FE  70 01     BVS $F801                       A:E8 X:33 Y:81 P:27 SP:F9 CYC:4288
F800  60        RTS                             A:E8 X:33 Y:81 P:27 SP:F9 CYC:4290
D293  C8        INY                             A:E8 X:33 Y:81 P:27 SP:FB CYC:4296
D294  A9 AA     LDA #$AA                        A:E8 X:33 Y:82 P:A5 SP:FB CYC:4298
D296  85 78     STA $78 = EF                    A:AA X:33 Y:82 P:A5 SP:FB CYC:4300
D298  20 04 F8  JSR $F804                       A:AA X:33 Y:82 P:A5 SP:FB CYC:4303
F804  18        CLC                             A:AA X:33 Y:82 P:A5 SP:F9 CYC:4309
F805  24 01     BIT $01 = FF                    A:AA X:33 Y:82 P:A4 SP:F9 CYC:4311
F807  A9 5F     LDA #$5F                        A:AA X:33 Y:82 P:E4 SP:F9 CYC:4314
F809  60        RTS                             A:5F X:33 Y:82 P:64 SP:F9 CYC:4316
D29B  45 78     EOR $78 = AA                    A:5F X:33 Y:82 P:64 SP:FB CYC:4322
D29D  20 0A F8  JSR $F80A                       A:F5 X:33 Y:82 P:E4 SP:FB CYC:4325
F80A  B0 09     BCS $F815                       A:F5 X:33 Y:82 P:E4 SP:F9 CYC:4331
F80C  10 07     BPL $F815                       A:F5 X:33 Y:82 P:E4 SP:F9 CYC:4333
F80E  C9 F5     CMP #$F5                        A:F5 X:33 Y:82 P:E4 SP:F9 CYC:4335
F810  D0 03     BNE $F815                       A:F5 X:33 Y:82 P:67 SP:F9 CYC:4337
F812  50 01     BVC $F815                       A:F5 X:33 Y:82 P:67 SP:F9 CYC:4339
F814  60        RTS                             A:F5 X:33 Y:82 P:67 SP:F9 CYC:4341
D2A0  C8        INY                             A:F5 X:33 Y:82 P:67 SP:FB CYC:4347
D2A1  A9 70     LDA #$70                        A:F5 X:33 Y:83 P:E5 SP:FB CYC:4349
D2A3  85 78     STA $78 = AA                    A:70 X:33 Y:83 P:65 SP:FB CYC:4351
D2A5  20 18 F8  JSR $F818                       A:70 X:33 Y:83 P:65 SP:FB CYC:4354
F818  38        SEC                             A:70 X:33 Y:83 P:65 SP:F9 CYC:4360
F819  B8        CLV                             A:70 X:33 Y:83 P:65 SP:F9 CYC:4362
F81A  A9 70     LDA #$70                        A:70 X:33 Y:83 P:25 SP:F9 CYC:4364
F81C  60        RTS                             A:70 X:33 Y:83 P:25 SP:F9 CYC:4366
D2A8  45 78     EOR $78 = 70                    A:70 X:33 Y:83 P:25 SP:FB CYC:4372
D2AA  20 1D F8  JSR $F81D                       A:0 X:33 Y:83 P:27 SP:FB CYC:4375
F81D  D0 07     BNE $F826                       A:0 X:33 Y:83 P:27 SP:F9 CYC:4381
F81F  70 05     BVS $F826                       A:0 X:33 Y:83 P:27 SP:F9 CYC:4383
F821  90 03     BCC $F826                       A:0 X:33 Y:83 P:27 SP:F9 CYC:4385
F823  30 01     BMI $F826                       A:0 X:33 Y:83 P:27 SP:F9 CYC:4387
F825  60        RTS                             A:0 X:33 Y:83 P:27 SP:F9 CYC:4389
D2AD  C8        INY                             A:0 X:33 Y:83 P:27 SP:FB CYC:4395
D2AE  A9 69     LDA #$69                        A:0 X:33 Y:84 P:A5 SP:FB CYC:4397
D2B0  85 78     STA $78 = 70                    A:69 X:33 Y:84 P:25 SP:FB CYC:4399
D2B2  20 29 F8  JSR $F829                       A:69 X:33 Y:84 P:25 SP:FB CYC:4402
F829  18        CLC                             A:69 X:33 Y:84 P:25 SP:F9 CYC:4408
F82A  24 01     BIT $01 = FF                    A:69 X:33 Y:84 P:24 SP:F9 CYC:4410
F82C  A9 00     LDA #$00                        A:69 X:33 Y:84 P:E4 SP:F9 CYC:4413
F82E  60        RTS                             A:0 X:33 Y:84 P:66 SP:F9 CYC:4415
D2B5  65 78     ADC $78 = 69                    A:0 X:33 Y:84 P:66 SP:FB CYC:4421
D2B7  20 2F F8  JSR $F82F                       A:69 X:33 Y:84 P:24 SP:FB CYC:4424
F82F  30 09     BMI $F83A                       A:69 X:33 Y:84 P:24 SP:F9 CYC:4430
F831  B0 07     BCS $F83A                       A:69 X:33 Y:84 P:24 SP:F9 CYC:4432
F833  C9 69     CMP #$69                        A:69 X:33 Y:84 P:24 SP:F9 CYC:4434
F835  D0 03     BNE $F83A                       A:69 X:33 Y:84 P:27 SP:F9 CYC:4436
F837  70 01     BVS $F83A                       A:69 X:33 Y:84 P:27 SP:F9 CYC:4438
F839  60        RTS                             A:69 X:33 Y:84 P:27 SP:F9 CYC:4440
D2BA  C8        INY                             A:69 X:33 Y:84 P:27 SP:FB CYC:4446
D2BB  20 3D F8  JSR $F83D                       A:69 X:33 Y:85 P:A5 SP:FB CYC:4448
F83D  38        SEC                             A:69 X:33 Y:85 P:A5 SP:F9 CYC:4454
F83E  24 01     BIT $01 = FF                    A:69 X:33 Y:85 P:A5 SP:F9 CYC:4456
F840  A9 00     LDA #$00                        A:69 X:33 Y:85 P:E5 SP:F9 CYC:4459
F842  60        RTS                             A:0 X:33 Y:85 P:67 SP:F9 CYC:4461
D2BE  65 78     ADC $78 = 69                    A:0 X:33 Y:85 P:67 SP:FB CYC:4467
D2C0  20 43 F8  JSR $F843                       A:6A X:33 Y:85 P:24 SP:FB CYC:4470
F843  30 09     BMI $F84E                       A:6A X:33 Y:85 P:24 SP:F9 CYC:4476
F845  B0 07     BCS $F84E                       A:6A X:33 Y:85 P:24 SP:F9 CYC:4478
F847  C9 6A     CMP #$6A                        A:6A X:33 Y:85 P:24 SP:F9 CYC:4480
F849  D0 03     BNE $F84E                       A:6A X:33 Y:85 P:27 SP:F9 CYC:4482
F84B  70 01     BVS $F84E                       A:6A X:33 Y:85 P:27 SP:F9 CYC:4484
F84D  60        RTS                             A:6A X:33 Y:85 P:27 SP:F9 CYC:4486
D2C3  C8        INY                             A:6A X:33 Y:85 P:27 SP:FB CYC:4492
D2C4  A9 7F     LDA #$7F                        A:6A X:33 Y:86 P:A5 SP:FB CYC:4494
D2C6  85 78     STA $78 = 69                    A:7F X:33 Y:86 P:25 SP:FB CYC:4496
D2C8  20 51 F8  JSR $F851                       A:7F X:33 Y:86 P:25 SP:FB CYC:4499
F851  38        SEC                             A:7F X:33 Y:86 P:25 SP:F9 CYC:4505
F852  B8        CLV                             A:7F X:33 Y:86 P:25 SP:F9 CYC:4507
F853  A9 7F     LDA #$7F                        A:7F X:33 Y:86 P:25 SP:F9 CYC:4509
F855  60        RTS                             A:7F X:33 Y:86 P:25 SP:F9 CYC:4511
D2CB  65 78     ADC $78 = 7F                    A:7F X:33 Y:86 P:25 SP:FB CYC:4517
D2CD  20 56 F8  JSR $F856                       A:FF X:33 Y:86 P:E4 SP:FB CYC:4520
F856  10 09     BPL $F861                       A:FF X:33 Y:86 P:E4 SP:F9 CYC:4526
F858  B0 07     BCS $F861                       A:FF X:33 Y:86 P:E4 SP:F9 CYC:4528
F85A  C9 FF     CMP #$FF                        A:FF X:33 Y:86 P:E4 SP:F9 CYC:4530
F85C  D0 03     BNE $F861                       A:FF X:33 Y:86 P:67 SP:F9 CYC:4532
F85E  50 01     BVC $F861                       A:FF X:33 Y:86 P:67 SP:F9 CYC:4534
F860  60        RTS                             A:FF X:33 Y:86 P:67 SP:F9 CYC:4536
D2D0  C8        INY                             A:FF X:33 Y:86 P:67 SP:FB CYC:4542
D2D1  A9 80     LDA #$80                        A:FF X:33 Y:87 P:E5 SP:FB CYC:4544
D2D3  85 78     STA $78 = 7F                    A:80 X:33 Y:87 P:E5 SP:FB CYC:4546
D2D5  20 64 F8  JSR $F864                       A:80 X:33 Y:87 P:E5 SP:FB CYC:4549
F864  18        CLC                             A:80 X:33 Y:87 P:E5 SP:F9 CYC:4555
F865  24 01     BIT $01 = FF                    A:80 X:33 Y:87 P:E4 SP:F9 CYC:4557
F867  A9 7F     LDA #$7F                        A:80 X:33 Y:87 P:E4 SP:F9 CYC:4560
F869  60        RTS                             A:7F X:33 Y:87 P:64 SP:F9 CYC:4562
D2D8  65 78     ADC $78 = 80                    A:7F X:33 Y:87 P:64 SP:FB CYC:4568
D2DA  20 6A F8  JSR $F86A                       A:FF X:33 Y:87 P:A4 SP:FB CYC:4571
F86A  10 09     BPL $F875                       A:FF X:33 Y:87 P:A4 SP:F9 CYC:4577
F86C  B0 07     BCS $F875                       A:FF X:33 Y:87 P:A4 SP:F9 CYC:4579
F86E  C9 FF     CMP #$FF                        A:FF X:33 Y:87 P:A4 SP:F9 CYC:4581
F870  D0 03     BNE $F875                       A:FF X:33 Y:87 P:27 SP:F9 CYC:4583
F872  70 01     BVS $F875                       A:FF X:33 Y:87 P:27 SP:F9 CYC:4585
F874  60        RTS                             A:FF X:33 Y:87 P:27 SP:F9 CYC:4587
D2DD  C8        INY                             A:FF X:33 Y:87 P:27 SP:FB CYC:4593
D2DE  20 78 F8  JSR $F878                       A:FF X:33 Y:88 P:A5 SP:FB CYC:4595
F878  38        SEC                             A:FF X:33 Y:88 P:A5 SP:F9 CYC:4601
F879  B8        CLV                             A:FF X:33 Y:88 P:A5 SP:F9 CYC:4603
F87A  A9 7F     LDA #$7F                        A:FF X:33 Y:88 P:A5 SP:F9 CYC:4605
F87C  60        RTS                             A:7F X:33 Y:88 P:25 SP:F9 CYC:4607
D2E1  65 78     ADC $78 = 80                    A:7F X:33 Y:88 P:25 SP:FB CYC:4613
D2E3  20 7D F8  JSR $F87D                       A:0 X:33 Y:88 P:27 SP:FB CYC:4616
F87D  D0 07     BNE $F886                       A:0 X:33 Y:88 P:27 SP:F9 CYC:4622
F87F  30 05     BMI $F886                       A:0 X:33 Y:88 P:27 SP:F9 CYC:4624
F881  70 03     BVS $F886                       A:0 X:33 Y:88 P:27 SP:F9 CYC:4626
F883  90 01     BCC $F886                       A:0 X:33 Y:88 P:27 SP:F9 CYC:4628
F885  60        RTS                             A:0 X:33 Y:88 P:27 SP:F9 CYC:4630
D2E6  C8        INY                             A:0 X:33 Y:88 P:27 SP:FB CYC:4636
D2E7  A9 40     LDA #$40                        A:0 X:33 Y:89 P:A5 SP:FB CYC:4638
D2E9  85 78     STA $78 = 80                    A:40 X:33 Y:89 P:25 SP:FB CYC:4640
D2EB  20 89 F8  JSR $F889                       A:40 X:33 Y:89 P:25 SP:FB CYC:4643
F889  24 01     BIT $01 = FF                    A:40 X:33 Y:89 P:25 SP:F9 CYC:4649
F88B  A9 40     LDA #$40                        A:40 X:33 Y:89 P:E5 SP:F9 CYC:4652
F88D  60        RTS                             A:40 X:33 Y:89 P:65 SP:F9 CYC:4654
D2EE  C5 78     CMP $78 = 40                    A:40 X:33 Y:89 P:65 SP:FB CYC:4660
D2F0  20 8E F8  JSR $F88E                       A:40 X:33 Y:89 P:67 SP:FB CYC:4663
F88E  30 07     BMI $F897                       A:40 X:33 Y:89 P:67 SP:F9 CYC:4669
F890  90 05     BCC $F897                       A:40 X:33 Y:89 P:67 SP:F9 CYC:4671
F892  D0 03     BNE $F897                       A:40 X:33 Y:89 P:67 SP:F9 CYC:4673
F894  50 01     BVC $F897                       A:40 X:33 Y:89 P:67 SP:F9 CYC:4675
F896  60        RTS                             A:40 X:33 Y:89 P:67 SP:F9 CYC:4677
D2F3  C8        INY                             A:40 X:33 Y:89 P:67 SP:FB CYC:4683
D2F4  48        PHA                             A:40 X:33 Y:8A P:E5 SP:FB CYC:4685
D2F5  A9 3F     LDA #$3F                        A:40 X:33 Y:8A P:E5 SP:FA CYC:4688
D2F7  85 78     STA $78 = 40                    A:3F X:33 Y:8A P:65 SP:FA CYC:4690
D2F9  68        PLA                             A:3F X:33 Y:8A P:65 SP:FA CYC:4693
D2FA  20 9A F8  JSR $F89A                       A:40 X:33 Y:8A P:65 SP:FB CYC:4697
F89A  B8        CLV                             A:40 X:33 Y:8A P:65 SP:F9 CYC:4703
F89B  60        RTS                             A:40 X:33 Y:8A P:25 SP:F9 CYC:4705
D2FD  C5 78     CMP $78 = 3F                    A:40 X:33 Y:8A P:25 SP:FB CYC:4711
D2FF  20 9C F8  JSR $F89C                       A:40 X:33 Y:8A P:25 SP:FB CYC:4714
F89C  F0 07     BEQ $F8A5                       A:40 X:33 Y:8A P:25 SP:F9 CYC:4720
F89E  30 05     BMI $F8A5                       A:40 X:33 Y:8A P:25 SP:F9 CYC:4722
F8A0  90 03     BCC $F8A5                       A:40 X:33 Y:8A P:25 SP:F9 CYC:4724
F8A2  70 01     BVS $F8A5                       A:40 X:33 Y:8A P:25 SP:F9 CYC:4726
F8A4  60        RTS                             A:40 X:33 Y:8A P:25 SP:F9 CYC:4728
D302  C8        INY                             A:40 X:33 Y:8A P:25 SP:FB CYC:4734
D303  48        PHA                             A:40 X:33 Y:8B P:A5 SP:FB CYC:4736
D304  A9 41     LDA #$41                        A:40 X:33 Y:8B P:A5 SP:FA CYC:4739
D306  85 78     STA $78 = 3F                    A:41 X:33 Y:8B P:25 SP:FA CYC:4741
D308  68        PLA                             A:41 X:33 Y:8B P:25 SP:FA CYC:4744
D309  C5 78     CMP $78 = 41                    A:40 X:33 Y:8B P:25 SP:FB CYC:4748
D30B  20 A8 F8  JSR $F8A8                       A:40 X:33 Y:8B P:A4 SP:FB CYC:4751
F8A8  F0 05     BEQ $F8AF                       A:40 X:33 Y:8B P:A4 SP:F9 CYC:4757
F8AA  10 03     BPL $F8AF                       A:40 X:33 Y:8B P:A4 SP:F9 CYC:4759
F8AC  10 01     BPL $F8AF                       A:40 X:33 Y:8B P:A4 SP:F9 CYC:4761
F8AE  60        RTS                             A:40 X:33 Y:8B P:A4 SP:F9 CYC:4763
D30E  C8        INY                             A:40 X:33 Y:8B P:A4 SP:FB CYC:4769
D30F  48        PHA                             A:40 X:33 Y:8C P:A4 SP:FB CYC:4771
D310  A9 00     LDA #$00                        A:40 X:33 Y:8C P:A4 SP:FA CYC:4774
D312  85 78     STA $78 = 41                    A:0 X:33 Y:8C P:26 SP:FA CYC:4776
D314  68        PLA                             A:0 X:33 Y:8C P:26 SP:FA CYC:4779
D315  20 B2 F8  JSR $F8B2                       A:40 X:33 Y:8C P:24 SP:FB CYC:4783
F8B2  A9 80     LDA #$80                        A:40 X:33 Y:8C P:24 SP:F9 CYC:4789
F8B4  60        RTS                             A:80 X:33 Y:8C P:A4 SP:F9 CYC:4791
D318  C5 78     CMP $78 = 00                    A:80 X:33 Y:8C P:A4 SP:FB CYC:4797
D31A  20 B5 F8  JSR $F8B5                       A:80 X:33 Y:8C P:A5 SP:FB CYC:4800
F8B5  F0 05     BEQ $F8BC                       A:80 X:33 Y:8C P:A5 SP:F9 CYC:4806
F8B7  10 03     BPL $F8BC                       A:80 X:33 Y:8C P:A5 SP:F9 CYC:4808
F8B9  90 01     BCC $F8BC                       A:80 X:33 Y:8C P:A5 SP:F9 CYC:4810
F8BB  60        RTS                             A:80 X:33 Y:8C P:A5 SP:F9 CYC:4812
D31D  C8        INY                             A:80 X:33 Y:8C P:A5 SP:FB CYC:4818
D31E  48        PHA                             A:80 X:33 Y:8D P:A5 SP:FB CYC:4820
D31F  A9 80     LDA #$80                        A:80 X:33 Y:8D P:A5 SP:FA CYC:4823
D321  85 78     STA $78 = 00                    A:80 X:33 Y:8D P:A5 SP:FA CYC:4825
D323  68        PLA                             A:80 X:33 Y:8D P:A5 SP:FA CYC:4828
D324  C5 78     CMP $78 = 80                    A:80 X:33 Y:8D P:A5 SP:FB CYC:4832
D326  20 BF F8  JSR $F8BF                       A:80 X:33 Y:8D P:27 SP:FB CYC:4835
F8BF  D0 05     BNE $F8C6                       A:80 X:33 Y:8D P:27 SP:F9 CYC:4841
F8C1  30 03     BMI $F8C6                       A:80 X:33 Y:8D P:27 SP:F9 CYC:4843
F8C3  90 01     BCC $F8C6                       A:80 X:33 Y:8D P:27 SP:F9 CYC:4845
F8C5  60        RTS                             A:80 X:33 Y:8D P:27 SP:F9 CYC:4847
D329  C8        INY                             A:80 X:33 Y:8D P:27 SP:FB CYC:4853
D32A  48        PHA                             A:80 X:33 Y:8E P:A5 SP:FB CYC:4855
D32B  A9 81     LDA #$81                        A:80 X:33 Y:8E P:A5 SP:FA CYC:4858
D32D  85 78     STA $78 = 80                    A:81 X:33 Y:8E P:A5 SP:FA CYC:4860
D32F  68        PLA                             A:81 X:33 Y:8E P:A5 SP:FA CYC:4863
D330  C5 78     CMP $78 = 81                    A:80 X:33 Y:8E P:A5 SP:FB CYC:4867
D332  20 C9 F8  JSR $F8C9                       A:80 X:33 Y:8E P:A4 SP:FB CYC:4870
F8C9  B0 05     BCS $F8D0                       A:80 X:33 Y:8E P:A4 SP:F9 CYC:4876
F8CB  F0 03     BEQ $F8D0                       A:80 X:33 Y:8E P:A4 SP:F9 CYC:4878
F8CD  10 01     BPL $F8D0                       A:80 X:33 Y:8E P:A4 SP:F9 CYC:4880
F8CF  60        RTS                             A:80 X:33 Y:8E P:A4 SP:F9 CYC:4882
D335  C8        INY                             A:80 X:33 Y:8E P:A4 SP:FB CYC:4888
D336  48        PHA                             A:80 X:33 Y:8F P:A4 SP:FB CYC:4890
D337  A9 7F     LDA #$7F                        A:80 X:33 Y:8F P:A4 SP:FA CYC:4893
D339  85 78     STA $78 = 81                    A:7F X:33 Y:8F P:24 SP:FA CYC:4895
D33B  68        PLA                             A:7F X:33 Y:8F P:24 SP:FA CYC:4898
D33C  C5 78     CMP $78 = 7F                    A:80 X:33 Y:8F P:A4 SP:FB CYC:4902
D33E  20 D3 F8  JSR $F8D3                       A:80 X:33 Y:8F P:25 SP:FB CYC:4905
F8D3  90 05     BCC $F8DA                       A:80 X:33 Y:8F P:25 SP:F9 CYC:4911
F8D5  F0 03     BEQ $F8DA                       A:80 X:33 Y:8F P:25 SP:F9 CYC:4913
F8D7  30 01     BMI $F8DA                       A:80 X:33 Y:8F P:25 SP:F9 CYC:4915
F8D9  60        RTS                             A:80 X:33 Y:8F P:25 SP:F9 CYC:4917
D341  C8        INY                             A:80 X:33 Y:8F P:25 SP:FB CYC:4923
D342  A9 40     LDA #$40                        A:80 X:33 Y:90 P:A5 SP:FB CYC:4925
D344  85 78     STA $78 = 7F                    A:40 X:33 Y:90 P:25 SP:FB CYC:4927
D346  20 31 F9  JSR $F931                       A:40 X:33 Y:90 P:25 SP:FB CYC:4930
F931  24 01     BIT $01 = FF                    A:40 X:33 Y:90 P:25 SP:F9 CYC:4936
F933  A9 40     LDA #$40                        A:40 X:33 Y:90 P:E5 SP:F9 CYC:4939
F935  38        SEC                             A:40 X:33 Y:90 P:65 SP:F9 CYC:4941
F936  60        RTS                             A:40 X:33 Y:90 P:65 SP:F9 CYC:4943
D349  E5 78     SBC $78 = 40                    A:40 X:33 Y:90 P:65 SP:FB CYC:4949
D34B  20 37 F9  JSR $F937                       A:0 X:33 Y:90 P:27 SP:FB CYC:4952
F937  30 0B     BMI $F944                       A:0 X:33 Y:90 P:27 SP:F9 CYC:4958
F939  90 09     BCC $F944                       A:0 X:33 Y:90 P:27 SP:F9 CYC:4960
F93B  D0 07     BNE $F944                       A:0 X:33 Y:90 P:27 SP:F9 CYC:4962
F93D  70 05     BVS $F944                       A:0 X:33 Y:90 P:27 SP:F9 CYC:4964
F93F  C9 00     CMP #$00                        A:0 X:33 Y:90 P:27 SP:F9 CYC:4966
F941  D0 01     BNE $F944                       A:0 X:33 Y:90 P:27 SP:F9 CYC:4968
F943  60        RTS                             A:0 X:33 Y:90 P:27 SP:F9 CYC:4970
D34E  C8        INY                             A:0 X:33 Y:90 P:27 SP:FB CYC:4976
D34F  A9 3F     LDA #$3F                        A:0 X:33 Y:91 P:A5 SP:FB CYC:4978
D351  85 78     STA $78 = 40                    A:3F X:33 Y:91 P:25 SP:FB CYC:4980
D353  20 47 F9  JSR $F947                       A:3F X:33 Y:91 P:25 SP:FB CYC:4983
F947  B8        CLV                             A:3F X:33 Y:91 P:25 SP:F9 CYC:4989
F948  38        SEC                             A:3F X:33 Y:91 P:25 SP:F9 CYC:4991
F949  A9 40     LDA #$40                        A:3F X:33 Y:91 P:25 SP:F9 CYC:4993
F94B  60        RTS                             A:40 X:33 Y:91 P:25 SP:F9 CYC:4995
D356  E5 78     SBC $78 = 3F                    A:40 X:33 Y:91 P:25 SP:FB CYC:5001
D358  20 4C F9  JSR $F94C                       A:1 X:33 Y:91 P:25 SP:FB CYC:5004
F94C  F0 0B     BEQ $F959                       A:1 X:33 Y:91 P:25 SP:F9 CYC:5010
F94E  30 09     BMI $F959                       A:1 X:33 Y:91 P:25 SP:F9 CYC:5012
F950  90 07     BCC $F959                       A:1 X:33 Y:91 P:25 SP:F9 CYC:5014
F952  70 05     BVS $F959                       A:1 X:33 Y:91 P:25 SP:F9 CYC:5016
F954  C9 01     CMP #$01                        A:1 X:33 Y:91 P:25 SP:F9 CYC:5018
F956  D0 01     BNE $F959                       A:1 X:33 Y:91 P:27 SP:F9 CYC:5020
F958  60        RTS                             A:1 X:33 Y:91 P:27 SP:F9 CYC:5022
D35B  C8        INY                             A:1 X:33 Y:91 P:27 SP:FB CYC:5028
D35C  A9 41     LDA #$41                        A:1 X:33 Y:92 P:A5 SP:FB CYC:5030
D35E  85 78     STA $78 = 3F                    A:41 X:33 Y:92 P:25 SP:FB CYC:5032
D360  20 5C F9  JSR $F95C                       A:41 X:33 Y:92 P:25 SP:FB CYC:5035
F95C  A9 40     LDA #$40                        A:41 X:33 Y:92 P:25 SP:F9 CYC:5041
F95E  38        SEC                             A:40 X:33 Y:92 P:25 SP:F9 CYC:5043
F95F  24 01     BIT $01 = FF                    A:40 X:33 Y:92 P:25 SP:F9 CYC:5045
F961  60        RTS                             A:40 X:33 Y:92 P:E5 SP:F9 CYC:5048
D363  E5 78     SBC $78 = 41                    A:40 X:33 Y:92 P:E5 SP:FB CYC:5054
D365  20 62 F9  JSR $F962                       A:FF X:33 Y:92 P:A4 SP:FB CYC:5057
F962  B0 0B     BCS $F96F                       A:FF X:33 Y:92 P:A4 SP:F9 CYC:5063
F964  F0 09     BEQ $F96F                       A:FF X:33 Y:92 P:A4 SP:F9 CYC:5065
F966  10 07     BPL $F96F                       A:FF X:33 Y:92 P:A4 SP:F9 CYC:5067
F968  70 05     BVS $F96F                       A:FF X:33 Y:92 P:A4 SP:F9 CYC:5069
F96A  C9 FF     CMP #$FF                        A:FF X:33 Y:92 P:A4 SP:F9 CYC:5071
F96C  D0 01     BNE $F96F                       A:FF X:33 Y:92 P:27 SP:F9 CYC:5073
F96E  60        RTS                             A:FF X:33 Y:92 P:27 SP:F9 CYC:5075
D368  C8        INY                             A:FF X:33 Y:92 P:27 SP:FB CYC:5081
D369  A9 00     LDA #$00                        A:FF X:33 Y:93 P:A5 SP:FB CYC:5083
D36B  85 78     STA $78 = 41                    A:0 X:33 Y:93 P:27 SP:FB CYC:5085
D36D  20 72 F9  JSR $F972                       A:0 X:33 Y:93 P:27 SP:FB CYC:5088
F972  18        CLC                             A:0 X:33 Y:93 P:27 SP:F9 CYC:5094
F973  A9 80     LDA #$80                        A:0 X:33 Y:93 P:26 SP:F9 CYC:5096
F975  60        RTS                             A:80 X:33 Y:93 P:A4 SP:F9 CYC:5098
D370  E5 78     SBC $78 = 00                    A:80 X:33 Y:93 P:A4 SP:FB CYC:5104
D372  20 76 F9  JSR $F976                       A:7F X:33 Y:93 P:65 SP:FB CYC:5107
F976  90 05     BCC $F97D                       A:7F X:33 Y:93 P:65 SP:F9 CYC:5113
F978  C9 7F     CMP #$7F                        A:7F X:33 Y:93 P:65 SP:F9 CYC:5115
F97A  D0 01     BNE $F97D                       A:7F X:33 Y:93 P:67 SP:F9 CYC:5117
F97C  60        RTS                             A:7F X:33 Y:93 P:67 SP:F9 CYC:5119
D375  C8        INY                             A:7F X:33 Y:93 P:67 SP:FB CYC:5125
D376  A9 7F     LDA #$7F                        A:7F X:33 Y:94 P:E5 SP:FB CYC:5127
D378  85 78     STA $78 = 00                    A:7F X:33 Y:94 P:65 SP:FB CYC:5129
D37A  20 80 F9  JSR $F980                       A:7F X:33 Y:94 P:65 SP:FB CYC:5132
F980  38        SEC                             A:7F X:33 Y:94 P:65 SP:F9 CYC:5138
F981  A9 81     LDA #$81                        A:7F X:33 Y:94 P:65 SP:F9 CYC:5140
F983  60        RTS                             A:81 X:33 Y:94 P:E5 SP:F9 CYC:5142
D37D  E5 78     SBC $78 = 7F                    A:81 X:33 Y:94 P:E5 SP:FB CYC:5148
D37F  20 84 F9  JSR $F984                       A:2 X:33 Y:94 P:65 SP:FB CYC:5151
F984  50 07     BVC $F98D                       A:2 X:33 Y:94 P:65 SP:F9 CYC:5157
F986  90 05     BCC $F98D                       A:2 X:33 Y:94 P:65 SP:F9 CYC:5159
F988  C9 02     CMP #$02                        A:2 X:33 Y:94 P:65 SP:F9 CYC:5161
F98A  D0 01     BNE $F98D                       A:2 X:33 Y:94 P:67 SP:F9 CYC:5163
F98C  60        RTS                             A:2 X:33 Y:94 P:67 SP:F9 CYC:5165
D382  C8        INY                             A:2 X:33 Y:94 P:67 SP:FB CYC:5171
D383  A9 40     LDA #$40                        A:2 X:33 Y:95 P:E5 SP:FB CYC:5173
D385  85 78     STA $78 = 7F                    A:40 X:33 Y:95 P:65 SP:FB CYC:5175
D387  20 89 F8  JSR $F889                       A:40 X:33 Y:95 P:65 SP:FB CYC:5178
F889  24 01     BIT $01 = FF                    A:40 X:33 Y:95 P:65 SP:F9 CYC:5184
F88B  A9 40     LDA #$40                        A:40 X:33 Y:95 P:E5 SP:F9 CYC:5187
F88D  60        RTS                             A:40 X:33 Y:95 P:65 SP:F9 CYC:5189
D38A  AA        TAX                             A:40 X:33 Y:95 P:65 SP:FB CYC:5195
D38B  E4 78     CPX $78 = 40                    A:40 X:40 Y:95 P:65 SP:FB CYC:5197
D38D  20 8E F8  JSR $F88E                       A:40 X:40 Y:95 P:67 SP:FB CYC:5200
F88E  30 07     BMI $F897                       A:40 X:40 Y:95 P:67 SP:F9 CYC:5206
F890  90 05     BCC $F897                       A:40 X:40 Y:95 P:67 SP:F9 CYC:5208
F892  D0 03     BNE $F897                       A:40 X:40 Y:95 P:67 SP:F9 CYC:5210
F894  50 01     BVC $F897                       A:40 X:40 Y:95 P:67 SP:F9 CYC:5212
F896  60        RTS                             A:40 X:40 Y:95 P:67 SP:F9 CYC:5214
D390  C8        INY                             A:40 X:40 Y:95 P:67 SP:FB CYC:5220
D391  A9 3F     LDA #$3F                        A:40 X:40 Y:96 P:E5 SP:FB CYC:5222
D393  85 78     STA $78 = 40                    A:3F X:40 Y:96 P:65 SP:FB CYC:5224
D395  20 9A F8  JSR $F89A                       A:3F X:40 Y:96 P:65 SP:FB CYC:5227
F89A  B8        CLV                             A:3F X:40 Y:96 P:65 SP:F9 CYC:5233
F89B  60        RTS                             A:3F X:40 Y:96 P:25 SP:F9 CYC:5235
D398  E4 78     CPX $78 = 3F                    A:3F X:40 Y:96 P:25 SP:FB CYC:5241
D39A  20 9C F8  JSR $F89C                       A:3F X:40 Y:96 P:25 SP:FB CYC:5244
F89C  F0 07     BEQ $F8A5                       A:3F X:40 Y:96 P:25 SP:F9 CYC:5250
F89E  30 05     BMI $F8A5                       A:3F X:40 Y:96 P:25 SP:F9 CYC:5252
F8A0  90 03     BCC $F8A5                       A:3F X:40 Y:96 P:25 SP:F9 CYC:5254
F8A2  70 01     BVS $F8A5                       A:3F X:40 Y:96 P:25 SP:F9 CYC:5256
F8A4  60        RTS                             A:3F X:40 Y:96 P:25 SP:F9 CYC:5258
D39D  C8        INY                             A:3F X:40 Y:96 P:25 SP:FB CYC:5264
D39E  A9 41     LDA #$41                        A:3F X:40 Y:97 P:A5 SP:FB CYC:5266
D3A0  85 78     STA $78 = 3F                    A:41 X:40 Y:97 P:25 SP:FB CYC:5268
D3A2  E4 78     CPX $78 = 41                    A:41 X:40 Y:97 P:25 SP:FB CYC:5271
D3A4  20 A8 F8  JSR $F8A8                       A:41 X:40 Y:97 P:A4 SP:FB CYC:5274
F8A8  F0 05     BEQ $F8AF                       A:41 X:40 Y:97 P:A4 SP:F9 CYC:5280
F8AA  10 03     BPL $F8AF                       A:41 X:40 Y:97 P:A4 SP:F9 CYC:5282
F8AC  10 01     BPL $F8AF                       A:41 X:40 Y:97 P:A4 SP:F9 CYC:5284
F8AE  60        RTS                             A:41 X:40 Y:97 P:A4 SP:F9 CYC:5286
D3A7  C8        INY                             A:41 X:40 Y:97 P:A4 SP:FB CYC:5292
D3A8  A9 00     LDA #$00                        A:41 X:40 Y:98 P:A4 SP:FB CYC:5294
D3AA  85 78     STA $78 = 41                    A:0 X:40 Y:98 P:26 SP:FB CYC:5296
D3AC  20 B2 F8  JSR $F8B2                       A:0 X:40 Y:98 P:26 SP:FB CYC:5299
F8B2  A9 80     LDA #$80                        A:0 X:40 Y:98 P:26 SP:F9 CYC:5305
F8B4  60        RTS                             A:80 X:40 Y:98 P:A4 SP:F9 CYC:5307
D3AF  AA        TAX                             A:80 X:40 Y:98 P:A4 SP:FB CYC:5313
D3B0  E4 78     CPX $78 = 00                    A:80 X:80 Y:98 P:A4 SP:FB CYC:5315
D3B2  20 B5 F8  JSR $F8B5                       A:80 X:80 Y:98 P:A5 SP:FB CYC:5318
F8B5  F0 05     BEQ $F8BC                       A:80 X:80 Y:98 P:A5 SP:F9 CYC:5324
F8B7  10 03     BPL $F8BC                       A:80 X:80 Y:98 P:A5 SP:F9 CYC:5326
F8B9  90 01     BCC $F8BC                       A:80 X:80 Y:98 P:A5 SP:F9 CYC:5328
F8BB  60        RTS                             A:80 X:80 Y:98 P:A5 SP:F9 CYC:5330
D3B5  C8        INY                             A:80 X:80 Y:98 P:A5 SP:FB CYC:5336
D3B6  A9 80     LDA #$80                        A:80 X:80 Y:99 P:A5 SP:FB CYC:5338
D3B8  85 78     STA $78 = 00                    A:80 X:80 Y:99 P:A5 SP:FB CYC:5340
D3BA  E4 78     CPX $78 = 80                    A:80 X:80 Y:99 P:A5 SP:FB CYC:5343
D3BC  20 BF F8  JSR $F8BF                       A:80 X:80 Y:99 P:27 SP:FB CYC:5346
F8BF  D0 05     BNE $F8C6                       A:80 X:80 Y:99 P:27 SP:F9 CYC:5352
F8C1  30 03     BMI $F8C6                       A:80 X:80 Y:99 P:27 SP:F9 CYC:5354
F8C3  90 01     BCC $F8C6                       A:80 X:80 Y:99 P:27 SP:F9 CYC:5356
F8C5  60        RTS                             A:80 X:80 Y:99 P:27 SP:F9 CYC:5358
D3BF  C8        INY                             A:80 X:80 Y:99 P:27 SP:FB CYC:5364
D3C0  A9 81     LDA #$81                        A:80 X:80 Y:9A P:A5 SP:FB CYC:5366
D3C2  85 78     STA $78 = 80                    A:81 X:80 Y:9A P:A5 SP:FB CYC:5368
D3C4  E4 78     CPX $78 = 81                    A:81 X:80 Y:9A P:A5 SP:FB CYC:5371
D3C6  20 C9 F8  JSR $F8C9                       A:81 X:80 Y:9A P:A4 SP:FB CYC:5374
F8C9  B0 05     BCS $F8D0                       A:81 X:80 Y:9A P:A4 SP:F9 CYC:5380
F8CB  F0 03     BEQ $F8D0                       A:81 X:80 Y:9A P:A4 SP:F9 CYC:5382
F8CD  10 01     BPL $F8D0                       A:81 X:80 Y:9A P:A4 SP:F9 CYC:5384
F8CF  60        RTS                             A:81 X:80 Y:9A P:A4 SP:F9 CYC:5386
D3C9  C8        INY                             A:81 X:80 Y:9A P:A4 SP:FB CYC:5392
D3CA  A9 7F     LDA #$7F                        A:81 X:80 Y:9B P:A4 SP:FB CYC:5394
D3CC  85 78     STA $78 = 81                    A:7F X:80 Y:9B P:24 SP:FB CYC:5396
D3CE  E4 78     CPX $78 = 7F                    A:7F X:80 Y:9B P:24 SP:FB CYC:5399
D3D0  20 D3 F8  JSR $F8D3                       A:7F X:80 Y:9B P:25 SP:FB CYC:5402
F8D3  90 05     BCC $F8DA                       A:7F X:80 Y:9B P:25 SP:F9 CYC:5408
F8D5  F0 03     BEQ $F8DA                       A:7F X:80 Y:9B P:25 SP:F9 CYC:5410
F8D7  30 01     BMI $F8DA                       A:7F X:80 Y:9B P:25 SP:F9 CYC:5412
F8D9  60        RTS                             A:7F X:80 Y:9B P:25 SP:F9 CYC:5414
D3D3  C8        INY                             A:7F X:80 Y:9B P:25 SP:FB CYC:5420
D3D4  98        TYA                             A:7F X:80 Y:9C P:A5 SP:FB CYC:5422
D3D5  AA        TAX                             A:9C X:80 Y:9C P:A5 SP:FB CYC:5424
D3D6  A9 40     LDA #$40                        A:9C X:9C Y:9C P:A5 SP:FB CYC:5426
D3D8  85 78     STA $78 = 7F                    A:40 X:9C Y:9C P:25 SP:FB CYC:5428
D3DA  20 DD F8  JSR $F8DD                       A:40 X:9C Y:9C P:25 SP:FB CYC:5431
F8DD  24 01     BIT $01 = FF                    A:40 X:9C Y:9C P:25 SP:F9 CYC:5437
F8DF  A0 40     LDY #$40                        A:40 X:9C Y:9C P:E5 SP:F9 CYC:5440
F8E1  60        RTS                             A:40 X:9C Y:40 P:65 SP:F9 CYC:5442
D3DD  C4 78     CPY $78 = 40                    A:40 X:9C Y:40 P:65 SP:FB CYC:5448
D3DF  20 E2 F8  JSR $F8E2                       A:40 X:9C Y:40 P:67 SP:FB CYC:5451
F8E2  30 07     BMI $F8EB                       A:40 X:9C Y:40 P:67 SP:F9 CYC:5457
F8E4  90 05     BCC $F8EB                       A:40 X:9C Y:40 P:67 SP:F9 CYC:5459
F8E6  D0 03     BNE $F8EB                       A:40 X:9C Y:40 P:67 SP:F9 CYC:5461
F8E8  50 01     BVC $F8EB                       A:40 X:9C Y:40 P:67 SP:F9 CYC:5463
F8EA  60        RTS                             A:40 X:9C Y:40 P:67 SP:F9 CYC:5465
D3E2  E8        INX                             A:40 X:9C Y:40 P:67 SP:FB CYC:5471
D3E3  A9 3F     LDA #$3F                        A:40 X:9D Y:40 P:E5 SP:FB CYC:5473
D3E5  85 78     STA $78 = 40                    A:3F X:9D Y:40 P:65 SP:FB CYC:5475
D3E7  20 EE F8  JSR $F8EE                       A:3F X:9D Y:40 P:65 SP:FB CYC:5478
F8EE  B8        CLV                             A:3F X:9D Y:40 P:65 SP:F9 CYC:5484
F8EF  60        RTS                             A:3F X:9D Y:40 P:25 SP:F9 CYC:5486
D3EA  C4 78     CPY $78 = 3F                    A:3F X:9D Y:40 P:25 SP:FB CYC:5492
D3EC  20 F0 F8  JSR $F8F0                       A:3F X:9D Y:40 P:25 SP:FB CYC:5495
F8F0  F0 07     BEQ $F8F9                       A:3F X:9D Y:40 P:25 SP:F9 CYC:5501
F8F2  30 05     BMI $F8F9                       A:3F X:9D Y:40 P:25 SP:F9 CYC:5503
F8F4  90 03     BCC $F8F9                       A:3F X:9D Y:40 P:25 SP:F9 CYC:5505
F8F6  70 01     BVS $F8F9                       A:3F X:9D Y:40 P:25 SP:F9 CYC:5507
F8F8  60        RTS                             A:3F X:9D Y:40 P:25 SP:F9 CYC:5509
D3EF  E8        INX                             A:3F X:9D Y:40 P:25 SP:FB CYC:5515
D3F0  A9 41     LDA #$41                        A:3F X:9E Y:40 P:A5 SP:FB CYC:5517
D3F2  85 78     STA $78 = 3F                    A:41 X:9E Y:40 P:25 SP:FB CYC:5519
D3F4  C4 78     CPY $78 = 41                    A:41 X:9E Y:40 P:25 SP:FB CYC:5522
D3F6  20 FC F8  JSR $F8FC                       A:41 X:9E Y:40 P:A4 SP:FB CYC:5525
F8FC  F0 05     BEQ $F903                       A:41 X:9E Y:40 P:A4 SP:F9 CYC:5531
F8FE  10 03     BPL $F903                       A:41 X:9E Y:40 P:A4 SP:F9 CYC:5533
F900  10 01     BPL $F903                       A:41 X:9E Y:40 P:A4 SP:F9 CYC:5535
F902  60        RTS                             A:41 X:9E Y:40 P:A4 SP:F9 CYC:5537
D3F9  E8        INX                             A:41 X:9E Y:40 P:A4 SP:FB CYC:5543
D3FA  A9 00     LDA #$00                        A:41 X:9F Y:40 P:A4 SP:FB CYC:5545
D3FC  85 78     STA $78 = 41                    A:0 X:9F Y:40 P:26 SP:FB CYC:5547
D3FE  20 06 F9  JSR $F906                       A:0 X:9F Y:40 P:26 SP:FB CYC:5550
F906  A0 80     LDY #$80                        A:0 X:9F Y:40 P:26 SP:F9 CYC:5556
F908  60        RTS                             A:0 X:9F Y:80 P:A4 SP:F9 CYC:5558
D401  C4 78     CPY $78 = 00                    A:0 X:9F Y:80 P:A4 SP:FB CYC:5564
D403  20 09 F9  JSR $F909                       A:0 X:9F Y:80 P:A5 SP:FB CYC:5567
F909  F0 05     BEQ $F910                       A:0 X:9F Y:80 P:A5 SP:F9 CYC:5573
F90B  10 03     BPL $F910                       A:0 X:9F Y:80 P:A5 SP:F9 CYC:5575
F90D  90 01     BCC $F910                       A:0 X:9F Y:80 P:A5 SP:F9 CYC:5577
F90F  60        RTS                             A:0 X:9F Y:80 P:A5 SP:F9 CYC:5579
D406  E8        INX                             A:0 X:9F Y:80 P:A5 SP:FB CYC:5585
D407  A9 80     LDA #$80                        A:0 X:A0 Y:80 P:A5 SP:FB CYC:5587
D409  85 78     STA $78 = 00                    A:80 X:A0 Y:80 P:A5 SP:FB CYC:5589
D40B  C4 78     CPY $78 = 80                    A:80 X:A0 Y:80 P:A5 SP:FB CYC:5592
D40D  20 13 F9  JSR $F913                       A:80 X:A0 Y:80 P:27 SP:FB CYC:5595
F913  D0 05     BNE $F91A                       A:80 X:A0 Y:80 P:27 SP:F9 CYC:5601
F915  30 03     BMI $F91A                       A:80 X:A0 Y:80 P:27 SP:F9 CYC:5603
F917  90 01     BCC $F91A                       A:80 X:A0 Y:80 P:27 SP:F9 CYC:5605
F919  60        RTS                             A:80 X:A0 Y:80 P:27 SP:F9 CYC:5607
D410  E8        INX                             A:80 X:A0 Y:80 P:27 SP:FB CYC:5613
D411  A9 81     LDA #$81                        A:80 X:A1 Y:80 P:A5 SP:FB CYC:5615
D413  85 78     STA $78 = 80                    A:81 X:A1 Y:80 P:A5 SP:FB CYC:5617
D415  C4 78     CPY $78 = 81                    A:81 X:A1 Y:80 P:A5 SP:FB CYC:5620
D417  20 1D F9  JSR $F91D                       A:81 X:A1 Y:80 P:A4 SP:FB CYC:5623
F91D  B0 05     BCS $F924                       A:81 X:A1 Y:80 P:A4 SP:F9 CYC:5629
F91F  F0 03     BEQ $F924                       A:81 X:A1 Y:80 P:A4 SP:F9 CYC:5631
F921  10 01     BPL $F924                       A:81 X:A1 Y:80 P:A4 SP:F9 CYC:5633
F923  60        RTS                             A:81 X:A1 Y:80 P:A4 SP:F9 CYC:5635
D41A  E8        INX                             A:81 X:A1 Y:80 P:A4 SP:FB CYC:5641
D41B  A9 7F     LDA #$7F                        A:81 X:A2 Y:80 P:A4 SP:FB CYC:5643
D41D  85 78     STA $78 = 81                    A:7F X:A2 Y:80 P:24 SP:FB CYC:5645
D41F  C4 78     CPY $78 = 7F                    A:7F X:A2 Y:80 P:24 SP:FB CYC:5648
D421  20 27 F9  JSR $F927                       A:7F X:A2 Y:80 P:25 SP:FB CYC:5651
F927  90 05     BCC $F92E                       A:7F X:A2 Y:80 P:25 SP:F9 CYC:5657
F929  F0 03     BEQ $F92E                       A:7F X:A2 Y:80 P:25 SP:F9 CYC:5659
F92B  30 01     BMI $F92E                       A:7F X:A2 Y:80 P:25 SP:F9 CYC:5661
F92D  60        RTS                             A:7F X:A2 Y:80 P:25 SP:F9 CYC:5663
D424  E8        INX                             A:7F X:A2 Y:80 P:25 SP:FB CYC:5669
D425  8A        TXA                             A:7F X:A3 Y:80 P:A5 SP:FB CYC:5671
D426  A8        TAY                             A:A3 X:A3 Y:80 P:A5 SP:FB CYC:5673
D427  20 90 F9  JSR $F990                       A:A3 X:A3 Y:A3 P:A5 SP:FB CYC:5675
F990  A2 55     LDX #$55                        A:A3 X:A3 Y:A3 P:A5 SP:F9 CYC:5681
F992  A9 FF     LDA #$FF                        A:A3 X:55 Y:A3 P:25 SP:F9 CYC:5683
F994  85 01     STA $01 = FF                    A:FF X:55 Y:A3 P:A5 SP:F9 CYC:5685
F996  EA        NOP                             A:FF X:55 Y:A3 P:A5 SP:F9 CYC:5688
F997  24 01     BIT $01 = FF                    A:FF X:55 Y:A3 P:A5 SP:F9 CYC:5690
F999  38        SEC                             A:FF X:55 Y:A3 P:E5 SP:F9 CYC:5693
F99A  A9 01     LDA #$01                        A:FF X:55 Y:A3 P:E5 SP:F9 CYC:5695
F99C  60        RTS                             A:1 X:55 Y:A3 P:65 SP:F9 CYC:5697
D42A  85 78     STA $78 = 7F                    A:1 X:55 Y:A3 P:65 SP:FB CYC:5703
D42C  46 78     LSR $78 = 01                    A:1 X:55 Y:A3 P:65 SP:FB CYC:5706
D42E  A5 78     LDA $78 = 00                    A:1 X:55 Y:A3 P:67 SP:FB CYC:5711
D430  20 9D F9  JSR $F99D                       A:0 X:55 Y:A3 P:67 SP:FB CYC:5714
F99D  90 1B     BCC $F9BA                       A:0 X:55 Y:A3 P:67 SP:F9 CYC:5720
F99F  D0 19     BNE $F9BA                       A:0 X:55 Y:A3 P:67 SP:F9 CYC:5722
F9A1  30 17     BMI $F9BA                       A:0 X:55 Y:A3 P:67 SP:F9 CYC:5724
F9A3  50 15     BVC $F9BA                       A:0 X:55 Y:A3 P:67 SP:F9 CYC:5726
F9A5  C9 00     CMP #$00                        A:0 X:55 Y:A3 P:67 SP:F9 CYC:5728
F9A7  D0 11     BNE $F9BA                       A:0 X:55 Y:A3 P:67 SP:F9 CYC:5730
F9A9  B8        CLV                             A:0 X:55 Y:A3 P:67 SP:F9 CYC:5732
F9AA  A9 AA     LDA #$AA                        A:0 X:55 Y:A3 P:27 SP:F9 CYC:5734
F9AC  60        RTS                             A:AA X:55 Y:A3 P:A5 SP:F9 CYC:5736
D433  C8        INY                             A:AA X:55 Y:A3 P:A5 SP:FB CYC:5742
D434  85 78     STA $78 = 00                    A:AA X:55 Y:A4 P:A5 SP:FB CYC:5744
D436  46 78     LSR $78 = AA                    A:AA X:55 Y:A4 P:A5 SP:FB CYC:5747
D438  A5 78     LDA $78 = 55                    A:AA X:55 Y:A4 P:24 SP:FB CYC:5752
D43A  20 AD F9  JSR $F9AD                       A:55 X:55 Y:A4 P:24 SP:FB CYC:5755
F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9 CYC:5761
F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9 CYC:5763
F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9 CYC:5765
F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9 CYC:5767
F9B5  C9 55     CMP #$55                        A:55 X:55 Y:A4 P:24 SP:F9 CYC:5769
F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:A4 P:27 SP:F9 CYC:5771
F9B9  60        RTS                             A:55 X:55 Y:A4 P:27 SP:F9 CYC:5773
D43D  C8        INY                             A:55 X:55 Y:A4 P:27 SP:FB CYC:5779
D43E  20 BD F9  JSR $F9BD                       A:55 X:55 Y:A5 P:A5 SP:FB CYC:5781
F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:A5 P:A5 SP:F9 CYC:5787
F9BF  38        SEC                             A:55 X:55 Y:A5 P:E5 SP:F9 CYC:5790
F9C0  A9 80     LDA #$80                        A:55 X:55 Y:A5 P:E5 SP:F9 CYC:5792
F9C2  60        RTS                             A:80 X:55 Y:A5 P:E5 SP:F9 CYC:5794
D441  85 78     STA $78 = 55                    A:80 X:55 Y:A5 P:E5 SP:FB CYC:5800
D443  06 78     ASL $78 = 80                    A:80 X:55 Y:A5 P:E5 SP:FB CYC:5803
D445  A5 78     LDA $78 = 00                    A:80 X:55 Y:A5 P:67 SP:FB CYC:5808
D447  20 C3 F9  JSR $F9C3                       A:0 X:55 Y:A5 P:67 SP:FB CYC:5811
F9C3  90 1C     BCC $F9E1                       A:0 X:55 Y:A5 P:67 SP:F9 CYC:5817
F9C5  D0 1A     BNE $F9E1                       A:0 X:55 Y:A5 P:67 SP:F9 CYC:5819
F9C7  30 18     BMI $F9E1                       A:0 X:55 Y:A5 P:67 SP:F9 CYC:5821
F9C9  50 16     BVC $F9E1                       A:0 X:55 Y:A5 P:67 SP:F9 CYC:5823
F9CB  C9 00     CMP #$00                        A:0 X:55 Y:A5 P:67 SP:F9 CYC:5825
F9CD  D0 12     BNE $F9E1                       A:0 X:55 Y:A5 P:67 SP:F9 CYC:5827
F9CF  B8        CLV                             A:0 X:55 Y:A5 P:67 SP:F9 CYC:5829
F9D0  A9 55     LDA #$55                        A:0 X:55 Y:A5 P:27 SP:F9 CYC:5831
F9D2  38        SEC                             A:55 X:55 Y:A5 P:25 SP:F9 CYC:5833
F9D3  60        RTS                             A:55 X:55 Y:A5 P:25 SP:F9 CYC:5835
D44A  C8        INY                             A:55 X:55 Y:A5 P:25 SP:FB CYC:5841
D44B  85 78     STA $78 = 00                    A:55 X:55 Y:A6 P:A5 SP:FB CYC:5843
D44D  06 78     ASL $78 = 55                    A:55 X:55 Y:A6 P:A5 SP:FB CYC:5846
D44F  A5 78     LDA $78 = AA                    A:55 X:55 Y:A6 P:A4 SP:FB CYC:5851
D451  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:A6 P:A4 SP:FB CYC:5854
F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9 CYC:5860
F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9 CYC:5862
F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9 CYC:5864
F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9 CYC:5866
F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:A6 P:A4 SP:F9 CYC:5868
F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:A6 P:27 SP:F9 CYC:5870
F9E0  60        RTS                             A:AA X:55 Y:A6 P:27 SP:F9 CYC:5872
D454  C8        INY                             A:AA X:55 Y:A6 P:27 SP:FB CYC:5878
D455  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:A7 P:A5 SP:FB CYC:5880
F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:A7 P:A5 SP:F9 CYC:5886
F9E6  38        SEC                             A:AA X:55 Y:A7 P:E5 SP:F9 CYC:5889
F9E7  A9 01     LDA #$01                        A:AA X:55 Y:A7 P:E5 SP:F9 CYC:5891
F9E9  60        RTS                             A:1 X:55 Y:A7 P:65 SP:F9 CYC:5893
D458  85 78     STA $78 = AA                    A:1 X:55 Y:A7 P:65 SP:FB CYC:5899
D45A  66 78     ROR $78 = 01                    A:1 X:55 Y:A7 P:65 SP:FB CYC:5902
D45C  A5 78     LDA $78 = 80                    A:1 X:55 Y:A7 P:E5 SP:FB CYC:5907
D45E  20 EA F9  JSR $F9EA                       A:80 X:55 Y:A7 P:E5 SP:FB CYC:5910
F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9 CYC:5916
F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9 CYC:5918
F9EE  10 18     BPL $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9 CYC:5920
F9F0  50 16     BVC $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9 CYC:5922
F9F2  C9 80     CMP #$80                        A:80 X:55 Y:A7 P:E5 SP:F9 CYC:5924
F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:A7 P:67 SP:F9 CYC:5926
F9F6  B8        CLV                             A:80 X:55 Y:A7 P:67 SP:F9 CYC:5928
F9F7  18        CLC                             A:80 X:55 Y:A7 P:27 SP:F9 CYC:5930
F9F8  A9 55     LDA #$55                        A:80 X:55 Y:A7 P:26 SP:F9 CYC:5932
F9FA  60        RTS                             A:55 X:55 Y:A7 P:24 SP:F9 CYC:5934
D461  C8        INY                             A:55 X:55 Y:A7 P:24 SP:FB CYC:5940
D462  85 78     STA $78 = 80                    A:55 X:55 Y:A8 P:A4 SP:FB CYC:5942
D464  66 78     ROR $78 = 55                    A:55 X:55 Y:A8 P:A4 SP:FB CYC:5945
D466  A5 78     LDA $78 = 2A                    A:55 X:55 Y:A8 P:25 SP:FB CYC:5950
D468  20 FB F9  JSR $F9FB                       A:2A X:55 Y:A8 P:25 SP:FB CYC:5953
F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:A8 P:25 SP:F9 CYC:5959
F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:A8 P:25 SP:F9 CYC:5961
F9FF  30 07     BMI $FA08                       A:2A X:55 Y:A8 P:25 SP:F9 CYC:5963
FA01  70 05     BVS $FA08                       A:2A X:55 Y:A8 P:25 SP:F9 CYC:5965
FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:A8 P:25 SP:F9 CYC:5967
FA05  D0 01     BNE $FA08                       A:2A X:55 Y:A8 P:27 SP:F9 CYC:5969
FA07  60        RTS                             A:2A X:55 Y:A8 P:27 SP:F9 CYC:5971
D46B  C8        INY                             A:2A X:55 Y:A8 P:27 SP:FB CYC:5977
D46C  20 0A FA  JSR $FA0A                       A:2A X:55 Y:A9 P:A5 SP:FB CYC:5979
FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:A9 P:A5 SP:F9 CYC:5985
FA0C  38        SEC                             A:2A X:55 Y:A9 P:E5 SP:F9 CYC:5988
FA0D  A9 80     LDA #$80                        A:2A X:55 Y:A9 P:E5 SP:F9 CYC:5990
FA0F  60        RTS                             A:80 X:55 Y:A9 P:E5 SP:F9 CYC:5992
D46F  85 78     STA $78 = 2A                    A:80 X:55 Y:A9 P:E5 SP:FB CYC:5998
D471  26 78     ROL $78 = 80                    A:80 X:55 Y:A9 P:E5 SP:FB CYC:6001
D473  A5 78     LDA $78 = 01                    A:80 X:55 Y:A9 P:65 SP:FB CYC:6006
D475  20 10 FA  JSR $FA10                       A:1 X:55 Y:A9 P:65 SP:FB CYC:6009
FA10  90 1C     BCC $FA2E                       A:1 X:55 Y:A9 P:65 SP:F9 CYC:6015
FA12  F0 1A     BEQ $FA2E                       A:1 X:55 Y:A9 P:65 SP:F9 CYC:6017
FA14  30 18     BMI $FA2E                       A:1 X:55 Y:A9 P:65 SP:F9 CYC:6019
FA16  50 16     BVC $FA2E                       A:1 X:55 Y:A9 P:65 SP:F9 CYC:6021
FA18  C9 01     CMP #$01                        A:1 X:55 Y:A9 P:65 SP:F9 CYC:6023
FA1A  D0 12     BNE $FA2E                       A:1 X:55 Y:A9 P:67 SP:F9 CYC:6025
FA1C  B8        CLV                             A:1 X:55 Y:A9 P:67 SP:F9 CYC:6027
FA1D  18        CLC                             A:1 X:55 Y:A9 P:27 SP:F9 CYC:6029
FA1E  A9 55     LDA #$55                        A:1 X:55 Y:A9 P:26 SP:F9 CYC:6031
FA20  60        RTS                             A:55 X:55 Y:A9 P:24 SP:F9 CYC:6033
D478  C8        INY                             A:55 X:55 Y:A9 P:24 SP:FB CYC:6039
D479  85 78     STA $78 = 01                    A:55 X:55 Y:AA P:A4 SP:FB CYC:6041
D47B  26 78     ROL $78 = 55                    A:55 X:55 Y:AA P:A4 SP:FB CYC:6044
D47D  A5 78     LDA $78 = AA                    A:55 X:55 Y:AA P:A4 SP:FB CYC:6049
D47F  20 21 FA  JSR $FA21                       A:AA X:55 Y:AA P:A4 SP:FB CYC:6052
FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9 CYC:6058
FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9 CYC:6060
FA25  10 07     BPL $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9 CYC:6062
FA27  70 05     BVS $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9 CYC:6064
FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:AA P:A4 SP:F9 CYC:6066
FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:AA P:27 SP:F9 CYC:6068
FA2D  60        RTS                             A:AA X:55 Y:AA P:27 SP:F9 CYC:6070
D482  A9 FF     LDA #$FF                        A:AA X:55 Y:AA P:27 SP:FB CYC:6076
D484  85 78     STA $78 = AA                    A:FF X:55 Y:AA P:A5 SP:FB CYC:6078
D486  85 01     STA $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB CYC:6081
D488  24 01     BIT $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB CYC:6084
D48A  38        SEC                             A:FF X:55 Y:AA P:E5 SP:FB CYC:6087
D48B  E6 78     INC $78 = FF                    A:FF X:55 Y:AA P:E5 SP:FB CYC:6089
D48D  D0 0C     BNE $D49B                       A:FF X:55 Y:AA P:67 SP:FB CYC:6094
D48F  30 0A     BMI $D49B                       A:FF X:55 Y:AA P:67 SP:FB CYC:6096
D491  50 08     BVC $D49B                       A:FF X:55 Y:AA P:67 SP:FB CYC:6098
D493  90 06     BCC $D49B                       A:FF X:55 Y:AA P:67 SP:FB CYC:6100
D495  A5 78     LDA $78 = 00                    A:FF X:55 Y:AA P:67 SP:FB CYC:6102
D497  C9 00     CMP #$00                        A:0 X:55 Y:AA P:67 SP:FB CYC:6105
D499  F0 04     BEQ $D49F                       A:0 X:55 Y:AA P:67 SP:FB CYC:6107
D49F  A9 7F     LDA #$7F                        A:0 X:55 Y:AA P:67 SP:FB CYC:6110
D4A1  85 78     STA $78 = 00                    A:7F X:55 Y:AA P:65 SP:FB CYC:6112
D4A3  B8        CLV                             A:7F X:55 Y:AA P:65 SP:FB CYC:6115
D4A4  18        CLC                             A:7F X:55 Y:AA P:25 SP:FB CYC:6117
D4A5  E6 78     INC $78 = 7F                    A:7F X:55 Y:AA P:24 SP:FB CYC:6119
D4A7  F0 0C     BEQ $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB CYC:6124
D4A9  10 0A     BPL $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB CYC:6126
D4AB  70 08     BVS $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB CYC:6128
D4AD  B0 06     BCS $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB CYC:6130
D4AF  A5 78     LDA $78 = 80                    A:7F X:55 Y:AA P:A4 SP:FB CYC:6132
D4B1  C9 80     CMP #$80                        A:80 X:55 Y:AA P:A4 SP:FB CYC:6135
D4B3  F0 04     BEQ $D4B9                       A:80 X:55 Y:AA P:27 SP:FB CYC:6137
D4B9  A9 00     LDA #$00                        A:80 X:55 Y:AA P:27 SP:FB CYC:6140
D4BB  85 78     STA $78 = 80                    A:0 X:55 Y:AA P:27 SP:FB CYC:6142
D4BD  24 01     BIT $01 = FF                    A:0 X:55 Y:AA P:27 SP:FB CYC:6145
D4BF  38        SEC                             A:0 X:55 Y:AA P:E7 SP:FB CYC:6148
D4C0  C6 78     DEC $78 = 00                    A:0 X:55 Y:AA P:E7 SP:FB CYC:6150
D4C2  F0 0C     BEQ $D4D0                       A:0 X:55 Y:AA P:E5 SP:FB CYC:6155
D4C4  10 0A     BPL $D4D0                       A:0 X:55 Y:AA P:E5 SP:FB CYC:6157
D4C6  50 08     BVC $D4D0                       A:0 X:55 Y:AA P:E5 SP:FB CYC:6159
D4C8  90 06     BCC $D4D0                       A:0 X:55 Y:AA P:E5 SP:FB CYC:6161
D4CA  A5 78     LDA $78 = FF                    A:0 X:55 Y:AA P:E5 SP:FB CYC:6163
D4CC  C9 FF     CMP #$FF                        A:FF X:55 Y:AA P:E5 SP:FB CYC:6166
D4CE  F0 04     BEQ $D4D4                       A:FF X:55 Y:AA P:67 SP:FB CYC:6168
D4D4  A9 80     LDA #$80                        A:FF X:55 Y:AA P:67 SP:FB CYC:6171
D4D6  85 78     STA $78 = FF                    A:80 X:55 Y:AA P:E5 SP:FB CYC:6173
D4D8  B8        CLV                             A:80 X:55 Y:AA P:E5 SP:FB CYC:6176
D4D9  18        CLC                             A:80 X:55 Y:AA P:A5 SP:FB CYC:6178
D4DA  C6 78     DEC $78 = 80                    A:80 X:55 Y:AA P:A4 SP:FB CYC:6180
D4DC  F0 0C     BEQ $D4EA                       A:80 X:55 Y:AA P:24 SP:FB CYC:6185
D4DE  30 0A     BMI $D4EA                       A:80 X:55 Y:AA P:24 SP:FB CYC:6187
D4E0  70 08     BVS $D4EA                       A:80 X:55 Y:AA P:24 SP:FB CYC:6189
D4E2  B0 06     BCS $D4EA                       A:80 X:55 Y:AA P:24 SP:FB CYC:6191
D4E4  A5 78     LDA $78 = 7F                    A:80 X:55 Y:AA P:24 SP:FB CYC:6193
D4E6  C9 7F     CMP #$7F                        A:7F X:55 Y:AA P:24 SP:FB CYC:6196
D4E8  F0 04     BEQ $D4EE                       A:7F X:55 Y:AA P:27 SP:FB CYC:6198
D4EE  A9 01     LDA #$01                        A:7F X:55 Y:AA P:27 SP:FB CYC:6201
D4F0  85 78     STA $78 = 7F                    A:1 X:55 Y:AA P:25 SP:FB CYC:6203
D4F2  C6 78     DEC $78 = 01                    A:1 X:55 Y:AA P:25 SP:FB CYC:6206
D4F4  F0 04     BEQ $D4FA                       A:1 X:55 Y:AA P:27 SP:FB CYC:6211
D4FA  60        RTS                             A:1 X:55 Y:AA P:27 SP:FB CYC:6214
C615  20 FB D4  JSR $D4FB                       A:1 X:55 Y:AA P:27 SP:FD CYC:6220
D4FB  A9 55     LDA #$55                        A:1 X:55 Y:AA P:27 SP:FB CYC:6226
D4FD  8D 78 06  STA $0678 = 00                  A:55 X:55 Y:AA P:25 SP:FB CYC:6228
D500  A9 FF     LDA #$FF                        A:55 X:55 Y:AA P:25 SP:FB CYC:6232
D502  85 01     STA $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB CYC:6234
D504  24 01     BIT $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB CYC:6237
D506  A0 11     LDY #$11                        A:FF X:55 Y:AA P:E5 SP:FB CYC:6240
D508  A2 23     LDX #$23                        A:FF X:55 Y:11 P:65 SP:FB CYC:6242
D50A  A9 00     LDA #$00                        A:FF X:23 Y:11 P:65 SP:FB CYC:6244
D50C  AD 78 06  LDA $0678 = 55                  A:0 X:23 Y:11 P:67 SP:FB CYC:6246
D50F  F0 10     BEQ $D521                       A:55 X:23 Y:11 P:65 SP:FB CYC:6250
D511  30 0E     BMI $D521                       A:55 X:23 Y:11 P:65 SP:FB CYC:6252
D513  C9 55     CMP #$55                        A:55 X:23 Y:11 P:65 SP:FB CYC:6254
D515  D0 0A     BNE $D521                       A:55 X:23 Y:11 P:67 SP:FB CYC:6256
D517  C0 11     CPY #$11                        A:55 X:23 Y:11 P:67 SP:FB CYC:6258
D519  D0 06     BNE $D521                       A:55 X:23 Y:11 P:67 SP:FB CYC:6260
D51B  E0 23     CPX #$23                        A:55 X:23 Y:11 P:67 SP:FB CYC:6262
D51D  50 02     BVC $D521                       A:55 X:23 Y:11 P:67 SP:FB CYC:6264
D51F  F0 04     BEQ $D525                       A:55 X:23 Y:11 P:67 SP:FB CYC:6266
D525  A9 46     LDA #$46                        A:55 X:23 Y:11 P:67 SP:FB CYC:6269
D527  24 01     BIT $01 = FF                    A:46 X:23 Y:11 P:65 SP:FB CYC:6271
D529  8D 78 06  STA $0678 = 55                  A:46 X:23 Y:11 P:E5 SP:FB CYC:6274
D52C  F0 0B     BEQ $D539                       A:46 X:23 Y:11 P:E5 SP:FB CYC:6278
D52E  10 09     BPL $D539                       A:46 X:23 Y:11 P:E5 SP:FB CYC:6280
D530  50 07     BVC $D539                       A:46 X:23 Y:11 P:E5 SP:FB CYC:6282
D532  AD 78 06  LDA $0678 = 46                  A:46 X:23 Y:11 P:E5 SP:FB CYC:6284
D535  C9 46     CMP #$46                        A:46 X:23 Y:11 P:65 SP:FB CYC:6288
D537  F0 04     BEQ $D53D                       A:46 X:23 Y:11 P:67 SP:FB CYC:6290
D53D  A9 55     LDA #$55                        A:46 X:23 Y:11 P:67 SP:FB CYC:6293
D53F  8D 78 06  STA $0678 = 46                  A:55 X:23 Y:11 P:65 SP:FB CYC:6295
D542  24 01     BIT $01 = FF                    A:55 X:23 Y:11 P:65 SP:FB CYC:6299
D544  A9 11     LDA #$11                        A:55 X:23 Y:11 P:E5 SP:FB CYC:6302
D546  A2 23     LDX #$23                        A:11 X:23 Y:11 P:65 SP:FB CYC:6304
D548  A0 00     LDY #$00                        A:11 X:23 Y:11 P:65 SP:FB CYC:6306
D54A  AC 78 06  LDY $0678 = 55                  A:11 X:23 Y:0 P:67 SP:FB CYC:6308
D54D  F0 10     BEQ $D55F                       A:11 X:23 Y:55 P:65 SP:FB CYC:6312
D54F  30 0E     BMI $D55F                       A:11 X:23 Y:55 P:65 SP:FB CYC:6314
D551  C0 55     CPY #$55                        A:11 X:23 Y:55 P:65 SP:FB CYC:6316
D553  D0 0A     BNE $D55F                       A:11 X:23 Y:55 P:67 SP:FB CYC:6318
D555  C9 11     CMP #$11                        A:11 X:23 Y:55 P:67 SP:FB CYC:6320
D557  D0 06     BNE $D55F                       A:11 X:23 Y:55 P:67 SP:FB CYC:6322
D559  E0 23     CPX #$23                        A:11 X:23 Y:55 P:67 SP:FB CYC:6324
D55B  50 02     BVC $D55F                       A:11 X:23 Y:55 P:67 SP:FB CYC:6326
D55D  F0 04     BEQ $D563                       A:11 X:23 Y:55 P:67 SP:FB CYC:6328
D563  A0 46     LDY #$46                        A:11 X:23 Y:55 P:67 SP:FB CYC:6331
D565  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:65 SP:FB CYC:6333
D567  8C 78 06  STY $0678 = 55                  A:11 X:23 Y:46 P:E5 SP:FB CYC:6336
D56A  F0 0B     BEQ $D577                       A:11 X:23 Y:46 P:E5 SP:FB CYC:6340
D56C  10 09     BPL $D577                       A:11 X:23 Y:46 P:E5 SP:FB CYC:6342
D56E  50 07     BVC $D577                       A:11 X:23 Y:46 P:E5 SP:FB CYC:6344
D570  AC 78 06  LDY $0678 = 46                  A:11 X:23 Y:46 P:E5 SP:FB CYC:6346
D573  C0 46     CPY #$46                        A:11 X:23 Y:46 P:65 SP:FB CYC:6350
D575  F0 04     BEQ $D57B                       A:11 X:23 Y:46 P:67 SP:FB CYC:6352
D57B  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:67 SP:FB CYC:6355
D57D  A9 55     LDA #$55                        A:11 X:23 Y:46 P:E5 SP:FB CYC:6358
D57F  8D 78 06  STA $0678 = 46                  A:55 X:23 Y:46 P:65 SP:FB CYC:6360
D582  A0 11     LDY #$11                        A:55 X:23 Y:46 P:65 SP:FB CYC:6364
D584  A9 23     LDA #$23                        A:55 X:23 Y:11 P:65 SP:FB CYC:6366
D586  A2 00     LDX #$00                        A:23 X:23 Y:11 P:65 SP:FB CYC:6368
D588  AE 78 06  LDX $0678 = 55                  A:23 X:0 Y:11 P:67 SP:FB CYC:6370
D58B  F0 10     BEQ $D59D                       A:23 X:55 Y:11 P:65 SP:FB CYC:6374
D58D  30 0E     BMI $D59D                       A:23 X:55 Y:11 P:65 SP:FB CYC:6376
D58F  E0 55     CPX #$55                        A:23 X:55 Y:11 P:65 SP:FB CYC:6378
D591  D0 0A     BNE $D59D                       A:23 X:55 Y:11 P:67 SP:FB CYC:6380
D593  C0 11     CPY #$11                        A:23 X:55 Y:11 P:67 SP:FB CYC:6382
D595  D0 06     BNE $D59D                       A:23 X:55 Y:11 P:67 SP:FB CYC:6384
D597  C9 23     CMP #$23                        A:23 X:55 Y:11 P:67 SP:FB CYC:6386
D599  50 02     BVC $D59D                       A:23 X:55 Y:11 P:67 SP:FB CYC:6388
D59B  F0 04     BEQ $D5A1                       A:23 X:55 Y:11 P:67 SP:FB CYC:6390
D5A1  A2 46     LDX #$46                        A:23 X:55 Y:11 P:67 SP:FB CYC:6393
D5A3  24 01     BIT $01 = FF                    A:23 X:46 Y:11 P:65 SP:FB CYC:6395
D5A5  8E 78 06  STX $0678 = 55                  A:23 X:46 Y:11 P:E5 SP:FB CYC:6398
D5A8  F0 0B     BEQ $D5B5                       A:23 X:46 Y:11 P:E5 SP:FB CYC:6402
D5AA  10 09     BPL $D5B5                       A:23 X:46 Y:11 P:E5 SP:FB CYC:6404
D5AC  50 07     BVC $D5B5                       A:23 X:46 Y:11 P:E5 SP:FB CYC:6406
D5AE  AE 78 06  LDX $0678 = 46                  A:23 X:46 Y:11 P:E5 SP:FB CYC:6408
D5B1  E0 46     CPX #$46                        A:23 X:46 Y:11 P:65 SP:FB CYC:6412
D5B3  F0 04     BEQ $D5B9                       A:23 X:46 Y:11 P:67 SP:FB CYC:6414
D5B9  A9 C0     LDA #$C0                        A:23 X:46 Y:11 P:67 SP:FB CYC:6417
D5BB  8D 78 06  STA $0678 = 46                  A:C0 X:46 Y:11 P:E5 SP:FB CYC:6419
D5BE  A2 33     LDX #$33                        A:C0 X:46 Y:11 P:E5 SP:FB CYC:6423
D5C0  A0 88     LDY #$88                        A:C0 X:33 Y:11 P:65 SP:FB CYC:6425
D5C2  A9 05     LDA #$05                        A:C0 X:33 Y:88 P:E5 SP:FB CYC:6427
D5C4  2C 78 06  BIT $0678 = C0                  A:5 X:33 Y:88 P:65 SP:FB CYC:6429
D5C7  10 10     BPL $D5D9                       A:5 X:33 Y:88 P:E7 SP:FB CYC:6433
D5C9  50 0E     BVC $D5D9                       A:5 X:33 Y:88 P:E7 SP:FB CYC:6435
D5CB  D0 0C     BNE $D5D9                       A:5 X:33 Y:88 P:E7 SP:FB CYC:6437
D5CD  C9 05     CMP #$05                        A:5 X:33 Y:88 P:E7 SP:FB CYC:6439
D5CF  D0 08     BNE $D5D9                       A:5 X:33 Y:88 P:67 SP:FB CYC:6441
D5D1  E0 33     CPX #$33                        A:5 X:33 Y:88 P:67 SP:FB CYC:6443
D5D3  D0 04     BNE $D5D9                       A:5 X:33 Y:88 P:67 SP:FB CYC:6445
D5D5  C0 88     CPY #$88                        A:5 X:33 Y:88 P:67 SP:FB CYC:6447
D5D7  F0 04     BEQ $D5DD                       A:5 X:33 Y:88 P:67 SP:FB CYC:6449
D5DD  A9 03     LDA #$03                        A:5 X:33 Y:88 P:67 SP:FB CYC:6452
D5DF  8D 78 06  STA $0678 = C0                  A:3 X:33 Y:88 P:65 SP:FB CYC:6454
D5E2  A9 01     LDA #$01                        A:3 X:33 Y:88 P:65 SP:FB CYC:6458
D5E4  2C 78 06  BIT $0678 = 03                  A:1 X:33 Y:88 P:65 SP:FB CYC:6460
D5E7  30 08     BMI $D5F1                       A:1 X:33 Y:88 P:25 SP:FB CYC:6464
D5E9  70 06     BVS $D5F1                       A:1 X:33 Y:88 P:25 SP:FB CYC:6466
D5EB  F0 04     BEQ $D5F1                       A:1 X:33 Y:88 P:25 SP:FB CYC:6468
D5ED  C9 01     CMP #$01                        A:1 X:33 Y:88 P:25 SP:FB CYC:6470
D5EF  F0 04     BEQ $D5F5                       A:1 X:33 Y:88 P:27 SP:FB CYC:6472
D5F5  A0 B8     LDY #$B8                        A:1 X:33 Y:88 P:27 SP:FB CYC:6475
D5F7  A9 AA     LDA #$AA                        A:1 X:33 Y:B8 P:A5 SP:FB CYC:6477
D5F9  8D 78 06  STA $0678 = 03                  A:AA X:33 Y:B8 P:A5 SP:FB CYC:6479
D5FC  20 B6 F7  JSR $F7B6                       A:AA X:33 Y:B8 P:A5 SP:FB CYC:6483
F7B6  18        CLC                             A:AA X:33 Y:B8 P:A5 SP:F9 CYC:6489
F7B7  A9 FF     LDA #$FF                        A:AA X:33 Y:B8 P:A4 SP:F9 CYC:6491
F7B9  85 01     STA $01 = FF                    A:FF X:33 Y:B8 P:A4 SP:F9 CYC:6493
F7BB  24 01     BIT $01 = FF                    A:FF X:33 Y:B8 P:A4 SP:F9 CYC:6496
F7BD  A9 55     LDA #$55                        A:FF X:33 Y:B8 P:E4 SP:F9 CYC:6499
F7BF  60        RTS                             A:55 X:33 Y:B8 P:64 SP:F9 CYC:6501
D5FF  0D 78 06  ORA $0678 = AA                  A:55 X:33 Y:B8 P:64 SP:FB CYC:6507
D602  20 C0 F7  JSR $F7C0                       A:FF X:33 Y:B8 P:E4 SP:FB CYC:6511
F7C0  B0 09     BCS $F7CB                       A:FF X:33 Y:B8 P:E4 SP:F9 CYC:6517
F7C2  10 07     BPL $F7CB                       A:FF X:33 Y:B8 P:E4 SP:F9 CYC:6519
F7C4  C9 FF     CMP #$FF                        A:FF X:33 Y:B8 P:E4 SP:F9 CYC:6521
F7C6  D0 03     BNE $F7CB                       A:FF X:33 Y:B8 P:67 SP:F9 CYC:6523
F7C8  50 01     BVC $F7CB                       A:FF X:33 Y:B8 P:67 SP:F9 CYC:6525
F7CA  60        RTS                             A:FF X:33 Y:B8 P:67 SP:F9 CYC:6527
D605  C8        INY                             A:FF X:33 Y:B8 P:67 SP:FB CYC:6533
D606  A9 00     LDA #$00                        A:FF X:33 Y:B9 P:E5 SP:FB CYC:6535
D608  8D 78 06  STA $0678 = AA                  A:0 X:33 Y:B9 P:67 SP:FB CYC:6537
D60B  20 CE F7  JSR $F7CE                       A:0 X:33 Y:B9 P:67 SP:FB CYC:6541
F7CE  38        SEC                             A:0 X:33 Y:B9 P:67 SP:F9 CYC:6547
F7CF  B8        CLV                             A:0 X:33 Y:B9 P:67 SP:F9 CYC:6549
F7D0  A9 00     LDA #$00                        A:0 X:33 Y:B9 P:27 SP:F9 CYC:6551
F7D2  60        RTS                             A:0 X:33 Y:B9 P:27 SP:F9 CYC:6553
D60E  0D 78 06  ORA $0678 = 00                  A:0 X:33 Y:B9 P:27 SP:FB CYC:6559
D611  20 D3 F7  JSR $F7D3                       A:0 X:33 Y:B9 P:27 SP:FB CYC:6563
F7D3  D0 07     BNE $F7DC                       A:0 X:33 Y:B9 P:27 SP:F9 CYC:6569
F7D5  70 05     BVS $F7DC                       A:0 X:33 Y:B9 P:27 SP:F9 CYC:6571
F7D7  90 03     BCC $F7DC                       A:0 X:33 Y:B9 P:27 SP:F9 CYC:6573
F7D9  30 01     BMI $F7DC                       A:0 X:33 Y:B9 P:27 SP:F9 CYC:6575
F7DB  60        RTS                             A:0 X:33 Y:B9 P:27 SP:F9 CYC:6577
D614  C8        INY                             A:0 X:33 Y:B9 P:27 SP:FB CYC:6583
D615  A9 AA     LDA #$AA                        A:0 X:33 Y:BA P:A5 SP:FB CYC:6585
D617  8D 78 06  STA $0678 = 00                  A:AA X:33 Y:BA P:A5 SP:FB CYC:6587
D61A  20 DF F7  JSR $F7DF                       A:AA X:33 Y:BA P:A5 SP:FB CYC:6591
F7DF  18        CLC                             A:AA X:33 Y:BA P:A5 SP:F9 CYC:6597
F7E0  24 01     BIT $01 = FF                    A:AA X:33 Y:BA P:A4 SP:F9 CYC:6599
F7E2  A9 55     LDA #$55                        A:AA X:33 Y:BA P:E4 SP:F9 CYC:6602
F7E4  60        RTS                             A:55 X:33 Y:BA P:64 SP:F9 CYC:6604
D61D  2D 78 06  AND $0678 = AA                  A:55 X:33 Y:BA P:64 SP:FB CYC:6610
D620  20 E5 F7  JSR $F7E5                       A:0 X:33 Y:BA P:66 SP:FB CYC:6614
F7E5  D0 07     BNE $F7EE                       A:0 X:33 Y:BA P:66 SP:F9 CYC:6620
F7E7  50 05     BVC $F7EE                       A:0 X:33 Y:BA P:66 SP:F9 CYC:6622
F7E9  B0 03     BCS $F7EE                       A:0 X:33 Y:BA P:66 SP:F9 CYC:6624
F7EB  30 01     BMI $F7EE                       A:0 X:33 Y:BA P:66 SP:F9 CYC:6626
F7ED  60        RTS                             A:0 X:33 Y:BA P:66 SP:F9 CYC:6628
D623  C8        INY                             A:0 X:33 Y:BA P:66 SP:FB CYC:6634
D624  A9 EF     LDA #$EF                        A:0 X:33 Y:BB P:E4 SP:FB CYC:6636
D626  8D 78 06  STA $0678 = AA                  A:EF X:33 Y:BB P:E4 SP:FB CYC:6638
D629  20 F1 F7  JSR $F7F1                       A:EF X:33 Y:BB P:E4 SP:FB CYC:6642
F7F1  38        SEC                             A:EF X:33 Y:BB P:E4 SP:F9 CYC:6648
F7F2  B8        CLV                             A:EF X:33 Y:BB P:E5 SP:F9 CYC:6650
F7F3  A9 F8     LDA #$F8                        A:EF X:33 Y:BB P:A5 SP:F9 CYC:6652
F7F5  60        RTS                             A:F8 X:33 Y:BB P:A5 SP:F9 CYC:6654
D62C  2D 78 06  AND $0678 = EF                  A:F8 X:33 Y:BB P:A5 SP:FB CYC:6660
D62F  20 F6 F7  JSR $F7F6                       A:E8 X:33 Y:BB P:A5 SP:FB CYC:6664
F7F6  90 09     BCC $F801                       A:E8 X:33 Y:BB P:A5 SP:F9 CYC:6670
F7F8  10 07     BPL $F801                       A:E8 X:33 Y:BB P:A5 SP:F9 CYC:6672
F7FA  C9 E8     CMP #$E8                        A:E8 X:33 Y:BB P:A5 SP:F9 CYC:6674
F7FC  D0 03     BNE $F801                       A:E8 X:33 Y:BB P:27 SP:F9 CYC:6676
F7FE  70 01     BVS $F801                       A:E8 X:33 Y:BB P:27 SP:F9 CYC:6678
F800  60        RTS                             A:E8 X:33 Y:BB P:27 SP:F9 CYC:6680
D632  C8        INY                             A:E8 X:33 Y:BB P:27 SP:FB CYC:6686
D633  A9 AA     LDA #$AA                        A:E8 X:33 Y:BC P:A5 SP:FB CYC:6688
D635  8D 78 06  STA $0678 = EF                  A:AA X:33 Y:BC P:A5 SP:FB CYC:6690
D638  20 04 F8  JSR $F804                       A:AA X:33 Y:BC P:A5 SP:FB CYC:6694
F804  18        CLC                             A:AA X:33 Y:BC P:A5 SP:F9 CYC:6700
F805  24 01     BIT $01 = FF                    A:AA X:33 Y:BC P:A4 SP:F9 CYC:6702
F807  A9 5F     LDA #$5F                        A:AA X:33 Y:BC P:E4 SP:F9 CYC:6705
F809  60        RTS                             A:5F X:33 Y:BC P:64 SP:F9 CYC:6707
D63B  4D 78 06  EOR $0678 = AA                  A:5F X:33 Y:BC P:64 SP:FB CYC:6713
D63E  20 0A F8  JSR $F80A                       A:F5 X:33 Y:BC P:E4 SP:FB CYC:6717
F80A  B0 09     BCS $F815                       A:F5 X:33 Y:BC P:E4 SP:F9 CYC:6723
F80C  10 07     BPL $F815                       A:F5 X:33 Y:BC P:E4 SP:F9 CYC:6725
F80E  C9 F5     CMP #$F5                        A:F5 X:33 Y:BC P:E4 SP:F9 CYC:6727
F810  D0 03     BNE $F815                       A:F5 X:33 Y:BC P:67 SP:F9 CYC:6729
F812  50 01     BVC $F815                       A:F5 X:33 Y:BC P:67 SP:F9 CYC:6731
F814  60        RTS                             A:F5 X:33 Y:BC P:67 SP:F9 CYC:6733
D641  C8        INY                             A:F5 X:33 Y:BC P:67 SP:FB CYC:6739
D642  A9 70     LDA #$70                        A:F5 X:33 Y:BD P:E5 SP:FB CYC:6741
D644  8D 78 06  STA $0678 = AA                  A:70 X:33 Y:BD P:65 SP:FB CYC:6743
D647  20 18 F8  JSR $F818                       A:70 X:33 Y:BD P:65 SP:FB CYC:6747
F818  38        SEC                             A:70 X:33 Y:BD P:65 SP:F9 CYC:6753
F819  B8        CLV                             A:70 X:33 Y:BD P:65 SP:F9 CYC:6755
F81A  A9 70     LDA #$70                        A:70 X:33 Y:BD P:25 SP:F9 CYC:6757
F81C  60        RTS                             A:70 X:33 Y:BD P:25 SP:F9 CYC:6759
D64A  4D 78 06  EOR $0678 = 70                  A:70 X:33 Y:BD P:25 SP:FB CYC:6765
D64D  20 1D F8  JSR $F81D                       A:0 X:33 Y:BD P:27 SP:FB CYC:6769
F81D  D0 07     BNE $F826                       A:0 X:33 Y:BD P:27 SP:F9 CYC:6775
F81F  70 05     BVS $F826                       A:0 X:33 Y:BD P:27 SP:F9 CYC:6777
F821  90 03     BCC $F826                       A:0 X:33 Y:BD P:27 SP:F9 CYC:6779
F823  30 01     BMI $F826                       A:0 X:33 Y:BD P:27 SP:F9 CYC:6781
F825  60        RTS                             A:0 X:33 Y:BD P:27 SP:F9 CYC:6783
D650  C8        INY                             A:0 X:33 Y:BD P:27 SP:FB CYC:6789
D651  A9 69     LDA #$69                        A:0 X:33 Y:BE P:A5 SP:FB CYC:6791
D653  8D 78 06  STA $0678 = 70                  A:69 X:33 Y:BE P:25 SP:FB CYC:6793
D656  20 29 F8  JSR $F829                       A:69 X:33 Y:BE P:25 SP:FB CYC:6797
F829  18        CLC                             A:69 X:33 Y:BE P:25 SP:F9 CYC:6803
F82A  24 01     BIT $01 = FF                    A:69 X:33 Y:BE P:24 SP:F9 CYC:6805
F82C  A9 00     LDA #$00                        A:69 X:33 Y:BE P:E4 SP:F9 CYC:6808
F82E  60        RTS                             A:0 X:33 Y:BE P:66 SP:F9 CYC:6810
D659  6D 78 06  ADC $0678 = 69                  A:0 X:33 Y:BE P:66 SP:FB CYC:6816
D65C  20 2F F8  JSR $F82F                       A:69 X:33 Y:BE P:24 SP:FB CYC:6820
F82F  30 09     BMI $F83A                       A:69 X:33 Y:BE P:24 SP:F9 CYC:6826
F831  B0 07     BCS $F83A                       A:69 X:33 Y:BE P:24 SP:F9 CYC:6828
F833  C9 69     CMP #$69                        A:69 X:33 Y:BE P:24 SP:F9 CYC:6830
F835  D0 03     BNE $F83A                       A:69 X:33 Y:BE P:27 SP:F9 CYC:6832
F837  70 01     BVS $F83A                       A:69 X:33 Y:BE P:27 SP:F9 CYC:6834
F839  60        RTS                             A:69 X:33 Y:BE P:27 SP:F9 CYC:6836
D65F  C8        INY                             A:69 X:33 Y:BE P:27 SP:FB CYC:6842
D660  20 3D F8  JSR $F83D                       A:69 X:33 Y:BF P:A5 SP:FB CYC:6844
F83D  38        SEC                             A:69 X:33 Y:BF P:A5 SP:F9 CYC:6850
F83E  24 01     BIT $01 = FF                    A:69 X:33 Y:BF P:A5 SP:F9 CYC:6852
F840  A9 00     LDA #$00                        A:69 X:33 Y:BF P:E5 SP:F9 CYC:6855
F842  60        RTS                             A:0 X:33 Y:BF P:67 SP:F9 CYC:6857
D663  6D 78 06  ADC $0678 = 69                  A:0 X:33 Y:BF P:67 SP:FB CYC:6863
D666  20 43 F8  JSR $F843                       A:6A X:33 Y:BF P:24 SP:FB CYC:6867
F843  30 09     BMI $F84E                       A:6A X:33 Y:BF P:24 SP:F9 CYC:6873
F845  B0 07     BCS $F84E                       A:6A X:33 Y:BF P:24 SP:F9 CYC:6875
F847  C9 6A     CMP #$6A                        A:6A X:33 Y:BF P:24 SP:F9 CYC:6877
F849  D0 03     BNE $F84E                       A:6A X:33 Y:BF P:27 SP:F9 CYC:6879
F84B  70 01     BVS $F84E                       A:6A X:33 Y:BF P:27 SP:F9 CYC:6881
F84D  60        RTS                             A:6A X:33 Y:BF P:27 SP:F9 CYC:6883
D669  C8        INY                             A:6A X:33 Y:BF P:27 SP:FB CYC:6889
D66A  A9 7F     LDA #$7F                        A:6A X:33 Y:C0 P:A5 SP:FB CYC:6891
D66C  8D 78 06  STA $0678 = 69                  A:7F X:33 Y:C0 P:25 SP:FB CYC:6893
D66F  20 51 F8  JSR $F851                       A:7F X:33 Y:C0 P:25 SP:FB CYC:6897
F851  38        SEC                             A:7F X:33 Y:C0 P:25 SP:F9 CYC:6903
F852  B8        CLV                             A:7F X:33 Y:C0 P:25 SP:F9 CYC:6905
F853  A9 7F     LDA #$7F                        A:7F X:33 Y:C0 P:25 SP:F9 CYC:6907
F855  60        RTS                             A:7F X:33 Y:C0 P:25 SP:F9 CYC:6909
D672  6D 78 06  ADC $0678 = 7F                  A:7F X:33 Y:C0 P:25 SP:FB CYC:6915
D675  20 56 F8  JSR $F856                       A:FF X:33 Y:C0 P:E4 SP:FB CYC:6919
F856  10 09     BPL $F861                       A:FF X:33 Y:C0 P:E4 SP:F9 CYC:6925
F858  B0 07     BCS $F861                       A:FF X:33 Y:C0 P:E4 SP:F9 CYC:6927
F85A  C9 FF     CMP #$FF                        A:FF X:33 Y:C0 P:E4 SP:F9 CYC:6929
F85C  D0 03     BNE $F861                       A:FF X:33 Y:C0 P:67 SP:F9 CYC:6931
F85E  50 01     BVC $F861                       A:FF X:33 Y:C0 P:67 SP:F9 CYC:6933
F860  60        RTS                             A:FF X:33 Y:C0 P:67 SP:F9 CYC:6935
D678  C8        INY                             A:FF X:33 Y:C0 P:67 SP:FB CYC:6941
D679  A9 80     LDA #$80                        A:FF X:33 Y:C1 P:E5 SP:FB CYC:6943
D67B  8D 78 06  STA $0678 = 7F                  A:80 X:33 Y:C1 P:E5 SP:FB CYC:6945
D67E  20 64 F8  JSR $F864                       A:80 X:33 Y:C1 P:E5 SP:FB CYC:6949
F864  18        CLC                             A:80 X:33 Y:C1 P:E5 SP:F9 CYC:6955
F865  24 01     BIT $01 = FF                    A:80 X:33 Y:C1 P:E4 SP:F9 CYC:6957
F867  A9 7F     LDA #$7F                        A:80 X:33 Y:C1 P:E4 SP:F9 CYC:6960
F869  60        RTS                             A:7F X:33 Y:C1 P:64 SP:F9 CYC:6962
D681  6D 78 06  ADC $0678 = 80                  A:7F X:33 Y:C1 P:64 SP:FB CYC:6968
D684  20 6A F8  JSR $F86A                       A:FF X:33 Y:C1 P:A4 SP:FB CYC:6972
F86A  10 09     BPL $F875                       A:FF X:33 Y:C1 P:A4 SP:F9 CYC:6978
F86C  B0 07     BCS $F875                       A:FF X:33 Y:C1 P:A4 SP:F9 CYC:6980
F86E  C9 FF     CMP #$FF                        A:FF X:33 Y:C1 P:A4 SP:F9 CYC:6982
F870  D0 03     BNE $F875                       A:FF X:33 Y:C1 P:27 SP:F9 CYC:6984
F872  70 01     BVS $F875                       A:FF X:33 Y:C1 P:27 SP:F9 CYC:6986
F874  60        RTS                             A:FF X:33 Y:C1 P:27 SP:F9 CYC:6988
D687  C8        INY                             A:FF X:33 Y:C1 P:27 SP:FB CYC:6994
D688  20 78 F8  JSR $F878                       A:FF X:33 Y:C2 P:A5 SP:FB CYC:6996
F878  38        SEC                             A:FF X:33 Y:C2 P:A5 SP:F9 CYC:7002
F879  B8        CLV                             A:FF X:33 Y:C2 P:A5 SP:F9 CYC:7004
F87A  A9 7F     LDA #$7F                        A:FF X:33 Y:C2 P:A5 SP:F9 CYC:7006
F87C  60        RTS                             A:7F X:33 Y:C2 P:25 SP:F9 CYC:7008
D68B  6D 78 06  ADC $0678 = 80                  A:7F X:33 Y:C2 P:25 SP:FB CYC:7014
D68E  20 7D F8  JSR $F87D                       A:0 X:33 Y:C2 P:27 SP:FB CYC:7018
F87D  D0 07     BNE $F886                       A:0 X:33 Y:C2 P:27 SP:F9 CYC:7024
F87F  30 05     BMI $F886                       A:0 X:33 Y:C2 P:27 SP:F9 CYC:7026
F881  70 03     BVS $F886                       A:0 X:33 Y:C2 P:27 SP:F9 CYC:7028
F883  90 01     BCC $F886                       A:0 X:33 Y:C2 P:27 SP:F9 CYC:7030
F885  60        RTS                             A:0 X:33 Y:C2 P:27 SP:F9 CYC:7032
D691  C8        INY                             A:0 X:33 Y:C2 P:27 SP:FB CYC:7038
D692  A9 40     LDA #$40                        A:0 X:33 Y:C3 P:A5 SP:FB CYC:7040
D694  8D 78 06  STA $0678 = 80                  A:40 X:33 Y:C3 P:25 SP:FB CYC:7042
D697  20 89 F8  JSR $F889                       A:40 X:33 Y:C3 P:25 SP:FB CYC:7046
F889  24 01     BIT $01 = FF                    A:40 X:33 Y:C3 P:25 SP:F9 CYC:7052
F88B  A9 40     LDA #$40                        A:40 X:33 Y:C3 P:E5 SP:F9 CYC:7055
F88D  60        RTS                             A:40 X:33 Y:C3 P:65 SP:F9 CYC:7057
D69A  CD 78 06  CMP $0678 = 40                  A:40 X:33 Y:C3 P:65 SP:FB CYC:7063
D69D  20 8E F8  JSR $F88E                       A:40 X:33 Y:C3 P:67 SP:FB CYC:7067
F88E  30 07     BMI $F897                       A:40 X:33 Y:C3 P:67 SP:F9 CYC:7073
F890  90 05     BCC $F897                       A:40 X:33 Y:C3 P:67 SP:F9 CYC:7075
F892  D0 03     BNE $F897                       A:40 X:33 Y:C3 P:67 SP:F9 CYC:7077
F894  50 01     BVC $F897                       A:40 X:33 Y:C3 P:67 SP:F9 CYC:7079
F896  60        RTS                             A:40 X:33 Y:C3 P:67 SP:F9 CYC:7081
D6A0  C8        INY                             A:40 X:33 Y:C3 P:67 SP:FB CYC:7087
D6A1  48        PHA                             A:40 X:33 Y:C4 P:E5 SP:FB CYC:7089
D6A2  A9 3F     LDA #$3F                        A:40 X:33 Y:C4 P:E5 SP:FA CYC:7092
D6A4  8D 78 06  STA $0678 = 40                  A:3F X:33 Y:C4 P:65 SP:FA CYC:7094
D6A7  68        PLA                             A:3F X:33 Y:C4 P:65 SP:FA CYC:7098
D6A8  20 9A F8  JSR $F89A                       A:40 X:33 Y:C4 P:65 SP:FB CYC:7102
F89A  B8        CLV                             A:40 X:33 Y:C4 P:65 SP:F9 CYC:7108
F89B  60        RTS                             A:40 X:33 Y:C4 P:25 SP:F9 CYC:7110
D6AB  CD 78 06  CMP $0678 = 3F                  A:40 X:33 Y:C4 P:25 SP:FB CYC:7116
D6AE  20 9C F8  JSR $F89C                       A:40 X:33 Y:C4 P:25 SP:FB CYC:7120
F89C  F0 07     BEQ $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9 CYC:7126
F89E  30 05     BMI $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9 CYC:7128
F8A0  90 03     BCC $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9 CYC:7130
F8A2  70 01     BVS $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9 CYC:7132
F8A4  60        RTS                             A:40 X:33 Y:C4 P:25 SP:F9 CYC:7134
D6B1  C8        INY                             A:40 X:33 Y:C4 P:25 SP:FB CYC:7140
D6B2  48        PHA                             A:40 X:33 Y:C5 P:A5 SP:FB CYC:7142
D6B3  A9 41     LDA #$41                        A:40 X:33 Y:C5 P:A5 SP:FA CYC:7145
D6B5  8D 78 06  STA $0678 = 3F                  A:41 X:33 Y:C5 P:25 SP:FA CYC:7147
D6B8  68        PLA                             A:41 X:33 Y:C5 P:25 SP:FA CYC:7151
D6B9  CD 78 06  CMP $0678 = 41                  A:40 X:33 Y:C5 P:25 SP:FB CYC:7155
D6BC  20 A8 F8  JSR $F8A8                       A:40 X:33 Y:C5 P:A4 SP:FB CYC:7159
F8A8  F0 05     BEQ $F8AF                       A:40 X:33 Y:C5 P:A4 SP:F9 CYC:7165
F8AA  10 03     BPL $F8AF                       A:40 X:33 Y:C5 P:A4 SP:F9 CYC:7167
F8AC  10 01     BPL $F8AF                       A:40 X:33 Y:C5 P:A4 SP:F9 CYC:7169
F8AE  60        RTS                             A:40 X:33 Y:C5 P:A4 SP:F9 CYC:7171
D6BF  C8        INY                             A:40 X:33 Y:C5 P:A4 SP:FB CYC:7177
D6C0  48        PHA                             A:40 X:33 Y:C6 P:A4 SP:FB CYC:7179
D6C1  A9 00     LDA #$00                        A:40 X:33 Y:C6 P:A4 SP:FA CYC:7182
D6C3  8D 78 06  STA $0678 = 41                  A:0 X:33 Y:C6 P:26 SP:FA CYC:7184
D6C6  68        PLA                             A:0 X:33 Y:C6 P:26 SP:FA CYC:7188
D6C7  20 B2 F8  JSR $F8B2                       A:40 X:33 Y:C6 P:24 SP:FB CYC:7192
F8B2  A9 80     LDA #$80                        A:40 X:33 Y:C6 P:24 SP:F9 CYC:7198
F8B4  60        RTS                             A:80 X:33 Y:C6 P:A4 SP:F9 CYC:7200
D6CA  CD 78 06  CMP $0678 = 00                  A:80 X:33 Y:C6 P:A4 SP:FB CYC:7206
D6CD  20 B5 F8  JSR $F8B5                       A:80 X:33 Y:C6 P:A5 SP:FB CYC:7210
F8B5  F0 05     BEQ $F8BC                       A:80 X:33 Y:C6 P:A5 SP:F9 CYC:7216
F8B7  10 03     BPL $F8BC                       A:80 X:33 Y:C6 P:A5 SP:F9 CYC:7218
F8B9  90 01     BCC $F8BC                       A:80 X:33 Y:C6 P:A5 SP:F9 CYC:7220
F8BB  60        RTS                             A:80 X:33 Y:C6 P:A5 SP:F9 CYC:7222
D6D0  C8        INY                             A:80 X:33 Y:C6 P:A5 SP:FB CYC:7228
D6D1  48        PHA                             A:80 X:33 Y:C7 P:A5 SP:FB CYC:7230
D6D2  A9 80     LDA #$80                        A:80 X:33 Y:C7 P:A5 SP:FA CYC:7233
D6D4  8D 78 06  STA $0678 = 00                  A:80 X:33 Y:C7 P:A5 SP:FA CYC:7235
D6D7  68        PLA                             A:80 X:33 Y:C7 P:A5 SP:FA CYC:7239
D6D8  CD 78 06  CMP $0678 = 80                  A:80 X:33 Y:C7 P:A5 SP:FB CYC:7243
D6DB  20 BF F8  JSR $F8BF                       A:80 X:33 Y:C7 P:27 SP:FB CYC:7247
F8BF  D0 05     BNE $F8C6                       A:80 X:33 Y:C7 P:27 SP:F9 CYC:7253
F8C1  30 03     BMI $F8C6                       A:80 X:33 Y:C7 P:27 SP:F9 CYC:7255
F8C3  90 01     BCC $F8C6                       A:80 X:33 Y:C7 P:27 SP:F9 CYC:7257
F8C5  60        RTS                             A:80 X:33 Y:C7 P:27 SP:F9 CYC:7259
D6DE  C8        INY                             A:80 X:33 Y:C7 P:27 SP:FB CYC:7265
D6DF  48        PHA                             A:80 X:33 Y:C8 P:A5 SP:FB CYC:7267
D6E0  A9 81     LDA #$81                        A:80 X:33 Y:C8 P:A5 SP:FA CYC:7270
D6E2  8D 78 06  STA $0678 = 80                  A:81 X:33 Y:C8 P:A5 SP:FA CYC:7272
D6E5  68        PLA                             A:81 X:33 Y:C8 P:A5 SP:FA CYC:7276
D6E6  CD 78 06  CMP $0678 = 81                  A:80 X:33 Y:C8 P:A5 SP:FB CYC:7280
D6E9  20 C9 F8  JSR $F8C9                       A:80 X:33 Y:C8 P:A4 SP:FB CYC:7284
F8C9  B0 05     BCS $F8D0                       A:80 X:33 Y:C8 P:A4 SP:F9 CYC:7290
F8CB  F0 03     BEQ $F8D0                       A:80 X:33 Y:C8 P:A4 SP:F9 CYC:7292
F8CD  10 01     BPL $F8D0                       A:80 X:33 Y:C8 P:A4 SP:F9 CYC:7294
F8CF  60        RTS                             A:80 X:33 Y:C8 P:A4 SP:F9 CYC:7296
D6EC  C8        INY                             A:80 X:33 Y:C8 P:A4 SP:FB CYC:7302
D6ED  48        PHA                             A:80 X:33 Y:C9 P:A4 SP:FB CYC:7304
D6EE  A9 7F     LDA #$7F                        A:80 X:33 Y:C9 P:A4 SP:FA CYC:7307
D6F0  8D 78 06  STA $0678 = 81                  A:7F X:33 Y:C9 P:24 SP:FA CYC:7309
D6F3  68        PLA                             A:7F X:33 Y:C9 P:24 SP:FA CYC:7313
D6F4  CD 78 06  CMP $0678 = 7F                  A:80 X:33 Y:C9 P:A4 SP:FB CYC:7317
D6F7  20 D3 F8  JSR $F8D3                       A:80 X:33 Y:C9 P:25 SP:FB CYC:7321
F8D3  90 05     BCC $F8DA                       A:80 X:33 Y:C9 P:25 SP:F9 CYC:7327
F8D5  F0 03     BEQ $F8DA                       A:80 X:33 Y:C9 P:25 SP:F9 CYC:7329
F8D7  30 01     BMI $F8DA                       A:80 X:33 Y:C9 P:25 SP:F9 CYC:7331
F8D9  60        RTS                             A:80 X:33 Y:C9 P:25 SP:F9 CYC:7333
D6FA  C8        INY                             A:80 X:33 Y:C9 P:25 SP:FB CYC:7339
D6FB  A9 40     LDA #$40                        A:80 X:33 Y:CA P:A5 SP:FB CYC:7341
D6FD  8D 78 06  STA $0678 = 7F                  A:40 X:33 Y:CA P:25 SP:FB CYC:7343
D700  20 31 F9  JSR $F931                       A:40 X:33 Y:CA P:25 SP:FB CYC:7347
F931  24 01     BIT $01 = FF                    A:40 X:33 Y:CA P:25 SP:F9 CYC:7353
F933  A9 40     LDA #$40                        A:40 X:33 Y:CA P:E5 SP:F9 CYC:7356
F935  38        SEC                             A:40 X:33 Y:CA P:65 SP:F9 CYC:7358
F936  60        RTS                             A:40 X:33 Y:CA P:65 SP:F9 CYC:7360
D703  ED 78 06  SBC $0678 = 40                  A:40 X:33 Y:CA P:65 SP:FB CYC:7366
D706  20 37 F9  JSR $F937                       A:0 X:33 Y:CA P:27 SP:FB CYC:7370
F937  30 0B     BMI $F944                       A:0 X:33 Y:CA P:27 SP:F9 CYC:7376
F939  90 09     BCC $F944                       A:0 X:33 Y:CA P:27 SP:F9 CYC:7378
F93B  D0 07     BNE $F944                       A:0 X:33 Y:CA P:27 SP:F9 CYC:7380
F93D  70 05     BVS $F944                       A:0 X:33 Y:CA P:27 SP:F9 CYC:7382
F93F  C9 00     CMP #$00                        A:0 X:33 Y:CA P:27 SP:F9 CYC:7384
F941  D0 01     BNE $F944                       A:0 X:33 Y:CA P:27 SP:F9 CYC:7386
F943  60        RTS                             A:0 X:33 Y:CA P:27 SP:F9 CYC:7388
D709  C8        INY                             A:0 X:33 Y:CA P:27 SP:FB CYC:7394
D70A  A9 3F     LDA #$3F                        A:0 X:33 Y:CB P:A5 SP:FB CYC:7396
D70C  8D 78 06  STA $0678 = 40                  A:3F X:33 Y:CB P:25 SP:FB CYC:7398
D70F  20 47 F9  JSR $F947                       A:3F X:33 Y:CB P:25 SP:FB CYC:7402
F947  B8        CLV                             A:3F X:33 Y:CB P:25 SP:F9 CYC:7408
F948  38        SEC                             A:3F X:33 Y:CB P:25 SP:F9 CYC:7410
F949  A9 40     LDA #$40                        A:3F X:33 Y:CB P:25 SP:F9 CYC:7412
F94B  60        RTS                             A:40 X:33 Y:CB P:25 SP:F9 CYC:7414
D712  ED 78 06  SBC $0678 = 3F                  A:40 X:33 Y:CB P:25 SP:FB CYC:7420
D715  20 4C F9  JSR $F94C                       A:1 X:33 Y:CB P:25 SP:FB CYC:7424
F94C  F0 0B     BEQ $F959                       A:1 X:33 Y:CB P:25 SP:F9 CYC:7430
F94E  30 09     BMI $F959                       A:1 X:33 Y:CB P:25 SP:F9 CYC:7432
F950  90 07     BCC $F959                       A:1 X:33 Y:CB P:25 SP:F9 CYC:7434
F952  70 05     BVS $F959                       A:1 X:33 Y:CB P:25 SP:F9 CYC:7436
F954  C9 01     CMP #$01                        A:1 X:33 Y:CB P:25 SP:F9 CYC:7438
F956  D0 01     BNE $F959                       A:1 X:33 Y:CB P:27 SP:F9 CYC:7440
F958  60        RTS                             A:1 X:33 Y:CB P:27 SP:F9 CYC:7442
D718  C8        INY                             A:1 X:33 Y:CB P:27 SP:FB CYC:7448
D719  A9 41     LDA #$41                        A:1 X:33 Y:CC P:A5 SP:FB CYC:7450
D71B  8D 78 06  STA $0678 = 3F                  A:41 X:33 Y:CC P:25 SP:FB CYC:7452
D71E  20 5C F9  JSR $F95C                       A:41 X:33 Y:CC P:25 SP:FB CYC:7456
F95C  A9 40     LDA #$40                        A:41 X:33 Y:CC P:25 SP:F9 CYC:7462
F95E  38        SEC                             A:40 X:33 Y:CC P:25 SP:F9 CYC:7464
F95F  24 01     BIT $01 = FF                    A:40 X:33 Y:CC P:25 SP:F9 CYC:7466
F961  60        RTS                             A:40 X:33 Y:CC P:E5 SP:F9 CYC:7469
D721  ED 78 06  SBC $0678 = 41                  A:40 X:33 Y:CC P:E5 SP:FB CYC:7475
D724  20 62 F9  JSR $F962                       A:FF X:33 Y:CC P:A4 SP:FB CYC:7479
F962  B0 0B     BCS $F96F                       A:FF X:33 Y:CC P:A4 SP:F9 CYC:7485
F964  F0 09     BEQ $F96F                       A:FF X:33 Y:CC P:A4 SP:F9 CYC:7487
F966  10 07     BPL $F96F                       A:FF X:33 Y:CC P:A4 SP:F9 CYC:7489
F968  70 05     BVS $F96F                       A:FF X:33 Y:CC P:A4 SP:F9 CYC:7491
F96A  C9 FF     CMP #$FF                        A:FF X:33 Y:CC P:A4 SP:F9 CYC:7493
F96C  D0 01     BNE $F96F                       A:FF X:33 Y:CC P:27 SP:F9 CYC:7495
F96E  60        RTS                             A:FF X:33 Y:CC P:27 SP:F9 CYC:7497
D727  C8        INY                             A:FF X:33 Y:CC P:27 SP:FB CYC:7503
D728  A9 00     LDA #$00                        A:FF X:33 Y:CD P:A5 SP:FB CYC:7505
D72A  8D 78 06  STA $0678 = 41                  A:0 X:33 Y:CD P:27 SP:FB CYC:7507
D72D  20 72 F9  JSR $F972                       A:0 X:33 Y:CD P:27 SP:FB CYC:7511
F972  18        CLC                             A:0 X:33 Y:CD P:27 SP:F9 CYC:7517
F973  A9 80     LDA #$80                        A:0 X:33 Y:CD P:26 SP:F9 CYC:7519
F975  60        RTS                             A:80 X:33 Y:CD P:A4 SP:F9 CYC:7521
D730  ED 78 06  SBC $0678 = 00                  A:80 X:33 Y:CD P:A4 SP:FB CYC:7527
D733  20 76 F9  JSR $F976                       A:7F X:33 Y:CD P:65 SP:FB CYC:7531
F976  90 05     BCC $F97D                       A:7F X:33 Y:CD P:65 SP:F9 CYC:7537
F978  C9 7F     CMP #$7F                        A:7F X:33 Y:CD P:65 SP:F9 CYC:7539
F97A  D0 01     BNE $F97D                       A:7F X:33 Y:CD P:67 SP:F9 CYC:7541
F97C  60        RTS                             A:7F X:33 Y:CD P:67 SP:F9 CYC:7543
D736  C8        INY                             A:7F X:33 Y:CD P:67 SP:FB CYC:7549
D737  A9 7F     LDA #$7F                        A:7F X:33 Y:CE P:E5 SP:FB CYC:7551
D739  8D 78 06  STA $0678 = 00                  A:7F X:33 Y:CE P:65 SP:FB CYC:7553
D73C  20 80 F9  JSR $F980                       A:7F X:33 Y:CE P:65 SP:FB CYC:7557
F980  38        SEC                             A:7F X:33 Y:CE P:65 SP:F9 CYC:7563
F981  A9 81     LDA #$81                        A:7F X:33 Y:CE P:65 SP:F9 CYC:7565
F983  60        RTS                             A:81 X:33 Y:CE P:E5 SP:F9 CYC:7567
D73F  ED 78 06  SBC $0678 = 7F                  A:81 X:33 Y:CE P:E5 SP:FB CYC:7573
D742  20 84 F9  JSR $F984                       A:2 X:33 Y:CE P:65 SP:FB CYC:7577
F984  50 07     BVC $F98D                       A:2 X:33 Y:CE P:65 SP:F9 CYC:7583
F986  90 05     BCC $F98D                       A:2 X:33 Y:CE P:65 SP:F9 CYC:7585
F988  C9 02     CMP #$02                        A:2 X:33 Y:CE P:65 SP:F9 CYC:7587
F98A  D0 01     BNE $F98D                       A:2 X:33 Y:CE P:67 SP:F9 CYC:7589
F98C  60        RTS                             A:2 X:33 Y:CE P:67 SP:F9 CYC:7591
D745  C8        INY                             A:2 X:33 Y:CE P:67 SP:FB CYC:7597
D746  A9 40     LDA #$40                        A:2 X:33 Y:CF P:E5 SP:FB CYC:7599
D748  8D 78 06  STA $0678 = 7F                  A:40 X:33 Y:CF P:65 SP:FB CYC:7601
D74B  20 89 F8  JSR $F889                       A:40 X:33 Y:CF P:65 SP:FB CYC:7605
F889  24 01     BIT $01 = FF                    A:40 X:33 Y:CF P:65 SP:F9 CYC:7611
F88B  A9 40     LDA #$40                        A:40 X:33 Y:CF P:E5 SP:F9 CYC:7614
F88D  60        RTS                             A:40 X:33 Y:CF P:65 SP:F9 CYC:7616
D74E  AA        TAX                             A:40 X:33 Y:CF P:65 SP:FB CYC:7622
D74F  EC 78 06  CPX $0678 = 40                  A:40 X:40 Y:CF P:65 SP:FB CYC:7624
D752  20 8E F8  JSR $F88E                       A:40 X:40 Y:CF P:67 SP:FB CYC:7628
F88E  30 07     BMI $F897                       A:40 X:40 Y:CF P:67 SP:F9 CYC:7634
F890  90 05     BCC $F897                       A:40 X:40 Y:CF P:67 SP:F9 CYC:7636
F892  D0 03     BNE $F897                       A:40 X:40 Y:CF P:67 SP:F9 CYC:7638
F894  50 01     BVC $F897                       A:40 X:40 Y:CF P:67 SP:F9 CYC:7640
F896  60        RTS                             A:40 X:40 Y:CF P:67 SP:F9 CYC:7642
D755  C8        INY                             A:40 X:40 Y:CF P:67 SP:FB CYC:7648
D756  A9 3F     LDA #$3F                        A:40 X:40 Y:D0 P:E5 SP:FB CYC:7650
D758  8D 78 06  STA $0678 = 40                  A:3F X:40 Y:D0 P:65 SP:FB CYC:7652
D75B  20 9A F8  JSR $F89A                       A:3F X:40 Y:D0 P:65 SP:FB CYC:7656
F89A  B8        CLV                             A:3F X:40 Y:D0 P:65 SP:F9 CYC:7662
F89B  60        RTS                             A:3F X:40 Y:D0 P:25 SP:F9 CYC:7664
D75E  EC 78 06  CPX $0678 = 3F                  A:3F X:40 Y:D0 P:25 SP:FB CYC:7670
D761  20 9C F8  JSR $F89C                       A:3F X:40 Y:D0 P:25 SP:FB CYC:7674
F89C  F0 07     BEQ $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9 CYC:7680
F89E  30 05     BMI $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9 CYC:7682
F8A0  90 03     BCC $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9 CYC:7684
F8A2  70 01     BVS $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9 CYC:7686
F8A4  60        RTS                             A:3F X:40 Y:D0 P:25 SP:F9 CYC:7688
D764  C8        INY                             A:3F X:40 Y:D0 P:25 SP:FB CYC:7694
D765  A9 41     LDA #$41                        A:3F X:40 Y:D1 P:A5 SP:FB CYC:7696
D767  8D 78 06  STA $0678 = 3F                  A:41 X:40 Y:D1 P:25 SP:FB CYC:7698
D76A  EC 78 06  CPX $0678 = 41                  A:41 X:40 Y:D1 P:25 SP:FB CYC:7702
D76D  20 A8 F8  JSR $F8A8                       A:41 X:40 Y:D1 P:A4 SP:FB CYC:7706
F8A8  F0 05     BEQ $F8AF                       A:41 X:40 Y:D1 P:A4 SP:F9 CYC:7712
F8AA  10 03     BPL $F8AF                       A:41 X:40 Y:D1 P:A4 SP:F9 CYC:7714
F8AC  10 01     BPL $F8AF                       A:41 X:40 Y:D1 P:A4 SP:F9 CYC:7716
F8AE  60        RTS                             A:41 X:40 Y:D1 P:A4 SP:F9 CYC:7718
D770  C8        INY                             A:41 X:40 Y:D1 P:A4 SP:FB CYC:7724
D771  A9 00     LDA #$00                        A:41 X:40 Y:D2 P:A4 SP:FB CYC:7726
D773  8D 78 06  STA $0678 = 41                  A:0 X:40 Y:D2 P:26 SP:FB CYC:7728
D776  20 B2 F8  JSR $F8B2                       A:0 X:40 Y:D2 P:26 SP:FB CYC:7732
F8B2  A9 80     LDA #$80                        A:0 X:40 Y:D2 P:26 SP:F9 CYC:7738
F8B4  60        RTS                             A:80 X:40 Y:D2 P:A4 SP:F9 CYC:7740
D779  AA        TAX                             A:80 X:40 Y:D2 P:A4 SP:FB CYC:7746
D77A  EC 78 06  CPX $0678 = 00                  A:80 X:80 Y:D2 P:A4 SP:FB CYC:7748
D77D  20 B5 F8  JSR $F8B5                       A:80 X:80 Y:D2 P:A5 SP:FB CYC:7752
F8B5  F0 05     BEQ $F8BC                       A:80 X:80 Y:D2 P:A5 SP:F9 CYC:7758
F8B7  10 03     BPL $F8BC                       A:80 X:80 Y:D2 P:A5 SP:F9 CYC:7760
F8B9  90 01     BCC $F8BC                       A:80 X:80 Y:D2 P:A5 SP:F9 CYC:7762
F8BB  60        RTS                             A:80 X:80 Y:D2 P:A5 SP:F9 CYC:7764
D780  C8        INY                             A:80 X:80 Y:D2 P:A5 SP:FB CYC:7770
D781  A9 80     LDA #$80                        A:80 X:80 Y:D3 P:A5 SP:FB CYC:7772
D783  8D 78 06  STA $0678 = 00                  A:80 X:80 Y:D3 P:A5 SP:FB CYC:7774
D786  EC 78 06  CPX $0678 = 80                  A:80 X:80 Y:D3 P:A5 SP:FB CYC:7778
D789  20 BF F8  JSR $F8BF                       A:80 X:80 Y:D3 P:27 SP:FB CYC:7782
F8BF  D0 05     BNE $F8C6                       A:80 X:80 Y:D3 P:27 SP:F9 CYC:7788
F8C1  30 03     BMI $F8C6                       A:80 X:80 Y:D3 P:27 SP:F9 CYC:7790
F8C3  90 01     BCC $F8C6                       A:80 X:80 Y:D3 P:27 SP:F9 CYC:7792
F8C5  60        RTS                             A:80 X:80 Y:D3 P:27 SP:F9 CYC:7794
D78C  C8        INY                             A:80 X:80 Y:D3 P:27 SP:FB CYC:7800
D78D  A9 81     LDA #$81                        A:80 X:80 Y:D4 P:A5 SP:FB CYC:7802
D78F  8D 78 06  STA $0678 = 80                  A:81 X:80 Y:D4 P:A5 SP:FB CYC:7804
D792  EC 78 06  CPX $0678 = 81                  A:81 X:80 Y:D4 P:A5 SP:FB CYC:7808
D795  20 C9 F8  JSR $F8C9                       A:81 X:80 Y:D4 P:A4 SP:FB CYC:7812
F8C9  B0 05     BCS $F8D0                       A:81 X:80 Y:D4 P:A4 SP:F9 CYC:7818
F8CB  F0 03     BEQ $F8D0                       A:81 X:80 Y:D4 P:A4 SP:F9 CYC:7820
F8CD  10 01     BPL $F8D0                       A:81 X:80 Y:D4 P:A4 SP:F9 CYC:7822
F8CF  60        RTS                             A:81 X:80 Y:D4 P:A4 SP:F9 CYC:7824
D798  C8        INY                             A:81 X:80 Y:D4 P:A4 SP:FB CYC:7830
D799  A9 7F     LDA #$7F                        A:81 X:80 Y:D5 P:A4 SP:FB CYC:7832
D79B  8D 78 06  STA $0678 = 81                  A:7F X:80 Y:D5 P:24 SP:FB CYC:7834
D79E  EC 78 06  CPX $0678 = 7F                  A:7F X:80 Y:D5 P:24 SP:FB CYC:7838
D7A1  20 D3 F8  JSR $F8D3                       A:7F X:80 Y:D5 P:25 SP:FB CYC:7842
F8D3  90 05     BCC $F8DA                       A:7F X:80 Y:D5 P:25 SP:F9 CYC:7848
F8D5  F0 03     BEQ $F8DA                       A:7F X:80 Y:D5 P:25 SP:F9 CYC:7850
F8D7  30 01     BMI $F8DA                       A:7F X:80 Y:D5 P:25 SP:F9 CYC:7852
F8D9  60        RTS                             A:7F X:80 Y:D5 P:25 SP:F9 CYC:7854
D7A4  C8        INY                             A:7F X:80 Y:D5 P:25 SP:FB CYC:7860
D7A5  98        TYA                             A:7F X:80 Y:D6 P:A5 SP:FB CYC:7862
D7A6  AA        TAX                             A:D6 X:80 Y:D6 P:A5 SP:FB CYC:7864
D7A7  A9 40     LDA #$40                        A:D6 X:D6 Y:D6 P:A5 SP:FB CYC:7866
D7A9  8D 78 06  STA $0678 = 7F                  A:40 X:D6 Y:D6 P:25 SP:FB CYC:7868
D7AC  20 DD F8  JSR $F8DD                       A:40 X:D6 Y:D6 P:25 SP:FB CYC:7872
F8DD  24 01     BIT $01 = FF                    A:40 X:D6 Y:D6 P:25 SP:F9 CYC:7878
F8DF  A0 40     LDY #$40                        A:40 X:D6 Y:D6 P:E5 SP:F9 CYC:7881
F8E1  60        RTS                             A:40 X:D6 Y:40 P:65 SP:F9 CYC:7883
D7AF  CC 78 06  CPY $0678 = 40                  A:40 X:D6 Y:40 P:65 SP:FB CYC:7889
D7B2  20 E2 F8  JSR $F8E2                       A:40 X:D6 Y:40 P:67 SP:FB CYC:7893
F8E2  30 07     BMI $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9 CYC:7899
F8E4  90 05     BCC $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9 CYC:7901
F8E6  D0 03     BNE $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9 CYC:7903
F8E8  50 01     BVC $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9 CYC:7905
F8EA  60        RTS                             A:40 X:D6 Y:40 P:67 SP:F9 CYC:7907
D7B5  E8        INX                             A:40 X:D6 Y:40 P:67 SP:FB CYC:7913
D7B6  A9 3F     LDA #$3F                        A:40 X:D7 Y:40 P:E5 SP:FB CYC:7915
D7B8  8D 78 06  STA $0678 = 40                  A:3F X:D7 Y:40 P:65 SP:FB CYC:7917
D7BB  20 EE F8  JSR $F8EE                       A:3F X:D7 Y:40 P:65 SP:FB CYC:7921
F8EE  B8        CLV                             A:3F X:D7 Y:40 P:65 SP:F9 CYC:7927
F8EF  60        RTS                             A:3F X:D7 Y:40 P:25 SP:F9 CYC:7929
D7BE  CC 78 06  CPY $0678 = 3F                  A:3F X:D7 Y:40 P:25 SP:FB CYC:7935
D7C1  20 F0 F8  JSR $F8F0                       A:3F X:D7 Y:40 P:25 SP:FB CYC:7939
F8F0  F0 07     BEQ $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9 CYC:7945
F8F2  30 05     BMI $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9 CYC:7947
F8F4  90 03     BCC $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9 CYC:7949
F8F6  70 01     BVS $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9 CYC:7951
F8F8  60        RTS                             A:3F X:D7 Y:40 P:25 SP:F9 CYC:7953
D7C4  E8        INX                             A:3F X:D7 Y:40 P:25 SP:FB CYC:7959
D7C5  A9 41     LDA #$41                        A:3F X:D8 Y:40 P:A5 SP:FB CYC:7961
D7C7  8D 78 06  STA $0678 = 3F                  A:41 X:D8 Y:40 P:25 SP:FB CYC:7963
D7CA  CC 78 06  CPY $0678 = 41                  A:41 X:D8 Y:40 P:25 SP:FB CYC:7967
D7CD  20 FC F8  JSR $F8FC                       A:41 X:D8 Y:40 P:A4 SP:FB CYC:7971
F8FC  F0 05     BEQ $F903                       A:41 X:D8 Y:40 P:A4 SP:F9 CYC:7977
F8FE  10 03     BPL $F903                       A:41 X:D8 Y:40 P:A4 SP:F9 CYC:7979
F900  10 01     BPL $F903                       A:41 X:D8 Y:40 P:A4 SP:F9 CYC:7981
F902  60        RTS                             A:41 X:D8 Y:40 P:A4 SP:F9 CYC:7983
D7D0  E8        INX                             A:41 X:D8 Y:40 P:A4 SP:FB CYC:7989
D7D1  A9 00     LDA #$00                        A:41 X:D9 Y:40 P:A4 SP:FB CYC:7991
D7D3  8D 78 06  STA $0678 = 41                  A:0 X:D9 Y:40 P:26 SP:FB CYC:7993
D7D6  20 06 F9  JSR $F906                       A:0 X:D9 Y:40 P:26 SP:FB CYC:7997
F906  A0 80     LDY #$80                        A:0 X:D9 Y:40 P:26 SP:F9 CYC:8003
F908  60        RTS                             A:0 X:D9 Y:80 P:A4 SP:F9 CYC:8005
D7D9  CC 78 06  CPY $0678 = 00                  A:0 X:D9 Y:80 P:A4 SP:FB CYC:8011
D7DC  20 09 F9  JSR $F909                       A:0 X:D9 Y:80 P:A5 SP:FB CYC:8015
F909  F0 05     BEQ $F910                       A:0 X:D9 Y:80 P:A5 SP:F9 CYC:8021
F90B  10 03     BPL $F910                       A:0 X:D9 Y:80 P:A5 SP:F9 CYC:8023
F90D  90 01     BCC $F910                       A:0 X:D9 Y:80 P:A5 SP:F9 CYC:8025
F90F  60        RTS                             A:0 X:D9 Y:80 P:A5 SP:F9 CYC:8027
D7DF  E8        INX                             A:0 X:D9 Y:80 P:A5 SP:FB CYC:8033
D7E0  A9 80     LDA #$80                        A:0 X:DA Y:80 P:A5 SP:FB CYC:8035
D7E2  8D 78 06  STA $0678 = 00                  A:80 X:DA Y:80 P:A5 SP:FB CYC:8037
D7E5  CC 78 06  CPY $0678 = 80                  A:80 X:DA Y:80 P:A5 SP:FB CYC:8041
D7E8  20 13 F9  JSR $F913                       A:80 X:DA Y:80 P:27 SP:FB CYC:8045
F913  D0 05     BNE $F91A                       A:80 X:DA Y:80 P:27 SP:F9 CYC:8051
F915  30 03     BMI $F91A                       A:80 X:DA Y:80 P:27 SP:F9 CYC:8053
F917  90 01     BCC $F91A                       A:80 X:DA Y:80 P:27 SP:F9 CYC:8055
F919  60        RTS                             A:80 X:DA Y:80 P:27 SP:F9 CYC:8057
D7EB  E8        INX                             A:80 X:DA Y:80 P:27 SP:FB CYC:8063
D7EC  A9 81     LDA #$81                        A:80 X:DB Y:80 P:A5 SP:FB CYC:8065
D7EE  8D 78 06  STA $0678 = 80                  A:81 X:DB Y:80 P:A5 SP:FB CYC:8067
D7F1  CC 78 06  CPY $0678 = 81                  A:81 X:DB Y:80 P:A5 SP:FB CYC:8071
D7F4  20 1D F9  JSR $F91D                       A:81 X:DB Y:80 P:A4 SP:FB CYC:8075
F91D  B0 05     BCS $F924                       A:81 X:DB Y:80 P:A4 SP:F9 CYC:8081
F91F  F0 03     BEQ $F924                       A:81 X:DB Y:80 P:A4 SP:F9 CYC:8083
F921  10 01     BPL $F924                       A:81 X:DB Y:80 P:A4 SP:F9 CYC:8085
F923  60        RTS                             A:81 X:DB Y:80 P:A4 SP:F9 CYC:8087
D7F7  E8        INX                             A:81 X:DB Y:80 P:A4 SP:FB CYC:8093
D7F8  A9 7F     LDA #$7F                        A:81 X:DC Y:80 P:A4 SP:FB CYC:8095
D7FA  8D 78 06  STA $0678 = 81                  A:7F X:DC Y:80 P:24 SP:FB CYC:8097
D7FD  CC 78 06  CPY $0678 = 7F                  A:7F X:DC Y:80 P:24 SP:FB CYC:8101
D800  20 27 F9  JSR $F927                       A:7F X:DC Y:80 P:25 SP:FB CYC:8105
F927  90 05     BCC $F92E                       A:7F X:DC Y:80 P:25 SP:F9 CYC:8111
F929  F0 03     BEQ $F92E                       A:7F X:DC Y:80 P:25 SP:F9 CYC:8113
F92B  30 01     BMI $F92E                       A:7F X:DC Y:80 P:25 SP:F9 CYC:8115
F92D  60        RTS                             A:7F X:DC Y:80 P:25 SP:F9 CYC:8117
D803  E8        INX                             A:7F X:DC Y:80 P:25 SP:FB CYC:8123
D804  8A        TXA                             A:7F X:DD Y:80 P:A5 SP:FB CYC:8125
D805  A8        TAY                             A:DD X:DD Y:80 P:A5 SP:FB CYC:8127
D806  20 90 F9  JSR $F990                       A:DD X:DD Y:DD P:A5 SP:FB CYC:8129
F990  A2 55     LDX #$55                        A:DD X:DD Y:DD P:A5 SP:F9 CYC:8135
F992  A9 FF     LDA #$FF                        A:DD X:55 Y:DD P:25 SP:F9 CYC:8137
F994  85 01     STA $01 = FF                    A:FF X:55 Y:DD P:A5 SP:F9 CYC:8139
F996  EA        NOP                             A:FF X:55 Y:DD P:A5 SP:F9 CYC:8142
F997  24 01     BIT $01 = FF                    A:FF X:55 Y:DD P:A5 SP:F9 CYC:8144
F999  38        SEC                             A:FF X:55 Y:DD P:E5 SP:F9 CYC:8147
F99A  A9 01     LDA #$01                        A:FF X:55 Y:DD P:E5 SP:F9 CYC:8149
F99C  60        RTS                             A:1 X:55 Y:DD P:65 SP:F9 CYC:8151
D809  8D 78 06  STA $0678 = 7F                  A:1 X:55 Y:DD P:65 SP:FB CYC:8157
D80C  4E 78 06  LSR $0678 = 01                  A:1 X:55 Y:DD P:65 SP:FB CYC:8161
D80F  AD 78 06  LDA $0678 = 00                  A:1 X:55 Y:DD P:67 SP:FB CYC:8167
D812  20 9D F9  JSR $F99D                       A:0 X:55 Y:DD P:67 SP:FB CYC:8171
F99D  90 1B     BCC $F9BA                       A:0 X:55 Y:DD P:67 SP:F9 CYC:8177
F99F  D0 19     BNE $F9BA                       A:0 X:55 Y:DD P:67 SP:F9 CYC:8179
F9A1  30 17     BMI $F9BA                       A:0 X:55 Y:DD P:67 SP:F9 CYC:8181
F9A3  50 15     BVC $F9BA                       A:0 X:55 Y:DD P:67 SP:F9 CYC:8183
F9A5  C9 00     CMP #$00                        A:0 X:55 Y:DD P:67 SP:F9 CYC:8185
F9A7  D0 11     BNE $F9BA                       A:0 X:55 Y:DD P:67 SP:F9 CYC:8187
F9A9  B8        CLV                             A:0 X:55 Y:DD P:67 SP:F9 CYC:8189
F9AA  A9 AA     LDA #$AA                        A:0 X:55 Y:DD P:27 SP:F9 CYC:8191
F9AC  60        RTS                             A:AA X:55 Y:DD P:A5 SP:F9 CYC:8193
D815  C8        INY                             A:AA X:55 Y:DD P:A5 SP:FB CYC:8199
D816  8D 78 06  STA $0678 = 00                  A:AA X:55 Y:DE P:A5 SP:FB CYC:8201
D819  4E 78 06  LSR $0678 = AA                  A:AA X:55 Y:DE P:A5 SP:FB CYC:8205
D81C  AD 78 06  LDA $0678 = 55                  A:AA X:55 Y:DE P:24 SP:FB CYC:8211
D81F  20 AD F9  JSR $F9AD                       A:55 X:55 Y:DE P:24 SP:FB CYC:8215
F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:DE P:24 SP:F9 CYC:8221
F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:DE P:24 SP:F9 CYC:8223
F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:DE P:24 SP:F9 CYC:8225
F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:DE P:24 SP:F9 CYC:8227
F9B5  C9 55     CMP #$55                        A:55 X:55 Y:DE P:24 SP:F9 CYC:8229
F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:DE P:27 SP:F9 CYC:8231
F9B9  60        RTS                             A:55 X:55 Y:DE P:27 SP:F9 CYC:8233
D822  C8        INY                             A:55 X:55 Y:DE P:27 SP:FB CYC:8239
D823  20 BD F9  JSR $F9BD                       A:55 X:55 Y:DF P:A5 SP:FB CYC:8241
F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:DF P:A5 SP:F9 CYC:8247
F9BF  38        SEC                             A:55 X:55 Y:DF P:E5 SP:F9 CYC:8250
F9C0  A9 80     LDA #$80                        A:55 X:55 Y:DF P:E5 SP:F9 CYC:8252
F9C2  60        RTS                             A:80 X:55 Y:DF P:E5 SP:F9 CYC:8254
D826  8D 78 06  STA $0678 = 55                  A:80 X:55 Y:DF P:E5 SP:FB CYC:8260
D829  0E 78 06  ASL $0678 = 80                  A:80 X:55 Y:DF P:E5 SP:FB CYC:8264
D82C  AD 78 06  LDA $0678 = 00                  A:80 X:55 Y:DF P:67 SP:FB CYC:8270
D82F  20 C3 F9  JSR $F9C3                       A:0 X:55 Y:DF P:67 SP:FB CYC:8274
F9C3  90 1C     BCC $F9E1                       A:0 X:55 Y:DF P:67 SP:F9 CYC:8280
F9C5  D0 1A     BNE $F9E1                       A:0 X:55 Y:DF P:67 SP:F9 CYC:8282
F9C7  30 18     BMI $F9E1                       A:0 X:55 Y:DF P:67 SP:F9 CYC:8284
F9C9  50 16     BVC $F9E1                       A:0 X:55 Y:DF P:67 SP:F9 CYC:8286
F9CB  C9 00     CMP #$00                        A:0 X:55 Y:DF P:67 SP:F9 CYC:8288
F9CD  D0 12     BNE $F9E1                       A:0 X:55 Y:DF P:67 SP:F9 CYC:8290
F9CF  B8        CLV                             A:0 X:55 Y:DF P:67 SP:F9 CYC:8292
F9D0  A9 55     LDA #$55                        A:0 X:55 Y:DF P:27 SP:F9 CYC:8294
F9D2  38        SEC                             A:55 X:55 Y:DF P:25 SP:F9 CYC:8296
F9D3  60        RTS                             A:55 X:55 Y:DF P:25 SP:F9 CYC:8298
D832  C8        INY                             A:55 X:55 Y:DF P:25 SP:FB CYC:8304
D833  8D 78 06  STA $0678 = 00                  A:55 X:55 Y:E0 P:A5 SP:FB CYC:8306
D836  0E 78 06  ASL $0678 = 55                  A:55 X:55 Y:E0 P:A5 SP:FB CYC:8310
D839  AD 78 06  LDA $0678 = AA                  A:55 X:55 Y:E0 P:A4 SP:FB CYC:8316
D83C  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:E0 P:A4 SP:FB CYC:8320
F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9 CYC:8326
F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9 CYC:8328
F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9 CYC:8330
F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9 CYC:8332
F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:E0 P:A4 SP:F9 CYC:8334
F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:E0 P:27 SP:F9 CYC:8336
F9E0  60        RTS                             A:AA X:55 Y:E0 P:27 SP:F9 CYC:8338
D83F  C8        INY                             A:AA X:55 Y:E0 P:27 SP:FB CYC:8344
D840  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:E1 P:A5 SP:FB CYC:8346
F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:E1 P:A5 SP:F9 CYC:8352
F9E6  38        SEC                             A:AA X:55 Y:E1 P:E5 SP:F9 CYC:8355
F9E7  A9 01     LDA #$01                        A:AA X:55 Y:E1 P:E5 SP:F9 CYC:8357
F9E9  60        RTS                             A:1 X:55 Y:E1 P:65 SP:F9 CYC:8359
D843  8D 78 06  STA $0678 = AA                  A:1 X:55 Y:E1 P:65 SP:FB CYC:8365
D846  6E 78 06  ROR $0678 = 01                  A:1 X:55 Y:E1 P:65 SP:FB CYC:8369
D849  AD 78 06  LDA $0678 = 80                  A:1 X:55 Y:E1 P:E5 SP:FB CYC:8375
D84C  20 EA F9  JSR $F9EA                       A:80 X:55 Y:E1 P:E5 SP:FB CYC:8379
F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9 CYC:8385
F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9 CYC:8387
F9EE  10 18     BPL $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9 CYC:8389
F9F0  50 16     BVC $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9 CYC:8391
F9F2  C9 80     CMP #$80                        A:80 X:55 Y:E1 P:E5 SP:F9 CYC:8393
F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:E1 P:67 SP:F9 CYC:8395
F9F6  B8        CLV                             A:80 X:55 Y:E1 P:67 SP:F9 CYC:8397
F9F7  18        CLC                             A:80 X:55 Y:E1 P:27 SP:F9 CYC:8399
F9F8  A9 55     LDA #$55                        A:80 X:55 Y:E1 P:26 SP:F9 CYC:8401
F9FA  60        RTS                             A:55 X:55 Y:E1 P:24 SP:F9 CYC:8403
D84F  C8        INY                             A:55 X:55 Y:E1 P:24 SP:FB CYC:8409
D850  8D 78 06  STA $0678 = 80                  A:55 X:55 Y:E2 P:A4 SP:FB CYC:8411
D853  6E 78 06  ROR $0678 = 55                  A:55 X:55 Y:E2 P:A4 SP:FB CYC:8415
D856  AD 78 06  LDA $0678 = 2A                  A:55 X:55 Y:E2 P:25 SP:FB CYC:8421
D859  20 FB F9  JSR $F9FB                       A:2A X:55 Y:E2 P:25 SP:FB CYC:8425
F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:E2 P:25 SP:F9 CYC:8431
F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:E2 P:25 SP:F9 CYC:8433
F9FF  30 07     BMI $FA08                       A:2A X:55 Y:E2 P:25 SP:F9 CYC:8435
FA01  70 05     BVS $FA08                       A:2A X:55 Y:E2 P:25 SP:F9 CYC:8437
FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:E2 P:25 SP:F9 CYC:8439
FA05  D0 01     BNE $FA08                       A:2A X:55 Y:E2 P:27 SP:F9 CYC:8441
FA07  60        RTS                             A:2A X:55 Y:E2 P:27 SP:F9 CYC:8443
D85C  C8        INY                             A:2A X:55 Y:E2 P:27 SP:FB CYC:8449
D85D  20 0A FA  JSR $FA0A                       A:2A X:55 Y:E3 P:A5 SP:FB CYC:8451
FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:E3 P:A5 SP:F9 CYC:8457
FA0C  38        SEC                             A:2A X:55 Y:E3 P:E5 SP:F9 CYC:8460
FA0D  A9 80     LDA #$80                        A:2A X:55 Y:E3 P:E5 SP:F9 CYC:8462
FA0F  60        RTS                             A:80 X:55 Y:E3 P:E5 SP:F9 CYC:8464
D860  8D 78 06  STA $0678 = 2A                  A:80 X:55 Y:E3 P:E5 SP:FB CYC:8470
D863  2E 78 06  ROL $0678 = 80                  A:80 X:55 Y:E3 P:E5 SP:FB CYC:8474
D866  AD 78 06  LDA $0678 = 01                  A:80 X:55 Y:E3 P:65 SP:FB CYC:8480
D869  20 10 FA  JSR $FA10                       A:1 X:55 Y:E3 P:65 SP:FB CYC:8484
FA10  90 1C     BCC $FA2E                       A:1 X:55 Y:E3 P:65 SP:F9 CYC:8490
FA12  F0 1A     BEQ $FA2E                       A:1 X:55 Y:E3 P:65 SP:F9 CYC:8492
FA14  30 18     BMI $FA2E                       A:1 X:55 Y:E3 P:65 SP:F9 CYC:8494
FA16  50 16     BVC $FA2E                       A:1 X:55 Y:E3 P:65 SP:F9 CYC:8496
FA18  C9 01     CMP #$01                        A:1 X:55 Y:E3 P:65 SP:F9 CYC:8498
FA1A  D0 12     BNE $FA2E                       A:1 X:55 Y:E3 P:67 SP:F9 CYC:8500
FA1C  B8        CLV                             A:1 X:55 Y:E3 P:67 SP:F9 CYC:8502
FA1D  18        CLC                             A:1 X:55 Y:E3 P:27 SP:F9 CYC:8504
FA1E  A9 55     LDA #$55                        A:1 X:55 Y:E3 P:26 SP:F9 CYC:8506
FA20  60        RTS                             A:55 X:55 Y:E3 P:24 SP:F9 CYC:8508
D86C  C8        INY                             A:55 X:55 Y:E3 P:24 SP:FB CYC:8514
D86D  8D 78 06  STA $0678 = 01                  A:55 X:55 Y:E4 P:A4 SP:FB CYC:8516
D870  2E 78 06  ROL $0678 = 55                  A:55 X:55 Y:E4 P:A4 SP:FB CYC:8520
D873  AD 78 06  LDA $0678 = AA                  A:55 X:55 Y:E4 P:A4 SP:FB CYC:8526
D876  20 21 FA  JSR $FA21                       A:AA X:55 Y:E4 P:A4 SP:FB CYC:8530
FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9 CYC:8536
FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9 CYC:8538
FA25  10 07     BPL $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9 CYC:8540
FA27  70 05     BVS $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9 CYC:8542
FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:E4 P:A4 SP:F9 CYC:8544
FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:E4 P:27 SP:F9 CYC:8546
FA2D  60        RTS                             A:AA X:55 Y:E4 P:27 SP:F9 CYC:8548
D879  A9 FF     LDA #$FF                        A:AA X:55 Y:E4 P:27 SP:FB CYC:8554
D87B  8D 78 06  STA $0678 = AA                  A:FF X:55 Y:E4 P:A5 SP:FB CYC:8556
D87E  85 01     STA $01 = FF                    A:FF X:55 Y:E4 P:A5 SP:FB CYC:8560
D880  24 01     BIT $01 = FF                    A:FF X:55 Y:E4 P:A5 SP:FB CYC:8563
D882  38        SEC                             A:FF X:55 Y:E4 P:E5 SP:FB CYC:8566
D883  EE 78 06  INC $0678 = FF                  A:FF X:55 Y:E4 P:E5 SP:FB CYC:8568
D886  D0 0D     BNE $D895                       A:FF X:55 Y:E4 P:67 SP:FB CYC:8574
D888  30 0B     BMI $D895                       A:FF X:55 Y:E4 P:67 SP:FB CYC:8576
D88A  50 09     BVC $D895                       A:FF X:55 Y:E4 P:67 SP:FB CYC:8578
D88C  90 07     BCC $D895                       A:FF X:55 Y:E4 P:67 SP:FB CYC:8580
D88E  AD 78 06  LDA $0678 = 00                  A:FF X:55 Y:E4 P:67 SP:FB CYC:8582
D891  C9 00     CMP #$00                        A:0 X:55 Y:E4 P:67 SP:FB CYC:8586
D893  F0 04     BEQ $D899                       A:0 X:55 Y:E4 P:67 SP:FB CYC:8588
D899  A9 7F     LDA #$7F                        A:0 X:55 Y:E4 P:67 SP:FB CYC:8591
D89B  8D 78 06  STA $0678 = 00                  A:7F X:55 Y:E4 P:65 SP:FB CYC:8593
D89E  B8        CLV                             A:7F X:55 Y:E4 P:65 SP:FB CYC:8597
D89F  18        CLC                             A:7F X:55 Y:E4 P:25 SP:FB CYC:8599
D8A0  EE 78 06  INC $0678 = 7F                  A:7F X:55 Y:E4 P:24 SP:FB CYC:8601
D8A3  F0 0D     BEQ $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB CYC:8607
D8A5  10 0B     BPL $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB CYC:8609
D8A7  70 09     BVS $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB CYC:8611
D8A9  B0 07     BCS $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB CYC:8613
D8AB  AD 78 06  LDA $0678 = 80                  A:7F X:55 Y:E4 P:A4 SP:FB CYC:8615
D8AE  C9 80     CMP #$80                        A:80 X:55 Y:E4 P:A4 SP:FB CYC:8619
D8B0  F0 04     BEQ $D8B6                       A:80 X:55 Y:E4 P:27 SP:FB CYC:8621
D8B6  A9 00     LDA #$00                        A:80 X:55 Y:E4 P:27 SP:FB CYC:8624
D8B8  8D 78 06  STA $0678 = 80                  A:0 X:55 Y:E4 P:27 SP:FB CYC:8626
D8BB  24 01     BIT $01 = FF                    A:0 X:55 Y:E4 P:27 SP:FB CYC:8630
D8BD  38        SEC                             A:0 X:55 Y:E4 P:E7 SP:FB CYC:8633
D8BE  CE 78 06  DEC $0678 = 00                  A:0 X:55 Y:E4 P:E7 SP:FB CYC:8635
D8C1  F0 0D     BEQ $D8D0                       A:0 X:55 Y:E4 P:E5 SP:FB CYC:8641
D8C3  10 0B     BPL $D8D0                       A:0 X:55 Y:E4 P:E5 SP:FB CYC:8643
D8C5  50 09     BVC $D8D0                       A:0 X:55 Y:E4 P:E5 SP:FB CYC:8645
D8C7  90 07     BCC $D8D0                       A:0 X:55 Y:E4 P:E5 SP:FB CYC:8647
D8C9  AD 78 06  LDA $0678 = FF                  A:0 X:55 Y:E4 P:E5 SP:FB CYC:8649
D8CC  C9 FF     CMP #$FF                        A:FF X:55 Y:E4 P:E5 SP:FB CYC:8653
D8CE  F0 04     BEQ $D8D4                       A:FF X:55 Y:E4 P:67 SP:FB CYC:8655
D8D4  A9 80     LDA #$80                        A:FF X:55 Y:E4 P:67 SP:FB CYC:8658
D8D6  8D 78 06  STA $0678 = FF                  A:80 X:55 Y:E4 P:E5 SP:FB CYC:8660
D8D9  B8        CLV                             A:80 X:55 Y:E4 P:E5 SP:FB CYC:8664
D8DA  18        CLC                             A:80 X:55 Y:E4 P:A5 SP:FB CYC:8666
D8DB  CE 78 06  DEC $0678 = 80                  A:80 X:55 Y:E4 P:A4 SP:FB CYC:8668
D8DE  F0 0D     BEQ $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB CYC:8674
D8E0  30 0B     BMI $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB CYC:8676
D8E2  70 09     BVS $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB CYC:8678
D8E4  B0 07     BCS $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB CYC:8680
D8E6  AD 78 06  LDA $0678 = 7F                  A:80 X:55 Y:E4 P:24 SP:FB CYC:8682
D8E9  C9 7F     CMP #$7F                        A:7F X:55 Y:E4 P:24 SP:FB CYC:8686
D8EB  F0 04     BEQ $D8F1                       A:7F X:55 Y:E4 P:27 SP:FB CYC:8688
D8F1  A9 01     LDA #$01                        A:7F X:55 Y:E4 P:27 SP:FB CYC:8691
D8F3  8D 78 06  STA $0678 = 7F                  A:1 X:55 Y:E4 P:25 SP:FB CYC:8693
D8F6  CE 78 06  DEC $0678 = 01                  A:1 X:55 Y:E4 P:25 SP:FB CYC:8697
D8F9  F0 04     BEQ $D8FF                       A:1 X:55 Y:E4 P:27 SP:FB CYC:8703
D8FF  60        RTS                             A:1 X:55 Y:E4 P:27 SP:FB CYC:8706
C618  20 00 D9  JSR $D900                       A:1 X:55 Y:E4 P:27 SP:FD CYC:8712
D900  A9 A3     LDA #$A3                        A:1 X:55 Y:E4 P:27 SP:FB CYC:8718
D902  85 33     STA $33 = 00                    A:A3 X:55 Y:E4 P:A5 SP:FB CYC:8720
D904  A9 89     LDA #$89                        A:A3 X:55 Y:E4 P:A5 SP:FB CYC:8723
D906  8D 00 03  STA $0300 = 70                  A:89 X:55 Y:E4 P:A5 SP:FB CYC:8725
D909  A9 12     LDA #$12                        A:89 X:55 Y:E4 P:A5 SP:FB CYC:8729
D90B  8D 45 02  STA $0245 = 00                  A:12 X:55 Y:E4 P:25 SP:FB CYC:8731
D90E  A9 FF     LDA #$FF                        A:12 X:55 Y:E4 P:25 SP:FB CYC:8735
D910  85 01     STA $01 = FF                    A:FF X:55 Y:E4 P:A5 SP:FB CYC:8737
D912  A2 65     LDX #$65                        A:FF X:55 Y:E4 P:A5 SP:FB CYC:8740
D914  A9 00     LDA #$00                        A:FF X:65 Y:E4 P:25 SP:FB CYC:8742
D916  85 89     STA $89 = 00                    A:0 X:65 Y:E4 P:27 SP:FB CYC:8744
D918  A9 03     LDA #$03                        A:0 X:65 Y:E4 P:27 SP:FB CYC:8747
D91A  85 8A     STA $8A = 00                    A:3 X:65 Y:E4 P:25 SP:FB CYC:8749
D91C  A0 00     LDY #$00                        A:3 X:65 Y:E4 P:25 SP:FB CYC:8752
D91E  38        SEC                             A:3 X:65 Y:0 P:27 SP:FB CYC:8754
D91F  A9 00     LDA #$00                        A:3 X:65 Y:0 P:27 SP:FB CYC:8756
D921  B8        CLV                             A:0 X:65 Y:0 P:27 SP:FB CYC:8758
D922  B1 89     LDA ($89),Y = 0300 @ 0300 = 89  A:0 X:65 Y:0 P:27 SP:FB CYC:8760
D924  F0 0C     BEQ $D932                       A:89 X:65 Y:0 P:A5 SP:FB CYC:8765
D926  90 0A     BCC $D932                       A:89 X:65 Y:0 P:A5 SP:FB CYC:8767
D928  70 08     BVS $D932                       A:89 X:65 Y:0 P:A5 SP:FB CYC:8769
D92A  C9 89     CMP #$89                        A:89 X:65 Y:0 P:A5 SP:FB CYC:8771
D92C  D0 04     BNE $D932                       A:89 X:65 Y:0 P:27 SP:FB CYC:8773
D92E  E0 65     CPX #$65                        A:89 X:65 Y:0 P:27 SP:FB CYC:8775
D930  F0 04     BEQ $D936                       A:89 X:65 Y:0 P:27 SP:FB CYC:8777
D936  A9 FF     LDA #$FF                        A:89 X:65 Y:0 P:27 SP:FB CYC:8780
D938  85 97     STA $97 = 00                    A:FF X:65 Y:0 P:A5 SP:FB CYC:8782
D93A  85 98     STA $98 = 00                    A:FF X:65 Y:0 P:A5 SP:FB CYC:8785
D93C  24 98     BIT $98 = FF                    A:FF X:65 Y:0 P:A5 SP:FB CYC:8788
D93E  A0 34     LDY #$34                        A:FF X:65 Y:0 P:E5 SP:FB CYC:8791
D940  B1 97     LDA ($97),Y = FFFF @ 0033 = A3  A:FF X:65 Y:34 P:65 SP:FB CYC:8793
D942  C9 A3     CMP #$A3                        A:A3 X:65 Y:34 P:E5 SP:FB CYC:8799
D944  D0 02     BNE $D948                       A:A3 X:65 Y:34 P:67 SP:FB CYC:8801
D946  B0 04     BCS $D94C                       A:A3 X:65 Y:34 P:67 SP:FB CYC:8803
D94C  A5 00     LDA $00 = 00                    A:A3 X:65 Y:34 P:67 SP:FB CYC:8806
D94E  48        PHA                             A:0 X:65 Y:34 P:67 SP:FB CYC:8809
D94F  A9 46     LDA #$46                        A:0 X:65 Y:34 P:67 SP:FA CYC:8812
D951  85 FF     STA $FF = 00                    A:46 X:65 Y:34 P:65 SP:FA CYC:8814
D953  A9 01     LDA #$01                        A:46 X:65 Y:34 P:65 SP:FA CYC:8817
D955  85 00     STA $00 = 00                    A:1 X:65 Y:34 P:65 SP:FA CYC:8819
D957  A0 FF     LDY #$FF                        A:1 X:65 Y:34 P:65 SP:FA CYC:8822
D959  B1 FF     LDA ($FF),Y = 0146 @ 0245 = 12  A:1 X:65 Y:FF P:E5 SP:FA CYC:8824
D95B  C9 12     CMP #$12                        A:12 X:65 Y:FF P:65 SP:FA CYC:8830
D95D  F0 04     BEQ $D963                       A:12 X:65 Y:FF P:67 SP:FA CYC:8832
D963  68        PLA                             A:12 X:65 Y:FF P:67 SP:FA CYC:8835
D964  85 00     STA $00 = 01                    A:0 X:65 Y:FF P:67 SP:FB CYC:8839
D966  A2 ED     LDX #$ED                        A:0 X:65 Y:FF P:67 SP:FB CYC:8842
D968  A9 00     LDA #$00                        A:0 X:ED Y:FF P:E5 SP:FB CYC:8844
D96A  85 33     STA $33 = A3                    A:0 X:ED Y:FF P:67 SP:FB CYC:8846
D96C  A9 04     LDA #$04                        A:0 X:ED Y:FF P:67 SP:FB CYC:8849
D96E  85 34     STA $34 = 00                    A:4 X:ED Y:FF P:65 SP:FB CYC:8851
D970  A0 00     LDY #$00                        A:4 X:ED Y:FF P:65 SP:FB CYC:8854
D972  18        CLC                             A:4 X:ED Y:0 P:67 SP:FB CYC:8856
D973  A9 FF     LDA #$FF                        A:4 X:ED Y:0 P:66 SP:FB CYC:8858
D975  85 01     STA $01 = FF                    A:FF X:ED Y:0 P:E4 SP:FB CYC:8860
D977  24 01     BIT $01 = FF                    A:FF X:ED Y:0 P:E4 SP:FB CYC:8863
D979  A9 AA     LDA #$AA                        A:FF X:ED Y:0 P:E4 SP:FB CYC:8866
D97B  8D 00 04  STA $0400 = AD                  A:AA X:ED Y:0 P:E4 SP:FB CYC:8868
D97E  A9 55     LDA #$55                        A:AA X:ED Y:0 P:E4 SP:FB CYC:8872
D980  11 33     ORA ($33),Y = 0400 @ 0400 = AA  A:55 X:ED Y:0 P:64 SP:FB CYC:8874
D982  B0 08     BCS $D98C                       A:FF X:ED Y:0 P:E4 SP:FB CYC:8879
D984  10 06     BPL $D98C                       A:FF X:ED Y:0 P:E4 SP:FB CYC:8881
D986  C9 FF     CMP #$FF                        A:FF X:ED Y:0 P:E4 SP:FB CYC:8883
D988  D0 02     BNE $D98C                       A:FF X:ED Y:0 P:67 SP:FB CYC:8885
D98A  70 02     BVS $D98E                       A:FF X:ED Y:0 P:67 SP:FB CYC:8887
D98E  E8        INX                             A:FF X:ED Y:0 P:67 SP:FB CYC:8890
D98F  38        SEC                             A:FF X:EE Y:0 P:E5 SP:FB CYC:8892
D990  B8        CLV                             A:FF X:EE Y:0 P:E5 SP:FB CYC:8894
D991  A9 00     LDA #$00                        A:FF X:EE Y:0 P:A5 SP:FB CYC:8896
D993  11 33     ORA ($33),Y = 0400 @ 0400 = AA  A:0 X:EE Y:0 P:27 SP:FB CYC:8898
D995  F0 06     BEQ $D99D                       A:AA X:EE Y:0 P:A5 SP:FB CYC:8903
D997  70 04     BVS $D99D                       A:AA X:EE Y:0 P:A5 SP:FB CYC:8905
D999  90 02     BCC $D99D                       A:AA X:EE Y:0 P:A5 SP:FB CYC:8907
D99B  30 02     BMI $D99F                       A:AA X:EE Y:0 P:A5 SP:FB CYC:8909
D99F  E8        INX                             A:AA X:EE Y:0 P:A5 SP:FB CYC:8912
D9A0  18        CLC                             A:AA X:EF Y:0 P:A5 SP:FB CYC:8914
D9A1  24 01     BIT $01 = FF                    A:AA X:EF Y:0 P:A4 SP:FB CYC:8916
D9A3  A9 55     LDA #$55                        A:AA X:EF Y:0 P:E4 SP:FB CYC:8919
D9A5  31 33     AND ($33),Y = 0400 @ 0400 = AA  A:55 X:EF Y:0 P:64 SP:FB CYC:8921
D9A7  D0 06     BNE $D9AF                       A:0 X:EF Y:0 P:66 SP:FB CYC:8926
D9A9  50 04     BVC $D9AF                       A:0 X:EF Y:0 P:66 SP:FB CYC:8928
D9AB  B0 02     BCS $D9AF                       A:0 X:EF Y:0 P:66 SP:FB CYC:8930
D9AD  10 02     BPL $D9B1                       A:0 X:EF Y:0 P:66 SP:FB CYC:8932
D9B1  E8        INX                             A:0 X:EF Y:0 P:66 SP:FB CYC:8935
D9B2  38        SEC                             A:0 X:F0 Y:0 P:E4 SP:FB CYC:8937
D9B3  B8        CLV                             A:0 X:F0 Y:0 P:E5 SP:FB CYC:8939
D9B4  A9 EF     LDA #$EF                        A:0 X:F0 Y:0 P:A5 SP:FB CYC:8941
D9B6  8D 00 04  STA $0400 = AA                  A:EF X:F0 Y:0 P:A5 SP:FB CYC:8943
D9B9  A9 F8     LDA #$F8                        A:EF X:F0 Y:0 P:A5 SP:FB CYC:8947
D9BB  31 33     AND ($33),Y = 0400 @ 0400 = EF  A:F8 X:F0 Y:0 P:A5 SP:FB CYC:8949
D9BD  90 08     BCC $D9C7                       A:E8 X:F0 Y:0 P:A5 SP:FB CYC:8954
D9BF  10 06     BPL $D9C7                       A:E8 X:F0 Y:0 P:A5 SP:FB CYC:8956
D9C1  C9 E8     CMP #$E8                        A:E8 X:F0 Y:0 P:A5 SP:FB CYC:8958
D9C3  D0 02     BNE $D9C7                       A:E8 X:F0 Y:0 P:27 SP:FB CYC:8960
D9C5  50 02     BVC $D9C9                       A:E8 X:F0 Y:0 P:27 SP:FB CYC:8962
D9C9  E8        INX                             A:E8 X:F0 Y:0 P:27 SP:FB CYC:8965
D9CA  18        CLC                             A:E8 X:F1 Y:0 P:A5 SP:FB CYC:8967
D9CB  24 01     BIT $01 = FF                    A:E8 X:F1 Y:0 P:A4 SP:FB CYC:8969
D9CD  A9 AA     LDA #$AA                        A:E8 X:F1 Y:0 P:E4 SP:FB CYC:8972
D9CF  8D 00 04  STA $0400 = EF                  A:AA X:F1 Y:0 P:E4 SP:FB CYC:8974
D9D2  A9 5F     LDA #$5F                        A:AA X:F1 Y:0 P:E4 SP:FB CYC:8978
D9D4  51 33     EOR ($33),Y = 0400 @ 0400 = AA  A:5F X:F1 Y:0 P:64 SP:FB CYC:8980
D9D6  B0 08     BCS $D9E0                       A:F5 X:F1 Y:0 P:E4 SP:FB CYC:8985
D9D8  10 06     BPL $D9E0                       A:F5 X:F1 Y:0 P:E4 SP:FB CYC:8987
D9DA  C9 F5     CMP #$F5                        A:F5 X:F1 Y:0 P:E4 SP:FB CYC:8989
D9DC  D0 02     BNE $D9E0                       A:F5 X:F1 Y:0 P:67 SP:FB CYC:8991
D9DE  70 02     BVS $D9E2                       A:F5 X:F1 Y:0 P:67 SP:FB CYC:8993
D9E2  E8        INX                             A:F5 X:F1 Y:0 P:67 SP:FB CYC:8996
D9E3  38        SEC                             A:F5 X:F2 Y:0 P:E5 SP:FB CYC:8998
D9E4  B8        CLV                             A:F5 X:F2 Y:0 P:E5 SP:FB CYC:9000
D9E5  A9 70     LDA #$70                        A:F5 X:F2 Y:0 P:A5 SP:FB CYC:9002
D9E7  8D 00 04  STA $0400 = AA                  A:70 X:F2 Y:0 P:25 SP:FB CYC:9004
D9EA  51 33     EOR ($33),Y = 0400 @ 0400 = 70  A:70 X:F2 Y:0 P:25 SP:FB CYC:9008
D9EC  D0 06     BNE $D9F4                       A:0 X:F2 Y:0 P:27 SP:FB CYC:9013
D9EE  70 04     BVS $D9F4                       A:0 X:F2 Y:0 P:27 SP:FB CYC:9015
D9F0  90 02     BCC $D9F4                       A:0 X:F2 Y:0 P:27 SP:FB CYC:9017
D9F2  10 02     BPL $D9F6                       A:0 X:F2 Y:0 P:27 SP:FB CYC:9019
D9F6  E8        INX                             A:0 X:F2 Y:0 P:27 SP:FB CYC:9022
D9F7  18        CLC                             A:0 X:F3 Y:0 P:A5 SP:FB CYC:9024
D9F8  24 01     BIT $01 = FF                    A:0 X:F3 Y:0 P:A4 SP:FB CYC:9026
D9FA  A9 69     LDA #$69                        A:0 X:F3 Y:0 P:E6 SP:FB CYC:9029
D9FC  8D 00 04  STA $0400 = 70                  A:69 X:F3 Y:0 P:64 SP:FB CYC:9031
D9FF  A9 00     LDA #$00                        A:69 X:F3 Y:0 P:64 SP:FB CYC:9035
DA01  71 33     ADC ($33),Y = 0400 @ 0400 = 69  A:0 X:F3 Y:0 P:66 SP:FB CYC:9037
DA03  30 08     BMI $DA0D                       A:69 X:F3 Y:0 P:24 SP:FB CYC:9042
DA05  B0 06     BCS $DA0D                       A:69 X:F3 Y:0 P:24 SP:FB CYC:9044
DA07  C9 69     CMP #$69                        A:69 X:F3 Y:0 P:24 SP:FB CYC:9046
DA09  D0 02     BNE $DA0D                       A:69 X:F3 Y:0 P:27 SP:FB CYC:9048
DA0B  50 02     BVC $DA0F                       A:69 X:F3 Y:0 P:27 SP:FB CYC:9050
DA0F  E8        INX                             A:69 X:F3 Y:0 P:27 SP:FB CYC:9053
DA10  38        SEC                             A:69 X:F4 Y:0 P:A5 SP:FB CYC:9055
DA11  24 01     BIT $01 = FF                    A:69 X:F4 Y:0 P:A5 SP:FB CYC:9057
DA13  A9 00     LDA #$00                        A:69 X:F4 Y:0 P:E5 SP:FB CYC:9060
DA15  71 33     ADC ($33),Y = 0400 @ 0400 = 69  A:0 X:F4 Y:0 P:67 SP:FB CYC:9062
DA17  30 08     BMI $DA21                       A:6A X:F4 Y:0 P:24 SP:FB CYC:9067
DA19  B0 06     BCS $DA21                       A:6A X:F4 Y:0 P:24 SP:FB CYC:9069
DA1B  C9 6A     CMP #$6A                        A:6A X:F4 Y:0 P:24 SP:FB CYC:9071
DA1D  D0 02     BNE $DA21                       A:6A X:F4 Y:0 P:27 SP:FB CYC:9073
DA1F  50 02     BVC $DA23                       A:6A X:F4 Y:0 P:27 SP:FB CYC:9075
DA23  E8        INX                             A:6A X:F4 Y:0 P:27 SP:FB CYC:9078
DA24  38        SEC                             A:6A X:F5 Y:0 P:A5 SP:FB CYC:9080
DA25  B8        CLV                             A:6A X:F5 Y:0 P:A5 SP:FB CYC:9082
DA26  A9 7F     LDA #$7F                        A:6A X:F5 Y:0 P:A5 SP:FB CYC:9084
DA28  8D 00 04  STA $0400 = 69                  A:7F X:F5 Y:0 P:25 SP:FB CYC:9086
DA2B  71 33     ADC ($33),Y = 0400 @ 0400 = 7F  A:7F X:F5 Y:0 P:25 SP:FB CYC:9090
DA2D  10 08     BPL $DA37                       A:FF X:F5 Y:0 P:E4 SP:FB CYC:9095
DA2F  B0 06     BCS $DA37                       A:FF X:F5 Y:0 P:E4 SP:FB CYC:9097
DA31  C9 FF     CMP #$FF                        A:FF X:F5 Y:0 P:E4 SP:FB CYC:9099
DA33  D0 02     BNE $DA37                       A:FF X:F5 Y:0 P:67 SP:FB CYC:9101
DA35  70 02     BVS $DA39                       A:FF X:F5 Y:0 P:67 SP:FB CYC:9103
DA39  E8        INX                             A:FF X:F5 Y:0 P:67 SP:FB CYC:9106
DA3A  18        CLC                             A:FF X:F6 Y:0 P:E5 SP:FB CYC:9108
DA3B  24 01     BIT $01 = FF                    A:FF X:F6 Y:0 P:E4 SP:FB CYC:9110
DA3D  A9 80     LDA #$80                        A:FF X:F6 Y:0 P:E4 SP:FB CYC:9113
DA3F  8D 00 04  STA $0400 = 7F                  A:80 X:F6 Y:0 P:E4 SP:FB CYC:9115
DA42  A9 7F     LDA #$7F                        A:80 X:F6 Y:0 P:E4 SP:FB CYC:9119
DA44  71 33     ADC ($33),Y = 0400 @ 0400 = 80  A:7F X:F6 Y:0 P:64 SP:FB CYC:9121
DA46  10 08     BPL $DA50                       A:FF X:F6 Y:0 P:A4 SP:FB CYC:9126
DA48  B0 06     BCS $DA50                       A:FF X:F6 Y:0 P:A4 SP:FB CYC:9128
DA4A  C9 FF     CMP #$FF                        A:FF X:F6 Y:0 P:A4 SP:FB CYC:9130
DA4C  D0 02     BNE $DA50                       A:FF X:F6 Y:0 P:27 SP:FB CYC:9132
DA4E  50 02     BVC $DA52                       A:FF X:F6 Y:0 P:27 SP:FB CYC:9134
DA52  E8        INX                             A:FF X:F6 Y:0 P:27 SP:FB CYC:9137
DA53  38        SEC                             A:FF X:F7 Y:0 P:A5 SP:FB CYC:9139
DA54  B8        CLV                             A:FF X:F7 Y:0 P:A5 SP:FB CYC:9141
DA55  A9 80     LDA #$80                        A:FF X:F7 Y:0 P:A5 SP:FB CYC:9143
DA57  8D 00 04  STA $0400 = 80                  A:80 X:F7 Y:0 P:A5 SP:FB CYC:9145
DA5A  A9 7F     LDA #$7F                        A:80 X:F7 Y:0 P:A5 SP:FB CYC:9149
DA5C  71 33     ADC ($33),Y = 0400 @ 0400 = 80  A:7F X:F7 Y:0 P:25 SP:FB CYC:9151
DA5E  D0 06     BNE $DA66                       A:0 X:F7 Y:0 P:27 SP:FB CYC:9156
DA60  30 04     BMI $DA66                       A:0 X:F7 Y:0 P:27 SP:FB CYC:9158
DA62  70 02     BVS $DA66                       A:0 X:F7 Y:0 P:27 SP:FB CYC:9160
DA64  B0 02     BCS $DA68                       A:0 X:F7 Y:0 P:27 SP:FB CYC:9162
DA68  E8        INX                             A:0 X:F7 Y:0 P:27 SP:FB CYC:9165
DA69  24 01     BIT $01 = FF                    A:0 X:F8 Y:0 P:A5 SP:FB CYC:9167
DA6B  A9 40     LDA #$40                        A:0 X:F8 Y:0 P:E7 SP:FB CYC:9170
DA6D  8D 00 04  STA $0400 = 80                  A:40 X:F8 Y:0 P:65 SP:FB CYC:9172
DA70  D1 33     CMP ($33),Y = 0400 @ 0400 = 40  A:40 X:F8 Y:0 P:65 SP:FB CYC:9176
DA72  30 06     BMI $DA7A                       A:40 X:F8 Y:0 P:67 SP:FB CYC:9181
DA74  90 04     BCC $DA7A                       A:40 X:F8 Y:0 P:67 SP:FB CYC:9183
DA76  D0 02     BNE $DA7A                       A:40 X:F8 Y:0 P:67 SP:FB CYC:9185
DA78  70 02     BVS $DA7C                       A:40 X:F8 Y:0 P:67 SP:FB CYC:9187
DA7C  E8        INX                             A:40 X:F8 Y:0 P:67 SP:FB CYC:9190
DA7D  B8        CLV                             A:40 X:F9 Y:0 P:E5 SP:FB CYC:9192
DA7E  CE 00 04  DEC $0400 = 40                  A:40 X:F9 Y:0 P:A5 SP:FB CYC:9194
DA81  D1 33     CMP ($33),Y = 0400 @ 0400 = 3F  A:40 X:F9 Y:0 P:25 SP:FB CYC:9200
DA83  F0 06     BEQ $DA8B                       A:40 X:F9 Y:0 P:25 SP:FB CYC:9205
DA85  30 04     BMI $DA8B                       A:40 X:F9 Y:0 P:25 SP:FB CYC:9207
DA87  90 02     BCC $DA8B                       A:40 X:F9 Y:0 P:25 SP:FB CYC:9209
DA89  50 02     BVC $DA8D                       A:40 X:F9 Y:0 P:25 SP:FB CYC:9211
DA8D  E8        INX                             A:40 X:F9 Y:0 P:25 SP:FB CYC:9214
DA8E  EE 00 04  INC $0400 = 3F                  A:40 X:FA Y:0 P:A5 SP:FB CYC:9216
DA91  EE 00 04  INC $0400 = 40                  A:40 X:FA Y:0 P:25 SP:FB CYC:9222
DA94  D1 33     CMP ($33),Y = 0400 @ 0400 = 41  A:40 X:FA Y:0 P:25 SP:FB CYC:9228
DA96  F0 02     BEQ $DA9A                       A:40 X:FA Y:0 P:A4 SP:FB CYC:9233
DA98  30 02     BMI $DA9C                       A:40 X:FA Y:0 P:A4 SP:FB CYC:9235
DA9C  E8        INX                             A:40 X:FA Y:0 P:A4 SP:FB CYC:9238
DA9D  A9 00     LDA #$00                        A:40 X:FB Y:0 P:A4 SP:FB CYC:9240
DA9F  8D 00 04  STA $0400 = 41                  A:0 X:FB Y:0 P:26 SP:FB CYC:9242
DAA2  A9 80     LDA #$80                        A:0 X:FB Y:0 P:26 SP:FB CYC:9246
DAA4  D1 33     CMP ($33),Y = 0400 @ 0400 = 00  A:80 X:FB Y:0 P:A4 SP:FB CYC:9248
DAA6  F0 04     BEQ $DAAC                       A:80 X:FB Y:0 P:A5 SP:FB CYC:9253
DAA8  10 02     BPL $DAAC                       A:80 X:FB Y:0 P:A5 SP:FB CYC:9255
DAAA  B0 02     BCS $DAAE                       A:80 X:FB Y:0 P:A5 SP:FB CYC:9257
DAAE  E8        INX                             A:80 X:FB Y:0 P:A5 SP:FB CYC:9260
DAAF  A0 80     LDY #$80                        A:80 X:FC Y:0 P:A5 SP:FB CYC:9262
DAB1  8C 00 04  STY $0400 = 00                  A:80 X:FC Y:80 P:A5 SP:FB CYC:9264
DAB4  A0 00     LDY #$00                        A:80 X:FC Y:80 P:A5 SP:FB CYC:9268
DAB6  D1 33     CMP ($33),Y = 0400 @ 0400 = 80  A:80 X:FC Y:0 P:27 SP:FB CYC:9270
DAB8  D0 04     BNE $DABE                       A:80 X:FC Y:0 P:27 SP:FB CYC:9275
DABA  30 02     BMI $DABE                       A:80 X:FC Y:0 P:27 SP:FB CYC:9277
DABC  B0 02     BCS $DAC0                       A:80 X:FC Y:0 P:27 SP:FB CYC:9279
DAC0  E8        INX                             A:80 X:FC Y:0 P:27 SP:FB CYC:9282
DAC1  EE 00 04  INC $0400 = 80                  A:80 X:FD Y:0 P:A5 SP:FB CYC:9284
DAC4  D1 33     CMP ($33),Y = 0400 @ 0400 = 81  A:80 X:FD Y:0 P:A5 SP:FB CYC:9290
DAC6  B0 04     BCS $DACC                       A:80 X:FD Y:0 P:A4 SP:FB CYC:9295
DAC8  F0 02     BEQ $DACC                       A:80 X:FD Y:0 P:A4 SP:FB CYC:9297
DACA  30 02     BMI $DACE                       A:80 X:FD Y:0 P:A4 SP:FB CYC:9299
DACE  E8        INX                             A:80 X:FD Y:0 P:A4 SP:FB CYC:9302
DACF  CE 00 04  DEC $0400 = 81                  A:80 X:FE Y:0 P:A4 SP:FB CYC:9304
DAD2  CE 00 04  DEC $0400 = 80                  A:80 X:FE Y:0 P:A4 SP:FB CYC:9310
DAD5  D1 33     CMP ($33),Y = 0400 @ 0400 = 7F  A:80 X:FE Y:0 P:24 SP:FB CYC:9316
DAD7  90 04     BCC $DADD                       A:80 X:FE Y:0 P:25 SP:FB CYC:9321
DAD9  F0 02     BEQ $DADD                       A:80 X:FE Y:0 P:25 SP:FB CYC:9323
DADB  10 02     BPL $DADF                       A:80 X:FE Y:0 P:25 SP:FB CYC:9325
DADF  60        RTS                             A:80 X:FE Y:0 P:25 SP:FB CYC:9328
C61B  A5 00     LDA $00 = 00                    A:80 X:FE Y:0 P:25 SP:FD CYC:9334
C61D  85 10     STA $10 = 00                    A:0 X:FE Y:0 P:27 SP:FD CYC:9337
C61F  A9 00     LDA #$00                        A:0 X:FE Y:0 P:27 SP:FD CYC:9340
C621  85 00     STA $00 = 00                    A:0 X:FE Y:0 P:27 SP:FD CYC:9342
C623  20 E0 DA  JSR $DAE0                       A:0 X:FE Y:0 P:27 SP:FD CYC:9345
DAE0  A9 00     LDA #$00                        A:0 X:FE Y:0 P:27 SP:FB CYC:9351
DAE2  85 33     STA $33 = 00                    A:0 X:FE Y:0 P:27 SP:FB CYC:9353
DAE4  A9 04     LDA #$04                        A:0 X:FE Y:0 P:27 SP:FB CYC:9356
DAE6  85 34     STA $34 = 04                    A:4 X:FE Y:0 P:25 SP:FB CYC:9358
DAE8  A0 00     LDY #$00                        A:4 X:FE Y:0 P:25 SP:FB CYC:9361
DAEA  A2 01     LDX #$01                        A:4 X:FE Y:0 P:27 SP:FB CYC:9363
DAEC  24 01     BIT $01 = FF                    A:4 X:1 Y:0 P:25 SP:FB CYC:9365
DAEE  A9 40     LDA #$40                        A:4 X:1 Y:0 P:E5 SP:FB CYC:9368
DAF0  8D 00 04  STA $0400 = 7F                  A:40 X:1 Y:0 P:65 SP:FB CYC:9370
DAF3  38        SEC                             A:40 X:1 Y:0 P:65 SP:FB CYC:9374
DAF4  F1 33     SBC ($33),Y = 0400 @ 0400 = 40  A:40 X:1 Y:0 P:65 SP:FB CYC:9376
DAF6  30 0A     BMI $DB02                       A:0 X:1 Y:0 P:27 SP:FB CYC:9381
DAF8  90 08     BCC $DB02                       A:0 X:1 Y:0 P:27 SP:FB CYC:9383
DAFA  D0 06     BNE $DB02                       A:0 X:1 Y:0 P:27 SP:FB CYC:9385
DAFC  70 04     BVS $DB02                       A:0 X:1 Y:0 P:27 SP:FB CYC:9387
DAFE  C9 00     CMP #$00                        A:0 X:1 Y:0 P:27 SP:FB CYC:9389
DB00  F0 02     BEQ $DB04                       A:0 X:1 Y:0 P:27 SP:FB CYC:9391
DB04  E8        INX                             A:0 X:1 Y:0 P:27 SP:FB CYC:9394
DB05  B8        CLV                             A:0 X:2 Y:0 P:25 SP:FB CYC:9396
DB06  38        SEC                             A:0 X:2 Y:0 P:25 SP:FB CYC:9398
DB07  A9 40     LDA #$40                        A:0 X:2 Y:0 P:25 SP:FB CYC:9400
DB09  CE 00 04  DEC $0400 = 40                  A:40 X:2 Y:0 P:25 SP:FB CYC:9402
DB0C  F1 33     SBC ($33),Y = 0400 @ 0400 = 3F  A:40 X:2 Y:0 P:25 SP:FB CYC:9408
DB0E  F0 0A     BEQ $DB1A                       A:1 X:2 Y:0 P:25 SP:FB CYC:9413
DB10  30 08     BMI $DB1A                       A:1 X:2 Y:0 P:25 SP:FB CYC:9415
DB12  90 06     BCC $DB1A                       A:1 X:2 Y:0 P:25 SP:FB CYC:9417
DB14  70 04     BVS $DB1A                       A:1 X:2 Y:0 P:25 SP:FB CYC:9419
DB16  C9 01     CMP #$01                        A:1 X:2 Y:0 P:25 SP:FB CYC:9421
DB18  F0 02     BEQ $DB1C                       A:1 X:2 Y:0 P:27 SP:FB CYC:9423
DB1C  E8        INX                             A:1 X:2 Y:0 P:27 SP:FB CYC:9426
DB1D  A9 40     LDA #$40                        A:1 X:3 Y:0 P:25 SP:FB CYC:9428
DB1F  38        SEC                             A:40 X:3 Y:0 P:25 SP:FB CYC:9430
DB20  24 01     BIT $01 = FF                    A:40 X:3 Y:0 P:25 SP:FB CYC:9432
DB22  EE 00 04  INC $0400 = 3F                  A:40 X:3 Y:0 P:E5 SP:FB CYC:9435
DB25  EE 00 04  INC $0400 = 40                  A:40 X:3 Y:0 P:65 SP:FB CYC:9441
DB28  F1 33     SBC ($33),Y = 0400 @ 0400 = 41  A:40 X:3 Y:0 P:65 SP:FB CYC:9447
DB2A  B0 0A     BCS $DB36                       A:FF X:3 Y:0 P:A4 SP:FB CYC:9452
DB2C  F0 08     BEQ $DB36                       A:FF X:3 Y:0 P:A4 SP:FB CYC:9454
DB2E  10 06     BPL $DB36                       A:FF X:3 Y:0 P:A4 SP:FB CYC:9456
DB30  70 04     BVS $DB36                       A:FF X:3 Y:0 P:A4 SP:FB CYC:9458
DB32  C9 FF     CMP #$FF                        A:FF X:3 Y:0 P:A4 SP:FB CYC:9460
DB34  F0 02     BEQ $DB38                       A:FF X:3 Y:0 P:27 SP:FB CYC:9462
DB38  E8        INX                             A:FF X:3 Y:0 P:27 SP:FB CYC:9465
DB39  18        CLC                             A:FF X:4 Y:0 P:25 SP:FB CYC:9467
DB3A  A9 00     LDA #$00                        A:FF X:4 Y:0 P:24 SP:FB CYC:9469
DB3C  8D 00 04  STA $0400 = 41                  A:0 X:4 Y:0 P:26 SP:FB CYC:9471
DB3F  A9 80     LDA #$80                        A:0 X:4 Y:0 P:26 SP:FB CYC:9475
DB41  F1 33     SBC ($33),Y = 0400 @ 0400 = 00  A:80 X:4 Y:0 P:A4 SP:FB CYC:9477
DB43  90 04     BCC $DB49                       A:7F X:4 Y:0 P:65 SP:FB CYC:9482
DB45  C9 7F     CMP #$7F                        A:7F X:4 Y:0 P:65 SP:FB CYC:9484
DB47  F0 02     BEQ $DB4B                       A:7F X:4 Y:0 P:67 SP:FB CYC:9486
DB4B  E8        INX                             A:7F X:4 Y:0 P:67 SP:FB CYC:9489
DB4C  38        SEC                             A:7F X:5 Y:0 P:65 SP:FB CYC:9491
DB4D  A9 7F     LDA #$7F                        A:7F X:5 Y:0 P:65 SP:FB CYC:9493
DB4F  8D 00 04  STA $0400 = 00                  A:7F X:5 Y:0 P:65 SP:FB CYC:9495
DB52  A9 81     LDA #$81                        A:7F X:5 Y:0 P:65 SP:FB CYC:9499
DB54  F1 33     SBC ($33),Y = 0400 @ 0400 = 7F  A:81 X:5 Y:0 P:E5 SP:FB CYC:9501
DB56  50 06     BVC $DB5E                       A:2 X:5 Y:0 P:65 SP:FB CYC:9506
DB58  90 04     BCC $DB5E                       A:2 X:5 Y:0 P:65 SP:FB CYC:9508
DB5A  C9 02     CMP #$02                        A:2 X:5 Y:0 P:65 SP:FB CYC:9510
DB5C  F0 02     BEQ $DB60                       A:2 X:5 Y:0 P:67 SP:FB CYC:9512
DB60  E8        INX                             A:2 X:5 Y:0 P:67 SP:FB CYC:9515
DB61  A9 00     LDA #$00                        A:2 X:6 Y:0 P:65 SP:FB CYC:9517
DB63  A9 87     LDA #$87                        A:0 X:6 Y:0 P:67 SP:FB CYC:9519
DB65  91 33     STA ($33),Y = 0400 @ 0400 = 7F  A:87 X:6 Y:0 P:E5 SP:FB CYC:9521
DB67  AD 00 04  LDA $0400 = 87                  A:87 X:6 Y:0 P:E5 SP:FB CYC:9527
DB6A  C9 87     CMP #$87                        A:87 X:6 Y:0 P:E5 SP:FB CYC:9531
DB6C  F0 02     BEQ $DB70                       A:87 X:6 Y:0 P:67 SP:FB CYC:9533
DB70  E8        INX                             A:87 X:6 Y:0 P:67 SP:FB CYC:9536
DB71  A9 7E     LDA #$7E                        A:87 X:7 Y:0 P:65 SP:FB CYC:9538
DB73  8D 00 02  STA $0200 = 7F                  A:7E X:7 Y:0 P:65 SP:FB CYC:9540
DB76  A9 DB     LDA #$DB                        A:7E X:7 Y:0 P:65 SP:FB CYC:9544
DB78  8D 01 02  STA $0201 = 00                  A:DB X:7 Y:0 P:E5 SP:FB CYC:9546
DB7B  6C 00 02  JMP ($0200) = DB7E              A:DB X:7 Y:0 P:E5 SP:FB CYC:9550
DB7E  A9 00     LDA #$00                        A:DB X:7 Y:0 P:E5 SP:FB CYC:9555
DB80  8D FF 02  STA $02FF = 00                  A:0 X:7 Y:0 P:67 SP:FB CYC:9557
DB83  A9 01     LDA #$01                        A:0 X:7 Y:0 P:67 SP:FB CYC:9561
DB85  8D 00 03  STA $0300 = 89                  A:1 X:7 Y:0 P:65 SP:FB CYC:9563
DB88  A9 03     LDA #$03                        A:1 X:7 Y:0 P:65 SP:FB CYC:9567
DB8A  8D 00 02  STA $0200 = 7E                  A:3 X:7 Y:0 P:65 SP:FB CYC:9569
DB8D  A9 A9     LDA #$A9                        A:3 X:7 Y:0 P:65 SP:FB CYC:9573
DB8F  8D 00 01  STA $0100 = 00                  A:A9 X:7 Y:0 P:E5 SP:FB CYC:9575
DB92  A9 55     LDA #$55                        A:A9 X:7 Y:0 P:E5 SP:FB CYC:9579
DB94  8D 01 01  STA $0101 = 00                  A:55 X:7 Y:0 P:65 SP:FB CYC:9581
DB97  A9 60     LDA #$60                        A:55 X:7 Y:0 P:65 SP:FB CYC:9585
DB99  8D 02 01  STA $0102 = 00                  A:60 X:7 Y:0 P:65 SP:FB CYC:9587
DB9C  A9 A9     LDA #$A9                        A:60 X:7 Y:0 P:65 SP:FB CYC:9591
DB9E  8D 00 03  STA $0300 = 01                  A:A9 X:7 Y:0 P:E5 SP:FB CYC:9593
DBA1  A9 AA     LDA #$AA                        A:A9 X:7 Y:0 P:E5 SP:FB CYC:9597
DBA3  8D 01 03  STA $0301 = 00                  A:AA X:7 Y:0 P:E5 SP:FB CYC:9599
DBA6  A9 60     LDA #$60                        A:AA X:7 Y:0 P:E5 SP:FB CYC:9603
DBA8  8D 02 03  STA $0302 = 00                  A:60 X:7 Y:0 P:65 SP:FB CYC:9605
DBAB  20 B5 DB  JSR $DBB5                       A:60 X:7 Y:0 P:65 SP:FB CYC:9609
DBB5  6C FF 02  JMP ($02FF) = 0300              A:60 X:7 Y:0 P:65 SP:F9 CYC:9615
0300  A9 AA     LDA #$AA                        A:60 X:7 Y:0 P:65 SP:F9 CYC:9620
0302  60        RTS                             A:AA X:7 Y:0 P:E5 SP:F9 CYC:9622
DBAE  C9 AA     CMP #$AA                        A:AA X:7 Y:0 P:E5 SP:FB CYC:9628
DBB0  F0 02     BEQ $DBB4                       A:AA X:7 Y:0 P:67 SP:FB CYC:9630
DBB4  60        RTS                             A:AA X:7 Y:0 P:67 SP:FB CYC:9633
C626  20 4A DF  JSR $DF4A                       A:AA X:7 Y:0 P:67 SP:FD CYC:9639
DF4A  A9 89     LDA #$89                        A:AA X:7 Y:0 P:67 SP:FB CYC:9645
DF4C  8D 00 03  STA $0300 = A9                  A:89 X:7 Y:0 P:E5 SP:FB CYC:9647
DF4F  A9 A3     LDA #$A3                        A:89 X:7 Y:0 P:E5 SP:FB CYC:9651
DF51  85 33     STA $33 = 00                    A:A3 X:7 Y:0 P:E5 SP:FB CYC:9653
DF53  A9 12     LDA #$12                        A:A3 X:7 Y:0 P:E5 SP:FB CYC:9656
DF55  8D 45 02  STA $0245 = 12                  A:12 X:7 Y:0 P:65 SP:FB CYC:9658
DF58  A2 65     LDX #$65                        A:12 X:7 Y:0 P:65 SP:FB CYC:9662
DF5A  A0 00     LDY #$00                        A:12 X:65 Y:0 P:65 SP:FB CYC:9664
DF5C  38        SEC                             A:12 X:65 Y:0 P:67 SP:FB CYC:9666
DF5D  A9 00     LDA #$00                        A:12 X:65 Y:0 P:67 SP:FB CYC:9668
DF5F  B8        CLV                             A:0 X:65 Y:0 P:67 SP:FB CYC:9670
DF60  B9 00 03  LDA $0300,Y @ 0300 = 89         A:0 X:65 Y:0 P:27 SP:FB CYC:9672
DF63  F0 0C     BEQ $DF71                       A:89 X:65 Y:0 P:A5 SP:FB CYC:9676
DF65  90 0A     BCC $DF71                       A:89 X:65 Y:0 P:A5 SP:FB CYC:9678
DF67  70 08     BVS $DF71                       A:89 X:65 Y:0 P:A5 SP:FB CYC:9680
DF69  C9 89     CMP #$89                        A:89 X:65 Y:0 P:A5 SP:FB CYC:9682
DF6B  D0 04     BNE $DF71                       A:89 X:65 Y:0 P:27 SP:FB CYC:9684
DF6D  E0 65     CPX #$65                        A:89 X:65 Y:0 P:27 SP:FB CYC:9686
DF6F  F0 04     BEQ $DF75                       A:89 X:65 Y:0 P:27 SP:FB CYC:9688
DF75  A9 FF     LDA #$FF                        A:89 X:65 Y:0 P:27 SP:FB CYC:9691
DF77  85 01     STA $01 = FF                    A:FF X:65 Y:0 P:A5 SP:FB CYC:9693
DF79  24 01     BIT $01 = FF                    A:FF X:65 Y:0 P:A5 SP:FB CYC:9696
DF7B  A0 34     LDY #$34                        A:FF X:65 Y:0 P:E5 SP:FB CYC:9699
DF7D  B9 FF FF  LDA $FFFF,Y @ 0033 = A3         A:FF X:65 Y:34 P:65 SP:FB CYC:9701
DF80  C9 A3     CMP #$A3                        A:A3 X:65 Y:34 P:E5 SP:FB CYC:9706
DF82  D0 02     BNE $DF86                       A:A3 X:65 Y:34 P:67 SP:FB CYC:9708
DF84  B0 04     BCS $DF8A                       A:A3 X:65 Y:34 P:67 SP:FB CYC:9710
DF8A  A9 46     LDA #$46                        A:A3 X:65 Y:34 P:67 SP:FB CYC:9713
DF8C  85 FF     STA $FF = 46                    A:46 X:65 Y:34 P:65 SP:FB CYC:9715
DF8E  A0 FF     LDY #$FF                        A:46 X:65 Y:34 P:65 SP:FB CYC:9718
DF90  B9 46 01  LDA $0146,Y @ 0245 = 12         A:46 X:65 Y:FF P:E5 SP:FB CYC:9720
DF93  C9 12     CMP #$12                        A:12 X:65 Y:FF P:65 SP:FB CYC:9725
DF95  F0 04     BEQ $DF9B                       A:12 X:65 Y:FF P:67 SP:FB CYC:9727
DF9B  A2 39     LDX #$39                        A:12 X:65 Y:FF P:67 SP:FB CYC:9730
DF9D  18        CLC                             A:12 X:39 Y:FF P:65 SP:FB CYC:9732
DF9E  A9 FF     LDA #$FF                        A:12 X:39 Y:FF P:64 SP:FB CYC:9734
DFA0  85 01     STA $01 = FF                    A:FF X:39 Y:FF P:E4 SP:FB CYC:9736
DFA2  24 01     BIT $01 = FF                    A:FF X:39 Y:FF P:E4 SP:FB CYC:9739
DFA4  A9 AA     LDA #$AA                        A:FF X:39 Y:FF P:E4 SP:FB CYC:9742
DFA6  8D 00 04  STA $0400 = 87                  A:AA X:39 Y:FF P:E4 SP:FB CYC:9744
DFA9  A9 55     LDA #$55                        A:AA X:39 Y:FF P:E4 SP:FB CYC:9748
DFAB  A0 00     LDY #$00                        A:55 X:39 Y:FF P:64 SP:FB CYC:9750
DFAD  19 00 04  ORA $0400,Y @ 0400 = AA         A:55 X:39 Y:0 P:66 SP:FB CYC:9752
DFB0  B0 08     BCS $DFBA                       A:FF X:39 Y:0 P:E4 SP:FB CYC:9756
DFB2  10 06     BPL $DFBA                       A:FF X:39 Y:0 P:E4 SP:FB CYC:9758
DFB4  C9 FF     CMP #$FF                        A:FF X:39 Y:0 P:E4 SP:FB CYC:9760
DFB6  D0 02     BNE $DFBA                       A:FF X:39 Y:0 P:67 SP:FB CYC:9762
DFB8  70 02     BVS $DFBC                       A:FF X:39 Y:0 P:67 SP:FB CYC:9764
DFBC  E8        INX                             A:FF X:39 Y:0 P:67 SP:FB CYC:9767
DFBD  38        SEC                             A:FF X:3A Y:0 P:65 SP:FB CYC:9769
DFBE  B8        CLV                             A:FF X:3A Y:0 P:65 SP:FB CYC:9771
DFBF  A9 00     LDA #$00                        A:FF X:3A Y:0 P:25 SP:FB CYC:9773
DFC1  19 00 04  ORA $0400,Y @ 0400 = AA         A:0 X:3A Y:0 P:27 SP:FB CYC:9775
DFC4  F0 06     BEQ $DFCC                       A:AA X:3A Y:0 P:A5 SP:FB CYC:9779
DFC6  70 04     BVS $DFCC                       A:AA X:3A Y:0 P:A5 SP:FB CYC:9781
DFC8  90 02     BCC $DFCC                       A:AA X:3A Y:0 P:A5 SP:FB CYC:9783
DFCA  30 02     BMI $DFCE                       A:AA X:3A Y:0 P:A5 SP:FB CYC:9785
DFCE  E8        INX                             A:AA X:3A Y:0 P:A5 SP:FB CYC:9788
DFCF  18        CLC                             A:AA X:3B Y:0 P:25 SP:FB CYC:9790
DFD0  24 01     BIT $01 = FF                    A:AA X:3B Y:0 P:24 SP:FB CYC:9792
DFD2  A9 55     LDA #$55                        A:AA X:3B Y:0 P:E4 SP:FB CYC:9795
DFD4  39 00 04  AND $0400,Y @ 0400 = AA         A:55 X:3B Y:0 P:64 SP:FB CYC:9797
DFD7  D0 06     BNE $DFDF                       A:0 X:3B Y:0 P:66 SP:FB CYC:9801
DFD9  50 04     BVC $DFDF                       A:0 X:3B Y:0 P:66 SP:FB CYC:9803
DFDB  B0 02     BCS $DFDF                       A:0 X:3B Y:0 P:66 SP:FB CYC:9805
DFDD  10 02     BPL $DFE1                       A:0 X:3B Y:0 P:66 SP:FB CYC:9807
DFE1  E8        INX                             A:0 X:3B Y:0 P:66 SP:FB CYC:9810
DFE2  38        SEC                             A:0 X:3C Y:0 P:64 SP:FB CYC:9812
DFE3  B8        CLV                             A:0 X:3C Y:0 P:65 SP:FB CYC:9814
DFE4  A9 EF     LDA #$EF                        A:0 X:3C Y:0 P:25 SP:FB CYC:9816
DFE6  8D 00 04  STA $0400 = AA                  A:EF X:3C Y:0 P:A5 SP:FB CYC:9818
DFE9  A9 F8     LDA #$F8                        A:EF X:3C Y:0 P:A5 SP:FB CYC:9822
DFEB  39 00 04  AND $0400,Y @ 0400 = EF         A:F8 X:3C Y:0 P:A5 SP:FB CYC:9824
DFEE  90 08     BCC $DFF8                       A:E8 X:3C Y:0 P:A5 SP:FB CYC:9828
DFF0  10 06     BPL $DFF8                       A:E8 X:3C Y:0 P:A5 SP:FB CYC:9830
DFF2  C9 E8     CMP #$E8                        A:E8 X:3C Y:0 P:A5 SP:FB CYC:9832
DFF4  D0 02     BNE $DFF8                       A:E8 X:3C Y:0 P:27 SP:FB CYC:9834
DFF6  50 02     BVC $DFFA                       A:E8 X:3C Y:0 P:27 SP:FB CYC:9836
DFFA  E8        INX                             A:E8 X:3C Y:0 P:27 SP:FB CYC:9839
DFFB  18        CLC                             A:E8 X:3D Y:0 P:25 SP:FB CYC:9841
DFFC  24 01     BIT $01 = FF                    A:E8 X:3D Y:0 P:24 SP:FB CYC:9843
DFFE  A9 AA     LDA #$AA                        A:E8 X:3D Y:0 P:E4 SP:FB CYC:9846
E000  8D 00 04  STA $0400 = EF                  A:AA X:3D Y:0 P:E4 SP:FB CYC:9848
E003  A9 5F     LDA #$5F                        A:AA X:3D Y:0 P:E4 SP:FB CYC:9852
E005  59 00 04  EOR $0400,Y @ 0400 = AA         A:5F X:3D Y:0 P:64 SP:FB CYC:9854
E008  B0 08     BCS $E012                       A:F5 X:3D Y:0 P:E4 SP:FB CYC:9858
E00A  10 06     BPL $E012                       A:F5 X:3D Y:0 P:E4 SP:FB CYC:9860
E00C  C9 F5     CMP #$F5                        A:F5 X:3D Y:0 P:E4 SP:FB CYC:9862
E00E  D0 02     BNE $E012                       A:F5 X:3D Y:0 P:67 SP:FB CYC:9864
E010  70 02     BVS $E014                       A:F5 X:3D Y:0 P:67 SP:FB CYC:9866
E014  E8        INX                             A:F5 X:3D Y:0 P:67 SP:FB CYC:9869
E015  38        SEC                             A:F5 X:3E Y:0 P:65 SP:FB CYC:9871
E016  B8        CLV                             A:F5 X:3E Y:0 P:65 SP:FB CYC:9873
E017  A9 70     LDA #$70                        A:F5 X:3E Y:0 P:25 SP:FB CYC:9875
E019  8D 00 04  STA $0400 = AA                  A:70 X:3E Y:0 P:25 SP:FB CYC:9877
E01C  59 00 04  EOR $0400,Y @ 0400 = 70         A:70 X:3E Y:0 P:25 SP:FB CYC:9881
E01F  D0 06     BNE $E027                       A:0 X:3E Y:0 P:27 SP:FB CYC:9885
E021  70 04     BVS $E027                       A:0 X:3E Y:0 P:27 SP:FB CYC:9887
E023  90 02     BCC $E027                       A:0 X:3E Y:0 P:27 SP:FB CYC:9889
E025  10 02     BPL $E029                       A:0 X:3E Y:0 P:27 SP:FB CYC:9891
E029  E8        INX                             A:0 X:3E Y:0 P:27 SP:FB CYC:9894
E02A  18        CLC                             A:0 X:3F Y:0 P:25 SP:FB CYC:9896
E02B  24 01     BIT $01 = FF                    A:0 X:3F Y:0 P:24 SP:FB CYC:9898
E02D  A9 69     LDA #$69                        A:0 X:3F Y:0 P:E6 SP:FB CYC:9901
E02F  8D 00 04  STA $0400 = 70                  A:69 X:3F Y:0 P:64 SP:FB CYC:9903
E032  A9 00     LDA #$00                        A:69 X:3F Y:0 P:64 SP:FB CYC:9907
E034  79 00 04  ADC $0400,Y @ 0400 = 69         A:0 X:3F Y:0 P:66 SP:FB CYC:9909
E037  30 08     BMI $E041                       A:69 X:3F Y:0 P:24 SP:FB CYC:9913
E039  B0 06     BCS $E041                       A:69 X:3F Y:0 P:24 SP:FB CYC:9915
E03B  C9 69     CMP #$69                        A:69 X:3F Y:0 P:24 SP:FB CYC:9917
E03D  D0 02     BNE $E041                       A:69 X:3F Y:0 P:27 SP:FB CYC:9919
E03F  50 02     BVC $E043                       A:69 X:3F Y:0 P:27 SP:FB CYC:9921
E043  E8        INX                             A:69 X:3F Y:0 P:27 SP:FB CYC:9924
E044  38        SEC                             A:69 X:40 Y:0 P:25 SP:FB CYC:9926
E045  24 01     BIT $01 = FF                    A:69 X:40 Y:0 P:25 SP:FB CYC:9928
E047  A9 00     LDA #$00                        A:69 X:40 Y:0 P:E5 SP:FB CYC:9931
E049  79 00 04  ADC $0400,Y @ 0400 = 69         A:0 X:40 Y:0 P:67 SP:FB CYC:9933
E04C  30 08     BMI $E056                       A:6A X:40 Y:0 P:24 SP:FB CYC:9937
E04E  B0 06     BCS $E056                       A:6A X:40 Y:0 P:24 SP:FB CYC:9939
E050  C9 6A     CMP #$6A                        A:6A X:40 Y:0 P:24 SP:FB CYC:9941
E052  D0 02     BNE $E056                       A:6A X:40 Y:0 P:27 SP:FB CYC:9943
E054  50 02     BVC $E058                       A:6A X:40 Y:0 P:27 SP:FB CYC:9945
E058  E8        INX                             A:6A X:40 Y:0 P:27 SP:FB CYC:9948
E059  38        SEC                             A:6A X:41 Y:0 P:25 SP:FB CYC:9950
E05A  B8        CLV                             A:6A X:41 Y:0 P:25 SP:FB CYC:9952
E05B  A9 7F     LDA #$7F                        A:6A X:41 Y:0 P:25 SP:FB CYC:9954
E05D  8D 00 04  STA $0400 = 69                  A:7F X:41 Y:0 P:25 SP:FB CYC:9956
E060  79 00 04  ADC $0400,Y @ 0400 = 7F         A:7F X:41 Y:0 P:25 SP:FB CYC:9960
E063  10 08     BPL $E06D                       A:FF X:41 Y:0 P:E4 SP:FB CYC:9964
E065  B0 06     BCS $E06D                       A:FF X:41 Y:0 P:E4 SP:FB CYC:9966
E067  C9 FF     CMP #$FF                        A:FF X:41 Y:0 P:E4 SP:FB CYC:9968
E069  D0 02     BNE $E06D                       A:FF X:41 Y:0 P:67 SP:FB CYC:9970
E06B  70 02     BVS $E06F                       A:FF X:41 Y:0 P:67 SP:FB CYC:9972
E06F  E8        INX                             A:FF X:41 Y:0 P:67 SP:FB CYC:9975
E070  18        CLC                             A:FF X:42 Y:0 P:65 SP:FB CYC:9977
E071  24 01     BIT $01 = FF                    A:FF X:42 Y:0 P:64 SP:FB CYC:9979
E073  A9 80     LDA #$80                        A:FF X:42 Y:0 P:E4 SP:FB CYC:9982
E075  8D 00 04  STA $0400 = 7F                  A:80 X:42 Y:0 P:E4 SP:FB CYC:9984
E078  A9 7F     LDA #$7F                        A:80 X:42 Y:0 P:E4 SP:FB CYC:9988
E07A  79 00 04  ADC $0400,Y @ 0400 = 80         A:7F X:42 Y:0 P:64 SP:FB CYC:9990
E07D  10 08     BPL $E087                       A:FF X:42 Y:0 P:A4 SP:FB CYC:9994
E07F  B0 06     BCS $E087                       A:FF X:42 Y:0 P:A4 SP:FB CYC:9996
E081  C9 FF     CMP #$FF                        A:FF X:42 Y:0 P:A4 SP:FB CYC:9998
E083  D0 02     BNE $E087                       A:FF X:42 Y:0 P:27 SP:FB CYC:10000
E085  50 02     BVC $E089                       A:FF X:42 Y:0 P:27 SP:FB CYC:10002
E089  E8        INX                             A:FF X:42 Y:0 P:27 SP:FB CYC:10005
E08A  38        SEC                             A:FF X:43 Y:0 P:25 SP:FB CYC:10007
E08B  B8        CLV                             A:FF X:43 Y:0 P:25 SP:FB CYC:10009
E08C  A9 80     LDA #$80                        A:FF X:43 Y:0 P:25 SP:FB CYC:10011
E08E  8D 00 04  STA $0400 = 80                  A:80 X:43 Y:0 P:A5 SP:FB CYC:10013
E091  A9 7F     LDA #$7F                        A:80 X:43 Y:0 P:A5 SP:FB CYC:10017
E093  79 00 04  ADC $0400,Y @ 0400 = 80         A:7F X:43 Y:0 P:25 SP:FB CYC:10019
E096  D0 06     BNE $E09E                       A:0 X:43 Y:0 P:27 SP:FB CYC:10023
E098  30 04     BMI $E09E                       A:0 X:43 Y:0 P:27 SP:FB CYC:10025
E09A  70 02     BVS $E09E                       A:0 X:43 Y:0 P:27 SP:FB CYC:10027
E09C  B0 02     BCS $E0A0                       A:0 X:43 Y:0 P:27 SP:FB CYC:10029
E0A0  E8        INX                             A:0 X:43 Y:0 P:27 SP:FB CYC:10032
E0A1  24 01     BIT $01 = FF                    A:0 X:44 Y:0 P:25 SP:FB CYC:10034
E0A3  A9 40     LDA #$40                        A:0 X:44 Y:0 P:E7 SP:FB CYC:10037
E0A5  8D 00 04  STA $0400 = 80                  A:40 X:44 Y:0 P:65 SP:FB CYC:10039
E0A8  D9 00 04  CMP $0400,Y @ 0400 = 40         A:40 X:44 Y:0 P:65 SP:FB CYC:10043
E0AB  30 06     BMI $E0B3                       A:40 X:44 Y:0 P:67 SP:FB CYC:10047
E0AD  90 04     BCC $E0B3                       A:40 X:44 Y:0 P:67 SP:FB CYC:10049
E0AF  D0 02     BNE $E0B3                       A:40 X:44 Y:0 P:67 SP:FB CYC:10051
E0B1  70 02     BVS $E0B5                       A:40 X:44 Y:0 P:67 SP:FB CYC:10053
E0B5  E8        INX                             A:40 X:44 Y:0 P:67 SP:FB CYC:10056
E0B6  B8        CLV                             A:40 X:45 Y:0 P:65 SP:FB CYC:10058
E0B7  CE 00 04  DEC $0400 = 40                  A:40 X:45 Y:0 P:25 SP:FB CYC:10060
E0BA  D9 00 04  CMP $0400,Y @ 0400 = 3F         A:40 X:45 Y:0 P:25 SP:FB CYC:10066
E0BD  F0 06     BEQ $E0C5                       A:40 X:45 Y:0 P:25 SP:FB CYC:10070
E0BF  30 04     BMI $E0C5                       A:40 X:45 Y:0 P:25 SP:FB CYC:10072
E0C1  90 02     BCC $E0C5                       A:40 X:45 Y:0 P:25 SP:FB CYC:10074
E0C3  50 02     BVC $E0C7                       A:40 X:45 Y:0 P:25 SP:FB CYC:10076
E0C7  E8        INX                             A:40 X:45 Y:0 P:25 SP:FB CYC:10079
E0C8  EE 00 04  INC $0400 = 3F                  A:40 X:46 Y:0 P:25 SP:FB CYC:10081
E0CB  EE 00 04  INC $0400 = 40                  A:40 X:46 Y:0 P:25 SP:FB CYC:10087
E0CE  D9 00 04  CMP $0400,Y @ 0400 = 41         A:40 X:46 Y:0 P:25 SP:FB CYC:10093
E0D1  F0 02     BEQ $E0D5                       A:40 X:46 Y:0 P:A4 SP:FB CYC:10097
E0D3  30 02     BMI $E0D7                       A:40 X:46 Y:0 P:A4 SP:FB CYC:10099
E0D7  E8        INX                             A:40 X:46 Y:0 P:A4 SP:FB CYC:10102
E0D8  A9 00     LDA #$00                        A:40 X:47 Y:0 P:24 SP:FB CYC:10104
E0DA  8D 00 04  STA $0400 = 41                  A:0 X:47 Y:0 P:26 SP:FB CYC:10106
E0DD  A9 80     LDA #$80                        A:0 X:47 Y:0 P:26 SP:FB CYC:10110
E0DF  D9 00 04  CMP $0400,Y @ 0400 = 00         A:80 X:47 Y:0 P:A4 SP:FB CYC:10112
E0E2  F0 04     BEQ $E0E8                       A:80 X:47 Y:0 P:A5 SP:FB CYC:10116
E0E4  10 02     BPL $E0E8                       A:80 X:47 Y:0 P:A5 SP:FB CYC:10118
E0E6  B0 02     BCS $E0EA                       A:80 X:47 Y:0 P:A5 SP:FB CYC:10120
E0EA  E8        INX                             A:80 X:47 Y:0 P:A5 SP:FB CYC:10123
E0EB  A0 80     LDY #$80                        A:80 X:48 Y:0 P:25 SP:FB CYC:10125
E0ED  8C 00 04  STY $0400 = 00                  A:80 X:48 Y:80 P:A5 SP:FB CYC:10127
E0F0  A0 00     LDY #$00                        A:80 X:48 Y:80 P:A5 SP:FB CYC:10131
E0F2  D9 00 04  CMP $0400,Y @ 0400 = 80         A:80 X:48 Y:0 P:27 SP:FB CYC:10133
E0F5  D0 04     BNE $E0FB                       A:80 X:48 Y:0 P:27 SP:FB CYC:10137
E0F7  30 02     BMI $E0FB                       A:80 X:48 Y:0 P:27 SP:FB CYC:10139
E0F9  B0 02     BCS $E0FD                       A:80 X:48 Y:0 P:27 SP:FB CYC:10141
E0FD  E8        INX                             A:80 X:48 Y:0 P:27 SP:FB CYC:10144
E0FE  EE 00 04  INC $0400 = 80                  A:80 X:49 Y:0 P:25 SP:FB CYC:10146
E101  D9 00 04  CMP $0400,Y @ 0400 = 81         A:80 X:49 Y:0 P:A5 SP:FB CYC:10152
E104  B0 04     BCS $E10A                       A:80 X:49 Y:0 P:A4 SP:FB CYC:10156
E106  F0 02     BEQ $E10A                       A:80 X:49 Y:0 P:A4 SP:FB CYC:10158
E108  30 02     BMI $E10C                       A:80 X:49 Y:0 P:A4 SP:FB CYC:10160
E10C  E8        INX                             A:80 X:49 Y:0 P:A4 SP:FB CYC:10163
E10D  CE 00 04  DEC $0400 = 81                  A:80 X:4A Y:0 P:24 SP:FB CYC:10165
E110  CE 00 04  DEC $0400 = 80                  A:80 X:4A Y:0 P:A4 SP:FB CYC:10171
E113  D9 00 04  CMP $0400,Y @ 0400 = 7F         A:80 X:4A Y:0 P:24 SP:FB CYC:10177
E116  90 04     BCC $E11C                       A:80 X:4A Y:0 P:25 SP:FB CYC:10181
E118  F0 02     BEQ $E11C                       A:80 X:4A Y:0 P:25 SP:FB CYC:10183
E11A  10 02     BPL $E11E                       A:80 X:4A Y:0 P:25 SP:FB CYC:10185
E11E  E8        INX                             A:80 X:4A Y:0 P:25 SP:FB CYC:10188
E11F  24 01     BIT $01 = FF                    A:80 X:4B Y:0 P:25 SP:FB CYC:10190
E121  A9 40     LDA #$40                        A:80 X:4B Y:0 P:E5 SP:FB CYC:10193
E123  8D 00 04  STA $0400 = 7F                  A:40 X:4B Y:0 P:65 SP:FB CYC:10195
E126  38        SEC                             A:40 X:4B Y:0 P:65 SP:FB CYC:10199
E127  F9 00 04  SBC $0400,Y @ 0400 = 40         A:40 X:4B Y:0 P:65 SP:FB CYC:10201
E12A  30 0A     BMI $E136                       A:0 X:4B Y:0 P:27 SP:FB CYC:10205
E12C  90 08     BCC $E136                       A:0 X:4B Y:0 P:27 SP:FB CYC:10207
E12E  D0 06     BNE $E136                       A:0 X:4B Y:0 P:27 SP:FB CYC:10209
E130  70 04     BVS $E136                       A:0 X:4B Y:0 P:27 SP:FB CYC:10211
E132  C9 00     CMP #$00                        A:0 X:4B Y:0 P:27 SP:FB CYC:10213
E134  F0 02     BEQ $E138                       A:0 X:4B Y:0 P:27 SP:FB CYC:10215
E138  E8        INX                             A:0 X:4B Y:0 P:27 SP:FB CYC:10218
E139  B8        CLV                             A:0 X:4C Y:0 P:25 SP:FB CYC:10220
E13A  38        SEC                             A:0 X:4C Y:0 P:25 SP:FB CYC:10222
E13B  A9 40     LDA #$40                        A:0 X:4C Y:0 P:25 SP:FB CYC:10224
E13D  CE 00 04  DEC $0400 = 40                  A:40 X:4C Y:0 P:25 SP:FB CYC:10226
E140  F9 00 04  SBC $0400,Y @ 0400 = 3F         A:40 X:4C Y:0 P:25 SP:FB CYC:10232
E143  F0 0A     BEQ $E14F                       A:1 X:4C Y:0 P:25 SP:FB CYC:10236
E145  30 08     BMI $E14F                       A:1 X:4C Y:0 P:25 SP:FB CYC:10238
E147  90 06     BCC $E14F                       A:1 X:4C Y:0 P:25 SP:FB CYC:10240
E149  70 04     BVS $E14F                       A:1 X:4C Y:0 P:25 SP:FB CYC:10242
E14B  C9 01     CMP #$01                        A:1 X:4C Y:0 P:25 SP:FB CYC:10244
E14D  F0 02     BEQ $E151                       A:1 X:4C Y:0 P:27 SP:FB CYC:10246
E151  E8        INX                             A:1 X:4C Y:0 P:27 SP:FB CYC:10249
E152  A9 40     LDA #$40                        A:1 X:4D Y:0 P:25 SP:FB CYC:10251
E154  38        SEC                             A:40 X:4D Y:0 P:25 SP:FB CYC:10253
E155  24 01     BIT $01 = FF                    A:40 X:4D Y:0 P:25 SP:FB CYC:10255
E157  EE 00 04  INC $0400 = 3F                  A:40 X:4D Y:0 P:E5 SP:FB CYC:10258
E15A  EE 00 04  INC $0400 = 40                  A:40 X:4D Y:0 P:65 SP:FB CYC:10264
E15D  F9 00 04  SBC $0400,Y @ 0400 = 41         A:40 X:4D Y:0 P:65 SP:FB CYC:10270
E160  B0 0A     BCS $E16C                       A:FF X:4D Y:0 P:A4 SP:FB CYC:10274
E162  F0 08     BEQ $E16C                       A:FF X:4D Y:0 P:A4 SP:FB CYC:10276
E164  10 06     BPL $E16C                       A:FF X:4D Y:0 P:A4 SP:FB CYC:10278
E166  70 04     BVS $E16C                       A:FF X:4D Y:0 P:A4 SP:FB CYC:10280
E168  C9 FF     CMP #$FF                        A:FF X:4D Y:0 P:A4 SP:FB CYC:10282
E16A  F0 02     BEQ $E16E                       A:FF X:4D Y:0 P:27 SP:FB CYC:10284
E16E  E8        INX                             A:FF X:4D Y:0 P:27 SP:FB CYC:10287
E16F  18        CLC                             A:FF X:4E Y:0 P:25 SP:FB CYC:10289
E170  A9 00     LDA #$00                        A:FF X:4E Y:0 P:24 SP:FB CYC:10291
E172  8D 00 04  STA $0400 = 41                  A:0 X:4E Y:0 P:26 SP:FB CYC:10293
E175  A9 80     LDA #$80                        A:0 X:4E Y:0 P:26 SP:FB CYC:10297
E177  F9 00 04  SBC $0400,Y @ 0400 = 00         A:80 X:4E Y:0 P:A4 SP:FB CYC:10299
E17A  90 04     BCC $E180                       A:7F X:4E Y:0 P:65 SP:FB CYC:10303
E17C  C9 7F     CMP #$7F                        A:7F X:4E Y:0 P:65 SP:FB CYC:10305
E17E  F0 02     BEQ $E182                       A:7F X:4E Y:0 P:67 SP:FB CYC:10307
E182  E8        INX                             A:7F X:4E Y:0 P:67 SP:FB CYC:10310
E183  38        SEC                             A:7F X:4F Y:0 P:65 SP:FB CYC:10312
E184  A9 7F     LDA #$7F                        A:7F X:4F Y:0 P:65 SP:FB CYC:10314
E186  8D 00 04  STA $0400 = 00                  A:7F X:4F Y:0 P:65 SP:FB CYC:10316
E189  A9 81     LDA #$81                        A:7F X:4F Y:0 P:65 SP:FB CYC:10320
E18B  F9 00 04  SBC $0400,Y @ 0400 = 7F         A:81 X:4F Y:0 P:E5 SP:FB CYC:10322
E18E  50 06     BVC $E196                       A:2 X:4F Y:0 P:65 SP:FB CYC:10326
E190  90 04     BCC $E196                       A:2 X:4F Y:0 P:65 SP:FB CYC:10328
E192  C9 02     CMP #$02                        A:2 X:4F Y:0 P:65 SP:FB CYC:10330
E194  F0 02     BEQ $E198                       A:2 X:4F Y:0 P:67 SP:FB CYC:10332
E198  E8        INX                             A:2 X:4F Y:0 P:67 SP:FB CYC:10335
E199  A9 00     LDA #$00                        A:2 X:50 Y:0 P:65 SP:FB CYC:10337
E19B  A9 87     LDA #$87                        A:0 X:50 Y:0 P:67 SP:FB CYC:10339
E19D  99 00 04  STA $0400,Y @ 0400 = 7F         A:87 X:50 Y:0 P:E5 SP:FB CYC:10341
E1A0  AD 00 04  LDA $0400 = 87                  A:87 X:50 Y:0 P:E5 SP:FB CYC:10346
E1A3  C9 87     CMP #$87                        A:87 X:50 Y:0 P:E5 SP:FB CYC:10350
E1A5  F0 02     BEQ $E1A9                       A:87 X:50 Y:0 P:67 SP:FB CYC:10352
E1A9  60        RTS                             A:87 X:50 Y:0 P:67 SP:FB CYC:10355
C629  20 B8 DB  JSR $DBB8                       A:87 X:50 Y:0 P:67 SP:FD CYC:10361
DBB8  A9 FF     LDA #$FF                        A:87 X:50 Y:0 P:67 SP:FB CYC:10367
DBBA  85 01     STA $01 = FF                    A:FF X:50 Y:0 P:E5 SP:FB CYC:10369
DBBC  A9 AA     LDA #$AA                        A:FF X:50 Y:0 P:E5 SP:FB CYC:10372
DBBE  85 33     STA $33 = A3                    A:AA X:50 Y:0 P:E5 SP:FB CYC:10374
DBC0  A9 BB     LDA #$BB                        A:AA X:50 Y:0 P:E5 SP:FB CYC:10377
DBC2  85 89     STA $89 = 00                    A:BB X:50 Y:0 P:E5 SP:FB CYC:10379
DBC4  A2 00     LDX #$00                        A:BB X:50 Y:0 P:E5 SP:FB CYC:10382
DBC6  A9 66     LDA #$66                        A:BB X:0 Y:0 P:67 SP:FB CYC:10384
DBC8  24 01     BIT $01 = FF                    A:66 X:0 Y:0 P:65 SP:FB CYC:10386
DBCA  38        SEC                             A:66 X:0 Y:0 P:E5 SP:FB CYC:10389
DBCB  A0 00     LDY #$00                        A:66 X:0 Y:0 P:E5 SP:FB CYC:10391
DBCD  B4 33     LDY $33,X @ 33 = AA             A:66 X:0 Y:0 P:67 SP:FB CYC:10393
DBCF  10 12     BPL $DBE3                       A:66 X:0 Y:AA P:E5 SP:FB CYC:10397
DBD1  F0 10     BEQ $DBE3                       A:66 X:0 Y:AA P:E5 SP:FB CYC:10399
DBD3  50 0E     BVC $DBE3                       A:66 X:0 Y:AA P:E5 SP:FB CYC:10401
DBD5  90 0C     BCC $DBE3                       A:66 X:0 Y:AA P:E5 SP:FB CYC:10403
DBD7  C9 66     CMP #$66                        A:66 X:0 Y:AA P:E5 SP:FB CYC:10405
DBD9  D0 08     BNE $DBE3                       A:66 X:0 Y:AA P:67 SP:FB CYC:10407
DBDB  E0 00     CPX #$00                        A:66 X:0 Y:AA P:67 SP:FB CYC:10409
DBDD  D0 04     BNE $DBE3                       A:66 X:0 Y:AA P:67 SP:FB CYC:10411
DBDF  C0 AA     CPY #$AA                        A:66 X:0 Y:AA P:67 SP:FB CYC:10413
DBE1  F0 04     BEQ $DBE7                       A:66 X:0 Y:AA P:67 SP:FB CYC:10415
DBE7  A2 8A     LDX #$8A                        A:66 X:0 Y:AA P:67 SP:FB CYC:10418
DBE9  A9 66     LDA #$66                        A:66 X:8A Y:AA P:E5 SP:FB CYC:10420
DBEB  B8        CLV                             A:66 X:8A Y:AA P:65 SP:FB CYC:10422
DBEC  18        CLC                             A:66 X:8A Y:AA P:25 SP:FB CYC:10424
DBED  A0 00     LDY #$00                        A:66 X:8A Y:AA P:24 SP:FB CYC:10426
DBEF  B4 FF     LDY $FF,X @ 89 = BB             A:66 X:8A Y:0 P:26 SP:FB CYC:10428
DBF1  10 12     BPL $DC05                       A:66 X:8A Y:BB P:A4 SP:FB CYC:10432
DBF3  F0 10     BEQ $DC05                       A:66 X:8A Y:BB P:A4 SP:FB CYC:10434
DBF5  70 0E     BVS $DC05                       A:66 X:8A Y:BB P:A4 SP:FB CYC:10436
DBF7  B0 0C     BCS $DC05                       A:66 X:8A Y:BB P:A4 SP:FB CYC:10438
DBF9  C0 BB     CPY #$BB                        A:66 X:8A Y:BB P:A4 SP:FB CYC:10440
DBFB  D0 08     BNE $DC05                       A:66 X:8A Y:BB P:27 SP:FB CYC:10442
DBFD  C9 66     CMP #$66                        A:66 X:8A Y:BB P:27 SP:FB CYC:10444
DBFF  D0 04     BNE $DC05                       A:66 X:8A Y:BB P:27 SP:FB CYC:10446
DC01  E0 8A     CPX #$8A                        A:66 X:8A Y:BB P:27 SP:FB CYC:10448
DC03  F0 04     BEQ $DC09                       A:66 X:8A Y:BB P:27 SP:FB CYC:10450
DC09  24 01     BIT $01 = FF                    A:66 X:8A Y:BB P:27 SP:FB CYC:10453
DC0B  38        SEC                             A:66 X:8A Y:BB P:E5 SP:FB CYC:10456
DC0C  A0 44     LDY #$44                        A:66 X:8A Y:BB P:E5 SP:FB CYC:10458
DC0E  A2 00     LDX #$00                        A:66 X:8A Y:44 P:65 SP:FB CYC:10460
DC10  94 33     STY $33,X @ 33 = AA             A:66 X:0 Y:44 P:67 SP:FB CYC:10462
DC12  A5 33     LDA $33 = 44                    A:66 X:0 Y:44 P:67 SP:FB CYC:10466
DC14  90 18     BCC $DC2E                       A:44 X:0 Y:44 P:65 SP:FB CYC:10469
DC16  C9 44     CMP #$44                        A:44 X:0 Y:44 P:65 SP:FB CYC:10471
DC18  D0 14     BNE $DC2E                       A:44 X:0 Y:44 P:67 SP:FB CYC:10473
DC1A  50 12     BVC $DC2E                       A:44 X:0 Y:44 P:67 SP:FB CYC:10475
DC1C  18        CLC                             A:44 X:0 Y:44 P:67 SP:FB CYC:10477
DC1D  B8        CLV                             A:44 X:0 Y:44 P:66 SP:FB CYC:10479
DC1E  A0 99     LDY #$99                        A:44 X:0 Y:44 P:26 SP:FB CYC:10481
DC20  A2 80     LDX #$80                        A:44 X:0 Y:99 P:A4 SP:FB CYC:10483
DC22  94 85     STY $85,X @ 05 = 00             A:44 X:80 Y:99 P:A4 SP:FB CYC:10485
DC24  A5 05     LDA $05 = 99                    A:44 X:80 Y:99 P:A4 SP:FB CYC:10489
DC26  B0 06     BCS $DC2E                       A:99 X:80 Y:99 P:A4 SP:FB CYC:10492
DC28  C9 99     CMP #$99                        A:99 X:80 Y:99 P:A4 SP:FB CYC:10494
DC2A  D0 02     BNE $DC2E                       A:99 X:80 Y:99 P:27 SP:FB CYC:10496
DC2C  50 04     BVC $DC32                       A:99 X:80 Y:99 P:27 SP:FB CYC:10498
DC32  A0 0B     LDY #$0B                        A:99 X:80 Y:99 P:27 SP:FB CYC:10501
DC34  A9 AA     LDA #$AA                        A:99 X:80 Y:B P:25 SP:FB CYC:10503
DC36  A2 78     LDX #$78                        A:AA X:80 Y:B P:A5 SP:FB CYC:10505
DC38  85 78     STA $78 = 00                    A:AA X:78 Y:B P:25 SP:FB CYC:10507
DC3A  20 B6 F7  JSR $F7B6                       A:AA X:78 Y:B P:25 SP:FB CYC:10510
F7B6  18        CLC                             A:AA X:78 Y:B P:25 SP:F9 CYC:10516
F7B7  A9 FF     LDA #$FF                        A:AA X:78 Y:B P:24 SP:F9 CYC:10518
F7B9  85 01     STA $01 = FF                    A:FF X:78 Y:B P:A4 SP:F9 CYC:10520
F7BB  24 01     BIT $01 = FF                    A:FF X:78 Y:B P:A4 SP:F9 CYC:10523
F7BD  A9 55     LDA #$55                        A:FF X:78 Y:B P:E4 SP:F9 CYC:10526
F7BF  60        RTS                             A:55 X:78 Y:B P:64 SP:F9 CYC:10528
DC3D  15 00     ORA $00,X @ 78 = AA             A:55 X:78 Y:B P:64 SP:FB CYC:10534
DC3F  20 C0 F7  JSR $F7C0                       A:FF X:78 Y:B P:E4 SP:FB CYC:10538
F7C0  B0 09     BCS $F7CB                       A:FF X:78 Y:B P:E4 SP:F9 CYC:10544
F7C2  10 07     BPL $F7CB                       A:FF X:78 Y:B P:E4 SP:F9 CYC:10546
F7C4  C9 FF     CMP #$FF                        A:FF X:78 Y:B P:E4 SP:F9 CYC:10548
F7C6  D0 03     BNE $F7CB                       A:FF X:78 Y:B P:67 SP:F9 CYC:10550
F7C8  50 01     BVC $F7CB                       A:FF X:78 Y:B P:67 SP:F9 CYC:10552
F7CA  60        RTS                             A:FF X:78 Y:B P:67 SP:F9 CYC:10554
DC42  C8        INY                             A:FF X:78 Y:B P:67 SP:FB CYC:10560
DC43  A9 00     LDA #$00                        A:FF X:78 Y:C P:65 SP:FB CYC:10562
DC45  85 78     STA $78 = AA                    A:0 X:78 Y:C P:67 SP:FB CYC:10564
DC47  20 CE F7  JSR $F7CE                       A:0 X:78 Y:C P:67 SP:FB CYC:10567
F7CE  38        SEC                             A:0 X:78 Y:C P:67 SP:F9 CYC:10573
F7CF  B8        CLV                             A:0 X:78 Y:C P:67 SP:F9 CYC:10575
F7D0  A9 00     LDA #$00                        A:0 X:78 Y:C P:27 SP:F9 CYC:10577
F7D2  60        RTS                             A:0 X:78 Y:C P:27 SP:F9 CYC:10579
DC4A  15 00     ORA $00,X @ 78 = 00             A:0 X:78 Y:C P:27 SP:FB CYC:10585
DC4C  20 D3 F7  JSR $F7D3                       A:0 X:78 Y:C P:27 SP:FB CYC:10589
F7D3  D0 07     BNE $F7DC                       A:0 X:78 Y:C P:27 SP:F9 CYC:10595
F7D5  70 05     BVS $F7DC                       A:0 X:78 Y:C P:27 SP:F9 CYC:10597
F7D7  90 03     BCC $F7DC                       A:0 X:78 Y:C P:27 SP:F9 CYC:10599
F7D9  30 01     BMI $F7DC                       A:0 X:78 Y:C P:27 SP:F9 CYC:10601
F7DB  60        RTS                             A:0 X:78 Y:C P:27 SP:F9 CYC:10603
DC4F  C8        INY                             A:0 X:78 Y:C P:27 SP:FB CYC:10609
DC50  A9 AA     LDA #$AA                        A:0 X:78 Y:D P:25 SP:FB CYC:10611
DC52  85 78     STA $78 = 00                    A:AA X:78 Y:D P:A5 SP:FB CYC:10613
DC54  20 DF F7  JSR $F7DF                       A:AA X:78 Y:D P:A5 SP:FB CYC:10616
F7DF  18        CLC                             A:AA X:78 Y:D P:A5 SP:F9 CYC:10622
F7E0  24 01     BIT $01 = FF                    A:AA X:78 Y:D P:A4 SP:F9 CYC:10624
F7E2  A9 55     LDA #$55                        A:AA X:78 Y:D P:E4 SP:F9 CYC:10627
F7E4  60        RTS                             A:55 X:78 Y:D P:64 SP:F9 CYC:10629
DC57  35 00     AND $00,X @ 78 = AA             A:55 X:78 Y:D P:64 SP:FB CYC:10635
DC59  20 E5 F7  JSR $F7E5                       A:0 X:78 Y:D P:66 SP:FB CYC:10639
F7E5  D0 07     BNE $F7EE                       A:0 X:78 Y:D P:66 SP:F9 CYC:10645
F7E7  50 05     BVC $F7EE                       A:0 X:78 Y:D P:66 SP:F9 CYC:10647
F7E9  B0 03     BCS $F7EE                       A:0 X:78 Y:D P:66 SP:F9 CYC:10649
F7EB  30 01     BMI $F7EE                       A:0 X:78 Y:D P:66 SP:F9 CYC:10651
F7ED  60        RTS                             A:0 X:78 Y:D P:66 SP:F9 CYC:10653
DC5C  C8        INY                             A:0 X:78 Y:D P:66 SP:FB CYC:10659
DC5D  A9 EF     LDA #$EF                        A:0 X:78 Y:E P:64 SP:FB CYC:10661
DC5F  85 78     STA $78 = AA                    A:EF X:78 Y:E P:E4 SP:FB CYC:10663
DC61  20 F1 F7  JSR $F7F1                       A:EF X:78 Y:E P:E4 SP:FB CYC:10666
F7F1  38        SEC                             A:EF X:78 Y:E P:E4 SP:F9 CYC:10672
F7F2  B8        CLV                             A:EF X:78 Y:E P:E5 SP:F9 CYC:10674
F7F3  A9 F8     LDA #$F8                        A:EF X:78 Y:E P:A5 SP:F9 CYC:10676
F7F5  60        RTS                             A:F8 X:78 Y:E P:A5 SP:F9 CYC:10678
DC64  35 00     AND $00,X @ 78 = EF             A:F8 X:78 Y:E P:A5 SP:FB CYC:10684
DC66  20 F6 F7  JSR $F7F6                       A:E8 X:78 Y:E P:A5 SP:FB CYC:10688
F7F6  90 09     BCC $F801                       A:E8 X:78 Y:E P:A5 SP:F9 CYC:10694
F7F8  10 07     BPL $F801                       A:E8 X:78 Y:E P:A5 SP:F9 CYC:10696
F7FA  C9 E8     CMP #$E8                        A:E8 X:78 Y:E P:A5 SP:F9 CYC:10698
F7FC  D0 03     BNE $F801                       A:E8 X:78 Y:E P:27 SP:F9 CYC:10700
F7FE  70 01     BVS $F801                       A:E8 X:78 Y:E P:27 SP:F9 CYC:10702
F800  60        RTS                             A:E8 X:78 Y:E P:27 SP:F9 CYC:10704
DC69  C8        INY                             A:E8 X:78 Y:E P:27 SP:FB CYC:10710
DC6A  A9 AA     LDA #$AA                        A:E8 X:78 Y:F P:25 SP:FB CYC:10712
DC6C  85 78     STA $78 = EF                    A:AA X:78 Y:F P:A5 SP:FB CYC:10714
DC6E  20 04 F8  JSR $F804                       A:AA X:78 Y:F P:A5 SP:FB CYC:10717
F804  18        CLC                             A:AA X:78 Y:F P:A5 SP:F9 CYC:10723
F805  24 01     BIT $01 = FF                    A:AA X:78 Y:F P:A4 SP:F9 CYC:10725
F807  A9 5F     LDA #$5F                        A:AA X:78 Y:F P:E4 SP:F9 CYC:10728
F809  60        RTS                             A:5F X:78 Y:F P:64 SP:F9 CYC:10730
DC71  55 00     EOR $00,X @ 78 = AA             A:5F X:78 Y:F P:64 SP:FB CYC:10736
DC73  20 0A F8  JSR $F80A                       A:F5 X:78 Y:F P:E4 SP:FB CYC:10740
F80A  B0 09     BCS $F815                       A:F5 X:78 Y:F P:E4 SP:F9 CYC:10746
F80C  10 07     BPL $F815                       A:F5 X:78 Y:F P:E4 SP:F9 CYC:10748
F80E  C9 F5     CMP #$F5                        A:F5 X:78 Y:F P:E4 SP:F9 CYC:10750
F810  D0 03     BNE $F815                       A:F5 X:78 Y:F P:67 SP:F9 CYC:10752
F812  50 01     BVC $F815                       A:F5 X:78 Y:F P:67 SP:F9 CYC:10754
F814  60        RTS                             A:F5 X:78 Y:F P:67 SP:F9 CYC:10756
DC76  C8        INY                             A:F5 X:78 Y:F P:67 SP:FB CYC:10762
DC77  A9 70     LDA #$70                        A:F5 X:78 Y:10 P:65 SP:FB CYC:10764
DC79  85 78     STA $78 = AA                    A:70 X:78 Y:10 P:65 SP:FB CYC:10766
DC7B  20 18 F8  JSR $F818                       A:70 X:78 Y:10 P:65 SP:FB CYC:10769
F818  38        SEC                             A:70 X:78 Y:10 P:65 SP:F9 CYC:10775
F819  B8        CLV                             A:70 X:78 Y:10 P:65 SP:F9 CYC:10777
F81A  A9 70     LDA #$70                        A:70 X:78 Y:10 P:25 SP:F9 CYC:10779
F81C  60        RTS                             A:70 X:78 Y:10 P:25 SP:F9 CYC:10781
DC7E  55 00     EOR $00,X @ 78 = 70             A:70 X:78 Y:10 P:25 SP:FB CYC:10787
DC80  20 1D F8  JSR $F81D                       A:0 X:78 Y:10 P:27 SP:FB CYC:10791
F81D  D0 07     BNE $F826                       A:0 X:78 Y:10 P:27 SP:F9 CYC:10797
F81F  70 05     BVS $F826                       A:0 X:78 Y:10 P:27 SP:F9 CYC:10799
F821  90 03     BCC $F826                       A:0 X:78 Y:10 P:27 SP:F9 CYC:10801
F823  30 01     BMI $F826                       A:0 X:78 Y:10 P:27 SP:F9 CYC:10803
F825  60        RTS                             A:0 X:78 Y:10 P:27 SP:F9 CYC:10805
DC83  C8        INY                             A:0 X:78 Y:10 P:27 SP:FB CYC:10811
DC84  A9 69     LDA #$69                        A:0 X:78 Y:11 P:25 SP:FB CYC:10813
DC86  85 78     STA $78 = 70                    A:69 X:78 Y:11 P:25 SP:FB CYC:10815
DC88  20 29 F8  JSR $F829                       A:69 X:78 Y:11 P:25 SP:FB CYC:10818
F829  18        CLC                             A:69 X:78 Y:11 P:25 SP:F9 CYC:10824
F82A  24 01     BIT $01 = FF                    A:69 X:78 Y:11 P:24 SP:F9 CYC:10826
F82C  A9 00     LDA #$00                        A:69 X:78 Y:11 P:E4 SP:F9 CYC:10829
F82E  60        RTS                             A:0 X:78 Y:11 P:66 SP:F9 CYC:10831
DC8B  75 00     ADC $00,X @ 78 = 69             A:0 X:78 Y:11 P:66 SP:FB CYC:10837
DC8D  20 2F F8  JSR $F82F                       A:69 X:78 Y:11 P:24 SP:FB CYC:10841
F82F  30 09     BMI $F83A                       A:69 X:78 Y:11 P:24 SP:F9 CYC:10847
F831  B0 07     BCS $F83A                       A:69 X:78 Y:11 P:24 SP:F9 CYC:10849
F833  C9 69     CMP #$69                        A:69 X:78 Y:11 P:24 SP:F9 CYC:10851
F835  D0 03     BNE $F83A                       A:69 X:78 Y:11 P:27 SP:F9 CYC:10853
F837  70 01     BVS $F83A                       A:69 X:78 Y:11 P:27 SP:F9 CYC:10855
F839  60        RTS                             A:69 X:78 Y:11 P:27 SP:F9 CYC:10857
DC90  C8        INY                             A:69 X:78 Y:11 P:27 SP:FB CYC:10863
DC91  20 3D F8  JSR $F83D                       A:69 X:78 Y:12 P:25 SP:FB CYC:10865
F83D  38        SEC                             A:69 X:78 Y:12 P:25 SP:F9 CYC:10871
F83E  24 01     BIT $01 = FF                    A:69 X:78 Y:12 P:25 SP:F9 CYC:10873
F840  A9 00     LDA #$00                        A:69 X:78 Y:12 P:E5 SP:F9 CYC:10876
F842  60        RTS                             A:0 X:78 Y:12 P:67 SP:F9 CYC:10878
DC94  75 00     ADC $00,X @ 78 = 69             A:0 X:78 Y:12 P:67 SP:FB CYC:10884
DC96  20 43 F8  JSR $F843                       A:6A X:78 Y:12 P:24 SP:FB CYC:10888
F843  30 09     BMI $F84E                       A:6A X:78 Y:12 P:24 SP:F9 CYC:10894
F845  B0 07     BCS $F84E                       A:6A X:78 Y:12 P:24 SP:F9 CYC:10896
F847  C9 6A     CMP #$6A                        A:6A X:78 Y:12 P:24 SP:F9 CYC:10898
F849  D0 03     BNE $F84E                       A:6A X:78 Y:12 P:27 SP:F9 CYC:10900
F84B  70 01     BVS $F84E                       A:6A X:78 Y:12 P:27 SP:F9 CYC:10902
F84D  60        RTS                             A:6A X:78 Y:12 P:27 SP:F9 CYC:10904
DC99  C8        INY                             A:6A X:78 Y:12 P:27 SP:FB CYC:10910
DC9A  A9 7F     LDA #$7F                        A:6A X:78 Y:13 P:25 SP:FB CYC:10912
DC9C  85 78     STA $78 = 69                    A:7F X:78 Y:13 P:25 SP:FB CYC:10914
DC9E  20 51 F8  JSR $F851                       A:7F X:78 Y:13 P:25 SP:FB CYC:10917
F851  38        SEC                             A:7F X:78 Y:13 P:25 SP:F9 CYC:10923
F852  B8        CLV                             A:7F X:78 Y:13 P:25 SP:F9 CYC:10925
F853  A9 7F     LDA #$7F                        A:7F X:78 Y:13 P:25 SP:F9 CYC:10927
F855  60        RTS                             A:7F X:78 Y:13 P:25 SP:F9 CYC:10929
DCA1  75 00     ADC $00,X @ 78 = 7F             A:7F X:78 Y:13 P:25 SP:FB CYC:10935
DCA3  20 56 F8  JSR $F856                       A:FF X:78 Y:13 P:E4 SP:FB CYC:10939
F856  10 09     BPL $F861                       A:FF X:78 Y:13 P:E4 SP:F9 CYC:10945
F858  B0 07     BCS $F861                       A:FF X:78 Y:13 P:E4 SP:F9 CYC:10947
F85A  C9 FF     CMP #$FF                        A:FF X:78 Y:13 P:E4 SP:F9 CYC:10949
F85C  D0 03     BNE $F861                       A:FF X:78 Y:13 P:67 SP:F9 CYC:10951
F85E  50 01     BVC $F861                       A:FF X:78 Y:13 P:67 SP:F9 CYC:10953
F860  60        RTS                             A:FF X:78 Y:13 P:67 SP:F9 CYC:10955
DCA6  C8        INY                             A:FF X:78 Y:13 P:67 SP:FB CYC:10961
DCA7  A9 80     LDA #$80                        A:FF X:78 Y:14 P:65 SP:FB CYC:10963
DCA9  85 78     STA $78 = 7F                    A:80 X:78 Y:14 P:E5 SP:FB CYC:10965
DCAB  20 64 F8  JSR $F864                       A:80 X:78 Y:14 P:E5 SP:FB CYC:10968
F864  18        CLC                             A:80 X:78 Y:14 P:E5 SP:F9 CYC:10974
F865  24 01     BIT $01 = FF                    A:80 X:78 Y:14 P:E4 SP:F9 CYC:10976
F867  A9 7F     LDA #$7F                        A:80 X:78 Y:14 P:E4 SP:F9 CYC:10979
F869  60        RTS                             A:7F X:78 Y:14 P:64 SP:F9 CYC:10981
DCAE  75 00     ADC $00,X @ 78 = 80             A:7F X:78 Y:14 P:64 SP:FB CYC:10987
DCB0  20 6A F8  JSR $F86A                       A:FF X:78 Y:14 P:A4 SP:FB CYC:10991
F86A  10 09     BPL $F875                       A:FF X:78 Y:14 P:A4 SP:F9 CYC:10997
F86C  B0 07     BCS $F875                       A:FF X:78 Y:14 P:A4 SP:F9 CYC:10999
F86E  C9 FF     CMP #$FF                        A:FF X:78 Y:14 P:A4 SP:F9 CYC:11001
F870  D0 03     BNE $F875                       A:FF X:78 Y:14 P:27 SP:F9 CYC:11003
F872  70 01     BVS $F875                       A:FF X:78 Y:14 P:27 SP:F9 CYC:11005
F874  60        RTS                             A:FF X:78 Y:14 P:27 SP:F9 CYC:11007
DCB3  C8        INY                             A:FF X:78 Y:14 P:27 SP:FB CYC:11013
DCB4  20 78 F8  JSR $F878                       A:FF X:78 Y:15 P:25 SP:FB CYC:11015
F878  38        SEC                             A:FF X:78 Y:15 P:25 SP:F9 CYC:11021
F879  B8        CLV                             A:FF X:78 Y:15 P:25 SP:F9 CYC:11023
F87A  A9 7F     LDA #$7F                        A:FF X:78 Y:15 P:25 SP:F9 CYC:11025
F87C  60        RTS                             A:7F X:78 Y:15 P:25 SP:F9 CYC:11027
DCB7  75 00     ADC $00,X @ 78 = 80             A:7F X:78 Y:15 P:25 SP:FB CYC:11033
DCB9  20 7D F8  JSR $F87D                       A:0 X:78 Y:15 P:27 SP:FB CYC:11037
F87D  D0 07     BNE $F886                       A:0 X:78 Y:15 P:27 SP:F9 CYC:11043
F87F  30 05     BMI $F886                       A:0 X:78 Y:15 P:27 SP:F9 CYC:11045
F881  70 03     BVS $F886                       A:0 X:78 Y:15 P:27 SP:F9 CYC:11047
F883  90 01     BCC $F886                       A:0 X:78 Y:15 P:27 SP:F9 CYC:11049
F885  60        RTS                             A:0 X:78 Y:15 P:27 SP:F9 CYC:11051
DCBC  C8        INY                             A:0 X:78 Y:15 P:27 SP:FB CYC:11057
DCBD  A9 40     LDA #$40                        A:0 X:78 Y:16 P:25 SP:FB CYC:11059
DCBF  85 78     STA $78 = 80                    A:40 X:78 Y:16 P:25 SP:FB CYC:11061
DCC1  20 89 F8  JSR $F889                       A:40 X:78 Y:16 P:25 SP:FB CYC:11064
F889  24 01     BIT $01 = FF                    A:40 X:78 Y:16 P:25 SP:F9 CYC:11070
F88B  A9 40     LDA #$40                        A:40 X:78 Y:16 P:E5 SP:F9 CYC:11073
F88D  60        RTS                             A:40 X:78 Y:16 P:65 SP:F9 CYC:11075
DCC4  D5 00     CMP $00,X @ 78 = 40             A:40 X:78 Y:16 P:65 SP:FB CYC:11081
DCC6  20 8E F8  JSR $F88E                       A:40 X:78 Y:16 P:67 SP:FB CYC:11085
F88E  30 07     BMI $F897                       A:40 X:78 Y:16 P:67 SP:F9 CYC:11091
F890  90 05     BCC $F897                       A:40 X:78 Y:16 P:67 SP:F9 CYC:11093
F892  D0 03     BNE $F897                       A:40 X:78 Y:16 P:67 SP:F9 CYC:11095
F894  50 01     BVC $F897                       A:40 X:78 Y:16 P:67 SP:F9 CYC:11097
F896  60        RTS                             A:40 X:78 Y:16 P:67 SP:F9 CYC:11099
DCC9  C8        INY                             A:40 X:78 Y:16 P:67 SP:FB CYC:11105
DCCA  48        PHA                             A:40 X:78 Y:17 P:65 SP:FB CYC:11107
DCCB  A9 3F     LDA #$3F                        A:40 X:78 Y:17 P:65 SP:FA CYC:11110
DCCD  85 78     STA $78 = 40                    A:3F X:78 Y:17 P:65 SP:FA CYC:11112
DCCF  68        PLA                             A:3F X:78 Y:17 P:65 SP:FA CYC:11115
DCD0  20 9A F8  JSR $F89A                       A:40 X:78 Y:17 P:65 SP:FB CYC:11119
F89A  B8        CLV                             A:40 X:78 Y:17 P:65 SP:F9 CYC:11125
F89B  60        RTS                             A:40 X:78 Y:17 P:25 SP:F9 CYC:11127
DCD3  D5 00     CMP $00,X @ 78 = 3F             A:40 X:78 Y:17 P:25 SP:FB CYC:11133
DCD5  20 9C F8  JSR $F89C                       A:40 X:78 Y:17 P:25 SP:FB CYC:11137
F89C  F0 07     BEQ $F8A5                       A:40 X:78 Y:17 P:25 SP:F9 CYC:11143
F89E  30 05     BMI $F8A5                       A:40 X:78 Y:17 P:25 SP:F9 CYC:11145
F8A0  90 03     BCC $F8A5                       A:40 X:78 Y:17 P:25 SP:F9 CYC:11147
F8A2  70 01     BVS $F8A5                       A:40 X:78 Y:17 P:25 SP:F9 CYC:11149
F8A4  60        RTS                             A:40 X:78 Y:17 P:25 SP:F9 CYC:11151
DCD8  C8        INY                             A:40 X:78 Y:17 P:25 SP:FB CYC:11157
DCD9  48        PHA                             A:40 X:78 Y:18 P:25 SP:FB CYC:11159
DCDA  A9 41     LDA #$41                        A:40 X:78 Y:18 P:25 SP:FA CYC:11162
DCDC  85 78     STA $78 = 3F                    A:41 X:78 Y:18 P:25 SP:FA CYC:11164
DCDE  68        PLA                             A:41 X:78 Y:18 P:25 SP:FA CYC:11167
DCDF  D5 00     CMP $00,X @ 78 = 41             A:40 X:78 Y:18 P:25 SP:FB CYC:11171
DCE1  20 A8 F8  JSR $F8A8                       A:40 X:78 Y:18 P:A4 SP:FB CYC:11175
F8A8  F0 05     BEQ $F8AF                       A:40 X:78 Y:18 P:A4 SP:F9 CYC:11181
F8AA  10 03     BPL $F8AF                       A:40 X:78 Y:18 P:A4 SP:F9 CYC:11183
F8AC  10 01     BPL $F8AF                       A:40 X:78 Y:18 P:A4 SP:F9 CYC:11185
F8AE  60        RTS                             A:40 X:78 Y:18 P:A4 SP:F9 CYC:11187
DCE4  C8        INY                             A:40 X:78 Y:18 P:A4 SP:FB CYC:11193
DCE5  48        PHA                             A:40 X:78 Y:19 P:24 SP:FB CYC:11195
DCE6  A9 00     LDA #$00                        A:40 X:78 Y:19 P:24 SP:FA CYC:11198
DCE8  85 78     STA $78 = 41                    A:0 X:78 Y:19 P:26 SP:FA CYC:11200
DCEA  68        PLA                             A:0 X:78 Y:19 P:26 SP:FA CYC:11203
DCEB  20 B2 F8  JSR $F8B2                       A:40 X:78 Y:19 P:24 SP:FB CYC:11207
F8B2  A9 80     LDA #$80                        A:40 X:78 Y:19 P:24 SP:F9 CYC:11213
F8B4  60        RTS                             A:80 X:78 Y:19 P:A4 SP:F9 CYC:11215
DCEE  D5 00     CMP $00,X @ 78 = 00             A:80 X:78 Y:19 P:A4 SP:FB CYC:11221
DCF0  20 B5 F8  JSR $F8B5                       A:80 X:78 Y:19 P:A5 SP:FB CYC:11225
F8B5  F0 05     BEQ $F8BC                       A:80 X:78 Y:19 P:A5 SP:F9 CYC:11231
F8B7  10 03     BPL $F8BC                       A:80 X:78 Y:19 P:A5 SP:F9 CYC:11233
F8B9  90 01     BCC $F8BC                       A:80 X:78 Y:19 P:A5 SP:F9 CYC:11235
F8BB  60        RTS                             A:80 X:78 Y:19 P:A5 SP:F9 CYC:11237
DCF3  C8        INY                             A:80 X:78 Y:19 P:A5 SP:FB CYC:11243
DCF4  48        PHA                             A:80 X:78 Y:1A P:25 SP:FB CYC:11245
DCF5  A9 80     LDA #$80                        A:80 X:78 Y:1A P:25 SP:FA CYC:11248
DCF7  85 78     STA $78 = 00                    A:80 X:78 Y:1A P:A5 SP:FA CYC:11250
DCF9  68        PLA                             A:80 X:78 Y:1A P:A5 SP:FA CYC:11253
DCFA  D5 00     CMP $00,X @ 78 = 80             A:80 X:78 Y:1A P:A5 SP:FB CYC:11257
DCFC  20 BF F8  JSR $F8BF                       A:80 X:78 Y:1A P:27 SP:FB CYC:11261
F8BF  D0 05     BNE $F8C6                       A:80 X:78 Y:1A P:27 SP:F9 CYC:11267
F8C1  30 03     BMI $F8C6                       A:80 X:78 Y:1A P:27 SP:F9 CYC:11269
F8C3  90 01     BCC $F8C6                       A:80 X:78 Y:1A P:27 SP:F9 CYC:11271
F8C5  60        RTS                             A:80 X:78 Y:1A P:27 SP:F9 CYC:11273
DCFF  C8        INY                             A:80 X:78 Y:1A P:27 SP:FB CYC:11279
DD00  48        PHA                             A:80 X:78 Y:1B P:25 SP:FB CYC:11281
DD01  A9 81     LDA #$81                        A:80 X:78 Y:1B P:25 SP:FA CYC:11284
DD03  85 78     STA $78 = 80                    A:81 X:78 Y:1B P:A5 SP:FA CYC:11286
DD05  68        PLA                             A:81 X:78 Y:1B P:A5 SP:FA CYC:11289
DD06  D5 00     CMP $00,X @ 78 = 81             A:80 X:78 Y:1B P:A5 SP:FB CYC:11293
DD08  20 C9 F8  JSR $F8C9                       A:80 X:78 Y:1B P:A4 SP:FB CYC:11297
F8C9  B0 05     BCS $F8D0                       A:80 X:78 Y:1B P:A4 SP:F9 CYC:11303
F8CB  F0 03     BEQ $F8D0                       A:80 X:78 Y:1B P:A4 SP:F9 CYC:11305
F8CD  10 01     BPL $F8D0                       A:80 X:78 Y:1B P:A4 SP:F9 CYC:11307
F8CF  60        RTS                             A:80 X:78 Y:1B P:A4 SP:F9 CYC:11309
DD0B  C8        INY                             A:80 X:78 Y:1B P:A4 SP:FB CYC:11315
DD0C  48        PHA                             A:80 X:78 Y:1C P:24 SP:FB CYC:11317
DD0D  A9 7F     LDA #$7F                        A:80 X:78 Y:1C P:24 SP:FA CYC:11320
DD0F  85 78     STA $78 = 81                    A:7F X:78 Y:1C P:24 SP:FA CYC:11322
DD11  68        PLA                             A:7F X:78 Y:1C P:24 SP:FA CYC:11325
DD12  D5 00     CMP $00,X @ 78 = 7F             A:80 X:78 Y:1C P:A4 SP:FB CYC:11329
DD14  20 D3 F8  JSR $F8D3                       A:80 X:78 Y:1C P:25 SP:FB CYC:11333
F8D3  90 05     BCC $F8DA                       A:80 X:78 Y:1C P:25 SP:F9 CYC:11339
F8D5  F0 03     BEQ $F8DA                       A:80 X:78 Y:1C P:25 SP:F9 CYC:11341
F8D7  30 01     BMI $F8DA                       A:80 X:78 Y:1C P:25 SP:F9 CYC:11343
F8D9  60        RTS                             A:80 X:78 Y:1C P:25 SP:F9 CYC:11345
DD17  C8        INY                             A:80 X:78 Y:1C P:25 SP:FB CYC:11351
DD18  A9 40     LDA #$40                        A:80 X:78 Y:1D P:25 SP:FB CYC:11353
DD1A  85 78     STA $78 = 7F                    A:40 X:78 Y:1D P:25 SP:FB CYC:11355
DD1C  20 31 F9  JSR $F931                       A:40 X:78 Y:1D P:25 SP:FB CYC:11358
F931  24 01     BIT $01 = FF                    A:40 X:78 Y:1D P:25 SP:F9 CYC:11364
F933  A9 40     LDA #$40                        A:40 X:78 Y:1D P:E5 SP:F9 CYC:11367
F935  38        SEC                             A:40 X:78 Y:1D P:65 SP:F9 CYC:11369
F936  60        RTS                             A:40 X:78 Y:1D P:65 SP:F9 CYC:11371
DD1F  F5 00     SBC $00,X @ 78 = 40             A:40 X:78 Y:1D P:65 SP:FB CYC:11377
DD21  20 37 F9  JSR $F937                       A:0 X:78 Y:1D P:27 SP:FB CYC:11381
F937  30 0B     BMI $F944                       A:0 X:78 Y:1D P:27 SP:F9 CYC:11387
F939  90 09     BCC $F944                       A:0 X:78 Y:1D P:27 SP:F9 CYC:11389
F93B  D0 07     BNE $F944                       A:0 X:78 Y:1D P:27 SP:F9 CYC:11391
F93D  70 05     BVS $F944                       A:0 X:78 Y:1D P:27 SP:F9 CYC:11393
F93F  C9 00     CMP #$00                        A:0 X:78 Y:1D P:27 SP:F9 CYC:11395
F941  D0 01     BNE $F944                       A:0 X:78 Y:1D P:27 SP:F9 CYC:11397
F943  60        RTS                             A:0 X:78 Y:1D P:27 SP:F9 CYC:11399
DD24  C8        INY                             A:0 X:78 Y:1D P:27 SP:FB CYC:11405
DD25  A9 3F     LDA #$3F                        A:0 X:78 Y:1E P:25 SP:FB CYC:11407
DD27  85 78     STA $78 = 40                    A:3F X:78 Y:1E P:25 SP:FB CYC:11409
DD29  20 47 F9  JSR $F947                       A:3F X:78 Y:1E P:25 SP:FB CYC:11412
F947  B8        CLV                             A:3F X:78 Y:1E P:25 SP:F9 CYC:11418
F948  38        SEC                             A:3F X:78 Y:1E P:25 SP:F9 CYC:11420
F949  A9 40     LDA #$40                        A:3F X:78 Y:1E P:25 SP:F9 CYC:11422
F94B  60        RTS                             A:40 X:78 Y:1E P:25 SP:F9 CYC:11424
DD2C  F5 00     SBC $00,X @ 78 = 3F             A:40 X:78 Y:1E P:25 SP:FB CYC:11430
DD2E  20 4C F9  JSR $F94C                       A:1 X:78 Y:1E P:25 SP:FB CYC:11434
F94C  F0 0B     BEQ $F959                       A:1 X:78 Y:1E P:25 SP:F9 CYC:11440
F94E  30 09     BMI $F959                       A:1 X:78 Y:1E P:25 SP:F9 CYC:11442
F950  90 07     BCC $F959                       A:1 X:78 Y:1E P:25 SP:F9 CYC:11444
F952  70 05     BVS $F959                       A:1 X:78 Y:1E P:25 SP:F9 CYC:11446
F954  C9 01     CMP #$01                        A:1 X:78 Y:1E P:25 SP:F9 CYC:11448
F956  D0 01     BNE $F959                       A:1 X:78 Y:1E P:27 SP:F9 CYC:11450
F958  60        RTS                             A:1 X:78 Y:1E P:27 SP:F9 CYC:11452
DD31  C8        INY                             A:1 X:78 Y:1E P:27 SP:FB CYC:11458
DD32  A9 41     LDA #$41                        A:1 X:78 Y:1F P:25 SP:FB CYC:11460
DD34  85 78     STA $78 = 3F                    A:41 X:78 Y:1F P:25 SP:FB CYC:11462
DD36  20 5C F9  JSR $F95C                       A:41 X:78 Y:1F P:25 SP:FB CYC:11465
F95C  A9 40     LDA #$40                        A:41 X:78 Y:1F P:25 SP:F9 CYC:11471
F95E  38        SEC                             A:40 X:78 Y:1F P:25 SP:F9 CYC:11473
F95F  24 01     BIT $01 = FF                    A:40 X:78 Y:1F P:25 SP:F9 CYC:11475
F961  60        RTS                             A:40 X:78 Y:1F P:E5 SP:F9 CYC:11478
DD39  F5 00     SBC $00,X @ 78 = 41             A:40 X:78 Y:1F P:E5 SP:FB CYC:11484
DD3B  20 62 F9  JSR $F962                       A:FF X:78 Y:1F P:A4 SP:FB CYC:11488
F962  B0 0B     BCS $F96F                       A:FF X:78 Y:1F P:A4 SP:F9 CYC:11494
F964  F0 09     BEQ $F96F                       A:FF X:78 Y:1F P:A4 SP:F9 CYC:11496
F966  10 07     BPL $F96F                       A:FF X:78 Y:1F P:A4 SP:F9 CYC:11498
F968  70 05     BVS $F96F                       A:FF X:78 Y:1F P:A4 SP:F9 CYC:11500
F96A  C9 FF     CMP #$FF                        A:FF X:78 Y:1F P:A4 SP:F9 CYC:11502
F96C  D0 01     BNE $F96F                       A:FF X:78 Y:1F P:27 SP:F9 CYC:11504
F96E  60        RTS                             A:FF X:78 Y:1F P:27 SP:F9 CYC:11506
DD3E  C8        INY                             A:FF X:78 Y:1F P:27 SP:FB CYC:11512
DD3F  A9 00     LDA #$00                        A:FF X:78 Y:20 P:25 SP:FB CYC:11514
DD41  85 78     STA $78 = 41                    A:0 X:78 Y:20 P:27 SP:FB CYC:11516
DD43  20 72 F9  JSR $F972                       A:0 X:78 Y:20 P:27 SP:FB CYC:11519
F972  18        CLC                             A:0 X:78 Y:20 P:27 SP:F9 CYC:11525
F973  A9 80     LDA #$80                        A:0 X:78 Y:20 P:26 SP:F9 CYC:11527
F975  60        RTS                             A:80 X:78 Y:20 P:A4 SP:F9 CYC:11529
DD46  F5 00     SBC $00,X @ 78 = 00             A:80 X:78 Y:20 P:A4 SP:FB CYC:11535
DD48  20 76 F9  JSR $F976                       A:7F X:78 Y:20 P:65 SP:FB CYC:11539
F976  90 05     BCC $F97D                       A:7F X:78 Y:20 P:65 SP:F9 CYC:11545
F978  C9 7F     CMP #$7F                        A:7F X:78 Y:20 P:65 SP:F9 CYC:11547
F97A  D0 01     BNE $F97D                       A:7F X:78 Y:20 P:67 SP:F9 CYC:11549
F97C  60        RTS                             A:7F X:78 Y:20 P:67 SP:F9 CYC:11551
DD4B  C8        INY                             A:7F X:78 Y:20 P:67 SP:FB CYC:11557
DD4C  A9 7F     LDA #$7F                        A:7F X:78 Y:21 P:65 SP:FB CYC:11559
DD4E  85 78     STA $78 = 00                    A:7F X:78 Y:21 P:65 SP:FB CYC:11561
DD50  20 80 F9  JSR $F980                       A:7F X:78 Y:21 P:65 SP:FB CYC:11564
F980  38        SEC                             A:7F X:78 Y:21 P:65 SP:F9 CYC:11570
F981  A9 81     LDA #$81                        A:7F X:78 Y:21 P:65 SP:F9 CYC:11572
F983  60        RTS                             A:81 X:78 Y:21 P:E5 SP:F9 CYC:11574
DD53  F5 00     SBC $00,X @ 78 = 7F             A:81 X:78 Y:21 P:E5 SP:FB CYC:11580
DD55  20 84 F9  JSR $F984                       A:2 X:78 Y:21 P:65 SP:FB CYC:11584
F984  50 07     BVC $F98D                       A:2 X:78 Y:21 P:65 SP:F9 CYC:11590
F986  90 05     BCC $F98D                       A:2 X:78 Y:21 P:65 SP:F9 CYC:11592
F988  C9 02     CMP #$02                        A:2 X:78 Y:21 P:65 SP:F9 CYC:11594
F98A  D0 01     BNE $F98D                       A:2 X:78 Y:21 P:67 SP:F9 CYC:11596
F98C  60        RTS                             A:2 X:78 Y:21 P:67 SP:F9 CYC:11598
DD58  A9 AA     LDA #$AA                        A:2 X:78 Y:21 P:67 SP:FB CYC:11604
DD5A  85 33     STA $33 = 44                    A:AA X:78 Y:21 P:E5 SP:FB CYC:11606
DD5C  A9 BB     LDA #$BB                        A:AA X:78 Y:21 P:E5 SP:FB CYC:11609
DD5E  85 89     STA $89 = BB                    A:BB X:78 Y:21 P:E5 SP:FB CYC:11611
DD60  A2 00     LDX #$00                        A:BB X:78 Y:21 P:E5 SP:FB CYC:11614
DD62  A0 66     LDY #$66                        A:BB X:0 Y:21 P:67 SP:FB CYC:11616
DD64  24 01     BIT $01 = FF                    A:BB X:0 Y:66 P:65 SP:FB CYC:11618
DD66  38        SEC                             A:BB X:0 Y:66 P:E5 SP:FB CYC:11621
DD67  A9 00     LDA #$00                        A:BB X:0 Y:66 P:E5 SP:FB CYC:11623
DD69  B5 33     LDA $33,X @ 33 = AA             A:0 X:0 Y:66 P:67 SP:FB CYC:11625
DD6B  10 12     BPL $DD7F                       A:AA X:0 Y:66 P:E5 SP:FB CYC:11629
DD6D  F0 10     BEQ $DD7F                       A:AA X:0 Y:66 P:E5 SP:FB CYC:11631
DD6F  50 0E     BVC $DD7F                       A:AA X:0 Y:66 P:E5 SP:FB CYC:11633
DD71  90 0C     BCC $DD7F                       A:AA X:0 Y:66 P:E5 SP:FB CYC:11635
DD73  C0 66     CPY #$66                        A:AA X:0 Y:66 P:E5 SP:FB CYC:11637
DD75  D0 08     BNE $DD7F                       A:AA X:0 Y:66 P:67 SP:FB CYC:11639
DD77  E0 00     CPX #$00                        A:AA X:0 Y:66 P:67 SP:FB CYC:11641
DD79  D0 04     BNE $DD7F                       A:AA X:0 Y:66 P:67 SP:FB CYC:11643
DD7B  C9 AA     CMP #$AA                        A:AA X:0 Y:66 P:67 SP:FB CYC:11645
DD7D  F0 04     BEQ $DD83                       A:AA X:0 Y:66 P:67 SP:FB CYC:11647
DD83  A2 8A     LDX #$8A                        A:AA X:0 Y:66 P:67 SP:FB CYC:11650
DD85  A0 66     LDY #$66                        A:AA X:8A Y:66 P:E5 SP:FB CYC:11652
DD87  B8        CLV                             A:AA X:8A Y:66 P:65 SP:FB CYC:11654
DD88  18        CLC                             A:AA X:8A Y:66 P:25 SP:FB CYC:11656
DD89  A9 00     LDA #$00                        A:AA X:8A Y:66 P:24 SP:FB CYC:11658
DD8B  B5 FF     LDA $FF,X @ 89 = BB             A:0 X:8A Y:66 P:26 SP:FB CYC:11660
DD8D  10 12     BPL $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB CYC:11664
DD8F  F0 10     BEQ $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB CYC:11666
DD91  70 0E     BVS $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB CYC:11668
DD93  B0 0C     BCS $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB CYC:11670
DD95  C9 BB     CMP #$BB                        A:BB X:8A Y:66 P:A4 SP:FB CYC:11672
DD97  D0 08     BNE $DDA1                       A:BB X:8A Y:66 P:27 SP:FB CYC:11674
DD99  C0 66     CPY #$66                        A:BB X:8A Y:66 P:27 SP:FB CYC:11676
DD9B  D0 04     BNE $DDA1                       A:BB X:8A Y:66 P:27 SP:FB CYC:11678
DD9D  E0 8A     CPX #$8A                        A:BB X:8A Y:66 P:27 SP:FB CYC:11680
DD9F  F0 04     BEQ $DDA5                       A:BB X:8A Y:66 P:27 SP:FB CYC:11682
DDA5  24 01     BIT $01 = FF                    A:BB X:8A Y:66 P:27 SP:FB CYC:11685
DDA7  38        SEC                             A:BB X:8A Y:66 P:E5 SP:FB CYC:11688
DDA8  A9 44     LDA #$44                        A:BB X:8A Y:66 P:E5 SP:FB CYC:11690
DDAA  A2 00     LDX #$00                        A:44 X:8A Y:66 P:65 SP:FB CYC:11692
DDAC  95 33     STA $33,X @ 33 = AA             A:44 X:0 Y:66 P:67 SP:FB CYC:11694
DDAE  A5 33     LDA $33 = 44                    A:44 X:0 Y:66 P:67 SP:FB CYC:11698
DDB0  90 18     BCC $DDCA                       A:44 X:0 Y:66 P:65 SP:FB CYC:11701
DDB2  C9 44     CMP #$44                        A:44 X:0 Y:66 P:65 SP:FB CYC:11703
DDB4  D0 14     BNE $DDCA                       A:44 X:0 Y:66 P:67 SP:FB CYC:11705
DDB6  50 12     BVC $DDCA                       A:44 X:0 Y:66 P:67 SP:FB CYC:11707
DDB8  18        CLC                             A:44 X:0 Y:66 P:67 SP:FB CYC:11709
DDB9  B8        CLV                             A:44 X:0 Y:66 P:66 SP:FB CYC:11711
DDBA  A9 99     LDA #$99                        A:44 X:0 Y:66 P:26 SP:FB CYC:11713
DDBC  A2 80     LDX #$80                        A:99 X:0 Y:66 P:A4 SP:FB CYC:11715
DDBE  95 85     STA $85,X @ 05 = 99             A:99 X:80 Y:66 P:A4 SP:FB CYC:11717
DDC0  A5 05     LDA $05 = 99                    A:99 X:80 Y:66 P:A4 SP:FB CYC:11721
DDC2  B0 06     BCS $DDCA                       A:99 X:80 Y:66 P:A4 SP:FB CYC:11724
DDC4  C9 99     CMP #$99                        A:99 X:80 Y:66 P:A4 SP:FB CYC:11726
DDC6  D0 02     BNE $DDCA                       A:99 X:80 Y:66 P:27 SP:FB CYC:11728
DDC8  50 04     BVC $DDCE                       A:99 X:80 Y:66 P:27 SP:FB CYC:11730
DDCE  A0 25     LDY #$25                        A:99 X:80 Y:66 P:27 SP:FB CYC:11733
DDD0  A2 78     LDX #$78                        A:99 X:80 Y:25 P:25 SP:FB CYC:11735
DDD2  20 90 F9  JSR $F990                       A:99 X:78 Y:25 P:25 SP:FB CYC:11737
F990  A2 55     LDX #$55                        A:99 X:78 Y:25 P:25 SP:F9 CYC:11743
F992  A9 FF     LDA #$FF                        A:99 X:55 Y:25 P:25 SP:F9 CYC:11745
F994  85 01     STA $01 = FF                    A:FF X:55 Y:25 P:A5 SP:F9 CYC:11747
F996  EA        NOP                             A:FF X:55 Y:25 P:A5 SP:F9 CYC:11750
F997  24 01     BIT $01 = FF                    A:FF X:55 Y:25 P:A5 SP:F9 CYC:11752
F999  38        SEC                             A:FF X:55 Y:25 P:E5 SP:F9 CYC:11755
F99A  A9 01     LDA #$01                        A:FF X:55 Y:25 P:E5 SP:F9 CYC:11757
F99C  60        RTS                             A:1 X:55 Y:25 P:65 SP:F9 CYC:11759
DDD5  95 00     STA $00,X @ 55 = 00             A:1 X:55 Y:25 P:65 SP:FB CYC:11765
DDD7  56 00     LSR $00,X @ 55 = 01             A:1 X:55 Y:25 P:65 SP:FB CYC:11769
DDD9  B5 00     LDA $00,X @ 55 = 00             A:1 X:55 Y:25 P:67 SP:FB CYC:11775
DDDB  20 9D F9  JSR $F99D                       A:0 X:55 Y:25 P:67 SP:FB CYC:11779
F99D  90 1B     BCC $F9BA                       A:0 X:55 Y:25 P:67 SP:F9 CYC:11785
F99F  D0 19     BNE $F9BA                       A:0 X:55 Y:25 P:67 SP:F9 CYC:11787
F9A1  30 17     BMI $F9BA                       A:0 X:55 Y:25 P:67 SP:F9 CYC:11789
F9A3  50 15     BVC $F9BA                       A:0 X:55 Y:25 P:67 SP:F9 CYC:11791
F9A5  C9 00     CMP #$00                        A:0 X:55 Y:25 P:67 SP:F9 CYC:11793
F9A7  D0 11     BNE $F9BA                       A:0 X:55 Y:25 P:67 SP:F9 CYC:11795
F9A9  B8        CLV                             A:0 X:55 Y:25 P:67 SP:F9 CYC:11797
F9AA  A9 AA     LDA #$AA                        A:0 X:55 Y:25 P:27 SP:F9 CYC:11799
F9AC  60        RTS                             A:AA X:55 Y:25 P:A5 SP:F9 CYC:11801
DDDE  C8        INY                             A:AA X:55 Y:25 P:A5 SP:FB CYC:11807
DDDF  95 00     STA $00,X @ 55 = 00             A:AA X:55 Y:26 P:25 SP:FB CYC:11809
DDE1  56 00     LSR $00,X @ 55 = AA             A:AA X:55 Y:26 P:25 SP:FB CYC:11813
DDE3  B5 00     LDA $00,X @ 55 = 55             A:AA X:55 Y:26 P:24 SP:FB CYC:11819
DDE5  20 AD F9  JSR $F9AD                       A:55 X:55 Y:26 P:24 SP:FB CYC:11823
F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:26 P:24 SP:F9 CYC:11829
F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:26 P:24 SP:F9 CYC:11831
F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:26 P:24 SP:F9 CYC:11833
F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:26 P:24 SP:F9 CYC:11835
F9B5  C9 55     CMP #$55                        A:55 X:55 Y:26 P:24 SP:F9 CYC:11837
F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:26 P:27 SP:F9 CYC:11839
F9B9  60        RTS                             A:55 X:55 Y:26 P:27 SP:F9 CYC:11841
DDE8  C8        INY                             A:55 X:55 Y:26 P:27 SP:FB CYC:11847
DDE9  20 BD F9  JSR $F9BD                       A:55 X:55 Y:27 P:25 SP:FB CYC:11849
F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:27 P:25 SP:F9 CYC:11855
F9BF  38        SEC                             A:55 X:55 Y:27 P:E5 SP:F9 CYC:11858
F9C0  A9 80     LDA #$80                        A:55 X:55 Y:27 P:E5 SP:F9 CYC:11860
F9C2  60        RTS                             A:80 X:55 Y:27 P:E5 SP:F9 CYC:11862
DDEC  95 00     STA $00,X @ 55 = 55             A:80 X:55 Y:27 P:E5 SP:FB CYC:11868
DDEE  16 00     ASL $00,X @ 55 = 80             A:80 X:55 Y:27 P:E5 SP:FB CYC:11872
DDF0  B5 00     LDA $00,X @ 55 = 00             A:80 X:55 Y:27 P:67 SP:FB CYC:11878
DDF2  20 C3 F9  JSR $F9C3                       A:0 X:55 Y:27 P:67 SP:FB CYC:11882
F9C3  90 1C     BCC $F9E1                       A:0 X:55 Y:27 P:67 SP:F9 CYC:11888
F9C5  D0 1A     BNE $F9E1                       A:0 X:55 Y:27 P:67 SP:F9 CYC:11890
F9C7  30 18     BMI $F9E1                       A:0 X:55 Y:27 P:67 SP:F9 CYC:11892
F9C9  50 16     BVC $F9E1                       A:0 X:55 Y:27 P:67 SP:F9 CYC:11894
F9CB  C9 00     CMP #$00                        A:0 X:55 Y:27 P:67 SP:F9 CYC:11896
F9CD  D0 12     BNE $F9E1                       A:0 X:55 Y:27 P:67 SP:F9 CYC:11898
F9CF  B8        CLV                             A:0 X:55 Y:27 P:67 SP:F9 CYC:11900
F9D0  A9 55     LDA #$55                        A:0 X:55 Y:27 P:27 SP:F9 CYC:11902
F9D2  38        SEC                             A:55 X:55 Y:27 P:25 SP:F9 CYC:11904
F9D3  60        RTS                             A:55 X:55 Y:27 P:25 SP:F9 CYC:11906
DDF5  C8        INY                             A:55 X:55 Y:27 P:25 SP:FB CYC:11912
DDF6  95 00     STA $00,X @ 55 = 00             A:55 X:55 Y:28 P:25 SP:FB CYC:11914
DDF8  16 00     ASL $00,X @ 55 = 55             A:55 X:55 Y:28 P:25 SP:FB CYC:11918
DDFA  B5 00     LDA $00,X @ 55 = AA             A:55 X:55 Y:28 P:A4 SP:FB CYC:11924
DDFC  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:28 P:A4 SP:FB CYC:11928
F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9 CYC:11934
F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9 CYC:11936
F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9 CYC:11938
F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9 CYC:11940
F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:28 P:A4 SP:F9 CYC:11942
F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:28 P:27 SP:F9 CYC:11944
F9E0  60        RTS                             A:AA X:55 Y:28 P:27 SP:F9 CYC:11946
DDFF  C8        INY                             A:AA X:55 Y:28 P:27 SP:FB CYC:11952
DE00  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:29 P:25 SP:FB CYC:11954
F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:29 P:25 SP:F9 CYC:11960
F9E6  38        SEC                             A:AA X:55 Y:29 P:E5 SP:F9 CYC:11963
F9E7  A9 01     LDA #$01                        A:AA X:55 Y:29 P:E5 SP:F9 CYC:11965
F9E9  60        RTS                             A:1 X:55 Y:29 P:65 SP:F9 CYC:11967
DE03  95 00     STA $00,X @ 55 = AA             A:1 X:55 Y:29 P:65 SP:FB CYC:11973
DE05  76 00     ROR $00,X @ 55 = 01             A:1 X:55 Y:29 P:65 SP:FB CYC:11977
DE07  B5 00     LDA $00,X @ 55 = 80             A:1 X:55 Y:29 P:E5 SP:FB CYC:11983
DE09  20 EA F9  JSR $F9EA                       A:80 X:55 Y:29 P:E5 SP:FB CYC:11987
F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:29 P:E5 SP:F9 CYC:11993
F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:29 P:E5 SP:F9 CYC:11995
F9EE  10 18     BPL $FA08                       A:80 X:55 Y:29 P:E5 SP:F9 CYC:11997
F9F0  50 16     BVC $FA08                       A:80 X:55 Y:29 P:E5 SP:F9 CYC:11999
F9F2  C9 80     CMP #$80                        A:80 X:55 Y:29 P:E5 SP:F9 CYC:12001
F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:29 P:67 SP:F9 CYC:12003
F9F6  B8        CLV                             A:80 X:55 Y:29 P:67 SP:F9 CYC:12005
F9F7  18        CLC                             A:80 X:55 Y:29 P:27 SP:F9 CYC:12007
F9F8  A9 55     LDA #$55                        A:80 X:55 Y:29 P:26 SP:F9 CYC:12009
F9FA  60        RTS                             A:55 X:55 Y:29 P:24 SP:F9 CYC:12011
DE0C  C8        INY                             A:55 X:55 Y:29 P:24 SP:FB CYC:12017
DE0D  95 00     STA $00,X @ 55 = 80             A:55 X:55 Y:2A P:24 SP:FB CYC:12019
DE0F  76 00     ROR $00,X @ 55 = 55             A:55 X:55 Y:2A P:24 SP:FB CYC:12023
DE11  B5 00     LDA $00,X @ 55 = 2A             A:55 X:55 Y:2A P:25 SP:FB CYC:12029
DE13  20 FB F9  JSR $F9FB                       A:2A X:55 Y:2A P:25 SP:FB CYC:12033
F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:2A P:25 SP:F9 CYC:12039
F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:2A P:25 SP:F9 CYC:12041
F9FF  30 07     BMI $FA08                       A:2A X:55 Y:2A P:25 SP:F9 CYC:12043
FA01  70 05     BVS $FA08                       A:2A X:55 Y:2A P:25 SP:F9 CYC:12045
FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:2A P:25 SP:F9 CYC:12047
FA05  D0 01     BNE $FA08                       A:2A X:55 Y:2A P:27 SP:F9 CYC:12049
FA07  60        RTS                             A:2A X:55 Y:2A P:27 SP:F9 CYC:12051
DE16  C8        INY                             A:2A X:55 Y:2A P:27 SP:FB CYC:12057
DE17  20 0A FA  JSR $FA0A                       A:2A X:55 Y:2B P:25 SP:FB CYC:12059
FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:2B P:25 SP:F9 CYC:12065
FA0C  38        SEC                             A:2A X:55 Y:2B P:E5 SP:F9 CYC:12068
FA0D  A9 80     LDA #$80                        A:2A X:55 Y:2B P:E5 SP:F9 CYC:12070
FA0F  60        RTS                             A:80 X:55 Y:2B P:E5 SP:F9 CYC:12072
DE1A  95 00     STA $00,X @ 55 = 2A             A:80 X:55 Y:2B P:E5 SP:FB CYC:12078
DE1C  36 00     ROL $00,X @ 55 = 80             A:80 X:55 Y:2B P:E5 SP:FB CYC:12082
DE1E  B5 00     LDA $00,X @ 55 = 01             A:80 X:55 Y:2B P:65 SP:FB CYC:12088
DE20  20 10 FA  JSR $FA10                       A:1 X:55 Y:2B P:65 SP:FB CYC:12092
FA10  90 1C     BCC $FA2E                       A:1 X:55 Y:2B P:65 SP:F9 CYC:12098
FA12  F0 1A     BEQ $FA2E                       A:1 X:55 Y:2B P:65 SP:F9 CYC:12100
FA14  30 18     BMI $FA2E                       A:1 X:55 Y:2B P:65 SP:F9 CYC:12102
FA16  50 16     BVC $FA2E                       A:1 X:55 Y:2B P:65 SP:F9 CYC:12104
FA18  C9 01     CMP #$01                        A:1 X:55 Y:2B P:65 SP:F9 CYC:12106
FA1A  D0 12     BNE $FA2E                       A:1 X:55 Y:2B P:67 SP:F9 CYC:12108
FA1C  B8        CLV                             A:1 X:55 Y:2B P:67 SP:F9 CYC:12110
FA1D  18        CLC                             A:1 X:55 Y:2B P:27 SP:F9 CYC:12112
FA1E  A9 55     LDA #$55                        A:1 X:55 Y:2B P:26 SP:F9 CYC:12114
FA20  60        RTS                             A:55 X:55 Y:2B P:24 SP:F9 CYC:12116
DE23  C8        INY                             A:55 X:55 Y:2B P:24 SP:FB CYC:12122
DE24  95 00     STA $00,X @ 55 = 01             A:55 X:55 Y:2C P:24 SP:FB CYC:12124
DE26  36 00     ROL $00,X @ 55 = 55             A:55 X:55 Y:2C P:24 SP:FB CYC:12128
DE28  B5 00     LDA $00,X @ 55 = AA             A:55 X:55 Y:2C P:A4 SP:FB CYC:12134
DE2A  20 21 FA  JSR $FA21                       A:AA X:55 Y:2C P:A4 SP:FB CYC:12138
FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9 CYC:12144
FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9 CYC:12146
FA25  10 07     BPL $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9 CYC:12148
FA27  70 05     BVS $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9 CYC:12150
FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:2C P:A4 SP:F9 CYC:12152
FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:2C P:27 SP:F9 CYC:12154
FA2D  60        RTS                             A:AA X:55 Y:2C P:27 SP:F9 CYC:12156
DE2D  A9 FF     LDA #$FF                        A:AA X:55 Y:2C P:27 SP:FB CYC:12162
DE2F  95 00     STA $00,X @ 55 = AA             A:FF X:55 Y:2C P:A5 SP:FB CYC:12164
DE31  85 01     STA $01 = FF                    A:FF X:55 Y:2C P:A5 SP:FB CYC:12168
DE33  24 01     BIT $01 = FF                    A:FF X:55 Y:2C P:A5 SP:FB CYC:12171
DE35  38        SEC                             A:FF X:55 Y:2C P:E5 SP:FB CYC:12174
DE36  F6 00     INC $00,X @ 55 = FF             A:FF X:55 Y:2C P:E5 SP:FB CYC:12176
DE38  D0 0C     BNE $DE46                       A:FF X:55 Y:2C P:67 SP:FB CYC:12182
DE3A  30 0A     BMI $DE46                       A:FF X:55 Y:2C P:67 SP:FB CYC:12184
DE3C  50 08     BVC $DE46                       A:FF X:55 Y:2C P:67 SP:FB CYC:12186
DE3E  90 06     BCC $DE46                       A:FF X:55 Y:2C P:67 SP:FB CYC:12188
DE40  B5 00     LDA $00,X @ 55 = 00             A:FF X:55 Y:2C P:67 SP:FB CYC:12190
DE42  C9 00     CMP #$00                        A:0 X:55 Y:2C P:67 SP:FB CYC:12194
DE44  F0 04     BEQ $DE4A                       A:0 X:55 Y:2C P:67 SP:FB CYC:12196
DE4A  A9 7F     LDA #$7F                        A:0 X:55 Y:2C P:67 SP:FB CYC:12199
DE4C  95 00     STA $00,X @ 55 = 00             A:7F X:55 Y:2C P:65 SP:FB CYC:12201
DE4E  B8        CLV                             A:7F X:55 Y:2C P:65 SP:FB CYC:12205
DE4F  18        CLC                             A:7F X:55 Y:2C P:25 SP:FB CYC:12207
DE50  F6 00     INC $00,X @ 55 = 7F             A:7F X:55 Y:2C P:24 SP:FB CYC:12209
DE52  F0 0C     BEQ $DE60                       A:7F X:55 Y:2C P:A4 SP:FB CYC:12215
DE54  10 0A     BPL $DE60                       A:7F X:55 Y:2C P:A4 SP:FB CYC:12217
DE56  70 08     BVS $DE60                       A:7F X:55 Y:2C P:A4 SP:FB CYC:12219
DE58  B0 06     BCS $DE60                       A:7F X:55 Y:2C P:A4 SP:FB CYC:12221
DE5A  B5 00     LDA $00,X @ 55 = 80             A:7F X:55 Y:2C P:A4 SP:FB CYC:12223
DE5C  C9 80     CMP #$80                        A:80 X:55 Y:2C P:A4 SP:FB CYC:12227
DE5E  F0 04     BEQ $DE64                       A:80 X:55 Y:2C P:27 SP:FB CYC:12229
DE64  A9 00     LDA #$00                        A:80 X:55 Y:2C P:27 SP:FB CYC:12232
DE66  95 00     STA $00,X @ 55 = 80             A:0 X:55 Y:2C P:27 SP:FB CYC:12234
DE68  24 01     BIT $01 = FF                    A:0 X:55 Y:2C P:27 SP:FB CYC:12238
DE6A  38        SEC                             A:0 X:55 Y:2C P:E7 SP:FB CYC:12241
DE6B  D6 00     DEC $00,X @ 55 = 00             A:0 X:55 Y:2C P:E7 SP:FB CYC:12243
DE6D  F0 0C     BEQ $DE7B                       A:0 X:55 Y:2C P:E5 SP:FB CYC:12249
DE6F  10 0A     BPL $DE7B                       A:0 X:55 Y:2C P:E5 SP:FB CYC:12251
DE71  50 08     BVC $DE7B                       A:0 X:55 Y:2C P:E5 SP:FB CYC:12253
DE73  90 06     BCC $DE7B                       A:0 X:55 Y:2C P:E5 SP:FB CYC:12255
DE75  B5 00     LDA $00,X @ 55 = FF             A:0 X:55 Y:2C P:E5 SP:FB CYC:12257
DE77  C9 FF     CMP #$FF                        A:FF X:55 Y:2C P:E5 SP:FB CYC:12261
DE79  F0 04     BEQ $DE7F                       A:FF X:55 Y:2C P:67 SP:FB CYC:12263
DE7F  A9 80     LDA #$80                        A:FF X:55 Y:2C P:67 SP:FB CYC:12266
DE81  95 00     STA $00,X @ 55 = FF             A:80 X:55 Y:2C P:E5 SP:FB CYC:12268
DE83  B8        CLV                             A:80 X:55 Y:2C P:E5 SP:FB CYC:12272
DE84  18        CLC                             A:80 X:55 Y:2C P:A5 SP:FB CYC:12274
DE85  D6 00     DEC $00,X @ 55 = 80             A:80 X:55 Y:2C P:A4 SP:FB CYC:12276
DE87  F0 0C     BEQ $DE95                       A:80 X:55 Y:2C P:24 SP:FB CYC:12282
DE89  30 0A     BMI $DE95                       A:80 X:55 Y:2C P:24 SP:FB CYC:12284
DE8B  70 08     BVS $DE95                       A:80 X:55 Y:2C P:24 SP:FB CYC:12286
DE8D  B0 06     BCS $DE95                       A:80 X:55 Y:2C P:24 SP:FB CYC:12288
DE8F  B5 00     LDA $00,X @ 55 = 7F             A:80 X:55 Y:2C P:24 SP:FB CYC:12290
DE91  C9 7F     CMP #$7F                        A:7F X:55 Y:2C P:24 SP:FB CYC:12294
DE93  F0 04     BEQ $DE99                       A:7F X:55 Y:2C P:27 SP:FB CYC:12296
DE99  A9 01     LDA #$01                        A:7F X:55 Y:2C P:27 SP:FB CYC:12299
DE9B  95 00     STA $00,X @ 55 = 7F             A:1 X:55 Y:2C P:25 SP:FB CYC:12301
DE9D  D6 00     DEC $00,X @ 55 = 01             A:1 X:55 Y:2C P:25 SP:FB CYC:12305
DE9F  F0 04     BEQ $DEA5                       A:1 X:55 Y:2C P:27 SP:FB CYC:12311
DEA5  A9 33     LDA #$33                        A:1 X:55 Y:2C P:27 SP:FB CYC:12314
DEA7  85 78     STA $78 = 7F                    A:33 X:55 Y:2C P:25 SP:FB CYC:12316
DEA9  A9 44     LDA #$44                        A:33 X:55 Y:2C P:25 SP:FB CYC:12319
DEAB  A0 78     LDY #$78                        A:44 X:55 Y:2C P:25 SP:FB CYC:12321
DEAD  A2 00     LDX #$00                        A:44 X:55 Y:78 P:25 SP:FB CYC:12323
DEAF  38        SEC                             A:44 X:0 Y:78 P:27 SP:FB CYC:12325
DEB0  24 01     BIT $01 = FF                    A:44 X:0 Y:78 P:27 SP:FB CYC:12327
DEB2  B6 00     LDX $00,Y @ 78 = 33             A:44 X:0 Y:78 P:E5 SP:FB CYC:12330
DEB4  90 12     BCC $DEC8                       A:44 X:33 Y:78 P:65 SP:FB CYC:12334
DEB6  50 10     BVC $DEC8                       A:44 X:33 Y:78 P:65 SP:FB CYC:12336
DEB8  30 0E     BMI $DEC8                       A:44 X:33 Y:78 P:65 SP:FB CYC:12338
DEBA  F0 0C     BEQ $DEC8                       A:44 X:33 Y:78 P:65 SP:FB CYC:12340
DEBC  E0 33     CPX #$33                        A:44 X:33 Y:78 P:65 SP:FB CYC:12342
DEBE  D0 08     BNE $DEC8                       A:44 X:33 Y:78 P:67 SP:FB CYC:12344
DEC0  C0 78     CPY #$78                        A:44 X:33 Y:78 P:67 SP:FB CYC:12346
DEC2  D0 04     BNE $DEC8                       A:44 X:33 Y:78 P:67 SP:FB CYC:12348
DEC4  C9 44     CMP #$44                        A:44 X:33 Y:78 P:67 SP:FB CYC:12350
DEC6  F0 04     BEQ $DECC                       A:44 X:33 Y:78 P:67 SP:FB CYC:12352
DECC  A9 97     LDA #$97                        A:44 X:33 Y:78 P:67 SP:FB CYC:12355
DECE  85 7F     STA $7F = 00                    A:97 X:33 Y:78 P:E5 SP:FB CYC:12357
DED0  A9 47     LDA #$47                        A:97 X:33 Y:78 P:E5 SP:FB CYC:12360
DED2  A0 FF     LDY #$FF                        A:47 X:33 Y:78 P:65 SP:FB CYC:12362
DED4  A2 00     LDX #$00                        A:47 X:33 Y:FF P:E5 SP:FB CYC:12364
DED6  18        CLC                             A:47 X:0 Y:FF P:67 SP:FB CYC:12366
DED7  B8        CLV                             A:47 X:0 Y:FF P:66 SP:FB CYC:12368
DED8  B6 80     LDX $80,Y @ 7F = 97             A:47 X:0 Y:FF P:26 SP:FB CYC:12370
DEDA  B0 12     BCS $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB CYC:12374
DEDC  70 10     BVS $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB CYC:12376
DEDE  10 0E     BPL $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB CYC:12378
DEE0  F0 0C     BEQ $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB CYC:12380
DEE2  E0 97     CPX #$97                        A:47 X:97 Y:FF P:A4 SP:FB CYC:12382
DEE4  D0 08     BNE $DEEE                       A:47 X:97 Y:FF P:27 SP:FB CYC:12384
DEE6  C0 FF     CPY #$FF                        A:47 X:97 Y:FF P:27 SP:FB CYC:12386
DEE8  D0 04     BNE $DEEE                       A:47 X:97 Y:FF P:27 SP:FB CYC:12388
DEEA  C9 47     CMP #$47                        A:47 X:97 Y:FF P:27 SP:FB CYC:12390
DEEC  F0 04     BEQ $DEF2                       A:47 X:97 Y:FF P:27 SP:FB CYC:12392
DEF2  A9 00     LDA #$00                        A:47 X:97 Y:FF P:27 SP:FB CYC:12395
DEF4  85 7F     STA $7F = 97                    A:0 X:97 Y:FF P:27 SP:FB CYC:12397
DEF6  A9 47     LDA #$47                        A:0 X:97 Y:FF P:27 SP:FB CYC:12400
DEF8  A0 FF     LDY #$FF                        A:47 X:97 Y:FF P:25 SP:FB CYC:12402
DEFA  A2 69     LDX #$69                        A:47 X:97 Y:FF P:A5 SP:FB CYC:12404
DEFC  18        CLC                             A:47 X:69 Y:FF P:25 SP:FB CYC:12406
DEFD  B8        CLV                             A:47 X:69 Y:FF P:24 SP:FB CYC:12408
DEFE  96 80     STX $80,Y @ 7F = 00             A:47 X:69 Y:FF P:24 SP:FB CYC:12410
DF00  B0 18     BCS $DF1A                       A:47 X:69 Y:FF P:24 SP:FB CYC:12414
DF02  70 16     BVS $DF1A                       A:47 X:69 Y:FF P:24 SP:FB CYC:12416
DF04  30 14     BMI $DF1A                       A:47 X:69 Y:FF P:24 SP:FB CYC:12418
DF06  F0 12     BEQ $DF1A                       A:47 X:69 Y:FF P:24 SP:FB CYC:12420
DF08  E0 69     CPX #$69                        A:47 X:69 Y:FF P:24 SP:FB CYC:12422
DF0A  D0 0E     BNE $DF1A                       A:47 X:69 Y:FF P:27 SP:FB CYC:12424
DF0C  C0 FF     CPY #$FF                        A:47 X:69 Y:FF P:27 SP:FB CYC:12426
DF0E  D0 0A     BNE $DF1A                       A:47 X:69 Y:FF P:27 SP:FB CYC:12428
DF10  C9 47     CMP #$47                        A:47 X:69 Y:FF P:27 SP:FB CYC:12430
DF12  D0 06     BNE $DF1A                       A:47 X:69 Y:FF P:27 SP:FB CYC:12432
DF14  A5 7F     LDA $7F = 69                    A:47 X:69 Y:FF P:27 SP:FB CYC:12434
DF16  C9 69     CMP #$69                        A:69 X:69 Y:FF P:25 SP:FB CYC:12437
DF18  F0 04     BEQ $DF1E                       A:69 X:69 Y:FF P:27 SP:FB CYC:12439
DF1E  A9 F5     LDA #$F5                        A:69 X:69 Y:FF P:27 SP:FB CYC:12442
DF20  85 4F     STA $4F = 00                    A:F5 X:69 Y:FF P:A5 SP:FB CYC:12444
DF22  A9 47     LDA #$47                        A:F5 X:69 Y:FF P:A5 SP:FB CYC:12447
DF24  A0 4F     LDY #$4F                        A:47 X:69 Y:FF P:25 SP:FB CYC:12449
DF26  24 01     BIT $01 = FF                    A:47 X:69 Y:4F P:25 SP:FB CYC:12451
DF28  A2 00     LDX #$00                        A:47 X:69 Y:4F P:E5 SP:FB CYC:12454
DF2A  38        SEC                             A:47 X:0 Y:4F P:67 SP:FB CYC:12456
DF2B  96 00     STX $00,Y @ 4F = F5             A:47 X:0 Y:4F P:67 SP:FB CYC:12458
DF2D  90 16     BCC $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12462
DF2F  50 14     BVC $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12464
DF31  30 12     BMI $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12466
DF33  D0 10     BNE $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12468
DF35  E0 00     CPX #$00                        A:47 X:0 Y:4F P:67 SP:FB CYC:12470
DF37  D0 0C     BNE $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12472
DF39  C0 4F     CPY #$4F                        A:47 X:0 Y:4F P:67 SP:FB CYC:12474
DF3B  D0 08     BNE $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12476
DF3D  C9 47     CMP #$47                        A:47 X:0 Y:4F P:67 SP:FB CYC:12478
DF3F  D0 04     BNE $DF45                       A:47 X:0 Y:4F P:67 SP:FB CYC:12480
DF41  A5 4F     LDA $4F = 00                    A:47 X:0 Y:4F P:67 SP:FB CYC:12482
DF43  F0 04     BEQ $DF49                       A:0 X:0 Y:4F P:67 SP:FB CYC:12485
DF49  60        RTS                             A:0 X:0 Y:4F P:67 SP:FB CYC:12488
C62C  20 AA E1  JSR $E1AA                       A:0 X:0 Y:4F P:67 SP:FD CYC:12494
E1AA  A9 FF     LDA #$FF                        A:0 X:0 Y:4F P:67 SP:FB CYC:12500
E1AC  85 01     STA $01 = FF                    A:FF X:0 Y:4F P:E5 SP:FB CYC:12502
E1AE  A9 AA     LDA #$AA                        A:FF X:0 Y:4F P:E5 SP:FB CYC:12505
E1B0  8D 33 06  STA $0633 = 00                  A:AA X:0 Y:4F P:E5 SP:FB CYC:12507
E1B3  A9 BB     LDA #$BB                        A:AA X:0 Y:4F P:E5 SP:FB CYC:12511
E1B5  8D 89 06  STA $0689 = 00                  A:BB X:0 Y:4F P:E5 SP:FB CYC:12513
E1B8  A2 00     LDX #$00                        A:BB X:0 Y:4F P:E5 SP:FB CYC:12517
E1BA  A9 66     LDA #$66                        A:BB X:0 Y:4F P:67 SP:FB CYC:12519
E1BC  24 01     BIT $01 = FF                    A:66 X:0 Y:4F P:65 SP:FB CYC:12521
E1BE  38        SEC                             A:66 X:0 Y:4F P:E5 SP:FB CYC:12524
E1BF  A0 00     LDY #$00                        A:66 X:0 Y:4F P:E5 SP:FB CYC:12526
E1C1  BC 33 06  LDY $0633,X @ 0633 = AA         A:66 X:0 Y:0 P:67 SP:FB CYC:12528
E1C4  10 12     BPL $E1D8                       A:66 X:0 Y:AA P:E5 SP:FB CYC:12532
E1C6  F0 10     BEQ $E1D8                       A:66 X:0 Y:AA P:E5 SP:FB CYC:12534
E1C8  50 0E     BVC $E1D8                       A:66 X:0 Y:AA P:E5 SP:FB CYC:12536
E1CA  90 0C     BCC $E1D8                       A:66 X:0 Y:AA P:E5 SP:FB CYC:12538
E1CC  C9 66     CMP #$66                        A:66 X:0 Y:AA P:E5 SP:FB CYC:12540
E1CE  D0 08     BNE $E1D8                       A:66 X:0 Y:AA P:67 SP:FB CYC:12542
E1D0  E0 00     CPX #$00                        A:66 X:0 Y:AA P:67 SP:FB CYC:12544
E1D2  D0 04     BNE $E1D8                       A:66 X:0 Y:AA P:67 SP:FB CYC:12546
E1D4  C0 AA     CPY #$AA                        A:66 X:0 Y:AA P:67 SP:FB CYC:12548
E1D6  F0 04     BEQ $E1DC                       A:66 X:0 Y:AA P:67 SP:FB CYC:12550
E1DC  A2 8A     LDX #$8A                        A:66 X:0 Y:AA P:67 SP:FB CYC:12553
E1DE  A9 66     LDA #$66                        A:66 X:8A Y:AA P:E5 SP:FB CYC:12555
E1E0  B8        CLV                             A:66 X:8A Y:AA P:65 SP:FB CYC:12557
E1E1  18        CLC                             A:66 X:8A Y:AA P:25 SP:FB CYC:12559
E1E2  A0 00     LDY #$00                        A:66 X:8A Y:AA P:24 SP:FB CYC:12561
E1E4  BC FF 05  LDY $05FF,X @ 0689 = BB         A:66 X:8A Y:0 P:26 SP:FB CYC:12563
E1E7  10 12     BPL $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB CYC:12568
E1E9  F0 10     BEQ $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB CYC:12570
E1EB  70 0E     BVS $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB CYC:12572
E1ED  B0 0C     BCS $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB CYC:12574
E1EF  C0 BB     CPY #$BB                        A:66 X:8A Y:BB P:A4 SP:FB CYC:12576
E1F1  D0 08     BNE $E1FB                       A:66 X:8A Y:BB P:27 SP:FB CYC:12578
E1F3  C9 66     CMP #$66                        A:66 X:8A Y:BB P:27 SP:FB CYC:12580
E1F5  D0 04     BNE $E1FB                       A:66 X:8A Y:BB P:27 SP:FB CYC:12582
E1F7  E0 8A     CPX #$8A                        A:66 X:8A Y:BB P:27 SP:FB CYC:12584
E1F9  F0 04     BEQ $E1FF                       A:66 X:8A Y:BB P:27 SP:FB CYC:12586
E1FF  A0 53     LDY #$53                        A:66 X:8A Y:BB P:27 SP:FB CYC:12589
E201  A9 AA     LDA #$AA                        A:66 X:8A Y:53 P:25 SP:FB CYC:12591
E203  A2 78     LDX #$78                        A:AA X:8A Y:53 P:A5 SP:FB CYC:12593
E205  8D 78 06  STA $0678 = 00                  A:AA X:78 Y:53 P:25 SP:FB CYC:12595
E208  20 B6 F7  JSR $F7B6                       A:AA X:78 Y:53 P:25 SP:FB CYC:12599
F7B6  18        CLC                             A:AA X:78 Y:53 P:25 SP:F9 CYC:12605
F7B7  A9 FF     LDA #$FF                        A:AA X:78 Y:53 P:24 SP:F9 CYC:12607
F7B9  85 01     STA $01 = FF                    A:FF X:78 Y:53 P:A4 SP:F9 CYC:12609
F7BB  24 01     BIT $01 = FF                    A:FF X:78 Y:53 P:A4 SP:F9 CYC:12612
F7BD  A9 55     LDA #$55                        A:FF X:78 Y:53 P:E4 SP:F9 CYC:12615
F7BF  60        RTS                             A:55 X:78 Y:53 P:64 SP:F9 CYC:12617
E20B  1D 00 06  ORA $0600,X @ 0678 = AA         A:55 X:78 Y:53 P:64 SP:FB CYC:12623
E20E  20 C0 F7  JSR $F7C0                       A:FF X:78 Y:53 P:E4 SP:FB CYC:12627
F7C0  B0 09     BCS $F7CB                       A:FF X:78 Y:53 P:E4 SP:F9 CYC:12633
F7C2  10 07     BPL $F7CB                       A:FF X:78 Y:53 P:E4 SP:F9 CYC:12635
F7C4  C9 FF     CMP #$FF                        A:FF X:78 Y:53 P:E4 SP:F9 CYC:12637
F7C6  D0 03     BNE $F7CB                       A:FF X:78 Y:53 P:67 SP:F9 CYC:12639
F7C8  50 01     BVC $F7CB                       A:FF X:78 Y:53 P:67 SP:F9 CYC:12641
F7CA  60        RTS                             A:FF X:78 Y:53 P:67 SP:F9 CYC:12643
E211  C8        INY                             A:FF X:78 Y:53 P:67 SP:FB CYC:12649
E212  A9 00     LDA #$00                        A:FF X:78 Y:54 P:65 SP:FB CYC:12651
E214  8D 78 06  STA $0678 = AA                  A:0 X:78 Y:54 P:67 SP:FB CYC:12653
E217  20 CE F7  JSR $F7CE                       A:0 X:78 Y:54 P:67 SP:FB CYC:12657
F7CE  38        SEC                             A:0 X:78 Y:54 P:67 SP:F9 CYC:12663
F7CF  B8        CLV                             A:0 X:78 Y:54 P:67 SP:F9 CYC:12665
F7D0  A9 00     LDA #$00                        A:0 X:78 Y:54 P:27 SP:F9 CYC:12667
F7D2  60        RTS                             A:0 X:78 Y:54 P:27 SP:F9 CYC:12669
E21A  1D 00 06  ORA $0600,X @ 0678 = 00         A:0 X:78 Y:54 P:27 SP:FB CYC:12675
E21D  20 D3 F7  JSR $F7D3                       A:0 X:78 Y:54 P:27 SP:FB CYC:12679
F7D3  D0 07     BNE $F7DC                       A:0 X:78 Y:54 P:27 SP:F9 CYC:12685
F7D5  70 05     BVS $F7DC                       A:0 X:78 Y:54 P:27 SP:F9 CYC:12687
F7D7  90 03     BCC $F7DC                       A:0 X:78 Y:54 P:27 SP:F9 CYC:12689
F7D9  30 01     BMI $F7DC                       A:0 X:78 Y:54 P:27 SP:F9 CYC:12691
F7DB  60        RTS                             A:0 X:78 Y:54 P:27 SP:F9 CYC:12693
E220  C8        INY                             A:0 X:78 Y:54 P:27 SP:FB CYC:12699
E221  A9 AA     LDA #$AA                        A:0 X:78 Y:55 P:25 SP:FB CYC:12701
E223  8D 78 06  STA $0678 = 00                  A:AA X:78 Y:55 P:A5 SP:FB CYC:12703
E226  20 DF F7  JSR $F7DF                       A:AA X:78 Y:55 P:A5 SP:FB CYC:12707
F7DF  18        CLC                             A:AA X:78 Y:55 P:A5 SP:F9 CYC:12713
F7E0  24 01     BIT $01 = FF                    A:AA X:78 Y:55 P:A4 SP:F9 CYC:12715
F7E2  A9 55     LDA #$55                        A:AA X:78 Y:55 P:E4 SP:F9 CYC:12718
F7E4  60        RTS                             A:55 X:78 Y:55 P:64 SP:F9 CYC:12720
E229  3D 00 06  AND $0600,X @ 0678 = AA         A:55 X:78 Y:55 P:64 SP:FB CYC:12726
E22C  20 E5 F7  JSR $F7E5                       A:0 X:78 Y:55 P:66 SP:FB CYC:12730
F7E5  D0 07     BNE $F7EE                       A:0 X:78 Y:55 P:66 SP:F9 CYC:12736
F7E7  50 05     BVC $F7EE                       A:0 X:78 Y:55 P:66 SP:F9 CYC:12738
F7E9  B0 03     BCS $F7EE                       A:0 X:78 Y:55 P:66 SP:F9 CYC:12740
F7EB  30 01     BMI $F7EE                       A:0 X:78 Y:55 P:66 SP:F9 CYC:12742
F7ED  60        RTS                             A:0 X:78 Y:55 P:66 SP:F9 CYC:12744
E22F  C8        INY                             A:0 X:78 Y:55 P:66 SP:FB CYC:12750
E230  A9 EF     LDA #$EF                        A:0 X:78 Y:56 P:64 SP:FB CYC:12752
E232  8D 78 06  STA $0678 = AA                  A:EF X:78 Y:56 P:E4 SP:FB CYC:12754
E235  20 F1 F7  JSR $F7F1                       A:EF X:78 Y:56 P:E4 SP:FB CYC:12758
F7F1  38        SEC                             A:EF X:78 Y:56 P:E4 SP:F9 CYC:12764
F7F2  B8        CLV                             A:EF X:78 Y:56 P:E5 SP:F9 CYC:12766
F7F3  A9 F8     LDA #$F8                        A:EF X:78 Y:56 P:A5 SP:F9 CYC:12768
F7F5  60        RTS                             A:F8 X:78 Y:56 P:A5 SP:F9 CYC:12770
E238  3D 00 06  AND $0600,X @ 0678 = EF         A:F8 X:78 Y:56 P:A5 SP:FB CYC:12776
E23B  20 F6 F7  JSR $F7F6                       A:E8 X:78 Y:56 P:A5 SP:FB CYC:12780
F7F6  90 09     BCC $F801                       A:E8 X:78 Y:56 P:A5 SP:F9 CYC:12786
F7F8  10 07     BPL $F801                       A:E8 X:78 Y:56 P:A5 SP:F9 CYC:12788
F7FA  C9 E8     CMP #$E8                        A:E8 X:78 Y:56 P:A5 SP:F9 CYC:12790
F7FC  D0 03     BNE $F801                       A:E8 X:78 Y:56 P:27 SP:F9 CYC:12792
F7FE  70 01     BVS $F801                       A:E8 X:78 Y:56 P:27 SP:F9 CYC:12794
F800  60        RTS                             A:E8 X:78 Y:56 P:27 SP:F9 CYC:12796
E23E  C8        INY                             A:E8 X:78 Y:56 P:27 SP:FB CYC:12802
E23F  A9 AA     LDA #$AA                        A:E8 X:78 Y:57 P:25 SP:FB CYC:12804
E241  8D 78 06  STA $0678 = EF                  A:AA X:78 Y:57 P:A5 SP:FB CYC:12806
E244  20 04 F8  JSR $F804                       A:AA X:78 Y:57 P:A5 SP:FB CYC:12810
F804  18        CLC                             A:AA X:78 Y:57 P:A5 SP:F9 CYC:12816
F805  24 01     BIT $01 = FF                    A:AA X:78 Y:57 P:A4 SP:F9 CYC:12818
F807  A9 5F     LDA #$5F                        A:AA X:78 Y:57 P:E4 SP:F9 CYC:12821
F809  60        RTS                             A:5F X:78 Y:57 P:64 SP:F9 CYC:12823
E247  5D 00 06  EOR $0600,X @ 0678 = AA         A:5F X:78 Y:57 P:64 SP:FB CYC:12829
E24A  20 0A F8  JSR $F80A                       A:F5 X:78 Y:57 P:E4 SP:FB CYC:12833
F80A  B0 09     BCS $F815                       A:F5 X:78 Y:57 P:E4 SP:F9 CYC:12839
F80C  10 07     BPL $F815                       A:F5 X:78 Y:57 P:E4 SP:F9 CYC:12841
F80E  C9 F5     CMP #$F5                        A:F5 X:78 Y:57 P:E4 SP:F9 CYC:12843
F810  D0 03     BNE $F815                       A:F5 X:78 Y:57 P:67 SP:F9 CYC:12845
F812  50 01     BVC $F815                       A:F5 X:78 Y:57 P:67 SP:F9 CYC:12847
F814  60        RTS                             A:F5 X:78 Y:57 P:67 SP:F9 CYC:12849
E24D  C8        INY                             A:F5 X:78 Y:57 P:67 SP:FB CYC:12855
E24E  A9 70     LDA #$70                        A:F5 X:78 Y:58 P:65 SP:FB CYC:12857
E250  8D 78 06  STA $0678 = AA                  A:70 X:78 Y:58 P:65 SP:FB CYC:12859
E253  20 18 F8  JSR $F818                       A:70 X:78 Y:58 P:65 SP:FB CYC:12863
F818  38        SEC                             A:70 X:78 Y:58 P:65 SP:F9 CYC:12869
F819  B8        CLV                             A:70 X:78 Y:58 P:65 SP:F9 CYC:12871
F81A  A9 70     LDA #$70                        A:70 X:78 Y:58 P:25 SP:F9 CYC:12873
F81C  60        RTS                             A:70 X:78 Y:58 P:25 SP:F9 CYC:12875
E256  5D 00 06  EOR $0600,X @ 0678 = 70         A:70 X:78 Y:58 P:25 SP:FB CYC:12881
E259  20 1D F8  JSR $F81D                       A:0 X:78 Y:58 P:27 SP:FB CYC:12885
F81D  D0 07     BNE $F826                       A:0 X:78 Y:58 P:27 SP:F9 CYC:12891
F81F  70 05     BVS $F826                       A:0 X:78 Y:58 P:27 SP:F9 CYC:12893
F821  90 03     BCC $F826                       A:0 X:78 Y:58 P:27 SP:F9 CYC:12895
F823  30 01     BMI $F826                       A:0 X:78 Y:58 P:27 SP:F9 CYC:12897
F825  60        RTS                             A:0 X:78 Y:58 P:27 SP:F9 CYC:12899
E25C  C8        INY                             A:0 X:78 Y:58 P:27 SP:FB CYC:12905
E25D  A9 69     LDA #$69                        A:0 X:78 Y:59 P:25 SP:FB CYC:12907
E25F  8D 78 06  STA $0678 = 70                  A:69 X:78 Y:59 P:25 SP:FB CYC:12909
E262  20 29 F8  JSR $F829                       A:69 X:78 Y:59 P:25 SP:FB CYC:12913
F829  18        CLC                             A:69 X:78 Y:59 P:25 SP:F9 CYC:12919
F82A  24 01     BIT $01 = FF                    A:69 X:78 Y:59 P:24 SP:F9 CYC:12921
F82C  A9 00     LDA #$00                        A:69 X:78 Y:59 P:E4 SP:F9 CYC:12924
F82E  60        RTS                             A:0 X:78 Y:59 P:66 SP:F9 CYC:12926
E265  7D 00 06  ADC $0600,X @ 0678 = 69         A:0 X:78 Y:59 P:66 SP:FB CYC:12932
E268  20 2F F8  JSR $F82F                       A:69 X:78 Y:59 P:24 SP:FB CYC:12936
F82F  30 09     BMI $F83A                       A:69 X:78 Y:59 P:24 SP:F9 CYC:12942
F831  B0 07     BCS $F83A                       A:69 X:78 Y:59 P:24 SP:F9 CYC:12944
F833  C9 69     CMP #$69                        A:69 X:78 Y:59 P:24 SP:F9 CYC:12946
F835  D0 03     BNE $F83A                       A:69 X:78 Y:59 P:27 SP:F9 CYC:12948
F837  70 01     BVS $F83A                       A:69 X:78 Y:59 P:27 SP:F9 CYC:12950
F839  60        RTS                             A:69 X:78 Y:59 P:27 SP:F9 CYC:12952
E26B  C8        INY                             A:69 X:78 Y:59 P:27 SP:FB CYC:12958
E26C  20 3D F8  JSR $F83D                       A:69 X:78 Y:5A P:25 SP:FB CYC:12960
F83D  38        SEC                             A:69 X:78 Y:5A P:25 SP:F9 CYC:12966
F83E  24 01     BIT $01 = FF                    A:69 X:78 Y:5A P:25 SP:F9 CYC:12968
F840  A9 00     LDA #$00                        A:69 X:78 Y:5A P:E5 SP:F9 CYC:12971
F842  60        RTS                             A:0 X:78 Y:5A P:67 SP:F9 CYC:12973
E26F  7D 00 06  ADC $0600,X @ 0678 = 69         A:0 X:78 Y:5A P:67 SP:FB CYC:12979
E272  20 43 F8  JSR $F843                       A:6A X:78 Y:5A P:24 SP:FB CYC:12983
F843  30 09     BMI $F84E                       A:6A X:78 Y:5A P:24 SP:F9 CYC:12989
F845  B0 07     BCS $F84E                       A:6A X:78 Y:5A P:24 SP:F9 CYC:12991
F847  C9 6A     CMP #$6A                        A:6A X:78 Y:5A P:24 SP:F9 CYC:12993
F849  D0 03     BNE $F84E                       A:6A X:78 Y:5A P:27 SP:F9 CYC:12995
F84B  70 01     BVS $F84E                       A:6A X:78 Y:5A P:27 SP:F9 CYC:12997
F84D  60        RTS                             A:6A X:78 Y:5A P:27 SP:F9 CYC:12999
E275  C8        INY                             A:6A X:78 Y:5A P:27 SP:FB CYC:13005
E276  A9 7F     LDA #$7F                        A:6A X:78 Y:5B P:25 SP:FB CYC:13007
E278  8D 78 06  STA $0678 = 69                  A:7F X:78 Y:5B P:25 SP:FB CYC:13009
E27B  20 51 F8  JSR $F851                       A:7F X:78 Y:5B P:25 SP:FB CYC:13013
F851  38        SEC                             A:7F X:78 Y:5B P:25 SP:F9 CYC:13019
F852  B8        CLV                             A:7F X:78 Y:5B P:25 SP:F9 CYC:13021
F853  A9 7F     LDA #$7F                        A:7F X:78 Y:5B P:25 SP:F9 CYC:13023
F855  60        RTS                             A:7F X:78 Y:5B P:25 SP:F9 CYC:13025
E27E  7D 00 06  ADC $0600,X @ 0678 = 7F         A:7F X:78 Y:5B P:25 SP:FB CYC:13031
E281  20 56 F8  JSR $F856                       A:FF X:78 Y:5B P:E4 SP:FB CYC:13035
F856  10 09     BPL $F861                       A:FF X:78 Y:5B P:E4 SP:F9 CYC:13041
F858  B0 07     BCS $F861                       A:FF X:78 Y:5B P:E4 SP:F9 CYC:13043
F85A  C9 FF     CMP #$FF                        A:FF X:78 Y:5B P:E4 SP:F9 CYC:13045
F85C  D0 03     BNE $F861                       A:FF X:78 Y:5B P:67 SP:F9 CYC:13047
F85E  50 01     BVC $F861                       A:FF X:78 Y:5B P:67 SP:F9 CYC:13049
F860  60        RTS                             A:FF X:78 Y:5B P:67 SP:F9 CYC:13051
E284  C8        INY                             A:FF X:78 Y:5B P:67 SP:FB CYC:13057
E285  A9 80     LDA #$80                        A:FF X:78 Y:5C P:65 SP:FB CYC:13059
E287  8D 78 06  STA $0678 = 7F                  A:80 X:78 Y:5C P:E5 SP:FB CYC:13061
E28A  20 64 F8  JSR $F864                       A:80 X:78 Y:5C P:E5 SP:FB CYC:13065
F864  18        CLC                             A:80 X:78 Y:5C P:E5 SP:F9 CYC:13071
F865  24 01     BIT $01 = FF                    A:80 X:78 Y:5C P:E4 SP:F9 CYC:13073
F867  A9 7F     LDA #$7F                        A:80 X:78 Y:5C P:E4 SP:F9 CYC:13076
F869  60        RTS                             A:7F X:78 Y:5C P:64 SP:F9 CYC:13078
E28D  7D 00 06  ADC $0600,X @ 0678 = 80         A:7F X:78 Y:5C P:64 SP:FB CYC:13084
E290  20 6A F8  JSR $F86A                       A:FF X:78 Y:5C P:A4 SP:FB CYC:13088
F86A  10 09     BPL $F875                       A:FF X:78 Y:5C P:A4 SP:F9 CYC:13094
F86C  B0 07     BCS $F875                       A:FF X:78 Y:5C P:A4 SP:F9 CYC:13096
F86E  C9 FF     CMP #$FF                        A:FF X:78 Y:5C P:A4 SP:F9 CYC:13098
F870  D0 03     BNE $F875                       A:FF X:78 Y:5C P:27 SP:F9 CYC:13100
F872  70 01     BVS $F875                       A:FF X:78 Y:5C P:27 SP:F9 CYC:13102
F874  60        RTS                             A:FF X:78 Y:5C P:27 SP:F9 CYC:13104
E293  C8        INY                             A:FF X:78 Y:5C P:27 SP:FB CYC:13110
E294  20 78 F8  JSR $F878                       A:FF X:78 Y:5D P:25 SP:FB CYC:13112
F878  38        SEC                             A:FF X:78 Y:5D P:25 SP:F9 CYC:13118
F879  B8        CLV                             A:FF X:78 Y:5D P:25 SP:F9 CYC:13120
F87A  A9 7F     LDA #$7F                        A:FF X:78 Y:5D P:25 SP:F9 CYC:13122
F87C  60        RTS                             A:7F X:78 Y:5D P:25 SP:F9 CYC:13124
E297  7D 00 06  ADC $0600,X @ 0678 = 80         A:7F X:78 Y:5D P:25 SP:FB CYC:13130
E29A  20 7D F8  JSR $F87D                       A:0 X:78 Y:5D P:27 SP:FB CYC:13134
F87D  D0 07     BNE $F886                       A:0 X:78 Y:5D P:27 SP:F9 CYC:13140
F87F  30 05     BMI $F886                       A:0 X:78 Y:5D P:27 SP:F9 CYC:13142
F881  70 03     BVS $F886                       A:0 X:78 Y:5D P:27 SP:F9 CYC:13144
F883  90 01     BCC $F886                       A:0 X:78 Y:5D P:27 SP:F9 CYC:13146
F885  60        RTS                             A:0 X:78 Y:5D P:27 SP:F9 CYC:13148
E29D  C8        INY                             A:0 X:78 Y:5D P:27 SP:FB CYC:13154
E29E  A9 40     LDA #$40                        A:0 X:78 Y:5E P:25 SP:FB CYC:13156
E2A0  8D 78 06  STA $0678 = 80                  A:40 X:78 Y:5E P:25 SP:FB CYC:13158
E2A3  20 89 F8  JSR $F889                       A:40 X:78 Y:5E P:25 SP:FB CYC:13162
F889  24 01     BIT $01 = FF                    A:40 X:78 Y:5E P:25 SP:F9 CYC:13168
F88B  A9 40     LDA #$40                        A:40 X:78 Y:5E P:E5 SP:F9 CYC:13171
F88D  60        RTS                             A:40 X:78 Y:5E P:65 SP:F9 CYC:13173
E2A6  DD 00 06  CMP $0600,X @ 0678 = 40         A:40 X:78 Y:5E P:65 SP:FB CYC:13179
E2A9  20 8E F8  JSR $F88E                       A:40 X:78 Y:5E P:67 SP:FB CYC:13183
F88E  30 07     BMI $F897                       A:40 X:78 Y:5E P:67 SP:F9 CYC:13189
F890  90 05     BCC $F897                       A:40 X:78 Y:5E P:67 SP:F9 CYC:13191
F892  D0 03     BNE $F897                       A:40 X:78 Y:5E P:67 SP:F9 CYC:13193
F894  50 01     BVC $F897                       A:40 X:78 Y:5E P:67 SP:F9 CYC:13195
F896  60        RTS                             A:40 X:78 Y:5E P:67 SP:F9 CYC:13197
E2AC  C8        INY                             A:40 X:78 Y:5E P:67 SP:FB CYC:13203
E2AD  48        PHA                             A:40 X:78 Y:5F P:65 SP:FB CYC:13205
E2AE  A9 3F     LDA #$3F                        A:40 X:78 Y:5F P:65 SP:FA CYC:13208
E2B0  8D 78 06  STA $0678 = 40                  A:3F X:78 Y:5F P:65 SP:FA CYC:13210
E2B3  68        PLA                             A:3F X:78 Y:5F P:65 SP:FA CYC:13214
E2B4  20 9A F8  JSR $F89A                       A:40 X:78 Y:5F P:65 SP:FB CYC:13218
F89A  B8        CLV                             A:40 X:78 Y:5F P:65 SP:F9 CYC:13224
F89B  60        RTS                             A:40 X:78 Y:5F P:25 SP:F9 CYC:13226
E2B7  DD 00 06  CMP $0600,X @ 0678 = 3F         A:40 X:78 Y:5F P:25 SP:FB CYC:13232
E2BA  20 9C F8  JSR $F89C                       A:40 X:78 Y:5F P:25 SP:FB CYC:13236
F89C  F0 07     BEQ $F8A5                       A:40 X:78 Y:5F P:25 SP:F9 CYC:13242
F89E  30 05     BMI $F8A5                       A:40 X:78 Y:5F P:25 SP:F9 CYC:13244
F8A0  90 03     BCC $F8A5                       A:40 X:78 Y:5F P:25 SP:F9 CYC:13246
F8A2  70 01     BVS $F8A5                       A:40 X:78 Y:5F P:25 SP:F9 CYC:13248
F8A4  60        RTS                             A:40 X:78 Y:5F P:25 SP:F9 CYC:13250
E2BD  C8        INY                             A:40 X:78 Y:5F P:25 SP:FB CYC:13256
E2BE  48        PHA                             A:40 X:78 Y:60 P:25 SP:FB CYC:13258
E2BF  A9 41     LDA #$41                        A:40 X:78 Y:60 P:25 SP:FA CYC:13261
E2C1  8D 78 06  STA $0678 = 3F                  A:41 X:78 Y:60 P:25 SP:FA CYC:13263
E2C4  68        PLA                             A:41 X:78 Y:60 P:25 SP:FA CYC:13267
E2C5  DD 00 06  CMP $0600,X @ 0678 = 41         A:40 X:78 Y:60 P:25 SP:FB CYC:13271
E2C8  20 A8 F8  JSR $F8A8                       A:40 X:78 Y:60 P:A4 SP:FB CYC:13275
F8A8  F0 05     BEQ $F8AF                       A:40 X:78 Y:60 P:A4 SP:F9 CYC:13281
F8AA  10 03     BPL $F8AF                       A:40 X:78 Y:60 P:A4 SP:F9 CYC:13283
F8AC  10 01     BPL $F8AF                       A:40 X:78 Y:60 P:A4 SP:F9 CYC:13285
F8AE  60        RTS                             A:40 X:78 Y:60 P:A4 SP:F9 CYC:13287
E2CB  C8        INY                             A:40 X:78 Y:60 P:A4 SP:FB CYC:13293
E2CC  48        PHA                             A:40 X:78 Y:61 P:24 SP:FB CYC:13295
E2CD  A9 00     LDA #$00                        A:40 X:78 Y:61 P:24 SP:FA CYC:13298
E2CF  8D 78 06  STA $0678 = 41                  A:0 X:78 Y:61 P:26 SP:FA CYC:13300
E2D2  68        PLA                             A:0 X:78 Y:61 P:26 SP:FA CYC:13304
E2D3  20 B2 F8  JSR $F8B2                       A:40 X:78 Y:61 P:24 SP:FB CYC:13308
F8B2  A9 80     LDA #$80                        A:40 X:78 Y:61 P:24 SP:F9 CYC:13314
F8B4  60        RTS                             A:80 X:78 Y:61 P:A4 SP:F9 CYC:13316
E2D6  DD 00 06  CMP $0600,X @ 0678 = 00         A:80 X:78 Y:61 P:A4 SP:FB CYC:13322
E2D9  20 B5 F8  JSR $F8B5                       A:80 X:78 Y:61 P:A5 SP:FB CYC:13326
F8B5  F0 05     BEQ $F8BC                       A:80 X:78 Y:61 P:A5 SP:F9 CYC:13332
F8B7  10 03     BPL $F8BC                       A:80 X:78 Y:61 P:A5 SP:F9 CYC:13334
F8B9  90 01     BCC $F8BC                       A:80 X:78 Y:61 P:A5 SP:F9 CYC:13336
F8BB  60        RTS                             A:80 X:78 Y:61 P:A5 SP:F9 CYC:13338
E2DC  C8        INY                             A:80 X:78 Y:61 P:A5 SP:FB CYC:13344
E2DD  48        PHA                             A:80 X:78 Y:62 P:25 SP:FB CYC:13346
E2DE  A9 80     LDA #$80                        A:80 X:78 Y:62 P:25 SP:FA CYC:13349
E2E0  8D 78 06  STA $0678 = 00                  A:80 X:78 Y:62 P:A5 SP:FA CYC:13351
E2E3  68        PLA                             A:80 X:78 Y:62 P:A5 SP:FA CYC:13355
E2E4  DD 00 06  CMP $0600,X @ 0678 = 80         A:80 X:78 Y:62 P:A5 SP:FB CYC:13359
E2E7  20 BF F8  JSR $F8BF                       A:80 X:78 Y:62 P:27 SP:FB CYC:13363
F8BF  D0 05     BNE $F8C6                       A:80 X:78 Y:62 P:27 SP:F9 CYC:13369
F8C1  30 03     BMI $F8C6                       A:80 X:78 Y:62 P:27 SP:F9 CYC:13371
F8C3  90 01     BCC $F8C6                       A:80 X:78 Y:62 P:27 SP:F9 CYC:13373
F8C5  60        RTS                             A:80 X:78 Y:62 P:27 SP:F9 CYC:13375
E2EA  C8        INY                             A:80 X:78 Y:62 P:27 SP:FB CYC:13381
E2EB  48        PHA                             A:80 X:78 Y:63 P:25 SP:FB CYC:13383
E2EC  A9 81     LDA #$81                        A:80 X:78 Y:63 P:25 SP:FA CYC:13386
E2EE  8D 78 06  STA $0678 = 80                  A:81 X:78 Y:63 P:A5 SP:FA CYC:13388
E2F1  68        PLA                             A:81 X:78 Y:63 P:A5 SP:FA CYC:13392
E2F2  DD 00 06  CMP $0600,X @ 0678 = 81         A:80 X:78 Y:63 P:A5 SP:FB CYC:13396
E2F5  20 C9 F8  JSR $F8C9                       A:80 X:78 Y:63 P:A4 SP:FB CYC:13400
F8C9  B0 05     BCS $F8D0                       A:80 X:78 Y:63 P:A4 SP:F9 CYC:13406
F8CB  F0 03     BEQ $F8D0                       A:80 X:78 Y:63 P:A4 SP:F9 CYC:13408
F8CD  10 01     BPL $F8D0                       A:80 X:78 Y:63 P:A4 SP:F9 CYC:13410
F8CF  60        RTS                             A:80 X:78 Y:63 P:A4 SP:F9 CYC:13412
E2F8  C8        INY                             A:80 X:78 Y:63 P:A4 SP:FB CYC:13418
E2F9  48        PHA                             A:80 X:78 Y:64 P:24 SP:FB CYC:13420
E2FA  A9 7F     LDA #$7F                        A:80 X:78 Y:64 P:24 SP:FA CYC:13423
E2FC  8D 78 06  STA $0678 = 81                  A:7F X:78 Y:64 P:24 SP:FA CYC:13425
E2FF  68        PLA                             A:7F X:78 Y:64 P:24 SP:FA CYC:13429
E300  DD 00 06  CMP $0600,X @ 0678 = 7F         A:80 X:78 Y:64 P:A4 SP:FB CYC:13433
E303  20 D3 F8  JSR $F8D3                       A:80 X:78 Y:64 P:25 SP:FB CYC:13437
F8D3  90 05     BCC $F8DA                       A:80 X:78 Y:64 P:25 SP:F9 CYC:13443
F8D5  F0 03     BEQ $F8DA                       A:80 X:78 Y:64 P:25 SP:F9 CYC:13445
F8D7  30 01     BMI $F8DA                       A:80 X:78 Y:64 P:25 SP:F9 CYC:13447
F8D9  60        RTS                             A:80 X:78 Y:64 P:25 SP:F9 CYC:13449
E306  C8        INY                             A:80 X:78 Y:64 P:25 SP:FB CYC:13455
E307  A9 40     LDA #$40                        A:80 X:78 Y:65 P:25 SP:FB CYC:13457
E309  8D 78 06  STA $0678 = 7F                  A:40 X:78 Y:65 P:25 SP:FB CYC:13459
E30C  20 31 F9  JSR $F931                       A:40 X:78 Y:65 P:25 SP:FB CYC:13463
F931  24 01     BIT $01 = FF                    A:40 X:78 Y:65 P:25 SP:F9 CYC:13469
F933  A9 40     LDA #$40                        A:40 X:78 Y:65 P:E5 SP:F9 CYC:13472
F935  38        SEC                             A:40 X:78 Y:65 P:65 SP:F9 CYC:13474
F936  60        RTS                             A:40 X:78 Y:65 P:65 SP:F9 CYC:13476
E30F  FD 00 06  SBC $0600,X @ 0678 = 40         A:40 X:78 Y:65 P:65 SP:FB CYC:13482
E312  20 37 F9  JSR $F937                       A:0 X:78 Y:65 P:27 SP:FB CYC:13486
F937  30 0B     BMI $F944                       A:0 X:78 Y:65 P:27 SP:F9 CYC:13492
F939  90 09     BCC $F944                       A:0 X:78 Y:65 P:27 SP:F9 CYC:13494
F93B  D0 07     BNE $F944                       A:0 X:78 Y:65 P:27 SP:F9 CYC:13496
F93D  70 05     BVS $F944                       A:0 X:78 Y:65 P:27 SP:F9 CYC:13498
F93F  C9 00     CMP #$00                        A:0 X:78 Y:65 P:27 SP:F9 CYC:13500
F941  D0 01     BNE $F944                       A:0 X:78 Y:65 P:27 SP:F9 CYC:13502
F943  60        RTS                             A:0 X:78 Y:65 P:27 SP:F9 CYC:13504
E315  C8        INY                             A:0 X:78 Y:65 P:27 SP:FB CYC:13510
E316  A9 3F     LDA #$3F                        A:0 X:78 Y:66 P:25 SP:FB CYC:13512
E318  8D 78 06  STA $0678 = 40                  A:3F X:78 Y:66 P:25 SP:FB CYC:13514
E31B  20 47 F9  JSR $F947                       A:3F X:78 Y:66 P:25 SP:FB CYC:13518
F947  B8        CLV                             A:3F X:78 Y:66 P:25 SP:F9 CYC:13524
F948  38        SEC                             A:3F X:78 Y:66 P:25 SP:F9 CYC:13526
F949  A9 40     LDA #$40                        A:3F X:78 Y:66 P:25 SP:F9 CYC:13528
F94B  60        RTS                             A:40 X:78 Y:66 P:25 SP:F9 CYC:13530
E31E  FD 00 06  SBC $0600,X @ 0678 = 3F         A:40 X:78 Y:66 P:25 SP:FB CYC:13536
E321  20 4C F9  JSR $F94C                       A:1 X:78 Y:66 P:25 SP:FB CYC:13540
F94C  F0 0B     BEQ $F959                       A:1 X:78 Y:66 P:25 SP:F9 CYC:13546
F94E  30 09     BMI $F959                       A:1 X:78 Y:66 P:25 SP:F9 CYC:13548
F950  90 07     BCC $F959                       A:1 X:78 Y:66 P:25 SP:F9 CYC:13550
F952  70 05     BVS $F959                       A:1 X:78 Y:66 P:25 SP:F9 CYC:13552
F954  C9 01     CMP #$01                        A:1 X:78 Y:66 P:25 SP:F9 CYC:13554
F956  D0 01     BNE $F959                       A:1 X:78 Y:66 P:27 SP:F9 CYC:13556
F958  60        RTS                             A:1 X:78 Y:66 P:27 SP:F9 CYC:13558
E324  C8        INY                             A:1 X:78 Y:66 P:27 SP:FB CYC:13564
E325  A9 41     LDA #$41                        A:1 X:78 Y:67 P:25 SP:FB CYC:13566
E327  8D 78 06  STA $0678 = 3F                  A:41 X:78 Y:67 P:25 SP:FB CYC:13568
E32A  20 5C F9  JSR $F95C                       A:41 X:78 Y:67 P:25 SP:FB CYC:13572
F95C  A9 40     LDA #$40                        A:41 X:78 Y:67 P:25 SP:F9 CYC:13578
F95E  38        SEC                             A:40 X:78 Y:67 P:25 SP:F9 CYC:13580
F95F  24 01     BIT $01 = FF                    A:40 X:78 Y:67 P:25 SP:F9 CYC:13582
F961  60        RTS                             A:40 X:78 Y:67 P:E5 SP:F9 CYC:13585
E32D  FD 00 06  SBC $0600,X @ 0678 = 41         A:40 X:78 Y:67 P:E5 SP:FB CYC:13591
E330  20 62 F9  JSR $F962                       A:FF X:78 Y:67 P:A4 SP:FB CYC:13595
F962  B0 0B     BCS $F96F                       A:FF X:78 Y:67 P:A4 SP:F9 CYC:13601
F964  F0 09     BEQ $F96F                       A:FF X:78 Y:67 P:A4 SP:F9 CYC:13603
F966  10 07     BPL $F96F                       A:FF X:78 Y:67 P:A4 SP:F9 CYC:13605
F968  70 05     BVS $F96F                       A:FF X:78 Y:67 P:A4 SP:F9 CYC:13607
F96A  C9 FF     CMP #$FF                        A:FF X:78 Y:67 P:A4 SP:F9 CYC:13609
F96C  D0 01     BNE $F96F                       A:FF X:78 Y:67 P:27 SP:F9 CYC:13611
F96E  60        RTS                             A:FF X:78 Y:67 P:27 SP:F9 CYC:13613
E333  C8        INY                             A:FF X:78 Y:67 P:27 SP:FB CYC:13619
E334  A9 00     LDA #$00                        A:FF X:78 Y:68 P:25 SP:FB CYC:13621
E336  8D 78 06  STA $0678 = 41                  A:0 X:78 Y:68 P:27 SP:FB CYC:13623
E339  20 72 F9  JSR $F972                       A:0 X:78 Y:68 P:27 SP:FB CYC:13627
F972  18        CLC                             A:0 X:78 Y:68 P:27 SP:F9 CYC:13633
F973  A9 80     LDA #$80                        A:0 X:78 Y:68 P:26 SP:F9 CYC:13635
F975  60        RTS                             A:80 X:78 Y:68 P:A4 SP:F9 CYC:13637
E33C  FD 00 06  SBC $0600,X @ 0678 = 00         A:80 X:78 Y:68 P:A4 SP:FB CYC:13643
E33F  20 76 F9  JSR $F976                       A:7F X:78 Y:68 P:65 SP:FB CYC:13647
F976  90 05     BCC $F97D                       A:7F X:78 Y:68 P:65 SP:F9 CYC:13653
F978  C9 7F     CMP #$7F                        A:7F X:78 Y:68 P:65 SP:F9 CYC:13655
F97A  D0 01     BNE $F97D                       A:7F X:78 Y:68 P:67 SP:F9 CYC:13657
F97C  60        RTS                             A:7F X:78 Y:68 P:67 SP:F9 CYC:13659
E342  C8        INY                             A:7F X:78 Y:68 P:67 SP:FB CYC:13665
E343  A9 7F     LDA #$7F                        A:7F X:78 Y:69 P:65 SP:FB CYC:13667
E345  8D 78 06  STA $0678 = 00                  A:7F X:78 Y:69 P:65 SP:FB CYC:13669
E348  20 80 F9  JSR $F980                       A:7F X:78 Y:69 P:65 SP:FB CYC:13673
F980  38        SEC                             A:7F X:78 Y:69 P:65 SP:F9 CYC:13679
F981  A9 81     LDA #$81                        A:7F X:78 Y:69 P:65 SP:F9 CYC:13681
F983  60        RTS                             A:81 X:78 Y:69 P:E5 SP:F9 CYC:13683
E34B  FD 00 06  SBC $0600,X @ 0678 = 7F         A:81 X:78 Y:69 P:E5 SP:FB CYC:13689
E34E  20 84 F9  JSR $F984                       A:2 X:78 Y:69 P:65 SP:FB CYC:13693
F984  50 07     BVC $F98D                       A:2 X:78 Y:69 P:65 SP:F9 CYC:13699
F986  90 05     BCC $F98D                       A:2 X:78 Y:69 P:65 SP:F9 CYC:13701
F988  C9 02     CMP #$02                        A:2 X:78 Y:69 P:65 SP:F9 CYC:13703
F98A  D0 01     BNE $F98D                       A:2 X:78 Y:69 P:67 SP:F9 CYC:13705
F98C  60        RTS                             A:2 X:78 Y:69 P:67 SP:F9 CYC:13707
E351  A9 AA     LDA #$AA                        A:2 X:78 Y:69 P:67 SP:FB CYC:13713
E353  8D 33 06  STA $0633 = AA                  A:AA X:78 Y:69 P:E5 SP:FB CYC:13715
E356  A9 BB     LDA #$BB                        A:AA X:78 Y:69 P:E5 SP:FB CYC:13719
E358  8D 89 06  STA $0689 = BB                  A:BB X:78 Y:69 P:E5 SP:FB CYC:13721
E35B  A2 00     LDX #$00                        A:BB X:78 Y:69 P:E5 SP:FB CYC:13725
E35D  A0 66     LDY #$66                        A:BB X:0 Y:69 P:67 SP:FB CYC:13727
E35F  24 01     BIT $01 = FF                    A:BB X:0 Y:66 P:65 SP:FB CYC:13729
E361  38        SEC                             A:BB X:0 Y:66 P:E5 SP:FB CYC:13732
E362  A9 00     LDA #$00                        A:BB X:0 Y:66 P:E5 SP:FB CYC:13734
E364  BD 33 06  LDA $0633,X @ 0633 = AA         A:0 X:0 Y:66 P:67 SP:FB CYC:13736
E367  10 12     BPL $E37B                       A:AA X:0 Y:66 P:E5 SP:FB CYC:13740
E369  F0 10     BEQ $E37B                       A:AA X:0 Y:66 P:E5 SP:FB CYC:13742
E36B  50 0E     BVC $E37B                       A:AA X:0 Y:66 P:E5 SP:FB CYC:13744
E36D  90 0C     BCC $E37B                       A:AA X:0 Y:66 P:E5 SP:FB CYC:13746
E36F  C0 66     CPY #$66                        A:AA X:0 Y:66 P:E5 SP:FB CYC:13748
E371  D0 08     BNE $E37B                       A:AA X:0 Y:66 P:67 SP:FB CYC:13750
E373  E0 00     CPX #$00                        A:AA X:0 Y:66 P:67 SP:FB CYC:13752
E375  D0 04     BNE $E37B                       A:AA X:0 Y:66 P:67 SP:FB CYC:13754
E377  C9 AA     CMP #$AA                        A:AA X:0 Y:66 P:67 SP:FB CYC:13756
E379  F0 04     BEQ $E37F                       A:AA X:0 Y:66 P:67 SP:FB CYC:13758
E37F  A2 8A     LDX #$8A                        A:AA X:0 Y:66 P:67 SP:FB CYC:13761
E381  A0 66     LDY #$66                        A:AA X:8A Y:66 P:E5 SP:FB CYC:13763
E383  B8        CLV                             A:AA X:8A Y:66 P:65 SP:FB CYC:13765
E384  18        CLC                             A:AA X:8A Y:66 P:25 SP:FB CYC:13767
E385  A9 00     LDA #$00                        A:AA X:8A Y:66 P:24 SP:FB CYC:13769
E387  BD FF 05  LDA $05FF,X @ 0689 = BB         A:0 X:8A Y:66 P:26 SP:FB CYC:13771
E38A  10 12     BPL $E39E                       A:BB X:8A Y:66 P:A4 SP:FB CYC:13776
E38C  F0 10     BEQ $E39E                       A:BB X:8A Y:66 P:A4 SP:FB CYC:13778
E38E  70 0E     BVS $E39E                       A:BB X:8A Y:66 P:A4 SP:FB CYC:13780
E390  B0 0C     BCS $E39E                       A:BB X:8A Y:66 P:A4 SP:FB CYC:13782
E392  C9 BB     CMP #$BB                        A:BB X:8A Y:66 P:A4 SP:FB CYC:13784
E394  D0 08     BNE $E39E                       A:BB X:8A Y:66 P:27 SP:FB CYC:13786
E396  C0 66     CPY #$66                        A:BB X:8A Y:66 P:27 SP:FB CYC:13788
E398  D0 04     BNE $E39E                       A:BB X:8A Y:66 P:27 SP:FB CYC:13790
E39A  E0 8A     CPX #$8A                        A:BB X:8A Y:66 P:27 SP:FB CYC:13792
E39C  F0 04     BEQ $E3A2                       A:BB X:8A Y:66 P:27 SP:FB CYC:13794
E3A2  24 01     BIT $01 = FF                    A:BB X:8A Y:66 P:27 SP:FB CYC:13797
E3A4  38        SEC                             A:BB X:8A Y:66 P:E5 SP:FB CYC:13800
E3A5  A9 44     LDA #$44                        A:BB X:8A Y:66 P:E5 SP:FB CYC:13802
E3A7  A2 00     LDX #$00                        A:44 X:8A Y:66 P:65 SP:FB CYC:13804
E3A9  9D 33 06  STA $0633,X @ 0633 = AA         A:44 X:0 Y:66 P:67 SP:FB CYC:13806
E3AC  AD 33 06  LDA $0633 = 44                  A:44 X:0 Y:66 P:67 SP:FB CYC:13811
E3AF  90 1A     BCC $E3CB                       A:44 X:0 Y:66 P:65 SP:FB CYC:13815
E3B1  C9 44     CMP #$44                        A:44 X:0 Y:66 P:65 SP:FB CYC:13817
E3B3  D0 16     BNE $E3CB                       A:44 X:0 Y:66 P:67 SP:FB CYC:13819
E3B5  50 14     BVC $E3CB                       A:44 X:0 Y:66 P:67 SP:FB CYC:13821
E3B7  18        CLC                             A:44 X:0 Y:66 P:67 SP:FB CYC:13823
E3B8  B8        CLV                             A:44 X:0 Y:66 P:66 SP:FB CYC:13825
E3B9  A9 99     LDA #$99                        A:44 X:0 Y:66 P:26 SP:FB CYC:13827
E3BB  A2 80     LDX #$80                        A:99 X:0 Y:66 P:A4 SP:FB CYC:13829
E3BD  9D 85 05  STA $0585,X @ 0605 = 00         A:99 X:80 Y:66 P:A4 SP:FB CYC:13831
E3C0  AD 05 06  LDA $0605 = 99                  A:99 X:80 Y:66 P:A4 SP:FB CYC:13836
E3C3  B0 06     BCS $E3CB                       A:99 X:80 Y:66 P:A4 SP:FB CYC:13840
E3C5  C9 99     CMP #$99                        A:99 X:80 Y:66 P:A4 SP:FB CYC:13842
E3C7  D0 02     BNE $E3CB                       A:99 X:80 Y:66 P:27 SP:FB CYC:13844
E3C9  50 04     BVC $E3CF                       A:99 X:80 Y:66 P:27 SP:FB CYC:13846
E3CF  A0 6D     LDY #$6D                        A:99 X:80 Y:66 P:27 SP:FB CYC:13849
E3D1  A2 6D     LDX #$6D                        A:99 X:80 Y:6D P:25 SP:FB CYC:13851
E3D3  20 90 F9  JSR $F990                       A:99 X:6D Y:6D P:25 SP:FB CYC:13853
F990  A2 55     LDX #$55                        A:99 X:6D Y:6D P:25 SP:F9 CYC:13859
F992  A9 FF     LDA #$FF                        A:99 X:55 Y:6D P:25 SP:F9 CYC:13861
F994  85 01     STA $01 = FF                    A:FF X:55 Y:6D P:A5 SP:F9 CYC:13863
F996  EA        NOP                             A:FF X:55 Y:6D P:A5 SP:F9 CYC:13866
F997  24 01     BIT $01 = FF                    A:FF X:55 Y:6D P:A5 SP:F9 CYC:13868
F999  38        SEC                             A:FF X:55 Y:6D P:E5 SP:F9 CYC:13871
F99A  A9 01     LDA #$01                        A:FF X:55 Y:6D P:E5 SP:F9 CYC:13873
F99C  60        RTS                             A:1 X:55 Y:6D P:65 SP:F9 CYC:13875
E3D6  9D 00 06  STA $0600,X @ 0655 = 00         A:1 X:55 Y:6D P:65 SP:FB CYC:13881
E3D9  5E 00 06  LSR $0600,X @ 0655 = 01         A:1 X:55 Y:6D P:65 SP:FB CYC:13886
E3DC  BD 00 06  LDA $0600,X @ 0655 = 00         A:1 X:55 Y:6D P:67 SP:FB CYC:13893
E3DF  20 9D F9  JSR $F99D                       A:0 X:55 Y:6D P:67 SP:FB CYC:13897
F99D  90 1B     BCC $F9BA                       A:0 X:55 Y:6D P:67 SP:F9 CYC:13903
F99F  D0 19     BNE $F9BA                       A:0 X:55 Y:6D P:67 SP:F9 CYC:13905
F9A1  30 17     BMI $F9BA                       A:0 X:55 Y:6D P:67 SP:F9 CYC:13907
F9A3  50 15     BVC $F9BA                       A:0 X:55 Y:6D P:67 SP:F9 CYC:13909
F9A5  C9 00     CMP #$00                        A:0 X:55 Y:6D P:67 SP:F9 CYC:13911
F9A7  D0 11     BNE $F9BA                       A:0 X:55 Y:6D P:67 SP:F9 CYC:13913
F9A9  B8        CLV                             A:0 X:55 Y:6D P:67 SP:F9 CYC:13915
F9AA  A9 AA     LDA #$AA                        A:0 X:55 Y:6D P:27 SP:F9 CYC:13917
F9AC  60        RTS                             A:AA X:55 Y:6D P:A5 SP:F9 CYC:13919
E3E2  C8        INY                             A:AA X:55 Y:6D P:A5 SP:FB CYC:13925
E3E3  9D 00 06  STA $0600,X @ 0655 = 00         A:AA X:55 Y:6E P:25 SP:FB CYC:13927
E3E6  5E 00 06  LSR $0600,X @ 0655 = AA         A:AA X:55 Y:6E P:25 SP:FB CYC:13932
E3E9  BD 00 06  LDA $0600,X @ 0655 = 55         A:AA X:55 Y:6E P:24 SP:FB CYC:13939
E3EC  20 AD F9  JSR $F9AD                       A:55 X:55 Y:6E P:24 SP:FB CYC:13943
F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:6E P:24 SP:F9 CYC:13949
F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:6E P:24 SP:F9 CYC:13951
F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:6E P:24 SP:F9 CYC:13953
F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:6E P:24 SP:F9 CYC:13955
F9B5  C9 55     CMP #$55                        A:55 X:55 Y:6E P:24 SP:F9 CYC:13957
F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:6E P:27 SP:F9 CYC:13959
F9B9  60        RTS                             A:55 X:55 Y:6E P:27 SP:F9 CYC:13961
E3EF  C8        INY                             A:55 X:55 Y:6E P:27 SP:FB CYC:13967
E3F0  20 BD F9  JSR $F9BD                       A:55 X:55 Y:6F P:25 SP:FB CYC:13969
F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:6F P:25 SP:F9 CYC:13975
F9BF  38        SEC                             A:55 X:55 Y:6F P:E5 SP:F9 CYC:13978
F9C0  A9 80     LDA #$80                        A:55 X:55 Y:6F P:E5 SP:F9 CYC:13980
F9C2  60        RTS                             A:80 X:55 Y:6F P:E5 SP:F9 CYC:13982
E3F3  9D 00 06  STA $0600,X @ 0655 = 55         A:80 X:55 Y:6F P:E5 SP:FB CYC:13988
E3F6  1E 00 06  ASL $0600,X @ 0655 = 80         A:80 X:55 Y:6F P:E5 SP:FB CYC:13993
E3F9  BD 00 06  LDA $0600,X @ 0655 = 00         A:80 X:55 Y:6F P:67 SP:FB CYC:14000
E3FC  20 C3 F9  JSR $F9C3                       A:0 X:55 Y:6F P:67 SP:FB CYC:14004
F9C3  90 1C     BCC $F9E1                       A:0 X:55 Y:6F P:67 SP:F9 CYC:14010
F9C5  D0 1A     BNE $F9E1                       A:0 X:55 Y:6F P:67 SP:F9 CYC:14012
F9C7  30 18     BMI $F9E1                       A:0 X:55 Y:6F P:67 SP:F9 CYC:14014
F9C9  50 16     BVC $F9E1                       A:0 X:55 Y:6F P:67 SP:F9 CYC:14016
F9CB  C9 00     CMP #$00                        A:0 X:55 Y:6F P:67 SP:F9 CYC:14018
F9CD  D0 12     BNE $F9E1                       A:0 X:55 Y:6F P:67 SP:F9 CYC:14020
F9CF  B8        CLV                             A:0 X:55 Y:6F P:67 SP:F9 CYC:14022
F9D0  A9 55     LDA #$55                        A:0 X:55 Y:6F P:27 SP:F9 CYC:14024
F9D2  38        SEC                             A:55 X:55 Y:6F P:25 SP:F9 CYC:14026
F9D3  60        RTS                             A:55 X:55 Y:6F P:25 SP:F9 CYC:14028
E3FF  C8        INY                             A:55 X:55 Y:6F P:25 SP:FB CYC:14034
E400  9D 00 06  STA $0600,X @ 0655 = 00         A:55 X:55 Y:70 P:25 SP:FB CYC:14036
E403  1E 00 06  ASL $0600,X @ 0655 = 55         A:55 X:55 Y:70 P:25 SP:FB CYC:14041
E406  BD 00 06  LDA $0600,X @ 0655 = AA         A:55 X:55 Y:70 P:A4 SP:FB CYC:14048
E409  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:70 P:A4 SP:FB CYC:14052
F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9 CYC:14058
F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9 CYC:14060
F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9 CYC:14062
F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9 CYC:14064
F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:70 P:A4 SP:F9 CYC:14066
F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:70 P:27 SP:F9 CYC:14068
F9E0  60        RTS                             A:AA X:55 Y:70 P:27 SP:F9 CYC:14070
E40C  C8        INY                             A:AA X:55 Y:70 P:27 SP:FB CYC:14076
E40D  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:71 P:25 SP:FB CYC:14078
F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:71 P:25 SP:F9 CYC:14084
F9E6  38        SEC                             A:AA X:55 Y:71 P:E5 SP:F9 CYC:14087
F9E7  A9 01     LDA #$01                        A:AA X:55 Y:71 P:E5 SP:F9 CYC:14089
F9E9  60        RTS                             A:1 X:55 Y:71 P:65 SP:F9 CYC:14091
E410  9D 00 06  STA $0600,X @ 0655 = AA         A:1 X:55 Y:71 P:65 SP:FB CYC:14097
E413  7E 00 06  ROR $0600,X @ 0655 = 01         A:1 X:55 Y:71 P:65 SP:FB CYC:14102
E416  BD 00 06  LDA $0600,X @ 0655 = 80         A:1 X:55 Y:71 P:E5 SP:FB CYC:14109
E419  20 EA F9  JSR $F9EA                       A:80 X:55 Y:71 P:E5 SP:FB CYC:14113
F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:71 P:E5 SP:F9 CYC:14119
F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:71 P:E5 SP:F9 CYC:14121
F9EE  10 18     BPL $FA08                       A:80 X:55 Y:71 P:E5 SP:F9 CYC:14123
F9F0  50 16     BVC $FA08                       A:80 X:55 Y:71 P:E5 SP:F9 CYC:14125
F9F2  C9 80     CMP #$80                        A:80 X:55 Y:71 P:E5 SP:F9 CYC:14127
F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:71 P:67 SP:F9 CYC:14129
F9F6  B8        CLV                             A:80 X:55 Y:71 P:67 SP:F9 CYC:14131
F9F7  18        CLC                             A:80 X:55 Y:71 P:27 SP:F9 CYC:14133
F9F8  A9 55     LDA #$55                        A:80 X:55 Y:71 P:26 SP:F9 CYC:14135
F9FA  60        RTS                             A:55 X:55 Y:71 P:24 SP:F9 CYC:14137
E41C  C8        INY                             A:55 X:55 Y:71 P:24 SP:FB CYC:14143
E41D  9D 00 06  STA $0600,X @ 0655 = 80         A:55 X:55 Y:72 P:24 SP:FB CYC:14145
E420  7E 00 06  ROR $0600,X @ 0655 = 55         A:55 X:55 Y:72 P:24 SP:FB CYC:14150
E423  BD 00 06  LDA $0600,X @ 0655 = 2A         A:55 X:55 Y:72 P:25 SP:FB CYC:14157
E426  20 FB F9  JSR $F9FB                       A:2A X:55 Y:72 P:25 SP:FB CYC:14161
F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:72 P:25 SP:F9 CYC:14167
F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:72 P:25 SP:F9 CYC:14169
F9FF  30 07     BMI $FA08                       A:2A X:55 Y:72 P:25 SP:F9 CYC:14171
FA01  70 05     BVS $FA08                       A:2A X:55 Y:72 P:25 SP:F9 CYC:14173
FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:72 P:25 SP:F9 CYC:14175
FA05  D0 01     BNE $FA08                       A:2A X:55 Y:72 P:27 SP:F9 CYC:14177
FA07  60        RTS                             A:2A X:55 Y:72 P:27 SP:F9 CYC:14179
E429  C8        INY                             A:2A X:55 Y:72 P:27 SP:FB CYC:14185
E42A  20 0A FA  JSR $FA0A                       A:2A X:55 Y:73 P:25 SP:FB CYC:14187
FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:73 P:25 SP:F9 CYC:14193
FA0C  38        SEC                             A:2A X:55 Y:73 P:E5 SP:F9 CYC:14196
FA0D  A9 80     LDA #$80                        A:2A X:55 Y:73 P:E5 SP:F9 CYC:14198
FA0F  60        RTS                             A:80 X:55 Y:73 P:E5 SP:F9 CYC:14200
E42D  9D 00 06  STA $0600,X @ 0655 = 2A         A:80 X:55 Y:73 P:E5 SP:FB CYC:14206
E430  3E 00 06  ROL $0600,X @ 0655 = 80         A:80 X:55 Y:73 P:E5 SP:FB CYC:14211
E433  BD 00 06  LDA $0600,X @ 0655 = 01         A:80 X:55 Y:73 P:65 SP:FB CYC:14218
E436  20 10 FA  JSR $FA10                       A:1 X:55 Y:73 P:65 SP:FB CYC:14222
FA10  90 1C     BCC $FA2E                       A:1 X:55 Y:73 P:65 SP:F9 CYC:14228
FA12  F0 1A     BEQ $FA2E                       A:1 X:55 Y:73 P:65 SP:F9 CYC:14230
FA14  30 18     BMI $FA2E                       A:1 X:55 Y:73 P:65 SP:F9 CYC:14232
FA16  50 16     BVC $FA2E                       A:1 X:55 Y:73 P:65 SP:F9 CYC:14234
FA18  C9 01     CMP #$01                        A:1 X:55 Y:73 P:65 SP:F9 CYC:14236
FA1A  D0 12     BNE $FA2E                       A:1 X:55 Y:73 P:67 SP:F9 CYC:14238
FA1C  B8        CLV                             A:1 X:55 Y:73 P:67 SP:F9 CYC:14240
FA1D  18        CLC                             A:1 X:55 Y:73 P:27 SP:F9 CYC:14242
FA1E  A9 55     LDA #$55                        A:1 X:55 Y:73 P:26 SP:F9 CYC:14244
FA20  60        RTS                             A:55 X:55 Y:73 P:24 SP:F9 CYC:14246
E439  C8        INY                             A:55 X:55 Y:73 P:24 SP:FB CYC:14252
E43A  9D 00 06  STA $0600,X @ 0655 = 01         A:55 X:55 Y:74 P:24 SP:FB CYC:14254
E43D  3E 00 06  ROL $0600,X @ 0655 = 55         A:55 X:55 Y:74 P:24 SP:FB CYC:14259
E440  BD 00 06  LDA $0600,X @ 0655 = AA         A:55 X:55 Y:74 P:A4 SP:FB CYC:14266
E443  20 21 FA  JSR $FA21                       A:AA X:55 Y:74 P:A4 SP:FB CYC:14270
FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9 CYC:14276
FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9 CYC:14278
FA25  10 07     BPL $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9 CYC:14280
FA27  70 05     BVS $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9 CYC:14282
FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:74 P:A4 SP:F9 CYC:14284
FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:74 P:27 SP:F9 CYC:14286
FA2D  60        RTS                             A:AA X:55 Y:74 P:27 SP:F9 CYC:14288
E446  A9 FF     LDA #$FF                        A:AA X:55 Y:74 P:27 SP:FB CYC:14294
E448  9D 00 06  STA $0600,X @ 0655 = AA         A:FF X:55 Y:74 P:A5 SP:FB CYC:14296
E44B  85 01     STA $01 = FF                    A:FF X:55 Y:74 P:A5 SP:FB CYC:14301
E44D  24 01     BIT $01 = FF                    A:FF X:55 Y:74 P:A5 SP:FB CYC:14304
E44F  38        SEC                             A:FF X:55 Y:74 P:E5 SP:FB CYC:14307
E450  FE 00 06  INC $0600,X @ 0655 = FF         A:FF X:55 Y:74 P:E5 SP:FB CYC:14309
E453  D0 0D     BNE $E462                       A:FF X:55 Y:74 P:67 SP:FB CYC:14316
E455  30 0B     BMI $E462                       A:FF X:55 Y:74 P:67 SP:FB CYC:14318
E457  50 09     BVC $E462                       A:FF X:55 Y:74 P:67 SP:FB CYC:14320
E459  90 07     BCC $E462                       A:FF X:55 Y:74 P:67 SP:FB CYC:14322
E45B  BD 00 06  LDA $0600,X @ 0655 = 00         A:FF X:55 Y:74 P:67 SP:FB CYC:14324
E45E  C9 00     CMP #$00                        A:0 X:55 Y:74 P:67 SP:FB CYC:14328
E460  F0 04     BEQ $E466                       A:0 X:55 Y:74 P:67 SP:FB CYC:14330
E466  A9 7F     LDA #$7F                        A:0 X:55 Y:74 P:67 SP:FB CYC:14333
E468  9D 00 06  STA $0600,X @ 0655 = 00         A:7F X:55 Y:74 P:65 SP:FB CYC:14335
E46B  B8        CLV                             A:7F X:55 Y:74 P:65 SP:FB CYC:14340
E46C  18        CLC                             A:7F X:55 Y:74 P:25 SP:FB CYC:14342
E46D  FE 00 06  INC $0600,X @ 0655 = 7F         A:7F X:55 Y:74 P:24 SP:FB CYC:14344
E470  F0 0D     BEQ $E47F                       A:7F X:55 Y:74 P:A4 SP:FB CYC:14351
E472  10 0B     BPL $E47F                       A:7F X:55 Y:74 P:A4 SP:FB CYC:14353
E474  70 09     BVS $E47F                       A:7F X:55 Y:74 P:A4 SP:FB CYC:14355
E476  B0 07     BCS $E47F                       A:7F X:55 Y:74 P:A4 SP:FB CYC:14357
E478  BD 00 06  LDA $0600,X @ 0655 = 80         A:7F X:55 Y:74 P:A4 SP:FB CYC:14359
E47B  C9 80     CMP #$80                        A:80 X:55 Y:74 P:A4 SP:FB CYC:14363
E47D  F0 04     BEQ $E483                       A:80 X:55 Y:74 P:27 SP:FB CYC:14365
E483  A9 00     LDA #$00                        A:80 X:55 Y:74 P:27 SP:FB CYC:14368
E485  9D 00 06  STA $0600,X @ 0655 = 80         A:0 X:55 Y:74 P:27 SP:FB CYC:14370
E488  24 01     BIT $01 = FF                    A:0 X:55 Y:74 P:27 SP:FB CYC:14375
E48A  38        SEC                             A:0 X:55 Y:74 P:E7 SP:FB CYC:14378
E48B  DE 00 06  DEC $0600,X @ 0655 = 00         A:0 X:55 Y:74 P:E7 SP:FB CYC:14380
E48E  F0 0D     BEQ $E49D                       A:0 X:55 Y:74 P:E5 SP:FB CYC:14387
E490  10 0B     BPL $E49D                       A:0 X:55 Y:74 P:E5 SP:FB CYC:14389
E492  50 09     BVC $E49D                       A:0 X:55 Y:74 P:E5 SP:FB CYC:14391
E494  90 07     BCC $E49D                       A:0 X:55 Y:74 P:E5 SP:FB CYC:14393
E496  BD 00 06  LDA $0600,X @ 0655 = FF         A:0 X:55 Y:74 P:E5 SP:FB CYC:14395
E499  C9 FF     CMP #$FF                        A:FF X:55 Y:74 P:E5 SP:FB CYC:14399
E49B  F0 04     BEQ $E4A1                       A:FF X:55 Y:74 P:67 SP:FB CYC:14401
E4A1  A9 80     LDA #$80                        A:FF X:55 Y:74 P:67 SP:FB CYC:14404
E4A3  9D 00 06  STA $0600,X @ 0655 = FF         A:80 X:55 Y:74 P:E5 SP:FB CYC:14406
E4A6  B8        CLV                             A:80 X:55 Y:74 P:E5 SP:FB CYC:14411
E4A7  18        CLC                             A:80 X:55 Y:74 P:A5 SP:FB CYC:14413
E4A8  DE 00 06  DEC $0600,X @ 0655 = 80         A:80 X:55 Y:74 P:A4 SP:FB CYC:14415
E4AB  F0 0D     BEQ $E4BA                       A:80 X:55 Y:74 P:24 SP:FB CYC:14422
E4AD  30 0B     BMI $E4BA                       A:80 X:55 Y:74 P:24 SP:FB CYC:14424
E4AF  70 09     BVS $E4BA                       A:80 X:55 Y:74 P:24 SP:FB CYC:14426
E4B1  B0 07     BCS $E4BA                       A:80 X:55 Y:74 P:24 SP:FB CYC:14428
E4B3  BD 00 06  LDA $0600,X @ 0655 = 7F         A:80 X:55 Y:74 P:24 SP:FB CYC:14430
E4B6  C9 7F     CMP #$7F                        A:7F X:55 Y:74 P:24 SP:FB CYC:14434
E4B8  F0 04     BEQ $E4BE                       A:7F X:55 Y:74 P:27 SP:FB CYC:14436
E4BE  A9 01     LDA #$01                        A:7F X:55 Y:74 P:27 SP:FB CYC:14439
E4C0  9D 00 06  STA $0600,X @ 0655 = 7F         A:1 X:55 Y:74 P:25 SP:FB CYC:14441
E4C3  DE 00 06  DEC $0600,X @ 0655 = 01         A:1 X:55 Y:74 P:25 SP:FB CYC:14446
E4C6  F0 04     BEQ $E4CC                       A:1 X:55 Y:74 P:27 SP:FB CYC:14453
E4CC  A9 33     LDA #$33                        A:1 X:55 Y:74 P:27 SP:FB CYC:14456
E4CE  8D 78 06  STA $0678 = 7F                  A:33 X:55 Y:74 P:25 SP:FB CYC:14458
E4D1  A9 44     LDA #$44                        A:33 X:55 Y:74 P:25 SP:FB CYC:14462
E4D3  A0 78     LDY #$78                        A:44 X:55 Y:74 P:25 SP:FB CYC:14464
E4D5  A2 00     LDX #$00                        A:44 X:55 Y:78 P:25 SP:FB CYC:14466
E4D7  38        SEC                             A:44 X:0 Y:78 P:27 SP:FB CYC:14468
E4D8  24 01     BIT $01 = FF                    A:44 X:0 Y:78 P:27 SP:FB CYC:14470
E4DA  BE 00 06  LDX $0600,Y @ 0678 = 33         A:44 X:0 Y:78 P:E5 SP:FB CYC:14473
E4DD  90 12     BCC $E4F1                       A:44 X:33 Y:78 P:65 SP:FB CYC:14477
E4DF  50 10     BVC $E4F1                       A:44 X:33 Y:78 P:65 SP:FB CYC:14479
E4E1  30 0E     BMI $E4F1                       A:44 X:33 Y:78 P:65 SP:FB CYC:14481
E4E3  F0 0C     BEQ $E4F1                       A:44 X:33 Y:78 P:65 SP:FB CYC:14483
E4E5  E0 33     CPX #$33                        A:44 X:33 Y:78 P:65 SP:FB CYC:14485
E4E7  D0 08     BNE $E4F1                       A:44 X:33 Y:78 P:67 SP:FB CYC:14487
E4E9  C0 78     CPY #$78                        A:44 X:33 Y:78 P:67 SP:FB CYC:14489
E4EB  D0 04     BNE $E4F1                       A:44 X:33 Y:78 P:67 SP:FB CYC:14491
E4ED  C9 44     CMP #$44                        A:44 X:33 Y:78 P:67 SP:FB CYC:14493
E4EF  F0 04     BEQ $E4F5                       A:44 X:33 Y:78 P:67 SP:FB CYC:14495
E4F5  A9 97     LDA #$97                        A:44 X:33 Y:78 P:67 SP:FB CYC:14498
E4F7  8D 7F 06  STA $067F = 00                  A:97 X:33 Y:78 P:E5 SP:FB CYC:14500
E4FA  A9 47     LDA #$47                        A:97 X:33 Y:78 P:E5 SP:FB CYC:14504
E4FC  A0 FF     LDY #$FF                        A:47 X:33 Y:78 P:65 SP:FB CYC:14506
E4FE  A2 00     LDX #$00                        A:47 X:33 Y:FF P:E5 SP:FB CYC:14508
E500  18        CLC                             A:47 X:0 Y:FF P:67 SP:FB CYC:14510
E501  B8        CLV                             A:47 X:0 Y:FF P:66 SP:FB CYC:14512
E502  BE 80 05  LDX $0580,Y @ 067F = 97         A:47 X:0 Y:FF P:26 SP:FB CYC:14514
E505  B0 12     BCS $E519                       A:47 X:97 Y:FF P:A4 SP:FB CYC:14519
E507  70 10     BVS $E519                       A:47 X:97 Y:FF P:A4 SP:FB CYC:14521
E509  10 0E     BPL $E519                       A:47 X:97 Y:FF P:A4 SP:FB CYC:14523
E50B  F0 0C     BEQ $E519                       A:47 X:97 Y:FF P:A4 SP:FB CYC:14525
E50D  E0 97     CPX #$97                        A:47 X:97 Y:FF P:A4 SP:FB CYC:14527
E50F  D0 08     BNE $E519                       A:47 X:97 Y:FF P:27 SP:FB CYC:14529
E511  C0 FF     CPY #$FF                        A:47 X:97 Y:FF P:27 SP:FB CYC:14531
E513  D0 04     BNE $E519                       A:47 X:97 Y:FF P:27 SP:FB CYC:14533
E515  C9 47     CMP #$47                        A:47 X:97 Y:FF P:27 SP:FB CYC:14535
E517  F0 04     BEQ $E51D                       A:47 X:97 Y:FF P:27 SP:FB CYC:14537
E51D  60        RTS                             A:47 X:97 Y:FF P:27 SP:FB CYC:14540
C62F  20 A3 C6  JSR $C6A3                       A:47 X:97 Y:FF P:27 SP:FD CYC:14546
C6A3  A0 4E     LDY #$4E                        A:47 X:97 Y:FF P:27 SP:FB CYC:14552
C6A5  A9 FF     LDA #$FF                        A:47 X:97 Y:4E P:25 SP:FB CYC:14554
C6A7  85 01     STA $01 = FF                    A:FF X:97 Y:4E P:A5 SP:FB CYC:14556
C6A9  20 B0 C6  JSR $C6B0                       A:FF X:97 Y:4E P:A5 SP:FB CYC:14559
C6B0  A9 FF     LDA #$FF                        A:FF X:97 Y:4E P:A5 SP:F9 CYC:14565
C6B2  48        PHA                             A:FF X:97 Y:4E P:A5 SP:F9 CYC:14567
C6B3  A9 AA     LDA #$AA                        A:FF X:97 Y:4E P:A5 SP:F8 CYC:14570
C6B5  D0 05     BNE $C6BC                       A:AA X:97 Y:4E P:A5 SP:F8 CYC:14572
C6BC  28        PLP                             A:AA X:97 Y:4E P:A5 SP:F8 CYC:14575
C6BD  04 A9    *NOP $A9 = 00                    A:AA X:97 Y:4E P:EF SP:F9 CYC:14579
C6BF  44 A9    *NOP $A9 = 00                    A:AA X:97 Y:4E P:EF SP:F9 CYC:14582
C6C1  64 A9    *NOP $A9 = 00                    A:AA X:97 Y:4E P:EF SP:F9 CYC:14585
C6C3  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9 CYC:14588
C6C4  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9 CYC:14590
C6C5  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9 CYC:14592
C6C6  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9 CYC:14594
C6C7  08        PHP                             A:AA X:97 Y:4E P:EF SP:F9 CYC:14596
C6C8  48        PHA                             A:AA X:97 Y:4E P:EF SP:F8 CYC:14599
C6C9  0C A9 A9 *NOP $A9A9 = A9                  A:AA X:97 Y:4E P:EF SP:F7 CYC:14602
C6CC  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7 CYC:14606
C6CD  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7 CYC:14608
C6CE  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7 CYC:14610
C6CF  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7 CYC:14612
C6D0  08        PHP                             A:AA X:97 Y:4E P:EF SP:F7 CYC:14614
C6D1  48        PHA                             A:AA X:97 Y:4E P:EF SP:F6 CYC:14617
C6D2  14 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5 CYC:14620
C6D4  34 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5 CYC:14624
C6D6  54 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5 CYC:14628
C6D8  74 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5 CYC:14632
C6DA  D4 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5 CYC:14636
C6DC  F4 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5 CYC:14640
C6DE  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5 CYC:14644
C6DF  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5 CYC:14646
C6E0  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5 CYC:14648
C6E1  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5 CYC:14650
C6E2  08        PHP                             A:AA X:97 Y:4E P:EF SP:F5 CYC:14652
C6E3  48        PHA                             A:AA X:97 Y:4E P:EF SP:F4 CYC:14655
C6E4  1A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14658
C6E5  3A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14660
C6E6  5A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14662
C6E7  7A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14664
C6E8  DA       *NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14666
C6E9  FA       *NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14668
C6EA  80 89    *NOP #$89                        A:AA X:97 Y:4E P:EF SP:F3 CYC:14670
C6EC  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14672
C6ED  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14674
C6EE  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14676
C6EF  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14678
C6F0  08        PHP                             A:AA X:97 Y:4E P:EF SP:F3 CYC:14680
C6F1  48        PHA                             A:AA X:97 Y:4E P:EF SP:F2 CYC:14683
C6F2  1C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1 CYC:14686
C6F5  3C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1 CYC:14691
C6F8  5C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1 CYC:14696
C6FB  7C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1 CYC:14701
C6FE  DC A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1 CYC:14706
C701  FC A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1 CYC:14711
C704  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1 CYC:14716
C705  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1 CYC:14718
C706  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1 CYC:14720
C707  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1 CYC:14722
C708  08        PHP                             A:AA X:97 Y:4E P:EF SP:F1 CYC:14724
C709  48        PHA                             A:AA X:97 Y:4E P:EF SP:F0 CYC:14727
C70A  A2 05     LDX #$05                        A:AA X:97 Y:4E P:EF SP:EF CYC:14730
C70C  68        PLA                             A:AA X:5 Y:4E P:6D SP:EF CYC:14732
C70D  C9 55     CMP #$55                        A:AA X:5 Y:4E P:ED SP:F0 CYC:14736
C70F  F0 0A     BEQ $C71B                       A:AA X:5 Y:4E P:6D SP:F0 CYC:14738
C711  C9 AA     CMP #$AA                        A:AA X:5 Y:4E P:6D SP:F0 CYC:14740
C713  F0 06     BEQ $C71B                       A:AA X:5 Y:4E P:6F SP:F0 CYC:14742
C71B  68        PLA                             A:AA X:5 Y:4E P:6F SP:F0 CYC:14745
C71C  29 CB     AND #$CB                        A:FF X:5 Y:4E P:ED SP:F1 CYC:14749
C71E  C9 00     CMP #$00                        A:CB X:5 Y:4E P:ED SP:F1 CYC:14751
C720  F0 06     BEQ $C728                       A:CB X:5 Y:4E P:ED SP:F1 CYC:14753
C722  C9 CB     CMP #$CB                        A:CB X:5 Y:4E P:ED SP:F1 CYC:14755
C724  F0 02     BEQ $C728                       A:CB X:5 Y:4E P:6F SP:F1 CYC:14757
C728  C8        INY                             A:CB X:5 Y:4E P:6F SP:F1 CYC:14760
C729  CA        DEX                             A:CB X:5 Y:4F P:6D SP:F1 CYC:14762
C72A  D0 E0     BNE $C70C                       A:CB X:4 Y:4F P:6D SP:F1 CYC:14764
C70C  68        PLA                             A:CB X:4 Y:4F P:6D SP:F1 CYC:14767
C70D  C9 55     CMP #$55                        A:AA X:4 Y:4F P:ED SP:F2 CYC:14771
C70F  F0 0A     BEQ $C71B                       A:AA X:4 Y:4F P:6D SP:F2 CYC:14773
C711  C9 AA     CMP #$AA                        A:AA X:4 Y:4F P:6D SP:F2 CYC:14775
C713  F0 06     BEQ $C71B                       A:AA X:4 Y:4F P:6F SP:F2 CYC:14777
C71B  68        PLA                             A:AA X:4 Y:4F P:6F SP:F2 CYC:14780
C71C  29 CB     AND #$CB                        A:FF X:4 Y:4F P:ED SP:F3 CYC:14784
C71E  C9 00     CMP #$00                        A:CB X:4 Y:4F P:ED SP:F3 CYC:14786
C720  F0 06     BEQ $C728                       A:CB X:4 Y:4F P:ED SP:F3 CYC:14788
C722  C9 CB     CMP #$CB                        A:CB X:4 Y:4F P:ED SP:F3 CYC:14790
C724  F0 02     BEQ $C728                       A:CB X:4 Y:4F P:6F SP:F3 CYC:14792
C728  C8        INY                             A:CB X:4 Y:4F P:6F SP:F3 CYC:14795
C729  CA        DEX                             A:CB X:4 Y:50 P:6D SP:F3 CYC:14797
C72A  D0 E0     BNE $C70C                       A:CB X:3 Y:50 P:6D SP:F3 CYC:14799
C70C  68        PLA                             A:CB X:3 Y:50 P:6D SP:F3 CYC:14802
C70D  C9 55     CMP #$55                        A:AA X:3 Y:50 P:ED SP:F4 CYC:14806
C70F  F0 0A     BEQ $C71B                       A:AA X:3 Y:50 P:6D SP:F4 CYC:14808
C711  C9 AA     CMP #$AA                        A:AA X:3 Y:50 P:6D SP:F4 CYC:14810
C713  F0 06     BEQ $C71B                       A:AA X:3 Y:50 P:6F SP:F4 CYC:14812
C71B  68        PLA                             A:AA X:3 Y:50 P:6F SP:F4 CYC:14815
C71C  29 CB     AND #$CB                        A:FF X:3 Y:50 P:ED SP:F5 CYC:14819
C71E  C9 00     CMP #$00                        A:CB X:3 Y:50 P:ED SP:F5 CYC:14821
C720  F0 06     BEQ $C728                       A:CB X:3 Y:50 P:ED SP:F5 CYC:14823
C722  C9 CB     CMP #$CB                        A:CB X:3 Y:50 P:ED SP:F5 CYC:14825
C724  F0 02     BEQ $C728                       A:CB X:3 Y:50 P:6F SP:F5 CYC:14827
C728  C8        INY                             A:CB X:3 Y:50 P:6F SP:F5 CYC:14830
C729  CA        DEX                             A:CB X:3 Y:51 P:6D SP:F5 CYC:14832
C72A  D0 E0     BNE $C70C                       A:CB X:2 Y:51 P:6D SP:F5 CYC:14834
C70C  68        PLA                             A:CB X:2 Y:51 P:6D SP:F5 CYC:14837
C70D  C9 55     CMP #$55                        A:AA X:2 Y:51 P:ED SP:F6 CYC:14841
C70F  F0 0A     BEQ $C71B                       A:AA X:2 Y:51 P:6D SP:F6 CYC:14843
C711  C9 AA     CMP #$AA                        A:AA X:2 Y:51 P:6D SP:F6 CYC:14845
C713  F0 06     BEQ $C71B                       A:AA X:2 Y:51 P:6F SP:F6 CYC:14847
C71B  68        PLA                             A:AA X:2 Y:51 P:6F SP:F6 CYC:14850
C71C  29 CB     AND #$CB                        A:FF X:2 Y:51 P:ED SP:F7 CYC:14854
C71E  C9 00     CMP #$00                        A:CB X:2 Y:51 P:ED SP:F7 CYC:14856
C720  F0 06     BEQ $C728                       A:CB X:2 Y:51 P:ED SP:F7 CYC:14858
C722  C9 CB     CMP #$CB                        A:CB X:2 Y:51 P:ED SP:F7 CYC:14860
C724  F0 02     BEQ $C728                       A:CB X:2 Y:51 P:6F SP:F7 CYC:14862
C728  C8        INY                             A:CB X:2 Y:51 P:6F SP:F7 CYC:14865
C729  CA        DEX                             A:CB X:2 Y:52 P:6D SP:F7 CYC:14867
C72A  D0 E0     BNE $C70C                       A:CB X:1 Y:52 P:6D SP:F7 CYC:14869
C70C  68        PLA                             A:CB X:1 Y:52 P:6D SP:F7 CYC:14872
C70D  C9 55     CMP #$55                        A:AA X:1 Y:52 P:ED SP:F8 CYC:14876
C70F  F0 0A     BEQ $C71B                       A:AA X:1 Y:52 P:6D SP:F8 CYC:14878
C711  C9 AA     CMP #$AA                        A:AA X:1 Y:52 P:6D SP:F8 CYC:14880
C713  F0 06     BEQ $C71B                       A:AA X:1 Y:52 P:6F SP:F8 CYC:14882
C71B  68        PLA                             A:AA X:1 Y:52 P:6F SP:F8 CYC:14885
C71C  29 CB     AND #$CB                        A:FF X:1 Y:52 P:ED SP:F9 CYC:14889
C71E  C9 00     CMP #$00                        A:CB X:1 Y:52 P:ED SP:F9 CYC:14891
C720  F0 06     BEQ $C728                       A:CB X:1 Y:52 P:ED SP:F9 CYC:14893
C722  C9 CB     CMP #$CB                        A:CB X:1 Y:52 P:ED SP:F9 CYC:14895
C724  F0 02     BEQ $C728                       A:CB X:1 Y:52 P:6F SP:F9 CYC:14897
C728  C8        INY                             A:CB X:1 Y:52 P:6F SP:F9 CYC:14900
C729  CA        DEX                             A:CB X:1 Y:53 P:6D SP:F9 CYC:14902
C72A  D0 E0     BNE $C70C                       A:CB X:0 Y:53 P:6F SP:F9 CYC:14904
C72C  60        RTS                             A:CB X:0 Y:53 P:6F SP:F9 CYC:14906
C6AC  20 B7 C6  JSR $C6B7                       A:CB X:0 Y:53 P:6F SP:FB CYC:14912
C6B7  A9 34     LDA #$34                        A:CB X:0 Y:53 P:6F SP:F9 CYC:14918
C6B9  48        PHA                             A:34 X:0 Y:53 P:6D SP:F9 CYC:14920
C6BA  A9 55     LDA #$55                        A:34 X:0 Y:53 P:6D SP:F8 CYC:14923
C6BC  28        PLP                             A:55 X:0 Y:53 P:6D SP:F8 CYC:14925
C6BD  04 A9    *NOP $A9 = 00                    A:55 X:0 Y:53 P:24 SP:F9 CYC:14929
C6BF  44 A9    *NOP $A9 = 00                    A:55 X:0 Y:53 P:24 SP:F9 CYC:14932
C6C1  64 A9    *NOP $A9 = 00                    A:55 X:0 Y:53 P:24 SP:F9 CYC:14935
C6C3  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F9 CYC:14938
C6C4  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F9 CYC:14940
C6C5  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F9 CYC:14942
C6C6  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F9 CYC:14944
C6C7  08        PHP                             A:55 X:0 Y:53 P:24 SP:F9 CYC:14946
C6C8  48        PHA                             A:55 X:0 Y:53 P:24 SP:F8 CYC:14949
C6C9  0C A9 A9 *NOP $A9A9 = A9                  A:55 X:0 Y:53 P:24 SP:F7 CYC:14952
C6CC  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F7 CYC:14956
C6CD  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F7 CYC:14958
C6CE  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F7 CYC:14960
C6CF  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F7 CYC:14962
C6D0  08        PHP                             A:55 X:0 Y:53 P:24 SP:F7 CYC:14964
C6D1  48        PHA                             A:55 X:0 Y:53 P:24 SP:F6 CYC:14967
C6D2  14 A9    *NOP $A9,X @ A9 = 00             A:55 X:0 Y:53 P:24 SP:F5 CYC:14970
C6D4  34 A9    *NOP $A9,X @ A9 = 00             A:55 X:0 Y:53 P:24 SP:F5 CYC:14974
C6D6  54 A9    *NOP $A9,X @ A9 = 00             A:55 X:0 Y:53 P:24 SP:F5 CYC:14978
C6D8  74 A9    *NOP $A9,X @ A9 = 00             A:55 X:0 Y:53 P:24 SP:F5 CYC:14982
C6DA  D4 A9    *NOP $A9,X @ A9 = 00             A:55 X:0 Y:53 P:24 SP:F5 CYC:14986
C6DC  F4 A9    *NOP $A9,X @ A9 = 00             A:55 X:0 Y:53 P:24 SP:F5 CYC:14990
C6DE  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F5 CYC:14994
C6DF  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F5 CYC:14996
C6E0  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F5 CYC:14998
C6E1  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F5 CYC:15000
C6E2  08        PHP                             A:55 X:0 Y:53 P:24 SP:F5 CYC:15002
C6E3  48        PHA                             A:55 X:0 Y:53 P:24 SP:F4 CYC:15005
C6E4  1A       *NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15008
C6E5  3A       *NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15010
C6E6  5A       *NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15012
C6E7  7A       *NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15014
C6E8  DA       *NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15016
C6E9  FA       *NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15018
C6EA  80 89    *NOP #$89                        A:55 X:0 Y:53 P:24 SP:F3 CYC:15020
C6EC  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15022
C6ED  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15024
C6EE  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15026
C6EF  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15028
C6F0  08        PHP                             A:55 X:0 Y:53 P:24 SP:F3 CYC:15030
C6F1  48        PHA                             A:55 X:0 Y:53 P:24 SP:F2 CYC:15033
C6F2  1C A9 A9 *NOP $A9A9,X @ A9A9 = A9         A:55 X:0 Y:53 P:24 SP:F1 CYC:15036
C6F5  3C A9 A9 *NOP $A9A9,X @ A9A9 = A9         A:55 X:0 Y:53 P:24 SP:F1 CYC:15040
C6F8  5C A9 A9 *NOP $A9A9,X @ A9A9 = A9         A:55 X:0 Y:53 P:24 SP:F1 CYC:15044
C6FB  7C A9 A9 *NOP $A9A9,X @ A9A9 = A9         A:55 X:0 Y:53 P:24 SP:F1 CYC:15048
C6FE  DC A9 A9 *NOP $A9A9,X @ A9A9 = A9         A:55 X:0 Y:53 P:24 SP:F1 CYC:15052
C701  FC A9 A9 *NOP $A9A9,X @ A9A9 = A9         A:55 X:0 Y:53 P:24 SP:F1 CYC:15056
C704  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F1 CYC:15060
C705  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F1 CYC:15062
C706  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F1 CYC:15064
C707  EA        NOP                             A:55 X:0 Y:53 P:24 SP:F1 CYC:15066
C708  08        PHP                             A:55 X:0 Y:53 P:24 SP:F1 CYC:15068
C709  48        PHA                             A:55 X:0 Y:53 P:24 SP:F0 CYC:15071
C70A  A2 05     LDX #$05                        A:55 X:0 Y:53 P:24 SP:EF CYC:15074
C70C  68        PLA                             A:55 X:5 Y:53 P:24 SP:EF CYC:15076
C70D  C9 55     CMP #$55                        A:55 X:5 Y:53 P:24 SP:F0 CYC:15080
C70F  F0 0A     BEQ $C71B                       A:55 X:5 Y:53 P:27 SP:F0 CYC:15082
C71B  68        PLA                             A:55 X:5 Y:53 P:27 SP:F0 CYC:15085
C71C  29 CB     AND #$CB                        A:34 X:5 Y:53 P:25 SP:F1 CYC:15089
C71E  C9 00     CMP #$00                        A:0 X:5 Y:53 P:27 SP:F1 CYC:15091
C720  F0 06     BEQ $C728                       A:0 X:5 Y:53 P:27 SP:F1 CYC:15093
C728  C8        INY                             A:0 X:5 Y:53 P:27 SP:F1 CYC:15096
C729  CA        DEX                             A:0 X:5 Y:54 P:25 SP:F1 CYC:15098
C72A  D0 E0     BNE $C70C                       A:0 X:4 Y:54 P:25 SP:F1 CYC:15100
C70C  68        PLA                             A:0 X:4 Y:54 P:25 SP:F1 CYC:15103
C70D  C9 55     CMP #$55                        A:55 X:4 Y:54 P:25 SP:F2 CYC:15107
C70F  F0 0A     BEQ $C71B                       A:55 X:4 Y:54 P:27 SP:F2 CYC:15109
C71B  68        PLA                             A:55 X:4 Y:54 P:27 SP:F2 CYC:15112
C71C  29 CB     AND #$CB                        A:34 X:4 Y:54 P:25 SP:F3 CYC:15116
C71E  C9 00     CMP #$00                        A:0 X:4 Y:54 P:27 SP:F3 CYC:15118
C720  F0 06     BEQ $C728                       A:0 X:4 Y:54 P:27 SP:F3 CYC:15120
C728  C8        INY                             A:0 X:4 Y:54 P:27 SP:F3 CYC:15123
C729  CA        DEX                             A:0 X:4 Y:55 P:25 SP:F3 CYC:15125
C72A  D0 E0     BNE $C70C                       A:0 X:3 Y:55 P:25 SP:F3 CYC:15127
C70C  68        PLA                             A:0 X:3 Y:55 P:25 SP:F3 CYC:15130
C70D  C9 55     CMP #$55                        A:55 X:3 Y:55 P:25 SP:F4 CYC:15134
C70F  F0 0A     BEQ $C71B                       A:55 X:3 Y:55 P:27 SP:F4 CYC:15136
C71B  68        PLA                             A:55 X:3 Y:55 P:27 SP:F4 CYC:15139
C71C  29 CB     AND #$CB                        A:34 X:3 Y:55 P:25 SP:F5 CYC:15143
C71E  C9 00     CMP #$00                        A:0 X:3 Y:55 P:27 SP:F5 CYC:15145
C720  F0 06     BEQ $C728                       A:0 X:3 Y:55 P:27 SP:F5 CYC:15147
C728  C8        INY                             A:0 X:3 Y:55 P:27 SP:F5 CYC:15150
C729  CA        DEX                             A:0 X:3 Y:56 P:25 SP:F5 CYC:15152
C72A  D0 E0     BNE $C70C                       A:0 X:2 Y:56 P:25 SP:F5 CYC:15154
C70C  68        PLA                             A:0 X:2 Y:56 P:25 SP:F5 CYC:15157
C70D  C9 55     CMP #$55                        A:55 X:2 Y:56 P:25 SP:F6 CYC:15161
C70F  F0 0A     BEQ $C71B                       A:55 X:2 Y:56 P:27 SP:F6 CYC:15163
C71B  68        PLA                             A:55 X:2 Y:56 P:27 SP:F6 CYC:15166
C71C  29 CB     AND #$CB                        A:34 X:2 Y:56 P:25 SP:F7 CYC:15170
C71E  C9 00     CMP #$00                        A:0 X:2 Y:56 P:27 SP:F7 CYC:15172
C720  F0 06     BEQ $C728                       A:0 X:2 Y:56 P:27 SP:F7 CYC:15174
C728  C8        INY                             A:0 X:2 Y:56 P:27 SP:F7 CYC:15177
C729  CA        DEX                             A:0 X:2 Y:57 P:25 SP:F7 CYC:15179
C72A  D0 E0     BNE $C70C                       A:0 X:1 Y:57 P:25 SP:F7 CYC:15181
C70C  68        PLA                             A:0 X:1 Y:57 P:25 SP:F7 CYC:15184
C70D  C9 55     CMP #$55                        A:55 X:1 Y:57 P:25 SP:F8 CYC:15188
C70F  F0 0A     BEQ $C71B                       A:55 X:1 Y:57 P:27 SP:F8 CYC:15190
C71B  68        PLA                             A:55 X:1 Y:57 P:27 SP:F8 CYC:15193
C71C  29 CB     AND #$CB                        A:34 X:1 Y:57 P:25 SP:F9 CYC:15197
C71E  C9 00     CMP #$00                        A:0 X:1 Y:57 P:27 SP:F9 CYC:15199
C720  F0 06     BEQ $C728                       A:0 X:1 Y:57 P:27 SP:F9 CYC:15201
C728  C8        INY                             A:0 X:1 Y:57 P:27 SP:F9 CYC:15204
C729  CA        DEX                             A:0 X:1 Y:58 P:25 SP:F9 CYC:15206
C72A  D0 E0     BNE $C70C                       A:0 X:0 Y:58 P:27 SP:F9 CYC:15208
C72C  60        RTS                             A:0 X:0 Y:58 P:27 SP:F9 CYC:15210
C6AF  60        RTS                             A:0 X:0 Y:58 P:27 SP:FB CYC:15216
C632  20 1E E5  JSR $E51E                       A:0 X:0 Y:58 P:27 SP:FD CYC:15222
E51E  A9 55     LDA #$55                        A:0 X:0 Y:58 P:27 SP:FB CYC:15228
E520  8D 80 05  STA $0580 = 00                  A:55 X:0 Y:58 P:25 SP:FB CYC:15230
E523  A9 AA     LDA #$AA                        A:55 X:0 Y:58 P:25 SP:FB CYC:15234
E525  8D 32 04  STA $0432 = 00                  A:AA X:0 Y:58 P:A5 SP:FB CYC:15236
E528  A9 80     LDA #$80                        A:AA X:0 Y:58 P:A5 SP:FB CYC:15240
E52A  85 43     STA $43 = 00                    A:80 X:0 Y:58 P:A5 SP:FB CYC:15242
E52C  A9 05     LDA #$05                        A:80 X:0 Y:58 P:A5 SP:FB CYC:15245
E52E  85 44     STA $44 = 00                    A:5 X:0 Y:58 P:25 SP:FB CYC:15247
E530  A9 32     LDA #$32                        A:5 X:0 Y:58 P:25 SP:FB CYC:15250
E532  85 45     STA $45 = 00                    A:32 X:0 Y:58 P:25 SP:FB CYC:15252
E534  A9 04     LDA #$04                        A:32 X:0 Y:58 P:25 SP:FB CYC:15255
E536  85 46     STA $46 = 00                    A:4 X:0 Y:58 P:25 SP:FB CYC:15257
E538  A2 03     LDX #$03                        A:4 X:0 Y:58 P:25 SP:FB CYC:15260
E53A  A0 77     LDY #$77                        A:4 X:3 Y:58 P:25 SP:FB CYC:15262
E53C  A9 FF     LDA #$FF                        A:4 X:3 Y:77 P:25 SP:FB CYC:15264
E53E  85 01     STA $01 = FF                    A:FF X:3 Y:77 P:A5 SP:FB CYC:15266
E540  24 01     BIT $01 = FF                    A:FF X:3 Y:77 P:A5 SP:FB CYC:15269
E542  38        SEC                             A:FF X:3 Y:77 P:E5 SP:FB CYC:15272
E543  A9 00     LDA #$00                        A:FF X:3 Y:77 P:E5 SP:FB CYC:15274
E545  A3 40    *LAX ($40,X) @ 43 = 0580 = 55    A:0 X:3 Y:77 P:67 SP:FB CYC:15276
E547  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB CYC:15282
E548  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB CYC:15284
E549  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB CYC:15286
E54A  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB CYC:15288
E54B  F0 12     BEQ $E55F                       A:55 X:55 Y:77 P:65 SP:FB CYC:15290
E54D  30 10     BMI $E55F                       A:55 X:55 Y:77 P:65 SP:FB CYC:15292
E54F  50 0E     BVC $E55F                       A:55 X:55 Y:77 P:65 SP:FB CYC:15294
E551  90 0C     BCC $E55F                       A:55 X:55 Y:77 P:65 SP:FB CYC:15296
E553  C9 55     CMP #$55                        A:55 X:55 Y:77 P:65 SP:FB CYC:15298
E555  D0 08     BNE $E55F                       A:55 X:55 Y:77 P:67 SP:FB CYC:15300
E557  E0 55     CPX #$55                        A:55 X:55 Y:77 P:67 SP:FB CYC:15302
E559  D0 04     BNE $E55F                       A:55 X:55 Y:77 P:67 SP:FB CYC:15304
E55B  C0 77     CPY #$77                        A:55 X:55 Y:77 P:67 SP:FB CYC:15306
E55D  F0 04     BEQ $E563                       A:55 X:55 Y:77 P:67 SP:FB CYC:15308
E563  A2 05     LDX #$05                        A:55 X:55 Y:77 P:67 SP:FB CYC:15311
E565  A0 33     LDY #$33                        A:55 X:5 Y:77 P:65 SP:FB CYC:15313
E567  B8        CLV                             A:55 X:5 Y:33 P:65 SP:FB CYC:15315
E568  18        CLC                             A:55 X:5 Y:33 P:25 SP:FB CYC:15317
E569  A9 00     LDA #$00                        A:55 X:5 Y:33 P:24 SP:FB CYC:15319
E56B  A3 40    *LAX ($40,X) @ 45 = 0432 = AA    A:0 X:5 Y:33 P:26 SP:FB CYC:15321
E56D  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB CYC:15327
E56E  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB CYC:15329
E56F  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB CYC:15331
E570  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB CYC:15333
E571  F0 12     BEQ $E585                       A:AA X:AA Y:33 P:A4 SP:FB CYC:15335
E573  10 10     BPL $E585                       A:AA X:AA Y:33 P:A4 SP:FB CYC:15337
E575  70 0E     BVS $E585                       A:AA X:AA Y:33 P:A4 SP:FB CYC:15339
E577  B0 0C     BCS $E585                       A:AA X:AA Y:33 P:A4 SP:FB CYC:15341
E579  C9 AA     CMP #$AA                        A:AA X:AA Y:33 P:A4 SP:FB CYC:15343
E57B  D0 08     BNE $E585                       A:AA X:AA Y:33 P:27 SP:FB CYC:15345
E57D  E0 AA     CPX #$AA                        A:AA X:AA Y:33 P:27 SP:FB CYC:15347
E57F  D0 04     BNE $E585                       A:AA X:AA Y:33 P:27 SP:FB CYC:15349
E581  C0 33     CPY #$33                        A:AA X:AA Y:33 P:27 SP:FB CYC:15351
E583  F0 04     BEQ $E589                       A:AA X:AA Y:33 P:27 SP:FB CYC:15353
E589  A9 87     LDA #$87                        A:AA X:AA Y:33 P:27 SP:FB CYC:15356
E58B  85 67     STA $67 = 00                    A:87 X:AA Y:33 P:A5 SP:FB CYC:15358
E58D  A9 32     LDA #$32                        A:87 X:AA Y:33 P:A5 SP:FB CYC:15361
E58F  85 68     STA $68 = 00                    A:32 X:AA Y:33 P:25 SP:FB CYC:15363
E591  A0 57     LDY #$57                        A:32 X:AA Y:33 P:25 SP:FB CYC:15366
E593  24 01     BIT $01 = FF                    A:32 X:AA Y:57 P:25 SP:FB CYC:15368
E595  38        SEC                             A:32 X:AA Y:57 P:E5 SP:FB CYC:15371
E596  A9 00     LDA #$00                        A:32 X:AA Y:57 P:E5 SP:FB CYC:15373
E598  A7 67    *LAX $67 = 87                    A:0 X:AA Y:57 P:67 SP:FB CYC:15375
E59A  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15378
E59B  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15380
E59C  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15382
E59D  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15384
E59E  F0 12     BEQ $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15386
E5A0  10 10     BPL $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15388
E5A2  50 0E     BVC $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15390
E5A4  90 0C     BCC $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15392
E5A6  C9 87     CMP #$87                        A:87 X:87 Y:57 P:E5 SP:FB CYC:15394
E5A8  D0 08     BNE $E5B2                       A:87 X:87 Y:57 P:67 SP:FB CYC:15396
E5AA  E0 87     CPX #$87                        A:87 X:87 Y:57 P:67 SP:FB CYC:15398
E5AC  D0 04     BNE $E5B2                       A:87 X:87 Y:57 P:67 SP:FB CYC:15400
E5AE  C0 57     CPY #$57                        A:87 X:87 Y:57 P:67 SP:FB CYC:15402
E5B0  F0 04     BEQ $E5B6                       A:87 X:87 Y:57 P:67 SP:FB CYC:15404
E5B6  A0 53     LDY #$53                        A:87 X:87 Y:57 P:67 SP:FB CYC:15407
E5B8  B8        CLV                             A:87 X:87 Y:53 P:65 SP:FB CYC:15409
E5B9  18        CLC                             A:87 X:87 Y:53 P:25 SP:FB CYC:15411
E5BA  A9 00     LDA #$00                        A:87 X:87 Y:53 P:24 SP:FB CYC:15413
E5BC  A7 68    *LAX $68 = 32                    A:0 X:87 Y:53 P:26 SP:FB CYC:15415
E5BE  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15418
E5BF  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15420
E5C0  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15422
E5C1  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15424
E5C2  F0 12     BEQ $E5D6                       A:32 X:32 Y:53 P:24 SP:FB CYC:15426
E5C4  30 10     BMI $E5D6                       A:32 X:32 Y:53 P:24 SP:FB CYC:15428
E5C6  70 0E     BVS $E5D6                       A:32 X:32 Y:53 P:24 SP:FB CYC:15430
E5C8  B0 0C     BCS $E5D6                       A:32 X:32 Y:53 P:24 SP:FB CYC:15432
E5CA  C9 32     CMP #$32                        A:32 X:32 Y:53 P:24 SP:FB CYC:15434
E5CC  D0 08     BNE $E5D6                       A:32 X:32 Y:53 P:27 SP:FB CYC:15436
E5CE  E0 32     CPX #$32                        A:32 X:32 Y:53 P:27 SP:FB CYC:15438
E5D0  D0 04     BNE $E5D6                       A:32 X:32 Y:53 P:27 SP:FB CYC:15440
E5D2  C0 53     CPY #$53                        A:32 X:32 Y:53 P:27 SP:FB CYC:15442
E5D4  F0 04     BEQ $E5DA                       A:32 X:32 Y:53 P:27 SP:FB CYC:15444
E5DA  A9 87     LDA #$87                        A:32 X:32 Y:53 P:27 SP:FB CYC:15447
E5DC  8D 77 05  STA $0577 = 00                  A:87 X:32 Y:53 P:A5 SP:FB CYC:15449
E5DF  A9 32     LDA #$32                        A:87 X:32 Y:53 P:A5 SP:FB CYC:15453
E5E1  8D 78 05  STA $0578 = 00                  A:32 X:32 Y:53 P:25 SP:FB CYC:15455
E5E4  A0 57     LDY #$57                        A:32 X:32 Y:53 P:25 SP:FB CYC:15459
E5E6  24 01     BIT $01 = FF                    A:32 X:32 Y:57 P:25 SP:FB CYC:15461
E5E8  38        SEC                             A:32 X:32 Y:57 P:E5 SP:FB CYC:15464
E5E9  A9 00     LDA #$00                        A:32 X:32 Y:57 P:E5 SP:FB CYC:15466
E5EB  AF 77 05 *LAX $0577 = 87                  A:0 X:32 Y:57 P:67 SP:FB CYC:15468
E5EE  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15472
E5EF  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15474
E5F0  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15476
E5F1  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15478
E5F2  F0 12     BEQ $E606                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15480
E5F4  10 10     BPL $E606                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15482
E5F6  50 0E     BVC $E606                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15484
E5F8  90 0C     BCC $E606                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15486
E5FA  C9 87     CMP #$87                        A:87 X:87 Y:57 P:E5 SP:FB CYC:15488
E5FC  D0 08     BNE $E606                       A:87 X:87 Y:57 P:67 SP:FB CYC:15490
E5FE  E0 87     CPX #$87                        A:87 X:87 Y:57 P:67 SP:FB CYC:15492
E600  D0 04     BNE $E606                       A:87 X:87 Y:57 P:67 SP:FB CYC:15494
E602  C0 57     CPY #$57                        A:87 X:87 Y:57 P:67 SP:FB CYC:15496
E604  F0 04     BEQ $E60A                       A:87 X:87 Y:57 P:67 SP:FB CYC:15498
E60A  A0 53     LDY #$53                        A:87 X:87 Y:57 P:67 SP:FB CYC:15501
E60C  B8        CLV                             A:87 X:87 Y:53 P:65 SP:FB CYC:15503
E60D  18        CLC                             A:87 X:87 Y:53 P:25 SP:FB CYC:15505
E60E  A9 00     LDA #$00                        A:87 X:87 Y:53 P:24 SP:FB CYC:15507
E610  AF 78 05 *LAX $0578 = 32                  A:0 X:87 Y:53 P:26 SP:FB CYC:15509
E613  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15513
E614  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15515
E615  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15517
E616  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB CYC:15519
E617  F0 12     BEQ $E62B                       A:32 X:32 Y:53 P:24 SP:FB CYC:15521
E619  30 10     BMI $E62B                       A:32 X:32 Y:53 P:24 SP:FB CYC:15523
E61B  70 0E     BVS $E62B                       A:32 X:32 Y:53 P:24 SP:FB CYC:15525
E61D  B0 0C     BCS $E62B                       A:32 X:32 Y:53 P:24 SP:FB CYC:15527
E61F  C9 32     CMP #$32                        A:32 X:32 Y:53 P:24 SP:FB CYC:15529
E621  D0 08     BNE $E62B                       A:32 X:32 Y:53 P:27 SP:FB CYC:15531
E623  E0 32     CPX #$32                        A:32 X:32 Y:53 P:27 SP:FB CYC:15533
E625  D0 04     BNE $E62B                       A:32 X:32 Y:53 P:27 SP:FB CYC:15535
E627  C0 53     CPY #$53                        A:32 X:32 Y:53 P:27 SP:FB CYC:15537
E629  F0 04     BEQ $E62F                       A:32 X:32 Y:53 P:27 SP:FB CYC:15539
E62F  A9 FF     LDA #$FF                        A:32 X:32 Y:53 P:27 SP:FB CYC:15542
E631  85 43     STA $43 = 80                    A:FF X:32 Y:53 P:A5 SP:FB CYC:15544
E633  A9 04     LDA #$04                        A:FF X:32 Y:53 P:A5 SP:FB CYC:15547
E635  85 44     STA $44 = 05                    A:4 X:32 Y:53 P:25 SP:FB CYC:15549
E637  A9 32     LDA #$32                        A:4 X:32 Y:53 P:25 SP:FB CYC:15552
E639  85 45     STA $45 = 32                    A:32 X:32 Y:53 P:25 SP:FB CYC:15554
E63B  A9 04     LDA #$04                        A:32 X:32 Y:53 P:25 SP:FB CYC:15557
E63D  85 46     STA $46 = 04                    A:4 X:32 Y:53 P:25 SP:FB CYC:15559
E63F  A9 55     LDA #$55                        A:4 X:32 Y:53 P:25 SP:FB CYC:15562
E641  8D 80 05  STA $0580 = 55                  A:55 X:32 Y:53 P:25 SP:FB CYC:15564
E644  A9 AA     LDA #$AA                        A:55 X:32 Y:53 P:25 SP:FB CYC:15568
E646  8D 32 04  STA $0432 = AA                  A:AA X:32 Y:53 P:A5 SP:FB CYC:15570
E649  A2 03     LDX #$03                        A:AA X:32 Y:53 P:A5 SP:FB CYC:15574
E64B  A0 81     LDY #$81                        A:AA X:3 Y:53 P:25 SP:FB CYC:15576
E64D  24 01     BIT $01 = FF                    A:AA X:3 Y:81 P:A5 SP:FB CYC:15578
E64F  38        SEC                             A:AA X:3 Y:81 P:E5 SP:FB CYC:15581
E650  A9 00     LDA #$00                        A:AA X:3 Y:81 P:E5 SP:FB CYC:15583
E652  B3 43    *LAX ($43),Y = 04FF @ 0580 = 55  A:0 X:3 Y:81 P:67 SP:FB CYC:15585
E654  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB CYC:15591
E655  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB CYC:15593
E656  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB CYC:15595
E657  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB CYC:15597
E658  F0 12     BEQ $E66C                       A:55 X:55 Y:81 P:65 SP:FB CYC:15599
E65A  30 10     BMI $E66C                       A:55 X:55 Y:81 P:65 SP:FB CYC:15601
E65C  50 0E     BVC $E66C                       A:55 X:55 Y:81 P:65 SP:FB CYC:15603
E65E  90 0C     BCC $E66C                       A:55 X:55 Y:81 P:65 SP:FB CYC:15605
E660  C9 55     CMP #$55                        A:55 X:55 Y:81 P:65 SP:FB CYC:15607
E662  D0 08     BNE $E66C                       A:55 X:55 Y:81 P:67 SP:FB CYC:15609
E664  E0 55     CPX #$55                        A:55 X:55 Y:81 P:67 SP:FB CYC:15611
E666  D0 04     BNE $E66C                       A:55 X:55 Y:81 P:67 SP:FB CYC:15613
E668  C0 81     CPY #$81                        A:55 X:55 Y:81 P:67 SP:FB CYC:15615
E66A  F0 04     BEQ $E670                       A:55 X:55 Y:81 P:67 SP:FB CYC:15617
E670  A2 05     LDX #$05                        A:55 X:55 Y:81 P:67 SP:FB CYC:15620
E672  A0 00     LDY #$00                        A:55 X:5 Y:81 P:65 SP:FB CYC:15622
E674  B8        CLV                             A:55 X:5 Y:0 P:67 SP:FB CYC:15624
E675  18        CLC                             A:55 X:5 Y:0 P:27 SP:FB CYC:15626
E676  A9 00     LDA #$00                        A:55 X:5 Y:0 P:26 SP:FB CYC:15628
E678  B3 45    *LAX ($45),Y = 0432 @ 0432 = AA  A:0 X:5 Y:0 P:26 SP:FB CYC:15630
E67A  EA        NOP                             A:AA X:AA Y:0 P:A4 SP:FB CYC:15635
E67B  EA        NOP                             A:AA X:AA Y:0 P:A4 SP:FB CYC:15637
E67C  EA        NOP                             A:AA X:AA Y:0 P:A4 SP:FB CYC:15639
E67D  EA        NOP                             A:AA X:AA Y:0 P:A4 SP:FB CYC:15641
E67E  F0 12     BEQ $E692                       A:AA X:AA Y:0 P:A4 SP:FB CYC:15643
E680  10 10     BPL $E692                       A:AA X:AA Y:0 P:A4 SP:FB CYC:15645
E682  70 0E     BVS $E692                       A:AA X:AA Y:0 P:A4 SP:FB CYC:15647
E684  B0 0C     BCS $E692                       A:AA X:AA Y:0 P:A4 SP:FB CYC:15649
E686  C9 AA     CMP #$AA                        A:AA X:AA Y:0 P:A4 SP:FB CYC:15651
E688  D0 08     BNE $E692                       A:AA X:AA Y:0 P:27 SP:FB CYC:15653
E68A  E0 AA     CPX #$AA                        A:AA X:AA Y:0 P:27 SP:FB CYC:15655
E68C  D0 04     BNE $E692                       A:AA X:AA Y:0 P:27 SP:FB CYC:15657
E68E  C0 00     CPY #$00                        A:AA X:AA Y:0 P:27 SP:FB CYC:15659
E690  F0 04     BEQ $E696                       A:AA X:AA Y:0 P:27 SP:FB CYC:15661
E696  A9 87     LDA #$87                        A:AA X:AA Y:0 P:27 SP:FB CYC:15664
E698  85 67     STA $67 = 87                    A:87 X:AA Y:0 P:A5 SP:FB CYC:15666
E69A  A9 32     LDA #$32                        A:87 X:AA Y:0 P:A5 SP:FB CYC:15669
E69C  85 68     STA $68 = 32                    A:32 X:AA Y:0 P:25 SP:FB CYC:15671
E69E  A0 57     LDY #$57                        A:32 X:AA Y:0 P:25 SP:FB CYC:15674
E6A0  24 01     BIT $01 = FF                    A:32 X:AA Y:57 P:25 SP:FB CYC:15676
E6A2  38        SEC                             A:32 X:AA Y:57 P:E5 SP:FB CYC:15679
E6A3  A9 00     LDA #$00                        A:32 X:AA Y:57 P:E5 SP:FB CYC:15681
E6A5  B7 10    *LAX $10,Y @ 67 = 87             A:0 X:AA Y:57 P:67 SP:FB CYC:15683
E6A7  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15687
E6A8  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15689
E6A9  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15691
E6AA  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB CYC:15693
E6AB  F0 12     BEQ $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15695
E6AD  10 10     BPL $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15697
E6AF  50 0E     BVC $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15699
E6B1  90 0C     BCC $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB CYC:15701
E6B3  C9 87     CMP #$87                        A:87 X:87 Y:57 P:E5 SP:FB CYC:15703
E6B5  D0 08     BNE $E6BF                       A:87 X:87 Y:57 P:67 SP:FB CYC:15705
E6B7  E0 87     CPX #$87                        A:87 X:87 Y:57 P:67 SP:FB CYC:15707
E6B9  D0 04     BNE $E6BF                       A:87 X:87 Y:57 P:67 SP:FB CYC:15709
E6BB  C0 57     CPY #$57                        A:87 X:87 Y:57 P:67 SP:FB CYC:15711
E6BD  F0 04     BEQ $E6C3                       A:87 X:87 Y:57 P:67 SP:FB CYC:15713
E6C3  A0 FF     LDY #$FF                        A:87 X:87 Y:57 P:67 SP:FB CYC:15716
E6C5  B8        CLV                             A:87 X:87 Y:FF P:E5 SP:FB CYC:15718
E6C6  18        CLC                             A:87 X:87 Y:FF P:A5 SP:FB CYC:15720
E6C7  A9 00     LDA #$00                        A:87 X:87 Y:FF P:A4 SP:FB CYC:15722
E6C9  B7 69    *LAX $69,Y @ 68 = 32             A:0 X:87 Y:FF P:26 SP:FB CYC:15724
E6CB  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB CYC:15728
E6CC  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB CYC:15730
E6CD  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB CYC:15732
E6CE  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB CYC:15734
E6CF  F0 12     BEQ $E6E3                       A:32 X:32 Y:FF P:24 SP:FB CYC:15736
E6D1  30 10     BMI $E6E3                       A:32 X:32 Y:FF P:24 SP:FB CYC:15738
E6D3  70 0E     BVS $E6E3                       A:32 X:32 Y:FF P:24 SP:FB CYC:15740
E6D5  B0 0C     BCS $E6E3                       A:32 X:32 Y:FF P:24 SP:FB CYC:15742
E6D7  C9 32     CMP #$32                        A:32 X:32 Y:FF P:24 SP:FB CYC:15744
E6D9  D0 08     BNE $E6E3                       A:32 X:32 Y:FF P:27 SP:FB CYC:15746
E6DB  E0 32     CPX #$32                        A:32 X:32 Y:FF P:27 SP:FB CYC:15748
E6DD  D0 04     BNE $E6E3                       A:32 X:32 Y:FF P:27 SP:FB CYC:15750
E6DF  C0 FF     CPY #$FF                        A:32 X:32 Y:FF P:27 SP:FB CYC:15752
E6E1  F0 04     BEQ $E6E7                       A:32 X:32 Y:FF P:27 SP:FB CYC:15754
E6E7  A9 87     LDA #$87                        A:32 X:32 Y:FF P:27 SP:FB CYC:15757
E6E9  8D 87 05  STA $0587 = 00                  A:87 X:32 Y:FF P:A5 SP:FB CYC:15759
E6EC  A9 32     LDA #$32                        A:87 X:32 Y:FF P:A5 SP:FB CYC:15763
E6EE  8D 88 05  STA $0588 = 00                  A:32 X:32 Y:FF P:25 SP:FB CYC:15765
E6F1  A0 30     LDY #$30                        A:32 X:32 Y:FF P:25 SP:FB CYC:15769
E6F3  24 01     BIT $01 = FF                    A:32 X:32 Y:30 P:25 SP:FB CYC:15771
E6F5  38        SEC                             A:32 X:32 Y:30 P:E5 SP:FB CYC:15774
E6F6  A9 00     LDA #$00                        A:32 X:32 Y:30 P:E5 SP:FB CYC:15776
E6F8  BF 57 05 *LAX $0557,Y @ 0587 = 87         A:0 X:32 Y:30 P:67 SP:FB CYC:15778
E6FB  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB CYC:15782
E6FC  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB CYC:15784
E6FD  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB CYC:15786
E6FE  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB CYC:15788
E6FF  F0 12     BEQ $E713                       A:87 X:87 Y:30 P:E5 SP:FB CYC:15790
E701  10 10     BPL $E713                       A:87 X:87 Y:30 P:E5 SP:FB CYC:15792
E703  50 0E     BVC $E713                       A:87 X:87 Y:30 P:E5 SP:FB CYC:15794
E705  90 0C     BCC $E713                       A:87 X:87 Y:30 P:E5 SP:FB CYC:15796
E707  C9 87     CMP #$87                        A:87 X:87 Y:30 P:E5 SP:FB CYC:15798
E709  D0 08     BNE $E713                       A:87 X:87 Y:30 P:67 SP:FB CYC:15800
E70B  E0 87     CPX #$87                        A:87 X:87 Y:30 P:67 SP:FB CYC:15802
E70D  D0 04     BNE $E713                       A:87 X:87 Y:30 P:67 SP:FB CYC:15804
E70F  C0 30     CPY #$30                        A:87 X:87 Y:30 P:67 SP:FB CYC:15806
E711  F0 04     BEQ $E717                       A:87 X:87 Y:30 P:67 SP:FB CYC:15808
E717  A0 40     LDY #$40                        A:87 X:87 Y:30 P:67 SP:FB CYC:15811
E719  B8        CLV                             A:87 X:87 Y:40 P:65 SP:FB CYC:15813
E71A  18        CLC                             A:87 X:87 Y:40 P:25 SP:FB CYC:15815
E71B  A9 00     LDA #$00                        A:87 X:87 Y:40 P:24 SP:FB CYC:15817
E71D  BF 48 05 *LAX $0548,Y @ 0588 = 32         A:0 X:87 Y:40 P:26 SP:FB CYC:15819
E720  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB CYC:15823
E721  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB CYC:15825
E722  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB CYC:15827
E723  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB CYC:15829
E724  F0 12     BEQ $E738                       A:32 X:32 Y:40 P:24 SP:FB CYC:15831
E726  30 10     BMI $E738                       A:32 X:32 Y:40 P:24 SP:FB CYC:15833
E728  70 0E     BVS $E738                       A:32 X:32 Y:40 P:24 SP:FB CYC:15835
E72A  B0 0C     BCS $E738                       A:32 X:32 Y:40 P:24 SP:FB CYC:15837
E72C  C9 32     CMP #$32                        A:32 X:32 Y:40 P:24 SP:FB CYC:15839
E72E  D0 08     BNE $E738                       A:32 X:32 Y:40 P:27 SP:FB CYC:15841
E730  E0 32     CPX #$32                        A:32 X:32 Y:40 P:27 SP:FB CYC:15843
E732  D0 04     BNE $E738                       A:32 X:32 Y:40 P:27 SP:FB CYC:15845
E734  C0 40     CPY #$40                        A:32 X:32 Y:40 P:27 SP:FB CYC:15847
E736  F0 04     BEQ $E73C                       A:32 X:32 Y:40 P:27 SP:FB CYC:15849
E73C  60        RTS                             A:32 X:32 Y:40 P:27 SP:FB CYC:15852
C635  20 3D E7  JSR $E73D                       A:32 X:32 Y:40 P:27 SP:FD CYC:15858
E73D  A9 C0     LDA #$C0                        A:32 X:32 Y:40 P:27 SP:FB CYC:15864
E73F  85 01     STA $01 = FF                    A:C0 X:32 Y:40 P:A5 SP:FB CYC:15866
E741  A9 00     LDA #$00                        A:C0 X:32 Y:40 P:A5 SP:FB CYC:15869
E743  8D 89 04  STA $0489 = 00                  A:0 X:32 Y:40 P:27 SP:FB CYC:15871
E746  A9 89     LDA #$89                        A:0 X:32 Y:40 P:27 SP:FB CYC:15875
E748  85 60     STA $60 = 00                    A:89 X:32 Y:40 P:A5 SP:FB CYC:15877
E74A  A9 04     LDA #$04                        A:89 X:32 Y:40 P:A5 SP:FB CYC:15880
E74C  85 61     STA $61 = 00                    A:4 X:32 Y:40 P:25 SP:FB CYC:15882
E74E  A0 44     LDY #$44                        A:4 X:32 Y:40 P:25 SP:FB CYC:15885
E750  A2 17     LDX #$17                        A:4 X:32 Y:44 P:25 SP:FB CYC:15887
E752  A9 3E     LDA #$3E                        A:4 X:17 Y:44 P:25 SP:FB CYC:15889
E754  24 01     BIT $01 = C0                    A:3E X:17 Y:44 P:25 SP:FB CYC:15891
E756  18        CLC                             A:3E X:17 Y:44 P:E7 SP:FB CYC:15894
E757  83 49    *SAX ($49,X) @ 60 = 0489 = 00    A:3E X:17 Y:44 P:E6 SP:FB CYC:15896
E759  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB CYC:15902
E75A  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB CYC:15904
E75B  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB CYC:15906
E75C  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB CYC:15908
E75D  D0 19     BNE $E778                       A:3E X:17 Y:44 P:E6 SP:FB CYC:15910
E75F  B0 17     BCS $E778                       A:3E X:17 Y:44 P:E6 SP:FB CYC:15912
E761  50 15     BVC $E778                       A:3E X:17 Y:44 P:E6 SP:FB CYC:15914
E763  10 13     BPL $E778                       A:3E X:17 Y:44 P:E6 SP:FB CYC:15916
E765  C9 3E     CMP #$3E                        A:3E X:17 Y:44 P:E6 SP:FB CYC:15918
E767  D0 0F     BNE $E778                       A:3E X:17 Y:44 P:67 SP:FB CYC:15920
E769  C0 44     CPY #$44                        A:3E X:17 Y:44 P:67 SP:FB CYC:15922
E76B  D0 0B     BNE $E778                       A:3E X:17 Y:44 P:67 SP:FB CYC:15924
E76D  E0 17     CPX #$17                        A:3E X:17 Y:44 P:67 SP:FB CYC:15926
E76F  D0 07     BNE $E778                       A:3E X:17 Y:44 P:67 SP:FB CYC:15928
E771  AD 89 04  LDA $0489 = 16                  A:3E X:17 Y:44 P:67 SP:FB CYC:15930
E774  C9 16     CMP #$16                        A:16 X:17 Y:44 P:65 SP:FB CYC:15934
E776  F0 04     BEQ $E77C                       A:16 X:17 Y:44 P:67 SP:FB CYC:15936
E77C  A0 44     LDY #$44                        A:16 X:17 Y:44 P:67 SP:FB CYC:15939
E77E  A2 7A     LDX #$7A                        A:16 X:17 Y:44 P:65 SP:FB CYC:15941
E780  A9 66     LDA #$66                        A:16 X:7A Y:44 P:65 SP:FB CYC:15943
E782  38        SEC                             A:66 X:7A Y:44 P:65 SP:FB CYC:15945
E783  B8        CLV                             A:66 X:7A Y:44 P:65 SP:FB CYC:15947
E784  83 E6    *SAX ($E6,X) @ 60 = 0489 = 16    A:66 X:7A Y:44 P:25 SP:FB CYC:15949
E786  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB CYC:15955
E787  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB CYC:15957
E788  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB CYC:15959
E789  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB CYC:15961
E78A  F0 19     BEQ $E7A5                       A:66 X:7A Y:44 P:25 SP:FB CYC:15963
E78C  90 17     BCC $E7A5                       A:66 X:7A Y:44 P:25 SP:FB CYC:15965
E78E  70 15     BVS $E7A5                       A:66 X:7A Y:44 P:25 SP:FB CYC:15967
E790  30 13     BMI $E7A5                       A:66 X:7A Y:44 P:25 SP:FB CYC:15969
E792  C9 66     CMP #$66                        A:66 X:7A Y:44 P:25 SP:FB CYC:15971
E794  D0 0F     BNE $E7A5                       A:66 X:7A Y:44 P:27 SP:FB CYC:15973
E796  C0 44     CPY #$44                        A:66 X:7A Y:44 P:27 SP:FB CYC:15975
E798  D0 0B     BNE $E7A5                       A:66 X:7A Y:44 P:27 SP:FB CYC:15977
E79A  E0 7A     CPX #$7A                        A:66 X:7A Y:44 P:27 SP:FB CYC:15979
E79C  D0 07     BNE $E7A5                       A:66 X:7A Y:44 P:27 SP:FB CYC:15981
E79E  AD 89 04  LDA $0489 = 62                  A:66 X:7A Y:44 P:27 SP:FB CYC:15983
E7A1  C9 62     CMP #$62                        A:62 X:7A Y:44 P:25 SP:FB CYC:15987
E7A3  F0 04     BEQ $E7A9                       A:62 X:7A Y:44 P:27 SP:FB CYC:15989
E7A9  A9 FF     LDA #$FF                        A:62 X:7A Y:44 P:27 SP:FB CYC:15992
E7AB  85 49     STA $49 = 00                    A:FF X:7A Y:44 P:A5 SP:FB CYC:15994
E7AD  A0 44     LDY #$44                        A:FF X:7A Y:44 P:A5 SP:FB CYC:15997
E7AF  A2 AA     LDX #$AA                        A:FF X:7A Y:44 P:25 SP:FB CYC:15999
E7B1  A9 55     LDA #$55                        A:FF X:AA Y:44 P:A5 SP:FB CYC:16001
E7B3  24 01     BIT $01 = C0                    A:55 X:AA Y:44 P:25 SP:FB CYC:16003
E7B5  18        CLC                             A:55 X:AA Y:44 P:E5 SP:FB CYC:16006
E7B6  87 49    *SAX $49 = FF                    A:55 X:AA Y:44 P:E4 SP:FB CYC:16008
E7B8  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB CYC:16011
E7B9  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB CYC:16013
E7BA  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB CYC:16015
E7BB  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB CYC:16017
E7BC  F0 18     BEQ $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB CYC:16019
E7BE  B0 16     BCS $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB CYC:16021
E7C0  50 14     BVC $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB CYC:16023
E7C2  10 12     BPL $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB CYC:16025
E7C4  C9 55     CMP #$55                        A:55 X:AA Y:44 P:E4 SP:FB CYC:16027
E7C6  D0 0E     BNE $E7D6                       A:55 X:AA Y:44 P:67 SP:FB CYC:16029
E7C8  C0 44     CPY #$44                        A:55 X:AA Y:44 P:67 SP:FB CYC:16031
E7CA  D0 0A     BNE $E7D6                       A:55 X:AA Y:44 P:67 SP:FB CYC:16033
E7CC  E0 AA     CPX #$AA                        A:55 X:AA Y:44 P:67 SP:FB CYC:16035
E7CE  D0 06     BNE $E7D6                       A:55 X:AA Y:44 P:67 SP:FB CYC:16037
E7D0  A5 49     LDA $49 = 00                    A:55 X:AA Y:44 P:67 SP:FB CYC:16039
E7D2  C9 00     CMP #$00                        A:0 X:AA Y:44 P:67 SP:FB CYC:16042
E7D4  F0 04     BEQ $E7DA                       A:0 X:AA Y:44 P:67 SP:FB CYC:16044
E7DA  A9 00     LDA #$00                        A:0 X:AA Y:44 P:67 SP:FB CYC:16047
E7DC  85 56     STA $56 = 00                    A:0 X:AA Y:44 P:67 SP:FB CYC:16049
E7DE  A0 58     LDY #$58                        A:0 X:AA Y:44 P:67 SP:FB CYC:16052
E7E0  A2 EF     LDX #$EF                        A:0 X:AA Y:58 P:65 SP:FB CYC:16054
E7E2  A9 66     LDA #$66                        A:0 X:EF Y:58 P:E5 SP:FB CYC:16056
E7E4  38        SEC                             A:66 X:EF Y:58 P:65 SP:FB CYC:16058
E7E5  B8        CLV                             A:66 X:EF Y:58 P:65 SP:FB CYC:16060
E7E6  87 56    *SAX $56 = 00                    A:66 X:EF Y:58 P:25 SP:FB CYC:16062
E7E8  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB CYC:16065
E7E9  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB CYC:16067
E7EA  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB CYC:16069
E7EB  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB CYC:16071
E7EC  F0 18     BEQ $E806                       A:66 X:EF Y:58 P:25 SP:FB CYC:16073
E7EE  90 16     BCC $E806                       A:66 X:EF Y:58 P:25 SP:FB CYC:16075
E7F0  70 14     BVS $E806                       A:66 X:EF Y:58 P:25 SP:FB CYC:16077
E7F2  30 12     BMI $E806                       A:66 X:EF Y:58 P:25 SP:FB CYC:16079
E7F4  C9 66     CMP #$66                        A:66 X:EF Y:58 P:25 SP:FB CYC:16081
E7F6  D0 0E     BNE $E806                       A:66 X:EF Y:58 P:27 SP:FB CYC:16083
E7F8  C0 58     CPY #$58                        A:66 X:EF Y:58 P:27 SP:FB CYC:16085
E7FA  D0 0A     BNE $E806                       A:66 X:EF Y:58 P:27 SP:FB CYC:16087
E7FC  E0 EF     CPX #$EF                        A:66 X:EF Y:58 P:27 SP:FB CYC:16089
E7FE  D0 06     BNE $E806                       A:66 X:EF Y:58 P:27 SP:FB CYC:16091
E800  A5 56     LDA $56 = 66                    A:66 X:EF Y:58 P:27 SP:FB CYC:16093
E802  C9 66     CMP #$66                        A:66 X:EF Y:58 P:25 SP:FB CYC:16096
E804  F0 04     BEQ $E80A                       A:66 X:EF Y:58 P:27 SP:FB CYC:16098
E80A  A9 FF     LDA #$FF                        A:66 X:EF Y:58 P:27 SP:FB CYC:16101
E80C  8D 49 05  STA $0549 = 00                  A:FF X:EF Y:58 P:A5 SP:FB CYC:16103
E80F  A0 E5     LDY #$E5                        A:FF X:EF Y:58 P:A5 SP:FB CYC:16107
E811  A2 AF     LDX #$AF                        A:FF X:EF Y:E5 P:A5 SP:FB CYC:16109
E813  A9 F5     LDA #$F5                        A:FF X:AF Y:E5 P:A5 SP:FB CYC:16111
E815  24 01     BIT $01 = C0                    A:F5 X:AF Y:E5 P:A5 SP:FB CYC:16113
E817  18        CLC                             A:F5 X:AF Y:E5 P:E5 SP:FB CYC:16116
E818  8F 49 05 *SAX $0549 = FF                  A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16118
E81B  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16122
E81C  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16124
E81D  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16126
E81E  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16128
E81F  F0 19     BEQ $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16130
E821  B0 17     BCS $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16132
E823  50 15     BVC $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16134
E825  10 13     BPL $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16136
E827  C9 F5     CMP #$F5                        A:F5 X:AF Y:E5 P:E4 SP:FB CYC:16138
E829  D0 0F     BNE $E83A                       A:F5 X:AF Y:E5 P:67 SP:FB CYC:16140
E82B  C0 E5     CPY #$E5                        A:F5 X:AF Y:E5 P:67 SP:FB CYC:16142
E82D  D0 0B     BNE $E83A                       A:F5 X:AF Y:E5 P:67 SP:FB CYC:16144
E82F  E0 AF     CPX #$AF                        A:F5 X:AF Y:E5 P:67 SP:FB CYC:16146
E831  D0 07     BNE $E83A                       A:F5 X:AF Y:E5 P:67 SP:FB CYC:16148
E833  AD 49 05  LDA $0549 = A5                  A:F5 X:AF Y:E5 P:67 SP:FB CYC:16150
E836  C9 A5     CMP #$A5                        A:A5 X:AF Y:E5 P:E5 SP:FB CYC:16154
E838  F0 04     BEQ $E83E                       A:A5 X:AF Y:E5 P:67 SP:FB CYC:16156
E83E  A9 00     LDA #$00                        A:A5 X:AF Y:E5 P:67 SP:FB CYC:16159
E840  8D 56 05  STA $0556 = 00                  A:0 X:AF Y:E5 P:67 SP:FB CYC:16161
E843  A0 58     LDY #$58                        A:0 X:AF Y:E5 P:67 SP:FB CYC:16165
E845  A2 B3     LDX #$B3                        A:0 X:AF Y:58 P:65 SP:FB CYC:16167
E847  A9 97     LDA #$97                        A:0 X:B3 Y:58 P:E5 SP:FB CYC:16169
E849  38        SEC                             A:97 X:B3 Y:58 P:E5 SP:FB CYC:16171
E84A  B8        CLV                             A:97 X:B3 Y:58 P:E5 SP:FB CYC:16173
E84B  8F 56 05 *SAX $0556 = 00                  A:97 X:B3 Y:58 P:A5 SP:FB CYC:16175
E84E  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB CYC:16179
E84F  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB CYC:16181
E850  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB CYC:16183
E851  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB CYC:16185
E852  F0 19     BEQ $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB CYC:16187
E854  90 17     BCC $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB CYC:16189
E856  70 15     BVS $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB CYC:16191
E858  10 13     BPL $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB CYC:16193
E85A  C9 97     CMP #$97                        A:97 X:B3 Y:58 P:A5 SP:FB CYC:16195
E85C  D0 0F     BNE $E86D                       A:97 X:B3 Y:58 P:27 SP:FB CYC:16197
E85E  C0 58     CPY #$58                        A:97 X:B3 Y:58 P:27 SP:FB CYC:16199
E860  D0 0B     BNE $E86D                       A:97 X:B3 Y:58 P:27 SP:FB CYC:16201
E862  E0 B3     CPX #$B3                        A:97 X:B3 Y:58 P:27 SP:FB CYC:16203
E864  D0 07     BNE $E86D                       A:97 X:B3 Y:58 P:27 SP:FB CYC:16205
E866  AD 56 05  LDA $0556 = 93                  A:97 X:B3 Y:58 P:27 SP:FB CYC:16207
E869  C9 93     CMP #$93                        A:93 X:B3 Y:58 P:A5 SP:FB CYC:16211
E86B  F0 04     BEQ $E871                       A:93 X:B3 Y:58 P:27 SP:FB CYC:16213
E871  A9 FF     LDA #$FF                        A:93 X:B3 Y:58 P:27 SP:FB CYC:16216
E873  85 49     STA $49 = 00                    A:FF X:B3 Y:58 P:A5 SP:FB CYC:16218
E875  A0 FF     LDY #$FF                        A:FF X:B3 Y:58 P:A5 SP:FB CYC:16221
E877  A2 AA     LDX #$AA                        A:FF X:B3 Y:FF P:A5 SP:FB CYC:16223
E879  A9 55     LDA #$55                        A:FF X:AA Y:FF P:A5 SP:FB CYC:16225
E87B  24 01     BIT $01 = C0                    A:55 X:AA Y:FF P:25 SP:FB CYC:16227
E87D  18        CLC                             A:55 X:AA Y:FF P:E5 SP:FB CYC:16230
E87E  97 4A    *SAX $4A,Y @ 49 = FF             A:55 X:AA Y:FF P:E4 SP:FB CYC:16232
E880  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB CYC:16236
E881  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB CYC:16238
E882  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB CYC:16240
E883  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB CYC:16242
E884  F0 18     BEQ $E89E                       A:55 X:AA Y:FF P:E4 SP:FB CYC:16244
E886  B0 16     BCS $E89E                       A:55 X:AA Y:FF P:E4 SP:FB CYC:16246
E888  50 14     BVC $E89E                       A:55 X:AA Y:FF P:E4 SP:FB CYC:16248
E88A  10 12     BPL $E89E                       A:55 X:AA Y:FF P:E4 SP:FB CYC:16250
E88C  C9 55     CMP #$55                        A:55 X:AA Y:FF P:E4 SP:FB CYC:16252
E88E  D0 0E     BNE $E89E                       A:55 X:AA Y:FF P:67 SP:FB CYC:16254
E890  C0 FF     CPY #$FF                        A:55 X:AA Y:FF P:67 SP:FB CYC:16256
E892  D0 0A     BNE $E89E                       A:55 X:AA Y:FF P:67 SP:FB CYC:16258
E894  E0 AA     CPX #$AA                        A:55 X:AA Y:FF P:67 SP:FB CYC:16260
E896  D0 06     BNE $E89E                       A:55 X:AA Y:FF P:67 SP:FB CYC:16262
E898  A5 49     LDA $49 = 00                    A:55 X:AA Y:FF P:67 SP:FB CYC:16264
E89A  C9 00     CMP #$00                        A:0 X:AA Y:FF P:67 SP:FB CYC:16267
E89C  F0 04     BEQ $E8A2                       A:0 X:AA Y:FF P:67 SP:FB CYC:16269
E8A2  A9 00     LDA #$00                        A:0 X:AA Y:FF P:67 SP:FB CYC:16272
E8A4  85 56     STA $56 = 66                    A:0 X:AA Y:FF P:67 SP:FB CYC:16274
E8A6  A0 06     LDY #$06                        A:0 X:AA Y:FF P:67 SP:FB CYC:16277
E8A8  A2 EF     LDX #$EF                        A:0 X:AA Y:6 P:65 SP:FB CYC:16279
E8AA  A9 66     LDA #$66                        A:0 X:EF Y:6 P:E5 SP:FB CYC:16281
E8AC  38        SEC                             A:66 X:EF Y:6 P:65 SP:FB CYC:16283
E8AD  B8        CLV                             A:66 X:EF Y:6 P:65 SP:FB CYC:16285
E8AE  97 50    *SAX $50,Y @ 56 = 00             A:66 X:EF Y:6 P:25 SP:FB CYC:16287
E8B0  EA        NOP                             A:66 X:EF Y:6 P:25 SP:FB CYC:16291
E8B1  EA        NOP                             A:66 X:EF Y:6 P:25 SP:FB CYC:16293
E8B2  EA        NOP                             A:66 X:EF Y:6 P:25 SP:FB CYC:16295
E8B3  EA        NOP                             A:66 X:EF Y:6 P:25 SP:FB CYC:16297
E8B4  F0 18     BEQ $E8CE                       A:66 X:EF Y:6 P:25 SP:FB CYC:16299
E8B6  90 16     BCC $E8CE                       A:66 X:EF Y:6 P:25 SP:FB CYC:16301
E8B8  70 14     BVS $E8CE                       A:66 X:EF Y:6 P:25 SP:FB CYC:16303
E8BA  30 12     BMI $E8CE                       A:66 X:EF Y:6 P:25 SP:FB CYC:16305
E8BC  C9 66     CMP #$66                        A:66 X:EF Y:6 P:25 SP:FB CYC:16307
E8BE  D0 0E     BNE $E8CE                       A:66 X:EF Y:6 P:27 SP:FB CYC:16309
E8C0  C0 06     CPY #$06                        A:66 X:EF Y:6 P:27 SP:FB CYC:16311
E8C2  D0 0A     BNE $E8CE                       A:66 X:EF Y:6 P:27 SP:FB CYC:16313
E8C4  E0 EF     CPX #$EF                        A:66 X:EF Y:6 P:27 SP:FB CYC:16315
E8C6  D0 06     BNE $E8CE                       A:66 X:EF Y:6 P:27 SP:FB CYC:16317
E8C8  A5 56     LDA $56 = 66                    A:66 X:EF Y:6 P:27 SP:FB CYC:16319
E8CA  C9 66     CMP #$66                        A:66 X:EF Y:6 P:25 SP:FB CYC:16322
E8CC  F0 04     BEQ $E8D2                       A:66 X:EF Y:6 P:27 SP:FB CYC:16324
E8D2  60        RTS                             A:66 X:EF Y:6 P:27 SP:FB CYC:16327
C638  20 D3 E8  JSR $E8D3                       A:66 X:EF Y:6 P:27 SP:FD CYC:16333
E8D3  A0 90     LDY #$90                        A:66 X:EF Y:6 P:27 SP:FB CYC:16339
E8D5  20 31 F9  JSR $F931                       A:66 X:EF Y:90 P:A5 SP:FB CYC:16341
F931  24 01     BIT $01 = C0                    A:66 X:EF Y:90 P:A5 SP:F9 CYC:16347
F933  A9 40     LDA #$40                        A:66 X:EF Y:90 P:E5 SP:F9 CYC:16350
F935  38        SEC                             A:40 X:EF Y:90 P:65 SP:F9 CYC:16352
F936  60        RTS                             A:40 X:EF Y:90 P:65 SP:F9 CYC:16354
E8D8  EB 40    *SBC #$40                        A:40 X:EF Y:90 P:65 SP:FB CYC:16360
E8DA  EA        NOP                             A:0 X:EF Y:90 P:27 SP:FB CYC:16362
E8DB  EA        NOP                             A:0 X:EF Y:90 P:27 SP:FB CYC:16364
E8DC  EA        NOP                             A:0 X:EF Y:90 P:27 SP:FB CYC:16366
E8DD  EA        NOP                             A:0 X:EF Y:90 P:27 SP:FB CYC:16368
E8DE  20 37 F9  JSR $F937                       A:0 X:EF Y:90 P:27 SP:FB CYC:16370
F937  30 0B     BMI $F944                       A:0 X:EF Y:90 P:27 SP:F9 CYC:16376
F939  90 09     BCC $F944                       A:0 X:EF Y:90 P:27 SP:F9 CYC:16378
F93B  D0 07     BNE $F944                       A:0 X:EF Y:90 P:27 SP:F9 CYC:16380
F93D  70 05     BVS $F944                       A:0 X:EF Y:90 P:27 SP:F9 CYC:16382
F93F  C9 00     CMP #$00                        A:0 X:EF Y:90 P:27 SP:F9 CYC:16384
F941  D0 01     BNE $F944                       A:0 X:EF Y:90 P:27 SP:F9 CYC:16386
F943  60        RTS                             A:0 X:EF Y:90 P:27 SP:F9 CYC:16388
E8E1  C8        INY                             A:0 X:EF Y:90 P:27 SP:FB CYC:16394
E8E2  20 47 F9  JSR $F947                       A:0 X:EF Y:91 P:A5 SP:FB CYC:16396
F947  B8        CLV                             A:0 X:EF Y:91 P:A5 SP:F9 CYC:16402
F948  38        SEC                             A:0 X:EF Y:91 P:A5 SP:F9 CYC:16404
F949  A9 40     LDA #$40                        A:0 X:EF Y:91 P:A5 SP:F9 CYC:16406
F94B  60        RTS                             A:40 X:EF Y:91 P:25 SP:F9 CYC:16408
E8E5  EB 3F    *SBC #$3F                        A:40 X:EF Y:91 P:25 SP:FB CYC:16414
E8E7  EA        NOP                             A:1 X:EF Y:91 P:25 SP:FB CYC:16416
E8E8  EA        NOP                             A:1 X:EF Y:91 P:25 SP:FB CYC:16418
E8E9  EA        NOP                             A:1 X:EF Y:91 P:25 SP:FB CYC:16420
E8EA  EA        NOP                             A:1 X:EF Y:91 P:25 SP:FB CYC:16422
E8EB  20 4C F9  JSR $F94C                       A:1 X:EF Y:91 P:25 SP:FB CYC:16424
F94C  F0 0B     BEQ $F959                       A:1 X:EF Y:91 P:25 SP:F9 CYC:16430
F94E  30 09     BMI $F959                       A:1 X:EF Y:91 P:25 SP:F9 CYC:16432
F950  90 07     BCC $F959                       A:1 X:EF Y:91 P:25 SP:F9 CYC:16434
F952  70 05     BVS $F959                       A:1 X:EF Y:91 P:25 SP:F9 CYC:16436
F954  C9 01     CMP #$01                        A:1 X:EF Y:91 P:25 SP:F9 CYC:16438
F956  D0 01     BNE $F959                       A:1 X:EF Y:91 P:27 SP:F9 CYC:16440
F958  60        RTS                             A:1 X:EF Y:91 P:27 SP:F9 CYC:16442
E8EE  C8        INY                             A:1 X:EF Y:91 P:27 SP:FB CYC:16448
E8EF  20 5C F9  JSR $F95C                       A:1 X:EF Y:92 P:A5 SP:FB CYC:16450
F95C  A9 40     LDA #$40                        A:1 X:EF Y:92 P:A5 SP:F9 CYC:16456
F95E  38        SEC                             A:40 X:EF Y:92 P:25 SP:F9 CYC:16458
F95F  24 01     BIT $01 = C0                    A:40 X:EF Y:92 P:25 SP:F9 CYC:16460
F961  60        RTS                             A:40 X:EF Y:92 P:E5 SP:F9 CYC:16463
E8F2  EB 41    *SBC #$41                        A:40 X:EF Y:92 P:E5 SP:FB CYC:16469
E8F4  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB CYC:16471
E8F5  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB CYC:16473
E8F6  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB CYC:16475
E8F7  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB CYC:16477
E8F8  20 62 F9  JSR $F962                       A:FF X:EF Y:92 P:A4 SP:FB CYC:16479
F962  B0 0B     BCS $F96F                       A:FF X:EF Y:92 P:A4 SP:F9 CYC:16485
F964  F0 09     BEQ $F96F                       A:FF X:EF Y:92 P:A4 SP:F9 CYC:16487
F966  10 07     BPL $F96F                       A:FF X:EF Y:92 P:A4 SP:F9 CYC:16489
F968  70 05     BVS $F96F                       A:FF X:EF Y:92 P:A4 SP:F9 CYC:16491
F96A  C9 FF     CMP #$FF                        A:FF X:EF Y:92 P:A4 SP:F9 CYC:16493
F96C  D0 01     BNE $F96F                       A:FF X:EF Y:92 P:27 SP:F9 CYC:16495
F96E  60        RTS                             A:FF X:EF Y:92 P:27 SP:F9 CYC:16497
E8FB  C8        INY                             A:FF X:EF Y:92 P:27 SP:FB CYC:16503
E8FC  20 72 F9  JSR $F972                       A:FF X:EF Y:93 P:A5 SP:FB CYC:16505
F972  18        CLC                             A:FF X:EF Y:93 P:A5 SP:F9 CYC:16511
F973  A9 80     LDA #$80                        A:FF X:EF Y:93 P:A4 SP:F9 CYC:16513
F975  60        RTS                             A:80 X:EF Y:93 P:A4 SP:F9 CYC:16515
E8FF  EB 00    *SBC #$00                        A:80 X:EF Y:93 P:A4 SP:FB CYC:16521
E901  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB CYC:16523
E902  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB CYC:16525
E903  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB CYC:16527
E904  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB CYC:16529
E905  20 76 F9  JSR $F976                       A:7F X:EF Y:93 P:65 SP:FB CYC:16531
F976  90 05     BCC $F97D                       A:7F X:EF Y:93 P:65 SP:F9 CYC:16537
F978  C9 7F     CMP #$7F                        A:7F X:EF Y:93 P:65 SP:F9 CYC:16539
F97A  D0 01     BNE $F97D                       A:7F X:EF Y:93 P:67 SP:F9 CYC:16541
F97C  60        RTS                             A:7F X:EF Y:93 P:67 SP:F9 CYC:16543
E908  C8        INY                             A:7F X:EF Y:93 P:67 SP:FB CYC:16549
E909  20 80 F9  JSR $F980                       A:7F X:EF Y:94 P:E5 SP:FB CYC:16551
F980  38        SEC                             A:7F X:EF Y:94 P:E5 SP:F9 CYC:16557
F981  A9 81     LDA #$81                        A:7F X:EF Y:94 P:E5 SP:F9 CYC:16559
F983  60        RTS                             A:81 X:EF Y:94 P:E5 SP:F9 CYC:16561
E90C  EB 7F    *SBC #$7F                        A:81 X:EF Y:94 P:E5 SP:FB CYC:16567
E90E  EA        NOP                             A:2 X:EF Y:94 P:65 SP:FB CYC:16569
E90F  EA        NOP                             A:2 X:EF Y:94 P:65 SP:FB CYC:16571
E910  EA        NOP                             A:2 X:EF Y:94 P:65 SP:FB CYC:16573
E911  EA        NOP                             A:2 X:EF Y:94 P:65 SP:FB CYC:16575
E912  20 84 F9  JSR $F984                       A:2 X:EF Y:94 P:65 SP:FB CYC:16577
F984  50 07     BVC $F98D                       A:2 X:EF Y:94 P:65 SP:F9 CYC:16583
F986  90 05     BCC $F98D                       A:2 X:EF Y:94 P:65 SP:F9 CYC:16585
F988  C9 02     CMP #$02                        A:2 X:EF Y:94 P:65 SP:F9 CYC:16587
F98A  D0 01     BNE $F98D                       A:2 X:EF Y:94 P:67 SP:F9 CYC:16589
F98C  60        RTS                             A:2 X:EF Y:94 P:67 SP:F9 CYC:16591
E915  60        RTS                             A:2 X:EF Y:94 P:67 SP:FB CYC:16597
C63B  20 16 E9  JSR $E916                       A:2 X:EF Y:94 P:67 SP:FD CYC:16603
E916  A9 FF     LDA #$FF                        A:2 X:EF Y:94 P:67 SP:FB CYC:16609
E918  85 01     STA $01 = C0                    A:FF X:EF Y:94 P:E5 SP:FB CYC:16611
E91A  A0 95     LDY #$95                        A:FF X:EF Y:94 P:E5 SP:FB CYC:16614
E91C  A2 02     LDX #$02                        A:FF X:EF Y:95 P:E5 SP:FB CYC:16616
E91E  A9 47     LDA #$47                        A:FF X:2 Y:95 P:65 SP:FB CYC:16618
E920  85 47     STA $47 = 00                    A:47 X:2 Y:95 P:65 SP:FB CYC:16620
E922  A9 06     LDA #$06                        A:47 X:2 Y:95 P:65 SP:FB CYC:16623
E924  85 48     STA $48 = 00                    A:6 X:2 Y:95 P:65 SP:FB CYC:16625
E926  A9 EB     LDA #$EB                        A:6 X:2 Y:95 P:65 SP:FB CYC:16628
E928  8D 47 06  STA $0647 = 00                  A:EB X:2 Y:95 P:E5 SP:FB CYC:16630
E92B  20 31 FA  JSR $FA31                       A:EB X:2 Y:95 P:E5 SP:FB CYC:16634
FA31  24 01     BIT $01 = FF                    A:EB X:2 Y:95 P:E5 SP:F9 CYC:16640
FA33  18        CLC                             A:EB X:2 Y:95 P:E5 SP:F9 CYC:16643
FA34  A9 40     LDA #$40                        A:EB X:2 Y:95 P:E4 SP:F9 CYC:16645
FA36  60        RTS                             A:40 X:2 Y:95 P:64 SP:F9 CYC:16647
E92E  C3 45    *DCP ($45,X) @ 47 = 0647 = EB    A:40 X:2 Y:95 P:64 SP:FB CYC:16653
E930  EA        NOP                             A:40 X:2 Y:95 P:64 SP:FB CYC:16661
E931  EA        NOP                             A:40 X:2 Y:95 P:64 SP:FB CYC:16663
E932  EA        NOP                             A:40 X:2 Y:95 P:64 SP:FB CYC:16665
E933  EA        NOP                             A:40 X:2 Y:95 P:64 SP:FB CYC:16667
E934  20 37 FA  JSR $FA37                       A:40 X:2 Y:95 P:64 SP:FB CYC:16669
FA37  50 2C     BVC $FA65                       A:40 X:2 Y:95 P:64 SP:F9 CYC:16675
FA39  B0 2A     BCS $FA65                       A:40 X:2 Y:95 P:64 SP:F9 CYC:16677
FA3B  30 28     BMI $FA65                       A:40 X:2 Y:95 P:64 SP:F9 CYC:16679
FA3D  C9 40     CMP #$40                        A:40 X:2 Y:95 P:64 SP:F9 CYC:16681
FA3F  D0 24     BNE $FA65                       A:40 X:2 Y:95 P:67 SP:F9 CYC:16683
FA41  60        RTS                             A:40 X:2 Y:95 P:67 SP:F9 CYC:16685
E937  AD 47 06  LDA $0647 = EA                  A:40 X:2 Y:95 P:67 SP:FB CYC:16691
E93A  C9 EA     CMP #$EA                        A:EA X:2 Y:95 P:E5 SP:FB CYC:16695
E93C  F0 02     BEQ $E940                       A:EA X:2 Y:95 P:67 SP:FB CYC:16697
E940  C8        INY                             A:EA X:2 Y:95 P:67 SP:FB CYC:16700
E941  A9 00     LDA #$00                        A:EA X:2 Y:96 P:E5 SP:FB CYC:16702
E943  8D 47 06  STA $0647 = EA                  A:0 X:2 Y:96 P:67 SP:FB CYC:16704
E946  20 42 FA  JSR $FA42                       A:0 X:2 Y:96 P:67 SP:FB CYC:16708
FA42  B8        CLV                             A:0 X:2 Y:96 P:67 SP:F9 CYC:16714
FA43  38        SEC                             A:0 X:2 Y:96 P:27 SP:F9 CYC:16716
FA44  A9 FF     LDA #$FF                        A:0 X:2 Y:96 P:27 SP:F9 CYC:16718
FA46  60        RTS                             A:FF X:2 Y:96 P:A5 SP:F9 CYC:16720
E949  C3 45    *DCP ($45,X) @ 47 = 0647 = 00    A:FF X:2 Y:96 P:A5 SP:FB CYC:16726
E94B  EA        NOP                             A:FF X:2 Y:96 P:27 SP:FB CYC:16734
E94C  EA        NOP                             A:FF X:2 Y:96 P:27 SP:FB CYC:16736
E94D  EA        NOP                             A:FF X:2 Y:96 P:27 SP:FB CYC:16738
E94E  EA        NOP                             A:FF X:2 Y:96 P:27 SP:FB CYC:16740
E94F  20 47 FA  JSR $FA47                       A:FF X:2 Y:96 P:27 SP:FB CYC:16742
FA47  70 1C     BVS $FA65                       A:FF X:2 Y:96 P:27 SP:F9 CYC:16748
FA49  D0 1A     BNE $FA65                       A:FF X:2 Y:96 P:27 SP:F9 CYC:16750
FA4B  30 18     BMI $FA65                       A:FF X:2 Y:96 P:27 SP:F9 CYC:16752
FA4D  90 16     BCC $FA65                       A:FF X:2 Y:96 P:27 SP:F9 CYC:16754
FA4F  C9 FF     CMP #$FF                        A:FF X:2 Y:96 P:27 SP:F9 CYC:16756
FA51  D0 12     BNE $FA65                       A:FF X:2 Y:96 P:27 SP:F9 CYC:16758
FA53  60        RTS                             A:FF X:2 Y:96 P:27 SP:F9 CYC:16760
E952  AD 47 06  LDA $0647 = FF                  A:FF X:2 Y:96 P:27 SP:FB CYC:16766
E955  C9 FF     CMP #$FF                        A:FF X:2 Y:96 P:A5 SP:FB CYC:16770
E957  F0 02     BEQ $E95B                       A:FF X:2 Y:96 P:27 SP:FB CYC:16772
E95B  C8        INY                             A:FF X:2 Y:96 P:27 SP:FB CYC:16775
E95C  A9 37     LDA #$37                        A:FF X:2 Y:97 P:A5 SP:FB CYC:16777
E95E  8D 47 06  STA $0647 = FF                  A:37 X:2 Y:97 P:25 SP:FB CYC:16779
E961  20 54 FA  JSR $FA54                       A:37 X:2 Y:97 P:25 SP:FB CYC:16783
FA54  24 01     BIT $01 = FF                    A:37 X:2 Y:97 P:25 SP:F9 CYC:16789
FA56  A9 F0     LDA #$F0                        A:37 X:2 Y:97 P:E5 SP:F9 CYC:16792
FA58  60        RTS                             A:F0 X:2 Y:97 P:E5 SP:F9 CYC:16794
E964  C3 45    *DCP ($45,X) @ 47 = 0647 = 37    A:F0 X:2 Y:97 P:E5 SP:FB CYC:16800
E966  EA        NOP                             A:F0 X:2 Y:97 P:E5 SP:FB CYC:16808
E967  EA        NOP                             A:F0 X:2 Y:97 P:E5 SP:FB CYC:16810
E968  EA        NOP                             A:F0 X:2 Y:97 P:E5 SP:FB CYC:16812
E969  EA        NOP                             A:F0 X:2 Y:97 P:E5 SP:FB CYC:16814
E96A  20 59 FA  JSR $FA59                       A:F0 X:2 Y:97 P:E5 SP:FB CYC:16816
FA59  50 0A     BVC $FA65                       A:F0 X:2 Y:97 P:E5 SP:F9 CYC:16822
FA5B  F0 08     BEQ $FA65                       A:F0 X:2 Y:97 P:E5 SP:F9 CYC:16824
FA5D  10 06     BPL $FA65                       A:F0 X:2 Y:97 P:E5 SP:F9 CYC:16826
FA5F  90 04     BCC $FA65                       A:F0 X:2 Y:97 P:E5 SP:F9 CYC:16828
FA61  C9 F0     CMP #$F0                        A:F0 X:2 Y:97 P:E5 SP:F9 CYC:16830
FA63  F0 02     BEQ $FA67                       A:F0 X:2 Y:97 P:67 SP:F9 CYC:16832
FA67  60        RTS                             A:F0 X:2 Y:97 P:67 SP:F9 CYC:16835
E96D  AD 47 06  LDA $0647 = 36                  A:F0 X:2 Y:97 P:67 SP:FB CYC:16841
E970  C9 36     CMP #$36                        A:36 X:2 Y:97 P:65 SP:FB CYC:16845
E972  F0 02     BEQ $E976                       A:36 X:2 Y:97 P:67 SP:FB CYC:16847
E976  C8        INY                             A:36 X:2 Y:97 P:67 SP:FB CYC:16850
E977  A9 EB     LDA #$EB                        A:36 X:2 Y:98 P:E5 SP:FB CYC:16852
E979  85 47     STA $47 = 47                    A:EB X:2 Y:98 P:E5 SP:FB CYC:16854
E97B  20 31 FA  JSR $FA31                       A:EB X:2 Y:98 P:E5 SP:FB CYC:16857
FA31  24 01     BIT $01 = FF                    A:EB X:2 Y:98 P:E5 SP:F9 CYC:16863
FA33  18        CLC                             A:EB X:2 Y:98 P:E5 SP:F9 CYC:16866
FA34  A9 40     LDA #$40                        A:EB X:2 Y:98 P:E4 SP:F9 CYC:16868
FA36  60        RTS                             A:40 X:2 Y:98 P:64 SP:F9 CYC:16870
E97E  C7 47    *DCP $47 = EB                    A:40 X:2 Y:98 P:64 SP:FB CYC:16876
E980  EA        NOP                             A:40 X:2 Y:98 P:64 SP:FB CYC:16881
E981  EA        NOP                             A:40 X:2 Y:98 P:64 SP:FB CYC:16883
E982  EA        NOP                             A:40 X:2 Y:98 P:64 SP:FB CYC:16885
E983  EA        NOP                             A:40 X:2 Y:98 P:64 SP:FB CYC:16887
E984  20 37 FA  JSR $FA37                       A:40 X:2 Y:98 P:64 SP:FB CYC:16889
FA37  50 2C     BVC $FA65                       A:40 X:2 Y:98 P:64 SP:F9 CYC:16895
FA39  B0 2A     BCS $FA65                       A:40 X:2 Y:98 P:64 SP:F9 CYC:16897
FA3B  30 28     BMI $FA65                       A:40 X:2 Y:98 P:64 SP:F9 CYC:16899
FA3D  C9 40     CMP #$40                        A:40 X:2 Y:98 P:64 SP:F9 CYC:16901
FA3F  D0 24     BNE $FA65                       A:40 X:2 Y:98 P:67 SP:F9 CYC:16903
FA41  60        RTS                             A:40 X:2 Y:98 P:67 SP:F9 CYC:16905
E987  A5 47     LDA $47 = EA                    A:40 X:2 Y:98 P:67 SP:FB CYC:16911
E989  C9 EA     CMP #$EA                        A:EA X:2 Y:98 P:E5 SP:FB CYC:16914
E98B  F0 02     BEQ $E98F                       A:EA X:2 Y:98 P:67 SP:FB CYC:16916
E98F  C8        INY                             A:EA X:2 Y:98 P:67 SP:FB CYC:16919
E990  A9 00     LDA #$00                        A:EA X:2 Y:99 P:E5 SP:FB CYC:16921
E992  85 47     STA $47 = EA                    A:0 X:2 Y:99 P:67 SP:FB CYC:16923
E994  20 42 FA  JSR $FA42                       A:0 X:2 Y:99 P:67 SP:FB CYC:16926
FA42  B8        CLV                             A:0 X:2 Y:99 P:67 SP:F9 CYC:16932
FA43  38        SEC                             A:0 X:2 Y:99 P:27 SP:F9 CYC:16934
FA44  A9 FF     LDA #$FF                        A:0 X:2 Y:99 P:27 SP:F9 CYC:16936
FA46  60        RTS                             A:FF X:2 Y:99 P:A5 SP:F9 CYC:16938
E997  C7 47    *DCP $47 = 00                    A:FF X:2 Y:99 P:A5 SP:FB CYC:16944
E999  EA        NOP                             A:FF X:2 Y:99 P:27 SP:FB CYC:16949
E99A  EA        NOP                             A:FF X:2 Y:99 P:27 SP:FB CYC:16951
E99B  EA        NOP                             A:FF X:2 Y:99 P:27 SP:FB CYC:16953
E99C  EA        NOP                             A:FF X:2 Y:99 P:27 SP:FB CYC:16955
E99D  20 47 FA  JSR $FA47                       A:FF X:2 Y:99 P:27 SP:FB CYC:16957
FA47  70 1C     BVS $FA65                       A:FF X:2 Y:99 P:27 SP:F9 CYC:16963
FA49  D0 1A     BNE $FA65                       A:FF X:2 Y:99 P:27 SP:F9 CYC:16965
FA4B  30 18     BMI $FA65                       A:FF X:2 Y:99 P:27 SP:F9 CYC:16967
FA4D  90 16     BCC $FA65                       A:FF X:2 Y:99 P:27 SP:F9 CYC:16969
FA4F  C9 FF     CMP #$FF                        A:FF X:2 Y:99 P:27 SP:F9 CYC:16971
FA51  D0 12     BNE $FA65                       A:FF X:2 Y:99 P:27 SP:F9 CYC:16973
FA53  60        RTS                             A:FF X:2 Y:99 P:27 SP:F9 CYC:16975
E9A0  A5 47     LDA $47 = FF                    A:FF X:2 Y:99 P:27 SP:FB CYC:16981
E9A2  C9 FF     CMP #$FF                        A:FF X:2 Y:99 P:A5 SP:FB CYC:16984
E9A4  F0 02     BEQ $E9A8                       A:FF X:2 Y:99 P:27 SP:FB CYC:16986
E9A8  C8        INY                             A:FF X:2 Y:99 P:27 SP:FB CYC:16989
E9A9  A9 37     LDA #$37                        A:FF X:2 Y:9A P:A5 SP:FB CYC:16991
E9AB  85 47     STA $47 = FF                    A:37 X:2 Y:9A P:25 SP:FB CYC:16993
E9AD  20 54 FA  JSR $FA54                       A:37 X:2 Y:9A P:25 SP:FB CYC:16996
FA54  24 01     BIT $01 = FF                    A:37 X:2 Y:9A P:25 SP:F9 CYC:17002
FA56  A9 F0     LDA #$F0                        A:37 X:2 Y:9A P:E5 SP:F9 CYC:17005
FA58  60        RTS                             A:F0 X:2 Y:9A P:E5 SP:F9 CYC:17007
E9B0  C7 47    *DCP $47 = 37                    A:F0 X:2 Y:9A P:E5 SP:FB CYC:17013
E9B2  EA        NOP                             A:F0 X:2 Y:9A P:E5 SP:FB CYC:17018
E9B3  EA        NOP                             A:F0 X:2 Y:9A P:E5 SP:FB CYC:17020
E9B4  EA        NOP                             A:F0 X:2 Y:9A P:E5 SP:FB CYC:17022
E9B5  EA        NOP                             A:F0 X:2 Y:9A P:E5 SP:FB CYC:17024
E9B6  20 59 FA  JSR $FA59                       A:F0 X:2 Y:9A P:E5 SP:FB CYC:17026
FA59  50 0A     BVC $FA65                       A:F0 X:2 Y:9A P:E5 SP:F9 CYC:17032
FA5B  F0 08     BEQ $FA65                       A:F0 X:2 Y:9A P:E5 SP:F9 CYC:17034
FA5D  10 06     BPL $FA65                       A:F0 X:2 Y:9A P:E5 SP:F9 CYC:17036
FA5F  90 04     BCC $FA65                       A:F0 X:2 Y:9A P:E5 SP:F9 CYC:17038
FA61  C9 F0     CMP #$F0                        A:F0 X:2 Y:9A P:E5 SP:F9 CYC:17040
FA63  F0 02     BEQ $FA67                       A:F0 X:2 Y:9A P:67 SP:F9 CYC:17042
FA67  60        RTS                             A:F0 X:2 Y:9A P:67 SP:F9 CYC:17045
E9B9  A5 47     LDA $47 = 36                    A:F0 X:2 Y:9A P:67 SP:FB CYC:17051
E9BB  C9 36     CMP #$36                        A:36 X:2 Y:9A P:65 SP:FB CYC:17054
E9BD  F0 02     BEQ $E9C1                       A:36 X:2 Y:9A P:67 SP:FB CYC:17056
E9C1  C8        INY                             A:36 X:2 Y:9A P:67 SP:FB CYC:17059
E9C2  A9 EB     LDA #$EB                        A:36 X:2 Y:9B P:E5 SP:FB CYC:17061
E9C4  8D 47 06  STA $0647 = 36                  A:EB X:2 Y:9B P:E5 SP:FB CYC:17063
E9C7  20 31 FA  JSR $FA31                       A:EB X:2 Y:9B P:E5 SP:FB CYC:17067
FA31  24 01     BIT $01 = FF                    A:EB X:2 Y:9B P:E5 SP:F9 CYC:17073
FA33  18        CLC                             A:EB X:2 Y:9B P:E5 SP:F9 CYC:17076
FA34  A9 40     LDA #$40                        A:EB X:2 Y:9B P:E4 SP:F9 CYC:17078
FA36  60        RTS                             A:40 X:2 Y:9B P:64 SP:F9 CYC:17080
E9CA  CF 47 06 *DCP $0647 = EB                  A:40 X:2 Y:9B P:64 SP:FB CYC:17086
E9CD  EA        NOP                             A:40 X:2 Y:9B P:64 SP:FB CYC:17092
E9CE  EA        NOP                             A:40 X:2 Y:9B P:64 SP:FB CYC:17094
E9CF  EA        NOP                             A:40 X:2 Y:9B P:64 SP:FB CYC:17096
E9D0  EA        NOP                             A:40 X:2 Y:9B P:64 SP:FB CYC:17098
E9D1  20 37 FA  JSR $FA37                       A:40 X:2 Y:9B P:64 SP:FB CYC:17100
FA37  50 2C     BVC $FA65                       A:40 X:2 Y:9B P:64 SP:F9 CYC:17106
FA39  B0 2A     BCS $FA65                       A:40 X:2 Y:9B P:64 SP:F9 CYC:17108
FA3B  30 28     BMI $FA65                       A:40 X:2 Y:9B P:64 SP:F9 CYC:17110
FA3D  C9 40     CMP #$40                        A:40 X:2 Y:9B P:64 SP:F9 CYC:17112
FA3F  D0 24     BNE $FA65                       A:40 X:2 Y:9B P:67 SP:F9 CYC:17114
FA41  60        RTS                             A:40 X:2 Y:9B P:67 SP:F9 CYC:17116
E9D4  AD 47 06  LDA $0647 = EA                  A:40 X:2 Y:9B P:67 SP:FB CYC:17122
E9D7  C9 EA     CMP #$EA                        A:EA X:2 Y:9B P:E5 SP:FB CYC:17126
E9D9  F0 02     BEQ $E9DD                       A:EA X:2 Y:9B P:67 SP:FB CYC:17128
E9DD  C8        INY                             A:EA X:2 Y:9B P:67 SP:FB CYC:17131
E9DE  A9 00     LDA #$00                        A:EA X:2 Y:9C P:E5 SP:FB CYC:17133
E9E0  8D 47 06  STA $0647 = EA                  A:0 X:2 Y:9C P:67 SP:FB CYC:17135
E9E3  20 42 FA  JSR $FA42                       A:0 X:2 Y:9C P:67 SP:FB CYC:17139
FA42  B8        CLV                             A:0 X:2 Y:9C P:67 SP:F9 CYC:17145
FA43  38        SEC                             A:0 X:2 Y:9C P:27 SP:F9 CYC:17147
FA44  A9 FF     LDA #$FF                        A:0 X:2 Y:9C P:27 SP:F9 CYC:17149
FA46  60        RTS                             A:FF X:2 Y:9C P:A5 SP:F9 CYC:17151
E9E6  CF 47 06 *DCP $0647 = 00                  A:FF X:2 Y:9C P:A5 SP:FB CYC:17157
E9E9  EA        NOP                             A:FF X:2 Y:9C P:27 SP:FB CYC:17163
E9EA  EA        NOP                             A:FF X:2 Y:9C P:27 SP:FB CYC:17165
E9EB  EA        NOP                             A:FF X:2 Y:9C P:27 SP:FB CYC:17167
E9EC  EA        NOP                             A:FF X:2 Y:9C P:27 SP:FB CYC:17169
E9ED  20 47 FA  JSR $FA47                       A:FF X:2 Y:9C P:27 SP:FB CYC:17171
FA47  70 1C     BVS $FA65                       A:FF X:2 Y:9C P:27 SP:F9 CYC:17177
FA49  D0 1A     BNE $FA65                       A:FF X:2 Y:9C P:27 SP:F9 CYC:17179
FA4B  30 18     BMI $FA65                       A:FF X:2 Y:9C P:27 SP:F9 CYC:17181
FA4D  90 16     BCC $FA65                       A:FF X:2 Y:9C P:27 SP:F9 CYC:17183
FA4F  C9 FF     CMP #$FF                        A:FF X:2 Y:9C P:27 SP:F9 CYC:17185
FA51  D0 12     BNE $FA65                       A:FF X:2 Y:9C P:27 SP:F9 CYC:17187
FA53  60        RTS                             A:FF X:2 Y:9C P:27 SP:F9 CYC:17189
E9F0  AD 47 06  LDA $0647 = FF                  A:FF X:2 Y:9C P:27 SP:FB CYC:17195
E9F3  C9 FF     CMP #$FF                        A:FF X:2 Y:9C P:A5 SP:FB CYC:17199
E9F5  F0 02     BEQ $E9F9                       A:FF X:2 Y:9C P:27 SP:FB CYC:17201
E9F9  C8        INY                             A:FF X:2 Y:9C P:27 SP:FB CYC:17204
E9FA  A9 37     LDA #$37                        A:FF X:2 Y:9D P:A5 SP:FB CYC:17206
E9FC  8D 47 06  STA $0647 = FF                  A:37 X:2 Y:9D P:25 SP:FB CYC:17208
E9FF  20 54 FA  JSR $FA54                       A:37 X:2 Y:9D P:25 SP:FB CYC:17212
FA54  24 01     BIT $01 = FF                    A:37 X:2 Y:9D P:25 SP:F9 CYC:17218
FA56  A9 F0     LDA #$F0                        A:37 X:2 Y:9D P:E5 SP:F9 CYC:17221
FA58  60        RTS                             A:F0 X:2 Y:9D P:E5 SP:F9 CYC:17223
EA02  CF 47 06 *DCP $0647 = 37                  A:F0 X:2 Y:9D P:E5 SP:FB CYC:17229
EA05  EA        NOP                             A:F0 X:2 Y:9D P:E5 SP:FB CYC:17235
EA06  EA        NOP                             A:F0 X:2 Y:9D P:E5 SP:FB CYC:17237
EA07  EA        NOP                             A:F0 X:2 Y:9D P:E5 SP:FB CYC:17239
EA08  EA        NOP                             A:F0 X:2 Y:9D P:E5 SP:FB CYC:17241
EA09  20 59 FA  JSR $FA59                       A:F0 X:2 Y:9D P:E5 SP:FB CYC:17243
FA59  50 0A     BVC $FA65                       A:F0 X:2 Y:9D P:E5 SP:F9 CYC:17249
FA5B  F0 08     BEQ $FA65                       A:F0 X:2 Y:9D P:E5 SP:F9 CYC:17251
FA5D  10 06     BPL $FA65                       A:F0 X:2 Y:9D P:E5 SP:F9 CYC:17253
FA5F  90 04     BCC $FA65                       A:F0 X:2 Y:9D P:E5 SP:F9 CYC:17255
FA61  C9 F0     CMP #$F0                        A:F0 X:2 Y:9D P:E5 SP:F9 CYC:17257
FA63  F0 02     BEQ $FA67                       A:F0 X:2 Y:9D P:67 SP:F9 CYC:17259
FA67  60        RTS                             A:F0 X:2 Y:9D P:67 SP:F9 CYC:17262
EA0C  AD 47 06  LDA $0647 = 36                  A:F0 X:2 Y:9D P:67 SP:FB CYC:17268
EA0F  C9 36     CMP #$36                        A:36 X:2 Y:9D P:65 SP:FB CYC:17272
EA11  F0 02     BEQ $EA15                       A:36 X:2 Y:9D P:67 SP:FB CYC:17274
EA15  A9 EB     LDA #$EB                        A:36 X:2 Y:9D P:67 SP:FB CYC:17277
EA17  8D 47 06  STA $0647 = 36                  A:EB X:2 Y:9D P:E5 SP:FB CYC:17279
EA1A  A9 48     LDA #$48                        A:EB X:2 Y:9D P:E5 SP:FB CYC:17283
EA1C  85 45     STA $45 = 32                    A:48 X:2 Y:9D P:65 SP:FB CYC:17285
EA1E  A9 05     LDA #$05                        A:48 X:2 Y:9D P:65 SP:FB CYC:17288
EA20  85 46     STA $46 = 04                    A:5 X:2 Y:9D P:65 SP:FB CYC:17290
EA22  A0 FF     LDY #$FF                        A:5 X:2 Y:9D P:65 SP:FB CYC:17293
EA24  20 31 FA  JSR $FA31                       A:5 X:2 Y:FF P:E5 SP:FB CYC:17295
FA31  24 01     BIT $01 = FF                    A:5 X:2 Y:FF P:E5 SP:F9 CYC:17301
FA33  18        CLC                             A:5 X:2 Y:FF P:E5 SP:F9 CYC:17304
FA34  A9 40     LDA #$40                        A:5 X:2 Y:FF P:E4 SP:F9 CYC:17306
FA36  60        RTS                             A:40 X:2 Y:FF P:64 SP:F9 CYC:17308
EA27  D3 45    *DCP ($45),Y = 0548 @ 0647 = EB  A:40 X:2 Y:FF P:64 SP:FB CYC:17314
EA29  EA        NOP                             A:40 X:2 Y:FF P:64 SP:FB CYC:17322
EA2A  EA        NOP                             A:40 X:2 Y:FF P:64 SP:FB CYC:17324
EA2B  08        PHP                             A:40 X:2 Y:FF P:64 SP:FB CYC:17326
EA2C  48        PHA                             A:40 X:2 Y:FF P:64 SP:FA CYC:17329
EA2D  A0 9E     LDY #$9E                        A:40 X:2 Y:FF P:64 SP:F9 CYC:17332
EA2F  68        PLA                             A:40 X:2 Y:9E P:E4 SP:F9 CYC:17334
EA30  28        PLP                             A:40 X:2 Y:9E P:64 SP:FA CYC:17338
EA31  20 37 FA  JSR $FA37                       A:40 X:2 Y:9E P:64 SP:FB CYC:17342
FA37  50 2C     BVC $FA65                       A:40 X:2 Y:9E P:64 SP:F9 CYC:17348
FA39  B0 2A     BCS $FA65                       A:40 X:2 Y:9E P:64 SP:F9 CYC:17350
FA3B  30 28     BMI $FA65                       A:40 X:2 Y:9E P:64 SP:F9 CYC:17352
FA3D  C9 40     CMP #$40                        A:40 X:2 Y:9E P:64 SP:F9 CYC:17354
FA3F  D0 24     BNE $FA65                       A:40 X:2 Y:9E P:67 SP:F9 CYC:17356
FA41  60        RTS                             A:40 X:2 Y:9E P:67 SP:F9 CYC:17358
EA34  AD 47 06  LDA $0647 = EA                  A:40 X:2 Y:9E P:67 SP:FB CYC:17364
EA37  C9 EA     CMP #$EA                        A:EA X:2 Y:9E P:E5 SP:FB CYC:17368
EA39  F0 02     BEQ $EA3D                       A:EA X:2 Y:9E P:67 SP:FB CYC:17370
EA3D  A0 FF     LDY #$FF                        A:EA X:2 Y:9E P:67 SP:FB CYC:17373
EA3F  A9 00     LDA #$00                        A:EA X:2 Y:FF P:E5 SP:FB CYC:17375
EA41  8D 47 06  STA $0647 = EA                  A:0 X:2 Y:FF P:67 SP:FB CYC:17377
EA44  20 42 FA  JSR $FA42                       A:0 X:2 Y:FF P:67 SP:FB CYC:17381
FA42  B8        CLV                             A:0 X:2 Y:FF P:67 SP:F9 CYC:17387
FA43  38        SEC                             A:0 X:2 Y:FF P:27 SP:F9 CYC:17389
FA44  A9 FF     LDA #$FF                        A:0 X:2 Y:FF P:27 SP:F9 CYC:17391
FA46  60        RTS                             A:FF X:2 Y:FF P:A5 SP:F9 CYC:17393
EA47  D3 45    *DCP ($45),Y = 0548 @ 0647 = 00  A:FF X:2 Y:FF P:A5 SP:FB CYC:17399
EA49  EA        NOP                             A:FF X:2 Y:FF P:27 SP:FB CYC:17407
EA4A  EA        NOP                             A:FF X:2 Y:FF P:27 SP:FB CYC:17409
EA4B  08        PHP                             A:FF X:2 Y:FF P:27 SP:FB CYC:17411
EA4C  48        PHA                             A:FF X:2 Y:FF P:27 SP:FA CYC:17414
EA4D  A0 9F     LDY #$9F                        A:FF X:2 Y:FF P:27 SP:F9 CYC:17417
EA4F  68        PLA                             A:FF X:2 Y:9F P:A5 SP:F9 CYC:17419
EA50  28        PLP                             A:FF X:2 Y:9F P:A5 SP:FA CYC:17423
EA51  20 47 FA  JSR $FA47                       A:FF X:2 Y:9F P:27 SP:FB CYC:17427
FA47  70 1C     BVS $FA65                       A:FF X:2 Y:9F P:27 SP:F9 CYC:17433
FA49  D0 1A     BNE $FA65                       A:FF X:2 Y:9F P:27 SP:F9 CYC:17435
FA4B  30 18     BMI $FA65                       A:FF X:2 Y:9F P:27 SP:F9 CYC:17437
FA4D  90 16     BCC $FA65                       A:FF X:2 Y:9F P:27 SP:F9 CYC:17439
FA4F  C9 FF     CMP #$FF                        A:FF X:2 Y:9F P:27 SP:F9 CYC:17441
FA51  D0 12     BNE $FA65                       A:FF X:2 Y:9F P:27 SP:F9 CYC:17443
FA53  60        RTS                             A:FF X:2 Y:9F P:27 SP:F9 CYC:17445
EA54  AD 47 06  LDA $0647 = FF                  A:FF X:2 Y:9F P:27 SP:FB CYC:17451
EA57  C9 FF     CMP #$FF                        A:FF X:2 Y:9F P:A5 SP:FB CYC:17455
EA59  F0 02     BEQ $EA5D                       A:FF X:2 Y:9F P:27 SP:FB CYC:17457
EA5D  A0 FF     LDY #$FF                        A:FF X:2 Y:9F P:27 SP:FB CYC:17460
EA5F  A9 37     LDA #$37                        A:FF X:2 Y:FF P:A5 SP:FB CYC:17462
EA61  8D 47 06  STA $0647 = FF                  A:37 X:2 Y:FF P:25 SP:FB CYC:17464
EA64  20 54 FA  JSR $FA54                       A:37 X:2 Y:FF P:25 SP:FB CYC:17468
FA54  24 01     BIT $01 = FF                    A:37 X:2 Y:FF P:25 SP:F9 CYC:17474
FA56  A9 F0     LDA #$F0                        A:37 X:2 Y:FF P:E5 SP:F9 CYC:17477
FA58  60        RTS                             A:F0 X:2 Y:FF P:E5 SP:F9 CYC:17479
EA67  D3 45    *DCP ($45),Y = 0548 @ 0647 = 37  A:F0 X:2 Y:FF P:E5 SP:FB CYC:17485
EA69  EA        NOP                             A:F0 X:2 Y:FF P:E5 SP:FB CYC:17493
EA6A  EA        NOP                             A:F0 X:2 Y:FF P:E5 SP:FB CYC:17495
EA6B  08        PHP                             A:F0 X:2 Y:FF P:E5 SP:FB CYC:17497
EA6C  48        PHA                             A:F0 X:2 Y:FF P:E5 SP:FA CYC:17500
EA6D  A0 A0     LDY #$A0                        A:F0 X:2 Y:FF P:E5 SP:F9 CYC:17503
EA6F  68        PLA                             A:F0 X:2 Y:A0 P:E5 SP:F9 CYC:17505
EA70  28        PLP                             A:F0 X:2 Y:A0 P:E5 SP:FA CYC:17509
EA71  20 59 FA  JSR $FA59                       A:F0 X:2 Y:A0 P:E5 SP:FB CYC:17513
FA59  50 0A     BVC $FA65                       A:F0 X:2 Y:A0 P:E5 SP:F9 CYC:17519
FA5B  F0 08     BEQ $FA65                       A:F0 X:2 Y:A0 P:E5 SP:F9 CYC:17521
FA5D  10 06     BPL $FA65                       A:F0 X:2 Y:A0 P:E5 SP:F9 CYC:17523
FA5F  90 04     BCC $FA65                       A:F0 X:2 Y:A0 P:E5 SP:F9 CYC:17525
FA61  C9 F0     CMP #$F0                        A:F0 X:2 Y:A0 P:E5 SP:F9 CYC:17527
FA63  F0 02     BEQ $FA67                       A:F0 X:2 Y:A0 P:67 SP:F9 CYC:17529
FA67  60        RTS                             A:F0 X:2 Y:A0 P:67 SP:F9 CYC:17532
EA74  AD 47 06  LDA $0647 = 36                  A:F0 X:2 Y:A0 P:67 SP:FB CYC:17538
EA77  C9 36     CMP #$36                        A:36 X:2 Y:A0 P:65 SP:FB CYC:17542
EA79  F0 02     BEQ $EA7D                       A:36 X:2 Y:A0 P:67 SP:FB CYC:17544
EA7D  A0 A1     LDY #$A1                        A:36 X:2 Y:A0 P:67 SP:FB CYC:17547
EA7F  A2 FF     LDX #$FF                        A:36 X:2 Y:A1 P:E5 SP:FB CYC:17549
EA81  A9 EB     LDA #$EB                        A:36 X:FF Y:A1 P:E5 SP:FB CYC:17551
EA83  85 47     STA $47 = 36                    A:EB X:FF Y:A1 P:E5 SP:FB CYC:17553
EA85  20 31 FA  JSR $FA31                       A:EB X:FF Y:A1 P:E5 SP:FB CYC:17556
FA31  24 01     BIT $01 = FF                    A:EB X:FF Y:A1 P:E5 SP:F9 CYC:17562
FA33  18        CLC                             A:EB X:FF Y:A1 P:E5 SP:F9 CYC:17565
FA34  A9 40     LDA #$40                        A:EB X:FF Y:A1 P:E4 SP:F9 CYC:17567
FA36  60        RTS                             A:40 X:FF Y:A1 P:64 SP:F9 CYC:17569
EA88  D7 48    *DCP $48,X @ 47 = EB             A:40 X:FF Y:A1 P:64 SP:FB CYC:17575
EA8A  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB CYC:17581
EA8B  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB CYC:17583
EA8C  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB CYC:17585
EA8D  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB CYC:17587
EA8E  20 37 FA  JSR $FA37                       A:40 X:FF Y:A1 P:64 SP:FB CYC:17589
FA37  50 2C     BVC $FA65                       A:40 X:FF Y:A1 P:64 SP:F9 CYC:17595
FA39  B0 2A     BCS $FA65                       A:40 X:FF Y:A1 P:64 SP:F9 CYC:17597
FA3B  30 28     BMI $FA65                       A:40 X:FF Y:A1 P:64 SP:F9 CYC:17599
FA3D  C9 40     CMP #$40                        A:40 X:FF Y:A1 P:64 SP:F9 CYC:17601
FA3F  D0 24     BNE $FA65                       A:40 X:FF Y:A1 P:67 SP:F9 CYC:17603
FA41  60        RTS                             A:40 X:FF Y:A1 P:67 SP:F9 CYC:17605
EA91  A5 47     LDA $47 = EA                    A:40 X:FF Y:A1 P:67 SP:FB CYC:17611
EA93  C9 EA     CMP #$EA                        A:EA X:FF Y:A1 P:E5 SP:FB CYC:17614
EA95  F0 02     BEQ $EA99                       A:EA X:FF Y:A1 P:67 SP:FB CYC:17616
EA99  C8        INY                             A:EA X:FF Y:A1 P:67 SP:FB CYC:17619
EA9A  A9 00     LDA #$00                        A:EA X:FF Y:A2 P:E5 SP:FB CYC:17621
EA9C  85 47     STA $47 = EA                    A:0 X:FF Y:A2 P:67 SP:FB CYC:17623
EA9E  20 42 FA  JSR $FA42                       A:0 X:FF Y:A2 P:67 SP:FB CYC:17626
FA42  B8        CLV                             A:0 X:FF Y:A2 P:67 SP:F9 CYC:17632
FA43  38        SEC                             A:0 X:FF Y:A2 P:27 SP:F9 CYC:17634
FA44  A9 FF     LDA #$FF                        A:0 X:FF Y:A2 P:27 SP:F9 CYC:17636
FA46  60        RTS                             A:FF X:FF Y:A2 P:A5 SP:F9 CYC:17638
EAA1  D7 48    *DCP $48,X @ 47 = 00             A:FF X:FF Y:A2 P:A5 SP:FB CYC:17644
EAA3  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB CYC:17650
EAA4  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB CYC:17652
EAA5  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB CYC:17654
EAA6  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB CYC:17656
EAA7  20 47 FA  JSR $FA47                       A:FF X:FF Y:A2 P:27 SP:FB CYC:17658
FA47  70 1C     BVS $FA65                       A:FF X:FF Y:A2 P:27 SP:F9 CYC:17664
FA49  D0 1A     BNE $FA65                       A:FF X:FF Y:A2 P:27 SP:F9 CYC:17666
FA4B  30 18     BMI $FA65                       A:FF X:FF Y:A2 P:27 SP:F9 CYC:17668
FA4D  90 16     BCC $FA65                       A:FF X:FF Y:A2 P:27 SP:F9 CYC:17670
FA4F  C9 FF     CMP #$FF                        A:FF X:FF Y:A2 P:27 SP:F9 CYC:17672
FA51  D0 12     BNE $FA65                       A:FF X:FF Y:A2 P:27 SP:F9 CYC:17674
FA53  60        RTS                             A:FF X:FF Y:A2 P:27 SP:F9 CYC:17676
EAAA  A5 47     LDA $47 = FF                    A:FF X:FF Y:A2 P:27 SP:FB CYC:17682
EAAC  C9 FF     CMP #$FF                        A:FF X:FF Y:A2 P:A5 SP:FB CYC:17685
EAAE  F0 02     BEQ $EAB2                       A:FF X:FF Y:A2 P:27 SP:FB CYC:17687
EAB2  C8        INY                             A:FF X:FF Y:A2 P:27 SP:FB CYC:17690
EAB3  A9 37     LDA #$37                        A:FF X:FF Y:A3 P:A5 SP:FB CYC:17692
EAB5  85 47     STA $47 = FF                    A:37 X:FF Y:A3 P:25 SP:FB CYC:17694
EAB7  20 54 FA  JSR $FA54                       A:37 X:FF Y:A3 P:25 SP:FB CYC:17697
FA54  24 01     BIT $01 = FF                    A:37 X:FF Y:A3 P:25 SP:F9 CYC:17703
FA56  A9 F0     LDA #$F0                        A:37 X:FF Y:A3 P:E5 SP:F9 CYC:17706
FA58  60        RTS                             A:F0 X:FF Y:A3 P:E5 SP:F9 CYC:17708
EABA  D7 48    *DCP $48,X @ 47 = 37             A:F0 X:FF Y:A3 P:E5 SP:FB CYC:17714
EABC  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB CYC:17720
EABD  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB CYC:17722
EABE  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB CYC:17724
EABF  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB CYC:17726
EAC0  20 59 FA  JSR $FA59                       A:F0 X:FF Y:A3 P:E5 SP:FB CYC:17728
FA59  50 0A     BVC $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9 CYC:17734
FA5B  F0 08     BEQ $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9 CYC:17736
FA5D  10 06     BPL $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9 CYC:17738
FA5F  90 04     BCC $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9 CYC:17740
FA61  C9 F0     CMP #$F0                        A:F0 X:FF Y:A3 P:E5 SP:F9 CYC:17742
FA63  F0 02     BEQ $FA67                       A:F0 X:FF Y:A3 P:67 SP:F9 CYC:17744
FA67  60        RTS                             A:F0 X:FF Y:A3 P:67 SP:F9 CYC:17747
EAC3  A5 47     LDA $47 = 36                    A:F0 X:FF Y:A3 P:67 SP:FB CYC:17753
EAC5  C9 36     CMP #$36                        A:36 X:FF Y:A3 P:65 SP:FB CYC:17756
EAC7  F0 02     BEQ $EACB                       A:36 X:FF Y:A3 P:67 SP:FB CYC:17758
EACB  A9 EB     LDA #$EB                        A:36 X:FF Y:A3 P:67 SP:FB CYC:17761
EACD  8D 47 06  STA $0647 = 36                  A:EB X:FF Y:A3 P:E5 SP:FB CYC:17763
EAD0  A0 FF     LDY #$FF                        A:EB X:FF Y:A3 P:E5 SP:FB CYC:17767
EAD2  20 31 FA  JSR $FA31                       A:EB X:FF Y:FF P:E5 SP:FB CYC:17769
FA31  24 01     BIT $01 = FF                    A:EB X:FF Y:FF P:E5 SP:F9 CYC:17775
FA33  18        CLC                             A:EB X:FF Y:FF P:E5 SP:F9 CYC:17778
FA34  A9 40     LDA #$40                        A:EB X:FF Y:FF P:E4 SP:F9 CYC:17780
FA36  60        RTS                             A:40 X:FF Y:FF P:64 SP:F9 CYC:17782
EAD5  DB 48 05 *DCP $0548,Y @ 0647 = EB         A:40 X:FF Y:FF P:64 SP:FB CYC:17788
EAD8  EA        NOP                             A:40 X:FF Y:FF P:64 SP:FB CYC:17795
EAD9  EA        NOP                             A:40 X:FF Y:FF P:64 SP:FB CYC:17797
EADA  08        PHP                             A:40 X:FF Y:FF P:64 SP:FB CYC:17799
EADB  48        PHA                             A:40 X:FF Y:FF P:64 SP:FA CYC:17802
EADC  A0 A4     LDY #$A4                        A:40 X:FF Y:FF P:64 SP:F9 CYC:17805
EADE  68        PLA                             A:40 X:FF Y:A4 P:E4 SP:F9 CYC:17807
EADF  28        PLP                             A:40 X:FF Y:A4 P:64 SP:FA CYC:17811
EAE0  20 37 FA  JSR $FA37                       A:40 X:FF Y:A4 P:64 SP:FB CYC:17815
FA37  50 2C     BVC $FA65                       A:40 X:FF Y:A4 P:64 SP:F9 CYC:17821
FA39  B0 2A     BCS $FA65                       A:40 X:FF Y:A4 P:64 SP:F9 CYC:17823
FA3B  30 28     BMI $FA65                       A:40 X:FF Y:A4 P:64 SP:F9 CYC:17825
FA3D  C9 40     CMP #$40                        A:40 X:FF Y:A4 P:64 SP:F9 CYC:17827
FA3F  D0 24     BNE $FA65                       A:40 X:FF Y:A4 P:67 SP:F9 CYC:17829
FA41  60        RTS                             A:40 X:FF Y:A4 P:67 SP:F9 CYC:17831
EAE3  AD 47 06  LDA $0647 = EA                  A:40 X:FF Y:A4 P:67 SP:FB CYC:17837
EAE6  C9 EA     CMP #$EA                        A:EA X:FF Y:A4 P:E5 SP:FB CYC:17841
EAE8  F0 02     BEQ $EAEC                       A:EA X:FF Y:A4 P:67 SP:FB CYC:17843
EAEC  A0 FF     LDY #$FF                        A:EA X:FF Y:A4 P:67 SP:FB CYC:17846
EAEE  A9 00     LDA #$00                        A:EA X:FF Y:FF P:E5 SP:FB CYC:17848
EAF0  8D 47 06  STA $0647 = EA                  A:0 X:FF Y:FF P:67 SP:FB CYC:17850
EAF3  20 42 FA  JSR $FA42                       A:0 X:FF Y:FF P:67 SP:FB CYC:17854
FA42  B8        CLV                             A:0 X:FF Y:FF P:67 SP:F9 CYC:17860
FA43  38        SEC                             A:0 X:FF Y:FF P:27 SP:F9 CYC:17862
FA44  A9 FF     LDA #$FF                        A:0 X:FF Y:FF P:27 SP:F9 CYC:17864
FA46  60        RTS                             A:FF X:FF Y:FF P:A5 SP:F9 CYC:17866
EAF6  DB 48 05 *DCP $0548,Y @ 0647 = 00         A:FF X:FF Y:FF P:A5 SP:FB CYC:17872
EAF9  EA        NOP                             A:FF X:FF Y:FF P:27 SP:FB CYC:17879
EAFA  EA        NOP                             A:FF X:FF Y:FF P:27 SP:FB CYC:17881
EAFB  08        PHP                             A:FF X:FF Y:FF P:27 SP:FB CYC:17883
EAFC  48        PHA                             A:FF X:FF Y:FF P:27 SP:FA CYC:17886
EAFD  A0 A5     LDY #$A5                        A:FF X:FF Y:FF P:27 SP:F9 CYC:17889
EAFF  68        PLA                             A:FF X:FF Y:A5 P:A5 SP:F9 CYC:17891
EB00  28        PLP                             A:FF X:FF Y:A5 P:A5 SP:FA CYC:17895
EB01  20 47 FA  JSR $FA47                       A:FF X:FF Y:A5 P:27 SP:FB CYC:17899
FA47  70 1C     BVS $FA65                       A:FF X:FF Y:A5 P:27 SP:F9 CYC:17905
FA49  D0 1A     BNE $FA65                       A:FF X:FF Y:A5 P:27 SP:F9 CYC:17907
FA4B  30 18     BMI $FA65                       A:FF X:FF Y:A5 P:27 SP:F9 CYC:17909
FA4D  90 16     BCC $FA65                       A:FF X:FF Y:A5 P:27 SP:F9 CYC:17911
FA4F  C9 FF     CMP #$FF                        A:FF X:FF Y:A5 P:27 SP:F9 CYC:17913
FA51  D0 12     BNE $FA65                       A:FF X:FF Y:A5 P:27 SP:F9 CYC:17915
FA53  60        RTS                             A:FF X:FF Y:A5 P:27 SP:F9 CYC:17917
EB04  AD 47 06  LDA $0647 = FF                  A:FF X:FF Y:A5 P:27 SP:FB CYC:17923
EB07  C9 FF     CMP #$FF                        A:FF X:FF Y:A5 P:A5 SP:FB CYC:17927
EB09  F0 02     BEQ $EB0D                       A:FF X:FF Y:A5 P:27 SP:FB CYC:17929
EB0D  A0 FF     LDY #$FF                        A:FF X:FF Y:A5 P:27 SP:FB CYC:17932
EB0F  A9 37     LDA #$37                        A:FF X:FF Y:FF P:A5 SP:FB CYC:17934
EB11  8D 47 06  STA $0647 = FF                  A:37 X:FF Y:FF P:25 SP:FB CYC:17936
EB14  20 54 FA  JSR $FA54                       A:37 X:FF Y:FF P:25 SP:FB CYC:17940
FA54  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9 CYC:17946
FA56  A9 F0     LDA #$F0                        A:37 X:FF Y:FF P:E5 SP:F9 CYC:17949
FA58  60        RTS                             A:F0 X:FF Y:FF P:E5 SP:F9 CYC:17951
EB17  DB 48 05 *DCP $0548,Y @ 0647 = 37         A:F0 X:FF Y:FF P:E5 SP:FB CYC:17957
EB1A  EA        NOP                             A:F0 X:FF Y:FF P:E5 SP:FB CYC:17964
EB1B  EA        NOP                             A:F0 X:FF Y:FF P:E5 SP:FB CYC:17966
EB1C  08        PHP                             A:F0 X:FF Y:FF P:E5 SP:FB CYC:17968
EB1D  48        PHA                             A:F0 X:FF Y:FF P:E5 SP:FA CYC:17971
EB1E  A0 A6     LDY #$A6                        A:F0 X:FF Y:FF P:E5 SP:F9 CYC:17974
EB20  68        PLA                             A:F0 X:FF Y:A6 P:E5 SP:F9 CYC:17976
EB21  28        PLP                             A:F0 X:FF Y:A6 P:E5 SP:FA CYC:17980
EB22  20 59 FA  JSR $FA59                       A:F0 X:FF Y:A6 P:E5 SP:FB CYC:17984
FA59  50 0A     BVC $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9 CYC:17990
FA5B  F0 08     BEQ $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9 CYC:17992
FA5D  10 06     BPL $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9 CYC:17994
FA5F  90 04     BCC $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9 CYC:17996
FA61  C9 F0     CMP #$F0                        A:F0 X:FF Y:A6 P:E5 SP:F9 CYC:17998
FA63  F0 02     BEQ $FA67                       A:F0 X:FF Y:A6 P:67 SP:F9 CYC:18000
FA67  60        RTS                             A:F0 X:FF Y:A6 P:67 SP:F9 CYC:18003
EB25  AD 47 06  LDA $0647 = 36                  A:F0 X:FF Y:A6 P:67 SP:FB CYC:18009
EB28  C9 36     CMP #$36                        A:36 X:FF Y:A6 P:65 SP:FB CYC:18013
EB2A  F0 02     BEQ $EB2E                       A:36 X:FF Y:A6 P:67 SP:FB CYC:18015
EB2E  A0 A7     LDY #$A7                        A:36 X:FF Y:A6 P:67 SP:FB CYC:18018
EB30  A2 FF     LDX #$FF                        A:36 X:FF Y:A7 P:E5 SP:FB CYC:18020
EB32  A9 EB     LDA #$EB                        A:36 X:FF Y:A7 P:E5 SP:FB CYC:18022
EB34  8D 47 06  STA $0647 = 36                  A:EB X:FF Y:A7 P:E5 SP:FB CYC:18024
EB37  20 31 FA  JSR $FA31                       A:EB X:FF Y:A7 P:E5 SP:FB CYC:18028
FA31  24 01     BIT $01 = FF                    A:EB X:FF Y:A7 P:E5 SP:F9 CYC:18034
FA33  18        CLC                             A:EB X:FF Y:A7 P:E5 SP:F9 CYC:18037
FA34  A9 40     LDA #$40                        A:EB X:FF Y:A7 P:E4 SP:F9 CYC:18039
FA36  60        RTS                             A:40 X:FF Y:A7 P:64 SP:F9 CYC:18041
EB3A  DF 48 05 *DCP $0548,X @ 0647 = EB         A:40 X:FF Y:A7 P:64 SP:FB CYC:18047
EB3D  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB CYC:18054
EB3E  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB CYC:18056
EB3F  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB CYC:18058
EB40  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB CYC:18060
EB41  20 37 FA  JSR $FA37                       A:40 X:FF Y:A7 P:64 SP:FB CYC:18062
FA37  50 2C     BVC $FA65                       A:40 X:FF Y:A7 P:64 SP:F9 CYC:18068
FA39  B0 2A     BCS $FA65                       A:40 X:FF Y:A7 P:64 SP:F9 CYC:18070
FA3B  30 28     BMI $FA65                       A:40 X:FF Y:A7 P:64 SP:F9 CYC:18072
FA3D  C9 40     CMP #$40                        A:40 X:FF Y:A7 P:64 SP:F9 CYC:18074
FA3F  D0 24     BNE $FA65                       A:40 X:FF Y:A7 P:67 SP:F9 CYC:18076
FA41  60        RTS                             A:40 X:FF Y:A7 P:67 SP:F9 CYC:18078
EB44  AD 47 06  LDA $0647 = EA                  A:40 X:FF Y:A7 P:67 SP:FB CYC:18084
EB47  C9 EA     CMP #$EA                        A:EA X:FF Y:A7 P:E5 SP:FB CYC:18088
EB49  F0 02     BEQ $EB4D                       A:EA X:FF Y:A7 P:67 SP:FB CYC:18090
EB4D  C8        INY                             A:EA X:FF Y:A7 P:67 SP:FB CYC:18093
EB4E  A9 00     LDA #$00                        A:EA X:FF Y:A8 P:E5 SP:FB CYC:18095
EB50  8D 47 06  STA $0647 = EA                  A:0 X:FF Y:A8 P:67 SP:FB CYC:18097
EB53  20 42 FA  JSR $FA42                       A:0 X:FF Y:A8 P:67 SP:FB CYC:18101
FA42  B8        CLV                             A:0 X:FF Y:A8 P:67 SP:F9 CYC:18107
FA43  38        SEC                             A:0 X:FF Y:A8 P:27 SP:F9 CYC:18109
FA44  A9 FF     LDA #$FF                        A:0 X:FF Y:A8 P:27 SP:F9 CYC:18111
FA46  60        RTS                             A:FF X:FF Y:A8 P:A5 SP:F9 CYC:18113
EB56  DF 48 05 *DCP $0548,X @ 0647 = 00         A:FF X:FF Y:A8 P:A5 SP:FB CYC:18119
EB59  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB CYC:18126
EB5A  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB CYC:18128
EB5B  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB CYC:18130
EB5C  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB CYC:18132
EB5D  20 47 FA  JSR $FA47                       A:FF X:FF Y:A8 P:27 SP:FB CYC:18134
FA47  70 1C     BVS $FA65                       A:FF X:FF Y:A8 P:27 SP:F9 CYC:18140
FA49  D0 1A     BNE $FA65                       A:FF X:FF Y:A8 P:27 SP:F9 CYC:18142
FA4B  30 18     BMI $FA65                       A:FF X:FF Y:A8 P:27 SP:F9 CYC:18144
FA4D  90 16     BCC $FA65                       A:FF X:FF Y:A8 P:27 SP:F9 CYC:18146
FA4F  C9 FF     CMP #$FF                        A:FF X:FF Y:A8 P:27 SP:F9 CYC:18148
FA51  D0 12     BNE $FA65                       A:FF X:FF Y:A8 P:27 SP:F9 CYC:18150
FA53  60        RTS                             A:FF X:FF Y:A8 P:27 SP:F9 CYC:18152
EB60  AD 47 06  LDA $0647 = FF                  A:FF X:FF Y:A8 P:27 SP:FB CYC:18158
EB63  C9 FF     CMP #$FF                        A:FF X:FF Y:A8 P:A5 SP:FB CYC:18162
EB65  F0 02     BEQ $EB69                       A:FF X:FF Y:A8 P:27 SP:FB CYC:18164
EB69  C8        INY                             A:FF X:FF Y:A8 P:27 SP:FB CYC:18167
EB6A  A9 37     LDA #$37                        A:FF X:FF Y:A9 P:A5 SP:FB CYC:18169
EB6C  8D 47 06  STA $0647 = FF                  A:37 X:FF Y:A9 P:25 SP:FB CYC:18171
EB6F  20 54 FA  JSR $FA54                       A:37 X:FF Y:A9 P:25 SP:FB CYC:18175
FA54  24 01     BIT $01 = FF                    A:37 X:FF Y:A9 P:25 SP:F9 CYC:18181
FA56  A9 F0     LDA #$F0                        A:37 X:FF Y:A9 P:E5 SP:F9 CYC:18184
FA58  60        RTS                             A:F0 X:FF Y:A9 P:E5 SP:F9 CYC:18186
EB72  DF 48 05 *DCP $0548,X @ 0647 = 37         A:F0 X:FF Y:A9 P:E5 SP:FB CYC:18192
EB75  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB CYC:18199
EB76  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB CYC:18201
EB77  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB CYC:18203
EB78  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB CYC:18205
EB79  20 59 FA  JSR $FA59                       A:F0 X:FF Y:A9 P:E5 SP:FB CYC:18207
FA59  50 0A     BVC $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9 CYC:18213
FA5B  F0 08     BEQ $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9 CYC:18215
FA5D  10 06     BPL $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9 CYC:18217
FA5F  90 04     BCC $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9 CYC:18219
FA61  C9 F0     CMP #$F0                        A:F0 X:FF Y:A9 P:E5 SP:F9 CYC:18221
FA63  F0 02     BEQ $FA67                       A:F0 X:FF Y:A9 P:67 SP:F9 CYC:18223
FA67  60        RTS                             A:F0 X:FF Y:A9 P:67 SP:F9 CYC:18226
EB7C  AD 47 06  LDA $0647 = 36                  A:F0 X:FF Y:A9 P:67 SP:FB CYC:18232
EB7F  C9 36     CMP #$36                        A:36 X:FF Y:A9 P:65 SP:FB CYC:18236
EB81  F0 02     BEQ $EB85                       A:36 X:FF Y:A9 P:67 SP:FB CYC:18238
EB85  60        RTS                             A:36 X:FF Y:A9 P:67 SP:FB CYC:18241
C63E  20 86 EB  JSR $EB86                       A:36 X:FF Y:A9 P:67 SP:FD CYC:18247
EB86  A9 FF     LDA #$FF                        A:36 X:FF Y:A9 P:67 SP:FB CYC:18253
EB88  85 01     STA $01 = FF                    A:FF X:FF Y:A9 P:E5 SP:FB CYC:18255
EB8A  A0 AA     LDY #$AA                        A:FF X:FF Y:A9 P:E5 SP:FB CYC:18258
EB8C  A2 02     LDX #$02                        A:FF X:FF Y:AA P:E5 SP:FB CYC:18260
EB8E  A9 47     LDA #$47                        A:FF X:2 Y:AA P:65 SP:FB CYC:18262
EB90  85 47     STA $47 = 36                    A:47 X:2 Y:AA P:65 SP:FB CYC:18264
EB92  A9 06     LDA #$06                        A:47 X:2 Y:AA P:65 SP:FB CYC:18267
EB94  85 48     STA $48 = 06                    A:6 X:2 Y:AA P:65 SP:FB CYC:18269
EB96  A9 EB     LDA #$EB                        A:6 X:2 Y:AA P:65 SP:FB CYC:18272
EB98  8D 47 06  STA $0647 = 36                  A:EB X:2 Y:AA P:E5 SP:FB CYC:18274
EB9B  20 B1 FA  JSR $FAB1                       A:EB X:2 Y:AA P:E5 SP:FB CYC:18278
FAB1  24 01     BIT $01 = FF                    A:EB X:2 Y:AA P:E5 SP:F9 CYC:18284
FAB3  18        CLC                             A:EB X:2 Y:AA P:E5 SP:F9 CYC:18287
FAB4  A9 40     LDA #$40                        A:EB X:2 Y:AA P:E4 SP:F9 CYC:18289
FAB6  60        RTS                             A:40 X:2 Y:AA P:64 SP:F9 CYC:18291
EB9E  E3 45    *ISB ($45,X) @ 47 = 0647 = EB    A:40 X:2 Y:AA P:64 SP:FB CYC:18297
EBA0  EA        NOP                             A:53 X:2 Y:AA P:24 SP:FB CYC:18305
EBA1  EA        NOP                             A:53 X:2 Y:AA P:24 SP:FB CYC:18307
EBA2  EA        NOP                             A:53 X:2 Y:AA P:24 SP:FB CYC:18309
EBA3  EA        NOP                             A:53 X:2 Y:AA P:24 SP:FB CYC:18311
EBA4  20 B7 FA  JSR $FAB7                       A:53 X:2 Y:AA P:24 SP:FB CYC:18313
FAB7  70 2D     BVS $FAE6                       A:53 X:2 Y:AA P:24 SP:F9 CYC:18319
FAB9  B0 2B     BCS $FAE6                       A:53 X:2 Y:AA P:24 SP:F9 CYC:18321
FABB  30 29     BMI $FAE6                       A:53 X:2 Y:AA P:24 SP:F9 CYC:18323
FABD  C9 53     CMP #$53                        A:53 X:2 Y:AA P:24 SP:F9 CYC:18325
FABF  D0 25     BNE $FAE6                       A:53 X:2 Y:AA P:27 SP:F9 CYC:18327
FAC1  60        RTS                             A:53 X:2 Y:AA P:27 SP:F9 CYC:18329
EBA7  AD 47 06  LDA $0647 = EC                  A:53 X:2 Y:AA P:27 SP:FB CYC:18335
EBAA  C9 EC     CMP #$EC                        A:EC X:2 Y:AA P:A5 SP:FB CYC:18339
EBAC  F0 02     BEQ $EBB0                       A:EC X:2 Y:AA P:27 SP:FB CYC:18341
EBB0  C8        INY                             A:EC X:2 Y:AA P:27 SP:FB CYC:18344
EBB1  A9 FF     LDA #$FF                        A:EC X:2 Y:AB P:A5 SP:FB CYC:18346
EBB3  8D 47 06  STA $0647 = EC                  A:FF X:2 Y:AB P:A5 SP:FB CYC:18348
EBB6  20 C2 FA  JSR $FAC2                       A:FF X:2 Y:AB P:A5 SP:FB CYC:18352
FAC2  B8        CLV                             A:FF X:2 Y:AB P:A5 SP:F9 CYC:18358
FAC3  38        SEC                             A:FF X:2 Y:AB P:A5 SP:F9 CYC:18360
FAC4  A9 FF     LDA #$FF                        A:FF X:2 Y:AB P:A5 SP:F9 CYC:18362
FAC6  60        RTS                             A:FF X:2 Y:AB P:A5 SP:F9 CYC:18364
EBB9  E3 45    *ISB ($45,X) @ 47 = 0647 = FF    A:FF X:2 Y:AB P:A5 SP:FB CYC:18370
EBBB  EA        NOP                             A:FF X:2 Y:AB P:A5 SP:FB CYC:18378
EBBC  EA        NOP                             A:FF X:2 Y:AB P:A5 SP:FB CYC:18380
EBBD  EA        NOP                             A:FF X:2 Y:AB P:A5 SP:FB CYC:18382
EBBE  EA        NOP                             A:FF X:2 Y:AB P:A5 SP:FB CYC:18384
EBBF  20 C7 FA  JSR $FAC7                       A:FF X:2 Y:AB P:A5 SP:FB CYC:18386
FAC7  70 1D     BVS $FAE6                       A:FF X:2 Y:AB P:A5 SP:F9 CYC:18392
FAC9  F0 1B     BEQ $FAE6                       A:FF X:2 Y:AB P:A5 SP:F9 CYC:18394
FACB  10 19     BPL $FAE6                       A:FF X:2 Y:AB P:A5 SP:F9 CYC:18396
FACD  90 17     BCC $FAE6                       A:FF X:2 Y:AB P:A5 SP:F9 CYC:18398
FACF  C9 FF     CMP #$FF                        A:FF X:2 Y:AB P:A5 SP:F9 CYC:18400
FAD1  D0 13     BNE $FAE6                       A:FF X:2 Y:AB P:27 SP:F9 CYC:18402
FAD3  60        RTS                             A:FF X:2 Y:AB P:27 SP:F9 CYC:18404
EBC2  AD 47 06  LDA $0647 = 00                  A:FF X:2 Y:AB P:27 SP:FB CYC:18410
EBC5  C9 00     CMP #$00                        A:0 X:2 Y:AB P:27 SP:FB CYC:18414
EBC7  F0 02     BEQ $EBCB                       A:0 X:2 Y:AB P:27 SP:FB CYC:18416
EBCB  C8        INY                             A:0 X:2 Y:AB P:27 SP:FB CYC:18419
EBCC  A9 37     LDA #$37                        A:0 X:2 Y:AC P:A5 SP:FB CYC:18421
EBCE  8D 47 06  STA $0647 = 00                  A:37 X:2 Y:AC P:25 SP:FB CYC:18423
EBD1  20 D4 FA  JSR $FAD4                       A:37 X:2 Y:AC P:25 SP:FB CYC:18427
FAD4  24 01     BIT $01 = FF                    A:37 X:2 Y:AC P:25 SP:F9 CYC:18433
FAD6  38        SEC                             A:37 X:2 Y:AC P:E5 SP:F9 CYC:18436
FAD7  A9 F0     LDA #$F0                        A:37 X:2 Y:AC P:E5 SP:F9 CYC:18438
FAD9  60        RTS                             A:F0 X:2 Y:AC P:E5 SP:F9 CYC:18440
EBD4  E3 45    *ISB ($45,X) @ 47 = 0647 = 37    A:F0 X:2 Y:AC P:E5 SP:FB CYC:18446
EBD6  EA        NOP                             A:B8 X:2 Y:AC P:A5 SP:FB CYC:18454
EBD7  EA        NOP                             A:B8 X:2 Y:AC P:A5 SP:FB CYC:18456
EBD8  EA        NOP                             A:B8 X:2 Y:AC P:A5 SP:FB CYC:18458
EBD9  EA        NOP                             A:B8 X:2 Y:AC P:A5 SP:FB CYC:18460
EBDA  20 DA FA  JSR $FADA                       A:B8 X:2 Y:AC P:A5 SP:FB CYC:18462
FADA  70 0A     BVS $FAE6                       A:B8 X:2 Y:AC P:A5 SP:F9 CYC:18468
FADC  F0 08     BEQ $FAE6                       A:B8 X:2 Y:AC P:A5 SP:F9 CYC:18470
FADE  10 06     BPL $FAE6                       A:B8 X:2 Y:AC P:A5 SP:F9 CYC:18472
FAE0  90 04     BCC $FAE6                       A:B8 X:2 Y:AC P:A5 SP:F9 CYC:18474
FAE2  C9 B8     CMP #$B8                        A:B8 X:2 Y:AC P:A5 SP:F9 CYC:18476
FAE4  F0 02     BEQ $FAE8                       A:B8 X:2 Y:AC P:27 SP:F9 CYC:18478
FAE8  60        RTS                             A:B8 X:2 Y:AC P:27 SP:F9 CYC:18481
EBDD  AD 47 06  LDA $0647 = 38                  A:B8 X:2 Y:AC P:27 SP:FB CYC:18487
EBE0  C9 38     CMP #$38                        A:38 X:2 Y:AC P:25 SP:FB CYC:18491
EBE2  F0 02     BEQ $EBE6                       A:38 X:2 Y:AC P:27 SP:FB CYC:18493
EBE6  C8        INY                             A:38 X:2 Y:AC P:27 SP:FB CYC:18496
EBE7  A9 EB     LDA #$EB                        A:38 X:2 Y:AD P:A5 SP:FB CYC:18498
EBE9  85 47     STA $47 = 47                    A:EB X:2 Y:AD P:A5 SP:FB CYC:18500
EBEB  20 B1 FA  JSR $FAB1                       A:EB X:2 Y:AD P:A5 SP:FB CYC:18503
FAB1  24 01     BIT $01 = FF                    A:EB X:2 Y:AD P:A5 SP:F9 CYC:18509
FAB3  18        CLC                             A:EB X:2 Y:AD P:E5 SP:F9 CYC:18512
FAB4  A9 40     LDA #$40                        A:EB X:2 Y:AD P:E4 SP:F9 CYC:18514
FAB6  60        RTS                             A:40 X:2 Y:AD P:64 SP:F9 CYC:18516
EBEE  E7 47    *ISB $47 = EB                    A:40 X:2 Y:AD P:64 SP:FB CYC:18522
EBF0  EA        NOP                             A:53 X:2 Y:AD P:24 SP:FB CYC:18527
EBF1  EA        NOP                             A:53 X:2 Y:AD P:24 SP:FB CYC:18529
EBF2  EA        NOP                             A:53 X:2 Y:AD P:24 SP:FB CYC:18531
EBF3  EA        NOP                             A:53 X:2 Y:AD P:24 SP:FB CYC:18533
EBF4  20 B7 FA  JSR $FAB7                       A:53 X:2 Y:AD P:24 SP:FB CYC:18535
FAB7  70 2D     BVS $FAE6                       A:53 X:2 Y:AD P:24 SP:F9 CYC:18541
FAB9  B0 2B     BCS $FAE6                       A:53 X:2 Y:AD P:24 SP:F9 CYC:18543
FABB  30 29     BMI $FAE6                       A:53 X:2 Y:AD P:24 SP:F9 CYC:18545
FABD  C9 53     CMP #$53                        A:53 X:2 Y:AD P:24 SP:F9 CYC:18547
FABF  D0 25     BNE $FAE6                       A:53 X:2 Y:AD P:27 SP:F9 CYC:18549
FAC1  60        RTS                             A:53 X:2 Y:AD P:27 SP:F9 CYC:18551
EBF7  A5 47     LDA $47 = EC                    A:53 X:2 Y:AD P:27 SP:FB CYC:18557
EBF9  C9 EC     CMP #$EC                        A:EC X:2 Y:AD P:A5 SP:FB CYC:18560
EBFB  F0 02     BEQ $EBFF                       A:EC X:2 Y:AD P:27 SP:FB CYC:18562
EBFF  C8        INY                             A:EC X:2 Y:AD P:27 SP:FB CYC:18565
EC00  A9 FF     LDA #$FF                        A:EC X:2 Y:AE P:A5 SP:FB CYC:18567
EC02  85 47     STA $47 = EC                    A:FF X:2 Y:AE P:A5 SP:FB CYC:18569
EC04  20 C2 FA  JSR $FAC2                       A:FF X:2 Y:AE P:A5 SP:FB CYC:18572
FAC2  B8        CLV                             A:FF X:2 Y:AE P:A5 SP:F9 CYC:18578
FAC3  38        SEC                             A:FF X:2 Y:AE P:A5 SP:F9 CYC:18580
FAC4  A9 FF     LDA #$FF                        A:FF X:2 Y:AE P:A5 SP:F9 CYC:18582
FAC6  60        RTS                             A:FF X:2 Y:AE P:A5 SP:F9 CYC:18584
EC07  E7 47    *ISB $47 = FF                    A:FF X:2 Y:AE P:A5 SP:FB CYC:18590
EC09  EA        NOP                             A:FF X:2 Y:AE P:A5 SP:FB CYC:18595
EC0A  EA        NOP                             A:FF X:2 Y:AE P:A5 SP:FB CYC:18597
EC0B  EA        NOP                             A:FF X:2 Y:AE P:A5 SP:FB CYC:18599
EC0C  EA        NOP                             A:FF X:2 Y:AE P:A5 SP:FB CYC:18601
EC0D  20 C7 FA  JSR $FAC7                       A:FF X:2 Y:AE P:A5 SP:FB CYC:18603
FAC7  70 1D     BVS $FAE6                       A:FF X:2 Y:AE P:A5 SP:F9 CYC:18609
FAC9  F0 1B     BEQ $FAE6                       A:FF X:2 Y:AE P:A5 SP:F9 CYC:18611
FACB  10 19     BPL $FAE6                       A:FF X:2 Y:AE P:A5 SP:F9 CYC:18613
FACD  90 17     BCC $FAE6                       A:FF X:2 Y:AE P:A5 SP:F9 CYC:18615
FACF  C9 FF     CMP #$FF                        A:FF X:2 Y:AE P:A5 SP:F9 CYC:18617
FAD1  D0 13     BNE $FAE6                       A:FF X:2 Y:AE P:27 SP:F9 CYC:18619
FAD3  60        RTS                             A:FF X:2 Y:AE P:27 SP:F9 CYC:18621
EC10  A5 47     LDA $47 = 00                    A:FF X:2 Y:AE P:27 SP:FB CYC:18627
EC12  C9 00     CMP #$00                        A:0 X:2 Y:AE P:27 SP:FB CYC:18630
EC14  F0 02     BEQ $EC18                       A:0 X:2 Y:AE P:27 SP:FB CYC:18632
EC18  C8        INY                             A:0 X:2 Y:AE P:27 SP:FB CYC:18635
EC19  A9 37     LDA #$37                        A:0 X:2 Y:AF P:A5 SP:FB CYC:18637
EC1B  85 47     STA $47 = 00                    A:37 X:2 Y:AF P:25 SP:FB CYC:18639
EC1D  20 D4 FA  JSR $FAD4                       A:37 X:2 Y:AF P:25 SP:FB CYC:18642
FAD4  24 01     BIT $01 = FF                    A:37 X:2 Y:AF P:25 SP:F9 CYC:18648
FAD6  38        SEC                             A:37 X:2 Y:AF P:E5 SP:F9 CYC:18651
FAD7  A9 F0     LDA #$F0                        A:37 X:2 Y:AF P:E5 SP:F9 CYC:18653
FAD9  60        RTS                             A:F0 X:2 Y:AF P:E5 SP:F9 CYC:18655
EC20  E7 47    *ISB $47 = 37                    A:F0 X:2 Y:AF P:E5 SP:FB CYC:18661
EC22  EA        NOP                             A:B8 X:2 Y:AF P:A5 SP:FB CYC:18666
EC23  EA        NOP                             A:B8 X:2 Y:AF P:A5 SP:FB CYC:18668
EC24  EA        NOP                             A:B8 X:2 Y:AF P:A5 SP:FB CYC:18670
EC25  EA        NOP                             A:B8 X:2 Y:AF P:A5 SP:FB CYC:18672
EC26  20 DA FA  JSR $FADA                       A:B8 X:2 Y:AF P:A5 SP:FB CYC:18674
FADA  70 0A     BVS $FAE6                       A:B8 X:2 Y:AF P:A5 SP:F9 CYC:18680
FADC  F0 08     BEQ $FAE6                       A:B8 X:2 Y:AF P:A5 SP:F9 CYC:18682
FADE  10 06     BPL $FAE6                       A:B8 X:2 Y:AF P:A5 SP:F9 CYC:18684
FAE0  90 04     BCC $FAE6                       A:B8 X:2 Y:AF P:A5 SP:F9 CYC:18686
FAE2  C9 B8     CMP #$B8                        A:B8 X:2 Y:AF P:A5 SP:F9 CYC:18688
FAE4  F0 02     BEQ $FAE8                       A:B8 X:2 Y:AF P:27 SP:F9 CYC:18690
FAE8  60        RTS                             A:B8 X:2 Y:AF P:27 SP:F9 CYC:18693
EC29  A5 47     LDA $47 = 38                    A:B8 X:2 Y:AF P:27 SP:FB CYC:18699
EC2B  C9 38     CMP #$38                        A:38 X:2 Y:AF P:25 SP:FB CYC:18702
EC2D  F0 02     BEQ $EC31                       A:38 X:2 Y:AF P:27 SP:FB CYC:18704
EC31  C8        INY                             A:38 X:2 Y:AF P:27 SP:FB CYC:18707
EC32  A9 EB     LDA #$EB                        A:38 X:2 Y:B0 P:A5 SP:FB CYC:18709
EC34  8D 47 06  STA $0647 = 38                  A:EB X:2 Y:B0 P:A5 SP:FB CYC:18711
EC37  20 B1 FA  JSR $FAB1                       A:EB X:2 Y:B0 P:A5 SP:FB CYC:18715
FAB1  24 01     BIT $01 = FF                    A:EB X:2 Y:B0 P:A5 SP:F9 CYC:18721
FAB3  18        CLC                             A:EB X:2 Y:B0 P:E5 SP:F9 CYC:18724
FAB4  A9 40     LDA #$40                        A:EB X:2 Y:B0 P:E4 SP:F9 CYC:18726
FAB6  60        RTS                             A:40 X:2 Y:B0 P:64 SP:F9 CYC:18728
EC3A  EF 47 06 *ISB $0647 = EB                  A:40 X:2 Y:B0 P:64 SP:FB CYC:18734
EC3D  EA        NOP                             A:53 X:2 Y:B0 P:24 SP:FB CYC:18740
EC3E  EA        NOP                             A:53 X:2 Y:B0 P:24 SP:FB CYC:18742
EC3F  EA        NOP                             A:53 X:2 Y:B0 P:24 SP:FB CYC:18744
EC40  EA        NOP                             A:53 X:2 Y:B0 P:24 SP:FB CYC:18746
EC41  20 B7 FA  JSR $FAB7                       A:53 X:2 Y:B0 P:24 SP:FB CYC:18748
FAB7  70 2D     BVS $FAE6                       A:53 X:2 Y:B0 P:24 SP:F9 CYC:18754
FAB9  B0 2B     BCS $FAE6                       A:53 X:2 Y:B0 P:24 SP:F9 CYC:18756
FABB  30 29     BMI $FAE6                       A:53 X:2 Y:B0 P:24 SP:F9 CYC:18758
FABD  C9 53     CMP #$53                        A:53 X:2 Y:B0 P:24 SP:F9 CYC:18760
FABF  D0 25     BNE $FAE6                       A:53 X:2 Y:B0 P:27 SP:F9 CYC:18762
FAC1  60        RTS                             A:53 X:2 Y:B0 P:27 SP:F9 CYC:18764
EC44  AD 47 06  LDA $0647 = EC                  A:53 X:2 Y:B0 P:27 SP:FB CYC:18770
EC47  C9 EC     CMP #$EC                        A:EC X:2 Y:B0 P:A5 SP:FB CYC:18774
EC49  F0 02     BEQ $EC4D                       A:EC X:2 Y:B0 P:27 SP:FB CYC:18776
EC4D  C8        INY                             A:EC X:2 Y:B0 P:27 SP:FB CYC:18779
EC4E  A9 FF     LDA #$FF                        A:EC X:2 Y:B1 P:A5 SP:FB CYC:18781
EC50  8D 47 06  STA $0647 = EC                  A:FF X:2 Y:B1 P:A5 SP:FB CYC:18783
EC53  20 C2 FA  JSR $FAC2                       A:FF X:2 Y:B1 P:A5 SP:FB CYC:18787
FAC2  B8        CLV                             A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18793
FAC3  38        SEC                             A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18795
FAC4  A9 FF     LDA #$FF                        A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18797
FAC6  60        RTS                             A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18799
EC56  EF 47 06 *ISB $0647 = FF                  A:FF X:2 Y:B1 P:A5 SP:FB CYC:18805
EC59  EA        NOP                             A:FF X:2 Y:B1 P:A5 SP:FB CYC:18811
EC5A  EA        NOP                             A:FF X:2 Y:B1 P:A5 SP:FB CYC:18813
EC5B  EA        NOP                             A:FF X:2 Y:B1 P:A5 SP:FB CYC:18815
EC5C  EA        NOP                             A:FF X:2 Y:B1 P:A5 SP:FB CYC:18817
EC5D  20 C7 FA  JSR $FAC7                       A:FF X:2 Y:B1 P:A5 SP:FB CYC:18819
FAC7  70 1D     BVS $FAE6                       A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18825
FAC9  F0 1B     BEQ $FAE6                       A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18827
FACB  10 19     BPL $FAE6                       A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18829
FACD  90 17     BCC $FAE6                       A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18831
FACF  C9 FF     CMP #$FF                        A:FF X:2 Y:B1 P:A5 SP:F9 CYC:18833
FAD1  D0 13     BNE $FAE6                       A:FF X:2 Y:B1 P:27 SP:F9 CYC:18835
FAD3  60        RTS                             A:FF X:2 Y:B1 P:27 SP:F9 CYC:18837
EC60  AD 47 06  LDA $0647 = 00                  A:FF X:2 Y:B1 P:27 SP:FB CYC:18843
EC63  C9 00     CMP #$00                        A:0 X:2 Y:B1 P:27 SP:FB CYC:18847
EC65  F0 02     BEQ $EC69                       A:0 X:2 Y:B1 P:27 SP:FB CYC:18849
EC69  C8        INY                             A:0 X:2 Y:B1 P:27 SP:FB CYC:18852
EC6A  A9 37     LDA #$37                        A:0 X:2 Y:B2 P:A5 SP:FB CYC:18854
EC6C  8D 47 06  STA $0647 = 00                  A:37 X:2 Y:B2 P:25 SP:FB CYC:18856
EC6F  20 D4 FA  JSR $FAD4                       A:37 X:2 Y:B2 P:25 SP:FB CYC:18860
FAD4  24 01     BIT $01 = FF                    A:37 X:2 Y:B2 P:25 SP:F9 CYC:18866
FAD6  38        SEC                             A:37 X:2 Y:B2 P:E5 SP:F9 CYC:18869
FAD7  A9 F0     LDA #$F0                        A:37 X:2 Y:B2 P:E5 SP:F9 CYC:18871
FAD9  60        RTS                             A:F0 X:2 Y:B2 P:E5 SP:F9 CYC:18873
EC72  EF 47 06 *ISB $0647 = 37                  A:F0 X:2 Y:B2 P:E5 SP:FB CYC:18879
EC75  EA        NOP                             A:B8 X:2 Y:B2 P:A5 SP:FB CYC:18885
EC76  EA        NOP                             A:B8 X:2 Y:B2 P:A5 SP:FB CYC:18887
EC77  EA        NOP                             A:B8 X:2 Y:B2 P:A5 SP:FB CYC:18889
EC78  EA        NOP                             A:B8 X:2 Y:B2 P:A5 SP:FB CYC:18891
EC79  20 DA FA  JSR $FADA                       A:B8 X:2 Y:B2 P:A5 SP:FB CYC:18893
FADA  70 0A     BVS $FAE6                       A:B8 X:2 Y:B2 P:A5 SP:F9 CYC:18899
FADC  F0 08     BEQ $FAE6                       A:B8 X:2 Y:B2 P:A5 SP:F9 CYC:18901
FADE  10 06     BPL $FAE6                       A:B8 X:2 Y:B2 P:A5 SP:F9 CYC:18903
FAE0  90 04     BCC $FAE6                       A:B8 X:2 Y:B2 P:A5 SP:F9 CYC:18905
FAE2  C9 B8     CMP #$B8                        A:B8 X:2 Y:B2 P:A5 SP:F9 CYC:18907
FAE4  F0 02     BEQ $FAE8                       A:B8 X:2 Y:B2 P:27 SP:F9 CYC:18909
FAE8  60        RTS                             A:B8 X:2 Y:B2 P:27 SP:F9 CYC:18912
EC7C  AD 47 06  LDA $0647 = 38                  A:B8 X:2 Y:B2 P:27 SP:FB CYC:18918
EC7F  C9 38     CMP #$38                        A:38 X:2 Y:B2 P:25 SP:FB CYC:18922
EC81  F0 02     BEQ $EC85                       A:38 X:2 Y:B2 P:27 SP:FB CYC:18924
EC85  A9 EB     LDA #$EB                        A:38 X:2 Y:B2 P:27 SP:FB CYC:18927
EC87  8D 47 06  STA $0647 = 38                  A:EB X:2 Y:B2 P:A5 SP:FB CYC:18929
EC8A  A9 48     LDA #$48                        A:EB X:2 Y:B2 P:A5 SP:FB CYC:18933
EC8C  85 45     STA $45 = 48                    A:48 X:2 Y:B2 P:25 SP:FB CYC:18935
EC8E  A9 05     LDA #$05                        A:48 X:2 Y:B2 P:25 SP:FB CYC:18938
EC90  85 46     STA $46 = 05                    A:5 X:2 Y:B2 P:25 SP:FB CYC:18940
EC92  A0 FF     LDY #$FF                        A:5 X:2 Y:B2 P:25 SP:FB CYC:18943
EC94  20 B1 FA  JSR $FAB1                       A:5 X:2 Y:FF P:A5 SP:FB CYC:18945
FAB1  24 01     BIT $01 = FF                    A:5 X:2 Y:FF P:A5 SP:F9 CYC:18951
FAB3  18        CLC                             A:5 X:2 Y:FF P:E5 SP:F9 CYC:18954
FAB4  A9 40     LDA #$40                        A:5 X:2 Y:FF P:E4 SP:F9 CYC:18956
FAB6  60        RTS                             A:40 X:2 Y:FF P:64 SP:F9 CYC:18958
EC97  F3 45    *ISB ($45),Y = 0548 @ 0647 = EB  A:40 X:2 Y:FF P:64 SP:FB CYC:18964
EC99  EA        NOP                             A:53 X:2 Y:FF P:24 SP:FB CYC:18972
EC9A  EA        NOP                             A:53 X:2 Y:FF P:24 SP:FB CYC:18974
EC9B  08        PHP                             A:53 X:2 Y:FF P:24 SP:FB CYC:18976
EC9C  48        PHA                             A:53 X:2 Y:FF P:24 SP:FA CYC:18979
EC9D  A0 B3     LDY #$B3                        A:53 X:2 Y:FF P:24 SP:F9 CYC:18982
EC9F  68        PLA                             A:53 X:2 Y:B3 P:A4 SP:F9 CYC:18984
ECA0  28        PLP                             A:53 X:2 Y:B3 P:24 SP:FA CYC:18988
ECA1  20 B7 FA  JSR $FAB7                       A:53 X:2 Y:B3 P:24 SP:FB CYC:18992
FAB7  70 2D     BVS $FAE6                       A:53 X:2 Y:B3 P:24 SP:F9 CYC:18998
FAB9  B0 2B     BCS $FAE6                       A:53 X:2 Y:B3 P:24 SP:F9 CYC:19000
FABB  30 29     BMI $FAE6                       A:53 X:2 Y:B3 P:24 SP:F9 CYC:19002
FABD  C9 53     CMP #$53                        A:53 X:2 Y:B3 P:24 SP:F9 CYC:19004
FABF  D0 25     BNE $FAE6                       A:53 X:2 Y:B3 P:27 SP:F9 CYC:19006
FAC1  60        RTS                             A:53 X:2 Y:B3 P:27 SP:F9 CYC:19008
ECA4  AD 47 06  LDA $0647 = EC                  A:53 X:2 Y:B3 P:27 SP:FB CYC:19014
ECA7  C9 EC     CMP #$EC                        A:EC X:2 Y:B3 P:A5 SP:FB CYC:19018
ECA9  F0 02     BEQ $ECAD                       A:EC X:2 Y:B3 P:27 SP:FB CYC:19020
ECAD  A0 FF     LDY #$FF                        A:EC X:2 Y:B3 P:27 SP:FB CYC:19023
ECAF  A9 FF     LDA #$FF                        A:EC X:2 Y:FF P:A5 SP:FB CYC:19025
ECB1  8D 47 06  STA $0647 = EC                  A:FF X:2 Y:FF P:A5 SP:FB CYC:19027
ECB4  20 C2 FA  JSR $FAC2                       A:FF X:2 Y:FF P:A5 SP:FB CYC:19031
FAC2  B8        CLV                             A:FF X:2 Y:FF P:A5 SP:F9 CYC:19037
FAC3  38        SEC                             A:FF X:2 Y:FF P:A5 SP:F9 CYC:19039
FAC4  A9 FF     LDA #$FF                        A:FF X:2 Y:FF P:A5 SP:F9 CYC:19041
FAC6  60        RTS                             A:FF X:2 Y:FF P:A5 SP:F9 CYC:19043
ECB7  F3 45    *ISB ($45),Y = 0548 @ 0647 = FF  A:FF X:2 Y:FF P:A5 SP:FB CYC:19049
ECB9  EA        NOP                             A:FF X:2 Y:FF P:A5 SP:FB CYC:19057
ECBA  EA        NOP                             A:FF X:2 Y:FF P:A5 SP:FB CYC:19059
ECBB  08        PHP                             A:FF X:2 Y:FF P:A5 SP:FB CYC:19061
ECBC  48        PHA                             A:FF X:2 Y:FF P:A5 SP:FA CYC:19064
ECBD  A0 B4     LDY #$B4                        A:FF X:2 Y:FF P:A5 SP:F9 CYC:19067
ECBF  68        PLA                             A:FF X:2 Y:B4 P:A5 SP:F9 CYC:19069
ECC0  28        PLP                             A:FF X:2 Y:B4 P:A5 SP:FA CYC:19073
ECC1  20 C7 FA  JSR $FAC7                       A:FF X:2 Y:B4 P:A5 SP:FB CYC:19077
FAC7  70 1D     BVS $FAE6                       A:FF X:2 Y:B4 P:A5 SP:F9 CYC:19083
FAC9  F0 1B     BEQ $FAE6                       A:FF X:2 Y:B4 P:A5 SP:F9 CYC:19085
FACB  10 19     BPL $FAE6                       A:FF X:2 Y:B4 P:A5 SP:F9 CYC:19087
FACD  90 17     BCC $FAE6                       A:FF X:2 Y:B4 P:A5 SP:F9 CYC:19089
FACF  C9 FF     CMP #$FF                        A:FF X:2 Y:B4 P:A5 SP:F9 CYC:19091
FAD1  D0 13     BNE $FAE6                       A:FF X:2 Y:B4 P:27 SP:F9 CYC:19093
FAD3  60        RTS                             A:FF X:2 Y:B4 P:27 SP:F9 CYC:19095
ECC4  AD 47 06  LDA $0647 = 00                  A:FF X:2 Y:B4 P:27 SP:FB CYC:19101
ECC7  C9 00     CMP #$00                        A:0 X:2 Y:B4 P:27 SP:FB CYC:19105
ECC9  F0 02     BEQ $ECCD                       A:0 X:2 Y:B4 P:27 SP:FB CYC:19107
ECCD  A0 FF     LDY #$FF                        A:0 X:2 Y:B4 P:27 SP:FB CYC:19110
ECCF  A9 37     LDA #$37                        A:0 X:2 Y:FF P:A5 SP:FB CYC:19112
ECD1  8D 47 06  STA $0647 = 00                  A:37 X:2 Y:FF P:25 SP:FB CYC:19114
ECD4  20 D4 FA  JSR $FAD4                       A:37 X:2 Y:FF P:25 SP:FB CYC:19118
FAD4  24 01     BIT $01 = FF                    A:37 X:2 Y:FF P:25 SP:F9 CYC:19124
FAD6  38        SEC                             A:37 X:2 Y:FF P:E5 SP:F9 CYC:19127
FAD7  A9 F0     LDA #$F0                        A:37 X:2 Y:FF P:E5 SP:F9 CYC:19129
FAD9  60        RTS                             A:F0 X:2 Y:FF P:E5 SP:F9 CYC:19131
ECD7  F3 45    *ISB ($45),Y = 0548 @ 0647 = 37  A:F0 X:2 Y:FF P:E5 SP:FB CYC:19137
ECD9  EA        NOP                             A:B8 X:2 Y:FF P:A5 SP:FB CYC:19145
ECDA  EA        NOP                             A:B8 X:2 Y:FF P:A5 SP:FB CYC:19147
ECDB  08        PHP                             A:B8 X:2 Y:FF P:A5 SP:FB CYC:19149
ECDC  48        PHA                             A:B8 X:2 Y:FF P:A5 SP:FA CYC:19152
ECDD  A0 B5     LDY #$B5                        A:B8 X:2 Y:FF P:A5 SP:F9 CYC:19155
ECDF  68        PLA                             A:B8 X:2 Y:B5 P:A5 SP:F9 CYC:19157
ECE0  28        PLP                             A:B8 X:2 Y:B5 P:A5 SP:FA CYC:19161
ECE1  20 DA FA  JSR $FADA                       A:B8 X:2 Y:B5 P:A5 SP:FB CYC:19165
FADA  70 0A     BVS $FAE6                       A:B8 X:2 Y:B5 P:A5 SP:F9 CYC:19171
FADC  F0 08     BEQ $FAE6                       A:B8 X:2 Y:B5 P:A5 SP:F9 CYC:19173
FADE  10 06     BPL $FAE6                       A:B8 X:2 Y:B5 P:A5 SP:F9 CYC:19175
FAE0  90 04     BCC $FAE6                       A:B8 X:2 Y:B5 P:A5 SP:F9 CYC:19177
FAE2  C9 B8     CMP #$B8                        A:B8 X:2 Y:B5 P:A5 SP:F9 CYC:19179
FAE4  F0 02     BEQ $FAE8                       A:B8 X:2 Y:B5 P:27 SP:F9 CYC:19181
FAE8  60        RTS                             A:B8 X:2 Y:B5 P:27 SP:F9 CYC:19184
ECE4  AD 47 06  LDA $0647 = 38                  A:B8 X:2 Y:B5 P:27 SP:FB CYC:19190
ECE7  C9 38     CMP #$38                        A:38 X:2 Y:B5 P:25 SP:FB CYC:19194
ECE9  F0 02     BEQ $ECED                       A:38 X:2 Y:B5 P:27 SP:FB CYC:19196
ECED  A0 B6     LDY #$B6                        A:38 X:2 Y:B5 P:27 SP:FB CYC:19199
ECEF  A2 FF     LDX #$FF                        A:38 X:2 Y:B6 P:A5 SP:FB CYC:19201
ECF1  A9 EB     LDA #$EB                        A:38 X:FF Y:B6 P:A5 SP:FB CYC:19203
ECF3  85 47     STA $47 = 38                    A:EB X:FF Y:B6 P:A5 SP:FB CYC:19205
ECF5  20 B1 FA  JSR $FAB1                       A:EB X:FF Y:B6 P:A5 SP:FB CYC:19208
FAB1  24 01     BIT $01 = FF                    A:EB X:FF Y:B6 P:A5 SP:F9 CYC:19214
FAB3  18        CLC                             A:EB X:FF Y:B6 P:E5 SP:F9 CYC:19217
FAB4  A9 40     LDA #$40                        A:EB X:FF Y:B6 P:E4 SP:F9 CYC:19219
FAB6  60        RTS                             A:40 X:FF Y:B6 P:64 SP:F9 CYC:19221
ECF8  F7 48    *ISB $48,X @ 47 = EB             A:40 X:FF Y:B6 P:64 SP:FB CYC:19227
ECFA  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB CYC:19233
ECFB  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB CYC:19235
ECFC  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB CYC:19237
ECFD  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB CYC:19239
ECFE  20 B7 FA  JSR $FAB7                       A:53 X:FF Y:B6 P:24 SP:FB CYC:19241
FAB7  70 2D     BVS $FAE6                       A:53 X:FF Y:B6 P:24 SP:F9 CYC:19247
FAB9  B0 2B     BCS $FAE6                       A:53 X:FF Y:B6 P:24 SP:F9 CYC:19249
FABB  30 29     BMI $FAE6                       A:53 X:FF Y:B6 P:24 SP:F9 CYC:19251
FABD  C9 53     CMP #$53                        A:53 X:FF Y:B6 P:24 SP:F9 CYC:19253
FABF  D0 25     BNE $FAE6                       A:53 X:FF Y:B6 P:27 SP:F9 CYC:19255
FAC1  60        RTS                             A:53 X:FF Y:B6 P:27 SP:F9 CYC:19257
ED01  A5 47     LDA $47 = EC                    A:53 X:FF Y:B6 P:27 SP:FB CYC:19263
ED03  C9 EC     CMP #$EC                        A:EC X:FF Y:B6 P:A5 SP:FB CYC:19266
ED05  F0 02     BEQ $ED09                       A:EC X:FF Y:B6 P:27 SP:FB CYC:19268
ED09  C8        INY                             A:EC X:FF Y:B6 P:27 SP:FB CYC:19271
ED0A  A9 FF     LDA #$FF                        A:EC X:FF Y:B7 P:A5 SP:FB CYC:19273
ED0C  85 47     STA $47 = EC                    A:FF X:FF Y:B7 P:A5 SP:FB CYC:19275
ED0E  20 C2 FA  JSR $FAC2                       A:FF X:FF Y:B7 P:A5 SP:FB CYC:19278
FAC2  B8        CLV                             A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19284
FAC3  38        SEC                             A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19286
FAC4  A9 FF     LDA #$FF                        A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19288
FAC6  60        RTS                             A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19290
ED11  F7 48    *ISB $48,X @ 47 = FF             A:FF X:FF Y:B7 P:A5 SP:FB CYC:19296
ED13  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB CYC:19302
ED14  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB CYC:19304
ED15  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB CYC:19306
ED16  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB CYC:19308
ED17  20 C7 FA  JSR $FAC7                       A:FF X:FF Y:B7 P:A5 SP:FB CYC:19310
FAC7  70 1D     BVS $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19316
FAC9  F0 1B     BEQ $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19318
FACB  10 19     BPL $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19320
FACD  90 17     BCC $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19322
FACF  C9 FF     CMP #$FF                        A:FF X:FF Y:B7 P:A5 SP:F9 CYC:19324
FAD1  D0 13     BNE $FAE6                       A:FF X:FF Y:B7 P:27 SP:F9 CYC:19326
FAD3  60        RTS                             A:FF X:FF Y:B7 P:27 SP:F9 CYC:19328
ED1A  A5 47     LDA $47 = 00                    A:FF X:FF Y:B7 P:27 SP:FB CYC:19334
ED1C  C9 00     CMP #$00                        A:0 X:FF Y:B7 P:27 SP:FB CYC:19337
ED1E  F0 02     BEQ $ED22                       A:0 X:FF Y:B7 P:27 SP:FB CYC:19339
ED22  C8        INY                             A:0 X:FF Y:B7 P:27 SP:FB CYC:19342
ED23  A9 37     LDA #$37                        A:0 X:FF Y:B8 P:A5 SP:FB CYC:19344
ED25  85 47     STA $47 = 00                    A:37 X:FF Y:B8 P:25 SP:FB CYC:19346
ED27  20 D4 FA  JSR $FAD4                       A:37 X:FF Y:B8 P:25 SP:FB CYC:19349
FAD4  24 01     BIT $01 = FF                    A:37 X:FF Y:B8 P:25 SP:F9 CYC:19355
FAD6  38        SEC                             A:37 X:FF Y:B8 P:E5 SP:F9 CYC:19358
FAD7  A9 F0     LDA #$F0                        A:37 X:FF Y:B8 P:E5 SP:F9 CYC:19360
FAD9  60        RTS                             A:F0 X:FF Y:B8 P:E5 SP:F9 CYC:19362
ED2A  F7 48    *ISB $48,X @ 47 = 37             A:F0 X:FF Y:B8 P:E5 SP:FB CYC:19368
ED2C  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB CYC:19374
ED2D  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB CYC:19376
ED2E  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB CYC:19378
ED2F  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB CYC:19380
ED30  20 DA FA  JSR $FADA                       A:B8 X:FF Y:B8 P:A5 SP:FB CYC:19382
FADA  70 0A     BVS $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9 CYC:19388
FADC  F0 08     BEQ $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9 CYC:19390
FADE  10 06     BPL $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9 CYC:19392
FAE0  90 04     BCC $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9 CYC:19394
FAE2  C9 B8     CMP #$B8                        A:B8 X:FF Y:B8 P:A5 SP:F9 CYC:19396
FAE4  F0 02     BEQ $FAE8                       A:B8 X:FF Y:B8 P:27 SP:F9 CYC:19398
FAE8  60        RTS                             A:B8 X:FF Y:B8 P:27 SP:F9 CYC:19401
ED33  A5 47     LDA $47 = 38                    A:B8 X:FF Y:B8 P:27 SP:FB CYC:19407
ED35  C9 38     CMP #$38                        A:38 X:FF Y:B8 P:25 SP:FB CYC:19410
ED37  F0 02     BEQ $ED3B                       A:38 X:FF Y:B8 P:27 SP:FB CYC:19412
ED3B  A9 EB     LDA #$EB                        A:38 X:FF Y:B8 P:27 SP:FB CYC:19415
ED3D  8D 47 06  STA $0647 = 38                  A:EB X:FF Y:B8 P:A5 SP:FB CYC:19417
ED40  A0 FF     LDY #$FF                        A:EB X:FF Y:B8 P:A5 SP:FB CYC:19421
ED42  20 B1 FA  JSR $FAB1                       A:EB X:FF Y:FF P:A5 SP:FB CYC:19423
FAB1  24 01     BIT $01 = FF                    A:EB X:FF Y:FF P:A5 SP:F9 CYC:19429
FAB3  18        CLC                             A:EB X:FF Y:FF P:E5 SP:F9 CYC:19432
FAB4  A9 40     LDA #$40                        A:EB X:FF Y:FF P:E4 SP:F9 CYC:19434
FAB6  60        RTS                             A:40 X:FF Y:FF P:64 SP:F9 CYC:19436
ED45  FB 48 05 *ISB $0548,Y @ 0647 = EB         A:40 X:FF Y:FF P:64 SP:FB CYC:19442
ED48  EA        NOP                             A:53 X:FF Y:FF P:24 SP:FB CYC:19449
ED49  EA        NOP                             A:53 X:FF Y:FF P:24 SP:FB CYC:19451
ED4A  08        PHP                             A:53 X:FF Y:FF P:24 SP:FB CYC:19453
ED4B  48        PHA                             A:53 X:FF Y:FF P:24 SP:FA CYC:19456
ED4C  A0 B9     LDY #$B9                        A:53 X:FF Y:FF P:24 SP:F9 CYC:19459
ED4E  68        PLA                             A:53 X:FF Y:B9 P:A4 SP:F9 CYC:19461
ED4F  28        PLP                             A:53 X:FF Y:B9 P:24 SP:FA CYC:19465
ED50  20 B7 FA  JSR $FAB7                       A:53 X:FF Y:B9 P:24 SP:FB CYC:19469
FAB7  70 2D     BVS $FAE6                       A:53 X:FF Y:B9 P:24 SP:F9 CYC:19475
FAB9  B0 2B     BCS $FAE6                       A:53 X:FF Y:B9 P:24 SP:F9 CYC:19477
FABB  30 29     BMI $FAE6                       A:53 X:FF Y:B9 P:24 SP:F9 CYC:19479
FABD  C9 53     CMP #$53                        A:53 X:FF Y:B9 P:24 SP:F9 CYC:19481
FABF  D0 25     BNE $FAE6                       A:53 X:FF Y:B9 P:27 SP:F9 CYC:19483
FAC1  60        RTS                             A:53 X:FF Y:B9 P:27 SP:F9 CYC:19485
ED53  AD 47 06  LDA $0647 = EC                  A:53 X:FF Y:B9 P:27 SP:FB CYC:19491
ED56  C9 EC     CMP #$EC                        A:EC X:FF Y:B9 P:A5 SP:FB CYC:19495
ED58  F0 02     BEQ $ED5C                       A:EC X:FF Y:B9 P:27 SP:FB CYC:19497
ED5C  A0 FF     LDY #$FF                        A:EC X:FF Y:B9 P:27 SP:FB CYC:19500
ED5E  A9 FF     LDA #$FF                        A:EC X:FF Y:FF P:A5 SP:FB CYC:19502
ED60  8D 47 06  STA $0647 = EC                  A:FF X:FF Y:FF P:A5 SP:FB CYC:19504
ED63  20 C2 FA  JSR $FAC2                       A:FF X:FF Y:FF P:A5 SP:FB CYC:19508
FAC2  B8        CLV                             A:FF X:FF Y:FF P:A5 SP:F9 CYC:19514
FAC3  38        SEC                             A:FF X:FF Y:FF P:A5 SP:F9 CYC:19516
FAC4  A9 FF     LDA #$FF                        A:FF X:FF Y:FF P:A5 SP:F9 CYC:19518
FAC6  60        RTS                             A:FF X:FF Y:FF P:A5 SP:F9 CYC:19520
ED66  FB 48 05 *ISB $0548,Y @ 0647 = FF         A:FF X:FF Y:FF P:A5 SP:FB CYC:19526
ED69  EA        NOP                             A:FF X:FF Y:FF P:A5 SP:FB CYC:19533
ED6A  EA        NOP                             A:FF X:FF Y:FF P:A5 SP:FB CYC:19535
ED6B  08        PHP                             A:FF X:FF Y:FF P:A5 SP:FB CYC:19537
ED6C  48        PHA                             A:FF X:FF Y:FF P:A5 SP:FA CYC:19540
ED6D  A0 BA     LDY #$BA                        A:FF X:FF Y:FF P:A5 SP:F9 CYC:19543
ED6F  68        PLA                             A:FF X:FF Y:BA P:A5 SP:F9 CYC:19545
ED70  28        PLP                             A:FF X:FF Y:BA P:A5 SP:FA CYC:19549
ED71  20 C7 FA  JSR $FAC7                       A:FF X:FF Y:BA P:A5 SP:FB CYC:19553
FAC7  70 1D     BVS $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9 CYC:19559
FAC9  F0 1B     BEQ $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9 CYC:19561
FACB  10 19     BPL $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9 CYC:19563
FACD  90 17     BCC $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9 CYC:19565
FACF  C9 FF     CMP #$FF                        A:FF X:FF Y:BA P:A5 SP:F9 CYC:19567
FAD1  D0 13     BNE $FAE6                       A:FF X:FF Y:BA P:27 SP:F9 CYC:19569
FAD3  60        RTS                             A:FF X:FF Y:BA P:27 SP:F9 CYC:19571
ED74  AD 47 06  LDA $0647 = 00                  A:FF X:FF Y:BA P:27 SP:FB CYC:19577
ED77  C9 00     CMP #$00                        A:0 X:FF Y:BA P:27 SP:FB CYC:19581
ED79  F0 02     BEQ $ED7D                       A:0 X:FF Y:BA P:27 SP:FB CYC:19583
ED7D  A0 FF     LDY #$FF                        A:0 X:FF Y:BA P:27 SP:FB CYC:19586
ED7F  A9 37     LDA #$37                        A:0 X:FF Y:FF P:A5 SP:FB CYC:19588
ED81  8D 47 06  STA $0647 = 00                  A:37 X:FF Y:FF P:25 SP:FB CYC:19590
ED84  20 D4 FA  JSR $FAD4                       A:37 X:FF Y:FF P:25 SP:FB CYC:19594
FAD4  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9 CYC:19600
FAD6  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9 CYC:19603
FAD7  A9 F0     LDA #$F0                        A:37 X:FF Y:FF P:E5 SP:F9 CYC:19605
FAD9  60        RTS                             A:F0 X:FF Y:FF P:E5 SP:F9 CYC:19607
ED87  FB 48 05 *ISB $0548,Y @ 0647 = 37         A:F0 X:FF Y:FF P:E5 SP:FB CYC:19613
ED8A  EA        NOP                             A:B8 X:FF Y:FF P:A5 SP:FB CYC:19620
ED8B  EA        NOP                             A:B8 X:FF Y:FF P:A5 SP:FB CYC:19622
ED8C  08        PHP                             A:B8 X:FF Y:FF P:A5 SP:FB CYC:19624
ED8D  48        PHA                             A:B8 X:FF Y:FF P:A5 SP:FA CYC:19627
ED8E  A0 BB     LDY #$BB                        A:B8 X:FF Y:FF P:A5 SP:F9 CYC:19630
ED90  68        PLA                             A:B8 X:FF Y:BB P:A5 SP:F9 CYC:19632
ED91  28        PLP                             A:B8 X:FF Y:BB P:A5 SP:FA CYC:19636
ED92  20 DA FA  JSR $FADA                       A:B8 X:FF Y:BB P:A5 SP:FB CYC:19640
FADA  70 0A     BVS $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9 CYC:19646
FADC  F0 08     BEQ $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9 CYC:19648
FADE  10 06     BPL $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9 CYC:19650
FAE0  90 04     BCC $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9 CYC:19652
FAE2  C9 B8     CMP #$B8                        A:B8 X:FF Y:BB P:A5 SP:F9 CYC:19654
FAE4  F0 02     BEQ $FAE8                       A:B8 X:FF Y:BB P:27 SP:F9 CYC:19656
FAE8  60        RTS                             A:B8 X:FF Y:BB P:27 SP:F9 CYC:19659
ED95  AD 47 06  LDA $0647 = 38                  A:B8 X:FF Y:BB P:27 SP:FB CYC:19665
ED98  C9 38     CMP #$38                        A:38 X:FF Y:BB P:25 SP:FB CYC:19669
ED9A  F0 02     BEQ $ED9E                       A:38 X:FF Y:BB P:27 SP:FB CYC:19671
ED9E  A0 BC     LDY #$BC                        A:38 X:FF Y:BB P:27 SP:FB CYC:19674
EDA0  A2 FF     LDX #$FF                        A:38 X:FF Y:BC P:A5 SP:FB CYC:19676
EDA2  A9 EB     LDA #$EB                        A:38 X:FF Y:BC P:A5 SP:FB CYC:19678
EDA4  8D 47 06  STA $0647 = 38                  A:EB X:FF Y:BC P:A5 SP:FB CYC:19680
EDA7  20 B1 FA  JSR $FAB1                       A:EB X:FF Y:BC P:A5 SP:FB CYC:19684
FAB1  24 01     BIT $01 = FF                    A:EB X:FF Y:BC P:A5 SP:F9 CYC:19690
FAB3  18        CLC                             A:EB X:FF Y:BC P:E5 SP:F9 CYC:19693
FAB4  A9 40     LDA #$40                        A:EB X:FF Y:BC P:E4 SP:F9 CYC:19695
FAB6  60        RTS                             A:40 X:FF Y:BC P:64 SP:F9 CYC:19697
EDAA  FF 48 05 *ISB $0548,X @ 0647 = EB         A:40 X:FF Y:BC P:64 SP:FB CYC:19703
EDAD  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB CYC:19710
EDAE  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB CYC:19712
EDAF  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB CYC:19714
EDB0  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB CYC:19716
EDB1  20 B7 FA  JSR $FAB7                       A:53 X:FF Y:BC P:24 SP:FB CYC:19718
FAB7  70 2D     BVS $FAE6                       A:53 X:FF Y:BC P:24 SP:F9 CYC:19724
FAB9  B0 2B     BCS $FAE6                       A:53 X:FF Y:BC P:24 SP:F9 CYC:19726
FABB  30 29     BMI $FAE6                       A:53 X:FF Y:BC P:24 SP:F9 CYC:19728
FABD  C9 53     CMP #$53                        A:53 X:FF Y:BC P:24 SP:F9 CYC:19730
FABF  D0 25     BNE $FAE6                       A:53 X:FF Y:BC P:27 SP:F9 CYC:19732
FAC1  60        RTS                             A:53 X:FF Y:BC P:27 SP:F9 CYC:19734
EDB4  AD 47 06  LDA $0647 = EC                  A:53 X:FF Y:BC P:27 SP:FB CYC:19740
EDB7  C9 EC     CMP #$EC                        A:EC X:FF Y:BC P:A5 SP:FB CYC:19744
EDB9  F0 02     BEQ $EDBD                       A:EC X:FF Y:BC P:27 SP:FB CYC:19746
EDBD  C8        INY                             A:EC X:FF Y:BC P:27 SP:FB CYC:19749
EDBE  A9 FF     LDA #$FF                        A:EC X:FF Y:BD P:A5 SP:FB CYC:19751
EDC0  8D 47 06  STA $0647 = EC                  A:FF X:FF Y:BD P:A5 SP:FB CYC:19753
EDC3  20 C2 FA  JSR $FAC2                       A:FF X:FF Y:BD P:A5 SP:FB CYC:19757
FAC2  B8        CLV                             A:FF X:FF Y:BD P:A5 SP:F9 CYC:19763
FAC3  38        SEC                             A:FF X:FF Y:BD P:A5 SP:F9 CYC:19765
FAC4  A9 FF     LDA #$FF                        A:FF X:FF Y:BD P:A5 SP:F9 CYC:19767
FAC6  60        RTS                             A:FF X:FF Y:BD P:A5 SP:F9 CYC:19769
EDC6  FF 48 05 *ISB $0548,X @ 0647 = FF         A:FF X:FF Y:BD P:A5 SP:FB CYC:19775
EDC9  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB CYC:19782
EDCA  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB CYC:19784
EDCB  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB CYC:19786
EDCC  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB CYC:19788
EDCD  20 C7 FA  JSR $FAC7                       A:FF X:FF Y:BD P:A5 SP:FB CYC:19790
FAC7  70 1D     BVS $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9 CYC:19796
FAC9  F0 1B     BEQ $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9 CYC:19798
FACB  10 19     BPL $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9 CYC:19800
FACD  90 17     BCC $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9 CYC:19802
FACF  C9 FF     CMP #$FF                        A:FF X:FF Y:BD P:A5 SP:F9 CYC:19804
FAD1  D0 13     BNE $FAE6                       A:FF X:FF Y:BD P:27 SP:F9 CYC:19806
FAD3  60        RTS                             A:FF X:FF Y:BD P:27 SP:F9 CYC:19808
EDD0  AD 47 06  LDA $0647 = 00                  A:FF X:FF Y:BD P:27 SP:FB CYC:19814
EDD3  C9 00     CMP #$00                        A:0 X:FF Y:BD P:27 SP:FB CYC:19818
EDD5  F0 02     BEQ $EDD9                       A:0 X:FF Y:BD P:27 SP:FB CYC:19820
EDD9  C8        INY                             A:0 X:FF Y:BD P:27 SP:FB CYC:19823
EDDA  A9 37     LDA #$37                        A:0 X:FF Y:BE P:A5 SP:FB CYC:19825
EDDC  8D 47 06  STA $0647 = 00                  A:37 X:FF Y:BE P:25 SP:FB CYC:19827
EDDF  20 D4 FA  JSR $FAD4                       A:37 X:FF Y:BE P:25 SP:FB CYC:19831
FAD4  24 01     BIT $01 = FF                    A:37 X:FF Y:BE P:25 SP:F9 CYC:19837
FAD6  38        SEC                             A:37 X:FF Y:BE P:E5 SP:F9 CYC:19840
FAD7  A9 F0     LDA #$F0                        A:37 X:FF Y:BE P:E5 SP:F9 CYC:19842
FAD9  60        RTS                             A:F0 X:FF Y:BE P:E5 SP:F9 CYC:19844
EDE2  FF 48 05 *ISB $0548,X @ 0647 = 37         A:F0 X:FF Y:BE P:E5 SP:FB CYC:19850
EDE5  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB CYC:19857
EDE6  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB CYC:19859
EDE7  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB CYC:19861
EDE8  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB CYC:19863
EDE9  20 DA FA  JSR $FADA                       A:B8 X:FF Y:BE P:A5 SP:FB CYC:19865
FADA  70 0A     BVS $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9 CYC:19871
FADC  F0 08     BEQ $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9 CYC:19873
FADE  10 06     BPL $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9 CYC:19875
FAE0  90 04     BCC $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9 CYC:19877
FAE2  C9 B8     CMP #$B8                        A:B8 X:FF Y:BE P:A5 SP:F9 CYC:19879
FAE4  F0 02     BEQ $FAE8                       A:B8 X:FF Y:BE P:27 SP:F9 CYC:19881
FAE8  60        RTS                             A:B8 X:FF Y:BE P:27 SP:F9 CYC:19884
EDEC  AD 47 06  LDA $0647 = 38                  A:B8 X:FF Y:BE P:27 SP:FB CYC:19890
EDEF  C9 38     CMP #$38                        A:38 X:FF Y:BE P:25 SP:FB CYC:19894
EDF1  F0 02     BEQ $EDF5                       A:38 X:FF Y:BE P:27 SP:FB CYC:19896
EDF5  60        RTS                             A:38 X:FF Y:BE P:27 SP:FB CYC:19899
C641  20 F6 ED  JSR $EDF6                       A:38 X:FF Y:BE P:27 SP:FD CYC:19905
EDF6  A9 FF     LDA #$FF                        A:38 X:FF Y:BE P:27 SP:FB CYC:19911
EDF8  85 01     STA $01 = FF                    A:FF X:FF Y:BE P:A5 SP:FB CYC:19913
EDFA  A0 BF     LDY #$BF                        A:FF X:FF Y:BE P:A5 SP:FB CYC:19916
EDFC  A2 02     LDX #$02                        A:FF X:FF Y:BF P:A5 SP:FB CYC:19918
EDFE  A9 47     LDA #$47                        A:FF X:2 Y:BF P:25 SP:FB CYC:19920
EE00  85 47     STA $47 = 38                    A:47 X:2 Y:BF P:25 SP:FB CYC:19922
EE02  A9 06     LDA #$06                        A:47 X:2 Y:BF P:25 SP:FB CYC:19925
EE04  85 48     STA $48 = 06                    A:6 X:2 Y:BF P:25 SP:FB CYC:19927
EE06  A9 A5     LDA #$A5                        A:6 X:2 Y:BF P:25 SP:FB CYC:19930
EE08  8D 47 06  STA $0647 = 38                  A:A5 X:2 Y:BF P:A5 SP:FB CYC:19932
EE0B  20 7B FA  JSR $FA7B                       A:A5 X:2 Y:BF P:A5 SP:FB CYC:19936
FA7B  24 01     BIT $01 = FF                    A:A5 X:2 Y:BF P:A5 SP:F9 CYC:19942
FA7D  18        CLC                             A:A5 X:2 Y:BF P:E5 SP:F9 CYC:19945
FA7E  A9 B3     LDA #$B3                        A:A5 X:2 Y:BF P:E4 SP:F9 CYC:19947
FA80  60        RTS                             A:B3 X:2 Y:BF P:E4 SP:F9 CYC:19949
EE0E  03 45    *SLO ($45,X) @ 47 = 0647 = A5    A:B3 X:2 Y:BF P:E4 SP:FB CYC:19955
EE10  EA        NOP                             A:FB X:2 Y:BF P:E5 SP:FB CYC:19963
EE11  EA        NOP                             A:FB X:2 Y:BF P:E5 SP:FB CYC:19965
EE12  EA        NOP                             A:FB X:2 Y:BF P:E5 SP:FB CYC:19967
EE13  EA        NOP                             A:FB X:2 Y:BF P:E5 SP:FB CYC:19969
EE14  20 81 FA  JSR $FA81                       A:FB X:2 Y:BF P:E5 SP:FB CYC:19971
FA81  50 63     BVC $FAE6                       A:FB X:2 Y:BF P:E5 SP:F9 CYC:19977
FA83  90 61     BCC $FAE6                       A:FB X:2 Y:BF P:E5 SP:F9 CYC:19979
FA85  10 5F     BPL $FAE6                       A:FB X:2 Y:BF P:E5 SP:F9 CYC:19981
FA87  C9 FB     CMP #$FB                        A:FB X:2 Y:BF P:E5 SP:F9 CYC:19983
FA89  D0 5B     BNE $FAE6                       A:FB X:2 Y:BF P:67 SP:F9 CYC:19985
FA8B  60        RTS                             A:FB X:2 Y:BF P:67 SP:F9 CYC:19987
EE17  AD 47 06  LDA $0647 = 4A                  A:FB X:2 Y:BF P:67 SP:FB CYC:19993
EE1A  C9 4A     CMP #$4A                        A:4A X:2 Y:BF P:65 SP:FB CYC:19997
EE1C  F0 02     BEQ $EE20                       A:4A X:2 Y:BF P:67 SP:FB CYC:19999
EE20  C8        INY                             A:4A X:2 Y:BF P:67 SP:FB CYC:20002
EE21  A9 29     LDA #$29                        A:4A X:2 Y:C0 P:E5 SP:FB CYC:20004
EE23  8D 47 06  STA $0647 = 4A                  A:29 X:2 Y:C0 P:65 SP:FB CYC:20006
EE26  20 8C FA  JSR $FA8C                       A:29 X:2 Y:C0 P:65 SP:FB CYC:20010
FA8C  B8        CLV                             A:29 X:2 Y:C0 P:65 SP:F9 CYC:20016
FA8D  18        CLC                             A:29 X:2 Y:C0 P:25 SP:F9 CYC:20018
FA8E  A9 C3     LDA #$C3                        A:29 X:2 Y:C0 P:24 SP:F9 CYC:20020
FA90  60        RTS                             A:C3 X:2 Y:C0 P:A4 SP:F9 CYC:20022
EE29  03 45    *SLO ($45,X) @ 47 = 0647 = 29    A:C3 X:2 Y:C0 P:A4 SP:FB CYC:20028
EE2B  EA        NOP                             A:D3 X:2 Y:C0 P:A4 SP:FB CYC:20036
EE2C  EA        NOP                             A:D3 X:2 Y:C0 P:A4 SP:FB CYC:20038
EE2D  EA        NOP                             A:D3 X:2 Y:C0 P:A4 SP:FB CYC:20040
EE2E  EA        NOP                             A:D3 X:2 Y:C0 P:A4 SP:FB CYC:20042
EE2F  20 91 FA  JSR $FA91                       A:D3 X:2 Y:C0 P:A4 SP:FB CYC:20044
FA91  70 53     BVS $FAE6                       A:D3 X:2 Y:C0 P:A4 SP:F9 CYC:20050
FA93  F0 51     BEQ $FAE6                       A:D3 X:2 Y:C0 P:A4 SP:F9 CYC:20052
FA95  10 4F     BPL $FAE6                       A:D3 X:2 Y:C0 P:A4 SP:F9 CYC:20054
FA97  B0 4D     BCS $FAE6                       A:D3 X:2 Y:C0 P:A4 SP:F9 CYC:20056
FA99  C9 D3     CMP #$D3                        A:D3 X:2 Y:C0 P:A4 SP:F9 CYC:20058
FA9B  D0 49     BNE $FAE6                       A:D3 X:2 Y:C0 P:27 SP:F9 CYC:20060
FA9D  60        RTS                             A:D3 X:2 Y:C0 P:27 SP:F9 CYC:20062
EE32  AD 47 06  LDA $0647 = 52                  A:D3 X:2 Y:C0 P:27 SP:FB CYC:20068
EE35  C9 52     CMP #$52                        A:52 X:2 Y:C0 P:25 SP:FB CYC:20072
EE37  F0 02     BEQ $EE3B                       A:52 X:2 Y:C0 P:27 SP:FB CYC:20074
EE3B  C8        INY                             A:52 X:2 Y:C0 P:27 SP:FB CYC:20077
EE3C  A9 37     LDA #$37                        A:52 X:2 Y:C1 P:A5 SP:FB CYC:20079
EE3E  8D 47 06  STA $0647 = 52                  A:37 X:2 Y:C1 P:25 SP:FB CYC:20081
EE41  20 9E FA  JSR $FA9E                       A:37 X:2 Y:C1 P:25 SP:FB CYC:20085
FA9E  24 01     BIT $01 = FF                    A:37 X:2 Y:C1 P:25 SP:F9 CYC:20091
FAA0  38        SEC                             A:37 X:2 Y:C1 P:E5 SP:F9 CYC:20094
FAA1  A9 10     LDA #$10                        A:37 X:2 Y:C1 P:E5 SP:F9 CYC:20096
FAA3  60        RTS                             A:10 X:2 Y:C1 P:65 SP:F9 CYC:20098
EE44  03 45    *SLO ($45,X) @ 47 = 0647 = 37    A:10 X:2 Y:C1 P:65 SP:FB CYC:20104
EE46  EA        NOP                             A:7E X:2 Y:C1 P:64 SP:FB CYC:20112
EE47  EA        NOP                             A:7E X:2 Y:C1 P:64 SP:FB CYC:20114
EE48  EA        NOP                             A:7E X:2 Y:C1 P:64 SP:FB CYC:20116
EE49  EA        NOP                             A:7E X:2 Y:C1 P:64 SP:FB CYC:20118
EE4A  20 A4 FA  JSR $FAA4                       A:7E X:2 Y:C1 P:64 SP:FB CYC:20120
FAA4  50 40     BVC $FAE6                       A:7E X:2 Y:C1 P:64 SP:F9 CYC:20126
FAA6  F0 3E     BEQ $FAE6                       A:7E X:2 Y:C1 P:64 SP:F9 CYC:20128
FAA8  30 3C     BMI $FAE6                       A:7E X:2 Y:C1 P:64 SP:F9 CYC:20130
FAAA  B0 3A     BCS $FAE6                       A:7E X:2 Y:C1 P:64 SP:F9 CYC:20132
FAAC  C9 7E     CMP #$7E                        A:7E X:2 Y:C1 P:64 SP:F9 CYC:20134
FAAE  D0 36     BNE $FAE6                       A:7E X:2 Y:C1 P:67 SP:F9 CYC:20136
FAB0  60        RTS                             A:7E X:2 Y:C1 P:67 SP:F9 CYC:20138
EE4D  AD 47 06  LDA $0647 = 6E                  A:7E X:2 Y:C1 P:67 SP:FB CYC:20144
EE50  C9 6E     CMP #$6E                        A:6E X:2 Y:C1 P:65 SP:FB CYC:20148
EE52  F0 02     BEQ $EE56                       A:6E X:2 Y:C1 P:67 SP:FB CYC:20150
EE56  C8        INY                             A:6E X:2 Y:C1 P:67 SP:FB CYC:20153
EE57  A9 A5     LDA #$A5                        A:6E X:2 Y:C2 P:E5 SP:FB CYC:20155
EE59  85 47     STA $47 = 47                    A:A5 X:2 Y:C2 P:E5 SP:FB CYC:20157
EE5B  20 7B FA  JSR $FA7B                       A:A5 X:2 Y:C2 P:E5 SP:FB CYC:20160
FA7B  24 01     BIT $01 = FF                    A:A5 X:2 Y:C2 P:E5 SP:F9 CYC:20166
FA7D  18        CLC                             A:A5 X:2 Y:C2 P:E5 SP:F9 CYC:20169
FA7E  A9 B3     LDA #$B3                        A:A5 X:2 Y:C2 P:E4 SP:F9 CYC:20171
FA80  60        RTS                             A:B3 X:2 Y:C2 P:E4 SP:F9 CYC:20173
EE5E  07 47    *SLO $47 = A5                    A:B3 X:2 Y:C2 P:E4 SP:FB CYC:20179
EE60  EA        NOP                             A:FB X:2 Y:C2 P:E5 SP:FB CYC:20184
EE61  EA        NOP                             A:FB X:2 Y:C2 P:E5 SP:FB CYC:20186
EE62  EA        NOP                             A:FB X:2 Y:C2 P:E5 SP:FB CYC:20188
EE63  EA        NOP                             A:FB X:2 Y:C2 P:E5 SP:FB CYC:20190
EE64  20 81 FA  JSR $FA81                       A:FB X:2 Y:C2 P:E5 SP:FB CYC:20192
FA81  50 63     BVC $FAE6                       A:FB X:2 Y:C2 P:E5 SP:F9 CYC:20198
FA83  90 61     BCC $FAE6                       A:FB X:2 Y:C2 P:E5 SP:F9 CYC:20200
FA85  10 5F     BPL $FAE6                       A:FB X:2 Y:C2 P:E5 SP:F9 CYC:20202
FA87  C9 FB     CMP #$FB                        A:FB X:2 Y:C2 P:E5 SP:F9 CYC:20204
FA89  D0 5B     BNE $FAE6                       A:FB X:2 Y:C2 P:67 SP:F9 CYC:20206
FA8B  60        RTS                             A:FB X:2 Y:C2 P:67 SP:F9 CYC:20208
EE67  A5 47     LDA $47 = 4A                    A:FB X:2 Y:C2 P:67 SP:FB CYC:20214
EE69  C9 4A     CMP #$4A                        A:4A X:2 Y:C2 P:65 SP:FB CYC:20217
EE6B  F0 02     BEQ $EE6F                       A:4A X:2 Y:C2 P:67 SP:FB CYC:20219
EE6F  C8        INY                             A:4A X:2 Y:C2 P:67 SP:FB CYC:20222
EE70  A9 29     LDA #$29                        A:4A X:2 Y:C3 P:E5 SP:FB CYC:20224
EE72  85 47     STA $47 = 4A                    A:29 X:2 Y:C3 P:65 SP:FB CYC:20226
EE74  20 8C FA  JSR $FA8C                       A:29 X:2 Y:C3 P:65 SP:FB CYC:20229
FA8C  B8        CLV                             A:29 X:2 Y:C3 P:65 SP:F9 CYC:20235
FA8D  18        CLC                             A:29 X:2 Y:C3 P:25 SP:F9 CYC:20237
FA8E  A9 C3     LDA #$C3                        A:29 X:2 Y:C3 P:24 SP:F9 CYC:20239
FA90  60        RTS                             A:C3 X:2 Y:C3 P:A4 SP:F9 CYC:20241
EE77  07 47    *SLO $47 = 29                    A:C3 X:2 Y:C3 P:A4 SP:FB CYC:20247
EE79  EA        NOP                             A:D3 X:2 Y:C3 P:A4 SP:FB CYC:20252
EE7A  EA        NOP                             A:D3 X:2 Y:C3 P:A4 SP:FB CYC:20254
EE7B  EA        NOP                             A:D3 X:2 Y:C3 P:A4 SP:FB CYC:20256
EE7C  EA        NOP                             A:D3 X:2 Y:C3 P:A4 SP:FB CYC:20258
EE7D  20 91 FA  JSR $FA91                       A:D3 X:2 Y:C3 P:A4 SP:FB CYC:20260
FA91  70 53     BVS $FAE6                       A:D3 X:2 Y:C3 P:A4 SP:F9 CYC:20266
FA93  F0 51     BEQ $FAE6                       A:D3 X:2 Y:C3 P:A4 SP:F9 CYC:20268
FA95  10 4F     BPL $FAE6                       A:D3 X:2 Y:C3 P:A4 SP:F9 CYC:20270
FA97  B0 4D     BCS $FAE6                       A:D3 X:2 Y:C3 P:A4 SP:F9 CYC:20272
FA99  C9 D3     CMP #$D3                        A:D3 X:2 Y:C3 P:A4 SP:F9 CYC:20274
FA9B  D0 49     BNE $FAE6                       A:D3 X:2 Y:C3 P:27 SP:F9 CYC:20276
FA9D  60        RTS                             A:D3 X:2 Y:C3 P:27 SP:F9 CYC:20278
EE80  A5 47     LDA $47 = 52                    A:D3 X:2 Y:C3 P:27 SP:FB CYC:20284
EE82  C9 52     CMP #$52                        A:52 X:2 Y:C3 P:25 SP:FB CYC:20287
EE84  F0 02     BEQ $EE88                       A:52 X:2 Y:C3 P:27 SP:FB CYC:20289
EE88  C8        INY                             A:52 X:2 Y:C3 P:27 SP:FB CYC:20292
EE89  A9 37     LDA #$37                        A:52 X:2 Y:C4 P:A5 SP:FB CYC:20294
EE8B  85 47     STA $47 = 52                    A:37 X:2 Y:C4 P:25 SP:FB CYC:20296
EE8D  20 9E FA  JSR $FA9E                       A:37 X:2 Y:C4 P:25 SP:FB CYC:20299
FA9E  24 01     BIT $01 = FF                    A:37 X:2 Y:C4 P:25 SP:F9 CYC:20305
FAA0  38        SEC                             A:37 X:2 Y:C4 P:E5 SP:F9 CYC:20308
FAA1  A9 10     LDA #$10                        A:37 X:2 Y:C4 P:E5 SP:F9 CYC:20310
FAA3  60        RTS                             A:10 X:2 Y:C4 P:65 SP:F9 CYC:20312
EE90  07 47    *SLO $47 = 37                    A:10 X:2 Y:C4 P:65 SP:FB CYC:20318
EE92  EA        NOP                             A:7E X:2 Y:C4 P:64 SP:FB CYC:20323
EE93  EA        NOP                             A:7E X:2 Y:C4 P:64 SP:FB CYC:20325
EE94  EA        NOP                             A:7E X:2 Y:C4 P:64 SP:FB CYC:20327
EE95  EA        NOP                             A:7E X:2 Y:C4 P:64 SP:FB CYC:20329
EE96  20 A4 FA  JSR $FAA4                       A:7E X:2 Y:C4 P:64 SP:FB CYC:20331
FAA4  50 40     BVC $FAE6                       A:7E X:2 Y:C4 P:64 SP:F9 CYC:20337
FAA6  F0 3E     BEQ $FAE6                       A:7E X:2 Y:C4 P:64 SP:F9 CYC:20339
FAA8  30 3C     BMI $FAE6                       A:7E X:2 Y:C4 P:64 SP:F9 CYC:20341
FAAA  B0 3A     BCS $FAE6                       A:7E X:2 Y:C4 P:64 SP:F9 CYC:20343
FAAC  C9 7E     CMP #$7E                        A:7E X:2 Y:C4 P:64 SP:F9 CYC:20345
FAAE  D0 36     BNE $FAE6                       A:7E X:2 Y:C4 P:67 SP:F9 CYC:20347
FAB0  60        RTS                             A:7E X:2 Y:C4 P:67 SP:F9 CYC:20349
EE99  A5 47     LDA $47 = 6E                    A:7E X:2 Y:C4 P:67 SP:FB CYC:20355
EE9B  C9 6E     CMP #$6E                        A:6E X:2 Y:C4 P:65 SP:FB CYC:20358
EE9D  F0 02     BEQ $EEA1                       A:6E X:2 Y:C4 P:67 SP:FB CYC:20360
EEA1  C8        INY                             A:6E X:2 Y:C4 P:67 SP:FB CYC:20363
EEA2  A9 A5     LDA #$A5                        A:6E X:2 Y:C5 P:E5 SP:FB CYC:20365
EEA4  8D 47 06  STA $0647 = 6E                  A:A5 X:2 Y:C5 P:E5 SP:FB CYC:20367
EEA7  20 7B FA  JSR $FA7B                       A:A5 X:2 Y:C5 P:E5 SP:FB CYC:20371
FA7B  24 01     BIT $01 = FF                    A:A5 X:2 Y:C5 P:E5 SP:F9 CYC:20377
FA7D  18        CLC                             A:A5 X:2 Y:C5 P:E5 SP:F9 CYC:20380
FA7E  A9 B3     LDA #$B3                        A:A5 X:2 Y:C5 P:E4 SP:F9 CYC:20382
FA80  60        RTS                             A:B3 X:2 Y:C5 P:E4 SP:F9 CYC:20384
EEAA  0F 47 06 *SLO $0647 = A5                  A:B3 X:2 Y:C5 P:E4 SP:FB CYC:20390
EEAD  EA        NOP                             A:FB X:2 Y:C5 P:E5 SP:FB CYC:20396
EEAE  EA        NOP                             A:FB X:2 Y:C5 P:E5 SP:FB CYC:20398
EEAF  EA        NOP                             A:FB X:2 Y:C5 P:E5 SP:FB CYC:20400
EEB0  EA        NOP                             A:FB X:2 Y:C5 P:E5 SP:FB CYC:20402
EEB1  20 81 FA  JSR $FA81                       A:FB X:2 Y:C5 P:E5 SP:FB CYC:20404
FA81  50 63     BVC $FAE6                       A:FB X:2 Y:C5 P:E5 SP:F9 CYC:20410
FA83  90 61     BCC $FAE6                       A:FB X:2 Y:C5 P:E5 SP:F9 CYC:20412
FA85  10 5F     BPL $FAE6                       A:FB X:2 Y:C5 P:E5 SP:F9 CYC:20414
FA87  C9 FB     CMP #$FB                        A:FB X:2 Y:C5 P:E5 SP:F9 CYC:20416
FA89  D0 5B     BNE $FAE6                       A:FB X:2 Y:C5 P:67 SP:F9 CYC:20418
FA8B  60        RTS                             A:FB X:2 Y:C5 P:67 SP:F9 CYC:20420
EEB4  AD 47 06  LDA $0647 = 4A                  A:FB X:2 Y:C5 P:67 SP:FB CYC:20426
EEB7  C9 4A     CMP #$4A                        A:4A X:2 Y:C5 P:65 SP:FB CYC:20430
EEB9  F0 02     BEQ $EEBD                       A:4A X:2 Y:C5 P:67 SP:FB CYC:20432
EEBD  C8        INY                             A:4A X:2 Y:C5 P:67 SP:FB CYC:20435
EEBE  A9 29     LDA #$29                        A:4A X:2 Y:C6 P:E5 SP:FB CYC:20437
EEC0  8D 47 06  STA $0647 = 4A                  A:29 X:2 Y:C6 P:65 SP:FB CYC:20439
EEC3  20 8C FA  JSR $FA8C                       A:29 X:2 Y:C6 P:65 SP:FB CYC:20443
FA8C  B8        CLV                             A:29 X:2 Y:C6 P:65 SP:F9 CYC:20449
FA8D  18        CLC                             A:29 X:2 Y:C6 P:25 SP:F9 CYC:20451
FA8E  A9 C3     LDA #$C3                        A:29 X:2 Y:C6 P:24 SP:F9 CYC:20453
FA90  60        RTS                             A:C3 X:2 Y:C6 P:A4 SP:F9 CYC:20455
EEC6  0F 47 06 *SLO $0647 = 29                  A:C3 X:2 Y:C6 P:A4 SP:FB CYC:20461
EEC9  EA        NOP                             A:D3 X:2 Y:C6 P:A4 SP:FB CYC:20467
EECA  EA        NOP                             A:D3 X:2 Y:C6 P:A4 SP:FB CYC:20469
EECB  EA        NOP                             A:D3 X:2 Y:C6 P:A4 SP:FB CYC:20471
EECC  EA        NOP                             A:D3 X:2 Y:C6 P:A4 SP:FB CYC:20473
EECD  20 91 FA  JSR $FA91                       A:D3 X:2 Y:C6 P:A4 SP:FB CYC:20475
FA91  70 53     BVS $FAE6                       A:D3 X:2 Y:C6 P:A4 SP:F9 CYC:20481
FA93  F0 51     BEQ $FAE6                       A:D3 X:2 Y:C6 P:A4 SP:F9 CYC:20483
FA95  10 4F     BPL $FAE6                       A:D3 X:2 Y:C6 P:A4 SP:F9 CYC:20485
FA97  B0 4D     BCS $FAE6                       A:D3 X:2 Y:C6 P:A4 SP:F9 CYC:20487
FA99  C9 D3     CMP #$D3                        A:D3 X:2 Y:C6 P:A4 SP:F9 CYC:20489
FA9B  D0 49     BNE $FAE6                       A:D3 X:2 Y:C6 P:27 SP:F9 CYC:20491
FA9D  60        RTS                             A:D3 X:2 Y:C6 P:27 SP:F9 CYC:20493
EED0  AD 47 06  LDA $0647 = 52                  A:D3 X:2 Y:C6 P:27 SP:FB CYC:20499
EED3  C9 52     CMP #$52                        A:52 X:2 Y:C6 P:25 SP:FB CYC:20503
EED5  F0 02     BEQ $EED9                       A:52 X:2 Y:C6 P:27 SP:FB CYC:20505
EED9  C8        INY                             A:52 X:2 Y:C6 P:27 SP:FB CYC:20508
EEDA  A9 37     LDA #$37                        A:52 X:2 Y:C7 P:A5 SP:FB CYC:20510
EEDC  8D 47 06  STA $0647 = 52                  A:37 X:2 Y:C7 P:25 SP:FB CYC:20512
EEDF  20 9E FA  JSR $FA9E                       A:37 X:2 Y:C7 P:25 SP:FB CYC:20516
FA9E  24 01     BIT $01 = FF                    A:37 X:2 Y:C7 P:25 SP:F9 CYC:20522
FAA0  38        SEC                             A:37 X:2 Y:C7 P:E5 SP:F9 CYC:20525
FAA1  A9 10     LDA #$10                        A:37 X:2 Y:C7 P:E5 SP:F9 CYC:20527
FAA3  60        RTS                             A:10 X:2 Y:C7 P:65 SP:F9 CYC:20529
EEE2  0F 47 06 *SLO $0647 = 37                  A:10 X:2 Y:C7 P:65 SP:FB CYC:20535
EEE5  EA        NOP                             A:7E X:2 Y:C7 P:64 SP:FB CYC:20541
EEE6  EA        NOP                             A:7E X:2 Y:C7 P:64 SP:FB CYC:20543
EEE7  EA        NOP                             A:7E X:2 Y:C7 P:64 SP:FB CYC:20545
EEE8  EA        NOP                             A:7E X:2 Y:C7 P:64 SP:FB CYC:20547
EEE9  20 A4 FA  JSR $FAA4                       A:7E X:2 Y:C7 P:64 SP:FB CYC:20549
FAA4  50 40     BVC $FAE6                       A:7E X:2 Y:C7 P:64 SP:F9 CYC:20555
FAA6  F0 3E     BEQ $FAE6                       A:7E X:2 Y:C7 P:64 SP:F9 CYC:20557
FAA8  30 3C     BMI $FAE6                       A:7E X:2 Y:C7 P:64 SP:F9 CYC:20559
FAAA  B0 3A     BCS $FAE6                       A:7E X:2 Y:C7 P:64 SP:F9 CYC:20561
FAAC  C9 7E     CMP #$7E                        A:7E X:2 Y:C7 P:64 SP:F9 CYC:20563
FAAE  D0 36     BNE $FAE6                       A:7E X:2 Y:C7 P:67 SP:F9 CYC:20565
FAB0  60        RTS                             A:7E X:2 Y:C7 P:67 SP:F9 CYC:20567
EEEC  AD 47 06  LDA $0647 = 6E                  A:7E X:2 Y:C7 P:67 SP:FB CYC:20573
EEEF  C9 6E     CMP #$6E                        A:6E X:2 Y:C7 P:65 SP:FB CYC:20577
EEF1  F0 02     BEQ $EEF5                       A:6E X:2 Y:C7 P:67 SP:FB CYC:20579
EEF5  A9 A5     LDA #$A5                        A:6E X:2 Y:C7 P:67 SP:FB CYC:20582
EEF7  8D 47 06  STA $0647 = 6E                  A:A5 X:2 Y:C7 P:E5 SP:FB CYC:20584
EEFA  A9 48     LDA #$48                        A:A5 X:2 Y:C7 P:E5 SP:FB CYC:20588
EEFC  85 45     STA $45 = 48                    A:48 X:2 Y:C7 P:65 SP:FB CYC:20590
EEFE  A9 05     LDA #$05                        A:48 X:2 Y:C7 P:65 SP:FB CYC:20593
EF00  85 46     STA $46 = 05                    A:5 X:2 Y:C7 P:65 SP:FB CYC:20595
EF02  A0 FF     LDY #$FF                        A:5 X:2 Y:C7 P:65 SP:FB CYC:20598
EF04  20 7B FA  JSR $FA7B                       A:5 X:2 Y:FF P:E5 SP:FB CYC:20600
FA7B  24 01     BIT $01 = FF                    A:5 X:2 Y:FF P:E5 SP:F9 CYC:20606
FA7D  18        CLC                             A:5 X:2 Y:FF P:E5 SP:F9 CYC:20609
FA7E  A9 B3     LDA #$B3                        A:5 X:2 Y:FF P:E4 SP:F9 CYC:20611
FA80  60        RTS                             A:B3 X:2 Y:FF P:E4 SP:F9 CYC:20613
EF07  13 45    *SLO ($45),Y = 0548 @ 0647 = A5  A:B3 X:2 Y:FF P:E4 SP:FB CYC:20619
EF09  EA        NOP                             A:FB X:2 Y:FF P:E5 SP:FB CYC:20627
EF0A  EA        NOP                             A:FB X:2 Y:FF P:E5 SP:FB CYC:20629
EF0B  08        PHP                             A:FB X:2 Y:FF P:E5 SP:FB CYC:20631
EF0C  48        PHA                             A:FB X:2 Y:FF P:E5 SP:FA CYC:20634
EF0D  A0 C8     LDY #$C8                        A:FB X:2 Y:FF P:E5 SP:F9 CYC:20637
EF0F  68        PLA                             A:FB X:2 Y:C8 P:E5 SP:F9 CYC:20639
EF10  28        PLP                             A:FB X:2 Y:C8 P:E5 SP:FA CYC:20643
EF11  20 81 FA  JSR $FA81                       A:FB X:2 Y:C8 P:E5 SP:FB CYC:20647
FA81  50 63     BVC $FAE6                       A:FB X:2 Y:C8 P:E5 SP:F9 CYC:20653
FA83  90 61     BCC $FAE6                       A:FB X:2 Y:C8 P:E5 SP:F9 CYC:20655
FA85  10 5F     BPL $FAE6                       A:FB X:2 Y:C8 P:E5 SP:F9 CYC:20657
FA87  C9 FB     CMP #$FB                        A:FB X:2 Y:C8 P:E5 SP:F9 CYC:20659
FA89  D0 5B     BNE $FAE6                       A:FB X:2 Y:C8 P:67 SP:F9 CYC:20661
FA8B  60        RTS                             A:FB X:2 Y:C8 P:67 SP:F9 CYC:20663
EF14  AD 47 06  LDA $0647 = 4A                  A:FB X:2 Y:C8 P:67 SP:FB CYC:20669
EF17  C9 4A     CMP #$4A                        A:4A X:2 Y:C8 P:65 SP:FB CYC:20673
EF19  F0 02     BEQ $EF1D                       A:4A X:2 Y:C8 P:67 SP:FB CYC:20675
EF1D  A0 FF     LDY #$FF                        A:4A X:2 Y:C8 P:67 SP:FB CYC:20678
EF1F  A9 29     LDA #$29                        A:4A X:2 Y:FF P:E5 SP:FB CYC:20680
EF21  8D 47 06  STA $0647 = 4A                  A:29 X:2 Y:FF P:65 SP:FB CYC:20682
EF24  20 8C FA  JSR $FA8C                       A:29 X:2 Y:FF P:65 SP:FB CYC:20686
FA8C  B8        CLV                             A:29 X:2 Y:FF P:65 SP:F9 CYC:20692
FA8D  18        CLC                             A:29 X:2 Y:FF P:25 SP:F9 CYC:20694
FA8E  A9 C3     LDA #$C3                        A:29 X:2 Y:FF P:24 SP:F9 CYC:20696
FA90  60        RTS                             A:C3 X:2 Y:FF P:A4 SP:F9 CYC:20698
EF27  13 45    *SLO ($45),Y = 0548 @ 0647 = 29  A:C3 X:2 Y:FF P:A4 SP:FB CYC:20704
EF29  EA        NOP                             A:D3 X:2 Y:FF P:A4 SP:FB CYC:20712
EF2A  EA        NOP                             A:D3 X:2 Y:FF P:A4 SP:FB CYC:20714
EF2B  08        PHP                             A:D3 X:2 Y:FF P:A4 SP:FB CYC:20716
EF2C  48        PHA                             A:D3 X:2 Y:FF P:A4 SP:FA CYC:20719
EF2D  A0 C9     LDY #$C9                        A:D3 X:2 Y:FF P:A4 SP:F9 CYC:20722
EF2F  68        PLA                             A:D3 X:2 Y:C9 P:A4 SP:F9 CYC:20724
EF30  28        PLP                             A:D3 X:2 Y:C9 P:A4 SP:FA CYC:20728
EF31  20 91 FA  JSR $FA91                       A:D3 X:2 Y:C9 P:A4 SP:FB CYC:20732
FA91  70 53     BVS $FAE6                       A:D3 X:2 Y:C9 P:A4 SP:F9 CYC:20738
FA93  F0 51     BEQ $FAE6                       A:D3 X:2 Y:C9 P:A4 SP:F9 CYC:20740
FA95  10 4F     BPL $FAE6                       A:D3 X:2 Y:C9 P:A4 SP:F9 CYC:20742
FA97  B0 4D     BCS $FAE6                       A:D3 X:2 Y:C9 P:A4 SP:F9 CYC:20744
FA99  C9 D3     CMP #$D3                        A:D3 X:2 Y:C9 P:A4 SP:F9 CYC:20746
FA9B  D0 49     BNE $FAE6                       A:D3 X:2 Y:C9 P:27 SP:F9 CYC:20748
FA9D  60        RTS                             A:D3 X:2 Y:C9 P:27 SP:F9 CYC:20750
EF34  AD 47 06  LDA $0647 = 52                  A:D3 X:2 Y:C9 P:27 SP:FB CYC:20756
EF37  C9 52     CMP #$52                        A:52 X:2 Y:C9 P:25 SP:FB CYC:20760
EF39  F0 02     BEQ $EF3D                       A:52 X:2 Y:C9 P:27 SP:FB CYC:20762
EF3D  A0 FF     LDY #$FF                        A:52 X:2 Y:C9 P:27 SP:FB CYC:20765
EF3F  A9 37     LDA #$37                        A:52 X:2 Y:FF P:A5 SP:FB CYC:20767
EF41  8D 47 06  STA $0647 = 52                  A:37 X:2 Y:FF P:25 SP:FB CYC:20769
EF44  20 9E FA  JSR $FA9E                       A:37 X:2 Y:FF P:25 SP:FB CYC:20773
FA9E  24 01     BIT $01 = FF                    A:37 X:2 Y:FF P:25 SP:F9 CYC:20779
FAA0  38        SEC                             A:37 X:2 Y:FF P:E5 SP:F9 CYC:20782
FAA1  A9 10     LDA #$10                        A:37 X:2 Y:FF P:E5 SP:F9 CYC:20784
FAA3  60        RTS                             A:10 X:2 Y:FF P:65 SP:F9 CYC:20786
EF47  13 45    *SLO ($45),Y = 0548 @ 0647 = 37  A:10 X:2 Y:FF P:65 SP:FB CYC:20792
EF49  EA        NOP                             A:7E X:2 Y:FF P:64 SP:FB CYC:20800
EF4A  EA        NOP                             A:7E X:2 Y:FF P:64 SP:FB CYC:20802
EF4B  08        PHP                             A:7E X:2 Y:FF P:64 SP:FB CYC:20804
EF4C  48        PHA                             A:7E X:2 Y:FF P:64 SP:FA CYC:20807
EF4D  A0 CA     LDY #$CA                        A:7E X:2 Y:FF P:64 SP:F9 CYC:20810
EF4F  68        PLA                             A:7E X:2 Y:CA P:E4 SP:F9 CYC:20812
EF50  28        PLP                             A:7E X:2 Y:CA P:64 SP:FA CYC:20816
EF51  20 A4 FA  JSR $FAA4                       A:7E X:2 Y:CA P:64 SP:FB CYC:20820
FAA4  50 40     BVC $FAE6                       A:7E X:2 Y:CA P:64 SP:F9 CYC:20826
FAA6  F0 3E     BEQ $FAE6                       A:7E X:2 Y:CA P:64 SP:F9 CYC:20828
FAA8  30 3C     BMI $FAE6                       A:7E X:2 Y:CA P:64 SP:F9 CYC:20830
FAAA  B0 3A     BCS $FAE6                       A:7E X:2 Y:CA P:64 SP:F9 CYC:20832
FAAC  C9 7E     CMP #$7E                        A:7E X:2 Y:CA P:64 SP:F9 CYC:20834
FAAE  D0 36     BNE $FAE6                       A:7E X:2 Y:CA P:67 SP:F9 CYC:20836
FAB0  60        RTS                             A:7E X:2 Y:CA P:67 SP:F9 CYC:20838
EF54  AD 47 06  LDA $0647 = 6E                  A:7E X:2 Y:CA P:67 SP:FB CYC:20844
EF57  C9 6E     CMP #$6E                        A:6E X:2 Y:CA P:65 SP:FB CYC:20848
EF59  F0 02     BEQ $EF5D                       A:6E X:2 Y:CA P:67 SP:FB CYC:20850
EF5D  A0 CB     LDY #$CB                        A:6E X:2 Y:CA P:67 SP:FB CYC:20853
EF5F  A2 FF     LDX #$FF                        A:6E X:2 Y:CB P:E5 SP:FB CYC:20855
EF61  A9 A5     LDA #$A5                        A:6E X:FF Y:CB P:E5 SP:FB CYC:20857
EF63  85 47     STA $47 = 6E                    A:A5 X:FF Y:CB P:E5 SP:FB CYC:20859
EF65  20 7B FA  JSR $FA7B                       A:A5 X:FF Y:CB P:E5 SP:FB CYC:20862
FA7B  24 01     BIT $01 = FF                    A:A5 X:FF Y:CB P:E5 SP:F9 CYC:20868
FA7D  18        CLC                             A:A5 X:FF Y:CB P:E5 SP:F9 CYC:20871
FA7E  A9 B3     LDA #$B3                        A:A5 X:FF Y:CB P:E4 SP:F9 CYC:20873
FA80  60        RTS                             A:B3 X:FF Y:CB P:E4 SP:F9 CYC:20875
EF68  17 48    *SLO $48,X @ 47 = A5             A:B3 X:FF Y:CB P:E4 SP:FB CYC:20881
EF6A  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB CYC:20887
EF6B  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB CYC:20889
EF6C  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB CYC:20891
EF6D  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB CYC:20893
EF6E  20 81 FA  JSR $FA81                       A:FB X:FF Y:CB P:E5 SP:FB CYC:20895
FA81  50 63     BVC $FAE6                       A:FB X:FF Y:CB P:E5 SP:F9 CYC:20901
FA83  90 61     BCC $FAE6                       A:FB X:FF Y:CB P:E5 SP:F9 CYC:20903
FA85  10 5F     BPL $FAE6                       A:FB X:FF Y:CB P:E5 SP:F9 CYC:20905
FA87  C9 FB     CMP #$FB                        A:FB X:FF Y:CB P:E5 SP:F9 CYC:20907
FA89  D0 5B     BNE $FAE6                       A:FB X:FF Y:CB P:67 SP:F9 CYC:20909
FA8B  60        RTS                             A:FB X:FF Y:CB P:67 SP:F9 CYC:20911
EF71  A5 47     LDA $47 = 4A                    A:FB X:FF Y:CB P:67 SP:FB CYC:20917
EF73  C9 4A     CMP #$4A                        A:4A X:FF Y:CB P:65 SP:FB CYC:20920
EF75  F0 02     BEQ $EF79                       A:4A X:FF Y:CB P:67 SP:FB CYC:20922
EF79  C8        INY                             A:4A X:FF Y:CB P:67 SP:FB CYC:20925
EF7A  A9 29     LDA #$29                        A:4A X:FF Y:CC P:E5 SP:FB CYC:20927
EF7C  85 47     STA $47 = 4A                    A:29 X:FF Y:CC P:65 SP:FB CYC:20929
EF7E  20 8C FA  JSR $FA8C                       A:29 X:FF Y:CC P:65 SP:FB CYC:20932
FA8C  B8        CLV                             A:29 X:FF Y:CC P:65 SP:F9 CYC:20938
FA8D  18        CLC                             A:29 X:FF Y:CC P:25 SP:F9 CYC:20940
FA8E  A9 C3     LDA #$C3                        A:29 X:FF Y:CC P:24 SP:F9 CYC:20942
FA90  60        RTS                             A:C3 X:FF Y:CC P:A4 SP:F9 CYC:20944
EF81  17 48    *SLO $48,X @ 47 = 29             A:C3 X:FF Y:CC P:A4 SP:FB CYC:20950
EF83  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB CYC:20956
EF84  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB CYC:20958
EF85  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB CYC:20960
EF86  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB CYC:20962
EF87  20 91 FA  JSR $FA91                       A:D3 X:FF Y:CC P:A4 SP:FB CYC:20964
FA91  70 53     BVS $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9 CYC:20970
FA93  F0 51     BEQ $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9 CYC:20972
FA95  10 4F     BPL $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9 CYC:20974
FA97  B0 4D     BCS $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9 CYC:20976
FA99  C9 D3     CMP #$D3                        A:D3 X:FF Y:CC P:A4 SP:F9 CYC:20978
FA9B  D0 49     BNE $FAE6                       A:D3 X:FF Y:CC P:27 SP:F9 CYC:20980
FA9D  60        RTS                             A:D3 X:FF Y:CC P:27 SP:F9 CYC:20982
EF8A  A5 47     LDA $47 = 52                    A:D3 X:FF Y:CC P:27 SP:FB CYC:20988
EF8C  C9 52     CMP #$52                        A:52 X:FF Y:CC P:25 SP:FB CYC:20991
EF8E  F0 02     BEQ $EF92                       A:52 X:FF Y:CC P:27 SP:FB CYC:20993
EF92  C8        INY                             A:52 X:FF Y:CC P:27 SP:FB CYC:20996
EF93  A9 37     LDA #$37                        A:52 X:FF Y:CD P:A5 SP:FB CYC:20998
EF95  85 47     STA $47 = 52                    A:37 X:FF Y:CD P:25 SP:FB CYC:21000
EF97  20 9E FA  JSR $FA9E                       A:37 X:FF Y:CD P:25 SP:FB CYC:21003
FA9E  24 01     BIT $01 = FF                    A:37 X:FF Y:CD P:25 SP:F9 CYC:21009
FAA0  38        SEC                             A:37 X:FF Y:CD P:E5 SP:F9 CYC:21012
FAA1  A9 10     LDA #$10                        A:37 X:FF Y:CD P:E5 SP:F9 CYC:21014
FAA3  60        RTS                             A:10 X:FF Y:CD P:65 SP:F9 CYC:21016
EF9A  17 48    *SLO $48,X @ 47 = 37             A:10 X:FF Y:CD P:65 SP:FB CYC:21022
EF9C  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB CYC:21028
EF9D  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB CYC:21030
EF9E  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB CYC:21032
EF9F  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB CYC:21034
EFA0  20 A4 FA  JSR $FAA4                       A:7E X:FF Y:CD P:64 SP:FB CYC:21036
FAA4  50 40     BVC $FAE6                       A:7E X:FF Y:CD P:64 SP:F9 CYC:21042
FAA6  F0 3E     BEQ $FAE6                       A:7E X:FF Y:CD P:64 SP:F9 CYC:21044
FAA8  30 3C     BMI $FAE6                       A:7E X:FF Y:CD P:64 SP:F9 CYC:21046
FAAA  B0 3A     BCS $FAE6                       A:7E X:FF Y:CD P:64 SP:F9 CYC:21048
FAAC  C9 7E     CMP #$7E                        A:7E X:FF Y:CD P:64 SP:F9 CYC:21050
FAAE  D0 36     BNE $FAE6                       A:7E X:FF Y:CD P:67 SP:F9 CYC:21052
FAB0  60        RTS                             A:7E X:FF Y:CD P:67 SP:F9 CYC:21054
EFA3  A5 47     LDA $47 = 6E                    A:7E X:FF Y:CD P:67 SP:FB CYC:21060
EFA5  C9 6E     CMP #$6E                        A:6E X:FF Y:CD P:65 SP:FB CYC:21063
EFA7  F0 02     BEQ $EFAB                       A:6E X:FF Y:CD P:67 SP:FB CYC:21065
EFAB  A9 A5     LDA #$A5                        A:6E X:FF Y:CD P:67 SP:FB CYC:21068
EFAD  8D 47 06  STA $0647 = 6E                  A:A5 X:FF Y:CD P:E5 SP:FB CYC:21070
EFB0  A0 FF     LDY #$FF                        A:A5 X:FF Y:CD P:E5 SP:FB CYC:21074
EFB2  20 7B FA  JSR $FA7B                       A:A5 X:FF Y:FF P:E5 SP:FB CYC:21076
FA7B  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:E5 SP:F9 CYC:21082
FA7D  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9 CYC:21085
FA7E  A9 B3     LDA #$B3                        A:A5 X:FF Y:FF P:E4 SP:F9 CYC:21087
FA80  60        RTS                             A:B3 X:FF Y:FF P:E4 SP:F9 CYC:21089
EFB5  1B 48 05 *SLO $0548,Y @ 0647 = A5         A:B3 X:FF Y:FF P:E4 SP:FB CYC:21095
EFB8  EA        NOP                             A:FB X:FF Y:FF P:E5 SP:FB CYC:21102
EFB9  EA        NOP                             A:FB X:FF Y:FF P:E5 SP:FB CYC:21104
EFBA  08        PHP                             A:FB X:FF Y:FF P:E5 SP:FB CYC:21106
EFBB  48        PHA                             A:FB X:FF Y:FF P:E5 SP:FA CYC:21109
EFBC  A0 CE     LDY #$CE                        A:FB X:FF Y:FF P:E5 SP:F9 CYC:21112
EFBE  68        PLA                             A:FB X:FF Y:CE P:E5 SP:F9 CYC:21114
EFBF  28        PLP                             A:FB X:FF Y:CE P:E5 SP:FA CYC:21118
EFC0  20 81 FA  JSR $FA81                       A:FB X:FF Y:CE P:E5 SP:FB CYC:21122
FA81  50 63     BVC $FAE6                       A:FB X:FF Y:CE P:E5 SP:F9 CYC:21128
FA83  90 61     BCC $FAE6                       A:FB X:FF Y:CE P:E5 SP:F9 CYC:21130
FA85  10 5F     BPL $FAE6                       A:FB X:FF Y:CE P:E5 SP:F9 CYC:21132
FA87  C9 FB     CMP #$FB                        A:FB X:FF Y:CE P:E5 SP:F9 CYC:21134
FA89  D0 5B     BNE $FAE6                       A:FB X:FF Y:CE P:67 SP:F9 CYC:21136
FA8B  60        RTS                             A:FB X:FF Y:CE P:67 SP:F9 CYC:21138
EFC3  AD 47 06  LDA $0647 = 4A                  A:FB X:FF Y:CE P:67 SP:FB CYC:21144
EFC6  C9 4A     CMP #$4A                        A:4A X:FF Y:CE P:65 SP:FB CYC:21148
EFC8  F0 02     BEQ $EFCC                       A:4A X:FF Y:CE P:67 SP:FB CYC:21150
EFCC  A0 FF     LDY #$FF                        A:4A X:FF Y:CE P:67 SP:FB CYC:21153
EFCE  A9 29     LDA #$29                        A:4A X:FF Y:FF P:E5 SP:FB CYC:21155
EFD0  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:FF P:65 SP:FB CYC:21157
EFD3  20 8C FA  JSR $FA8C                       A:29 X:FF Y:FF P:65 SP:FB CYC:21161
FA8C  B8        CLV                             A:29 X:FF Y:FF P:65 SP:F9 CYC:21167
FA8D  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9 CYC:21169
FA8E  A9 C3     LDA #$C3                        A:29 X:FF Y:FF P:24 SP:F9 CYC:21171
FA90  60        RTS                             A:C3 X:FF Y:FF P:A4 SP:F9 CYC:21173
EFD6  1B 48 05 *SLO $0548,Y @ 0647 = 29         A:C3 X:FF Y:FF P:A4 SP:FB CYC:21179
EFD9  EA        NOP                             A:D3 X:FF Y:FF P:A4 SP:FB CYC:21186
EFDA  EA        NOP                             A:D3 X:FF Y:FF P:A4 SP:FB CYC:21188
EFDB  08        PHP                             A:D3 X:FF Y:FF P:A4 SP:FB CYC:21190
EFDC  48        PHA                             A:D3 X:FF Y:FF P:A4 SP:FA CYC:21193
EFDD  A0 CF     LDY #$CF                        A:D3 X:FF Y:FF P:A4 SP:F9 CYC:21196
EFDF  68        PLA                             A:D3 X:FF Y:CF P:A4 SP:F9 CYC:21198
EFE0  28        PLP                             A:D3 X:FF Y:CF P:A4 SP:FA CYC:21202
EFE1  20 91 FA  JSR $FA91                       A:D3 X:FF Y:CF P:A4 SP:FB CYC:21206
FA91  70 53     BVS $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9 CYC:21212
FA93  F0 51     BEQ $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9 CYC:21214
FA95  10 4F     BPL $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9 CYC:21216
FA97  B0 4D     BCS $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9 CYC:21218
FA99  C9 D3     CMP #$D3                        A:D3 X:FF Y:CF P:A4 SP:F9 CYC:21220
FA9B  D0 49     BNE $FAE6                       A:D3 X:FF Y:CF P:27 SP:F9 CYC:21222
FA9D  60        RTS                             A:D3 X:FF Y:CF P:27 SP:F9 CYC:21224
EFE4  AD 47 06  LDA $0647 = 52                  A:D3 X:FF Y:CF P:27 SP:FB CYC:21230
EFE7  C9 52     CMP #$52                        A:52 X:FF Y:CF P:25 SP:FB CYC:21234
EFE9  F0 02     BEQ $EFED                       A:52 X:FF Y:CF P:27 SP:FB CYC:21236
EFED  A0 FF     LDY #$FF                        A:52 X:FF Y:CF P:27 SP:FB CYC:21239
EFEF  A9 37     LDA #$37                        A:52 X:FF Y:FF P:A5 SP:FB CYC:21241
EFF1  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:FF P:25 SP:FB CYC:21243
EFF4  20 9E FA  JSR $FA9E                       A:37 X:FF Y:FF P:25 SP:FB CYC:21247
FA9E  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9 CYC:21253
FAA0  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9 CYC:21256
FAA1  A9 10     LDA #$10                        A:37 X:FF Y:FF P:E5 SP:F9 CYC:21258
FAA3  60        RTS                             A:10 X:FF Y:FF P:65 SP:F9 CYC:21260
EFF7  1B 48 05 *SLO $0548,Y @ 0647 = 37         A:10 X:FF Y:FF P:65 SP:FB CYC:21266
EFFA  EA        NOP                             A:7E X:FF Y:FF P:64 SP:FB CYC:21273
EFFB  EA        NOP                             A:7E X:FF Y:FF P:64 SP:FB CYC:21275
EFFC  08        PHP                             A:7E X:FF Y:FF P:64 SP:FB CYC:21277
EFFD  48        PHA                             A:7E X:FF Y:FF P:64 SP:FA CYC:21280
EFFE  A0 D0     LDY #$D0                        A:7E X:FF Y:FF P:64 SP:F9 CYC:21283
F000  68        PLA                             A:7E X:FF Y:D0 P:E4 SP:F9 CYC:21285
F001  28        PLP                             A:7E X:FF Y:D0 P:64 SP:FA CYC:21289
F002  20 A4 FA  JSR $FAA4                       A:7E X:FF Y:D0 P:64 SP:FB CYC:21293
FAA4  50 40     BVC $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9 CYC:21299
FAA6  F0 3E     BEQ $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9 CYC:21301
FAA8  30 3C     BMI $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9 CYC:21303
FAAA  B0 3A     BCS $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9 CYC:21305
FAAC  C9 7E     CMP #$7E                        A:7E X:FF Y:D0 P:64 SP:F9 CYC:21307
FAAE  D0 36     BNE $FAE6                       A:7E X:FF Y:D0 P:67 SP:F9 CYC:21309
FAB0  60        RTS                             A:7E X:FF Y:D0 P:67 SP:F9 CYC:21311
F005  AD 47 06  LDA $0647 = 6E                  A:7E X:FF Y:D0 P:67 SP:FB CYC:21317
F008  C9 6E     CMP #$6E                        A:6E X:FF Y:D0 P:65 SP:FB CYC:21321
F00A  F0 02     BEQ $F00E                       A:6E X:FF Y:D0 P:67 SP:FB CYC:21323
F00E  A0 D1     LDY #$D1                        A:6E X:FF Y:D0 P:67 SP:FB CYC:21326
F010  A2 FF     LDX #$FF                        A:6E X:FF Y:D1 P:E5 SP:FB CYC:21328
F012  A9 A5     LDA #$A5                        A:6E X:FF Y:D1 P:E5 SP:FB CYC:21330
F014  8D 47 06  STA $0647 = 6E                  A:A5 X:FF Y:D1 P:E5 SP:FB CYC:21332
F017  20 7B FA  JSR $FA7B                       A:A5 X:FF Y:D1 P:E5 SP:FB CYC:21336
FA7B  24 01     BIT $01 = FF                    A:A5 X:FF Y:D1 P:E5 SP:F9 CYC:21342
FA7D  18        CLC                             A:A5 X:FF Y:D1 P:E5 SP:F9 CYC:21345
FA7E  A9 B3     LDA #$B3                        A:A5 X:FF Y:D1 P:E4 SP:F9 CYC:21347
FA80  60        RTS                             A:B3 X:FF Y:D1 P:E4 SP:F9 CYC:21349
F01A  1F 48 05 *SLO $0548,X @ 0647 = A5         A:B3 X:FF Y:D1 P:E4 SP:FB CYC:21355
F01D  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB CYC:21362
F01E  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB CYC:21364
F01F  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB CYC:21366
F020  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB CYC:21368
F021  20 81 FA  JSR $FA81                       A:FB X:FF Y:D1 P:E5 SP:FB CYC:21370
FA81  50 63     BVC $FAE6                       A:FB X:FF Y:D1 P:E5 SP:F9 CYC:21376
FA83  90 61     BCC $FAE6                       A:FB X:FF Y:D1 P:E5 SP:F9 CYC:21378
FA85  10 5F     BPL $FAE6                       A:FB X:FF Y:D1 P:E5 SP:F9 CYC:21380
FA87  C9 FB     CMP #$FB                        A:FB X:FF Y:D1 P:E5 SP:F9 CYC:21382
FA89  D0 5B     BNE $FAE6                       A:FB X:FF Y:D1 P:67 SP:F9 CYC:21384
FA8B  60        RTS                             A:FB X:FF Y:D1 P:67 SP:F9 CYC:21386
F024  AD 47 06  LDA $0647 = 4A                  A:FB X:FF Y:D1 P:67 SP:FB CYC:21392
F027  C9 4A     CMP #$4A                        A:4A X:FF Y:D1 P:65 SP:FB CYC:21396
F029  F0 02     BEQ $F02D                       A:4A X:FF Y:D1 P:67 SP:FB CYC:21398
F02D  C8        INY                             A:4A X:FF Y:D1 P:67 SP:FB CYC:21401
F02E  A9 29     LDA #$29                        A:4A X:FF Y:D2 P:E5 SP:FB CYC:21403
F030  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:D2 P:65 SP:FB CYC:21405
F033  20 8C FA  JSR $FA8C                       A:29 X:FF Y:D2 P:65 SP:FB CYC:21409
FA8C  B8        CLV                             A:29 X:FF Y:D2 P:65 SP:F9 CYC:21415
FA8D  18        CLC                             A:29 X:FF Y:D2 P:25 SP:F9 CYC:21417
FA8E  A9 C3     LDA #$C3                        A:29 X:FF Y:D2 P:24 SP:F9 CYC:21419
FA90  60        RTS                             A:C3 X:FF Y:D2 P:A4 SP:F9 CYC:21421
F036  1F 48 05 *SLO $0548,X @ 0647 = 29         A:C3 X:FF Y:D2 P:A4 SP:FB CYC:21427
F039  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB CYC:21434
F03A  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB CYC:21436
F03B  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB CYC:21438
F03C  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB CYC:21440
F03D  20 91 FA  JSR $FA91                       A:D3 X:FF Y:D2 P:A4 SP:FB CYC:21442
FA91  70 53     BVS $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9 CYC:21448
FA93  F0 51     BEQ $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9 CYC:21450
FA95  10 4F     BPL $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9 CYC:21452
FA97  B0 4D     BCS $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9 CYC:21454
FA99  C9 D3     CMP #$D3                        A:D3 X:FF Y:D2 P:A4 SP:F9 CYC:21456
FA9B  D0 49     BNE $FAE6                       A:D3 X:FF Y:D2 P:27 SP:F9 CYC:21458
FA9D  60        RTS                             A:D3 X:FF Y:D2 P:27 SP:F9 CYC:21460
F040  AD 47 06  LDA $0647 = 52                  A:D3 X:FF Y:D2 P:27 SP:FB CYC:21466
F043  C9 52     CMP #$52                        A:52 X:FF Y:D2 P:25 SP:FB CYC:21470
F045  F0 02     BEQ $F049                       A:52 X:FF Y:D2 P:27 SP:FB CYC:21472
F049  C8        INY                             A:52 X:FF Y:D2 P:27 SP:FB CYC:21475
F04A  A9 37     LDA #$37                        A:52 X:FF Y:D3 P:A5 SP:FB CYC:21477
F04C  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:D3 P:25 SP:FB CYC:21479
F04F  20 9E FA  JSR $FA9E                       A:37 X:FF Y:D3 P:25 SP:FB CYC:21483
FA9E  24 01     BIT $01 = FF                    A:37 X:FF Y:D3 P:25 SP:F9 CYC:21489
FAA0  38        SEC                             A:37 X:FF Y:D3 P:E5 SP:F9 CYC:21492
FAA1  A9 10     LDA #$10                        A:37 X:FF Y:D3 P:E5 SP:F9 CYC:21494
FAA3  60        RTS                             A:10 X:FF Y:D3 P:65 SP:F9 CYC:21496
F052  1F 48 05 *SLO $0548,X @ 0647 = 37         A:10 X:FF Y:D3 P:65 SP:FB CYC:21502
F055  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB CYC:21509
F056  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB CYC:21511
F057  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB CYC:21513
F058  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB CYC:21515
F059  20 A4 FA  JSR $FAA4                       A:7E X:FF Y:D3 P:64 SP:FB CYC:21517
FAA4  50 40     BVC $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9 CYC:21523
FAA6  F0 3E     BEQ $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9 CYC:21525
FAA8  30 3C     BMI $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9 CYC:21527
FAAA  B0 3A     BCS $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9 CYC:21529
FAAC  C9 7E     CMP #$7E                        A:7E X:FF Y:D3 P:64 SP:F9 CYC:21531
FAAE  D0 36     BNE $FAE6                       A:7E X:FF Y:D3 P:67 SP:F9 CYC:21533
FAB0  60        RTS                             A:7E X:FF Y:D3 P:67 SP:F9 CYC:21535
F05C  AD 47 06  LDA $0647 = 6E                  A:7E X:FF Y:D3 P:67 SP:FB CYC:21541
F05F  C9 6E     CMP #$6E                        A:6E X:FF Y:D3 P:65 SP:FB CYC:21545
F061  F0 02     BEQ $F065                       A:6E X:FF Y:D3 P:67 SP:FB CYC:21547
F065  60        RTS                             A:6E X:FF Y:D3 P:67 SP:FB CYC:21550
C644  20 66 F0  JSR $F066                       A:6E X:FF Y:D3 P:67 SP:FD CYC:21556
F066  A9 FF     LDA #$FF                        A:6E X:FF Y:D3 P:67 SP:FB CYC:21562
F068  85 01     STA $01 = FF                    A:FF X:FF Y:D3 P:E5 SP:FB CYC:21564
F06A  A0 D4     LDY #$D4                        A:FF X:FF Y:D3 P:E5 SP:FB CYC:21567
F06C  A2 02     LDX #$02                        A:FF X:FF Y:D4 P:E5 SP:FB CYC:21569
F06E  A9 47     LDA #$47                        A:FF X:2 Y:D4 P:65 SP:FB CYC:21571
F070  85 47     STA $47 = 6E                    A:47 X:2 Y:D4 P:65 SP:FB CYC:21573
F072  A9 06     LDA #$06                        A:47 X:2 Y:D4 P:65 SP:FB CYC:21576
F074  85 48     STA $48 = 06                    A:6 X:2 Y:D4 P:65 SP:FB CYC:21578
F076  A9 A5     LDA #$A5                        A:6 X:2 Y:D4 P:65 SP:FB CYC:21581
F078  8D 47 06  STA $0647 = 6E                  A:A5 X:2 Y:D4 P:E5 SP:FB CYC:21583
F07B  20 53 FB  JSR $FB53                       A:A5 X:2 Y:D4 P:E5 SP:FB CYC:21587
FB53  24 01     BIT $01 = FF                    A:A5 X:2 Y:D4 P:E5 SP:F9 CYC:21593
FB55  18        CLC                             A:A5 X:2 Y:D4 P:E5 SP:F9 CYC:21596
FB56  A9 B3     LDA #$B3                        A:A5 X:2 Y:D4 P:E4 SP:F9 CYC:21598
FB58  60        RTS                             A:B3 X:2 Y:D4 P:E4 SP:F9 CYC:21600
F07E  23 45    *RLA ($45,X) @ 47 = 0647 = A5    A:B3 X:2 Y:D4 P:E4 SP:FB CYC:21606
F080  EA        NOP                             A:2 X:2 Y:D4 P:65 SP:FB CYC:21614
F081  EA        NOP                             A:2 X:2 Y:D4 P:65 SP:FB CYC:21616
F082  EA        NOP                             A:2 X:2 Y:D4 P:65 SP:FB CYC:21618
F083  EA        NOP                             A:2 X:2 Y:D4 P:65 SP:FB CYC:21620
F084  20 59 FB  JSR $FB59                       A:2 X:2 Y:D4 P:65 SP:FB CYC:21622
FB59  50 1A     BVC $FB75                       A:2 X:2 Y:D4 P:65 SP:F9 CYC:21628
FB5B  90 18     BCC $FB75                       A:2 X:2 Y:D4 P:65 SP:F9 CYC:21630
FB5D  30 16     BMI $FB75                       A:2 X:2 Y:D4 P:65 SP:F9 CYC:21632
FB5F  C9 02     CMP #$02                        A:2 X:2 Y:D4 P:65 SP:F9 CYC:21634
FB61  D0 12     BNE $FB75                       A:2 X:2 Y:D4 P:67 SP:F9 CYC:21636
FB63  60        RTS                             A:2 X:2 Y:D4 P:67 SP:F9 CYC:21638
F087  AD 47 06  LDA $0647 = 4A                  A:2 X:2 Y:D4 P:67 SP:FB CYC:21644
F08A  C9 4A     CMP #$4A                        A:4A X:2 Y:D4 P:65 SP:FB CYC:21648
F08C  F0 02     BEQ $F090                       A:4A X:2 Y:D4 P:67 SP:FB CYC:21650
F090  C8        INY                             A:4A X:2 Y:D4 P:67 SP:FB CYC:21653
F091  A9 29     LDA #$29                        A:4A X:2 Y:D5 P:E5 SP:FB CYC:21655
F093  8D 47 06  STA $0647 = 4A                  A:29 X:2 Y:D5 P:65 SP:FB CYC:21657
F096  20 64 FB  JSR $FB64                       A:29 X:2 Y:D5 P:65 SP:FB CYC:21661
FB64  B8        CLV                             A:29 X:2 Y:D5 P:65 SP:F9 CYC:21667
FB65  18        CLC                             A:29 X:2 Y:D5 P:25 SP:F9 CYC:21669
FB66  A9 42     LDA #$42                        A:29 X:2 Y:D5 P:24 SP:F9 CYC:21671
FB68  60        RTS                             A:42 X:2 Y:D5 P:24 SP:F9 CYC:21673
F099  23 45    *RLA ($45,X) @ 47 = 0647 = 29    A:42 X:2 Y:D5 P:24 SP:FB CYC:21679
F09B  EA        NOP                             A:42 X:2 Y:D5 P:24 SP:FB CYC:21687
F09C  EA        NOP                             A:42 X:2 Y:D5 P:24 SP:FB CYC:21689
F09D  EA        NOP                             A:42 X:2 Y:D5 P:24 SP:FB CYC:21691
F09E  EA        NOP                             A:42 X:2 Y:D5 P:24 SP:FB CYC:21693
F09F  20 69 FB  JSR $FB69                       A:42 X:2 Y:D5 P:24 SP:FB CYC:21695
FB69  70 0A     BVS $FB75                       A:42 X:2 Y:D5 P:24 SP:F9 CYC:21701
FB6B  F0 08     BEQ $FB75                       A:42 X:2 Y:D5 P:24 SP:F9 CYC:21703
FB6D  30 06     BMI $FB75                       A:42 X:2 Y:D5 P:24 SP:F9 CYC:21705
FB6F  B0 04     BCS $FB75                       A:42 X:2 Y:D5 P:24 SP:F9 CYC:21707
FB71  C9 42     CMP #$42                        A:42 X:2 Y:D5 P:24 SP:F9 CYC:21709
FB73  F0 02     BEQ $FB77                       A:42 X:2 Y:D5 P:27 SP:F9 CYC:21711
FB77  60        RTS                             A:42 X:2 Y:D5 P:27 SP:F9 CYC:21714
F0A2  AD 47 06  LDA $0647 = 52                  A:42 X:2 Y:D5 P:27 SP:FB CYC:21720
F0A5  C9 52     CMP #$52                        A:52 X:2 Y:D5 P:25 SP:FB CYC:21724
F0A7  F0 02     BEQ $F0AB                       A:52 X:2 Y:D5 P:27 SP:FB CYC:21726
F0AB  C8        INY                             A:52 X:2 Y:D5 P:27 SP:FB CYC:21729
F0AC  A9 37     LDA #$37                        A:52 X:2 Y:D6 P:A5 SP:FB CYC:21731
F0AE  8D 47 06  STA $0647 = 52                  A:37 X:2 Y:D6 P:25 SP:FB CYC:21733
F0B1  20 68 FA  JSR $FA68                       A:37 X:2 Y:D6 P:25 SP:FB CYC:21737
FA68  24 01     BIT $01 = FF                    A:37 X:2 Y:D6 P:25 SP:F9 CYC:21743
FA6A  38        SEC                             A:37 X:2 Y:D6 P:E5 SP:F9 CYC:21746
FA6B  A9 75     LDA #$75                        A:37 X:2 Y:D6 P:E5 SP:F9 CYC:21748
FA6D  60        RTS                             A:75 X:2 Y:D6 P:65 SP:F9 CYC:21750
F0B4  23 45    *RLA ($45,X) @ 47 = 0647 = 37    A:75 X:2 Y:D6 P:65 SP:FB CYC:21756
F0B6  EA        NOP                             A:65 X:2 Y:D6 P:64 SP:FB CYC:21764
F0B7  EA        NOP                             A:65 X:2 Y:D6 P:64 SP:FB CYC:21766
F0B8  EA        NOP                             A:65 X:2 Y:D6 P:64 SP:FB CYC:21768
F0B9  EA        NOP                             A:65 X:2 Y:D6 P:64 SP:FB CYC:21770
F0BA  20 6E FA  JSR $FA6E                       A:65 X:2 Y:D6 P:64 SP:FB CYC:21772
FA6E  50 76     BVC $FAE6                       A:65 X:2 Y:D6 P:64 SP:F9 CYC:21778
FA70  F0 74     BEQ $FAE6                       A:65 X:2 Y:D6 P:64 SP:F9 CYC:21780
FA72  30 72     BMI $FAE6                       A:65 X:2 Y:D6 P:64 SP:F9 CYC:21782
FA74  B0 70     BCS $FAE6                       A:65 X:2 Y:D6 P:64 SP:F9 CYC:21784
FA76  C9 65     CMP #$65                        A:65 X:2 Y:D6 P:64 SP:F9 CYC:21786
FA78  D0 6C     BNE $FAE6                       A:65 X:2 Y:D6 P:67 SP:F9 CYC:21788
FA7A  60        RTS                             A:65 X:2 Y:D6 P:67 SP:F9 CYC:21790
F0BD  AD 47 06  LDA $0647 = 6F                  A:65 X:2 Y:D6 P:67 SP:FB CYC:21796
F0C0  C9 6F     CMP #$6F                        A:6F X:2 Y:D6 P:65 SP:FB CYC:21800
F0C2  F0 02     BEQ $F0C6                       A:6F X:2 Y:D6 P:67 SP:FB CYC:21802
F0C6  C8        INY                             A:6F X:2 Y:D6 P:67 SP:FB CYC:21805
F0C7  A9 A5     LDA #$A5                        A:6F X:2 Y:D7 P:E5 SP:FB CYC:21807
F0C9  85 47     STA $47 = 47                    A:A5 X:2 Y:D7 P:E5 SP:FB CYC:21809
F0CB  20 53 FB  JSR $FB53                       A:A5 X:2 Y:D7 P:E5 SP:FB CYC:21812
FB53  24 01     BIT $01 = FF                    A:A5 X:2 Y:D7 P:E5 SP:F9 CYC:21818
FB55  18        CLC                             A:A5 X:2 Y:D7 P:E5 SP:F9 CYC:21821
FB56  A9 B3     LDA #$B3                        A:A5 X:2 Y:D7 P:E4 SP:F9 CYC:21823
FB58  60        RTS                             A:B3 X:2 Y:D7 P:E4 SP:F9 CYC:21825
F0CE  27 47    *RLA $47 = A5                    A:B3 X:2 Y:D7 P:E4 SP:FB CYC:21831
F0D0  EA        NOP                             A:2 X:2 Y:D7 P:65 SP:FB CYC:21836
F0D1  EA        NOP                             A:2 X:2 Y:D7 P:65 SP:FB CYC:21838
F0D2  EA        NOP                             A:2 X:2 Y:D7 P:65 SP:FB CYC:21840
F0D3  EA        NOP                             A:2 X:2 Y:D7 P:65 SP:FB CYC:21842
F0D4  20 59 FB  JSR $FB59                       A:2 X:2 Y:D7 P:65 SP:FB CYC:21844
FB59  50 1A     BVC $FB75                       A:2 X:2 Y:D7 P:65 SP:F9 CYC:21850
FB5B  90 18     BCC $FB75                       A:2 X:2 Y:D7 P:65 SP:F9 CYC:21852
FB5D  30 16     BMI $FB75                       A:2 X:2 Y:D7 P:65 SP:F9 CYC:21854
FB5F  C9 02     CMP #$02                        A:2 X:2 Y:D7 P:65 SP:F9 CYC:21856
FB61  D0 12     BNE $FB75                       A:2 X:2 Y:D7 P:67 SP:F9 CYC:21858
FB63  60        RTS                             A:2 X:2 Y:D7 P:67 SP:F9 CYC:21860
F0D7  A5 47     LDA $47 = 4A                    A:2 X:2 Y:D7 P:67 SP:FB CYC:21866
F0D9  C9 4A     CMP #$4A                        A:4A X:2 Y:D7 P:65 SP:FB CYC:21869
F0DB  F0 02     BEQ $F0DF                       A:4A X:2 Y:D7 P:67 SP:FB CYC:21871
F0DF  C8        INY                             A:4A X:2 Y:D7 P:67 SP:FB CYC:21874
F0E0  A9 29     LDA #$29                        A:4A X:2 Y:D8 P:E5 SP:FB CYC:21876
F0E2  85 47     STA $47 = 4A                    A:29 X:2 Y:D8 P:65 SP:FB CYC:21878
F0E4  20 64 FB  JSR $FB64                       A:29 X:2 Y:D8 P:65 SP:FB CYC:21881
FB64  B8        CLV                             A:29 X:2 Y:D8 P:65 SP:F9 CYC:21887
FB65  18        CLC                             A:29 X:2 Y:D8 P:25 SP:F9 CYC:21889
FB66  A9 42     LDA #$42                        A:29 X:2 Y:D8 P:24 SP:F9 CYC:21891
FB68  60        RTS                             A:42 X:2 Y:D8 P:24 SP:F9 CYC:21893
F0E7  27 47    *RLA $47 = 29                    A:42 X:2 Y:D8 P:24 SP:FB CYC:21899
F0E9  EA        NOP                             A:42 X:2 Y:D8 P:24 SP:FB CYC:21904
F0EA  EA        NOP                             A:42 X:2 Y:D8 P:24 SP:FB CYC:21906
F0EB  EA        NOP                             A:42 X:2 Y:D8 P:24 SP:FB CYC:21908
F0EC  EA        NOP                             A:42 X:2 Y:D8 P:24 SP:FB CYC:21910
F0ED  20 69 FB  JSR $FB69                       A:42 X:2 Y:D8 P:24 SP:FB CYC:21912
FB69  70 0A     BVS $FB75                       A:42 X:2 Y:D8 P:24 SP:F9 CYC:21918
FB6B  F0 08     BEQ $FB75                       A:42 X:2 Y:D8 P:24 SP:F9 CYC:21920
FB6D  30 06     BMI $FB75                       A:42 X:2 Y:D8 P:24 SP:F9 CYC:21922
FB6F  B0 04     BCS $FB75                       A:42 X:2 Y:D8 P:24 SP:F9 CYC:21924
FB71  C9 42     CMP #$42                        A:42 X:2 Y:D8 P:24 SP:F9 CYC:21926
FB73  F0 02     BEQ $FB77                       A:42 X:2 Y:D8 P:27 SP:F9 CYC:21928
FB77  60        RTS                             A:42 X:2 Y:D8 P:27 SP:F9 CYC:21931
F0F0  A5 47     LDA $47 = 52                    A:42 X:2 Y:D8 P:27 SP:FB CYC:21937
F0F2  C9 52     CMP #$52                        A:52 X:2 Y:D8 P:25 SP:FB CYC:21940
F0F4  F0 02     BEQ $F0F8                       A:52 X:2 Y:D8 P:27 SP:FB CYC:21942
F0F8  C8        INY                             A:52 X:2 Y:D8 P:27 SP:FB CYC:21945
F0F9  A9 37     LDA #$37                        A:52 X:2 Y:D9 P:A5 SP:FB CYC:21947
F0FB  85 47     STA $47 = 52                    A:37 X:2 Y:D9 P:25 SP:FB CYC:21949
F0FD  20 68 FA  JSR $FA68                       A:37 X:2 Y:D9 P:25 SP:FB CYC:21952
FA68  24 01     BIT $01 = FF                    A:37 X:2 Y:D9 P:25 SP:F9 CYC:21958
FA6A  38        SEC                             A:37 X:2 Y:D9 P:E5 SP:F9 CYC:21961
FA6B  A9 75     LDA #$75                        A:37 X:2 Y:D9 P:E5 SP:F9 CYC:21963
FA6D  60        RTS                             A:75 X:2 Y:D9 P:65 SP:F9 CYC:21965
F100  27 47    *RLA $47 = 37                    A:75 X:2 Y:D9 P:65 SP:FB CYC:21971
F102  EA        NOP                             A:65 X:2 Y:D9 P:64 SP:FB CYC:21976
F103  EA        NOP                             A:65 X:2 Y:D9 P:64 SP:FB CYC:21978
F104  EA        NOP                             A:65 X:2 Y:D9 P:64 SP:FB CYC:21980
F105  EA        NOP                             A:65 X:2 Y:D9 P:64 SP:FB CYC:21982
F106  20 6E FA  JSR $FA6E                       A:65 X:2 Y:D9 P:64 SP:FB CYC:21984
FA6E  50 76     BVC $FAE6                       A:65 X:2 Y:D9 P:64 SP:F9 CYC:21990
FA70  F0 74     BEQ $FAE6                       A:65 X:2 Y:D9 P:64 SP:F9 CYC:21992
FA72  30 72     BMI $FAE6                       A:65 X:2 Y:D9 P:64 SP:F9 CYC:21994
FA74  B0 70     BCS $FAE6                       A:65 X:2 Y:D9 P:64 SP:F9 CYC:21996
FA76  C9 65     CMP #$65                        A:65 X:2 Y:D9 P:64 SP:F9 CYC:21998
FA78  D0 6C     BNE $FAE6                       A:65 X:2 Y:D9 P:67 SP:F9 CYC:22000
FA7A  60        RTS                             A:65 X:2 Y:D9 P:67 SP:F9 CYC:22002
F109  A5 47     LDA $47 = 6F                    A:65 X:2 Y:D9 P:67 SP:FB CYC:22008
F10B  C9 6F     CMP #$6F                        A:6F X:2 Y:D9 P:65 SP:FB CYC:22011
F10D  F0 02     BEQ $F111                       A:6F X:2 Y:D9 P:67 SP:FB CYC:22013
F111  C8        INY                             A:6F X:2 Y:D9 P:67 SP:FB CYC:22016
F112  A9 A5     LDA #$A5                        A:6F X:2 Y:DA P:E5 SP:FB CYC:22018
F114  8D 47 06  STA $0647 = 6F                  A:A5 X:2 Y:DA P:E5 SP:FB CYC:22020
F117  20 53 FB  JSR $FB53                       A:A5 X:2 Y:DA P:E5 SP:FB CYC:22024
FB53  24 01     BIT $01 = FF                    A:A5 X:2 Y:DA P:E5 SP:F9 CYC:22030
FB55  18        CLC                             A:A5 X:2 Y:DA P:E5 SP:F9 CYC:22033
FB56  A9 B3     LDA #$B3                        A:A5 X:2 Y:DA P:E4 SP:F9 CYC:22035
FB58  60        RTS                             A:B3 X:2 Y:DA P:E4 SP:F9 CYC:22037
F11A  2F 47 06 *RLA $0647 = A5                  A:B3 X:2 Y:DA P:E4 SP:FB CYC:22043
F11D  EA        NOP                             A:2 X:2 Y:DA P:65 SP:FB CYC:22049
F11E  EA        NOP                             A:2 X:2 Y:DA P:65 SP:FB CYC:22051
F11F  EA        NOP                             A:2 X:2 Y:DA P:65 SP:FB CYC:22053
F120  EA        NOP                             A:2 X:2 Y:DA P:65 SP:FB CYC:22055
F121  20 59 FB  JSR $FB59                       A:2 X:2 Y:DA P:65 SP:FB CYC:22057
FB59  50 1A     BVC $FB75                       A:2 X:2 Y:DA P:65 SP:F9 CYC:22063
FB5B  90 18     BCC $FB75                       A:2 X:2 Y:DA P:65 SP:F9 CYC:22065
FB5D  30 16     BMI $FB75                       A:2 X:2 Y:DA P:65 SP:F9 CYC:22067
FB5F  C9 02     CMP #$02                        A:2 X:2 Y:DA P:65 SP:F9 CYC:22069
FB61  D0 12     BNE $FB75                       A:2 X:2 Y:DA P:67 SP:F9 CYC:22071
FB63  60        RTS                             A:2 X:2 Y:DA P:67 SP:F9 CYC:22073
F124  AD 47 06  LDA $0647 = 4A                  A:2 X:2 Y:DA P:67 SP:FB CYC:22079
F127  C9 4A     CMP #$4A                        A:4A X:2 Y:DA P:65 SP:FB CYC:22083
F129  F0 02     BEQ $F12D                       A:4A X:2 Y:DA P:67 SP:FB CYC:22085
F12D  C8        INY                             A:4A X:2 Y:DA P:67 SP:FB CYC:22088
F12E  A9 29     LDA #$29                        A:4A X:2 Y:DB P:E5 SP:FB CYC:22090
F130  8D 47 06  STA $0647 = 4A                  A:29 X:2 Y:DB P:65 SP:FB CYC:22092
F133  20 64 FB  JSR $FB64                       A:29 X:2 Y:DB P:65 SP:FB CYC:22096
FB64  B8        CLV                             A:29 X:2 Y:DB P:65 SP:F9 CYC:22102
FB65  18        CLC                             A:29 X:2 Y:DB P:25 SP:F9 CYC:22104
FB66  A9 42     LDA #$42                        A:29 X:2 Y:DB P:24 SP:F9 CYC:22106
FB68  60        RTS                             A:42 X:2 Y:DB P:24 SP:F9 CYC:22108
F136  2F 47 06 *RLA $0647 = 29                  A:42 X:2 Y:DB P:24 SP:FB CYC:22114
F139  EA        NOP                             A:42 X:2 Y:DB P:24 SP:FB CYC:22120
F13A  EA        NOP                             A:42 X:2 Y:DB P:24 SP:FB CYC:22122
F13B  EA        NOP                             A:42 X:2 Y:DB P:24 SP:FB CYC:22124
F13C  EA        NOP                             A:42 X:2 Y:DB P:24 SP:FB CYC:22126
F13D  20 69 FB  JSR $FB69                       A:42 X:2 Y:DB P:24 SP:FB CYC:22128
FB69  70 0A     BVS $FB75                       A:42 X:2 Y:DB P:24 SP:F9 CYC:22134
FB6B  F0 08     BEQ $FB75                       A:42 X:2 Y:DB P:24 SP:F9 CYC:22136
FB6D  30 06     BMI $FB75                       A:42 X:2 Y:DB P:24 SP:F9 CYC:22138
FB6F  B0 04     BCS $FB75                       A:42 X:2 Y:DB P:24 SP:F9 CYC:22140
FB71  C9 42     CMP #$42                        A:42 X:2 Y:DB P:24 SP:F9 CYC:22142
FB73  F0 02     BEQ $FB77                       A:42 X:2 Y:DB P:27 SP:F9 CYC:22144
FB77  60        RTS                             A:42 X:2 Y:DB P:27 SP:F9 CYC:22147
F140  AD 47 06  LDA $0647 = 52                  A:42 X:2 Y:DB P:27 SP:FB CYC:22153
F143  C9 52     CMP #$52                        A:52 X:2 Y:DB P:25 SP:FB CYC:22157
F145  F0 02     BEQ $F149                       A:52 X:2 Y:DB P:27 SP:FB CYC:22159
F149  C8        INY                             A:52 X:2 Y:DB P:27 SP:FB CYC:22162
F14A  A9 37     LDA #$37                        A:52 X:2 Y:DC P:A5 SP:FB CYC:22164
F14C  8D 47 06  STA $0647 = 52                  A:37 X:2 Y:DC P:25 SP:FB CYC:22166
F14F  20 68 FA  JSR $FA68                       A:37 X:2 Y:DC P:25 SP:FB CYC:22170
FA68  24 01     BIT $01 = FF                    A:37 X:2 Y:DC P:25 SP:F9 CYC:22176
FA6A  38        SEC                             A:37 X:2 Y:DC P:E5 SP:F9 CYC:22179
FA6B  A9 75     LDA #$75                        A:37 X:2 Y:DC P:E5 SP:F9 CYC:22181
FA6D  60        RTS                             A:75 X:2 Y:DC P:65 SP:F9 CYC:22183
F152  2F 47 06 *RLA $0647 = 37                  A:75 X:2 Y:DC P:65 SP:FB CYC:22189
F155  EA        NOP                             A:65 X:2 Y:DC P:64 SP:FB CYC:22195
F156  EA        NOP                             A:65 X:2 Y:DC P:64 SP:FB CYC:22197
F157  EA        NOP                             A:65 X:2 Y:DC P:64 SP:FB CYC:22199
F158  EA        NOP                             A:65 X:2 Y:DC P:64 SP:FB CYC:22201
F159  20 6E FA  JSR $FA6E                       A:65 X:2 Y:DC P:64 SP:FB CYC:22203
FA6E  50 76     BVC $FAE6                       A:65 X:2 Y:DC P:64 SP:F9 CYC:22209
FA70  F0 74     BEQ $FAE6                       A:65 X:2 Y:DC P:64 SP:F9 CYC:22211
FA72  30 72     BMI $FAE6                       A:65 X:2 Y:DC P:64 SP:F9 CYC:22213
FA74  B0 70     BCS $FAE6                       A:65 X:2 Y:DC P:64 SP:F9 CYC:22215
FA76  C9 65     CMP #$65                        A:65 X:2 Y:DC P:64 SP:F9 CYC:22217
FA78  D0 6C     BNE $FAE6                       A:65 X:2 Y:DC P:67 SP:F9 CYC:22219
FA7A  60        RTS                             A:65 X:2 Y:DC P:67 SP:F9 CYC:22221
F15C  AD 47 06  LDA $0647 = 6F                  A:65 X:2 Y:DC P:67 SP:FB CYC:22227
F15F  C9 6F     CMP #$6F                        A:6F X:2 Y:DC P:65 SP:FB CYC:22231
F161  F0 02     BEQ $F165                       A:6F X:2 Y:DC P:67 SP:FB CYC:22233
F165  A9 A5     LDA #$A5                        A:6F X:2 Y:DC P:67 SP:FB CYC:22236
F167  8D 47 06  STA $0647 = 6F                  A:A5 X:2 Y:DC P:E5 SP:FB CYC:22238
F16A  A9 48     LDA #$48                        A:A5 X:2 Y:DC P:E5 SP:FB CYC:22242
F16C  85 45     STA $45 = 48                    A:48 X:2 Y:DC P:65 SP:FB CYC:22244
F16E  A9 05     LDA #$05                        A:48 X:2 Y:DC P:65 SP:FB CYC:22247
F170  85 46     STA $46 = 05                    A:5 X:2 Y:DC P:65 SP:FB CYC:22249
F172  A0 FF     LDY #$FF                        A:5 X:2 Y:DC P:65 SP:FB CYC:22252
F174  20 53 FB  JSR $FB53                       A:5 X:2 Y:FF P:E5 SP:FB CYC:22254
FB53  24 01     BIT $01 = FF                    A:5 X:2 Y:FF P:E5 SP:F9 CYC:22260
FB55  18        CLC                             A:5 X:2 Y:FF P:E5 SP:F9 CYC:22263
FB56  A9 B3     LDA #$B3                        A:5 X:2 Y:FF P:E4 SP:F9 CYC:22265
FB58  60        RTS                             A:B3 X:2 Y:FF P:E4 SP:F9 CYC:22267
F177  33 45    *RLA ($45),Y = 0548 @ 0647 = A5  A:B3 X:2 Y:FF P:E4 SP:FB CYC:22273
F179  EA        NOP                             A:2 X:2 Y:FF P:65 SP:FB CYC:22281
F17A  EA        NOP                             A:2 X:2 Y:FF P:65 SP:FB CYC:22283
F17B  08        PHP                             A:2 X:2 Y:FF P:65 SP:FB CYC:22285
F17C  48        PHA                             A:2 X:2 Y:FF P:65 SP:FA CYC:22288
F17D  A0 DD     LDY #$DD                        A:2 X:2 Y:FF P:65 SP:F9 CYC:22291
F17F  68        PLA                             A:2 X:2 Y:DD P:E5 SP:F9 CYC:22293
F180  28        PLP                             A:2 X:2 Y:DD P:65 SP:FA CYC:22297
F181  20 59 FB  JSR $FB59                       A:2 X:2 Y:DD P:65 SP:FB CYC:22301
FB59  50 1A     BVC $FB75                       A:2 X:2 Y:DD P:65 SP:F9 CYC:22307
FB5B  90 18     BCC $FB75                       A:2 X:2 Y:DD P:65 SP:F9 CYC:22309
FB5D  30 16     BMI $FB75                       A:2 X:2 Y:DD P:65 SP:F9 CYC:22311
FB5F  C9 02     CMP #$02                        A:2 X:2 Y:DD P:65 SP:F9 CYC:22313
FB61  D0 12     BNE $FB75                       A:2 X:2 Y:DD P:67 SP:F9 CYC:22315
FB63  60        RTS                             A:2 X:2 Y:DD P:67 SP:F9 CYC:22317
F184  AD 47 06  LDA $0647 = 4A                  A:2 X:2 Y:DD P:67 SP:FB CYC:22323
F187  C9 4A     CMP #$4A                        A:4A X:2 Y:DD P:65 SP:FB CYC:22327
F189  F0 02     BEQ $F18D                       A:4A X:2 Y:DD P:67 SP:FB CYC:22329
F18D  A0 FF     LDY #$FF                        A:4A X:2 Y:DD P:67 SP:FB CYC:22332
F18F  A9 29     LDA #$29                        A:4A X:2 Y:FF P:E5 SP:FB CYC:22334
F191  8D 47 06  STA $0647 = 4A                  A:29 X:2 Y:FF P:65 SP:FB CYC:22336
F194  20 64 FB  JSR $FB64                       A:29 X:2 Y:FF P:65 SP:FB CYC:22340
FB64  B8        CLV                             A:29 X:2 Y:FF P:65 SP:F9 CYC:22346
FB65  18        CLC                             A:29 X:2 Y:FF P:25 SP:F9 CYC:22348
FB66  A9 42     LDA #$42                        A:29 X:2 Y:FF P:24 SP:F9 CYC:22350
FB68  60        RTS                             A:42 X:2 Y:FF P:24 SP:F9 CYC:22352
F197  33 45    *RLA ($45),Y = 0548 @ 0647 = 29  A:42 X:2 Y:FF P:24 SP:FB CYC:22358
F199  EA        NOP                             A:42 X:2 Y:FF P:24 SP:FB CYC:22366
F19A  EA        NOP                             A:42 X:2 Y:FF P:24 SP:FB CYC:22368
F19B  08        PHP                             A:42 X:2 Y:FF P:24 SP:FB CYC:22370
F19C  48        PHA                             A:42 X:2 Y:FF P:24 SP:FA CYC:22373
F19D  A0 DE     LDY #$DE                        A:42 X:2 Y:FF P:24 SP:F9 CYC:22376
F19F  68        PLA                             A:42 X:2 Y:DE P:A4 SP:F9 CYC:22378
F1A0  28        PLP                             A:42 X:2 Y:DE P:24 SP:FA CYC:22382
F1A1  20 69 FB  JSR $FB69                       A:42 X:2 Y:DE P:24 SP:FB CYC:22386
FB69  70 0A     BVS $FB75                       A:42 X:2 Y:DE P:24 SP:F9 CYC:22392
FB6B  F0 08     BEQ $FB75                       A:42 X:2 Y:DE P:24 SP:F9 CYC:22394
FB6D  30 06     BMI $FB75                       A:42 X:2 Y:DE P:24 SP:F9 CYC:22396
FB6F  B0 04     BCS $FB75                       A:42 X:2 Y:DE P:24 SP:F9 CYC:22398
FB71  C9 42     CMP #$42                        A:42 X:2 Y:DE P:24 SP:F9 CYC:22400
FB73  F0 02     BEQ $FB77                       A:42 X:2 Y:DE P:27 SP:F9 CYC:22402
FB77  60        RTS                             A:42 X:2 Y:DE P:27 SP:F9 CYC:22405
F1A4  AD 47 06  LDA $0647 = 52                  A:42 X:2 Y:DE P:27 SP:FB CYC:22411
F1A7  C9 52     CMP #$52                        A:52 X:2 Y:DE P:25 SP:FB CYC:22415
F1A9  F0 02     BEQ $F1AD                       A:52 X:2 Y:DE P:27 SP:FB CYC:22417
F1AD  A0 FF     LDY #$FF                        A:52 X:2 Y:DE P:27 SP:FB CYC:22420
F1AF  A9 37     LDA #$37                        A:52 X:2 Y:FF P:A5 SP:FB CYC:22422
F1B1  8D 47 06  STA $0647 = 52                  A:37 X:2 Y:FF P:25 SP:FB CYC:22424
F1B4  20 68 FA  JSR $FA68                       A:37 X:2 Y:FF P:25 SP:FB CYC:22428
FA68  24 01     BIT $01 = FF                    A:37 X:2 Y:FF P:25 SP:F9 CYC:22434
FA6A  38        SEC                             A:37 X:2 Y:FF P:E5 SP:F9 CYC:22437
FA6B  A9 75     LDA #$75                        A:37 X:2 Y:FF P:E5 SP:F9 CYC:22439
FA6D  60        RTS                             A:75 X:2 Y:FF P:65 SP:F9 CYC:22441
F1B7  33 45    *RLA ($45),Y = 0548 @ 0647 = 37  A:75 X:2 Y:FF P:65 SP:FB CYC:22447
F1B9  EA        NOP                             A:65 X:2 Y:FF P:64 SP:FB CYC:22455
F1BA  EA        NOP                             A:65 X:2 Y:FF P:64 SP:FB CYC:22457
F1BB  08        PHP                             A:65 X:2 Y:FF P:64 SP:FB CYC:22459
F1BC  48        PHA                             A:65 X:2 Y:FF P:64 SP:FA CYC:22462
F1BD  A0 DF     LDY #$DF                        A:65 X:2 Y:FF P:64 SP:F9 CYC:22465
F1BF  68        PLA                             A:65 X:2 Y:DF P:E4 SP:F9 CYC:22467
F1C0  28        PLP                             A:65 X:2 Y:DF P:64 SP:FA CYC:22471
F1C1  20 6E FA  JSR $FA6E                       A:65 X:2 Y:DF P:64 SP:FB CYC:22475
FA6E  50 76     BVC $FAE6                       A:65 X:2 Y:DF P:64 SP:F9 CYC:22481
FA70  F0 74     BEQ $FAE6                       A:65 X:2 Y:DF P:64 SP:F9 CYC:22483
FA72  30 72     BMI $FAE6                       A:65 X:2 Y:DF P:64 SP:F9 CYC:22485
FA74  B0 70     BCS $FAE6                       A:65 X:2 Y:DF P:64 SP:F9 CYC:22487
FA76  C9 65     CMP #$65                        A:65 X:2 Y:DF P:64 SP:F9 CYC:22489
FA78  D0 6C     BNE $FAE6                       A:65 X:2 Y:DF P:67 SP:F9 CYC:22491
FA7A  60        RTS                             A:65 X:2 Y:DF P:67 SP:F9 CYC:22493
F1C4  AD 47 06  LDA $0647 = 6F                  A:65 X:2 Y:DF P:67 SP:FB CYC:22499
F1C7  C9 6F     CMP #$6F                        A:6F X:2 Y:DF P:65 SP:FB CYC:22503
F1C9  F0 02     BEQ $F1CD                       A:6F X:2 Y:DF P:67 SP:FB CYC:22505
F1CD  A0 E0     LDY #$E0                        A:6F X:2 Y:DF P:67 SP:FB CYC:22508
F1CF  A2 FF     LDX #$FF                        A:6F X:2 Y:E0 P:E5 SP:FB CYC:22510
F1D1  A9 A5     LDA #$A5                        A:6F X:FF Y:E0 P:E5 SP:FB CYC:22512
F1D3  85 47     STA $47 = 6F                    A:A5 X:FF Y:E0 P:E5 SP:FB CYC:22514
F1D5  20 53 FB  JSR $FB53                       A:A5 X:FF Y:E0 P:E5 SP:FB CYC:22517
FB53  24 01     BIT $01 = FF                    A:A5 X:FF Y:E0 P:E5 SP:F9 CYC:22523
FB55  18        CLC                             A:A5 X:FF Y:E0 P:E5 SP:F9 CYC:22526
FB56  A9 B3     LDA #$B3                        A:A5 X:FF Y:E0 P:E4 SP:F9 CYC:22528
FB58  60        RTS                             A:B3 X:FF Y:E0 P:E4 SP:F9 CYC:22530
F1D8  37 48    *RLA $48,X @ 47 = A5             A:B3 X:FF Y:E0 P:E4 SP:FB CYC:22536
F1DA  EA        NOP                             A:2 X:FF Y:E0 P:65 SP:FB CYC:22542
F1DB  EA        NOP                             A:2 X:FF Y:E0 P:65 SP:FB CYC:22544
F1DC  EA        NOP                             A:2 X:FF Y:E0 P:65 SP:FB CYC:22546
F1DD  EA        NOP                             A:2 X:FF Y:E0 P:65 SP:FB CYC:22548
F1DE  20 59 FB  JSR $FB59                       A:2 X:FF Y:E0 P:65 SP:FB CYC:22550
FB59  50 1A     BVC $FB75                       A:2 X:FF Y:E0 P:65 SP:F9 CYC:22556
FB5B  90 18     BCC $FB75                       A:2 X:FF Y:E0 P:65 SP:F9 CYC:22558
FB5D  30 16     BMI $FB75                       A:2 X:FF Y:E0 P:65 SP:F9 CYC:22560
FB5F  C9 02     CMP #$02                        A:2 X:FF Y:E0 P:65 SP:F9 CYC:22562
FB61  D0 12     BNE $FB75                       A:2 X:FF Y:E0 P:67 SP:F9 CYC:22564
FB63  60        RTS                             A:2 X:FF Y:E0 P:67 SP:F9 CYC:22566
F1E1  A5 47     LDA $47 = 4A                    A:2 X:FF Y:E0 P:67 SP:FB CYC:22572
F1E3  C9 4A     CMP #$4A                        A:4A X:FF Y:E0 P:65 SP:FB CYC:22575
F1E5  F0 02     BEQ $F1E9                       A:4A X:FF Y:E0 P:67 SP:FB CYC:22577
F1E9  C8        INY                             A:4A X:FF Y:E0 P:67 SP:FB CYC:22580
F1EA  A9 29     LDA #$29                        A:4A X:FF Y:E1 P:E5 SP:FB CYC:22582
F1EC  85 47     STA $47 = 4A                    A:29 X:FF Y:E1 P:65 SP:FB CYC:22584
F1EE  20 64 FB  JSR $FB64                       A:29 X:FF Y:E1 P:65 SP:FB CYC:22587
FB64  B8        CLV                             A:29 X:FF Y:E1 P:65 SP:F9 CYC:22593
FB65  18        CLC                             A:29 X:FF Y:E1 P:25 SP:F9 CYC:22595
FB66  A9 42     LDA #$42                        A:29 X:FF Y:E1 P:24 SP:F9 CYC:22597
FB68  60        RTS                             A:42 X:FF Y:E1 P:24 SP:F9 CYC:22599
F1F1  37 48    *RLA $48,X @ 47 = 29             A:42 X:FF Y:E1 P:24 SP:FB CYC:22605
F1F3  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB CYC:22611
F1F4  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB CYC:22613
F1F5  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB CYC:22615
F1F6  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB CYC:22617
F1F7  20 69 FB  JSR $FB69                       A:42 X:FF Y:E1 P:24 SP:FB CYC:22619
FB69  70 0A     BVS $FB75                       A:42 X:FF Y:E1 P:24 SP:F9 CYC:22625
FB6B  F0 08     BEQ $FB75                       A:42 X:FF Y:E1 P:24 SP:F9 CYC:22627
FB6D  30 06     BMI $FB75                       A:42 X:FF Y:E1 P:24 SP:F9 CYC:22629
FB6F  B0 04     BCS $FB75                       A:42 X:FF Y:E1 P:24 SP:F9 CYC:22631
FB71  C9 42     CMP #$42                        A:42 X:FF Y:E1 P:24 SP:F9 CYC:22633
FB73  F0 02     BEQ $FB77                       A:42 X:FF Y:E1 P:27 SP:F9 CYC:22635
FB77  60        RTS                             A:42 X:FF Y:E1 P:27 SP:F9 CYC:22638
F1FA  A5 47     LDA $47 = 52                    A:42 X:FF Y:E1 P:27 SP:FB CYC:22644
F1FC  C9 52     CMP #$52                        A:52 X:FF Y:E1 P:25 SP:FB CYC:22647
F1FE  F0 02     BEQ $F202                       A:52 X:FF Y:E1 P:27 SP:FB CYC:22649
F202  C8        INY                             A:52 X:FF Y:E1 P:27 SP:FB CYC:22652
F203  A9 37     LDA #$37                        A:52 X:FF Y:E2 P:A5 SP:FB CYC:22654
F205  85 47     STA $47 = 52                    A:37 X:FF Y:E2 P:25 SP:FB CYC:22656
F207  20 68 FA  JSR $FA68                       A:37 X:FF Y:E2 P:25 SP:FB CYC:22659
FA68  24 01     BIT $01 = FF                    A:37 X:FF Y:E2 P:25 SP:F9 CYC:22665
FA6A  38        SEC                             A:37 X:FF Y:E2 P:E5 SP:F9 CYC:22668
FA6B  A9 75     LDA #$75                        A:37 X:FF Y:E2 P:E5 SP:F9 CYC:22670
FA6D  60        RTS                             A:75 X:FF Y:E2 P:65 SP:F9 CYC:22672
F20A  37 48    *RLA $48,X @ 47 = 37             A:75 X:FF Y:E2 P:65 SP:FB CYC:22678
F20C  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB CYC:22684
F20D  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB CYC:22686
F20E  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB CYC:22688
F20F  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB CYC:22690
F210  20 6E FA  JSR $FA6E                       A:65 X:FF Y:E2 P:64 SP:FB CYC:22692
FA6E  50 76     BVC $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9 CYC:22698
FA70  F0 74     BEQ $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9 CYC:22700
FA72  30 72     BMI $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9 CYC:22702
FA74  B0 70     BCS $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9 CYC:22704
FA76  C9 65     CMP #$65                        A:65 X:FF Y:E2 P:64 SP:F9 CYC:22706
FA78  D0 6C     BNE $FAE6                       A:65 X:FF Y:E2 P:67 SP:F9 CYC:22708
FA7A  60        RTS                             A:65 X:FF Y:E2 P:67 SP:F9 CYC:22710
F213  A5 47     LDA $47 = 6F                    A:65 X:FF Y:E2 P:67 SP:FB CYC:22716
F215  C9 6F     CMP #$6F                        A:6F X:FF Y:E2 P:65 SP:FB CYC:22719
F217  F0 02     BEQ $F21B                       A:6F X:FF Y:E2 P:67 SP:FB CYC:22721
F21B  A9 A5     LDA #$A5                        A:6F X:FF Y:E2 P:67 SP:FB CYC:22724
F21D  8D 47 06  STA $0647 = 6F                  A:A5 X:FF Y:E2 P:E5 SP:FB CYC:22726
F220  A0 FF     LDY #$FF                        A:A5 X:FF Y:E2 P:E5 SP:FB CYC:22730
F222  20 53 FB  JSR $FB53                       A:A5 X:FF Y:FF P:E5 SP:FB CYC:22732
FB53  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:E5 SP:F9 CYC:22738
FB55  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9 CYC:22741
FB56  A9 B3     LDA #$B3                        A:A5 X:FF Y:FF P:E4 SP:F9 CYC:22743
FB58  60        RTS                             A:B3 X:FF Y:FF P:E4 SP:F9 CYC:22745
F225  3B 48 05 *RLA $0548,Y @ 0647 = A5         A:B3 X:FF Y:FF P:E4 SP:FB CYC:22751
F228  EA        NOP                             A:2 X:FF Y:FF P:65 SP:FB CYC:22758
F229  EA        NOP                             A:2 X:FF Y:FF P:65 SP:FB CYC:22760
F22A  08        PHP                             A:2 X:FF Y:FF P:65 SP:FB CYC:22762
F22B  48        PHA                             A:2 X:FF Y:FF P:65 SP:FA CYC:22765
F22C  A0 E3     LDY #$E3                        A:2 X:FF Y:FF P:65 SP:F9 CYC:22768
F22E  68        PLA                             A:2 X:FF Y:E3 P:E5 SP:F9 CYC:22770
F22F  28        PLP                             A:2 X:FF Y:E3 P:65 SP:FA CYC:22774
F230  20 59 FB  JSR $FB59                       A:2 X:FF Y:E3 P:65 SP:FB CYC:22778
FB59  50 1A     BVC $FB75                       A:2 X:FF Y:E3 P:65 SP:F9 CYC:22784
FB5B  90 18     BCC $FB75                       A:2 X:FF Y:E3 P:65 SP:F9 CYC:22786
FB5D  30 16     BMI $FB75                       A:2 X:FF Y:E3 P:65 SP:F9 CYC:22788
FB5F  C9 02     CMP #$02                        A:2 X:FF Y:E3 P:65 SP:F9 CYC:22790
FB61  D0 12     BNE $FB75                       A:2 X:FF Y:E3 P:67 SP:F9 CYC:22792
FB63  60        RTS                             A:2 X:FF Y:E3 P:67 SP:F9 CYC:22794
F233  AD 47 06  LDA $0647 = 4A                  A:2 X:FF Y:E3 P:67 SP:FB CYC:22800
F236  C9 4A     CMP #$4A                        A:4A X:FF Y:E3 P:65 SP:FB CYC:22804
F238  F0 02     BEQ $F23C                       A:4A X:FF Y:E3 P:67 SP:FB CYC:22806
F23C  A0 FF     LDY #$FF                        A:4A X:FF Y:E3 P:67 SP:FB CYC:22809
F23E  A9 29     LDA #$29                        A:4A X:FF Y:FF P:E5 SP:FB CYC:22811
F240  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:FF P:65 SP:FB CYC:22813
F243  20 64 FB  JSR $FB64                       A:29 X:FF Y:FF P:65 SP:FB CYC:22817
FB64  B8        CLV                             A:29 X:FF Y:FF P:65 SP:F9 CYC:22823
FB65  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9 CYC:22825
FB66  A9 42     LDA #$42                        A:29 X:FF Y:FF P:24 SP:F9 CYC:22827
FB68  60        RTS                             A:42 X:FF Y:FF P:24 SP:F9 CYC:22829
F246  3B 48 05 *RLA $0548,Y @ 0647 = 29         A:42 X:FF Y:FF P:24 SP:FB CYC:22835
F249  EA        NOP                             A:42 X:FF Y:FF P:24 SP:FB CYC:22842
F24A  EA        NOP                             A:42 X:FF Y:FF P:24 SP:FB CYC:22844
F24B  08        PHP                             A:42 X:FF Y:FF P:24 SP:FB CYC:22846
F24C  48        PHA                             A:42 X:FF Y:FF P:24 SP:FA CYC:22849
F24D  A0 E4     LDY #$E4                        A:42 X:FF Y:FF P:24 SP:F9 CYC:22852
F24F  68        PLA                             A:42 X:FF Y:E4 P:A4 SP:F9 CYC:22854
F250  28        PLP                             A:42 X:FF Y:E4 P:24 SP:FA CYC:22858
F251  20 69 FB  JSR $FB69                       A:42 X:FF Y:E4 P:24 SP:FB CYC:22862
FB69  70 0A     BVS $FB75                       A:42 X:FF Y:E4 P:24 SP:F9 CYC:22868
FB6B  F0 08     BEQ $FB75                       A:42 X:FF Y:E4 P:24 SP:F9 CYC:22870
FB6D  30 06     BMI $FB75                       A:42 X:FF Y:E4 P:24 SP:F9 CYC:22872
FB6F  B0 04     BCS $FB75                       A:42 X:FF Y:E4 P:24 SP:F9 CYC:22874
FB71  C9 42     CMP #$42                        A:42 X:FF Y:E4 P:24 SP:F9 CYC:22876
FB73  F0 02     BEQ $FB77                       A:42 X:FF Y:E4 P:27 SP:F9 CYC:22878
FB77  60        RTS                             A:42 X:FF Y:E4 P:27 SP:F9 CYC:22881
F254  AD 47 06  LDA $0647 = 52                  A:42 X:FF Y:E4 P:27 SP:FB CYC:22887
F257  C9 52     CMP #$52                        A:52 X:FF Y:E4 P:25 SP:FB CYC:22891
F259  F0 02     BEQ $F25D                       A:52 X:FF Y:E4 P:27 SP:FB CYC:22893
F25D  A0 FF     LDY #$FF                        A:52 X:FF Y:E4 P:27 SP:FB CYC:22896
F25F  A9 37     LDA #$37                        A:52 X:FF Y:FF P:A5 SP:FB CYC:22898
F261  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:FF P:25 SP:FB CYC:22900
F264  20 68 FA  JSR $FA68                       A:37 X:FF Y:FF P:25 SP:FB CYC:22904
FA68  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9 CYC:22910
FA6A  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9 CYC:22913
FA6B  A9 75     LDA #$75                        A:37 X:FF Y:FF P:E5 SP:F9 CYC:22915
FA6D  60        RTS                             A:75 X:FF Y:FF P:65 SP:F9 CYC:22917
F267  3B 48 05 *RLA $0548,Y @ 0647 = 37         A:75 X:FF Y:FF P:65 SP:FB CYC:22923
F26A  EA        NOP                             A:65 X:FF Y:FF P:64 SP:FB CYC:22930
F26B  EA        NOP                             A:65 X:FF Y:FF P:64 SP:FB CYC:22932
F26C  08        PHP                             A:65 X:FF Y:FF P:64 SP:FB CYC:22934
F26D  48        PHA                             A:65 X:FF Y:FF P:64 SP:FA CYC:22937
F26E  A0 E5     LDY #$E5                        A:65 X:FF Y:FF P:64 SP:F9 CYC:22940
F270  68        PLA                             A:65 X:FF Y:E5 P:E4 SP:F9 CYC:22942
F271  28        PLP                             A:65 X:FF Y:E5 P:64 SP:FA CYC:22946
F272  20 6E FA  JSR $FA6E                       A:65 X:FF Y:E5 P:64 SP:FB CYC:22950
FA6E  50 76     BVC $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9 CYC:22956
FA70  F0 74     BEQ $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9 CYC:22958
FA72  30 72     BMI $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9 CYC:22960
FA74  B0 70     BCS $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9 CYC:22962
FA76  C9 65     CMP #$65                        A:65 X:FF Y:E5 P:64 SP:F9 CYC:22964
FA78  D0 6C     BNE $FAE6                       A:65 X:FF Y:E5 P:67 SP:F9 CYC:22966
FA7A  60        RTS                             A:65 X:FF Y:E5 P:67 SP:F9 CYC:22968
F275  AD 47 06  LDA $0647 = 6F                  A:65 X:FF Y:E5 P:67 SP:FB CYC:22974
F278  C9 6F     CMP #$6F                        A:6F X:FF Y:E5 P:65 SP:FB CYC:22978
F27A  F0 02     BEQ $F27E                       A:6F X:FF Y:E5 P:67 SP:FB CYC:22980
F27E  A0 E6     LDY #$E6                        A:6F X:FF Y:E5 P:67 SP:FB CYC:22983
F280  A2 FF     LDX #$FF                        A:6F X:FF Y:E6 P:E5 SP:FB CYC:22985
F282  A9 A5     LDA #$A5                        A:6F X:FF Y:E6 P:E5 SP:FB CYC:22987
F284  8D 47 06  STA $0647 = 6F                  A:A5 X:FF Y:E6 P:E5 SP:FB CYC:22989
F287  20 53 FB  JSR $FB53                       A:A5 X:FF Y:E6 P:E5 SP:FB CYC:22993
FB53  24 01     BIT $01 = FF                    A:A5 X:FF Y:E6 P:E5 SP:F9 CYC:22999
FB55  18        CLC                             A:A5 X:FF Y:E6 P:E5 SP:F9 CYC:23002
FB56  A9 B3     LDA #$B3                        A:A5 X:FF Y:E6 P:E4 SP:F9 CYC:23004
FB58  60        RTS                             A:B3 X:FF Y:E6 P:E4 SP:F9 CYC:23006
F28A  3F 48 05 *RLA $0548,X @ 0647 = A5         A:B3 X:FF Y:E6 P:E4 SP:FB CYC:23012
F28D  EA        NOP                             A:2 X:FF Y:E6 P:65 SP:FB CYC:23019
F28E  EA        NOP                             A:2 X:FF Y:E6 P:65 SP:FB CYC:23021
F28F  EA        NOP                             A:2 X:FF Y:E6 P:65 SP:FB CYC:23023
F290  EA        NOP                             A:2 X:FF Y:E6 P:65 SP:FB CYC:23025
F291  20 59 FB  JSR $FB59                       A:2 X:FF Y:E6 P:65 SP:FB CYC:23027
FB59  50 1A     BVC $FB75                       A:2 X:FF Y:E6 P:65 SP:F9 CYC:23033
FB5B  90 18     BCC $FB75                       A:2 X:FF Y:E6 P:65 SP:F9 CYC:23035
FB5D  30 16     BMI $FB75                       A:2 X:FF Y:E6 P:65 SP:F9 CYC:23037
FB5F  C9 02     CMP #$02                        A:2 X:FF Y:E6 P:65 SP:F9 CYC:23039
FB61  D0 12     BNE $FB75                       A:2 X:FF Y:E6 P:67 SP:F9 CYC:23041
FB63  60        RTS                             A:2 X:FF Y:E6 P:67 SP:F9 CYC:23043
F294  AD 47 06  LDA $0647 = 4A                  A:2 X:FF Y:E6 P:67 SP:FB CYC:23049
F297  C9 4A     CMP #$4A                        A:4A X:FF Y:E6 P:65 SP:FB CYC:23053
F299  F0 02     BEQ $F29D                       A:4A X:FF Y:E6 P:67 SP:FB CYC:23055
F29D  C8        INY                             A:4A X:FF Y:E6 P:67 SP:FB CYC:23058
F29E  A9 29     LDA #$29                        A:4A X:FF Y:E7 P:E5 SP:FB CYC:23060
F2A0  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:E7 P:65 SP:FB CYC:23062
F2A3  20 64 FB  JSR $FB64                       A:29 X:FF Y:E7 P:65 SP:FB CYC:23066
FB64  B8        CLV                             A:29 X:FF Y:E7 P:65 SP:F9 CYC:23072
FB65  18        CLC                             A:29 X:FF Y:E7 P:25 SP:F9 CYC:23074
FB66  A9 42     LDA #$42                        A:29 X:FF Y:E7 P:24 SP:F9 CYC:23076
FB68  60        RTS                             A:42 X:FF Y:E7 P:24 SP:F9 CYC:23078
F2A6  3F 48 05 *RLA $0548,X @ 0647 = 29         A:42 X:FF Y:E7 P:24 SP:FB CYC:23084
F2A9  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB CYC:23091
F2AA  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB CYC:23093
F2AB  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB CYC:23095
F2AC  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB CYC:23097
F2AD  20 69 FB  JSR $FB69                       A:42 X:FF Y:E7 P:24 SP:FB CYC:23099
FB69  70 0A     BVS $FB75                       A:42 X:FF Y:E7 P:24 SP:F9 CYC:23105
FB6B  F0 08     BEQ $FB75                       A:42 X:FF Y:E7 P:24 SP:F9 CYC:23107
FB6D  30 06     BMI $FB75                       A:42 X:FF Y:E7 P:24 SP:F9 CYC:23109
FB6F  B0 04     BCS $FB75                       A:42 X:FF Y:E7 P:24 SP:F9 CYC:23111
FB71  C9 42     CMP #$42                        A:42 X:FF Y:E7 P:24 SP:F9 CYC:23113
FB73  F0 02     BEQ $FB77                       A:42 X:FF Y:E7 P:27 SP:F9 CYC:23115
FB77  60        RTS                             A:42 X:FF Y:E7 P:27 SP:F9 CYC:23118
F2B0  AD 47 06  LDA $0647 = 52                  A:42 X:FF Y:E7 P:27 SP:FB CYC:23124
F2B3  C9 52     CMP #$52                        A:52 X:FF Y:E7 P:25 SP:FB CYC:23128
F2B5  F0 02     BEQ $F2B9                       A:52 X:FF Y:E7 P:27 SP:FB CYC:23130
F2B9  C8        INY                             A:52 X:FF Y:E7 P:27 SP:FB CYC:23133
F2BA  A9 37     LDA #$37                        A:52 X:FF Y:E8 P:A5 SP:FB CYC:23135
F2BC  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:E8 P:25 SP:FB CYC:23137
F2BF  20 68 FA  JSR $FA68                       A:37 X:FF Y:E8 P:25 SP:FB CYC:23141
FA68  24 01     BIT $01 = FF                    A:37 X:FF Y:E8 P:25 SP:F9 CYC:23147
FA6A  38        SEC                             A:37 X:FF Y:E8 P:E5 SP:F9 CYC:23150
FA6B  A9 75     LDA #$75                        A:37 X:FF Y:E8 P:E5 SP:F9 CYC:23152
FA6D  60        RTS                             A:75 X:FF Y:E8 P:65 SP:F9 CYC:23154
F2C2  3F 48 05 *RLA $0548,X @ 0647 = 37         A:75 X:FF Y:E8 P:65 SP:FB CYC:23160
F2C5  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB CYC:23167
F2C6  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB CYC:23169
F2C7  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB CYC:23171
F2C8  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB CYC:23173
F2C9  20 6E FA  JSR $FA6E                       A:65 X:FF Y:E8 P:64 SP:FB CYC:23175
FA6E  50 76     BVC $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9 CYC:23181
FA70  F0 74     BEQ $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9 CYC:23183
FA72  30 72     BMI $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9 CYC:23185
FA74  B0 70     BCS $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9 CYC:23187
FA76  C9 65     CMP #$65                        A:65 X:FF Y:E8 P:64 SP:F9 CYC:23189
FA78  D0 6C     BNE $FAE6                       A:65 X:FF Y:E8 P:67 SP:F9 CYC:23191
FA7A  60        RTS                             A:65 X:FF Y:E8 P:67 SP:F9 CYC:23193
F2CC  AD 47 06  LDA $0647 = 6F                  A:65 X:FF Y:E8 P:67 SP:FB CYC:23199
F2CF  C9 6F     CMP #$6F                        A:6F X:FF Y:E8 P:65 SP:FB CYC:23203
F2D1  F0 02     BEQ $F2D5                       A:6F X:FF Y:E8 P:67 SP:FB CYC:23205
F2D5  60        RTS                             A:6F X:FF Y:E8 P:67 SP:FB CYC:23208
C647  20 D6 F2  JSR $F2D6                       A:6F X:FF Y:E8 P:67 SP:FD CYC:23214
F2D6  A9 FF     LDA #$FF                        A:6F X:FF Y:E8 P:67 SP:FB CYC:23220
F2D8  85 01     STA $01 = FF                    A:FF X:FF Y:E8 P:E5 SP:FB CYC:23222
F2DA  A0 E9     LDY #$E9                        A:FF X:FF Y:E8 P:E5 SP:FB CYC:23225
F2DC  A2 02     LDX #$02                        A:FF X:FF Y:E9 P:E5 SP:FB CYC:23227
F2DE  A9 47     LDA #$47                        A:FF X:2 Y:E9 P:65 SP:FB CYC:23229
F2E0  85 47     STA $47 = 6F                    A:47 X:2 Y:E9 P:65 SP:FB CYC:23231
F2E2  A9 06     LDA #$06                        A:47 X:2 Y:E9 P:65 SP:FB CYC:23234
F2E4  85 48     STA $48 = 06                    A:6 X:2 Y:E9 P:65 SP:FB CYC:23236
F2E6  A9 A5     LDA #$A5                        A:6 X:2 Y:E9 P:65 SP:FB CYC:23239
F2E8  8D 47 06  STA $0647 = 6F                  A:A5 X:2 Y:E9 P:E5 SP:FB CYC:23241
F2EB  20 1D FB  JSR $FB1D                       A:A5 X:2 Y:E9 P:E5 SP:FB CYC:23245
FB1D  24 01     BIT $01 = FF                    A:A5 X:2 Y:E9 P:E5 SP:F9 CYC:23251
FB1F  18        CLC                             A:A5 X:2 Y:E9 P:E5 SP:F9 CYC:23254
FB20  A9 B3     LDA #$B3                        A:A5 X:2 Y:E9 P:E4 SP:F9 CYC:23256
FB22  60        RTS                             A:B3 X:2 Y:E9 P:E4 SP:F9 CYC:23258
F2EE  43 45    *SRE ($45,X) @ 47 = 0647 = A5    A:B3 X:2 Y:E9 P:E4 SP:FB CYC:23264
F2F0  EA        NOP                             A:E1 X:2 Y:E9 P:E5 SP:FB CYC:23272
F2F1  EA        NOP                             A:E1 X:2 Y:E9 P:E5 SP:FB CYC:23274
F2F2  EA        NOP                             A:E1 X:2 Y:E9 P:E5 SP:FB CYC:23276
F2F3  EA        NOP                             A:E1 X:2 Y:E9 P:E5 SP:FB CYC:23278
F2F4  20 23 FB  JSR $FB23                       A:E1 X:2 Y:E9 P:E5 SP:FB CYC:23280
FB23  50 50     BVC $FB75                       A:E1 X:2 Y:E9 P:E5 SP:F9 CYC:23286
FB25  90 4E     BCC $FB75                       A:E1 X:2 Y:E9 P:E5 SP:F9 CYC:23288
FB27  10 4C     BPL $FB75                       A:E1 X:2 Y:E9 P:E5 SP:F9 CYC:23290
FB29  C9 E1     CMP #$E1                        A:E1 X:2 Y:E9 P:E5 SP:F9 CYC:23292
FB2B  D0 48     BNE $FB75                       A:E1 X:2 Y:E9 P:67 SP:F9 CYC:23294
FB2D  60        RTS                             A:E1 X:2 Y:E9 P:67 SP:F9 CYC:23296
F2F7  AD 47 06  LDA $0647 = 52                  A:E1 X:2 Y:E9 P:67 SP:FB CYC:23302
F2FA  C9 52     CMP #$52                        A:52 X:2 Y:E9 P:65 SP:FB CYC:23306
F2FC  F0 02     BEQ $F300                       A:52 X:2 Y:E9 P:67 SP:FB CYC:23308
F300  C8        INY                             A:52 X:2 Y:E9 P:67 SP:FB CYC:23312
F301  A9 29     LDA #$29                        A:52 X:2 Y:EA P:E5 SP:FB CYC:23314
F303  8D 47 06  STA $0647 = 52                  A:29 X:2 Y:EA P:65 SP:FB CYC:23316
F306  20 2E FB  JSR $FB2E                       A:29 X:2 Y:EA P:65 SP:FB CYC:23320
FB2E  B8        CLV                             A:29 X:2 Y:EA P:65 SP:F9 CYC:23326
FB2F  18        CLC                             A:29 X:2 Y:EA P:25 SP:F9 CYC:23328
FB30  A9 42     LDA #$42                        A:29 X:2 Y:EA P:24 SP:F9 CYC:23330
FB32  60        RTS                             A:42 X:2 Y:EA P:24 SP:F9 CYC:23332
F309  43 45    *SRE ($45,X) @ 47 = 0647 = 29    A:42 X:2 Y:EA P:24 SP:FB CYC:23338
F30B  EA        NOP                             A:56 X:2 Y:EA P:25 SP:FB CYC:23346
F30C  EA        NOP                             A:56 X:2 Y:EA P:25 SP:FB CYC:23348
F30D  EA        NOP                             A:56 X:2 Y:EA P:25 SP:FB CYC:23350
F30E  EA        NOP                             A:56 X:2 Y:EA P:25 SP:FB CYC:23352
F30F  20 33 FB  JSR $FB33                       A:56 X:2 Y:EA P:25 SP:FB CYC:23354
FB33  70 40     BVS $FB75                       A:56 X:2 Y:EA P:25 SP:F9 CYC:23360
FB35  F0 3E     BEQ $FB75                       A:56 X:2 Y:EA P:25 SP:F9 CYC:23362
FB37  30 3C     BMI $FB75                       A:56 X:2 Y:EA P:25 SP:F9 CYC:23364
FB39  90 3A     BCC $FB75                       A:56 X:2 Y:EA P:25 SP:F9 CYC:23366
FB3B  C9 56     CMP #$56                        A:56 X:2 Y:EA P:25 SP:F9 CYC:23368
FB3D  D0 36     BNE $FB75                       A:56 X:2 Y:EA P:27 SP:F9 CYC:23370
FB3F  60        RTS                             A:56 X:2 Y:EA P:27 SP:F9 CYC:23372
F312  AD 47 06  LDA $0647 = 14                  A:56 X:2 Y:EA P:27 SP:FB CYC:23378
F315  C9 14     CMP #$14                        A:14 X:2 Y:EA P:25 SP:FB CYC:23382
F317  F0 02     BEQ $F31B                       A:14 X:2 Y:EA P:27 SP:FB CYC:23384
F31B  C8        INY                             A:14 X:2 Y:EA P:27 SP:FB CYC:23387
F31C  A9 37     LDA #$37                        A:14 X:2 Y:EB P:A5 SP:FB CYC:23389
F31E  8D 47 06  STA $0647 = 14                  A:37 X:2 Y:EB P:25 SP:FB CYC:23391
F321  20 40 FB  JSR $FB40                       A:37 X:2 Y:EB P:25 SP:FB CYC:23395
FB40  24 01     BIT $01 = FF                    A:37 X:2 Y:EB P:25 SP:F9 CYC:23401
FB42  38        SEC                             A:37 X:2 Y:EB P:E5 SP:F9 CYC:23404
FB43  A9 75     LDA #$75                        A:37 X:2 Y:EB P:E5 SP:F9 CYC:23406
FB45  60        RTS                             A:75 X:2 Y:EB P:65 SP:F9 CYC:23408
F324  43 45    *SRE ($45,X) @ 47 = 0647 = 37    A:75 X:2 Y:EB P:65 SP:FB CYC:23414
F326  EA        NOP                             A:6E X:2 Y:EB P:65 SP:FB CYC:23422
F327  EA        NOP                             A:6E X:2 Y:EB P:65 SP:FB CYC:23424
F328  EA        NOP                             A:6E X:2 Y:EB P:65 SP:FB CYC:23426
F329  EA        NOP                             A:6E X:2 Y:EB P:65 SP:FB CYC:23428
F32A  20 46 FB  JSR $FB46                       A:6E X:2 Y:EB P:65 SP:FB CYC:23430
FB46  50 2D     BVC $FB75                       A:6E X:2 Y:EB P:65 SP:F9 CYC:23436
FB48  F0 2B     BEQ $FB75                       A:6E X:2 Y:EB P:65 SP:F9 CYC:23438
FB4A  30 29     BMI $FB75                       A:6E X:2 Y:EB P:65 SP:F9 CYC:23440
FB4C  90 27     BCC $FB75                       A:6E X:2 Y:EB P:65 SP:F9 CYC:23442
FB4E  C9 6E     CMP #$6E                        A:6E X:2 Y:EB P:65 SP:F9 CYC:23444
FB50  D0 23     BNE $FB75                       A:6E X:2 Y:EB P:67 SP:F9 CYC:23446
FB52  60        RTS                             A:6E X:2 Y:EB P:67 SP:F9 CYC:23448
F32D  AD 47 06  LDA $0647 = 1B                  A:6E X:2 Y:EB P:67 SP:FB CYC:23454
F330  C9 1B     CMP #$1B                        A:1B X:2 Y:EB P:65 SP:FB CYC:23458
F332  F0 02     BEQ $F336                       A:1B X:2 Y:EB P:67 SP:FB CYC:23460
F336  C8        INY                             A:1B X:2 Y:EB P:67 SP:FB CYC:23463
F337  A9 A5     LDA #$A5                        A:1B X:2 Y:EC P:E5 SP:FB CYC:23465
F339  85 47     STA $47 = 47                    A:A5 X:2 Y:EC P:E5 SP:FB CYC:23467
F33B  20 1D FB  JSR $FB1D                       A:A5 X:2 Y:EC P:E5 SP:FB CYC:23470
FB1D  24 01     BIT $01 = FF                    A:A5 X:2 Y:EC P:E5 SP:F9 CYC:23476
FB1F  18        CLC                             A:A5 X:2 Y:EC P:E5 SP:F9 CYC:23479
FB20  A9 B3     LDA #$B3                        A:A5 X:2 Y:EC P:E4 SP:F9 CYC:23481
FB22  60        RTS                             A:B3 X:2 Y:EC P:E4 SP:F9 CYC:23483
F33E  47 47    *SRE $47 = A5                    A:B3 X:2 Y:EC P:E4 SP:FB CYC:23489
F340  EA        NOP                             A:E1 X:2 Y:EC P:E5 SP:FB CYC:23494
F341  EA        NOP                             A:E1 X:2 Y:EC P:E5 SP:FB CYC:23496
F342  EA        NOP                             A:E1 X:2 Y:EC P:E5 SP:FB CYC:23498
F343  EA        NOP                             A:E1 X:2 Y:EC P:E5 SP:FB CYC:23500
F344  20 23 FB  JSR $FB23                       A:E1 X:2 Y:EC P:E5 SP:FB CYC:23502
FB23  50 50     BVC $FB75                       A:E1 X:2 Y:EC P:E5 SP:F9 CYC:23508
FB25  90 4E     BCC $FB75                       A:E1 X:2 Y:EC P:E5 SP:F9 CYC:23510
FB27  10 4C     BPL $FB75                       A:E1 X:2 Y:EC P:E5 SP:F9 CYC:23512
FB29  C9 E1     CMP #$E1                        A:E1 X:2 Y:EC P:E5 SP:F9 CYC:23514
FB2B  D0 48     BNE $FB75                       A:E1 X:2 Y:EC P:67 SP:F9 CYC:23516
FB2D  60        RTS                             A:E1 X:2 Y:EC P:67 SP:F9 CYC:23518
F347  A5 47     LDA $47 = 52                    A:E1 X:2 Y:EC P:67 SP:FB CYC:23524
F349  C9 52     CMP #$52                        A:52 X:2 Y:EC P:65 SP:FB CYC:23527
F34B  F0 02     BEQ $F34F                       A:52 X:2 Y:EC P:67 SP:FB CYC:23529
F34F  C8        INY                             A:52 X:2 Y:EC P:67 SP:FB CYC:23532
F350  A9 29     LDA #$29                        A:52 X:2 Y:ED P:E5 SP:FB CYC:23534
F352  85 47     STA $47 = 52                    A:29 X:2 Y:ED P:65 SP:FB CYC:23536
F354  20 2E FB  JSR $FB2E                       A:29 X:2 Y:ED P:65 SP:FB CYC:23539
FB2E  B8        CLV                             A:29 X:2 Y:ED P:65 SP:F9 CYC:23545
FB2F  18        CLC                             A:29 X:2 Y:ED P:25 SP:F9 CYC:23547
FB30  A9 42     LDA #$42                        A:29 X:2 Y:ED P:24 SP:F9 CYC:23549
FB32  60        RTS                             A:42 X:2 Y:ED P:24 SP:F9 CYC:23551
F357  47 47    *SRE $47 = 29                    A:42 X:2 Y:ED P:24 SP:FB CYC:23557
F359  EA        NOP                             A:56 X:2 Y:ED P:25 SP:FB CYC:23562
F35A  EA        NOP                             A:56 X:2 Y:ED P:25 SP:FB CYC:23564
F35B  EA        NOP                             A:56 X:2 Y:ED P:25 SP:FB CYC:23566
F35C  EA        NOP                             A:56 X:2 Y:ED P:25 SP:FB CYC:23568
F35D  20 33 FB  JSR $FB33                       A:56 X:2 Y:ED P:25 SP:FB CYC:23570
FB33  70 40     BVS $FB75                       A:56 X:2 Y:ED P:25 SP:F9 CYC:23576
FB35  F0 3E     BEQ $FB75                       A:56 X:2 Y:ED P:25 SP:F9 CYC:23578
FB37  30 3C     BMI $FB75                       A:56 X:2 Y:ED P:25 SP:F9 CYC:23580
FB39  90 3A     BCC $FB75                       A:56 X:2 Y:ED P:25 SP:F9 CYC:23582
FB3B  C9 56     CMP #$56                        A:56 X:2 Y:ED P:25 SP:F9 CYC:23584
FB3D  D0 36     BNE $FB75                       A:56 X:2 Y:ED P:27 SP:F9 CYC:23586
FB3F  60        RTS                             A:56 X:2 Y:ED P:27 SP:F9 CYC:23588
F360  A5 47     LDA $47 = 14                    A:56 X:2 Y:ED P:27 SP:FB CYC:23594
F362  C9 14     CMP #$14                        A:14 X:2 Y:ED P:25 SP:FB CYC:23597
F364  F0 02     BEQ $F368                       A:14 X:2 Y:ED P:27 SP:FB CYC:23599
F368  C8        INY                             A:14 X:2 Y:ED P:27 SP:FB CYC:23602
F369  A9 37     LDA #$37                        A:14 X:2 Y:EE P:A5 SP:FB CYC:23604
F36B  85 47     STA $47 = 14                    A:37 X:2 Y:EE P:25 SP:FB CYC:23606
F36D  20 40 FB  JSR $FB40                       A:37 X:2 Y:EE P:25 SP:FB CYC:23609
FB40  24 01     BIT $01 = FF                    A:37 X:2 Y:EE P:25 SP:F9 CYC:23615
FB42  38        SEC                             A:37 X:2 Y:EE P:E5 SP:F9 CYC:23618
FB43  A9 75     LDA #$75                        A:37 X:2 Y:EE P:E5 SP:F9 CYC:23620
FB45  60        RTS                             A:75 X:2 Y:EE P:65 SP:F9 CYC:23622
F370  47 47    *SRE $47 = 37                    A:75 X:2 Y:EE P:65 SP:FB CYC:23628
F372  EA        NOP                             A:6E X:2 Y:EE P:65 SP:FB CYC:23633
F373  EA        NOP                             A:6E X:2 Y:EE P:65 SP:FB CYC:23635
F374  EA        NOP                             A:6E X:2 Y:EE P:65 SP:FB CYC:23637
F375  EA        NOP                             A:6E X:2 Y:EE P:65 SP:FB CYC:23639
F376  20 46 FB  JSR $FB46                       A:6E X:2 Y:EE P:65 SP:FB CYC:23641
FB46  50 2D     BVC $FB75                       A:6E X:2 Y:EE P:65 SP:F9 CYC:23647
FB48  F0 2B     BEQ $FB75                       A:6E X:2 Y:EE P:65 SP:F9 CYC:23649
FB4A  30 29     BMI $FB75                       A:6E X:2 Y:EE P:65 SP:F9 CYC:23651
FB4C  90 27     BCC $FB75                       A:6E X:2 Y:EE P:65 SP:F9 CYC:23653
FB4E  C9 6E     CMP #$6E                        A:6E X:2 Y:EE P:65 SP:F9 CYC:23655
FB50  D0 23     BNE $FB75                       A:6E X:2 Y:EE P:67 SP:F9 CYC:23657
FB52  60        RTS                             A:6E X:2 Y:EE P:67 SP:F9 CYC:23659
F379  A5 47     LDA $47 = 1B                    A:6E X:2 Y:EE P:67 SP:FB CYC:23665
F37B  C9 1B     CMP #$1B                        A:1B X:2 Y:EE P:65 SP:FB CYC:23668
F37D  F0 02     BEQ $F381                       A:1B X:2 Y:EE P:67 SP:FB CYC:23670
F381  C8        INY                             A:1B X:2 Y:EE P:67 SP:FB CYC:23673
F382  A9 A5     LDA #$A5                        A:1B X:2 Y:EF P:E5 SP:FB CYC:23675
F384  8D 47 06  STA $0647 = 1B                  A:A5 X:2 Y:EF P:E5 SP:FB CYC:23677
F387  20 1D FB  JSR $FB1D                       A:A5 X:2 Y:EF P:E5 SP:FB CYC:23681
FB1D  24 01     BIT $01 = FF                    A:A5 X:2 Y:EF P:E5 SP:F9 CYC:23687
FB1F  18        CLC                             A:A5 X:2 Y:EF P:E5 SP:F9 CYC:23690
FB20  A9 B3     LDA #$B3                        A:A5 X:2 Y:EF P:E4 SP:F9 CYC:23692
FB22  60        RTS                             A:B3 X:2 Y:EF P:E4 SP:F9 CYC:23694
F38A  4F 47 06 *SRE $0647 = A5                  A:B3 X:2 Y:EF P:E4 SP:FB CYC:23700
F38D  EA        NOP                             A:E1 X:2 Y:EF P:E5 SP:FB CYC:23706
F38E  EA        NOP                             A:E1 X:2 Y:EF P:E5 SP:FB CYC:23708
F38F  EA        NOP                             A:E1 X:2 Y:EF P:E5 SP:FB CYC:23710
F390  EA        NOP                             A:E1 X:2 Y:EF P:E5 SP:FB CYC:23712
F391  20 23 FB  JSR $FB23                       A:E1 X:2 Y:EF P:E5 SP:FB CYC:23714
FB23  50 50     BVC $FB75                       A:E1 X:2 Y:EF P:E5 SP:F9 CYC:23720
FB25  90 4E     BCC $FB75                       A:E1 X:2 Y:EF P:E5 SP:F9 CYC:23722
FB27  10 4C     BPL $FB75                       A:E1 X:2 Y:EF P:E5 SP:F9 CYC:23724
FB29  C9 E1     CMP #$E1                        A:E1 X:2 Y:EF P:E5 SP:F9 CYC:23726
FB2B  D0 48     BNE $FB75                       A:E1 X:2 Y:EF P:67 SP:F9 CYC:23728
FB2D  60        RTS                             A:E1 X:2 Y:EF P:67 SP:F9 CYC:23730
F394  AD 47 06  LDA $0647 = 52                  A:E1 X:2 Y:EF P:67 SP:FB CYC:23736
F397  C9 52     CMP #$52                        A:52 X:2 Y:EF P:65 SP:FB CYC:23740
F399  F0 02     BEQ $F39D                       A:52 X:2 Y:EF P:67 SP:FB CYC:23742
F39D  C8        INY                             A:52 X:2 Y:EF P:67 SP:FB CYC:23745
F39E  A9 29     LDA #$29                        A:52 X:2 Y:F0 P:E5 SP:FB CYC:23747
F3A0  8D 47 06  STA $0647 = 52                  A:29 X:2 Y:F0 P:65 SP:FB CYC:23749
F3A3  20 2E FB  JSR $FB2E                       A:29 X:2 Y:F0 P:65 SP:FB CYC:23753
FB2E  B8        CLV                             A:29 X:2 Y:F0 P:65 SP:F9 CYC:23759
FB2F  18        CLC                             A:29 X:2 Y:F0 P:25 SP:F9 CYC:23761
FB30  A9 42     LDA #$42                        A:29 X:2 Y:F0 P:24 SP:F9 CYC:23763
FB32  60        RTS                             A:42 X:2 Y:F0 P:24 SP:F9 CYC:23765
F3A6  4F 47 06 *SRE $0647 = 29                  A:42 X:2 Y:F0 P:24 SP:FB CYC:23771
F3A9  EA        NOP                             A:56 X:2 Y:F0 P:25 SP:FB CYC:23777
F3AA  EA        NOP                             A:56 X:2 Y:F0 P:25 SP:FB CYC:23779
F3AB  EA        NOP                             A:56 X:2 Y:F0 P:25 SP:FB CYC:23781
F3AC  EA        NOP                             A:56 X:2 Y:F0 P:25 SP:FB CYC:23783
F3AD  20 33 FB  JSR $FB33                       A:56 X:2 Y:F0 P:25 SP:FB CYC:23785
FB33  70 40     BVS $FB75                       A:56 X:2 Y:F0 P:25 SP:F9 CYC:23791
FB35  F0 3E     BEQ $FB75                       A:56 X:2 Y:F0 P:25 SP:F9 CYC:23793
FB37  30 3C     BMI $FB75                       A:56 X:2 Y:F0 P:25 SP:F9 CYC:23795
FB39  90 3A     BCC $FB75                       A:56 X:2 Y:F0 P:25 SP:F9 CYC:23797
FB3B  C9 56     CMP #$56                        A:56 X:2 Y:F0 P:25 SP:F9 CYC:23799
FB3D  D0 36     BNE $FB75                       A:56 X:2 Y:F0 P:27 SP:F9 CYC:23801
FB3F  60        RTS                             A:56 X:2 Y:F0 P:27 SP:F9 CYC:23803
F3B0  AD 47 06  LDA $0647 = 14                  A:56 X:2 Y:F0 P:27 SP:FB CYC:23809
F3B3  C9 14     CMP #$14                        A:14 X:2 Y:F0 P:25 SP:FB CYC:23813
F3B5  F0 02     BEQ $F3B9                       A:14 X:2 Y:F0 P:27 SP:FB CYC:23815
F3B9  C8        INY                             A:14 X:2 Y:F0 P:27 SP:FB CYC:23818
F3BA  A9 37     LDA #$37                        A:14 X:2 Y:F1 P:A5 SP:FB CYC:23820
F3BC  8D 47 06  STA $0647 = 14                  A:37 X:2 Y:F1 P:25 SP:FB CYC:23822
F3BF  20 40 FB  JSR $FB40                       A:37 X:2 Y:F1 P:25 SP:FB CYC:23826
FB40  24 01     BIT $01 = FF                    A:37 X:2 Y:F1 P:25 SP:F9 CYC:23832
FB42  38        SEC                             A:37 X:2 Y:F1 P:E5 SP:F9 CYC:23835
FB43  A9 75     LDA #$75                        A:37 X:2 Y:F1 P:E5 SP:F9 CYC:23837
FB45  60        RTS                             A:75 X:2 Y:F1 P:65 SP:F9 CYC:23839
F3C2  4F 47 06 *SRE $0647 = 37                  A:75 X:2 Y:F1 P:65 SP:FB CYC:23845
F3C5  EA        NOP                             A:6E X:2 Y:F1 P:65 SP:FB CYC:23851
F3C6  EA        NOP                             A:6E X:2 Y:F1 P:65 SP:FB CYC:23853
F3C7  EA        NOP                             A:6E X:2 Y:F1 P:65 SP:FB CYC:23855
F3C8  EA        NOP                             A:6E X:2 Y:F1 P:65 SP:FB CYC:23857
F3C9  20 46 FB  JSR $FB46                       A:6E X:2 Y:F1 P:65 SP:FB CYC:23859
FB46  50 2D     BVC $FB75                       A:6E X:2 Y:F1 P:65 SP:F9 CYC:23865
FB48  F0 2B     BEQ $FB75                       A:6E X:2 Y:F1 P:65 SP:F9 CYC:23867
FB4A  30 29     BMI $FB75                       A:6E X:2 Y:F1 P:65 SP:F9 CYC:23869
FB4C  90 27     BCC $FB75                       A:6E X:2 Y:F1 P:65 SP:F9 CYC:23871
FB4E  C9 6E     CMP #$6E                        A:6E X:2 Y:F1 P:65 SP:F9 CYC:23873
FB50  D0 23     BNE $FB75                       A:6E X:2 Y:F1 P:67 SP:F9 CYC:23875
FB52  60        RTS                             A:6E X:2 Y:F1 P:67 SP:F9 CYC:23877
F3CC  AD 47 06  LDA $0647 = 1B                  A:6E X:2 Y:F1 P:67 SP:FB CYC:23883
F3CF  C9 1B     CMP #$1B                        A:1B X:2 Y:F1 P:65 SP:FB CYC:23887
F3D1  F0 02     BEQ $F3D5                       A:1B X:2 Y:F1 P:67 SP:FB CYC:23889
F3D5  A9 A5     LDA #$A5                        A:1B X:2 Y:F1 P:67 SP:FB CYC:23892
F3D7  8D 47 06  STA $0647 = 1B                  A:A5 X:2 Y:F1 P:E5 SP:FB CYC:23894
F3DA  A9 48     LDA #$48                        A:A5 X:2 Y:F1 P:E5 SP:FB CYC:23898
F3DC  85 45     STA $45 = 48                    A:48 X:2 Y:F1 P:65 SP:FB CYC:23900
F3DE  A9 05     LDA #$05                        A:48 X:2 Y:F1 P:65 SP:FB CYC:23903
F3E0  85 46     STA $46 = 05                    A:5 X:2 Y:F1 P:65 SP:FB CYC:23905
F3E2  A0 FF     LDY #$FF                        A:5 X:2 Y:F1 P:65 SP:FB CYC:23908
F3E4  20 1D FB  JSR $FB1D                       A:5 X:2 Y:FF P:E5 SP:FB CYC:23910
FB1D  24 01     BIT $01 = FF                    A:5 X:2 Y:FF P:E5 SP:F9 CYC:23916
FB1F  18        CLC                             A:5 X:2 Y:FF P:E5 SP:F9 CYC:23919
FB20  A9 B3     LDA #$B3                        A:5 X:2 Y:FF P:E4 SP:F9 CYC:23921
FB22  60        RTS                             A:B3 X:2 Y:FF P:E4 SP:F9 CYC:23923
F3E7  53 45    *SRE ($45),Y = 0548 @ 0647 = A5  A:B3 X:2 Y:FF P:E4 SP:FB CYC:23929
F3E9  EA        NOP                             A:E1 X:2 Y:FF P:E5 SP:FB CYC:23937
F3EA  EA        NOP                             A:E1 X:2 Y:FF P:E5 SP:FB CYC:23939
F3EB  08        PHP                             A:E1 X:2 Y:FF P:E5 SP:FB CYC:23941
F3EC  48        PHA                             A:E1 X:2 Y:FF P:E5 SP:FA CYC:23944
F3ED  A0 F2     LDY #$F2                        A:E1 X:2 Y:FF P:E5 SP:F9 CYC:23947
F3EF  68        PLA                             A:E1 X:2 Y:F2 P:E5 SP:F9 CYC:23949
F3F0  28        PLP                             A:E1 X:2 Y:F2 P:E5 SP:FA CYC:23953
F3F1  20 23 FB  JSR $FB23                       A:E1 X:2 Y:F2 P:E5 SP:FB CYC:23957
FB23  50 50     BVC $FB75                       A:E1 X:2 Y:F2 P:E5 SP:F9 CYC:23963
FB25  90 4E     BCC $FB75                       A:E1 X:2 Y:F2 P:E5 SP:F9 CYC:23965
FB27  10 4C     BPL $FB75                       A:E1 X:2 Y:F2 P:E5 SP:F9 CYC:23967
FB29  C9 E1     CMP #$E1                        A:E1 X:2 Y:F2 P:E5 SP:F9 CYC:23969
FB2B  D0 48     BNE $FB75                       A:E1 X:2 Y:F2 P:67 SP:F9 CYC:23971
FB2D  60        RTS                             A:E1 X:2 Y:F2 P:67 SP:F9 CYC:23973
F3F4  AD 47 06  LDA $0647 = 52                  A:E1 X:2 Y:F2 P:67 SP:FB CYC:23979
F3F7  C9 52     CMP #$52                        A:52 X:2 Y:F2 P:65 SP:FB CYC:23983
F3F9  F0 02     BEQ $F3FD                       A:52 X:2 Y:F2 P:67 SP:FB CYC:23985
F3FD  A0 FF     LDY #$FF                        A:52 X:2 Y:F2 P:67 SP:FB CYC:23988
F3FF  A9 29     LDA #$29                        A:52 X:2 Y:FF P:E5 SP:FB CYC:23990
F401  8D 47 06  STA $0647 = 52                  A:29 X:2 Y:FF P:65 SP:FB CYC:23992
F404  20 2E FB  JSR $FB2E                       A:29 X:2 Y:FF P:65 SP:FB CYC:23996
FB2E  B8        CLV                             A:29 X:2 Y:FF P:65 SP:F9 CYC:24002
FB2F  18        CLC                             A:29 X:2 Y:FF P:25 SP:F9 CYC:24004
FB30  A9 42     LDA #$42                        A:29 X:2 Y:FF P:24 SP:F9 CYC:24006
FB32  60        RTS                             A:42 X:2 Y:FF P:24 SP:F9 CYC:24008
F407  53 45    *SRE ($45),Y = 0548 @ 0647 = 29  A:42 X:2 Y:FF P:24 SP:FB CYC:24014
F409  EA        NOP                             A:56 X:2 Y:FF P:25 SP:FB CYC:24022
F40A  EA        NOP                             A:56 X:2 Y:FF P:25 SP:FB CYC:24024
F40B  08        PHP                             A:56 X:2 Y:FF P:25 SP:FB CYC:24026
F40C  48        PHA                             A:56 X:2 Y:FF P:25 SP:FA CYC:24029
F40D  A0 F3     LDY #$F3                        A:56 X:2 Y:FF P:25 SP:F9 CYC:24032
F40F  68        PLA                             A:56 X:2 Y:F3 P:A5 SP:F9 CYC:24034
F410  28        PLP                             A:56 X:2 Y:F3 P:25 SP:FA CYC:24038
F411  20 33 FB  JSR $FB33                       A:56 X:2 Y:F3 P:25 SP:FB CYC:24042
FB33  70 40     BVS $FB75                       A:56 X:2 Y:F3 P:25 SP:F9 CYC:24048
FB35  F0 3E     BEQ $FB75                       A:56 X:2 Y:F3 P:25 SP:F9 CYC:24050
FB37  30 3C     BMI $FB75                       A:56 X:2 Y:F3 P:25 SP:F9 CYC:24052
FB39  90 3A     BCC $FB75                       A:56 X:2 Y:F3 P:25 SP:F9 CYC:24054
FB3B  C9 56     CMP #$56                        A:56 X:2 Y:F3 P:25 SP:F9 CYC:24056
FB3D  D0 36     BNE $FB75                       A:56 X:2 Y:F3 P:27 SP:F9 CYC:24058
FB3F  60        RTS                             A:56 X:2 Y:F3 P:27 SP:F9 CYC:24060
F414  AD 47 06  LDA $0647 = 14                  A:56 X:2 Y:F3 P:27 SP:FB CYC:24066
F417  C9 14     CMP #$14                        A:14 X:2 Y:F3 P:25 SP:FB CYC:24070
F419  F0 02     BEQ $F41D                       A:14 X:2 Y:F3 P:27 SP:FB CYC:24072
F41D  A0 FF     LDY #$FF                        A:14 X:2 Y:F3 P:27 SP:FB CYC:24075
F41F  A9 37     LDA #$37                        A:14 X:2 Y:FF P:A5 SP:FB CYC:24077
F421  8D 47 06  STA $0647 = 14                  A:37 X:2 Y:FF P:25 SP:FB CYC:24079
F424  20 40 FB  JSR $FB40                       A:37 X:2 Y:FF P:25 SP:FB CYC:24083
FB40  24 01     BIT $01 = FF                    A:37 X:2 Y:FF P:25 SP:F9 CYC:24089
FB42  38        SEC                             A:37 X:2 Y:FF P:E5 SP:F9 CYC:24092
FB43  A9 75     LDA #$75                        A:37 X:2 Y:FF P:E5 SP:F9 CYC:24094
FB45  60        RTS                             A:75 X:2 Y:FF P:65 SP:F9 CYC:24096
F427  53 45    *SRE ($45),Y = 0548 @ 0647 = 37  A:75 X:2 Y:FF P:65 SP:FB CYC:24102
F429  EA        NOP                             A:6E X:2 Y:FF P:65 SP:FB CYC:24110
F42A  EA        NOP                             A:6E X:2 Y:FF P:65 SP:FB CYC:24112
F42B  08        PHP                             A:6E X:2 Y:FF P:65 SP:FB CYC:24114
F42C  48        PHA                             A:6E X:2 Y:FF P:65 SP:FA CYC:24117
F42D  A0 F4     LDY #$F4                        A:6E X:2 Y:FF P:65 SP:F9 CYC:24120
F42F  68        PLA                             A:6E X:2 Y:F4 P:E5 SP:F9 CYC:24122
F430  28        PLP                             A:6E X:2 Y:F4 P:65 SP:FA CYC:24126
F431  20 46 FB  JSR $FB46                       A:6E X:2 Y:F4 P:65 SP:FB CYC:24130
FB46  50 2D     BVC $FB75                       A:6E X:2 Y:F4 P:65 SP:F9 CYC:24136
FB48  F0 2B     BEQ $FB75                       A:6E X:2 Y:F4 P:65 SP:F9 CYC:24138
FB4A  30 29     BMI $FB75                       A:6E X:2 Y:F4 P:65 SP:F9 CYC:24140
FB4C  90 27     BCC $FB75                       A:6E X:2 Y:F4 P:65 SP:F9 CYC:24142
FB4E  C9 6E     CMP #$6E                        A:6E X:2 Y:F4 P:65 SP:F9 CYC:24144
FB50  D0 23     BNE $FB75                       A:6E X:2 Y:F4 P:67 SP:F9 CYC:24146
FB52  60        RTS                             A:6E X:2 Y:F4 P:67 SP:F9 CYC:24148
F434  AD 47 06  LDA $0647 = 1B                  A:6E X:2 Y:F4 P:67 SP:FB CYC:24154
F437  C9 1B     CMP #$1B                        A:1B X:2 Y:F4 P:65 SP:FB CYC:24158
F439  F0 02     BEQ $F43D                       A:1B X:2 Y:F4 P:67 SP:FB CYC:24160
F43D  A0 F5     LDY #$F5                        A:1B X:2 Y:F4 P:67 SP:FB CYC:24163
F43F  A2 FF     LDX #$FF                        A:1B X:2 Y:F5 P:E5 SP:FB CYC:24165
F441  A9 A5     LDA #$A5                        A:1B X:FF Y:F5 P:E5 SP:FB CYC:24167
F443  85 47     STA $47 = 1B                    A:A5 X:FF Y:F5 P:E5 SP:FB CYC:24169
F445  20 1D FB  JSR $FB1D                       A:A5 X:FF Y:F5 P:E5 SP:FB CYC:24172
FB1D  24 01     BIT $01 = FF                    A:A5 X:FF Y:F5 P:E5 SP:F9 CYC:24178
FB1F  18        CLC                             A:A5 X:FF Y:F5 P:E5 SP:F9 CYC:24181
FB20  A9 B3     LDA #$B3                        A:A5 X:FF Y:F5 P:E4 SP:F9 CYC:24183
FB22  60        RTS                             A:B3 X:FF Y:F5 P:E4 SP:F9 CYC:24185
F448  57 48    *SRE $48,X @ 47 = A5             A:B3 X:FF Y:F5 P:E4 SP:FB CYC:24191
F44A  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB CYC:24197
F44B  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB CYC:24199
F44C  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB CYC:24201
F44D  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB CYC:24203
F44E  20 23 FB  JSR $FB23                       A:E1 X:FF Y:F5 P:E5 SP:FB CYC:24205
FB23  50 50     BVC $FB75                       A:E1 X:FF Y:F5 P:E5 SP:F9 CYC:24211
FB25  90 4E     BCC $FB75                       A:E1 X:FF Y:F5 P:E5 SP:F9 CYC:24213
FB27  10 4C     BPL $FB75                       A:E1 X:FF Y:F5 P:E5 SP:F9 CYC:24215
FB29  C9 E1     CMP #$E1                        A:E1 X:FF Y:F5 P:E5 SP:F9 CYC:24217
FB2B  D0 48     BNE $FB75                       A:E1 X:FF Y:F5 P:67 SP:F9 CYC:24219
FB2D  60        RTS                             A:E1 X:FF Y:F5 P:67 SP:F9 CYC:24221
F451  A5 47     LDA $47 = 52                    A:E1 X:FF Y:F5 P:67 SP:FB CYC:24227
F453  C9 52     CMP #$52                        A:52 X:FF Y:F5 P:65 SP:FB CYC:24230
F455  F0 02     BEQ $F459                       A:52 X:FF Y:F5 P:67 SP:FB CYC:24232
F459  C8        INY                             A:52 X:FF Y:F5 P:67 SP:FB CYC:24235
F45A  A9 29     LDA #$29                        A:52 X:FF Y:F6 P:E5 SP:FB CYC:24237
F45C  85 47     STA $47 = 52                    A:29 X:FF Y:F6 P:65 SP:FB CYC:24239
F45E  20 2E FB  JSR $FB2E                       A:29 X:FF Y:F6 P:65 SP:FB CYC:24242
FB2E  B8        CLV                             A:29 X:FF Y:F6 P:65 SP:F9 CYC:24248
FB2F  18        CLC                             A:29 X:FF Y:F6 P:25 SP:F9 CYC:24250
FB30  A9 42     LDA #$42                        A:29 X:FF Y:F6 P:24 SP:F9 CYC:24252
FB32  60        RTS                             A:42 X:FF Y:F6 P:24 SP:F9 CYC:24254
F461  57 48    *SRE $48,X @ 47 = 29             A:42 X:FF Y:F6 P:24 SP:FB CYC:24260
F463  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB CYC:24266
F464  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB CYC:24268
F465  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB CYC:24270
F466  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB CYC:24272
F467  20 33 FB  JSR $FB33                       A:56 X:FF Y:F6 P:25 SP:FB CYC:24274
FB33  70 40     BVS $FB75                       A:56 X:FF Y:F6 P:25 SP:F9 CYC:24280
FB35  F0 3E     BEQ $FB75                       A:56 X:FF Y:F6 P:25 SP:F9 CYC:24282
FB37  30 3C     BMI $FB75                       A:56 X:FF Y:F6 P:25 SP:F9 CYC:24284
FB39  90 3A     BCC $FB75                       A:56 X:FF Y:F6 P:25 SP:F9 CYC:24286
FB3B  C9 56     CMP #$56                        A:56 X:FF Y:F6 P:25 SP:F9 CYC:24288
FB3D  D0 36     BNE $FB75                       A:56 X:FF Y:F6 P:27 SP:F9 CYC:24290
FB3F  60        RTS                             A:56 X:FF Y:F6 P:27 SP:F9 CYC:24292
F46A  A5 47     LDA $47 = 14                    A:56 X:FF Y:F6 P:27 SP:FB CYC:24298
F46C  C9 14     CMP #$14                        A:14 X:FF Y:F6 P:25 SP:FB CYC:24301
F46E  F0 02     BEQ $F472                       A:14 X:FF Y:F6 P:27 SP:FB CYC:24303
F472  C8        INY                             A:14 X:FF Y:F6 P:27 SP:FB CYC:24306
F473  A9 37     LDA #$37                        A:14 X:FF Y:F7 P:A5 SP:FB CYC:24308
F475  85 47     STA $47 = 14                    A:37 X:FF Y:F7 P:25 SP:FB CYC:24310
F477  20 40 FB  JSR $FB40                       A:37 X:FF Y:F7 P:25 SP:FB CYC:24313
FB40  24 01     BIT $01 = FF                    A:37 X:FF Y:F7 P:25 SP:F9 CYC:24319
FB42  38        SEC                             A:37 X:FF Y:F7 P:E5 SP:F9 CYC:24322
FB43  A9 75     LDA #$75                        A:37 X:FF Y:F7 P:E5 SP:F9 CYC:24324
FB45  60        RTS                             A:75 X:FF Y:F7 P:65 SP:F9 CYC:24326
F47A  57 48    *SRE $48,X @ 47 = 37             A:75 X:FF Y:F7 P:65 SP:FB CYC:24332
F47C  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB CYC:24338
F47D  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB CYC:24340
F47E  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB CYC:24342
F47F  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB CYC:24344
F480  20 46 FB  JSR $FB46                       A:6E X:FF Y:F7 P:65 SP:FB CYC:24346
FB46  50 2D     BVC $FB75                       A:6E X:FF Y:F7 P:65 SP:F9 CYC:24352
FB48  F0 2B     BEQ $FB75                       A:6E X:FF Y:F7 P:65 SP:F9 CYC:24354
FB4A  30 29     BMI $FB75                       A:6E X:FF Y:F7 P:65 SP:F9 CYC:24356
FB4C  90 27     BCC $FB75                       A:6E X:FF Y:F7 P:65 SP:F9 CYC:24358
FB4E  C9 6E     CMP #$6E                        A:6E X:FF Y:F7 P:65 SP:F9 CYC:24360
FB50  D0 23     BNE $FB75                       A:6E X:FF Y:F7 P:67 SP:F9 CYC:24362
FB52  60        RTS                             A:6E X:FF Y:F7 P:67 SP:F9 CYC:24364
F483  A5 47     LDA $47 = 1B                    A:6E X:FF Y:F7 P:67 SP:FB CYC:24370
F485  C9 1B     CMP #$1B                        A:1B X:FF Y:F7 P:65 SP:FB CYC:24373
F487  F0 02     BEQ $F48B                       A:1B X:FF Y:F7 P:67 SP:FB CYC:24375
F48B  A9 A5     LDA #$A5                        A:1B X:FF Y:F7 P:67 SP:FB CYC:24378
F48D  8D 47 06  STA $0647 = 1B                  A:A5 X:FF Y:F7 P:E5 SP:FB CYC:24380
F490  A0 FF     LDY #$FF                        A:A5 X:FF Y:F7 P:E5 SP:FB CYC:24384
F492  20 1D FB  JSR $FB1D                       A:A5 X:FF Y:FF P:E5 SP:FB CYC:24386
FB1D  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:E5 SP:F9 CYC:24392
FB1F  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9 CYC:24395
FB20  A9 B3     LDA #$B3                        A:A5 X:FF Y:FF P:E4 SP:F9 CYC:24397
FB22  60        RTS                             A:B3 X:FF Y:FF P:E4 SP:F9 CYC:24399
F495  5B 48 05 *SRE $0548,Y @ 0647 = A5         A:B3 X:FF Y:FF P:E4 SP:FB CYC:24405
F498  EA        NOP                             A:E1 X:FF Y:FF P:E5 SP:FB CYC:24412
F499  EA        NOP                             A:E1 X:FF Y:FF P:E5 SP:FB CYC:24414
F49A  08        PHP                             A:E1 X:FF Y:FF P:E5 SP:FB CYC:24416
F49B  48        PHA                             A:E1 X:FF Y:FF P:E5 SP:FA CYC:24419
F49C  A0 F8     LDY #$F8                        A:E1 X:FF Y:FF P:E5 SP:F9 CYC:24422
F49E  68        PLA                             A:E1 X:FF Y:F8 P:E5 SP:F9 CYC:24424
F49F  28        PLP                             A:E1 X:FF Y:F8 P:E5 SP:FA CYC:24428
F4A0  20 23 FB  JSR $FB23                       A:E1 X:FF Y:F8 P:E5 SP:FB CYC:24432
FB23  50 50     BVC $FB75                       A:E1 X:FF Y:F8 P:E5 SP:F9 CYC:24438
FB25  90 4E     BCC $FB75                       A:E1 X:FF Y:F8 P:E5 SP:F9 CYC:24440
FB27  10 4C     BPL $FB75                       A:E1 X:FF Y:F8 P:E5 SP:F9 CYC:24442
FB29  C9 E1     CMP #$E1                        A:E1 X:FF Y:F8 P:E5 SP:F9 CYC:24444
FB2B  D0 48     BNE $FB75                       A:E1 X:FF Y:F8 P:67 SP:F9 CYC:24446
FB2D  60        RTS                             A:E1 X:FF Y:F8 P:67 SP:F9 CYC:24448
F4A3  AD 47 06  LDA $0647 = 52                  A:E1 X:FF Y:F8 P:67 SP:FB CYC:24454
F4A6  C9 52     CMP #$52                        A:52 X:FF Y:F8 P:65 SP:FB CYC:24458
F4A8  F0 02     BEQ $F4AC                       A:52 X:FF Y:F8 P:67 SP:FB CYC:24460
F4AC  A0 FF     LDY #$FF                        A:52 X:FF Y:F8 P:67 SP:FB CYC:24463
F4AE  A9 29     LDA #$29                        A:52 X:FF Y:FF P:E5 SP:FB CYC:24465
F4B0  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:FF P:65 SP:FB CYC:24467
F4B3  20 2E FB  JSR $FB2E                       A:29 X:FF Y:FF P:65 SP:FB CYC:24471
FB2E  B8        CLV                             A:29 X:FF Y:FF P:65 SP:F9 CYC:24477
FB2F  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9 CYC:24479
FB30  A9 42     LDA #$42                        A:29 X:FF Y:FF P:24 SP:F9 CYC:24481
FB32  60        RTS                             A:42 X:FF Y:FF P:24 SP:F9 CYC:24483
F4B6  5B 48 05 *SRE $0548,Y @ 0647 = 29         A:42 X:FF Y:FF P:24 SP:FB CYC:24489
F4B9  EA        NOP                             A:56 X:FF Y:FF P:25 SP:FB CYC:24496
F4BA  EA        NOP                             A:56 X:FF Y:FF P:25 SP:FB CYC:24498
F4BB  08        PHP                             A:56 X:FF Y:FF P:25 SP:FB CYC:24500
F4BC  48        PHA                             A:56 X:FF Y:FF P:25 SP:FA CYC:24503
F4BD  A0 F9     LDY #$F9                        A:56 X:FF Y:FF P:25 SP:F9 CYC:24506
F4BF  68        PLA                             A:56 X:FF Y:F9 P:A5 SP:F9 CYC:24508
F4C0  28        PLP                             A:56 X:FF Y:F9 P:25 SP:FA CYC:24512
F4C1  20 33 FB  JSR $FB33                       A:56 X:FF Y:F9 P:25 SP:FB CYC:24516
FB33  70 40     BVS $FB75                       A:56 X:FF Y:F9 P:25 SP:F9 CYC:24522
FB35  F0 3E     BEQ $FB75                       A:56 X:FF Y:F9 P:25 SP:F9 CYC:24524
FB37  30 3C     BMI $FB75                       A:56 X:FF Y:F9 P:25 SP:F9 CYC:24526
FB39  90 3A     BCC $FB75                       A:56 X:FF Y:F9 P:25 SP:F9 CYC:24528
FB3B  C9 56     CMP #$56                        A:56 X:FF Y:F9 P:25 SP:F9 CYC:24530
FB3D  D0 36     BNE $FB75                       A:56 X:FF Y:F9 P:27 SP:F9 CYC:24532
FB3F  60        RTS                             A:56 X:FF Y:F9 P:27 SP:F9 CYC:24534
F4C4  AD 47 06  LDA $0647 = 14                  A:56 X:FF Y:F9 P:27 SP:FB CYC:24540
F4C7  C9 14     CMP #$14                        A:14 X:FF Y:F9 P:25 SP:FB CYC:24544
F4C9  F0 02     BEQ $F4CD                       A:14 X:FF Y:F9 P:27 SP:FB CYC:24546
F4CD  A0 FF     LDY #$FF                        A:14 X:FF Y:F9 P:27 SP:FB CYC:24549
F4CF  A9 37     LDA #$37                        A:14 X:FF Y:FF P:A5 SP:FB CYC:24551
F4D1  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:FF P:25 SP:FB CYC:24553
F4D4  20 40 FB  JSR $FB40                       A:37 X:FF Y:FF P:25 SP:FB CYC:24557
FB40  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9 CYC:24563
FB42  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9 CYC:24566
FB43  A9 75     LDA #$75                        A:37 X:FF Y:FF P:E5 SP:F9 CYC:24568
FB45  60        RTS                             A:75 X:FF Y:FF P:65 SP:F9 CYC:24570
F4D7  5B 48 05 *SRE $0548,Y @ 0647 = 37         A:75 X:FF Y:FF P:65 SP:FB CYC:24576
F4DA  EA        NOP                             A:6E X:FF Y:FF P:65 SP:FB CYC:24583
F4DB  EA        NOP                             A:6E X:FF Y:FF P:65 SP:FB CYC:24585
F4DC  08        PHP                             A:6E X:FF Y:FF P:65 SP:FB CYC:24587
F4DD  48        PHA                             A:6E X:FF Y:FF P:65 SP:FA CYC:24590
F4DE  A0 FA     LDY #$FA                        A:6E X:FF Y:FF P:65 SP:F9 CYC:24593
F4E0  68        PLA                             A:6E X:FF Y:FA P:E5 SP:F9 CYC:24595
F4E1  28        PLP                             A:6E X:FF Y:FA P:65 SP:FA CYC:24599
F4E2  20 46 FB  JSR $FB46                       A:6E X:FF Y:FA P:65 SP:FB CYC:24603
FB46  50 2D     BVC $FB75                       A:6E X:FF Y:FA P:65 SP:F9 CYC:24609
FB48  F0 2B     BEQ $FB75                       A:6E X:FF Y:FA P:65 SP:F9 CYC:24611
FB4A  30 29     BMI $FB75                       A:6E X:FF Y:FA P:65 SP:F9 CYC:24613
FB4C  90 27     BCC $FB75                       A:6E X:FF Y:FA P:65 SP:F9 CYC:24615
FB4E  C9 6E     CMP #$6E                        A:6E X:FF Y:FA P:65 SP:F9 CYC:24617
FB50  D0 23     BNE $FB75                       A:6E X:FF Y:FA P:67 SP:F9 CYC:24619
FB52  60        RTS                             A:6E X:FF Y:FA P:67 SP:F9 CYC:24621
F4E5  AD 47 06  LDA $0647 = 1B                  A:6E X:FF Y:FA P:67 SP:FB CYC:24627
F4E8  C9 1B     CMP #$1B                        A:1B X:FF Y:FA P:65 SP:FB CYC:24631
F4EA  F0 02     BEQ $F4EE                       A:1B X:FF Y:FA P:67 SP:FB CYC:24633
F4EE  A0 FB     LDY #$FB                        A:1B X:FF Y:FA P:67 SP:FB CYC:24636
F4F0  A2 FF     LDX #$FF                        A:1B X:FF Y:FB P:E5 SP:FB CYC:24638
F4F2  A9 A5     LDA #$A5                        A:1B X:FF Y:FB P:E5 SP:FB CYC:24640
F4F4  8D 47 06  STA $0647 = 1B                  A:A5 X:FF Y:FB P:E5 SP:FB CYC:24642
F4F7  20 1D FB  JSR $FB1D                       A:A5 X:FF Y:FB P:E5 SP:FB CYC:24646
FB1D  24 01     BIT $01 = FF                    A:A5 X:FF Y:FB P:E5 SP:F9 CYC:24652
FB1F  18        CLC                             A:A5 X:FF Y:FB P:E5 SP:F9 CYC:24655
FB20  A9 B3     LDA #$B3                        A:A5 X:FF Y:FB P:E4 SP:F9 CYC:24657
FB22  60        RTS                             A:B3 X:FF Y:FB P:E4 SP:F9 CYC:24659
F4FA  5F 48 05 *SRE $0548,X @ 0647 = A5         A:B3 X:FF Y:FB P:E4 SP:FB CYC:24665
F4FD  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB CYC:24672
F4FE  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB CYC:24674
F4FF  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB CYC:24676
F500  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB CYC:24678
F501  20 23 FB  JSR $FB23                       A:E1 X:FF Y:FB P:E5 SP:FB CYC:24680
FB23  50 50     BVC $FB75                       A:E1 X:FF Y:FB P:E5 SP:F9 CYC:24686
FB25  90 4E     BCC $FB75                       A:E1 X:FF Y:FB P:E5 SP:F9 CYC:24688
FB27  10 4C     BPL $FB75                       A:E1 X:FF Y:FB P:E5 SP:F9 CYC:24690
FB29  C9 E1     CMP #$E1                        A:E1 X:FF Y:FB P:E5 SP:F9 CYC:24692
FB2B  D0 48     BNE $FB75                       A:E1 X:FF Y:FB P:67 SP:F9 CYC:24694
FB2D  60        RTS                             A:E1 X:FF Y:FB P:67 SP:F9 CYC:24696
F504  AD 47 06  LDA $0647 = 52                  A:E1 X:FF Y:FB P:67 SP:FB CYC:24702
F507  C9 52     CMP #$52                        A:52 X:FF Y:FB P:65 SP:FB CYC:24706
F509  F0 02     BEQ $F50D                       A:52 X:FF Y:FB P:67 SP:FB CYC:24708
F50D  C8        INY                             A:52 X:FF Y:FB P:67 SP:FB CYC:24711
F50E  A9 29     LDA #$29                        A:52 X:FF Y:FC P:E5 SP:FB CYC:24713
F510  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:FC P:65 SP:FB CYC:24715
F513  20 2E FB  JSR $FB2E                       A:29 X:FF Y:FC P:65 SP:FB CYC:24719
FB2E  B8        CLV                             A:29 X:FF Y:FC P:65 SP:F9 CYC:24725
FB2F  18        CLC                             A:29 X:FF Y:FC P:25 SP:F9 CYC:24727
FB30  A9 42     LDA #$42                        A:29 X:FF Y:FC P:24 SP:F9 CYC:24729
FB32  60        RTS                             A:42 X:FF Y:FC P:24 SP:F9 CYC:24731
F516  5F 48 05 *SRE $0548,X @ 0647 = 29         A:42 X:FF Y:FC P:24 SP:FB CYC:24737
F519  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB CYC:24744
F51A  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB CYC:24746
F51B  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB CYC:24748
F51C  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB CYC:24750
F51D  20 33 FB  JSR $FB33                       A:56 X:FF Y:FC P:25 SP:FB CYC:24752
FB33  70 40     BVS $FB75                       A:56 X:FF Y:FC P:25 SP:F9 CYC:24758
FB35  F0 3E     BEQ $FB75                       A:56 X:FF Y:FC P:25 SP:F9 CYC:24760
FB37  30 3C     BMI $FB75                       A:56 X:FF Y:FC P:25 SP:F9 CYC:24762
FB39  90 3A     BCC $FB75                       A:56 X:FF Y:FC P:25 SP:F9 CYC:24764
FB3B  C9 56     CMP #$56                        A:56 X:FF Y:FC P:25 SP:F9 CYC:24766
FB3D  D0 36     BNE $FB75                       A:56 X:FF Y:FC P:27 SP:F9 CYC:24768
FB3F  60        RTS                             A:56 X:FF Y:FC P:27 SP:F9 CYC:24770
F520  AD 47 06  LDA $0647 = 14                  A:56 X:FF Y:FC P:27 SP:FB CYC:24776
F523  C9 14     CMP #$14                        A:14 X:FF Y:FC P:25 SP:FB CYC:24780
F525  F0 02     BEQ $F529                       A:14 X:FF Y:FC P:27 SP:FB CYC:24782
F529  C8        INY                             A:14 X:FF Y:FC P:27 SP:FB CYC:24785
F52A  A9 37     LDA #$37                        A:14 X:FF Y:FD P:A5 SP:FB CYC:24787
F52C  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:FD P:25 SP:FB CYC:24789
F52F  20 40 FB  JSR $FB40                       A:37 X:FF Y:FD P:25 SP:FB CYC:24793
FB40  24 01     BIT $01 = FF                    A:37 X:FF Y:FD P:25 SP:F9 CYC:24799
FB42  38        SEC                             A:37 X:FF Y:FD P:E5 SP:F9 CYC:24802
FB43  A9 75     LDA #$75                        A:37 X:FF Y:FD P:E5 SP:F9 CYC:24804
FB45  60        RTS                             A:75 X:FF Y:FD P:65 SP:F9 CYC:24806
F532  5F 48 05 *SRE $0548,X @ 0647 = 37         A:75 X:FF Y:FD P:65 SP:FB CYC:24812
F535  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB CYC:24819
F536  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB CYC:24821
F537  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB CYC:24823
F538  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB CYC:24825
F539  20 46 FB  JSR $FB46                       A:6E X:FF Y:FD P:65 SP:FB CYC:24827
FB46  50 2D     BVC $FB75                       A:6E X:FF Y:FD P:65 SP:F9 CYC:24833
FB48  F0 2B     BEQ $FB75                       A:6E X:FF Y:FD P:65 SP:F9 CYC:24835
FB4A  30 29     BMI $FB75                       A:6E X:FF Y:FD P:65 SP:F9 CYC:24837
FB4C  90 27     BCC $FB75                       A:6E X:FF Y:FD P:65 SP:F9 CYC:24839
FB4E  C9 6E     CMP #$6E                        A:6E X:FF Y:FD P:65 SP:F9 CYC:24841
FB50  D0 23     BNE $FB75                       A:6E X:FF Y:FD P:67 SP:F9 CYC:24843
FB52  60        RTS                             A:6E X:FF Y:FD P:67 SP:F9 CYC:24845
F53C  AD 47 06  LDA $0647 = 1B                  A:6E X:FF Y:FD P:67 SP:FB CYC:24851
F53F  C9 1B     CMP #$1B                        A:1B X:FF Y:FD P:65 SP:FB CYC:24855
F541  F0 02     BEQ $F545                       A:1B X:FF Y:FD P:67 SP:FB CYC:24857
F545  60        RTS                             A:1B X:FF Y:FD P:67 SP:FB CYC:24860
C64A  A5 00     LDA $00 = 00                    A:1B X:FF Y:FD P:67 SP:FD CYC:24866
C64C  85 11     STA $11 = 00                    A:0 X:FF Y:FD P:67 SP:FD CYC:24869
C64E  A9 00     LDA #$00                        A:0 X:FF Y:FD P:67 SP:FD CYC:24872
C650  85 00     STA $00 = 00                    A:0 X:FF Y:FD P:67 SP:FD CYC:24874
C652  20 46 F5  JSR $F546                       A:0 X:FF Y:FD P:67 SP:FD CYC:24877
F546  A9 FF     LDA #$FF                        A:0 X:FF Y:FD P:67 SP:FB CYC:24883
F548  85 01     STA $01 = FF                    A:FF X:FF Y:FD P:E5 SP:FB CYC:24885
F54A  A0 01     LDY #$01                        A:FF X:FF Y:FD P:E5 SP:FB CYC:24888
F54C  A2 02     LDX #$02                        A:FF X:FF Y:1 P:65 SP:FB CYC:24890
F54E  A9 47     LDA #$47                        A:FF X:2 Y:1 P:65 SP:FB CYC:24892
F550  85 47     STA $47 = 1B                    A:47 X:2 Y:1 P:65 SP:FB CYC:24894
F552  A9 06     LDA #$06                        A:47 X:2 Y:1 P:65 SP:FB CYC:24897
F554  85 48     STA $48 = 06                    A:6 X:2 Y:1 P:65 SP:FB CYC:24899
F556  A9 A5     LDA #$A5                        A:6 X:2 Y:1 P:65 SP:FB CYC:24902
F558  8D 47 06  STA $0647 = 1B                  A:A5 X:2 Y:1 P:E5 SP:FB CYC:24904
F55B  20 E9 FA  JSR $FAE9                       A:A5 X:2 Y:1 P:E5 SP:FB CYC:24908
FAE9  24 01     BIT $01 = FF                    A:A5 X:2 Y:1 P:E5 SP:F9 CYC:24914
FAEB  18        CLC                             A:A5 X:2 Y:1 P:E5 SP:F9 CYC:24917
FAEC  A9 B2     LDA #$B2                        A:A5 X:2 Y:1 P:E4 SP:F9 CYC:24919
FAEE  60        RTS                             A:B2 X:2 Y:1 P:E4 SP:F9 CYC:24921
F55E  63 45    *RRA ($45,X) @ 47 = 0647 = A5    A:B2 X:2 Y:1 P:E4 SP:FB CYC:24927
F560  EA        NOP                             A:5 X:2 Y:1 P:25 SP:FB CYC:24935
F561  EA        NOP                             A:5 X:2 Y:1 P:25 SP:FB CYC:24937
F562  EA        NOP                             A:5 X:2 Y:1 P:25 SP:FB CYC:24939
F563  EA        NOP                             A:5 X:2 Y:1 P:25 SP:FB CYC:24941
F564  20 EF FA  JSR $FAEF                       A:5 X:2 Y:1 P:25 SP:FB CYC:24943
FAEF  70 2A     BVS $FB1B                       A:5 X:2 Y:1 P:25 SP:F9 CYC:24949
FAF1  90 28     BCC $FB1B                       A:5 X:2 Y:1 P:25 SP:F9 CYC:24951
FAF3  30 26     BMI $FB1B                       A:5 X:2 Y:1 P:25 SP:F9 CYC:24953
FAF5  C9 05     CMP #$05                        A:5 X:2 Y:1 P:25 SP:F9 CYC:24955
FAF7  D0 22     BNE $FB1B                       A:5 X:2 Y:1 P:27 SP:F9 CYC:24957
FAF9  60        RTS                             A:5 X:2 Y:1 P:27 SP:F9 CYC:24959
F567  AD 47 06  LDA $0647 = 52                  A:5 X:2 Y:1 P:27 SP:FB CYC:24965
F56A  C9 52     CMP #$52                        A:52 X:2 Y:1 P:25 SP:FB CYC:24969
F56C  F0 02     BEQ $F570                       A:52 X:2 Y:1 P:27 SP:FB CYC:24971
F570  C8        INY                             A:52 X:2 Y:1 P:27 SP:FB CYC:24974
F571  A9 29     LDA #$29                        A:52 X:2 Y:2 P:25 SP:FB CYC:24976
F573  8D 47 06  STA $0647 = 52                  A:29 X:2 Y:2 P:25 SP:FB CYC:24978
F576  20 FA FA  JSR $FAFA                       A:29 X:2 Y:2 P:25 SP:FB CYC:24982
FAFA  B8        CLV                             A:29 X:2 Y:2 P:25 SP:F9 CYC:24988
FAFB  18        CLC                             A:29 X:2 Y:2 P:25 SP:F9 CYC:24990
FAFC  A9 42     LDA #$42                        A:29 X:2 Y:2 P:24 SP:F9 CYC:24992
FAFE  60        RTS                             A:42 X:2 Y:2 P:24 SP:F9 CYC:24994
F579  63 45    *RRA ($45,X) @ 47 = 0647 = 29    A:42 X:2 Y:2 P:24 SP:FB CYC:25000
F57B  EA        NOP                             A:57 X:2 Y:2 P:24 SP:FB CYC:25008
F57C  EA        NOP                             A:57 X:2 Y:2 P:24 SP:FB CYC:25010
F57D  EA        NOP                             A:57 X:2 Y:2 P:24 SP:FB CYC:25012
F57E  EA        NOP                             A:57 X:2 Y:2 P:24 SP:FB CYC:25014
F57F  20 FF FA  JSR $FAFF                       A:57 X:2 Y:2 P:24 SP:FB CYC:25016
FAFF  70 1A     BVS $FB1B                       A:57 X:2 Y:2 P:24 SP:F9 CYC:25022
FB01  30 18     BMI $FB1B                       A:57 X:2 Y:2 P:24 SP:F9 CYC:25024
FB03  B0 16     BCS $FB1B                       A:57 X:2 Y:2 P:24 SP:F9 CYC:25026
FB05  C9 57     CMP #$57                        A:57 X:2 Y:2 P:24 SP:F9 CYC:25028
FB07  D0 12     BNE $FB1B                       A:57 X:2 Y:2 P:27 SP:F9 CYC:25030
FB09  60        RTS                             A:57 X:2 Y:2 P:27 SP:F9 CYC:25032
F582  AD 47 06  LDA $0647 = 14                  A:57 X:2 Y:2 P:27 SP:FB CYC:25038
F585  C9 14     CMP #$14                        A:14 X:2 Y:2 P:25 SP:FB CYC:25042
F587  F0 02     BEQ $F58B                       A:14 X:2 Y:2 P:27 SP:FB CYC:25044
F58B  C8        INY                             A:14 X:2 Y:2 P:27 SP:FB CYC:25047
F58C  A9 37     LDA #$37                        A:14 X:2 Y:3 P:25 SP:FB CYC:25049
F58E  8D 47 06  STA $0647 = 14                  A:37 X:2 Y:3 P:25 SP:FB CYC:25051
F591  20 0A FB  JSR $FB0A                       A:37 X:2 Y:3 P:25 SP:FB CYC:25055
FB0A  24 01     BIT $01 = FF                    A:37 X:2 Y:3 P:25 SP:F9 CYC:25061
FB0C  38        SEC                             A:37 X:2 Y:3 P:E5 SP:F9 CYC:25064
FB0D  A9 75     LDA #$75                        A:37 X:2 Y:3 P:E5 SP:F9 CYC:25066
FB0F  60        RTS                             A:75 X:2 Y:3 P:65 SP:F9 CYC:25068
F594  63 45    *RRA ($45,X) @ 47 = 0647 = 37    A:75 X:2 Y:3 P:65 SP:FB CYC:25074
F596  EA        NOP                             A:11 X:2 Y:3 P:25 SP:FB CYC:25082
F597  EA        NOP                             A:11 X:2 Y:3 P:25 SP:FB CYC:25084
F598  EA        NOP                             A:11 X:2 Y:3 P:25 SP:FB CYC:25086
F599  EA        NOP                             A:11 X:2 Y:3 P:25 SP:FB CYC:25088
F59A  20 10 FB  JSR $FB10                       A:11 X:2 Y:3 P:25 SP:FB CYC:25090
FB10  70 09     BVS $FB1B                       A:11 X:2 Y:3 P:25 SP:F9 CYC:25096
FB12  30 07     BMI $FB1B                       A:11 X:2 Y:3 P:25 SP:F9 CYC:25098
FB14  90 05     BCC $FB1B                       A:11 X:2 Y:3 P:25 SP:F9 CYC:25100
FB16  C9 11     CMP #$11                        A:11 X:2 Y:3 P:25 SP:F9 CYC:25102
FB18  D0 01     BNE $FB1B                       A:11 X:2 Y:3 P:27 SP:F9 CYC:25104
FB1A  60        RTS                             A:11 X:2 Y:3 P:27 SP:F9 CYC:25106
F59D  AD 47 06  LDA $0647 = 9B                  A:11 X:2 Y:3 P:27 SP:FB CYC:25112
F5A0  C9 9B     CMP #$9B                        A:9B X:2 Y:3 P:A5 SP:FB CYC:25116
F5A2  F0 02     BEQ $F5A6                       A:9B X:2 Y:3 P:27 SP:FB CYC:25118
F5A6  C8        INY                             A:9B X:2 Y:3 P:27 SP:FB CYC:25121
F5A7  A9 A5     LDA #$A5                        A:9B X:2 Y:4 P:25 SP:FB CYC:25123
F5A9  85 47     STA $47 = 47                    A:A5 X:2 Y:4 P:A5 SP:FB CYC:25125
F5AB  20 E9 FA  JSR $FAE9                       A:A5 X:2 Y:4 P:A5 SP:FB CYC:25128
FAE9  24 01     BIT $01 = FF                    A:A5 X:2 Y:4 P:A5 SP:F9 CYC:25134
FAEB  18        CLC                             A:A5 X:2 Y:4 P:E5 SP:F9 CYC:25137
FAEC  A9 B2     LDA #$B2                        A:A5 X:2 Y:4 P:E4 SP:F9 CYC:25139
FAEE  60        RTS                             A:B2 X:2 Y:4 P:E4 SP:F9 CYC:25141
F5AE  67 47    *RRA $47 = A5                    A:B2 X:2 Y:4 P:E4 SP:FB CYC:25147
F5B0  EA        NOP                             A:5 X:2 Y:4 P:25 SP:FB CYC:25152
F5B1  EA        NOP                             A:5 X:2 Y:4 P:25 SP:FB CYC:25154
F5B2  EA        NOP                             A:5 X:2 Y:4 P:25 SP:FB CYC:25156
F5B3  EA        NOP                             A:5 X:2 Y:4 P:25 SP:FB CYC:25158
F5B4  20 EF FA  JSR $FAEF                       A:5 X:2 Y:4 P:25 SP:FB CYC:25160
FAEF  70 2A     BVS $FB1B                       A:5 X:2 Y:4 P:25 SP:F9 CYC:25166
FAF1  90 28     BCC $FB1B                       A:5 X:2 Y:4 P:25 SP:F9 CYC:25168
FAF3  30 26     BMI $FB1B                       A:5 X:2 Y:4 P:25 SP:F9 CYC:25170
FAF5  C9 05     CMP #$05                        A:5 X:2 Y:4 P:25 SP:F9 CYC:25172
FAF7  D0 22     BNE $FB1B                       A:5 X:2 Y:4 P:27 SP:F9 CYC:25174
FAF9  60        RTS                             A:5 X:2 Y:4 P:27 SP:F9 CYC:25176
F5B7  A5 47     LDA $47 = 52                    A:5 X:2 Y:4 P:27 SP:FB CYC:25182
F5B9  C9 52     CMP #$52                        A:52 X:2 Y:4 P:25 SP:FB CYC:25185
F5BB  F0 02     BEQ $F5BF                       A:52 X:2 Y:4 P:27 SP:FB CYC:25187
F5BF  C8        INY                             A:52 X:2 Y:4 P:27 SP:FB CYC:25190
F5C0  A9 29     LDA #$29                        A:52 X:2 Y:5 P:25 SP:FB CYC:25192
F5C2  85 47     STA $47 = 52                    A:29 X:2 Y:5 P:25 SP:FB CYC:25194
F5C4  20 FA FA  JSR $FAFA                       A:29 X:2 Y:5 P:25 SP:FB CYC:25197
FAFA  B8        CLV                             A:29 X:2 Y:5 P:25 SP:F9 CYC:25203
FAFB  18        CLC                             A:29 X:2 Y:5 P:25 SP:F9 CYC:25205
FAFC  A9 42     LDA #$42                        A:29 X:2 Y:5 P:24 SP:F9 CYC:25207
FAFE  60        RTS                             A:42 X:2 Y:5 P:24 SP:F9 CYC:25209
F5C7  67 47    *RRA $47 = 29                    A:42 X:2 Y:5 P:24 SP:FB CYC:25215
F5C9  EA        NOP                             A:57 X:2 Y:5 P:24 SP:FB CYC:25220
F5CA  EA        NOP                             A:57 X:2 Y:5 P:24 SP:FB CYC:25222
F5CB  EA        NOP                             A:57 X:2 Y:5 P:24 SP:FB CYC:25224
F5CC  EA        NOP                             A:57 X:2 Y:5 P:24 SP:FB CYC:25226
F5CD  20 FF FA  JSR $FAFF                       A:57 X:2 Y:5 P:24 SP:FB CYC:25228
FAFF  70 1A     BVS $FB1B                       A:57 X:2 Y:5 P:24 SP:F9 CYC:25234
FB01  30 18     BMI $FB1B                       A:57 X:2 Y:5 P:24 SP:F9 CYC:25236
FB03  B0 16     BCS $FB1B                       A:57 X:2 Y:5 P:24 SP:F9 CYC:25238
FB05  C9 57     CMP #$57                        A:57 X:2 Y:5 P:24 SP:F9 CYC:25240
FB07  D0 12     BNE $FB1B                       A:57 X:2 Y:5 P:27 SP:F9 CYC:25242
FB09  60        RTS                             A:57 X:2 Y:5 P:27 SP:F9 CYC:25244
F5D0  A5 47     LDA $47 = 14                    A:57 X:2 Y:5 P:27 SP:FB CYC:25250
F5D2  C9 14     CMP #$14                        A:14 X:2 Y:5 P:25 SP:FB CYC:25253
F5D4  F0 02     BEQ $F5D8                       A:14 X:2 Y:5 P:27 SP:FB CYC:25255
F5D8  C8        INY                             A:14 X:2 Y:5 P:27 SP:FB CYC:25258
F5D9  A9 37     LDA #$37                        A:14 X:2 Y:6 P:25 SP:FB CYC:25260
F5DB  85 47     STA $47 = 14                    A:37 X:2 Y:6 P:25 SP:FB CYC:25262
F5DD  20 0A FB  JSR $FB0A                       A:37 X:2 Y:6 P:25 SP:FB CYC:25265
FB0A  24 01     BIT $01 = FF                    A:37 X:2 Y:6 P:25 SP:F9 CYC:25271
FB0C  38        SEC                             A:37 X:2 Y:6 P:E5 SP:F9 CYC:25274
FB0D  A9 75     LDA #$75                        A:37 X:2 Y:6 P:E5 SP:F9 CYC:25276
FB0F  60        RTS                             A:75 X:2 Y:6 P:65 SP:F9 CYC:25278
F5E0  67 47    *RRA $47 = 37                    A:75 X:2 Y:6 P:65 SP:FB CYC:25284
F5E2  EA        NOP                             A:11 X:2 Y:6 P:25 SP:FB CYC:25289
F5E3  EA        NOP                             A:11 X:2 Y:6 P:25 SP:FB CYC:25291
F5E4  EA        NOP                             A:11 X:2 Y:6 P:25 SP:FB CYC:25293
F5E5  EA        NOP                             A:11 X:2 Y:6 P:25 SP:FB CYC:25295
F5E6  20 10 FB  JSR $FB10                       A:11 X:2 Y:6 P:25 SP:FB CYC:25297
FB10  70 09     BVS $FB1B                       A:11 X:2 Y:6 P:25 SP:F9 CYC:25303
FB12  30 07     BMI $FB1B                       A:11 X:2 Y:6 P:25 SP:F9 CYC:25305
FB14  90 05     BCC $FB1B                       A:11 X:2 Y:6 P:25 SP:F9 CYC:25307
FB16  C9 11     CMP #$11                        A:11 X:2 Y:6 P:25 SP:F9 CYC:25309
FB18  D0 01     BNE $FB1B                       A:11 X:2 Y:6 P:27 SP:F9 CYC:25311
FB1A  60        RTS                             A:11 X:2 Y:6 P:27 SP:F9 CYC:25313
F5E9  A5 47     LDA $47 = 9B                    A:11 X:2 Y:6 P:27 SP:FB CYC:25319
F5EB  C9 9B     CMP #$9B                        A:9B X:2 Y:6 P:A5 SP:FB CYC:25322
F5ED  F0 02     BEQ $F5F1                       A:9B X:2 Y:6 P:27 SP:FB CYC:25324
F5F1  C8        INY                             A:9B X:2 Y:6 P:27 SP:FB CYC:25327
F5F2  A9 A5     LDA #$A5                        A:9B X:2 Y:7 P:25 SP:FB CYC:25329
F5F4  8D 47 06  STA $0647 = 9B                  A:A5 X:2 Y:7 P:A5 SP:FB CYC:25331
F5F7  20 E9 FA  JSR $FAE9                       A:A5 X:2 Y:7 P:A5 SP:FB CYC:25335
FAE9  24 01     BIT $01 = FF                    A:A5 X:2 Y:7 P:A5 SP:F9 CYC:25341
FAEB  18        CLC                             A:A5 X:2 Y:7 P:E5 SP:F9 CYC:25344
FAEC  A9 B2     LDA #$B2                        A:A5 X:2 Y:7 P:E4 SP:F9 CYC:25346
FAEE  60        RTS                             A:B2 X:2 Y:7 P:E4 SP:F9 CYC:25348
F5FA  6F 47 06 *RRA $0647 = A5                  A:B2 X:2 Y:7 P:E4 SP:FB CYC:25354
F5FD  EA        NOP                             A:5 X:2 Y:7 P:25 SP:FB CYC:25360
F5FE  EA        NOP                             A:5 X:2 Y:7 P:25 SP:FB CYC:25362
F5FF  EA        NOP                             A:5 X:2 Y:7 P:25 SP:FB CYC:25364
F600  EA        NOP                             A:5 X:2 Y:7 P:25 SP:FB CYC:25366
F601  20 EF FA  JSR $FAEF                       A:5 X:2 Y:7 P:25 SP:FB CYC:25368
FAEF  70 2A     BVS $FB1B                       A:5 X:2 Y:7 P:25 SP:F9 CYC:25374
FAF1  90 28     BCC $FB1B                       A:5 X:2 Y:7 P:25 SP:F9 CYC:25376
FAF3  30 26     BMI $FB1B                       A:5 X:2 Y:7 P:25 SP:F9 CYC:25378
FAF5  C9 05     CMP #$05                        A:5 X:2 Y:7 P:25 SP:F9 CYC:25380
FAF7  D0 22     BNE $FB1B                       A:5 X:2 Y:7 P:27 SP:F9 CYC:25382
FAF9  60        RTS                             A:5 X:2 Y:7 P:27 SP:F9 CYC:25384
F604  AD 47 06  LDA $0647 = 52                  A:5 X:2 Y:7 P:27 SP:FB CYC:25390
F607  C9 52     CMP #$52                        A:52 X:2 Y:7 P:25 SP:FB CYC:25394
F609  F0 02     BEQ $F60D                       A:52 X:2 Y:7 P:27 SP:FB CYC:25396
F60D  C8        INY                             A:52 X:2 Y:7 P:27 SP:FB CYC:25399
F60E  A9 29     LDA #$29                        A:52 X:2 Y:8 P:25 SP:FB CYC:25401
F610  8D 47 06  STA $0647 = 52                  A:29 X:2 Y:8 P:25 SP:FB CYC:25403
F613  20 FA FA  JSR $FAFA                       A:29 X:2 Y:8 P:25 SP:FB CYC:25407
FAFA  B8        CLV                             A:29 X:2 Y:8 P:25 SP:F9 CYC:25413
FAFB  18        CLC                             A:29 X:2 Y:8 P:25 SP:F9 CYC:25415
FAFC  A9 42     LDA #$42                        A:29 X:2 Y:8 P:24 SP:F9 CYC:25417
FAFE  60        RTS                             A:42 X:2 Y:8 P:24 SP:F9 CYC:25419
F616  6F 47 06 *RRA $0647 = 29                  A:42 X:2 Y:8 P:24 SP:FB CYC:25425
F619  EA        NOP                             A:57 X:2 Y:8 P:24 SP:FB CYC:25431
F61A  EA        NOP                             A:57 X:2 Y:8 P:24 SP:FB CYC:25433
F61B  EA        NOP                             A:57 X:2 Y:8 P:24 SP:FB CYC:25435
F61C  EA        NOP                             A:57 X:2 Y:8 P:24 SP:FB CYC:25437
F61D  20 FF FA  JSR $FAFF                       A:57 X:2 Y:8 P:24 SP:FB CYC:25439
FAFF  70 1A     BVS $FB1B                       A:57 X:2 Y:8 P:24 SP:F9 CYC:25445
FB01  30 18     BMI $FB1B                       A:57 X:2 Y:8 P:24 SP:F9 CYC:25447
FB03  B0 16     BCS $FB1B                       A:57 X:2 Y:8 P:24 SP:F9 CYC:25449
FB05  C9 57     CMP #$57                        A:57 X:2 Y:8 P:24 SP:F9 CYC:25451
FB07  D0 12     BNE $FB1B                       A:57 X:2 Y:8 P:27 SP:F9 CYC:25453
FB09  60        RTS                             A:57 X:2 Y:8 P:27 SP:F9 CYC:25455
F620  AD 47 06  LDA $0647 = 14                  A:57 X:2 Y:8 P:27 SP:FB CYC:25461
F623  C9 14     CMP #$14                        A:14 X:2 Y:8 P:25 SP:FB CYC:25465
F625  F0 02     BEQ $F629                       A:14 X:2 Y:8 P:27 SP:FB CYC:25467
F629  C8        INY                             A:14 X:2 Y:8 P:27 SP:FB CYC:25470
F62A  A9 37     LDA #$37                        A:14 X:2 Y:9 P:25 SP:FB CYC:25472
F62C  8D 47 06  STA $0647 = 14                  A:37 X:2 Y:9 P:25 SP:FB CYC:25474
F62F  20 0A FB  JSR $FB0A                       A:37 X:2 Y:9 P:25 SP:FB CYC:25478
FB0A  24 01     BIT $01 = FF                    A:37 X:2 Y:9 P:25 SP:F9 CYC:25484
FB0C  38        SEC                             A:37 X:2 Y:9 P:E5 SP:F9 CYC:25487
FB0D  A9 75     LDA #$75                        A:37 X:2 Y:9 P:E5 SP:F9 CYC:25489
FB0F  60        RTS                             A:75 X:2 Y:9 P:65 SP:F9 CYC:25491
F632  6F 47 06 *RRA $0647 = 37                  A:75 X:2 Y:9 P:65 SP:FB CYC:25497
F635  EA        NOP                             A:11 X:2 Y:9 P:25 SP:FB CYC:25503
F636  EA        NOP                             A:11 X:2 Y:9 P:25 SP:FB CYC:25505
F637  EA        NOP                             A:11 X:2 Y:9 P:25 SP:FB CYC:25507
F638  EA        NOP                             A:11 X:2 Y:9 P:25 SP:FB CYC:25509
F639  20 10 FB  JSR $FB10                       A:11 X:2 Y:9 P:25 SP:FB CYC:25511
FB10  70 09     BVS $FB1B                       A:11 X:2 Y:9 P:25 SP:F9 CYC:25517
FB12  30 07     BMI $FB1B                       A:11 X:2 Y:9 P:25 SP:F9 CYC:25519
FB14  90 05     BCC $FB1B                       A:11 X:2 Y:9 P:25 SP:F9 CYC:25521
FB16  C9 11     CMP #$11                        A:11 X:2 Y:9 P:25 SP:F9 CYC:25523
FB18  D0 01     BNE $FB1B                       A:11 X:2 Y:9 P:27 SP:F9 CYC:25525
FB1A  60        RTS                             A:11 X:2 Y:9 P:27 SP:F9 CYC:25527
F63C  AD 47 06  LDA $0647 = 9B                  A:11 X:2 Y:9 P:27 SP:FB CYC:25533
F63F  C9 9B     CMP #$9B                        A:9B X:2 Y:9 P:A5 SP:FB CYC:25537
F641  F0 02     BEQ $F645                       A:9B X:2 Y:9 P:27 SP:FB CYC:25539
F645  A9 A5     LDA #$A5                        A:9B X:2 Y:9 P:27 SP:FB CYC:25542
F647  8D 47 06  STA $0647 = 9B                  A:A5 X:2 Y:9 P:A5 SP:FB CYC:25544
F64A  A9 48     LDA #$48                        A:A5 X:2 Y:9 P:A5 SP:FB CYC:25548
F64C  85 45     STA $45 = 48                    A:48 X:2 Y:9 P:25 SP:FB CYC:25550
F64E  A9 05     LDA #$05                        A:48 X:2 Y:9 P:25 SP:FB CYC:25553
F650  85 46     STA $46 = 05                    A:5 X:2 Y:9 P:25 SP:FB CYC:25555
F652  A0 FF     LDY #$FF                        A:5 X:2 Y:9 P:25 SP:FB CYC:25558
F654  20 E9 FA  JSR $FAE9                       A:5 X:2 Y:FF P:A5 SP:FB CYC:25560
FAE9  24 01     BIT $01 = FF                    A:5 X:2 Y:FF P:A5 SP:F9 CYC:25566
FAEB  18        CLC                             A:5 X:2 Y:FF P:E5 SP:F9 CYC:25569
FAEC  A9 B2     LDA #$B2                        A:5 X:2 Y:FF P:E4 SP:F9 CYC:25571
FAEE  60        RTS                             A:B2 X:2 Y:FF P:E4 SP:F9 CYC:25573
F657  73 45    *RRA ($45),Y = 0548 @ 0647 = A5  A:B2 X:2 Y:FF P:E4 SP:FB CYC:25579
F659  EA        NOP                             A:5 X:2 Y:FF P:25 SP:FB CYC:25587
F65A  EA        NOP                             A:5 X:2 Y:FF P:25 SP:FB CYC:25589
F65B  08        PHP                             A:5 X:2 Y:FF P:25 SP:FB CYC:25591
F65C  48        PHA                             A:5 X:2 Y:FF P:25 SP:FA CYC:25594
F65D  A0 0A     LDY #$0A                        A:5 X:2 Y:FF P:25 SP:F9 CYC:25597
F65F  68        PLA                             A:5 X:2 Y:A P:25 SP:F9 CYC:25599
F660  28        PLP                             A:5 X:2 Y:A P:25 SP:FA CYC:25603
F661  20 EF FA  JSR $FAEF                       A:5 X:2 Y:A P:25 SP:FB CYC:25607
FAEF  70 2A     BVS $FB1B                       A:5 X:2 Y:A P:25 SP:F9 CYC:25613
FAF1  90 28     BCC $FB1B                       A:5 X:2 Y:A P:25 SP:F9 CYC:25615
FAF3  30 26     BMI $FB1B                       A:5 X:2 Y:A P:25 SP:F9 CYC:25617
FAF5  C9 05     CMP #$05                        A:5 X:2 Y:A P:25 SP:F9 CYC:25619
FAF7  D0 22     BNE $FB1B                       A:5 X:2 Y:A P:27 SP:F9 CYC:25621
FAF9  60        RTS                             A:5 X:2 Y:A P:27 SP:F9 CYC:25623
F664  AD 47 06  LDA $0647 = 52                  A:5 X:2 Y:A P:27 SP:FB CYC:25629
F667  C9 52     CMP #$52                        A:52 X:2 Y:A P:25 SP:FB CYC:25633
F669  F0 02     BEQ $F66D                       A:52 X:2 Y:A P:27 SP:FB CYC:25635
F66D  A0 FF     LDY #$FF                        A:52 X:2 Y:A P:27 SP:FB CYC:25638
F66F  A9 29     LDA #$29                        A:52 X:2 Y:FF P:A5 SP:FB CYC:25640
F671  8D 47 06  STA $0647 = 52                  A:29 X:2 Y:FF P:25 SP:FB CYC:25642
F674  20 FA FA  JSR $FAFA                       A:29 X:2 Y:FF P:25 SP:FB CYC:25646
FAFA  B8        CLV                             A:29 X:2 Y:FF P:25 SP:F9 CYC:25652
FAFB  18        CLC                             A:29 X:2 Y:FF P:25 SP:F9 CYC:25654
FAFC  A9 42     LDA #$42                        A:29 X:2 Y:FF P:24 SP:F9 CYC:25656
FAFE  60        RTS                             A:42 X:2 Y:FF P:24 SP:F9 CYC:25658
F677  73 45    *RRA ($45),Y = 0548 @ 0647 = 29  A:42 X:2 Y:FF P:24 SP:FB CYC:25664
F679  EA        NOP                             A:57 X:2 Y:FF P:24 SP:FB CYC:25672
F67A  EA        NOP                             A:57 X:2 Y:FF P:24 SP:FB CYC:25674
F67B  08        PHP                             A:57 X:2 Y:FF P:24 SP:FB CYC:25676
F67C  48        PHA                             A:57 X:2 Y:FF P:24 SP:FA CYC:25679
F67D  A0 0B     LDY #$0B                        A:57 X:2 Y:FF P:24 SP:F9 CYC:25682
F67F  68        PLA                             A:57 X:2 Y:B P:24 SP:F9 CYC:25684
F680  28        PLP                             A:57 X:2 Y:B P:24 SP:FA CYC:25688
F681  20 FF FA  JSR $FAFF                       A:57 X:2 Y:B P:24 SP:FB CYC:25692
FAFF  70 1A     BVS $FB1B                       A:57 X:2 Y:B P:24 SP:F9 CYC:25698
FB01  30 18     BMI $FB1B                       A:57 X:2 Y:B P:24 SP:F9 CYC:25700
FB03  B0 16     BCS $FB1B                       A:57 X:2 Y:B P:24 SP:F9 CYC:25702
FB05  C9 57     CMP #$57                        A:57 X:2 Y:B P:24 SP:F9 CYC:25704
FB07  D0 12     BNE $FB1B                       A:57 X:2 Y:B P:27 SP:F9 CYC:25706
FB09  60        RTS                             A:57 X:2 Y:B P:27 SP:F9 CYC:25708
F684  AD 47 06  LDA $0647 = 14                  A:57 X:2 Y:B P:27 SP:FB CYC:25714
F687  C9 14     CMP #$14                        A:14 X:2 Y:B P:25 SP:FB CYC:25718
F689  F0 02     BEQ $F68D                       A:14 X:2 Y:B P:27 SP:FB CYC:25720
F68D  A0 FF     LDY #$FF                        A:14 X:2 Y:B P:27 SP:FB CYC:25723
F68F  A9 37     LDA #$37                        A:14 X:2 Y:FF P:A5 SP:FB CYC:25725
F691  8D 47 06  STA $0647 = 14                  A:37 X:2 Y:FF P:25 SP:FB CYC:25727
F694  20 0A FB  JSR $FB0A                       A:37 X:2 Y:FF P:25 SP:FB CYC:25731
FB0A  24 01     BIT $01 = FF                    A:37 X:2 Y:FF P:25 SP:F9 CYC:25737
FB0C  38        SEC                             A:37 X:2 Y:FF P:E5 SP:F9 CYC:25740
FB0D  A9 75     LDA #$75                        A:37 X:2 Y:FF P:E5 SP:F9 CYC:25742
FB0F  60        RTS                             A:75 X:2 Y:FF P:65 SP:F9 CYC:25744
F697  73 45    *RRA ($45),Y = 0548 @ 0647 = 37  A:75 X:2 Y:FF P:65 SP:FB CYC:25750
F699  EA        NOP                             A:11 X:2 Y:FF P:25 SP:FB CYC:25758
F69A  EA        NOP                             A:11 X:2 Y:FF P:25 SP:FB CYC:25760
F69B  08        PHP                             A:11 X:2 Y:FF P:25 SP:FB CYC:25762
F69C  48        PHA                             A:11 X:2 Y:FF P:25 SP:FA CYC:25765
F69D  A0 0C     LDY #$0C                        A:11 X:2 Y:FF P:25 SP:F9 CYC:25768
F69F  68        PLA                             A:11 X:2 Y:C P:25 SP:F9 CYC:25770
F6A0  28        PLP                             A:11 X:2 Y:C P:25 SP:FA CYC:25774
F6A1  20 10 FB  JSR $FB10                       A:11 X:2 Y:C P:25 SP:FB CYC:25778
FB10  70 09     BVS $FB1B                       A:11 X:2 Y:C P:25 SP:F9 CYC:25784
FB12  30 07     BMI $FB1B                       A:11 X:2 Y:C P:25 SP:F9 CYC:25786
FB14  90 05     BCC $FB1B                       A:11 X:2 Y:C P:25 SP:F9 CYC:25788
FB16  C9 11     CMP #$11                        A:11 X:2 Y:C P:25 SP:F9 CYC:25790
FB18  D0 01     BNE $FB1B                       A:11 X:2 Y:C P:27 SP:F9 CYC:25792
FB1A  60        RTS                             A:11 X:2 Y:C P:27 SP:F9 CYC:25794
F6A4  AD 47 06  LDA $0647 = 9B                  A:11 X:2 Y:C P:27 SP:FB CYC:25800
F6A7  C9 9B     CMP #$9B                        A:9B X:2 Y:C P:A5 SP:FB CYC:25804
F6A9  F0 02     BEQ $F6AD                       A:9B X:2 Y:C P:27 SP:FB CYC:25806
F6AD  A0 0D     LDY #$0D                        A:9B X:2 Y:C P:27 SP:FB CYC:25809
F6AF  A2 FF     LDX #$FF                        A:9B X:2 Y:D P:25 SP:FB CYC:25811
F6B1  A9 A5     LDA #$A5                        A:9B X:FF Y:D P:A5 SP:FB CYC:25813
F6B3  85 47     STA $47 = 9B                    A:A5 X:FF Y:D P:A5 SP:FB CYC:25815
F6B5  20 E9 FA  JSR $FAE9                       A:A5 X:FF Y:D P:A5 SP:FB CYC:25818
FAE9  24 01     BIT $01 = FF                    A:A5 X:FF Y:D P:A5 SP:F9 CYC:25824
FAEB  18        CLC                             A:A5 X:FF Y:D P:E5 SP:F9 CYC:25827
FAEC  A9 B2     LDA #$B2                        A:A5 X:FF Y:D P:E4 SP:F9 CYC:25829
FAEE  60        RTS                             A:B2 X:FF Y:D P:E4 SP:F9 CYC:25831
F6B8  77 48    *RRA $48,X @ 47 = A5             A:B2 X:FF Y:D P:E4 SP:FB CYC:25837
F6BA  EA        NOP                             A:5 X:FF Y:D P:25 SP:FB CYC:25843
F6BB  EA        NOP                             A:5 X:FF Y:D P:25 SP:FB CYC:25845
F6BC  EA        NOP                             A:5 X:FF Y:D P:25 SP:FB CYC:25847
F6BD  EA        NOP                             A:5 X:FF Y:D P:25 SP:FB CYC:25849
F6BE  20 EF FA  JSR $FAEF                       A:5 X:FF Y:D P:25 SP:FB CYC:25851
FAEF  70 2A     BVS $FB1B                       A:5 X:FF Y:D P:25 SP:F9 CYC:25857
FAF1  90 28     BCC $FB1B                       A:5 X:FF Y:D P:25 SP:F9 CYC:25859
FAF3  30 26     BMI $FB1B                       A:5 X:FF Y:D P:25 SP:F9 CYC:25861
FAF5  C9 05     CMP #$05                        A:5 X:FF Y:D P:25 SP:F9 CYC:25863
FAF7  D0 22     BNE $FB1B                       A:5 X:FF Y:D P:27 SP:F9 CYC:25865
FAF9  60        RTS                             A:5 X:FF Y:D P:27 SP:F9 CYC:25867
F6C1  A5 47     LDA $47 = 52                    A:5 X:FF Y:D P:27 SP:FB CYC:25873
F6C3  C9 52     CMP #$52                        A:52 X:FF Y:D P:25 SP:FB CYC:25876
F6C5  F0 02     BEQ $F6C9                       A:52 X:FF Y:D P:27 SP:FB CYC:25878
F6C9  C8        INY                             A:52 X:FF Y:D P:27 SP:FB CYC:25881
F6CA  A9 29     LDA #$29                        A:52 X:FF Y:E P:25 SP:FB CYC:25883
F6CC  85 47     STA $47 = 52                    A:29 X:FF Y:E P:25 SP:FB CYC:25885
F6CE  20 FA FA  JSR $FAFA                       A:29 X:FF Y:E P:25 SP:FB CYC:25888
FAFA  B8        CLV                             A:29 X:FF Y:E P:25 SP:F9 CYC:25894
FAFB  18        CLC                             A:29 X:FF Y:E P:25 SP:F9 CYC:25896
FAFC  A9 42     LDA #$42                        A:29 X:FF Y:E P:24 SP:F9 CYC:25898
FAFE  60        RTS                             A:42 X:FF Y:E P:24 SP:F9 CYC:25900
F6D1  77 48    *RRA $48,X @ 47 = 29             A:42 X:FF Y:E P:24 SP:FB CYC:25906
F6D3  EA        NOP                             A:57 X:FF Y:E P:24 SP:FB CYC:25912
F6D4  EA        NOP                             A:57 X:FF Y:E P:24 SP:FB CYC:25914
F6D5  EA        NOP                             A:57 X:FF Y:E P:24 SP:FB CYC:25916
F6D6  EA        NOP                             A:57 X:FF Y:E P:24 SP:FB CYC:25918
F6D7  20 FF FA  JSR $FAFF                       A:57 X:FF Y:E P:24 SP:FB CYC:25920
FAFF  70 1A     BVS $FB1B                       A:57 X:FF Y:E P:24 SP:F9 CYC:25926
FB01  30 18     BMI $FB1B                       A:57 X:FF Y:E P:24 SP:F9 CYC:25928
FB03  B0 16     BCS $FB1B                       A:57 X:FF Y:E P:24 SP:F9 CYC:25930
FB05  C9 57     CMP #$57                        A:57 X:FF Y:E P:24 SP:F9 CYC:25932
FB07  D0 12     BNE $FB1B                       A:57 X:FF Y:E P:27 SP:F9 CYC:25934
FB09  60        RTS                             A:57 X:FF Y:E P:27 SP:F9 CYC:25936
F6DA  A5 47     LDA $47 = 14                    A:57 X:FF Y:E P:27 SP:FB CYC:25942
F6DC  C9 14     CMP #$14                        A:14 X:FF Y:E P:25 SP:FB CYC:25945
F6DE  F0 02     BEQ $F6E2                       A:14 X:FF Y:E P:27 SP:FB CYC:25947
F6E2  C8        INY                             A:14 X:FF Y:E P:27 SP:FB CYC:25950
F6E3  A9 37     LDA #$37                        A:14 X:FF Y:F P:25 SP:FB CYC:25952
F6E5  85 47     STA $47 = 14                    A:37 X:FF Y:F P:25 SP:FB CYC:25954
F6E7  20 0A FB  JSR $FB0A                       A:37 X:FF Y:F P:25 SP:FB CYC:25957
FB0A  24 01     BIT $01 = FF                    A:37 X:FF Y:F P:25 SP:F9 CYC:25963
FB0C  38        SEC                             A:37 X:FF Y:F P:E5 SP:F9 CYC:25966
FB0D  A9 75     LDA #$75                        A:37 X:FF Y:F P:E5 SP:F9 CYC:25968
FB0F  60        RTS                             A:75 X:FF Y:F P:65 SP:F9 CYC:25970
F6EA  77 48    *RRA $48,X @ 47 = 37             A:75 X:FF Y:F P:65 SP:FB CYC:25976
F6EC  EA        NOP                             A:11 X:FF Y:F P:25 SP:FB CYC:25982
F6ED  EA        NOP                             A:11 X:FF Y:F P:25 SP:FB CYC:25984
F6EE  EA        NOP                             A:11 X:FF Y:F P:25 SP:FB CYC:25986
F6EF  EA        NOP                             A:11 X:FF Y:F P:25 SP:FB CYC:25988
F6F0  20 10 FB  JSR $FB10                       A:11 X:FF Y:F P:25 SP:FB CYC:25990
FB10  70 09     BVS $FB1B                       A:11 X:FF Y:F P:25 SP:F9 CYC:25996
FB12  30 07     BMI $FB1B                       A:11 X:FF Y:F P:25 SP:F9 CYC:25998
FB14  90 05     BCC $FB1B                       A:11 X:FF Y:F P:25 SP:F9 CYC:26000
FB16  C9 11     CMP #$11                        A:11 X:FF Y:F P:25 SP:F9 CYC:26002
FB18  D0 01     BNE $FB1B                       A:11 X:FF Y:F P:27 SP:F9 CYC:26004
FB1A  60        RTS                             A:11 X:FF Y:F P:27 SP:F9 CYC:26006
F6F3  A5 47     LDA $47 = 9B                    A:11 X:FF Y:F P:27 SP:FB CYC:26012
F6F5  C9 9B     CMP #$9B                        A:9B X:FF Y:F P:A5 SP:FB CYC:26015
F6F7  F0 02     BEQ $F6FB                       A:9B X:FF Y:F P:27 SP:FB CYC:26017
F6FB  A9 A5     LDA #$A5                        A:9B X:FF Y:F P:27 SP:FB CYC:26020
F6FD  8D 47 06  STA $0647 = 9B                  A:A5 X:FF Y:F P:A5 SP:FB CYC:26022
F700  A0 FF     LDY #$FF                        A:A5 X:FF Y:F P:A5 SP:FB CYC:26026
F702  20 E9 FA  JSR $FAE9                       A:A5 X:FF Y:FF P:A5 SP:FB CYC:26028
FAE9  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:A5 SP:F9 CYC:26034
FAEB  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9 CYC:26037
FAEC  A9 B2     LDA #$B2                        A:A5 X:FF Y:FF P:E4 SP:F9 CYC:26039
FAEE  60        RTS                             A:B2 X:FF Y:FF P:E4 SP:F9 CYC:26041
F705  7B 48 05 *RRA $0548,Y @ 0647 = A5         A:B2 X:FF Y:FF P:E4 SP:FB CYC:26047
F708  EA        NOP                             A:5 X:FF Y:FF P:25 SP:FB CYC:26054
F709  EA        NOP                             A:5 X:FF Y:FF P:25 SP:FB CYC:26056
F70A  08        PHP                             A:5 X:FF Y:FF P:25 SP:FB CYC:26058
F70B  48        PHA                             A:5 X:FF Y:FF P:25 SP:FA CYC:26061
F70C  A0 10     LDY #$10                        A:5 X:FF Y:FF P:25 SP:F9 CYC:26064
F70E  68        PLA                             A:5 X:FF Y:10 P:25 SP:F9 CYC:26066
F70F  28        PLP                             A:5 X:FF Y:10 P:25 SP:FA CYC:26070
F710  20 EF FA  JSR $FAEF                       A:5 X:FF Y:10 P:25 SP:FB CYC:26074
FAEF  70 2A     BVS $FB1B                       A:5 X:FF Y:10 P:25 SP:F9 CYC:26080
FAF1  90 28     BCC $FB1B                       A:5 X:FF Y:10 P:25 SP:F9 CYC:26082
FAF3  30 26     BMI $FB1B                       A:5 X:FF Y:10 P:25 SP:F9 CYC:26084
FAF5  C9 05     CMP #$05                        A:5 X:FF Y:10 P:25 SP:F9 CYC:26086
FAF7  D0 22     BNE $FB1B                       A:5 X:FF Y:10 P:27 SP:F9 CYC:26088
FAF9  60        RTS                             A:5 X:FF Y:10 P:27 SP:F9 CYC:26090
F713  AD 47 06  LDA $0647 = 52                  A:5 X:FF Y:10 P:27 SP:FB CYC:26096
F716  C9 52     CMP #$52                        A:52 X:FF Y:10 P:25 SP:FB CYC:26100
F718  F0 02     BEQ $F71C                       A:52 X:FF Y:10 P:27 SP:FB CYC:26102
F71C  A0 FF     LDY #$FF                        A:52 X:FF Y:10 P:27 SP:FB CYC:26105
F71E  A9 29     LDA #$29                        A:52 X:FF Y:FF P:A5 SP:FB CYC:26107
F720  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:FF P:25 SP:FB CYC:26109
F723  20 FA FA  JSR $FAFA                       A:29 X:FF Y:FF P:25 SP:FB CYC:26113
FAFA  B8        CLV                             A:29 X:FF Y:FF P:25 SP:F9 CYC:26119
FAFB  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9 CYC:26121
FAFC  A9 42     LDA #$42                        A:29 X:FF Y:FF P:24 SP:F9 CYC:26123
FAFE  60        RTS                             A:42 X:FF Y:FF P:24 SP:F9 CYC:26125
F726  7B 48 05 *RRA $0548,Y @ 0647 = 29         A:42 X:FF Y:FF P:24 SP:FB CYC:26131
F729  EA        NOP                             A:57 X:FF Y:FF P:24 SP:FB CYC:26138
F72A  EA        NOP                             A:57 X:FF Y:FF P:24 SP:FB CYC:26140
F72B  08        PHP                             A:57 X:FF Y:FF P:24 SP:FB CYC:26142
F72C  48        PHA                             A:57 X:FF Y:FF P:24 SP:FA CYC:26145
F72D  A0 11     LDY #$11                        A:57 X:FF Y:FF P:24 SP:F9 CYC:26148
F72F  68        PLA                             A:57 X:FF Y:11 P:24 SP:F9 CYC:26150
F730  28        PLP                             A:57 X:FF Y:11 P:24 SP:FA CYC:26154
F731  20 FF FA  JSR $FAFF                       A:57 X:FF Y:11 P:24 SP:FB CYC:26158
FAFF  70 1A     BVS $FB1B                       A:57 X:FF Y:11 P:24 SP:F9 CYC:26164
FB01  30 18     BMI $FB1B                       A:57 X:FF Y:11 P:24 SP:F9 CYC:26166
FB03  B0 16     BCS $FB1B                       A:57 X:FF Y:11 P:24 SP:F9 CYC:26168
FB05  C9 57     CMP #$57                        A:57 X:FF Y:11 P:24 SP:F9 CYC:26170
FB07  D0 12     BNE $FB1B                       A:57 X:FF Y:11 P:27 SP:F9 CYC:26172
FB09  60        RTS                             A:57 X:FF Y:11 P:27 SP:F9 CYC:26174
F734  AD 47 06  LDA $0647 = 14                  A:57 X:FF Y:11 P:27 SP:FB CYC:26180
F737  C9 14     CMP #$14                        A:14 X:FF Y:11 P:25 SP:FB CYC:26184
F739  F0 02     BEQ $F73D                       A:14 X:FF Y:11 P:27 SP:FB CYC:26186
F73D  A0 FF     LDY #$FF                        A:14 X:FF Y:11 P:27 SP:FB CYC:26189
F73F  A9 37     LDA #$37                        A:14 X:FF Y:FF P:A5 SP:FB CYC:26191
F741  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:FF P:25 SP:FB CYC:26193
F744  20 0A FB  JSR $FB0A                       A:37 X:FF Y:FF P:25 SP:FB CYC:26197
FB0A  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9 CYC:26203
FB0C  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9 CYC:26206
FB0D  A9 75     LDA #$75                        A:37 X:FF Y:FF P:E5 SP:F9 CYC:26208
FB0F  60        RTS                             A:75 X:FF Y:FF P:65 SP:F9 CYC:26210
F747  7B 48 05 *RRA $0548,Y @ 0647 = 37         A:75 X:FF Y:FF P:65 SP:FB CYC:26216
F74A  EA        NOP                             A:11 X:FF Y:FF P:25 SP:FB CYC:26223
F74B  EA        NOP                             A:11 X:FF Y:FF P:25 SP:FB CYC:26225
F74C  08        PHP                             A:11 X:FF Y:FF P:25 SP:FB CYC:26227
F74D  48        PHA                             A:11 X:FF Y:FF P:25 SP:FA CYC:26230
F74E  A0 12     LDY #$12                        A:11 X:FF Y:FF P:25 SP:F9 CYC:26233
F750  68        PLA                             A:11 X:FF Y:12 P:25 SP:F9 CYC:26235
F751  28        PLP                             A:11 X:FF Y:12 P:25 SP:FA CYC:26239
F752  20 10 FB  JSR $FB10                       A:11 X:FF Y:12 P:25 SP:FB CYC:26243
FB10  70 09     BVS $FB1B                       A:11 X:FF Y:12 P:25 SP:F9 CYC:26249
FB12  30 07     BMI $FB1B                       A:11 X:FF Y:12 P:25 SP:F9 CYC:26251
FB14  90 05     BCC $FB1B                       A:11 X:FF Y:12 P:25 SP:F9 CYC:26253
FB16  C9 11     CMP #$11                        A:11 X:FF Y:12 P:25 SP:F9 CYC:26255
FB18  D0 01     BNE $FB1B                       A:11 X:FF Y:12 P:27 SP:F9 CYC:26257
FB1A  60        RTS                             A:11 X:FF Y:12 P:27 SP:F9 CYC:26259
F755  AD 47 06  LDA $0647 = 9B                  A:11 X:FF Y:12 P:27 SP:FB CYC:26265
F758  C9 9B     CMP #$9B                        A:9B X:FF Y:12 P:A5 SP:FB CYC:26269
F75A  F0 02     BEQ $F75E                       A:9B X:FF Y:12 P:27 SP:FB CYC:26271
F75E  A0 13     LDY #$13                        A:9B X:FF Y:12 P:27 SP:FB CYC:26274
F760  A2 FF     LDX #$FF                        A:9B X:FF Y:13 P:25 SP:FB CYC:26276
F762  A9 A5     LDA #$A5                        A:9B X:FF Y:13 P:A5 SP:FB CYC:26278
F764  8D 47 06  STA $0647 = 9B                  A:A5 X:FF Y:13 P:A5 SP:FB CYC:26280
F767  20 E9 FA  JSR $FAE9                       A:A5 X:FF Y:13 P:A5 SP:FB CYC:26284
FAE9  24 01     BIT $01 = FF                    A:A5 X:FF Y:13 P:A5 SP:F9 CYC:26290
FAEB  18        CLC                             A:A5 X:FF Y:13 P:E5 SP:F9 CYC:26293
FAEC  A9 B2     LDA #$B2                        A:A5 X:FF Y:13 P:E4 SP:F9 CYC:26295
FAEE  60        RTS                             A:B2 X:FF Y:13 P:E4 SP:F9 CYC:26297
F76A  7F 48 05 *RRA $0548,X @ 0647 = A5         A:B2 X:FF Y:13 P:E4 SP:FB CYC:26303
F76D  EA        NOP                             A:5 X:FF Y:13 P:25 SP:FB CYC:26310
F76E  EA        NOP                             A:5 X:FF Y:13 P:25 SP:FB CYC:26312
F76F  EA        NOP                             A:5 X:FF Y:13 P:25 SP:FB CYC:26314
F770  EA        NOP                             A:5 X:FF Y:13 P:25 SP:FB CYC:26316
F771  20 EF FA  JSR $FAEF                       A:5 X:FF Y:13 P:25 SP:FB CYC:26318
FAEF  70 2A     BVS $FB1B                       A:5 X:FF Y:13 P:25 SP:F9 CYC:26324
FAF1  90 28     BCC $FB1B                       A:5 X:FF Y:13 P:25 SP:F9 CYC:26326
FAF3  30 26     BMI $FB1B                       A:5 X:FF Y:13 P:25 SP:F9 CYC:26328
FAF5  C9 05     CMP #$05                        A:5 X:FF Y:13 P:25 SP:F9 CYC:26330
FAF7  D0 22     BNE $FB1B                       A:5 X:FF Y:13 P:27 SP:F9 CYC:26332
FAF9  60        RTS                             A:5 X:FF Y:13 P:27 SP:F9 CYC:26334
F774  AD 47 06  LDA $0647 = 52                  A:5 X:FF Y:13 P:27 SP:FB CYC:26340
F777  C9 52     CMP #$52                        A:52 X:FF Y:13 P:25 SP:FB CYC:26344
F779  F0 02     BEQ $F77D                       A:52 X:FF Y:13 P:27 SP:FB CYC:26346
F77D  C8        INY                             A:52 X:FF Y:13 P:27 SP:FB CYC:26349
F77E  A9 29     LDA #$29                        A:52 X:FF Y:14 P:25 SP:FB CYC:26351
F780  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:14 P:25 SP:FB CYC:26353
F783  20 FA FA  JSR $FAFA                       A:29 X:FF Y:14 P:25 SP:FB CYC:26357
FAFA  B8        CLV                             A:29 X:FF Y:14 P:25 SP:F9 CYC:26363
FAFB  18        CLC                             A:29 X:FF Y:14 P:25 SP:F9 CYC:26365
FAFC  A9 42     LDA #$42                        A:29 X:FF Y:14 P:24 SP:F9 CYC:26367
FAFE  60        RTS                             A:42 X:FF Y:14 P:24 SP:F9 CYC:26369
F786  7F 48 05 *RRA $0548,X @ 0647 = 29         A:42 X:FF Y:14 P:24 SP:FB CYC:26375
F789  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB CYC:26382
F78A  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB CYC:26384
F78B  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB CYC:26386
F78C  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB CYC:26388
F78D  20 FF FA  JSR $FAFF                       A:57 X:FF Y:14 P:24 SP:FB CYC:26390
FAFF  70 1A     BVS $FB1B                       A:57 X:FF Y:14 P:24 SP:F9 CYC:26396
FB01  30 18     BMI $FB1B                       A:57 X:FF Y:14 P:24 SP:F9 CYC:26398
FB03  B0 16     BCS $FB1B                       A:57 X:FF Y:14 P:24 SP:F9 CYC:26400
FB05  C9 57     CMP #$57                        A:57 X:FF Y:14 P:24 SP:F9 CYC:26402
FB07  D0 12     BNE $FB1B                       A:57 X:FF Y:14 P:27 SP:F9 CYC:26404
FB09  60        RTS                             A:57 X:FF Y:14 P:27 SP:F9 CYC:26406
F790  AD 47 06  LDA $0647 = 14                  A:57 X:FF Y:14 P:27 SP:FB CYC:26412
F793  C9 14     CMP #$14                        A:14 X:FF Y:14 P:25 SP:FB CYC:26416
F795  F0 02     BEQ $F799                       A:14 X:FF Y:14 P:27 SP:FB CYC:26418
F799  C8        INY                             A:14 X:FF Y:14 P:27 SP:FB CYC:26421
F79A  A9 37     LDA #$37                        A:14 X:FF Y:15 P:25 SP:FB CYC:26423
F79C  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:15 P:25 SP:FB CYC:26425
F79F  20 0A FB  JSR $FB0A                       A:37 X:FF Y:15 P:25 SP:FB CYC:26429
FB0A  24 01     BIT $01 = FF                    A:37 X:FF Y:15 P:25 SP:F9 CYC:26435
FB0C  38        SEC                             A:37 X:FF Y:15 P:E5 SP:F9 CYC:26438
FB0D  A9 75     LDA #$75                        A:37 X:FF Y:15 P:E5 SP:F9 CYC:26440
FB0F  60        RTS                             A:75 X:FF Y:15 P:65 SP:F9 CYC:26442
F7A2  7F 48 05 *RRA $0548,X @ 0647 = 37         A:75 X:FF Y:15 P:65 SP:FB CYC:26448
F7A5  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB CYC:26455
F7A6  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB CYC:26457
F7A7  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB CYC:26459
F7A8  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB CYC:26461
F7A9  20 10 FB  JSR $FB10                       A:11 X:FF Y:15 P:25 SP:FB CYC:26463
FB10  70 09     BVS $FB1B                       A:11 X:FF Y:15 P:25 SP:F9 CYC:26469
FB12  30 07     BMI $FB1B                       A:11 X:FF Y:15 P:25 SP:F9 CYC:26471
FB14  90 05     BCC $FB1B                       A:11 X:FF Y:15 P:25 SP:F9 CYC:26473
FB16  C9 11     CMP #$11                        A:11 X:FF Y:15 P:25 SP:F9 CYC:26475
FB18  D0 01     BNE $FB1B                       A:11 X:FF Y:15 P:27 SP:F9 CYC:26477
FB1A  60        RTS                             A:11 X:FF Y:15 P:27 SP:F9 CYC:26479
F7AC  AD 47 06  LDA $0647 = 9B                  A:11 X:FF Y:15 P:27 SP:FB CYC:26485
F7AF  C9 9B     CMP #$9B                        A:9B X:FF Y:15 P:A5 SP:FB CYC:26489
F7B1  F0 02     BEQ $F7B5                       A:9B X:FF Y:15 P:27 SP:FB CYC:26491
F7B5  60        RTS                             A:9B X:FF Y:15 P:27 SP:FB CYC:26494
C655  A5 00     LDA $00 = 00                    A:9B X:FF Y:15 P:27 SP:FD CYC:26500
C657  05 10     ORA $10 = 00                    A:0 X:FF Y:15 P:27 SP:FD CYC:26503
C659  05 11     ORA $11 = 00                    A:0 X:FF Y:15 P:27 SP:FD CYC:26506
C65B  F0 0E     BEQ $C66B                       A:0 X:FF Y:15 P:27 SP:FD CYC:26509
C66B  20 89 C6  JSR $C689                       A:0 X:FF Y:15 P:27 SP:FD CYC:26512
C689  A9 02     LDA #$02                        A:0 X:FF Y:15 P:27 SP:FB CYC:26518
C68B  8D 15 40  STA $4015 = FF                  A:2 X:FF Y:15 P:25 SP:FB CYC:26520
C68E  A9 3F     LDA #$3F                        A:2 X:FF Y:15 P:25 SP:FB CYC:26524
C690  8D 04 40  STA $4004 = FF                  A:3F X:FF Y:15 P:25 SP:FB CYC:26526
C693  A9 9A     LDA #$9A                        A:3F X:FF Y:15 P:25 SP:FB CYC:26530
C695  8D 05 40  STA $4005 = FF                  A:9A X:FF Y:15 P:A5 SP:FB CYC:26532
C698  A9 FF     LDA #$FF                        A:9A X:FF Y:15 P:A5 SP:FB CYC:26536
C69A  8D 06 40  STA $4006 = FF                  A:FF X:FF Y:15 P:A5 SP:FB CYC:26538
C69D  A9 00     LDA #$00                        A:FF X:FF Y:15 P:A5 SP:FB CYC:26542
C69F  8D 07 40  STA $4007 = FF                  A:0 X:FF Y:15 P:27 SP:FB CYC:26544
C6A2  60        RTS                             A:0 X:FF Y:15 P:27 SP:FB CYC:26548
C66E  60        RTS                             A:0 X:FF Y:15 P:27 SP:FD CYC:26554`
var truth_arr []string = strings.Split(truth, "\n")

func run_tests(in string, j int) bool {
	truth_fields := strings.Fields(truth_arr[j])
	test_fields := strings.Fields(in)
	failed := false
	for i := 0; i < 6; i++ {
		if strings.TrimRight(truth_fields[len(truth_fields)-i-1], "0") != strings.TrimRight(test_fields[len(test_fields)-i-1], "0") {
			println("\nDIFF:\noriginal", truth_arr[j], "\nyours:", in)
			failed = true
		}
	}
	return failed
}
