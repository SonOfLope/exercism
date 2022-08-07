export function hey(message: unknown): unknown {
  if (typeof message === "string") {
    if(message.trim().slice(-1) === "?"){
      if(message.toUpperCase() === message && message.toLowerCase() !== message) return "Calm down, I know what I'm doing!";
      return "Sure.";
    } 
    else if(message.toUpperCase() === message && message.toLowerCase() !== message) return "Whoa, chill out!";
    else if(message.trim() === "") return "Fine. Be that way!";
    else return "Whatever."; 
  }
}