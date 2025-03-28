#version 330

in vec2 fragTexCoord;
out vec4 finalColor;

uniform sampler2D texture0;

void main()
{
    vec4 texColor = texture(texture0, fragTexCoord);

    if (texColor.a != 0.0)
    {
        finalColor = vec4(1.0, 1.0, 1.0, 1.0);
    }
    else
    {
        finalColor = texColor;
    }
}