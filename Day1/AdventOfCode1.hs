import Data.List (intersect)

type Env = (Integer,Integer,Integer,Integer)

data Direction = NORTH | SOUTH | EAST | WEST
-- Advent of code #1
go :: Integer
go = step NORTH start theMap

start :: Env
start = (0,0,0,0)

step :: Direction -> Env -> [String] -> Integer
step _ (n,w,s,e) []           = abs (n-s) + abs (w-e)
step d env@(n,w,s,e) (ss:sss) = step newDirection (updateEnv newDirection env (read (tail ss))) sss
                            
  where
    updateEnv :: Direction -> Env -> Integer -> Env
    updateEnv NORTH (n,w,s,e) number = (n+number,w,s,e)
    updateEnv WEST (n,w,s,e) number = (n,w+number,s,e)
    updateEnv EAST (n,w,s,e) number = (n,w,s,e+number)
    updateEnv SOUTH (n,w,s,e) number = (n,w,s+number,e)

    newDirection :: Direction
    newDirection = updateDirection d $ head ss

    updateDirection :: Direction -> Char -> Direction
    updateDirection NORTH 'R' = EAST
    updateDirection NORTH 'L' = WEST
    updateDirection EAST 'R' = SOUTH
    updateDirection EAST 'L' = NORTH
    updateDirection WEST 'R' = NORTH
    updateDirection WEST 'L' = SOUTH
    updateDirection SOUTH 'R' = WEST
    updateDirection SOUTH 'L' = EAST

theMap :: [String]
theMap = words $ intersect "R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, R2, R3, R5, L2, L3, R3, R1, R3" "RL1234567890 "