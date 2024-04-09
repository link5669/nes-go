Reset:
  jsr _func_8606       
  jsr _func_8638                 

_func_8606:
  jsr _func_860d              
  jsr _func_862a                 
  rts                            

_func_860d:
  lda #$02                       
  sta $02           
  lda #$06                       
  sta $03
  lda #$11                       
  sta $10        
  lda #$10                       
  sta $12        
  lda #$0F                       
  sta $14                      
  lda #$04                       
  sta $11        
  sta $13                      
  sta $15                      
  rts                            

_func_862a:
  lda $FE                
  sta $00        
  lda $FE                
  and #$03                       
  clc                            
  adc #$02                       
  sta $01                
  rts                            

_func_8638:
  jsr _func_864d         
  jsr _func_868d            
  jsr _func_86c3    
  jsr _func_8719   
  jsr _func_8720        
  jsr _func_872d                 
  jmp _func_8638                 

_func_864d:
  lda $FF                      
  cmp #$77                       
  beq _label_8660                
  cmp #$64                       
  beq _label_866b                
  cmp #$73                       
  beq _label_8676                
  cmp #$61                       
  beq _label_8681                
  rts                            

_label_8660:
  lda #$04                       
  bit $02                
  bne _label_868c                
  lda #$01                       
  sta $02                
  rts                            

_label_866b:
  lda #$08                       
  bit $02                
  bne _label_868c                
  lda #$02                       
  sta $02                
  rts                            

_label_8676:
  lda #$01                       
  bit $02                
  bne _label_868c                
  lda #$04                       
  sta $02                
  rts                            

_label_8681:
  lda #$02                       
  bit $02                
  bne _label_868c                
  lda #$08                       
  sta $02                
  rts                            

_label_868c:
  rts                            

_func_868d:
  jsr _func_8694             
  jsr _func_86a8                 
  rts                            

_func_8694:
  lda $00        
  cmp $10        
  bne _label_86a7                
  lda $01                
  cmp $11        
  bne _label_86a7                
  inc $03                
  inc $03                
  jsr _func_862a                 

_label_86a7:
  rts                            

_func_86a8:
  ldx #$02                       

_label_86aa:
  lda $10,X      
  cmp $10        
  bne _label_86b6                
  lda $11,X      
  cmp $11        
  beq _label_86bf                

_label_86b6:
  inx                            
  inx                            
  cpx $03                
  beq _label_86c2                
  jmp _label_86aa                

_label_86bf:
  jmp _label_8735                

_label_86c2:
  rts                            

_func_86c3:
  ldx $03                
  dex                    
  txa                            

_label_86c7:
  lda $10,X      
  sta $12,X   
  dex                            
  bpl _label_86c7      
  lda $02                
  lsr                     
  bcs _label_86dc     
  lsr                   
  bcs _label_86ef                
  lsr                          
  bcs _label_86f8                
  lsr                          
  bcs _label_870b                

_label_86dc:
  lda $10        
  sec                            
  sbc #$20                       
  sta $10        
  bcc _label_86e6                
  rts                            

_label_86e6:
  dec $11        
  lda #$01                       
  cmp $11        
  beq _label_8716                
  rts                            

_label_86ef:
  inc $10     
  lda #$1F                       
  bit $10        
  beq _label_8716           
  rts                            

_label_86f8:
  lda $10        
  clc                            
  adc #$20                       
  sta $10      
  bcs _label_8702                
  rts                            

_label_8702:
  inc $11        
  lda #$06                       
  cmp $11        
  beq _label_8716                
  rts                            

_label_870b:
  dec $10        
  lda $10        
  and #$1F                       
  cmp #$1F                       
  beq _label_8716                
  rts                            

_label_8716:
  jmp _label_8735                

_func_8719:
  ldy #$00                
  lda $FE            
  sta ($00),Y      
  rts                            

_func_8720:
  ldx $03                
  lda #$00                       
  sta ($10,X)      
  ldx #$00                       
  lda #$01                       
  sta ($10,X)      
  rts                            

_func_872d:
  ldx #$00                       

_label_872f:
  nop                            
  nop                            
  dex                            
  bne _label_872f                
  rts                            

_label_8735:
  brk