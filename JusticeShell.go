package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Character struct {
	name  string
	ascii string
	color string
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Command line flags
	message := flag.String("message", "I am vengeance!", "the message to display")
	speaker := flag.String("speaker", "batman", "the DC character who's speaking (batman, joker, superman, wonder_woman)")
	respondWith := flag.String("respond", "", "optional: have another character respond (batman, joker, superman, wonder_woman)")
	flag.Parse()

	// Create the dialogue
	speakerChar := getCharacter(*speaker)
	speechBubble := createTextBubble(*message)

	// Print the first character's dialogue
	fmt.Println(speakerChar.color + speechBubble + "\033[0m")
	fmt.Println(speakerChar.color + speakerChar.ascii + "\033[0m")

	// If a responder is specified, generate and print their response
	if *respondWith != "" {
		responder := getCharacter(*respondWith)
		response := generateResponse(*respondWith, *speaker, *message)
		responseBubble := createTextBubble(response)

		fmt.Println("\n" + responder.color + responseBubble + "\033[0m")
		fmt.Println(responder.color + responder.ascii + "\033[0m")
	}
}

func createTextBubble(message string) string {
	lines := strings.Split(message, "\n")
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	border := strings.Repeat("-", maxLength+2)
	result := fmt.Sprintf(" %s\n", border)
	for _, line := range lines {
		result += fmt.Sprintf("< %s%s >\n", line, strings.Repeat(" ", maxLength-len(line)))
	}
	result += fmt.Sprintf(" %s", border)

	return result
}

func getCharacter(name string) Character {
	characters := map[string]Character{
		"batman": {
			name:  "Batman",
			color: "\033[90m", // Dark gray
			ascii: `
                         +%+                %%%                 
                       .%%%-               %%%%                 
                      -%%%%              .%%%%%                 
                     =%%%%%             =%%%%%%                 
                     %%%%%%%%%%%%%%%%%%%%%%%%%%                 
                   %%%%%%%%%%%%%%%%%%%%%%%%%%%%                 
                 %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                 
               =%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                 
              -%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%:                
              %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%+                
              %%%%%%%%%%%%%%%%%: :%%%%%%%%%%%%%=                
              * #%%%%%%%%*.    -%%%%%%%%%%%%%%%#                
             %%+ %%%%%%%:    =%%%%%%%%%%%%%%%%%%                
            -%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                
            .-.:%%%%%%*::...=:::::...%%%%%%%%%%%:               
             -....-...........#.....-%%%%%%%%%%%=               
             =...:....:+#*:....*....-%%%%%%%%%%%+               
             #..................:....%%%%%%%%%%%+               
             +.......................%%%%%%%%%%%*               
             +.......................%%%%%%%%%%%%               
             +......................-%%%%%%%%%%%%               
             :................:=*%%%%%%%%%%%%%%%%*              
                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%              
                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%              
                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%             
                  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%.           
                 %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%=::       
             .#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%:       
          %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%                                                                                      
`,
		},
		"joker": {
			name:  "Joker",
			color: "\033[35m", // Purple
			ascii: `
                               @@@@@@@@@%#***%@             
                                     @@@@@@@%*++%@          
                               @@@@@@@@@@@@@@@@#*+*%        
         @@  @@         @@@@@@@@@@@@@@@@@@@@@#++++++*%      
          @%@@@@@@@@@@@@@@@@@@@@@@@%#########%*+#@@*++#     
           %**#%@@@@@@@%**++++++++++++++*%@@%%%%%@@@#+*     
            @*+%%*+++++++*@@@@@@@@@@@%#*+++%@@@@@@@@%*%     
             @*+#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%#%      
              @#+#@@@@@@@@%-           =@@@@@@@@@@#@        
               @%+#@@@@@@@+              :@@@@@@@@          
                @#+#@@@@@@@%+.            #%+=:.=           
              @=   *@@@@@@@+                     *          
              @. :- +@@@@*                       :@         
              @: -:* @@#.  .:                     *         
               * -:  .       =.:+-.     .         =@        
               @=   -      ... -=   :-.   :*+ .=+#          
                @*:.       --::           :===.-#           
                 @=+    : .:.             ::.. :%           
                 %=*       . .#-    -.     :::.#            
                 =.*       .:   +*:  ..    -. #             
                %: -=      ..     .:--:=:  .=@              
                +   +.     -:    .   . .:=.=@               
               %.   .=.     +  :.::..: =: -%                
              @%=    :+.    :#*    - .  . *                 
              %%##*:  :=-    :##: :    +:+                  
             *=%%##*#*+--=-    -#+#**##*=@                  
            #=*#%%%#***#%@+=.           %                   
           #++**##%%%****#%%%#.        -@                   
          @+=****##%%%#***#%%%%#-      #                    
         @+=*******#%%%##**##%%%%%*.  -*#@@                 
        @+=*********#%%%%#***#%%%%*--=*++*#%@               
       @*=***********##%%%%****#%@=-#=**-----#@@            
      @*=*************##%%%%#**#%+=-*+=-++=---##@@          
      #=****************#%%%%#*#%=+-==-===++=-=##%@         
     @++*****************##%%%##*--*=+---+-+=+==+##@@
`,
		},
		"superman": {
			name:  "Superman",
			color: "\033[34m", // Blue
			ascii: `
         %%                                                                                   
            *=:  :..-                                                                               
          *-.  :  .  :+                                                                             
          #: .  :. .   : .+              ########%                                                  
           #-         ..  +              ###########                                                
            #:    .. .     *            #+.**++-   ##                                               
             +      .     -             #+ :-:  .  =*+                       #**++**#               
              #++:      .+#   **+==+*#  +-          .:                    #*+========+*#            
               #*+-    :++++#*++++====*==-           :                  #+============++*%          
               #*++++++++++++***+++*+==+==.  ..     .:**++*#*+#      %*+=======++*#                 
               ##++++++++++++++++++++*+=+=-         . =====+==+***##*+===++*#%                      
                #*++++++++++++++++++++++++=-       .   :-==+==*++++*++++++*#                        
                 %##*+++++++++++++++++++++=: ....    .  :==+=*+++++**+++++++++**#                   
                     ##*+++++++++++++++++*=:            :=+++++++++**+++++++++++**                  
                        ##*++++++++++++++++++=-::   .   -++++++++***+++:  ..:++++*#                 
                            ##***++++++++++++++++++++****+++++++*++=    .- .:++++*#                 
                            ##+++++*++++++++========+++++++++++*+++: .. :..:++++*#                  
                             %*+++++*+++++=--==-:.  ..:==-=+++*****++=+++++++++*#                   
                               #*++++*+++++=============-=++*#*+++++*********##                     
                                 #*+++**++++++==::-..:=++*++##++*##                                 
                                   ##**++++++++**++=++*++***+++++++**#                              
                                       #**+++++++++++=--=#*+++++++++++*#                            
                                          #+:..+*. :+*####++++++++++++++#                           
                                           %###########***#*+++++++++++++#                          
                                             #**+++++++++++*##*++++++++++*%                         
                                                ##*+++++++++*# #**+++++++*#                         
                                                      #***++***#*++*+++++*#                         
                                                       #*****+===**==++**#                          
                                                               *+=+*+==+#                           
                                                                *+=+*#*+*                           
                                                                   #%               
`,
		},
		"wonder_woman": {
			name:  "Wonder Woman",
			color: "\033[31m", // Red
			ascii: `               
                             @@@@@@@@@@@@@@@@                         
                          @@@@@@@@@@@@@@@@@@@@@@@                     
                        @@@@@@@@@@@@@@@@@@@@@@@@@@@@                  
                      @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@                
                     @@@@@@@@@@@@@@@@@@@@@@@#-@@@@@@@@@@              
                    @@@@@@@@@@@@@:::::::::::#-:@@@@@@@@@@@            
                   @@@@@@@@@@@@@::=@%:::::####+:+@@@@@@@@@@           
                   @@@@@@@*::+:::::::+@:::::*-::::@@@@@@@             
                  @@@@@@@@=:::::@@@:::==:::::=@@:@@@@@@               
                  @@@@@@@@@-::::::*::#=:::::=:-:=@@@@@                
                  @@@@@@@@--@::::::::::::-#@=::@@@@@@@                
                  @@@@@@@@@-@+:::-:::::::::::::::@@@@@                
                  @@@@@@@@@@@@+::::=:::::::::::@@@@@@@@               
                  @@@@@@@@@@@@@@=::::-:::::::+@:=@@@@@@@@             
                 @@@@@@@@@@@@@@@@++::::::@@@@@@@@@@@@@@@@@@           
                @@@@@@@@@@@@@@@@@+===+=*@@@@@@@@@@@@@@@@@@@@          
             @@@@@@@@@@@@@@@@@@@@+=====+%+=:::::::@@@@@@@@@@          
            @@@@@@@@@@@@%:::-============::::::::::+@@@@@@@@          
          @@@@@@@@@@@@@::::::::-===++++-=-::::::::::@@@@@@@@          
          @@@@@@@@@@@@:::::::::::::::::::::::::::-::=@@@@@@           
          @@@@@@@@@@@@:::::::::::::::::::::-::####=:=@@@@@            
          @@@@@@@@@@@@:::::::::::+####***##:::::::=:=@@@              
          @@@@@@@@@@@@-::::::::+@#-:::::::::::==::::=@                
           @@@@@@@@@@@@::::::::%+::::::::::::::::::::                 
             @@@@@@@@@@-:::::::%@%:::::::#*=:######:#                 
               @@@@@@@@@:::::::%@@@@@##############:@                 
                    @@@@@::::::@@@@@%#############+=                                        
`,
		},
	}

	if char, exists := characters[strings.ToLower(name)]; exists {
		return char
	}
	return characters["batman"] // Default to Batman if character not found
}

func generateResponse(responder, speaker, originalMessage string) string {
	// Map of possible responses for each character
	responses := map[string][]string{
		"batman": {
			"I am the night.",
			"Gotham needs me.",
			"Justice must be served.",
			"I'm not wearing hockey pads.",
		},
		"joker": {
			"HAHAHAHAHAHA!",
			"Why so serious?",
			"Let's put a smile on that face!",
			"This town deserves a better class of criminal.",
		},
		"superman": {
			"Truth, justice, and the American way!",
			"Up, up and away!",
			"This looks like a job for Superman!",
			"I'm here to help.",
		},
		"wonder_woman": {
			"Great Hera!",
			"We have to fight for those who cannot fight for themselves.",
			"I am Diana of Themyscira!",
			"Peace is always worth fighting for.",
		},
	}

	if responseList, exists := responses[strings.ToLower(responder)]; exists {
		return responseList[rand.Intn(len(responseList))]
	}
	return "..."
}
