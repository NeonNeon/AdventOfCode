import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Stream;

/**
 * Created by neon on 2016-12-02.
 */
public class Day2 {
    private boolean[][] problem1 = {{true,true,true},{true,true,true},{true,true,true}};
    private boolean[][] problem2 = {{false,false,true,false,false},{false,true,true,true,false},{true,true,true,true,true},{false,true,true,true,false},{false,false,true,false,false}};
    private MapProblem firstProblem;

    public Day2() {
        firstProblem = new MapProblem(problem2,0,2);
        try (BufferedReader br = new BufferedReader(new FileReader("/home/neon/Documents/AdventOfCode/Day2/src/input2"))) {
            String line;
            while ((line = br.readLine()) != null) {
               while (line.length() >= 1) {
                   firstProblem.move(getInstruction(line.charAt(0)));
                   if (line.length()>=1)
                       line = line.substring(1,line.length());
               }
                System.out.println(firstProblem.getPosition());
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

    }
    public Instruction getInstruction(char letter) {
        switch (letter) {
            case 'U':
                return Instruction.UP;
            case 'D':
                return Instruction.DOWN;
            case 'L':
                return Instruction.LEFT;
            case 'R':
                return Instruction.RIGHT;
        }
        return null;
    }

    public enum Instruction {
        DOWN,LEFT,RIGHT,UP;
    }
    private class MapProblem {
        private boolean[][] structure;
        private int xPosition, yPosition;
        public MapProblem(boolean[][] structure, int xStart, int yStart) {
            this.structure = structure;
            this.xPosition = xStart;
            this.yPosition = yStart;
        }
        public String getPosition() {
            return xPosition + "," + yPosition;
        }
        public void move(Instruction i) {
            switch (i) {
                case DOWN:
                    if (yPosition < structure[0].length - 1 && structure[yPosition + 1][xPosition])
                        yPosition++;
                    break;
                case UP:
                    if (yPosition > 0 && structure[yPosition - 1][xPosition])
                        yPosition--;
                    break;
                case LEFT:
                    if (xPosition > 0 && structure[yPosition][xPosition - 1])
                        xPosition--;
                    break;
                case RIGHT:
                    if (xPosition < structure.length - 1 && structure[yPosition][xPosition + 1])
                        xPosition++;
                    break;
                default:
                    break;
            }
        }
    }
    public static void main(String args[]) {
        Day2 go = new Day2();
    }
}
