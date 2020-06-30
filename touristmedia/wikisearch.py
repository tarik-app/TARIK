import wikipedia 

def wiki_search(site):
    result = wikipedia.search(site) 
    result = ", ".join(result)
    return result

print(wiki_search("The painted ladies"))
