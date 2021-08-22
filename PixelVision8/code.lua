level=1
start_y=104
mode="title"
wind=30.5
m = MousePosition()

function solid(x,y)
 return solids[Tile((x)/8,(y)/8).SpriteId]
end

function spring(x,y)
 return springs[Tile((x)/8,(y)/8).SpriteId]
end

function Init()
 solids={
  [1]=true,
  [2]=true,
  [3]=true,
  [5]=true,
  [6]=true,
  [7]=true,
  [8]=true,
  [11]=true,
  [12]=true,
  [14]=true,
  [15]=true,
  [16]=true,
  [17]=true,
  [18]=true
 }
 springs={
  [3]=true,
  [7]=true,
  [8]=true,
  [11]=true,
  [14]=true,
  [18]=true
 }
  p={
  x=16,
  y=start_y,
  vx=0,
  vy=0,
  s=0.35,
  sprite=19
 }
  mc=0
  sp=3
  f=0
end  

local ms=MetaSprite("player")

LoadTilemap("title-screen")
PlaySound(12) 
 
function Update(deltaTime)

 if mode=="title" then
  
  if Button(4,1) then
   mode="game"
   LoadTilemap("level-1")
   PlaySound(9)
  end 
   
 elseif mode=="ending" then
 
  if Button(4,1) then
   level=1
   start_y=104
   wind=30.5
   Init()
   mode="title"
   LoadTilemap("title-screen")   
   PlaySound(12)
  end
   
 elseif mode=="pause" then
 
  if Button(4,1) then
   mode="game"
   PlaySound(8)
  end
   
 elseif mode=="game" then
  p.s=0.35+(level-1)/15

  if p.vy==0 then
   p.vx=p.s
  else
   if jumping==true then
    p.vx=p.s*1.75
   else
    p.vx=0
   end
  end
  
  if solid(p.x+p.vx,p.y+p.vy) or solid(p.x+7+p.vx,p.y+p.vy) or solid(p.x+p.vx,p.y+7+p.vy) or solid(p.x+7+p.vx,p.y+7+p.vy) then
   p.vx=0
  end
  
  if solid(p.x,p.y+8+p.vy) or solid(p.x+7,p.y+8+p.vy) then
   p.vy=0
   jumping=false
  else
   p.vy=p.vy+0.2
  end
  
  if p.vy==0 and Tile(p.x/8,p.y/8+1).SpriteId==2 then 
   p.vy=-3.5 
   jumping=true
   PlaySound(4)
  end

  if p.vy<0 and (solid(p.x+p.vx,p.y+p.vy) or solid(p.x+7+p.vx,p.y+p.vy)) then
   p.vy=0
  end
  
  if Tile(p.x/8,p.y/8+1).SpriteId==9 then
   Init()
   LoadTilemap("level-" .. level)
   PlaySound(6)
  end
  
  if Tile(p.x/8,p.y/8).SpriteId==13 then
   if level==8 then
    mode="ending"
    LoadTilemap("end-screen")
    PlaySound(11,2)
   else
    p.x=16
    start_y=p.y
    level=level+1
    sp=3
    LoadTilemap("level-" .. level)
    PlaySound(7)
   end
  end
  
  p.x=p.x+p.vx
  p.y=p.y+p.vy


  if MouseButton(0) then
   if mc==0 then
    if spring(m.x,m.y) and sp>0 then
     Tile(m.x/8,m.y/8,2)
     sp=sp-1
     PlaySound(5)
    elseif Tile(m.x/8,m.y/8).SpriteId==2 and (p.vy>0 or jumping==false) then
     Tile(m.x/8,m.y/8,14)
     sp=sp+1
     PlaySound(5)
    end
   end
    mc=mc+1
   else
   mc=0
  end
  
  if Button(4,1) then
   mode="pause"
   PlaySound(8)
  end
  
  if sp>=3 then sp=3 end
  if mc>=1 then mc=1 end
  
  if p.vy==0 then
   p.sprite=19
   f=(f+p.s/5)%4
  elseif p.vy>0 then
   p.sprite=43
   f=0
  elseif p.vy<0 then
   p.sprite=44
   f=0     
  end
  
  wind=wind+0.1
  if wind>=30.5 then
   PlaySound(10,2)
   wind=0
  end
   
 end
 
 m=MousePosition()
 
end

function Draw()

  RedrawDisplay()

  if mode=="title" then
 
    DrawText("A Tiny Challenge",8*8,12*8+4,DrawMode.Sprite,"large",4)
    DrawText("PRESS A/X TO START",7*8,17*8,DrawMode.Sprite,"large",15) 
    DrawText("Mega Sparkmaster",10*8,21*8,DrawMode.Sprite,"medium",5,-2)  
    
   elseif mode=="ending" then
    
    DrawText("CONGRATULATIONS!",8*8+4,8*8,DrawMode.Sprite,"large",14)
    DrawText("You have finally",7*8+4,10*8,DrawMode.Sprite,"large",15)
    DrawText("reached the exit...",7*8+4,11*8,DrawMode.Sprite,"large",15)
    DrawText("Thanks for playing!",9*8+4,13*8-2,DrawMode.Sprite,"medium",14,-2)   
    DrawText("Press A/X to return",9*8+4,14*8-2,DrawMode.Sprite,"medium",11,-2)         
    
   elseif mode=="pause" then
    
    DrawSprite(p.sprite+f,p.x,p.y)
    DrawText("AREA "..level,13*8,16,DrawMode.Sprite,"large",15)
    DrawText("SPRINGS: "..sp,11*8,27*8,DrawMode.Sprite,"large",11)
    DrawText("PAUSE",13*8+5,15*8,DrawMode.Sprite,"large",15)            
    
   elseif mode=="game" then

    DrawSprite(41,math.floor(m.x/8)*8,math.floor(m.y/8)*8)
    DrawSprite(p.sprite+f,p.x,p.y)
    DrawText("AREA "..level,13*8,16,DrawMode.Sprite,"large",15)
    DrawText("SPRINGS: "..sp,11*8,27*8,DrawMode.Sprite,"large",11)
   
  end

  if m.x>0 and m.x<256 and m.y>0 and m.y<240 then
    DrawMetaSprite("mouse",m.x,m.y, false, false, DrawMode.SpriteAbove)
  end

end
