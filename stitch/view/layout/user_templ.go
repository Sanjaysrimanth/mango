// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func User() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<link rel=\"stylesheet\" href=\"/static/styles/user/template.css\"><link rel=\"stylesheet\" href=\"/static/styles/user/search.css\"><script src=\"/static/scripts/scroll_animation.js\" defer></script><body><div class=\"overflow\"><header class=\"surface\" style=\"position: relative;\"><div class=\"logo\"><input type=\"checkbox\" id=\"searchCheckbox\"> <label for=\"searchCheckbox\"><img src=\"/static/icons/search.svg\" alt=\"Search\"></label><div id=\"searchContainer\" class=\"searchContainer\"><form hx-get=\"/products\" hx-target=\"body\" hx-push-url=\"true\"><input type=\"text\" placeholder=\"Search...\" list=\"queries\" name=\"search\"> <datalist id=\"queries\"><option value=\"shirt\"></option> <option value=\"t-shirt\"></option> <option value=\"pants\"></option></datalist> <input type=\"submit\" class=\"button-primary\" value=\"Search\"></form><label for=\"searchCheckbox\"><img src=\"/static/icons/close.svg\" alt=\"Close\"></label></div><h1>STICTH</h1><a href=\"/cart\" hx-boost=\"true\"><img src=\"/static/icons/cart.svg\"></a></div><ul><li><a href=\"/\" hx-boost=\"true\">Home</a></li><li><a href=\"/products\" hx-boost=\"true\">All</a></li><li><a href=\"/collections\" hx-boost=\"true\">Collections</a></li><li><a href=\"/about\" hx-boost=\"true\">About us</a></li></ul></header><main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</main><footer class=\"surface\"><div><div class=\"policies\"><h3>Policies</h3><ul><li><a href=\"mailto:mail@gmail.com\">Contact us</a></li></ul></div><div class=\"logo\"><h1>STICTH</h1></div><div class=\"contacts\"><h3>Contact us</h3><ul><li>+91 8353692784</li></ul></div></div></footer></div></body>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
