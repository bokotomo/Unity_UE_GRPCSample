using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UI : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        GameObject testText = GameObject.Find("TestText");
        testText.GetComponent<Text>().text = "tomomomo";
        print("ok");
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
